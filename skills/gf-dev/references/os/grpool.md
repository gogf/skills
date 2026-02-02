# grpool

> Package: `github.com/gogf/gf/v2/os/grpool`

```go
import "github.com/gogf/gf/v2/os/grpool"
```

## 概述

Package grpool implements a goroutine reusable pool.

## 源文件

- `grpool.go`
- `grpool_pool.go`
- `grpool_supervisor.go`

## 类型定义

### Func

**类型**: type

Func is the pool function which contains context parameter.


### RecoverFunc

**类型**: type

RecoverFunc is the pool runtime panic recover function which contains context parameter.


### Pool

**类型**: struct

Pool manages the goroutines using pool.


**方法**:

- `func (p *Pool) Add(ctx context.Context, f Func) error`
  Add pushes a new job to the pool.
- `func (p *Pool) AddWithRecover(ctx context.Context, userFunc Func, recoverFunc RecoverFunc) error`
  AddWithRecover pushes a new job to the pool with specified recover function.
- `func (p *Pool) Cap() int`
  Cap returns the capacity of the pool.
- `func (p *Pool) Size() int`
  Size returns current goroutine count of the pool.
- `func (p *Pool) Jobs() int`
  Jobs returns current job count of the pool.
- `func (p *Pool) IsClosed() bool`
  IsClosed returns if pool is closed.
- `func (p *Pool) Close()`
  Close closes the goroutine pool, which makes all goroutines exit.

## 函数

### New

```go
func New(limit ...int) *Pool
```

New creates and returns a new goroutine pool object.
The parameter `limit` is used to limit the max goroutine count,
which is not limited in default.

### Add

```go
func Add(ctx context.Context, f Func) error
```

Add pushes a new job to the default goroutine pool.
The job will be executed asynchronously.

### AddWithRecover

```go
func AddWithRecover(ctx context.Context, userFunc Func, recoverFunc RecoverFunc) error
```

AddWithRecover pushes a new job to the default pool with specified recover function.

The optional `recoverFunc` is called when any panic during executing of `userFunc`.
If `recoverFunc` is not passed or given nil, it ignores the panic from `userFunc`.
The job will be executed asynchronously.

### Size

```go
func Size() int
```

Size returns current goroutine count of default goroutine pool.

### Jobs

```go
func Jobs() int
```

Jobs returns current job count of default goroutine pool.

## 内部依赖

- `container/glist`
- `container/gtype`
- `errors/gcode`
- `errors/gerror`
- `os/gtimer`
- `util/grand`

