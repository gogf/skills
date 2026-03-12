---
name: goframe-v2
description: GoFrame development skill. TRIGGER when writing/modifying Go files, implementing services, creating APIs, or database operations. DO NOT TRIGGER for frontend/shell scripts.
license: Apache-2.0
---

# Critical Conventions

## Project Development Standards
- For complete projects (HTTP/microservices), install GoFrame CLI and use `gf init` to create project scaffolding. See [Project Creation - init](./references/开发工具/项目创建-init.md) for details.
- Auto-generated code files (dao, do, entity) MUST NOT be manually created or modified per GoFrame conventions.
- Unless explicitly requested, do NOT use the `logic/` directory for business logic. Implement business logic directly in the `service/` directory.
- Reference complete project examples:
  - HTTP service best practice: [user-http-service](./examples/practices/user-http-service)
  - gRPC service best practice: [user-grpc-service](./examples/practices/user-grpc-service)

## Component Usage Standards
- Before creating new methods or variables, check if they already exist elsewhere and reuse existing implementations.
- Use the `gerror` component for all error handling to ensure complete stack traces for traceability.
- When exploring new components, prioritize GoFrame built-in components and reference best practice code from examples.
- **Database Operations MUST use DO objects** (`internal/model/do/`), never `g.Map` or `map[string]interface{}`. DO struct fields are `interface{}`; unset fields remain `nil` and are automatically ignored by the ORM:
  ```go
  // Good - use DO object
  dao.Users.Ctx(ctx).Where(cols.Id, id).Data(do.User{Uid: uid}).Update()

  // Good - conditional fields, unset fields are nil and ignored
  data := do.User{}
  if password != "" { data.PasswordHash = hash }
  if isAdmin != nil { data.IsAdmin = *isAdmin }
  dao.Users.Ctx(ctx).Where(cols.Id, id).Data(data).Update()

  // Good - explicitly set a column to NULL using gdb.Raw
  dao.Instances.Ctx(ctx).Where(cols.Id, id).Data(do.Instance{IdleSince: gdb.Raw("NULL")}).Update()

  // Bad - never use g.Map for database operations
  dao.Users.Ctx(ctx).Data(g.Map{cols.Uid: uid}).Update()
  ```
## Code Style Standards
- **Variable Declarations**: When defining multiple variables, use a `var` block to group them for better alignment and readability:
  ```go
  // Good - aligned and clean
  var (
      authSvc       *auth.Service
      bizCtxSvc     *bizctx.Service
      k8sSvc        *svcK8s.Service
      notebookSvc   *notebook.Service
      middlewareSvc *middleware.Service
  )

  // Avoid - scattered declarations
  authSvc := auth.New()
  bizCtxSvc := bizctx.New()
  k8sSvc := svcK8s.New()
  ```
- Apply this pattern when you have 3 or more related variable declarations in the same scope.

# GoFrame Documentation
Complete GoFrame development resources covering component design, usage, best practices, and considerations: [GoFrame Documentation](./references/README.MD)

# GoFrame Code Examples
Rich practical code examples covering HTTP services, gRPC services, and various project types: [GoFrame Examples](./examples/README.MD)