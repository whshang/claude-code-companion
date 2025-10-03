# 提交前检查清单

在提交代码到 Git 前，请确认以下事项：

## ✅ 敏感信息检查

- [ ] `config.yaml` 已添加到 `.gitignore`
- [ ] 没有任何测试脚本包含真实 API 密钥
- [ ] 日志文件 (`logs/`, `*.log`) 不会被提交
- [ ] 数据库文件 (`*.db`) 不会被提交

## 🔍 快速验证命令

```bash
# 1. 检查是否有敏感信息泄露
git diff | grep -E '(sk-ant-|sk-|cr_|Bearer [a-zA-Z0-9]{30,})'
# 应该返回空，如果有匹配，立即检查！

# 2. 查看即将提交的文件
git status

# 3. 确认 config.yaml 未被追踪
git ls-files | grep "^config\.yaml$"
# 应该返回空

# 4. 验证 .gitignore 生效
git check-ignore config.yaml test_*.sh logs/*.db
# 应该显示这些文件被忽略
```

## 📋 安全提交步骤

```bash
# 1. 查看当前更改
git status

# 2. 查看具体修改内容
git diff

# 3. 添加安全的文件
git add README.md CHANGELOG.md .gitignore config.yaml.example
git add internal/ web/ docs/

# 4. 提交（不包含 config.yaml）
git commit -m "feat: 添加 Codex 客户端支持

- 实现 Codex 格式自动检测和转换
- 支持客户端特定的端点路由
- 增强模型重写功能
- 完善响应验证逻辑

详见 CHANGELOG.md
"

# 5. 推送前再次确认
git log -1 --stat
git push origin master
```

## ⚠️ 如果不小心提交了密钥

如果已经提交了包含密钥的文件：

```bash
# 1. 立即撤销最后一次提交（如果还没 push）
git reset --soft HEAD^

# 2. 如果已经 push，需要强制覆盖（谨慎使用）
git reset --hard HEAD^
git push origin master --force

# 3. 更换所有泄露的 API 密钥！
```

## 📦 推荐的提交内容

### ✅ 应该提交的文件
- `README.md` - 项目文档
- `CHANGELOG.md` - 版本历史
- `config.yaml.example` - 配置模板
- `.gitignore` - 忽略规则
- `internal/` - 源代码
- `web/` - Web 资源
- `docs/` - 设计文档
- `go.mod`, `go.sum` - 依赖文件
- `Makefile` - 构建脚本

### ❌ 不应该提交的文件
- `config.yaml` - 包含真实密钥
- `config-*.yaml` - 临时配置
- `test_*.sh` - 包含密钥的测试脚本
- `logs/` - 日志文件
- `*.db` - 数据库文件
- `debug/` - 调试导出
- `claude-code-companion` - 编译后的二进制

## 🛡️ 额外的安全建议

1. **定期轮换密钥**：即使没有泄露，也应该定期更换 API 密钥
2. **使用环境变量**：考虑使用环境变量存储敏感信息
3. **本地配置**：为不同环境创建不同的配置文件（dev/staging/prod）
4. **审查历史**：定期检查 Git 历史，确保没有遗漏的敏感信息

## 📝 提交信息规范

使用语义化提交信息：

```
feat: 新功能
fix: 修复 bug
docs: 文档更新
style: 代码格式化
refactor: 重构
test: 测试相关
chore: 构建/工具变更
```

示例：
```
feat: 添加 Codex 客户端支持
fix: 修复 SSE 流 [DONE] 标记缺失问题
docs: 更新 README，添加 Codex 配置说明
refactor: 重构端点选择逻辑
```


