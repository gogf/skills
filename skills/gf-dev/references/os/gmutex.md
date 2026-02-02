# gmutex

> Package: `github.com/gogf/gf/v2/os/gmutex`

```go
import "github.com/gogf/gf/v2/os/gmutex"
```

## 概述

Package gmutex inherits and extends sync.Mutex and sync.RWMutex with more futures.

## 源文件

- `gmutex.go`
- `gmutex_mutex.go`
- `gmutex_rwmutex.go`

## 类型定义

### Mutex

**类型**: struct

Mutex is a high level Mutex, which implements more rich features for mutex.


**方法**:

- `func (m *Mutex) LockFunc(f func)`
  LockFunc locks the mutex for writing with given callback function `f`.
- `func (m *Mutex) TryLockFunc(f func) result bool`
  TryLockFunc tries locking the mutex for writing with given callback function ...

### RWMutex

**类型**: struct

RWMutex is a high level RWMutex, which implements more rich features for mutex.


**方法**:

- `func (m *RWMutex) LockFunc(f func)`
  LockFunc locks the mutex for writing with given callback function `f`.
- `func (m *RWMutex) RLockFunc(f func)`
  RLockFunc locks the mutex for reading with given callback function `f`.
- `func (m *RWMutex) TryLockFunc(f func) result bool`
  TryLockFunc tries locking the mutex for writing with given callback function ...
- `func (m *RWMutex) TryRLockFunc(f func) result bool`
  TryRLockFunc tries locking the mutex for reading with given callback function...

## 函数

### New

```go
func New() *RWMutex
```

New creates and returns a new mutex.

Deprecated: use Mutex or RWMutex instead.

