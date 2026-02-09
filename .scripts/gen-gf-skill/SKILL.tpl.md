---
name: {{SKILL_NAME}}
description: GoFrame开发框架专属技能集。为Go语言开发者提供完整的框架使用指南，涵盖命令行管理、配置管理、日志组件、错误处理、数据校验、类型转换、缓存管理、模板引擎、数据库ORM、I18n国际化等核心组件的最佳实践。包含项目工程化结构规范、开发模式指引、常见问题解决方案以及丰富的实战代码示例。适用于构建RESTful API、gRPC微服务、Web应用、CLI工具等各类Go项目，帮助开发者快速掌握GoFrame框架特性，提升开发效率和代码质量。
license: Apache-2.0
metadata:
  author: GoFrame
  version: {{SKILL_VERSION}}
---

# 重要规范

- 在使用GoFrame作为主要框架开发完整工程类型的项目时，应该首先需要安装GoFrame CLI开发工具。
- GoFrame工程规范中由开发工具自动维护的代码文件，不允许手动生成或修改，应当使用GoFrame CLI工具进行相关操作。
- 除非用户有明确明确要求，否则不自动使用GoFrame中的实验性特性。例如，默认在工程实践中使用service目录维护业务逻辑的具体实现，而不是使用logic目录。
- 在工程实践中，优先参考examples目录中的示例代码实现，其次参考给定参考文档中的文档资料。
- 参考了哪些示例代码和文档资料，需要回显给用户展示，以便用户能够清晰地知道自己参考了哪些资料来完成当前的开发任务。

# Go参考资料

{{SKILL_REFERENCES}}

# Go示例代码

{{SKILL_EXAMPLES}}