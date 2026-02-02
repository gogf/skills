# gmlock

> Package: `github.com/gogf/gf/v2/os/gmlock`

```go
import "github.com/gogf/gf/v2/os/gmlock"
```

## 概述

Package gmlock implements a concurrent-safe memory-based locker.

## 源文件

- `gmlock.go`
- `gmlock_locker.go`

## 类型定义

### Locker

**类型**: struct

Locker is a memory based locker.
Note that there's no cache expire mechanism for mutex in locker.
You need remove certain mutex manually when you do not want use it anymore.


**方法**:

- `func (l *Locker) Lock(key string)`
  Lock locks the `key` with writing lock.
- `func (l *Locker) TryLock(key string) bool`
  TryLock tries locking the `key` with writing lock,
- `func (l *Locker) Unlock(key string)`
  Unlock unlocks the writing lock of the `key`.
- `func (l *Locker) RLock(key string)`
  RLock locks the `key` with reading lock.
- `func (l *Locker) TryRLock(key string) bool`
  TryRLock tries locking the `key` with reading lock.
- `func (l *Locker) RUnlock(key string)`
  RUnlock unlocks the reading lock of the `key`.
- `func (l *Locker) LockFunc(key string, f func)`
  LockFunc locks the `key` with writing lock and callback function `f`.
- `func (l *Locker) RLockFunc(key string, f func)`
  RLockFunc locks the `key` with reading lock and callback function `f`.
- `func (l *Locker) TryLockFunc(key string, f func) bool`
  TryLockFunc locks the `key` with writing lock and callback function `f`.
- `func (l *Locker) TryRLockFunc(key string, f func) bool`
  TryRLockFunc locks the `key` with reading lock and callback function `f`.
- `func (l *Locker) Remove(key string)`
  Remove removes mutex with given `key` from locker.
- `func (l *Locker) Clear()`
  Clear removes all mutexes from locker.

## 函数

### Lock

```go
func Lock(key string)
```

Lock locks the `key` with writing lock.
If there's a write/reading lock the `key`,
it will blocks until the lock is released.

### TryLock

```go
func TryLock(key string) bool
```

TryLock tries locking the `key` with writing lock,
it returns true if success, or if there's a write/reading lock the `key`,
it returns false.

### Unlock

```go
func Unlock(key string)
```

Unlock unlocks the writing lock of the `key`.

### RLock

```go
func RLock(key string)
```

RLock locks the `key` with reading lock.
If there's a writing lock on `key`,
it will blocks until the writing lock is released.

### TryRLock

```go
func TryRLock(key string) bool
```

TryRLock tries locking the `key` with reading lock.
It returns true if success, or if there's a writing lock on `key`, it returns false.

### RUnlock

```go
func RUnlock(key string)
```

RUnlock unlocks the reading lock of the `key`.

### LockFunc

```go
func LockFunc(key string, f func)
```

LockFunc locks the `key` with writing lock and callback function `f`.
If there's a write/reading lock the `key`,
it will blocks until the lock is released.

It releases the lock after `f` is executed.

### RLockFunc

```go
func RLockFunc(key string, f func)
```

RLockFunc locks the `key` with reading lock and callback function `f`.
If there's a writing lock the `key`,
it will blocks until the lock is released.

It releases the lock after `f` is executed.

### TryLockFunc

```go
func TryLockFunc(key string, f func) bool
```

TryLockFunc locks the `key` with writing lock and callback function `f`.
It returns true if success, or else if there's a write/reading lock the `key`, it return false.

It releases the lock after `f` is executed.

### TryRLockFunc

```go
func TryRLockFunc(key string, f func) bool
```

TryRLockFunc locks the `key` with reading lock and callback function `f`.
It returns true if success, or else if there's a writing lock the `key`, it returns false.

It releases the lock after `f` is executed.

### Remove

```go
func Remove(key string)
```

Remove removes mutex with given `key`.

### New

```go
func New() *Locker
```

New creates and returns a new memory locker.
A memory locker can lock/unlock with dynamic string key.

## 内部依赖

- `container/gmap`

