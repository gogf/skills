---
description: "生成GoFrame Skills"
agent: "agent"
# model: "Claude Sonnet 4.5"
---

# 背景

我需要给GoFrame框架创建Agent Skills，以便Agent能更好地帮助用户使用GoFrame框架开发更规范、严谨、可靠和高质量的Go代码。

编写GoFrame Skills时的必要参考资料：
- GoFrame源码本地路径：/Users/john/Workspace/github/gogf/gf
- GoFrame框架使用手册：/Users/john/Workspace/github/gogf/gf-site/docs
- GoFrame框架示例代码：/Users/john/Workspace/github/gogf/examples

# 要求

需要严格遵循Agent Skills编写规范。

## 生成的Skill的目录路径

目录路径地址：../skills/.github/skills/gf-engineer

## Skill的目录结构规范

```text
gf-engineer/
├── SKILL.md       # 主要的Skill描述文件
└── references/    # Skill参考的相关资料
```

## SKILL.md文件内容结构规范

生成SKILL.md文件中需要包含以下关键内容：
- 顶部的YAML FrontMatter的description中需要注明，用户在使用Go语言开发时可以调用此Skill
- 最后以表格形式列出references目录下的文件其简要说明，展示的文件层级最多展示3级，例如：

| 文件路径 | 说明 |
| -------- | ---- |
| [核心组件/命令管理/命令管理-命令行对象](./references/核心组件/命令管理/命令管理-命令行对象.md) | 命令行对象的使用和管理 |
| [核心组件/命令管理/命令管理-基本概念](./references/核心组件/命令管理/命令管理-基本概念.md) | 命令管理的基本概念介绍 |

- 告诉AI Agent在编写代码时参考这些references资料内容，以保证生成的代码符合GoFrame框架的最佳实践和规范。

## 生成references目录下的文件

references目录下的内容是GoFrame框架官网文档的精简版本，目录结构与官网文档目录结构保持一致。精简规则如下：
- 每个文档只保留最核心的内容，不能原样拷贝已有文档内容。
- 文档顶部的frontmatter中如果存在noskill标签，则该文档不需要生成。
- 生成的文档内容不再包含顶部的frontmatter内容。
- 由于文档内容较多，你需要分拆为多个任务跟进文档的生成。



