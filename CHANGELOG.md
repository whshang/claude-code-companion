# Changelog

All notable changes to Claude Code Companion will be documented in this file.

## [2.1.0] - 2025-10-03

### 📖 Fixed - 文档错误修正

**README.md 修正**：
- ✅ 修复 `/help` 链接格式错误（缺少空格导致渲染问题）
- ✅ 修正 Codex 配置文件路径：`~/.codex/config.json` → `~/.codex/config.toml`
- ✅ 更新 Codex 配置格式为正确的 TOML 格式
- ✅ 添加详细的配置字段说明（model_provider, wire_api, requires_openai_auth 等）
- ✅ 移除环境变量配置方式（Codex 不支持）
- ✅ 移除重复的"一键生成配置"章节

**正确的 Codex 配置示例**：
```toml
model_provider = "cccc"
model = "gpt-5"

[model_providers.cccc]
name = "cccc"
base_url = "http://127.0.0.1:8080"
wire_api = "responses"
requires_openai_auth = true

[projects."/path/to/your/project"]
trust_level = "trusted"
```

### 🧠 Added - 智能参数学习与自动重试系统

**核心功能：零配置端点适配**
- **智能参数学习机制**
  - 从 400 错误响应中自动识别不支持的参数
  - 支持多种错误格式的解析（关键词匹配 + 正则提取）
  - 学习 `tools`、`tool_choice`、`functions`、`function_call` 等参数
  - 每个端点独立维护学习到的参数列表

- **即时自动重试策略**
  - 学习新参数后立即移除并重试当前请求
  - 递归调用 `proxyToEndpoint` 确保请求成功
  - 避免端点因参数不兼容被误判为故障
  - 详细日志记录学习和重试过程

- **线程安全的参数管理**
  - `sync.RWMutex` 保护参数列表并发访问
  - 实时参数检查方法 `IsParamUnsupported()`
  - 参数列表获取方法 `GetLearnedUnsupportedParams()`
  - 参数学习方法 `LearnUnsupportedParam()`

**技术实现细节**：
```go
// internal/proxy/proxy_logic.go:476-500
// 🎓 自动学习不支持的参数 - 基于400错误分析并重试
if resp.StatusCode == 400 {
    paramCountBefore := len(ep.GetLearnedUnsupportedParams())
    s.learnUnsupportedParamsFromError(decompressedBody, ep, finalRequestBody)
    paramCountAfter := len(ep.GetLearnedUnsupportedParams())
    if paramCountAfter > paramCountBefore {
        // 学到新参数，立即清理并重试
        cleanedBody, wasModified := s.autoRemoveUnsupportedParams(finalRequestBody, ep)
        if wasModified {
            return s.proxyToEndpoint(c, ep, path, cleanedBody, ...)
        }
    }
}
```

**用户收益**：
- ✅ 无需手动配置参数白名单
- ✅ 自动适配各端点的参数支持差异
- ✅ 首次 400 错误后自动修正并重试
- ✅ 避免端点被错误黑名单
- ✅ 提高请求成功率和系统稳定性

### 🔧 Fixed - Codex Native Support 检测修正

**问题背景**：
- Deepseek 端点在 `/chat/completions` 成功后被标记为 `native_codex_support = true`
- 导致后续 `/responses` 请求跳过转换，直接转发到端点
- 端点返回 404（不支持 `/responses` 路径），请求失败
- 端点被误判为故障并加入黑名单

**修复方案**：
```go
// internal/proxy/proxy_logic.go:818-827
// 只有当 /responses 路径成功时，才标记端点支持原生 Codex 格式
// /chat/completions 成功不代表支持 /responses
if formatDetection != nil && formatDetection.ClientType == utils.ClientCodex &&
   ep.EndpointType == "openai" {
    if inboundPath == "/responses" {
        s.updateEndpointCodexSupport(ep, true)
    }
}
```

**影响文件**：
- `internal/proxy/proxy_logic.go:818-827`

**用户收益**：
- ✅ 修复 Codex 请求 502 Bad Gateway 错误
- ✅ 正确识别端点的 Codex 格式支持能力
- ✅ 避免端点被错误黑名单

### 🛠️ Enhanced - 参数学习智能解析

**新增功能** (`internal/proxy/proxy_logic.go:1414-1507`):
```go
func (s *Server) learnUnsupportedParamsFromError(errorBody []byte, ep *endpoint.Endpoint, requestBody []byte) {
    // 1. 解析错误消息
    // 2. 关键词匹配（unsupported, not supported, invalid parameter）
    // 3. 正则提取参数名
    // 4. 验证请求体中是否包含该参数
    // 5. 学习并记录到端点
}
```

