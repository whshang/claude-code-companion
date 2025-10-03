[图片]
迄今为止最先进的代码助手： 
- Claude Code是为编写代码而生的Agent，您可以像与人对话一样，使用自然语言轻松实现的想法，无需任何学习成本。
- 只需等待一杯咖啡的时间，Claude Code就可以为您解决所有问题。
- CodeX已上线！！！口碑飙升，额度和cc通用

支持的 IDE
- Visual Studio Code（包括 Cursor 和 Windsurf 等流行分支）
- JetBrains IDEs（包括 PyCharm、WebStorm、IntelliJ 和 GoLand）
一、ClaudeCode国内镜像（稳定性/性价比/口碑之王🏆）
🌏网址：https://foxcode.hshwk.org/
（注册后使用兑换码激活）
✅ 纯正Max号池搭建：拒绝第三方掺假，确保每次调用都是不降智的官方品质。
🚫 无需魔法：国内外直连，秒速响应，稳定不封号
📊 （独家）缓存透明计费：每一笔记录都清晰可查，不乱扣费（支持任意第三方检测）
📊 （独家）独家技术：行业内命中缓存率超高，减少token调用（我们的token更耐用）
🚀 （独家）兼容官网各种报错（我们更少的报错）
🚀 （独家）支持在RooCode/Kilo Code/ CherryStudio/ ChatBox中使用（下面教程三）
🚀 支持镜像包方式安装：一键登录授权，无需修改配置文件，自动检测官方版本 并一键升级
🚀 支持官方包方式安装：配置环境变量
操作系统: macOS 10.15+ / Ubuntu 20.04+/Debian 10+ / Windows
硬件: 最少 4GB RAM
软件: Node.js 18+

1.1、❤️官方包，安装方式（二选一）
(如果使用上面镜像安装教程就不需要用此安装方式了)
安装官方 Claude Code
npm install -g @anthropic-ai/claude-code
claude --version
以下是 Windows、macOS 和 Linux 系统下设置 ANTHROPIC_BASE_URL 和 ANTHROPIC_API_KEY 环境变量的详细方法：
Windows 系统
方法1（永久设置）：配置settings.json
- 创建 ~/.claude/settings.json 文件，内容如下：
{
  "env": {
    "ANTHROPIC_AUTH_TOKEN": "替换为您的API Key",
    "ANTHROPIC_BASE_URL": "https://code.newcli.com/claude",
    "CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC": 1
  },
  "permissions": {
    "allow": [],
    "deny": []
  }
}
方法2：临时设置（仅当前终端有效）
- 在 PowerShell 或 CMD 中执行：
# PowerShell
$env:ANTHROPIC_BASE_URL="https://code.newcli.com/claude"
$env:ANTHROPIC_API_KEY="替换为您的API Key"

# CMD
set ANTHROPIC_BASE_URL=https://code.newcli.com/claude
set ANTHROPIC_API_KEY=替换为您的API Key
方法3：永久设置（全局生效）
1. 图形界面：
  - 右键「此电脑」→「属性」→「高级系统设置」→「环境变量」
  - 在「用户变量」或「系统变量」中新建：
    - 变量名：ANTHROPIC_BASE_URL
    - 变量值：https://code.newcli.com/claude
  - 同样方法添加 ANTHROPIC_API_KEY
    
2. PowerShell 永久设置：
[System.Environment]::SetEnvironmentVariable('ANTHROPIC_BASE_URL', 'https://code.newcli.com/claude', 'User')
[System.Environment]::SetEnvironmentVariable('ANTHROPIC_API_KEY', '替换为您的API Key', 'User')
  - 重启终端后生效。


---
macOS 系统
方法1（推荐）：配置settings.json
- 创建 ~/.claude/settings.json 文件，内容如下：
  
