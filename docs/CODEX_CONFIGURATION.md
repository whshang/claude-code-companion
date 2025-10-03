# Codex 配置指南

## 概述

Codex CLI 是一个基于 Anthropic API 的 AI 编程助手工具。CCCC 项目为其提供统一的代理转发服务。

## 重要发现（2025-10-02）

**Codex 使用 Anthropic API，而不是 OpenAI API！**

这意味着 Codex 和 Claude Code 共享相同的 API 基础设施，可以通过同一个代理服务器统一管理。

## 正确配置

### 环境变量

```bash
export ANTHROPIC_BASE_URL="http://127.0.0.1:8080"
```

**注意**：
- ✅ 使用 `ANTHROPIC_BASE_URL`
- ❌ 不要使用 `OPENAI_API_BASE` 或 `OPENAI_API_KEY`

### 配置文件

**路径**：`~/.codex/config.toml`

**格式**：JSON（虽然扩展名是 `.toml`）

**结构**：
```json
{
  "env": {
    "ANTHROPIC_BASE_URL": "http://127.0.0.1:8080"
  },
  "hooks": {},
  "permissions": {
    "allow": [],
    "deny": []
  }
}
```

### 默认端口

- Codex 默认代理端口：**8080**
- CCCC 代理服务器可配置端口（示例：8081）

## 自动配置脚本

CCCC Web UI（`/help` 页面）提供三种自动生成的配置脚本：

### 1. Codex 单独配置脚本

**文件名**：
- Windows: `cccc-codex.bat`
- macOS: `cccc-codex.command`
- Linux: `cccc-codex.sh`

**功能**：
- 设置 `ANTHROPIC_BASE_URL` 环境变量
- 创建或更新 `~/.codex/config.toml`
- 自动备份原有配置
- 启动 Codex 并传递命令行参数

**使用方法**：
```bash
# macOS/Linux
chmod +x cccc-codex.sh
./cccc-codex.sh

# Windows
cccc-codex.bat
```

### 2. 一键配置双客户端脚本

**文件名**：`cccc-setup.{bat,command,sh}`

**功能**：同时配置 Claude Code 和 Codex

**输出示例**：
```
========================================
 CCCC Setup - macOS
========================================

Configuring Claude Code...
----------------------------
Claude Code configured successfully

Configuring Codex...
--------------------
Backed up Codex config to: ~/.codex/config.toml.backup-2025-10-02T...
Codex configured successfully

========================================
 Setup Complete!
========================================

Claude Code: Use 'claude' command as usual
Codex: Use 'codex' command as usual

Both clients are now configured to use CCCC proxy at http://localhost:8081
```

## 配置对比

### Claude Code vs Codex

| 项目 | Claude Code | Codex |
|------|------------|-------|
| **环境变量** | `ANTHROPIC_BASE_URL` | `ANTHROPIC_BASE_URL` |
| **额外变量** | `ANTHROPIC_AUTH_TOKEN` 等 | 无 |
| **配置文件** | `~/.claude/settings.json` | `~/.codex/config.toml` |
| **文件格式** | JSON | JSON（扩展名 .toml） |
| **配置字段** | `env` | `env`, `hooks`, `permissions` |

## 手动配置

如果不使用自动脚本，可以手动配置：

### 步骤 1：创建配置目录

```bash
mkdir -p ~/.codex
```

### 步骤 2：创建配置文件

```bash
cat > ~/.codex/config.toml << 'EOF'
{
  "env": {
    "ANTHROPIC_BASE_URL": "http://127.0.0.1:8081"
  },
  "hooks": {},
  "permissions": {
    "allow": [],
    "deny": []
  }
}
EOF
```

### 步骤 3：设置环境变量（可选）

在 `~/.bashrc` 或 `~/.zshrc` 中添加：

```bash
export ANTHROPIC_BASE_URL="http://127.0.0.1:8081"
```

