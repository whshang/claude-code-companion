# 配置文件指南

## 📋 目录

- [快速开始](#快速开始)
- [端点配置](#端点配置)
- [高级功能](#高级功能)
- [常见问题](#常见问题)

## 🚀 快速开始

### 1. 基础配置结构

```yaml
server:
    host: 127.0.0.1    # 监听地址
    port: 8081         # 监听端口

endpoints:
    - name: my-endpoint
      url: https://api.example.com
      endpoint_type: openai  # 或 anthropic
      auth_type: auth_token
      auth_value: your-api-key
      enabled: true
      priority: 1
```

### 2. 当前启用的端点

- **88code-codex** (Priority 1) - Codex 专用
- **88code-cc** (Priority 2) - Claude Code 专用

## 🔧 端点配置

### 基础字段

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `name` | string | ✅ | 端点唯一标识符 |
| `url` | string | ✅ | 端点基础 URL（不带 `/v1`，见下方说明） |
| `endpoint_type` | string | ✅ | `openai` 或 `anthropic` |
| `auth_type` | string | ✅ | `auth_token` 或 `api_key` |
| `auth_value` | string | ✅ | API 密钥 |
| `enabled` | boolean | ✅ | 是否启用 |
| `priority` | integer | ✅ | 优先级（数字越小越高） |

### URL 配置原则

#### ✅ 推荐配置（简洁）

```yaml
# 大多数情况：不带 /v1
url: https://api.example.com

# 代理会自动拼接：
# - https://api.example.com/responses
# - https://api.example.com/chat/completions
```

#### 🔄 特殊情况

```yaml
# 仅当服务端明确要求完整路径时
url: https://www.88code.org/openai/v1

# 或使用 path_prefix 字段
url: https://api.example.com
path_prefix: /v1
```

### 可选字段

#### `supported_clients`
限制支持的客户端类型：

```yaml
supported_clients:
  - codex          # 仅支持 Codex
  - claude-code    # 仅支持 Claude Code
# 省略此字段 = 支持所有客户端
```

#### `model_rewrite`
模型名称重写：

```yaml
model_rewrite:
  enabled: true
  rules:
    - source_pattern: gpt-5*
      target_model: qwen3-coder
    - source_pattern: claude-*sonnet*
      target_model: kimi-k2
```

#### `path_prefix`
OpenAI 端点路径前缀：

```yaml
path_prefix: /v1
# 最终 URL: {url}{path_prefix}{request_path}
```

## 🎯 端点分组说明

### 🔥 主力端点
正在使用的主要服务，`enabled: true`

### 🔄 备用端点
故障转移或负载均衡用途，按需启用

### 🤖 国产大模型
Kimi、Deepseek、豆包等国产服务

### 🔧 自建/测试端点
个人服务器和实验性配置

### 🧪 测试账号
用于测试的临时账号

## 🔐 认证类型

### `auth_token`
发送 `Authorization: Bearer {token}`

```yaml
auth_type: auth_token
auth_value: sk-xxxxx
```

### `api_key`
发送 `x-api-key: {key}`

```yaml
auth_type: api_key
auth_value: sk-xxxxx
```

## 📊 优先级规则

1. **数字越小优先级越高**
   - Priority 1 = 最高优先级
   - Priority 18 = 最低优先级

2. **端点选择逻辑**
   - 按优先级排序
   - 过滤 `enabled: true`
   - 过滤健康状态
   - 匹配 `supported_clients`
   - 选择第一个可用端点

## 🛠️ 高级功能

### 1. 自动格式转换

代理自动处理以下转换：
- Codex `/responses` → OpenAI `/chat/completions`
- OpenAI ↔ Anthropic 格式互转
- 自动探测端点支持的格式

### 2. 健康检查

```yaml
timeouts:
    health_check_timeout: 30s  # 健康检查超时
    check_interval: 30s        # 检查间隔
    recovery_threshold: 0      # 恢复阈值
```

### 3. 日志配置

```yaml
logging:
    level: debug                    # debug, info, warn, error
    log_request_types: failed       # all, failed, none
    log_request_body: truncated     # full, truncated, none
    log_response_body: truncated
    log_directory: ./logs
```

## 📝 配置示例

### 示例 1: 标准 Codex 端点

```yaml
- name: my-codex-service
  url: https://api.example.com
  endpoint_type: openai
  auth_type: auth_token
  auth_value: sk-xxxxx
  enabled: true
  priority: 5
  supported_clients:
    - codex
```

### 示例 2: 带模型重写的端点

```yaml
- name: kimi-with-rewrite
  url: https://api.moonshot.cn/v1
  endpoint_type: openai
  auth_type: auth_token
  auth_value: sk-xxxxx
  enabled: true
  priority: 6
  model_rewrite:
    enabled: true
    rules:
      - source_pattern: gpt-*
        target_model: kimi-k2-0905-preview
```

### 示例 3: 多用途端点（无客户端限制）

```yaml
- name: universal-endpoint
  url: https://api.example.com
  endpoint_type: anthropic
  auth_type: auth_token
  auth_value: sk-xxxxx
  enabled: true
  priority: 10
  # 省略 supported_clients = 支持所有客户端
```

## ❓ 常见问题

### Q1: 如何添加新端点？

1. 复制现有端点配置
2. 修改 `name`, `url`, `auth_value`
3. 设置 `enabled: false` 先测试
4. 调整 `priority` 确定优先级
5. 重启代理服务

### Q2: URL 应该带 `/v1` 吗？

**推荐不带**，除非：
- 服务端明确要求完整路径
- 服务端 URL 本身就包含版本号（如 88code）

### Q3: 如何测试新端点？

```bash
# 使用测试脚本
./test-new-endpoint.sh your-endpoint-name

# 或直接 curl
curl -X POST "http://127.0.0.1:8081/responses" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer test" \
  -d '{"model":"gpt-5-codex",...}'
```

### Q4: 端点被 blacklist 了怎么办？

1. 检查日志: `tail -f logs/proxy.log`
2. 确认端点健康: 访问 `http://127.0.0.1:8081/admin/`
3. 重启服务清除 blacklist
4. 或等待自动恢复（recovery_threshold）

### Q5: 如何启用备用端点？

1. 编辑 `config.yaml`
2. 设置 `enabled: true`
3. 调整 `priority`（比主端点数字大）
4. 重启服务

### Q6: 模型重写何时生效？

模型重写在以下情况生效：
- 匹配 `source_pattern`（支持通配符 `*`）
- `enabled: true`
- 请求发送到该端点之前自动应用

## 🔗 相关文档

- [项目 README](../README.md)
- [88code 配置教程](../88code配置教程.md)
- [Codex 配置指南](./CODEX_CONFIGURATION.md)

## 💡 配置提示

1. ✅ **始终保留备用端点** - 故障转移很重要
2. ✅ **使用 `supported_clients` 分流** - 避免无效尝试
3. ✅ **优先级间隔留空间** - 方便后续插入新端点
4. ✅ **测试账号 `enabled: false`** - 避免意外消耗配额
5. ✅ **定期检查日志** - 及时发现问题

## 📊 推荐优先级分配

| 优先级范围 | 用途 |
|------------|------|
| 1-5 | 主力生产端点 |
| 6-10 | 国产大模型和备用端点 |
| 11-15 | 自建和测试服务 |
| 16-20 | 临时测试账号 |
