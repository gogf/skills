# gf 开发技能

## 基本信息

| 属性 | 值 |
|------|----|
| 技能名称 | gf-dev |
| 主模块路径 | github.com/gogf/gf/v2 |
| 子模块数量 | 31 |
| 数据来源 | go list -json + AST 解析 |
| 包数量 | 114 |

## 目录结构

```
gf-dev/
├── SKILL.md       # 主技能文档
├── README.md      # 本文件
└── references/    # 包详细文档
    ├── cmd/
    ├── container/
    ├── contrib/
    ├── crypto/
    ├── database/
    ├── debug/
    ├── encoding/
    ├── errors/
    ├── frame/
    ├── i18n/
    ├── net/
    ├── os/
    ├── root.md
    ├── test/
    ├── text/
    ├── util/
```

## 更新文档

运行分析工具重新生成：

```bash
go run .codebuddy/skills/go-skill-creator/scripts/analyze_module.go -name gf .
```
