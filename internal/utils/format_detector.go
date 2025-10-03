package utils

import (
	"encoding/json"
	"strings"
	"sync"
)

// RequestFormat represents the detected API format
type RequestFormat string

const (
	FormatAnthropic RequestFormat = "anthropic"
	FormatOpenAI    RequestFormat = "openai"
	FormatUnknown   RequestFormat = "unknown"
)

// ClientType represents the detected client type
type ClientType string

const (
	ClientClaudeCode ClientType = "claude-code"
	ClientCodex      ClientType = "codex"
	ClientUnknown    ClientType = "unknown"
)

// FormatDetectionResult contains the result of format detection
type FormatDetectionResult struct {
	Format      RequestFormat
	ClientType  ClientType
	Confidence  float64 // 0.0 - 1.0
	DetectedBy  string  // detection method used
}

// 简单的路径检测缓存，避免重复计算
var (
	pathDetectionCache = make(map[string]*FormatDetectionResult)
	cacheMutex         sync.RWMutex
	cacheMaxSize       = 1000 // 限制缓存大小，避免内存泄漏
)

// getCachedPathDetection 从缓存获取路径检测结果
func getCachedPathDetection(path string) (*FormatDetectionResult, bool) {
	cacheMutex.RLock()
	defer cacheMutex.RUnlock()
	result, exists := pathDetectionCache[path]
	return result, exists
}

// setCachedPathDetection 设置路径检测结果到缓存
func setCachedPathDetection(path string, result *FormatDetectionResult) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	// 简单的缓存淘汰策略：超过最大值时清空缓存
	if len(pathDetectionCache) >= cacheMaxSize {
		pathDetectionCache = make(map[string]*FormatDetectionResult)
	}

	pathDetectionCache[path] = result
}

// DetectRequestFormat automatically detects the API format from request path and body
func DetectRequestFormat(path string, requestBody []byte) *FormatDetectionResult {
	// 1. 先尝试从缓存获取路径检测结果
	if cached, exists := getCachedPathDetection(path); exists {
		return cached
	}
	result := &FormatDetectionResult{
		Format:     FormatUnknown,
		ClientType: ClientUnknown,
		Confidence: 0.0,
	}

	// 1. Path-based detection (highest confidence)
	// 使用更精确的路径匹配，避免误判

	// Anthropic API paths - 精确匹配端点路径
	if strings.HasSuffix(path, "/messages") || strings.HasSuffix(path, "/v1/messages") ||
		strings.HasSuffix(path, "/count_tokens") || strings.HasSuffix(path, "/v1/count_tokens") {
		result.Format = FormatAnthropic
		result.ClientType = ClientClaudeCode
		result.Confidence = 0.95
		result.DetectedBy = "path"
		setCachedPathDetection(path, result) // 缓存路径检测结果
		return result
	}

	// OpenAI API paths - 精确匹配端点路径（包含常见和新增路由）
	openaiPaths := []string{
		"/chat/completions",
		"/v1/chat/completions",
		"/completions",
		"/v1/completions",
		"/embeddings",
		"/v1/embeddings",
		"/models",
		"/v1/models",
		"/images/generations",
		"/v1/images/generations",
		"/audio/transcriptions",
		"/v1/audio/transcriptions",
		"/audio/translations",
		"/v1/audio/translations",
		"/audio/speech",
		"/v1/audio/speech",
		"/files",
		"/v1/files",
		"/fine_tuning",
		"/v1/fine_tuning",
		"/batches",
		"/v1/batches",
		"/responses",  // 新增：OpenAI responses API
		"/v1/responses",
		"/realtime",   // 新增：实时 API
		"/v1/realtime",
	}

	for _, openaiPath := range openaiPaths {
		if strings.HasSuffix(path, openaiPath) || strings.Contains(path, openaiPath+"/") {
			result.Format = FormatOpenAI
			result.ClientType = ClientCodex
			result.Confidence = 0.95
			result.DetectedBy = "path"
			setCachedPathDetection(path, result) // 缓存路径检测结果
			return result
		}
	}

	// 2. Body structure detection (medium confidence)
	if len(requestBody) > 0 {
		var reqData map[string]interface{}
		if err := json.Unmarshal(requestBody, &reqData); err == nil {
			bodyResult := detectFromBody(reqData)
			if bodyResult.Confidence > 0.3 { // 只有足够信心时才使用
				return bodyResult
			}
		}
	}

	// 3. 无法确定格式时保持 unknown，避免误判
	// 让后续代码根据端点类型决定是否需要转换
	result.Format = FormatUnknown
	result.ClientType = ClientUnknown
	result.Confidence = 0.0
	result.DetectedBy = "unknown"
	return result
}