{
  "env": {
    "ANTHROPIC_AUTH_TOKEN": "替换为您的API Key",
    "ANTHROPIC_BASE_URL": "https://code.newcli.com/claude",
    "CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC": 1
  },
  "permissions": {
    "allow": [],
    "deny": []
  }
}
方法2：临时设置（仅当前终端有效）
- 在 终端 中执行：
export ANTHROPIC_BASE_URL="https://code.newcli.com/claude"
export ANTHROPIC_API_KEY="替换为您的API Key"
方法3：永久设置
1. 编辑 shell 配置文件（根据使用的 shell 选择）：
# 如果是 bash（默认）
echo 'export ANTHROPIC_BASE_URL="https://code.newcli.com/claude"' >> ~/.bash_profile
echo 'export ANTHROPIC_API_KEY="替换为您的API Key"' >> ~/.bash_profile

# 如果是 zsh
echo 'export ANTHROPIC_BASE_URL="https://code.newcli.com/claude"' >> ~/.zshrc
echo 'export ANTHROPIC_API_KEY="替换为您的API Key"' >> ~/.zshrc
2. 立即生效：
source ~/.bash_profile  # 或 source ~/.zshrc


---
Linux 系统
方法1：临时设置（仅当前终端有效）
- 在 终端 中执行：
export ANTHROPIC_BASE_URL="https://code.newcli.com/claude"
export ANTHROPIC_API_KEY="替换为您的API Key"
方法2：永久设置
1. 编辑 shell 配置文件（根据使用的 shell 选择）：
# 如果是 bash
echo 'export ANTHROPIC_BASE_URL="https://code.newcli.com/claude"' >> ~/.bashrc
echo 'export ANTHROPIC_API_KEY="替换为您的API Key"' >> ~/.bashrc

# 如果是 zsh
echo 'export ANTHROPIC_BASE_URL="https://code.newcli.com/claude"' >> ~/.zshrc
echo 'export ANTHROPIC_API_KEY="替换为您的API Key"' >> ~/.zshrc
2. 立即生效：
source ~/.bashrc  # 或 source ~/.zshrc
方法3：配置settings.json
- 创建 ~/.claude/settings.json 文件，内容如下：
{
  "env": {
    "ANTHROPIC_AUTH_TOKEN": "替换为您的API Key",
    "ANTHROPIC_BASE_URL": "https://code.newcli.com/claude",
    "CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC": 1
  },
  "permissions": {
    "allow": [],
    "deny": []
  }
}

---
通用验证方法
在所有系统中，可以通过以下命令验证是否设置成功：
# macOS/Linux
echo $ANTHROPIC_BASE_URL
echo $ANTHROPIC_API_KEY

# Windows PowerShell
echo $env:ANTHROPIC_BASE_URL
echo $env:ANTHROPIC_API_KEY

# Windows CMD
echo %ANTHROPIC_BASE_URL%
echo %ANTHROPIC_API_KEY%
1.2、 ❤️镜像包，安装方式(二选一)
Mac安装教程
1、使用 npm 全局安装 CLI 工具
npm install -g https://code.newcli.com/install --registry=https://registry.npmmirror.com
2、在安装后，您可以访问您的项目文件夹，并在该目录下的终端输入以下命令直接运行 Claude Code
cd your-project-folder（你的项目目录）
claude
3、弹出登录即可使用

---
Windows教程（推荐！无需WSL，新1.0.51版本功能）
1. 前置要求：
  - Node.js（https://nodejs.org/）
  - 依赖 gitbash（https://git-scm.com/downloads）
2. 完成上述安装后，执行以下命令安装
npm install -g https://code.newcli.com/install --registry=https://registry.npmmirror.com


1.3、第三方客户端调用（已开放）
1. 在Roo Code/Kilo Code中使用
  - api供应商选择Anthropic
  - 填写api密钥，
  - 使用自定义基础URL，填写https://code.newcli.com/claude
[图片]
2. 在Cherry Studio中使用
  - 新增供应商平台
  - 类型选择Anthropic
  - 填写API密钥，API地址https://code.newcli.com/claude
  - 点击管理拉取模型加入
[图片]
[图片]
3. 在ChatBox中使用
  - 模型提供方选择Claude
  - 填写API密钥
  - API地址填写https://code.newcli.com/claude/v1
[图片]
二、CodeX安装教程, 新上线！！
本站claudecode 和 codex 额度通用，倍率一致
1、安装Codex
使用 npm 进行安装
npm install -g @openai/codex
2、配置文件
编辑文件 ~/.codex/config.toml
model_provider = "fox"
model = "gpt-5"
model_reasoning_effort = "high"
disable_response_storage = true


