# gpool

> Package: `github.com/gogf/gf/v2/container/gpool`

```go
import "github.com/gogf/gf/v2/container/gpool"
```

## 概述

Package gpool provides object-reusable concurrent-safe pool.

## 源文件

- `gpool.go`
- `gpool_t.go`

## 类型定义

### Pool

**类型**: struct

Pool is an Object-Reusable Pool.


**方法**:

- `func (p *Pool) Put(value any) error`
  Put puts an item to pool.
- `func (p *Pool) MustPut(value any)`
  MustPut puts an item to pool, it panics if any error occurs.
- `func (p *Pool) Clear()`
  Clear clears pool, which means it will remove all items from pool.
- `func (p *Pool) Get() (any, error)`
  Get picks and returns an item from pool. If the pool is empty and NewFunc is ...
- `func (p *Pool) Size() int`
  Size returns the count of available items of pool.
- `func (p *Pool) Close()`
  Close closes the pool. If `p` has ExpireFunc,

### NewFunc

**类型**: type

NewFunc Creation function for object.


### ExpireFunc

**类型**: type

ExpireFunc Destruction function for object.


### TPool

**类型**: struct

TPool is an Object-Reusable Pool.


### TPoolNewFunc

**类型**: type

TPoolNewFunc Creation function for object.


### TPoolExpireFunc

**类型**: type

TPoolExpireFunc Destruction function for object.


## 函数

### New

```go
func New(ttl time.Duration, newFunc NewFunc, expireFunc ...ExpireFunc) *Pool
```

New creates and returns a new object pool.
To ensure execution efficiency, the expiration time cannot be modified once it is set.

Note the expiration logic:
ttl = 0 : not expired;
ttl < 0 : immediate expired after use;
ttl > 0 : timeout expired;

### NewTPool

```go
func NewTPool(ttl time.Duration, newFunc , expireFunc ...) *
```

NewTPool creates and returns a new object pool.
To ensure execution efficiency, the expiration time cannot be modified once it is set.

Note the expiration logic:
ttl = 0 : not expired;
ttl < 0 : immediate expired after use;
ttl > 0 : timeout expired;

## 内部依赖

- `container/glist`
- `container/gtype`
- `errors/gcode`
- `errors/gerror`
- `os/gtime`
- `os/gtimer`

