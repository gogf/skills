# 使用示例

## 示例 1: 分析 GoFrame 项目

```bash
cd /path/to/gf
go run .codebuddy/skills/go-skill-creator/scripts/analyze_module.go -name gf -verbose
```

生成结构：
```
gf-skill/
├── SKILL.md          # 索引文档
└── references/
    ├── container/     # 10 个容器包
    ├── database/      # gdb, gredis
    ├── net/          # ghttp, gclient
    ├── os/           # 23 个 OS 包
    └── ...
```

## 示例 2: 分析 Gin 框架

```bash
cd /path/to/gin-project
go run .codebuddy/skills/go-skill-creator/scripts/analyze_module.go -name gin
```

## 示例 3: 包含 internal 包

```bash
go run analyze_module.go -internal -verbose
```

## 示例 4: 分析多模块项目

```bash
# monorepo 结构
project/
├── go.mod          # 主模块
├── core/           # 核心包
├── contrib/        # 贡献模块
│   ├── plugin1/go.mod
│   └── plugin2/go.mod
└── cmd/            # 命令行工具

# 自动分析所有模块
go run analyze_module.go -verbose

# 输出包含所有模块的包
```

## 生成的 Reference 文件示例

每个包的详细文档包含：

```markdown
# gbuild

> Package: `github.com/gogf/gf/v2/os/gbuild`

```go
import "github.com/gogf/gf/v2/os/gbuild"
```

## 源码位置

**导入路径**: `github.com/gogf/gf/v2/os/gbuild`

## 源文件

- `gbuild.go`

## 类型定义

### BuildInfo

**类型**: struct

BuildInfo maintains the built info of current binary.

## 函数

### Info

```go
func Info() BuildInfo
```

Info returns the basic built information of binary.
```

## 适用场景

- 为内部 Go 模块创建开发技能
- 为开源项目生成 API 参考
- 多模块项目的统一文档生成
- 团队知识沉淀