[model_providers.fox]
name = "fox"
base_url = "https://code.newcli.com/codex/v1"
wire_api = "responses"
requires_openai_auth = true
编辑文件 ~/.codex/auth.json
{
  "OPENAI_API_KEY": "替换为您的API Key"
}
切换新模型：
codex -m gpt-5-codex
CodeX在VsCode中使用
上方配置生效后，安装官方插件
[图片]

Codex Api在cherry中使用
[图片]
[图片]
[图片]
CodeX Api新增兼容老版本 /v1/chat/completions ，支持cherry和roo，kilo调用
四、ClaudeCode官方中文文档
https://docs.anthropic.com/zh-CN/docs/claude-code/quickstart

---
五、Claude Code功能
1. 直接进行交互：
- Claude Code 提供两种主要的交互方式：
  - 交互模式：运行 claude 启动 REPL 会话
  - 单次模式：使用 claude -p "查询" 进行快速命令
  - 您可以参考：
# 启动交互模式
claude

# 以初始查询启动
claude "解释这个项目"

# 运行单个命令并退出
claude -p "这个函数做什么？"

# 处理管道内容
cat logs.txt | claude -p "分析这些错误"
- 对于 Claude Code Client的常用参数和功能，您可以访问官方文档：CLI 使用和控制 - Anthropic
2. 支持连接到主流IDE
- 您可以直接在IDE中看到Claude Code的改动，在IDE中与其交互。
- 现在支持 VSCode 与 JetBrains
- 如果您使用Linux / MacOS，您可以直接使用该插件
  - 如果您使用VSCode，在VSCode的内置终端唤起Claude Code，插件将被自动安装
  - 如果您使用JetBrains，您需要通过此链接下载：Claude Code [Beta] - IntelliJ IDEs Plugin | Marketplace
- 您可能需要手动指定IDE，通过在Claude Code进行以下交互选择
> /ide
- 对于更多的用法，您可以参考Claude Code的官方文档：IDE integrations - Anthropic
3. 支持连接到Cursor：
方法一：直接安装插件
[图片]
方法二、基于wsl
使用本质：在cursor中本地连接Ubuntu终端使用Claude Code，可以可视化代码的操作！步骤如下：
序号
操作
图例
1
打开cursor
[图片]
2
点击左下角终端图标
[图片]
3
在弹出来的选项框里点击第三个；在弹出来的新选项框里点击Ubuntu选项，cursor就会自动连接Ubuntu系统。

[图片]
[图片]
[图片]

连接完成后显示
[图片]
 没有Connect to WSL using Distro选项
若打开只有2个选项，没有五个选项，原因是没有安装扩展，安装扩展之后重启即可。
[图片]
按下图依次点击，进入扩展界面。
[图片]
扩展详情页
[图片]
在搜索框里搜索WSL，找到图示这个扩展，不要选错，点击安装。安装过程需要翻墙，不然可能会因为网络安装失败。
[图片]
[图片]
此时再点击就有5个选项，选择第三个就行！
[图片]
4. 切换模型 Claude 4 Opus 与 Claude 4 Sonnet：
- 我们强烈推荐您使用Claude 4 Sonnet（默认的），其使用体验与Claude 4 Opus没有明显差别，但计费倍率仅为1/5。
- 如果更换请查阅官网指令
5. 压缩上下文以节省额度：
- Claude Code 通常会有长上下文，我们建议您使用以下斜杠命令来压缩以节省点数，较长的上下文往往需要更多点数。
/compact [instructions] #您可以添加说明
6. 恢复以前的对话：
- 使用以下命令可以恢复您上次的对话
claude --continue
  - 这会立即恢复您最近的对话，无需任何提示。
- 您如果需要显示时间，可以输入此命令
claude --resume
  - 这会显示一个交互式对话选择器，显示：
    - 对话开始时间
    - 初始提示或对话摘要
    - 消息数量
  - 使用箭头键导航并按Enter选择对话，您可以使用这个方法选择上下文。