### 步骤 4：验证配置

```bash
# 检查配置文件
cat ~/.codex/config.toml

# 检查环境变量
echo $ANTHROPIC_BASE_URL

# 测试 Codex
codex "帮我写一个 Hello World"
```

## 常见问题

### Q1: 为什么配置文件是 .toml 但内容是 JSON？

A: 这是 Codex 的设计决定。虽然文件扩展名是 `.toml`，但实际内容使用 JSON 格式。

### Q2: 需要设置 API Key 吗？

A: 不需要。Codex 通过 CCCC 代理服务器连接，API Key 由代理服务器管理。

### Q3: 配置不生效怎么办？

1. 检查代理服务器是否运行
2. 验证 `ANTHROPIC_BASE_URL` 指向正确的地址
3. 确认配置文件格式正确（有效的 JSON）
4. 尝试重新运行配置脚本

### Q4: 如何恢复旧配置？

配置脚本会自动创建备份文件：
```bash
ls ~/.codex/config.toml.backup-*
# 选择需要的备份恢复
cp ~/.codex/config.toml.backup-2025-10-02T... ~/.codex/config.toml
```

### Q5: 多个代理服务器如何切换？

修改 `config.toml` 中的 `ANTHROPIC_BASE_URL`：
```json
{
  "env": {
    "ANTHROPIC_BASE_URL": "http://other-proxy:8082"
  },
  ...
}
```

## 故障排查

### 问题：Codex 连接失败

**检查步骤**：

1. **验证代理服务器运行**：
   ```bash
   curl http://127.0.0.1:8081/health
   ```

2. **检查配置文件**：
   ```bash
   cat ~/.codex/config.toml | jq .
   # 应该能成功解析 JSON
   ```

3. **测试环境变量**：
   ```bash
   echo $ANTHROPIC_BASE_URL
   ```

4. **查看 Codex 日志**：
   ```bash
   codex --debug "test"
   ```

### 问题：Node.js 不可用

配置脚本需要 Node.js 来更新配置文件。如果没有安装：

**方案 A**：安装 Node.js
```bash
# macOS
brew install node

# Ubuntu/Debian
sudo apt install nodejs

# 或使用 nvm
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
nvm install node
```

**方案 B**：手动配置（见上方"手动配置"章节）

## 安全建议

1. **不要提交配置文件到版本控制**
   - `.codex/` 目录已在 `.gitignore` 中
   
2. **使用本地代理**
   - 推荐使用 `127.0.0.1` 或 `localhost`
   - 避免将配置暴露到公网

3. **定期备份配置**
   ```bash
   cp ~/.codex/config.toml ~/.codex/config.toml.backup-$(date +%Y%m%d)
   ```

4. **审查权限设置**
   ```json
   {
     "permissions": {
       "allow": ["read", "write"],
       "deny": ["network", "system"]
     }
   }
   ```

## 进阶配置

### Hooks 配置

```json
{
  "hooks": {
    "pre_request": "echo 'Sending request...'",
    "post_response": "echo 'Response received'"
  }
}
```

### 权限控制

```json
{
  "permissions": {
    "allow": [
      "read:./src/**",
      "write:./output/**"
    ],
    "deny": [
      "read:./secrets/**",
      "write:/etc/**"
    ]
  }
}
```

## 相关文档

- [SCRIPT_NAMING_CONVENTION.md](SCRIPT_NAMING_CONVENTION.md) - 脚本命名规范
- [README.md](../README.md) - 项目概述
- [CHANGELOG.md](../CHANGELOG.md) - 版本历史

## 更新历史

- **2025-10-02**: 修正配置信息，确认 Codex 使用 Anthropic API
- **2025-10-02**: 添加一键配置双客户端功能
- **2025-10-02**: 创建本文档（合并 CODEX_CONFIG_FIX_v2.md 和 CODEX_FIX_SUMMARY.md）