**支持的错误格式**：
- OpenAI 标准格式：`{"error": {"message": "...", "type": "..."}}`
- 简化格式：`{"message": "..."}`
- 纯文本格式：直接解析错误文本

**提取策略**：
- 关键词 + 参数名匹配：`"tools" parameter is not supported`
- 引号包裹提取：`'tool_choice' is invalid`
- 请求体验证：确保参数确实存在于请求中

### 🌍 Refactor - 国际化支持优化

**已实现语言**（9种）：
- 中文 (zh)
- 英文 (en)
- 日语 (ja)
- 韩语 (ko)
- 德语 (de)
- 西班牙语 (es)
- 法语 (fr)
- 意大利语 (it)
- 葡萄牙语 (pt)
- 俄语 (ru)

**技术改进**：
- Web UI 完整国际化
- 脚本生成器多语言支持
- 配置助手页面翻译
- 错误消息本地化

### 📊 Changed - 日志系统增强

**新增日志信息**：
- 参数学习事件记录
  ```
  [INFO] Learned new unsupported parameters, retrying with clean request
  endpoint: Deepseek-codex
  learned_count: 2
  ```
- 自动移除参数日志
  ```
  [DEBUG] Auto-removing unsupported parameters from request
  endpoint: Deepseek-codex
  removed: [tools, tool_choice]
  ```
- 重试请求追踪
  ```
  [DEBUG] Retrying request after removing learned unsupported parameters
  ```

**日志优化**：
- 更清晰的参数学习流程追踪
- 详细的重试原因说明
- 端点学习状态实时记录

### 🔍 Technical Details

**修改文件统计**：
```
94 files changed
4654 insertions(+)
851 deletions(-)
Net: +3803 lines
```

**核心变更文件**：
1. `internal/proxy/proxy_logic.go`
   - Line 3-22: 添加 regexp 导入
   - Line 476-500: 400错误学习和重试逻辑
   - Line 818-827: Native Codex 支持检测修正
   - Line 1414-1507: `learnUnsupportedParamsFromError()` 函数

2. `internal/endpoint/endpoint.go`
   - Line 83-91: 新增学习参数字段和互斥锁
   - Line 688-724: 参数管理方法实现

3. `internal/utils/endpoint_sorter.go`
   - Line 82-90: 标签过滤逻辑修正

### 📖 Documentation Updates

**新增/更新文档**：
- `README.md`: 新增智能参数学习系统介绍
- `CHANGELOG.md`: 完整的版本变更记录
- `CLAUDE.md`: Claude Code 项目说明
- `docs/`: 完整的技术文档目录

### ⚠️ Breaking Changes

无破坏性变更，完全向后兼容。

### 🚀 Migration Guide

从 v2.0.x 升级到 v2.1.0：

1. **无需配置变更**
   - 智能学习系统自动启用
   - 无需修改 `config.yaml`
   - 现有端点配置保持不变

2. **数据库自动升级**
   - 首次启动自动添加新字段
   - 无需手动迁移
   - 建议备份：`cp logs/logs.db logs/logs.db.backup`

3. **测试验证**
   ```bash
   # 重新编译
   go build -o cccc

   # 启动服务
   ./cccc -config config.yaml -port 8080

   # 观察日志确认学习系统工作
   tail -f logs/proxy.log | grep "Learned"
   ```

### 🎯 Performance Impact

**性能提升**：
- 首次 400 错误后立即修正，减少后续失败请求
- 避免端点被误判黑名单，提高可用端点数量
- 学习结果持久化（端点生命周期内），避免重复试错

**资源消耗**：
- 内存增量：每端点 ~100 bytes（参数列表）
- CPU 影响：可忽略（仅在 400 错误时触发）
- 网络影响：首次失败 + 1次重试，后续零额外请求

### 🐛 Known Issues

1. **参数学习限制**
   - 当前仅支持请求体级别参数（如 tools, tool_choice）
   - 不支持嵌套参数的学习
   - 未来版本计划支持更复杂的参数结构

2. **错误格式兼容性**
   - 依赖端点返回标准化的错误消息
   - 非标准错误格式可能无法正确提取参数名
   - 建议使用符合 OpenAI 规范的端点

### 🙏 Acknowledgments

感谢社区反馈的端点兼容性问题，推动了本次智能学习系统的开发。

---

## [2.0.1] - 2025-10-02