7. 处理图像信息：
- 您可以使用以下任何方法：
  1. 将图像拖放到Claude Code窗口中（在MacOS上）
  2. 复制图像并使用Ctrl+v粘贴到CLI中（在MacOS上）
  3. 提供图像路径
> 分析这个图像：/path/to/your/image.png
- 您可以完全使用自然语言要求他进行工作，如：
> 这是错误的截图。是什么导致了它？ 
> 这个图像显示了什么？ 
> 描述这个截图中的UI元素 
> 生成CSS以匹配这个设计模型 
> 什么HTML结构可以重新创建这个组件？ 
8. 深入思考：
- 您需要通过自然语言要求其进行深入思考
> 我需要使用OAuth2为我们的API实现一个新的身份验证系统。深入思考在我们的代码库中实现这一点的最佳方法。
> 思考这种方法中潜在的安全漏洞 
> 更深入地思考我们应该处理的边缘情况
  - 推荐您在使用复杂问题的时候使用这一功能，这也会消耗大量的额度点数。
9. 通过 Claude.md 存储重要记忆：
- 您可以使用以下命令设置一个CLAUDE.md文件来存储重要的项目信息、约定和常用命令。
> /init
  - 包括常用命令（构建、测试、lint）以避免重复搜索
  - 记录代码风格偏好和命名约定
  - 添加特定于您项目的重要架构模式
  - CLAUDE.md记忆可用于与团队共享的指令和您的个人偏好。
  - 更多关于记忆的设置，您可以访问此官方文档了解：Claude Code 概述 - Anthropic
  - 在官方文档中，此部分记录了记忆的常用用法：管理Claude的内存 - Anthropic
10. 自动化 CI 和基础设施工作流程
- Claude Code 提供非交互模式，用于无头执行。这在非交互上下文（如脚本、管道和 Github Actions）中运行 Claude Code 时特别有用。
- 使用 --print (-p) 在非交互模式下运行 Claude，如：
claude -p "使用最新更改更新 README" --allowedTools "Bash(git diff:*)" "Bash(git log:*)" Write --disallowedTools ..
11. 上下文通用协议（MCP）：
- 模型上下文协议(MCP)是一个开放协议，使LLM能够访问外部工具和数据源。
- 这是高级功能，您可以访问此文档获取更多配置信息：Introduction - Model Context Protocol
- Claude Code不仅支持接入MCP，同样支持作为MCP服务器等各类高级功能，您可以访问此文档获得更多信息：教程 - Anthropic
12. 使用Git工作树运行并行Claude Code会话：
- Claude Code 支持使用自然语言操作Git，如：
> 提交我的更改
> 创建一个 pr
> 哪个提交在去年十二月添加了 markdown 测试？
> 在 main 分支上变基并解决任何合并冲突
- 您可以使用工作树创建隔离的编码环境。
- 如果您您需要同时处理多个任务，并在Claude Code实例之间完全隔离代码，您可以使用此功能：
  - Git工作树允许您从同一存储库中检出多个分支到单独的目录。每个工作树都有自己的工作目录，文件是隔离的，同时共享相同的Git历史。在官方Git工作树文档中了解更多。
  - 创建新工作树
# 创建带有新分支的工作树 
git worktree add ../project-feature-a -b feature-a

# 或使用现有分支创建工作树
git worktree add ../project-bugfix bugfix-123
    - 这会创建一个包含存储库单独工作副本的新目录。
  - 在每个工作树中运行Claude Code
# 导航到您的工作树 
cd ../project-feature-a

# 在这个隔离环境中运行Claude Code
claude
  - 在另一个终端中：
cd ../project-bugfix
claude
  - 管理您的工作树
# 列出所有工作树
git worktree list

# 完成后移除工作树
git worktree remove ../project-feature-a
  - 每个工作树都有自己独立的文件状态，非常适合并行Claude Code会话
  - 在一个工作树中所做的更改不会影响其他工作树，防止Claude实例相互干扰
  - 所有工作树共享相同的Git历史和远程连接
  - 对于长时间运行的任务，您可以让Claude在一个工作树中工作，同时您在另一个工作树中继续开发
  - 使用描述性目录名称，以便轻松识别每个工作树的任务
  - 记得根据项目的设置在每个新工作树中初始化开发环境。根据您的技术栈，这可能包括：
    - JavaScript项目：运行依赖安装（npm install、yarn）
    - Python项目：设置虚拟环境或使用包管理器安装
    - 其他语言：遵循项目的标准设置流程
