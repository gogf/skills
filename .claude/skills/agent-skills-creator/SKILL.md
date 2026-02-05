---
name: agent-skills-creator
description: Agent Skills创建专家。帮助开发者遵循Agent Skills开放标准创建高质量的技能文件，包括正确的目录结构、FrontMatter配置和Markdown指令内容。当需要创建新的Agent Skill或修改现有Skill时使用此技能。
license: Apache-2.0
metadata:
  author: goframe
  version: "1.0.0"
---


这是一个帮助开发者创建符合`Agent Skills`开放标准的技能文件的专家指南。

# Agent Skills 简介

`Agent Skills`是一种轻量级的开放格式，用于扩展`AI`智能体的能力和专业知识。一个技能就是一个包含`SKILL.md`文件的文件夹，该文件包含元数据（至少包含`name`和`description`）和指导智能体如何执行特定任务的说明。

## 核心原理

`Skills`使用 **渐进式披露（Progressive Disclosure）** 来高效管理上下文：

1. **发现阶段**：启动时，智能体只加载每个可用技能的名称和描述，即`SKILL.md`的`frontmatter`
2. **激活阶段**：当任务匹配技能描述时，智能体将完整的 `SKILL.md` 指令读入上下文
3. **执行阶段**：智能体遵循指令，按需加载引用的文件或执行捆绑的代码

# 目录结构

一个标准的`Agent Skill`目录结构如下：

```text
skill-name/
├── SKILL.md          # 必需：指令 + 元数据
├── scripts/          # 可选：可执行代码
├── references/       # 可选：参考文档
└── assets/           # 可选：模板、资源
```

## 必需文件

- **SKILL.md**：唯一必需的文件，包含`YAML frontmatter`和`Markdown`指令

## 可选目录

