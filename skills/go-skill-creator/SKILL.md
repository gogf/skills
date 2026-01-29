---
name: go-skill-creator
description: Go 模块 Skill 生成器。支持分析本地项目或远程 Go 模块。通过 `go list -json` 和 AST 解析深度分析模块结构，自动生成多层级的 skill 文档。
---

# Go Skill Creator

为 Go 模块自动生成 AI 开发技能文档的工具。支持分析本地项目或从远程仓库下载指定版本的 Go 模块。

## 快速开始

```bash
# 分析本地项目
cd /path/to/project
go run ./scripts/analyze_module.go -name myproject -verbose

# 分析远程 Go 模块（自动下载到临时目录）
go run ./scripts/analyze_module.go -module github.com/gogf/gf/v2@latest -name gf -verbose
```

## 命名规则

```
skill-name = <project-name>-skill
```

**规则**：

- 使用小写
- 去除版本号 (v2, v3 → 简化)
- 使用 `-skill` 后缀
- 保持简洁 (通常 2-8 个字符)

| 项目路径 | 技能名 |
|---------|--------|
| github.com/gogf/gf/v2 | `gf-skill` |
| github.com/gin-gonic/gin | `gin-skill` |

## 命令行参数

```bash
go run ./scripts/analyze_module.go [options] [root_path]

Options:
  -root string     根项目目录 (默认: ".")
  -name string     技能名称 (默认: 从主模块提取)
  -output string   输出目录 (默认: 当前目录下的 <name>-skill)
  -module string   Go 模块引用，支持版本 (e.g., github.com/gogf/gf/v2@latest)
  -internal        包含 internal 包 (默认: false)
  -verbose         详细输出 (默认: false)
```

### 模块版本指定

- 如果不指定版本，默认为 `latest`
- 支持完整的语义化版本（如 `v2.0.0`）
- 当使用 `-module` 参数时，工具会在临时目录中创建 go.mod，下载模块，完成分析后自动清理

## 工作原理

### 本地项目分析

1. 扫描项目目录，查找所有 `go.mod` 文件
2. 使用 `go list -json` 获取模块和包的信息
3. 解析 Go 源文件的 AST，提取类型和函数定义
4. 生成 skill 文档

### 远程模块分析

1. 在临时目录中创建临时 `go.mod` 文件
2. 使用 `go mod download` 下载指定版本的模块到本地缓存
3. 使用 `go list` 获取模块的完整信息（包括版本号）
4. 分析模块的所有包和类型
5. 生成 skill 文档
6. 清理临时目录

这种方式支持分析任何已发布到 Go 模块仓库的项目，无需克隆源码。

## Description 优化流程

生成的 SKILL.md 中的 `description` 字段是自动生成的初始版本。根据 [skill-creator 标准](https://github.com/anthropics/anthropic-sdk-python/docs/skills)，description 应遵循以下格式，以便大模型准确判断何时使用该 Skill：

### Description 标准格式

```
{模块功能概述}。别名: {常见名称/简称}。导入路径: {Go 导入路径}。版本: {版本号}。

使用场景：当需要 {具体使用场景 1} 或 {具体使用场景 2} 等 {功能描述} 时使用。
```

### 优化步骤

1. **查看生成的初始 description**
   ```yaml
   description: Go 模块 `github.com/gogf/gf/v2` 的开发指南。基于源码深度分析，提供完整的包结构和 API 参考。
   ```

2. **请求大模型进行优化**
   
   将以下内容提交给大模型：
   ```
   请按照以下标准优化这个 Go 模块 Skill 的 description，确保 Claude 能准确理解何时使用此 Skill：
   
   模块名称: {skill-name}
   当前 description: {current-description}
   
   模块信息:
   - 主模块: {module-path}
   - 别名/常见名: {common-aliases}
   - 版本: {version}
   - 包数量: {package-count}
   - 核心功能: {core-features}
   - 主要使用场景: {use-cases}
   
   请生成符合以下要求的 description（100-150 字）：
   
   格式要求：
   1. 第一句：简洁的功能概述 + 别名
   2. 第二句：完整的 Go 导入路径（作为唯一标识符）
   3. 第三句：版本号（如果适用）
   4. 使用场景：列出 3-4 个具体的使用场景，说明何时应该使用此 Skill
   
   示例（参考格式）：
   "{模块简称} 是一个完整的 {功能描述}。别名: {别名}。Go 导入路径: \`{完整路径}\`。版本: {版本}。
   
   当需要以下任何操作时使用此 Skill：
   - (1) {场景 1}
   - (2) {场景 2}  
   - (3) {场景 3}
   - (4) {场景 4}"
   ```

3. **替换 description**
   
   将大模型优化后的描述替换到 SKILL.md 的 frontmatter 中：
   ```yaml
   ---
   name: {skill-name}
   description: {optimized-description}
   ---
   ```

### 优化示例

**原始描述**：
```
Go 模块 `github.com/gogf/gf/v2` 的开发指南。基于源码深度分析，提供完整的包结构和 API 参考。
```

**优化后描述**（参考）：
```
GoFrame (GF) 是企业级 Go 应用开发框架。别名: GoFrame, GF。Go 导入路径: `github.com/gogf/gf/v2`。版本: v2.x。

当需要以下任何操作时使用此 Skill：
- (1) Web 应用开发：构建 RESTful API、Web 服务、后端应用
- (2) 数据库操作：使用 ORM 进行增删改查、事务管理、数据验证
- (3) 日志和配置管理：配置管理、日志输出、性能监控
- (4) 工具函数和组件：使用提供的 114+ 包进行开发加速
```

### 关键要点

- **包含别名**：GoFrame、GF 等常用简称
- **完整路径**：`github.com/gogf/gf/v2` 作为唯一标识
- **具体场景**：不仅说"Web 开发"，要说"RESTful API 开发"、"数据库 ORM 操作"等
- **何时使用**：包含"何时使用"信息让大模型能正确判断触发条件

### 何时优化

- ✓ 新建 Skill 后进行首次优化
- ✓ 模块更新、版本升级后重新生成并优化
- ✓ 根据实际使用反馈进行迭代优化

## 详细信息

- **分析原理**：参见 [analysis.md](references/analysis.md)
- **使用示例**：参见 [examples.md](references/examples.md)