13. 其他的自然语言功能：
- 识别未文档化的代码
> 在auth模块中查找没有适当JSDoc注释的函数
- 生成文档
> 为auth.js中未文档化的函数添加JSDoc注释
- 理解陌生代码
> 支付处理系统做什么？
> 查找用户权限在哪里被检查
> 解释缓存层是如何工作的
- 智能编辑代码
> 为注册表单添加输入验证
> 重构日志记录器以使用新的 API
> 修复工作队列中的竞态条件
- 测试或编辑您的代码
> 运行 auth 模块的测试并修复失败
> 查找并修复安全漏洞
> 解释为什么这个测试失败了
14. 常见的斜杠命令：
暂时无法在飞书文档外展示此内容
15. 常用的快捷键：
- 使用 # 快速记忆
  - 通过以 # 开始输入来即时添加记忆
- 始终使用描述性变量名
  - 系统会提示你选择要将其存储在哪个记忆文件中。
- 终端中的换行
  - 使用以下方式输入多行命令：
  - 快速转义：输入 \ 后按 Enter
  - 键盘快捷键：Option+Enter（或配置后的 Shift+Enter）
    - 在终端中设置 Option+Enter：
      - 对于 Mac Terminal.app
        1. 打开设置 → 配置文件 → 键盘
        2. 勾选”将 Option 键用作 Meta 键”
      - 对于 iTerm2 和 VSCode 终端：
        1. 打开设置 → 配置文件 → 按键
        2. 在常规设置下，将左/右 Option 键设置为”Esc+”
      - iTerm2 和 VSCode 用户提示：在 Claude Code 中运行 /terminal-setup 以自动配置 Shift+Enter 作为更直观的替代方案。
      - 有关配置详情，请参见官方文档：设置中的终端设置。
- Vim 模式
  - Claude Code 支持一部分 Vim 键绑定，可以通过 /vim 启用或通过 /config 配置。
  - 支持的功能包括：
    - 模式切换：Esc（到 NORMAL），i/I，a/A，o/O（到 INSERT）
    - 导航：h/j/k/l，w/e/b，0/$/^，gg/G
    - 编辑：x，dw/de/db/dd/D，cw/ce/cb/cc/C，.（重复）
16. 常见的报错：
- 400 - invalid_request_error：您的请求格式或内容存在问题。我们也可能对下面未列出的其他 4XX 状态码使用此错误类型。
- 401 - authentication_error：您的 API 密钥存在问题。
- 403 - permission_error：您的 API 密钥没有使用指定资源的权限。
- 404 - not_found_error：未找到请求的资源。
- 413 - request_too_large：请求超过了允许的最大字节数。 建议使用/compact命令
- 429 - rate_limit_error：您的账户达到了速率限制。
- 500 - api_error：Anthropic 系统内部发生了意外错误。
- 529 - overloaded_error：Anthropic 的 API 暂时过载。
- 当 Anthropic API 在所有用户中遇到高流量时，可能会出现 529 错误。在极少数情况下，如果您的组织使用量急剧增加，您可能会看到此类错误。 为避免 529 错误，请逐步增加流量并保持一致的使用模式。
当通过 SSE 接收流式响应时，可能在返回 200 响应后发生错误，在这种情况下错误处理不会遵循这些标准机制。
17. 其他的高级功能：
- Claude Code可以被用作Claude用作类Unix工具：教程 - Anthropic
- Claude Code支持自定义斜杠指令：教程 - Anthropic
- Claude Code支持使用$ARGUMENTS添加命令参数：教程 - Anthropic
- Claude Code支持高级设置，您可以参考此文档：Claude Code 设置 - Anthropic
  1. 命令行参数
  2. 本地项目设置
  3. 共享项目设置
  4. 用户设置
- Claude Code的安全设置，请参考此官方文档：管理权限和安全 - Anthropic