- **scripts/**：包含智能体可以运行的可执行代码（`Python`、`Bash`、`JavaScript` 等）
- **references/**：包含智能体按需读取的额外文档（如 `REFERENCE.md`、特定领域文件等）
- **assets/**：包含静态资源（模板、图片、数据文件等）

# SKILL.md 文件格式

## FrontMatter 规范

`SKILL.md`文件必须在顶部包含`YAML frontmatter`，格式如下：

```yaml
---
name: skill-name
description: 描述这个技能的作用以及何时使用它。
---
```

### 必需字段

| 字段 | 约束条件 |
|------|----------|
| `name` | 最多`64`字符。只能使用小写字母、数字和连字符。不能以连字符开头或结尾。不能包含连续的连字符（`--`）。必须与父目录名称匹配。 |
| `description` | 最多`1024`字符。非空。应描述技能的功能、说明何时应该使用这个技能、应包含帮助智能体识别相关任务的特定关键词。 |

### 可选字段

```yaml
---
name: pdf-processing
description: 从PDF文件中提取文本和表格，填充表单，合并文档。处理PDF文档或用户提及PDF、表单、文档提取时使用。
license: Apache-2.0
compatibility: 需要 Python 3.8+ 和 pypdf 库
metadata:
  author: example-org
  version: "1.0.0"
  tags: pdf, document-processing
allowed-tools: Bash(python:*) Read Write
---
```

| 字段 | 说明 |
|------|------|
| `license` | 指定技能的许可协议，简短的许可证名称或文件引用 |
| `compatibility` | 说明环境要求（如所需工具、系统包等），最大`500`字符 |
| `metadata` | 存储作者、版本等额外元数据，键值对映射 |
| `allowed-tools` | 预批准可使用的工具（实验性），空格分隔的工具列表 |

## Markdown 指令内容

`FrontMatter`之后的`Markdown`正文包含技能指令。格式没有限制，编写任何有助于智能体有效执行任务的内容。

### 推荐的内容结构

```markdown
# 技能名称

# 何时使用此技能
清晰描述此技能适用的场景和任务类型。

# 如何使用
提供分步指令，帮助智能体理解执行流程。

# 常见示例
提供输入输出示例，帮助智能体理解预期结果。

# 注意事项
说明常见的边缘情况和错误处理方法。

# 参考资源
如果有额外的参考文档，在此列出：
- [详细技术参考](references/REFERENCE.md)
- [表单模板](references/FORMS.md)
```

### 内容编写建议

1. **保持主文件简洁**：建议`SKILL.md`保持在`500`行以内（少于`5000`tokens）
2. **使用相对路径引用**：引用其他文件时使用相对于技能根目录的路径
3. **避免深层嵌套**：文件引用保持在一级深度，避免深层嵌套的引用链
4. **提供具体示例**：包含输入输出示例，帮助智能体理解
5. **处理边缘情况**：说明常见的错误和如何处理

# 文件引用规范

在技能中引用其他文件时，使用相对路径：

```markdown
详细信息请参阅[参考指南](references/REFERENCE.md)。

运行提取脚本：
scripts/extract.py
```

建议保持文件引用在一级深度，避免复杂的引用链。

# Scripts 目录

如果技能需要执行代码，将脚本放在 `scripts/` 目录中。

## 脚本编写建议

- **自包含**：脚本应该是自包含的，或者清楚地记录依赖项
- **错误信息**：包含有用的错误消息
- **边缘情况处理**：优雅地处理边缘情况
- **文档注释**：添加清晰的注释说明脚本的功能

支持的语言取决于智能体实现，常见的包括`Python`、`Bash`和`JavaScript`。

# References 目录

包含智能体按需读取的额外文档：

- `REFERENCE.md` - 详细的技术参考
- `FORMS.md` - 表单模板或结构化数据格式
- 特定领域文件（如 `finance.md`、`legal.md` 等）

保持单个引用文件聚焦。智能体按需加载这些文件，较小的文件意味着更少的上下文使用。

# Assets 目录

包含静态资源：

- **模板**：文档模板、配置模板
- **图片**：图表、示例
- **数据文件**：查找表、模式定义

# 创建新技能的步骤

## 1. 确定技能范围

- 明确技能要解决的问题
- 确定目标用户和使用场景
- 定义技能的边界（做什么，不做什么）

## 2. 创建目录结构

```bash
mkdir -p my-skill/{scripts,references,assets}
touch my-skill/SKILL.md
```

## 3. 编写 FrontMatter

```yaml
---
name: my-skill
description: 清晰描述技能的功能和使用场景，包含相关关键词。
license: Apache-2.0
metadata:
  author: your-org
  version: "1.0.0"
---
```

## 4. 编写指令内容

按照推荐的结构编写`Markdown`指令：
- 说明何时使用
- 提供分步指令
- 包含示例
- 处理边缘情况

## 5. 添加支持文件

根据需要添加：
- 可执行脚本（`scripts/`）
- 参考文档（`references/`）
- 静态资源（`assets/`）

# 技能示例模板

## 基础技能模板

```markdown
---
name: example-skill
description: 简短描述技能的功能和使用场景。包含相关关键词帮助智能体识别任务。
license: Apache-2.0
metadata:
  author: your-org
  version: "1.0.0"
---

# 何时使用此技能

当用户需要...时使用此技能。

# 使用方法

## 步骤 1：准备

说明第一步需要做什么。

## 步骤 2：执行

说明具体执行步骤。

## 步骤 3：验证

说明如何验证结果。

# 示例

## 输入示例

\`\`\`
示例输入内容
\`\`\`

## 输出示例

\`\`\`
预期输出内容
\`\`\`

# 常见问题

## 问题 1：描述

解决方案...

## 问题 2：描述

解决方案...

# 参考资源

- [详细参考](references/REFERENCE.md)
```

## 带脚本的技能模板

```markdown
---
name: advanced-skill
description: 描述技能功能，说明何时使用。
license: Apache-2.0
compatibility: 需要 Python 3.8+ 和 requests 库
metadata:
  author: your-org
  version: "1.0.0"
allowed-tools: Bash(python:*) Read Write
---

# Advanced Skill

# 何时使用

当需要...时使用此技能。

# 前置要求

- Python 3.8 或更高版本
- 安装 requests 库：`pip install requests`

# 使用方法

## 1. 运行主脚本

\`\`\`bash
python scripts/main.py --input <file> --output <result>
\`\`\`

## 2. 检查结果

结果将保存在...

# 脚本说明

- `scripts/main.py` - 主处理脚本
- `scripts/utils.py` - 工具函数

详细的 API 文档请参阅 [references/API.md](references/API.md)。
```
