# gring

> Package: `github.com/gogf/gf/v2/container/gring`

```go
import "github.com/gogf/gf/v2/container/gring"
```

## 概述

Package gring provides a concurrent-safe/unsafe ring(circular lists).

## 源文件

- `gring.go`
- `gring_t.go`

## 类型定义

### Ring

**类型**: struct

Ring is a struct of ring structure.

Deprecated.


**方法**:

- `func (r *Ring) Val() any`
  Val returns the item's value of current position.
- `func (r *Ring) Len() int`
  Len returns the size of ring.
- `func (r *Ring) Cap() int`
  Cap returns the capacity of ring.
- `func (r *Ring) Set(value any) *Ring`
  Set sets value to the item of current position.
- `func (r *Ring) Put(value any) *Ring`
  Put sets `value` to current item of ring and moves position to next item.
- `func (r *Ring) Move(n int) *Ring`
  Move moves n % r.Len() elements backward (n < 0) or forward (n >= 0)
- `func (r *Ring) Prev() *Ring`
  Prev returns the previous ring element. r must not be empty.
- `func (r *Ring) Next() *Ring`
  Next returns the next ring element. r must not be empty.
- `func (r *Ring) Link(s *Ring) *Ring`
  Link connects ring r with ring s such that r.Next()
- `func (r *Ring) Unlink(n int) *Ring`
  Unlink removes n % r.Len() elements from the ring r, starting
- `func (r *Ring) RLockIteratorNext(f func)`
  RLockIteratorNext iterates and locks reading forward
- `func (r *Ring) RLockIteratorPrev(f func)`
  RLockIteratorPrev iterates and locks reading backward
- `func (r *Ring) SliceNext() []any`
  SliceNext returns a copy of all item values as slice forward from current pos...
- `func (r *Ring) SlicePrev() []any`
  SlicePrev returns a copy of all item values as slice backward from current po...

### TRing

**类型**: struct

TRing is a struct of ring structure.


## 函数

### New

```go
func New(cap int, safe ...bool) *Ring
```

New creates and returns a Ring structure of `cap` elements.
The optional parameter `safe` specifies whether using this structure in concurrent safety,
which is false in default.

Deprecated.

### NewTRing

```go
func NewTRing(cap int, safe ...bool) *
```

NewTRing creates and returns a Ring structure of `cap` elements.
The optional parameter `safe` specifies whether using this structure in concurrent safety,
which is false in default.

## 内部依赖

- `container/gtype`
- `internal/rwmutex`

