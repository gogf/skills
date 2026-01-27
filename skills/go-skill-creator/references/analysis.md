# 分析方法

## 模块发现

```bash
find . -name go.mod
```

自动扫描项目中所有 `go.mod` 文件，支持：

- 单模块项目
- 多模块 monorepo
- 嵌套子模块（如 contrib/ 下的独立模块）

## 包信息获取

```bash
go list -json ./...
```

获取每个模块的包信息：

| 字段 | 说明 |
|------|------|
| `ImportPath` | 导入路径 |
| `Name` | 包名 |
| `Doc` | 包文档 |
| `GoFiles` | 源文件列表 |
| `Imports` | 直接依赖 |
| `Deps` | 所有依赖 |

## AST 解析

使用 Go AST 解析器提取代码结构：

**类型定义**：
- struct
- interface
- type alias

**方法**：
- 带 receiver 的函数
- 方法签名
- 方法文档

**函数**：
- 包级导出函数
- 函数签名
- 函数文档

## 源码定位

每个包的 reference 文件包含源码位置信息：

```markdown
**导入路径**: `github.com/gogf/gf/v2/os/gbuild`

如需查看完整源码实现，可以通过以下命令获取包的实际路径：

```bash
go list -f '{{.Dir}}' github.com/gogf/gf/v2/os/gbuild
```
```

这样 AI 可以根据需要动态获取实际源码文件，而不是硬编码路径。