### Fixed - Codex 配置修正
- **修正 Codex 配置信息**
  - 确认 Codex 使用 `ANTHROPIC_BASE_URL` 而非 `OPENAI_API_BASE`
  - 配置文件路径修正为 `~/.codex/config.toml`（JSON 格式内容）
  - 配置结构包含 `env`, `hooks`, `permissions` 字段
  - 更新所有脚本生成逻辑（单独配置 + 一键配置）
  - 详见 [docs/CODEX_CONFIGURATION.md](docs/CODEX_CONFIGURATION.md)

### Changed - 文档结构优化
- **整理项目文档结构**
  - 根目录仅保留：README.md、CHANGELOG.md、CLAUDE.md、AGENTS.md
  - 所有技术文档移至 `docs/` 目录
  - 合并重复文档，移除临时文件
  - 新增 [docs/CODEX_CONFIGURATION.md](docs/CODEX_CONFIGURATION.md) - Codex 配置完整指南
  - 更新 [docs/SCRIPT_NAMING_CONVENTION.md](docs/SCRIPT_NAMING_CONVENTION.md) - 脚本命名规范

## [2.0.0] - 2025-10-02

### Added - 项目重命名与功能增强
- **项目更名为 CCCC (Claude Code and Codex Companion)**
  - 统一支持 Claude Code 和 Codex CLI 工具
  - 更新 README.md 和 Web UI 反映新定位
  - 致敬原项目 [kxn/claude-code-companion](https://github.com/kxn/claude-code-companion)
  
- **脚本命名统一规范**
  - Claude Code 脚本：`cccc-claude.{bat,sh,command}`
  - Codex 脚本：`cccc-codex.{bat,sh,command}`
  - 一键配置脚本：`cccc-setup.{bat,sh,command}`
  - 所有脚本统一使用 `cccc-` 前缀
  
- **一键配置双客户端功能**
  - Web UI 新增脚本类型选择器（单独配置 / 一键配置）
  - 支持同时配置 Claude Code 和 Codex
  - 自动更新两个客户端的配置文件
  - 智能备份原有配置
  - 详细的配置进度显示

### Enhanced - Web UI 改进
- **配置助手页面 (/help) 增强**
  - 客户端选择卡片（Claude Code / Codex）
  - 脚本类型选择器（单独配置 / 一键配置双客户端）
  - 动态文件名显示
  - 智能 UI 切换（选择一键配置时自动隐藏客户端选择）
  - URL 参数支持直接跳转到 Codex 配置（`?client=codex`）

## [Unreleased] - 2025-10-02

### Added - Codex 客户端支持
- **客户端类型检测系统**
  - 自动识别 Claude Code 和 Codex 客户端
  - 基于请求路径和请求体结构的智能检测
  - 支持客户端特定的端点路由
  
- **Codex 格式转换**
  - Codex `/responses` 格式自动转换为 OpenAI `/chat/completions` 格式
  - `instructions` 字段转换为 system 消息
  - `input` 数组转换为标准 messages 格式
  - 保留 `tools`、`tool_choice` 等 OpenAI 兼容字段
  
- **客户端特定端点配置**
  - `supported_clients` 字段支持端点级别的客户端过滤
  - 支持值：`claude-code`、`codex`、或留空表示支持所有客户端
  
- **增强的模型重写功能**
  - 支持通用端点的隐式模型重写
  - Claude Code 客户端：非 Claude 模型自动转为 claude-sonnet-4-20250514
  - Codex 客户端：非 GPT 模型自动转为 gpt-5
  - 保持与显式配置规则的兼容性

### Fixed - 端点健康检查和响应验证
- **SSE 流验证修复**
  - 修复 OpenAI SSE 流的 `[DONE]` 标记检测
  - 为 Codex 客户端自动添加缺失的 `[DONE]` 标记
  - 改进 finish_reason 检测逻辑
  
- **响应验证改进**
  - 严格验证 Anthropic 和 OpenAI 响应格式
  - 支持端点特定的验证规则（如 apis.iflow.cn 白名单）
  - 更好的错误消息和诊断信息
  
- **端点黑名单机制**
  - 自动黑名单连续失败的端点
  - 记录导致黑名单的请求 ID
  - 支持通用端点的回退机制

### Changed - 架构优化
- **端点选择逻辑重构**
  - 统一的端点选择接口（`EndpointSorter`）
  - 支持格式、客户端类型、标签的多维度筛选
  - 优化的优先级排序算法
  
- **请求处理流程优化**
  - 格式检测 → 客户端识别 → 端点选择 → 格式转换 → 模型重写
  - 每个步骤都有详细的日志记录
  - 支持请求重试和端点回退
  
- **日志系统增强**
  - 新增 `client_type`、`request_format`、`target_format` 字段
  - 记录格式转换状态和检测置信度
  - 改进的调试信息输出

### Technical Details

#### 文件结构变化
```
internal/
├── utils/
│   └── format_detector.go          # 新增：格式和客户端类型检测
├── modelrewrite/
│   └── rewriter.go                 # 增强：支持客户端类型感知的重写
├── endpoint/
│   ├── selector.go                 # 重构：多维度端点选择
│   └── manager.go                  # 增强：客户端类型过滤支持
├── proxy/
│   ├── proxy_logic.go              # 重大改进：Codex 转换和 SSE 修复
│   └── request_processing.go       # 新增：请求处理辅助函数
└── validator/
    └── response.go                 # 增强：严格的格式验证
```

#### 配置文件变化
```yaml
endpoints:
  - name: example-endpoint
    url: https://api.example.com
    endpoint_type: openai
    # 新增字段
    supported_clients:              # 可选：支持的客户端类型
      - codex
      - claude-code
    model_rewrite:
      enabled: true
      rules:
        - source_pattern: gpt-5*
          target_model: actual-model
```

#### 数据库 Schema 变化
```sql
-- request_logs 表新增字段
ALTER TABLE request_logs ADD COLUMN client_type TEXT DEFAULT '';
ALTER TABLE request_logs ADD COLUMN request_format TEXT DEFAULT '';
ALTER TABLE request_logs ADD COLUMN target_format TEXT DEFAULT '';
ALTER TABLE request_logs ADD COLUMN format_converted BOOLEAN DEFAULT false;
ALTER TABLE request_logs ADD COLUMN detection_confidence REAL DEFAULT 0;
ALTER TABLE request_logs ADD COLUMN detected_by TEXT DEFAULT '';
```

### Known Issues

1. **Codex 端点兼容性**
   - foxcode (code.newcli.com): SSE 流不完整，缺少 `[DONE]` 标记
   - 88code (88code.org): 部分请求返回 400 错误
   - 需要使用完全兼容 OpenAI 标准的端点

2. **响应验证严格性**
   - 当前验证器要求每个 SSE chunk 都包含 `id` 字段
   - 部分非标准端点可能被误判为无效
   - 建议：配置可靠的端点或调整验证规则

3. **工具调用支持**
   - Codex 的 `tools` 字段现已保留并正确传递
   - 但部分端点可能不完全支持 OpenAI 工具调用规范

### Migration Guide

#### 从旧版本升级

1. **配置文件迁移**
   ```bash
   # 备份现有配置
   cp config.yaml config.yaml.backup
   
   # 添加客户端支持（可选）
   # 在需要的端点下添加 supported_clients 字段
   ```

2. **数据库升级**
   ```bash
   # 数据库会自动迁移，但建议备份
   cp logs/logs.db logs/logs.db.backup
   ```

3. **端点配置检查**
   - 确认 Codex 端点配置了正确的 `endpoint_type: openai`
   - 验证模型重写规则是否符合新的客户端类型逻辑
   - 测试端点是否返回符合 OpenAI 规范的响应

#### 配置示例

**支持 Codex 的完整配置**:
```yaml
endpoints:
  - name: openai-compatible-endpoint
    url: https://api.example.com
    endpoint_type: openai
    path_prefix: "/v1"
    auth_type: auth_token
    auth_value: your-token-here
    enabled: true
    priority: 1
    supported_clients:
      - codex
    model_rewrite:
      enabled: true
      rules:
        - source_pattern: gpt-5*
          target_model: your-actual-model
```

### Performance Improvements
- 格式检测使用缓存机制，避免重复计算
- 优化端点选择算法，减少不必要的遍历
- 改进日志批量写入性能

### Security
- 确保测试脚本不包含在版本控制中
- 敏感配置使用示例文件（config.yaml.example）
- 日志中的认证信息自动脱敏

---

## [Previous Versions]

### Initial Release
- 基础的 Anthropic API 代理功能
- 多端点支持和健康检查
- 请求/响应日志记录
- Web 管理界面
- 标签路由系统

---

## Future Plans

### Planned Features
- [ ] 更多 AI 提供商支持（Google Gemini、Azure OpenAI）
- [ ] 请求限流和配额管理
- [ ] 实时性能监控和告警
- [ ] 端点自动健康评分
- [ ] 请求缓存支持
- [ ] WebSocket 支持

### Under Consideration
- GraphQL API 支持
- 请求变换规则引擎
- 多租户支持
- 插件系统

---

## Contributing
欢迎贡献代码！请查看 README.md 了解开发指南。

## License
See LICENSE file for details.

