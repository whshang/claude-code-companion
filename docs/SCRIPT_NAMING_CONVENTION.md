# 脚本命名规范与使用指南

## 脚本命名体系

CCCC (Claude Code and Codex Companion) 提供统一、规范的脚本命名体系，支持灵活配置单个或多个AI编程助手客户端。

## 脚本类型

### 1. Claude Code 专用脚本

用于配置和启动 Claude Code CLI 工具：

- **Windows**: `cccc-claude.bat`
- **macOS**: `cccc-claude.command`
- **Linux**: `cccc-claude.sh`

**功能**：
- 设置 Claude Code 所需的环境变量
- 更新 `~/.claude/settings.json` 配置文件
- 自动备份原有配置
- 启动 Claude Code 并传递所有命令行参数

**使用示例**：
```bash
# Windows
cccc-claude.bat

# macOS/Linux
./cccc-claude.sh
```

### 2. Codex 专用脚本

用于配置和启动 Codex CLI 工具：

- **Windows**: `cccc-codex.bat`
- **macOS**: `cccc-codex.command`
- **Linux**: `cccc-codex.sh`

**功能**：
- 设置 Codex 所需的环境变量 (`OPENAI_API_BASE`, `OPENAI_API_KEY`)
- 更新 `~/.codex/config.json` 配置文件（如有）
- 启动 Codex 并传递所有命令行参数

**使用示例**：
```bash
# Windows
cccc-codex.bat

# macOS/Linux
./cccc-codex.sh
```

### 3. 一键配置双客户端脚本

同时配置 Claude Code 和 Codex：

- **Windows**: `cccc-setup.bat`
- **macOS**: `cccc-setup.command`
- **Linux**: `cccc-setup.sh`

**功能**：
- 一次性配置两个客户端的所有设置
- 修改配置文件（`settings.json` 和 `config.json`）
- 设置持久化环境变量
- 自动备份原有配置
- 显示详细的配置进度和结果

**使用示例**：
```bash
# Windows
cccc-setup.bat

# macOS/Linux
./cccc-setup.sh
```

**输出示例**：
```
========================================
 CCCC Setup - macOS
========================================

Configuring Claude Code...
----------------------------
Backed up Claude settings to: /Users/xxx/.claude/settings.json.backup-2025-10-02T...
Claude Code configured successfully

Configuring Codex...
--------------------
Backed up Codex config to: /Users/xxx/.codex/config.json.backup-2025-10-02T...
Codex configured successfully

========================================
 Setup Complete!
========================================

Claude Code: Use 'claude' command as usual
Codex: Use 'codex' command as usual

Both clients are now configured to use CCCC proxy at http://localhost:8081
```

## 命名规范说明

### 前缀：`cccc-`
所有脚本统一使用 `cccc-` 前缀，代表 "Claude Code and Codex Companion"，便于识别和管理。

### 中间标识：
- `claude` - Claude Code 专用
- `codex` - Codex 专用
- `setup` - 一键配置双客户端

### 扩展名：
- `.bat` - Windows 批处理文件
- `.command` - macOS 可执行脚本（双击运行）
- `.sh` - Linux Shell 脚本

## Web UI 使用指南

在 CCCC 的 Web 管理界面（`/help` 页面），用户可以：

### 1. 选择脚本类型
- **单独配置**：生成单个客户端的启动脚本
- **一键配置双客户端**：生成同时配置两个客户端的设置脚本

### 2. 选择客户端（仅单独配置模式）
- **Claude Code**：生成 Claude Code 启动脚本
- **Codex**：生成 Codex 启动脚本

### 3. 选择操作系统
- Windows
- macOS
- Linux

### 4. 下载脚本
点击"下载脚本"按钮，自动生成并下载对应的配置脚本文件。

## 配置流程

### 方案 A：单独配置（适合只使用一个客户端）

1. 访问 CCCC Web UI 的 `/help` 页面
2. 选择"单独配置"
3. 选择要配置的客户端（Claude Code 或 Codex）
4. 选择操作系统
5. 下载并运行脚本
6. 直接使用 `claude` 或 `codex` 命令

### 方案 B：一键配置（适合同时使用两个客户端）

1. 访问 CCCC Web UI 的 `/help` 页面
2. 选择"一键配置双客户端"
3. 选择操作系统
4. 下载并运行 `cccc-setup` 脚本
5. 配置完成后，两个客户端都可以正常使用

## 技术细节

### Claude Code 配置

**环境变量**：
- `ANTHROPIC_BASE_URL`: CCCC 代理服务器地址
- `ANTHROPIC_AUTH_TOKEN`: 认证令牌（默认为 "hello"）
- `CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC`: 禁用非必要流量
- `API_TIMEOUT_MS`: API 超时时间（默认 600000ms）

**配置文件**：`~/.claude/settings.json`

### Codex 配置

**环境变量**：
- `ANTHROPIC_BASE_URL`: CCCC 代理服务器地址

**配置文件**：`~/.codex/config.toml`（JSON 格式内容）

**配置结构**：
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

**注意**：虽然文件扩展名为 `.toml`，但内容使用 JSON 格式。

## 安全注意事项

1. **不要提交脚本到版本控制**：所有 `cccc-*.bat/sh/command` 和 `test_*.sh` 文件已被添加到 `.gitignore`
2. **配置文件自动备份**：脚本会在修改配置文件前自动创建备份
3. **敏感信息保护**：默认使用占位符（如 "hello"），实际部署时由代理服务器处理真实凭据

## 故障排查

### 脚本无法执行（macOS/Linux）
```bash
chmod +x cccc-claude.sh
./cccc-claude.sh
```

### Node.js 未安装
脚本中的配置文件更新功能需要 Node.js。如果未安装，环境变量仍会被设置，但配置文件不会自动更新。可以手动编辑配置文件或安装 Node.js 后重新运行脚本。

### 配置未生效
1. 检查代理服务器是否正常运行
2. 确认脚本中的 `baseUrl` 是否正确
3. 查看脚本运行时的输出信息
4. 尝试运行一键配置脚本重新设置

## 更新日志

### v2.0 (2025-10-02)
- ✅ 统一脚本命名规范（`cccc-` 前缀）
- ✅ 新增一键配置双客户端功能
- ✅ Web UI 增强：脚本类型选择器
- ✅ 动态文件名显示
- ✅ 改进配置脚本错误处理

## 相关文档

- [README.md](README.md) - 项目概述
- [CHANGELOG.md](CHANGELOG.md) - 完整更新历史
- [PRE_COMMIT_CHECKLIST.md](PRE_COMMIT_CHECKLIST.md) - 提交前检查清单

