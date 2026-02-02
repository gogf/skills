# gctx

> Package: `github.com/gogf/gf/v2/os/gctx`

```go
import "github.com/gogf/gf/v2/os/gctx"
```

## 概述

Package gctx wraps context.Context and provides extra context features.

## 源文件

- `gctx.go`
- `gctx_never_done.go`

## 类型定义

### Ctx

**类型**: type

### StrKey

**类型**: type

## 函数

### New

```go
func New() context.Context
```

New creates and returns a context which contains context id.

### WithCtx

```go
func WithCtx(ctx context.Context) context.Context
```

WithCtx creates and returns a context containing context id upon given parent context `ctx`.

Deprecated: use WithSpan instead.

### WithSpan

```go
func WithSpan(ctx context.Context, spanName string) context.Context
```

WithSpan creates and returns a context containing span upon given parent context `ctx`.

### CtxId

```go
func CtxId(ctx context.Context) string
```

CtxId retrieves and returns the context id from context.

### SetInitCtx

```go
func SetInitCtx(ctx context.Context)
```

SetInitCtx sets custom initialization context.
Note that this function cannot be called in multiple goroutines.

### GetInitCtx

```go
func GetInitCtx() context.Context
```

GetInitCtx returns the initialization context.
Initialization context is used in `main` or `init` functions.

### NeverDone

```go
func NeverDone(ctx context.Context) context.Context
```

NeverDone wraps and returns a new context object that will be never done,
which forbids the context manually done, to make the context can be propagated
to asynchronous goroutines.

Note that, it does not affect the closing (canceling) of the parent context,
as it is a wrapper for its parent, which only affects the next context handling.

## 内部依赖

- `net/gtrace`

## 外部依赖

- `go.opentelemetry.io/otel`
- `go.opentelemetry.io/otel/propagation`