// detectFromBody detects format from request body structure
func detectFromBody(reqData map[string]interface{}) *FormatDetectionResult {
	result := &FormatDetectionResult{
		Format:     FormatUnknown,
		ClientType: ClientUnknown,
		Confidence: 0.0,
	}

	anthropicScore := 0.0
	openAIScore := 0.0

	// Anthropic format characteristics
	if _, hasSystem := reqData["system"]; hasSystem {
		anthropicScore += 0.3
	}

	if _, hasMaxTokens := reqData["max_tokens"]; hasMaxTokens {
		anthropicScore += 0.1
	}

	// Check for Anthropic-specific fields
	if _, hasThinking := reqData["thinking"]; hasThinking {
		anthropicScore += 0.2
	}

	// OpenAI format characteristics
	if messages, ok := reqData["messages"].([]interface{}); ok && len(messages) > 0 {
		if msg, ok := messages[0].(map[string]interface{}); ok {
			if role, ok := msg["role"].(string); ok {
				if role == "system" || role == "developer" {
					// OpenAI 格式的 system 消息在 messages 数组内
					openAIScore += 0.3
				} else if role == "user" || role == "assistant" {
					// Both formats can have user/assistant messages
					// Check for OpenAI-specific message structure
					if _, hasContent := msg["content"]; hasContent {
						openAIScore += 0.1
						anthropicScore += 0.1
					}
				}
			}
		}
	}

	// Codex-specific format detection (instructions field)
	// Codex 使用 instructions 字段代替 messages 数组
	if instructions, hasInstructions := reqData["instructions"]; hasInstructions {
		if _, ok := instructions.(string); ok {
			// 这是 Codex 特有的格式，需要转换为标准 OpenAI 格式
			// 注意：虽然是 OpenAI 兼容格式，但需要格式转换
			openAIScore += 0.5 // 高分表示是 OpenAI 格式家族
			result.Format = FormatOpenAI // Codex 是 OpenAI 的变体
			result.ClientType = ClientCodex
			result.Confidence = 0.95
			result.DetectedBy = "codex-instructions"
			return result // 立即返回，确保优先识别 Codex 格式
		}
	}

	// OpenAI-specific fields
	if _, hasMaxCompletionTokens := reqData["max_completion_tokens"]; hasMaxCompletionTokens {
		openAIScore += 0.2
	}

	if _, hasTopP := reqData["top_p"]; hasTopP {
		openAIScore += 0.1
		anthropicScore += 0.1 // Both support this
	}

	if _, hasFrequencyPenalty := reqData["frequency_penalty"]; hasFrequencyPenalty {
		openAIScore += 0.2 // OpenAI-specific
	}

	if _, hasPresencePenalty := reqData["presence_penalty"]; hasPresencePenalty {
		openAIScore += 0.2 // OpenAI-specific
	}

	// Determine format based on scores
	if anthropicScore > openAIScore && anthropicScore > 0.3 {
		result.Format = FormatAnthropic
		result.ClientType = ClientClaudeCode
		result.Confidence = anthropicScore
		result.DetectedBy = "body-structure"
	} else if openAIScore > anthropicScore && openAIScore > 0.3 {
		result.Format = FormatOpenAI
		result.ClientType = ClientCodex
		result.Confidence = openAIScore
		result.DetectedBy = "body-structure"
	}

	return result
}

// GetClientTypeName returns a human-readable client type name
func (c ClientType) String() string {
	switch c {
	case ClientClaudeCode:
		return "Claude Code"
	case ClientCodex:
		return "Codex"
	default:
		return "Unknown"
	}
}

// GetFormatName returns a human-readable format name
func (f RequestFormat) String() string {
	switch f {
	case FormatAnthropic:
		return "Anthropic"
	case FormatOpenAI:
		return "OpenAI"
	default:
		return "Unknown"
	}
}