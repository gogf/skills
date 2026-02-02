# gqueue

> Package: `github.com/gogf/gf/v2/container/gqueue`

```go
import "github.com/gogf/gf/v2/container/gqueue"
```

## 概述

Package gqueue provides dynamic/static concurrent-safe queue.

## 源文件

- `gqueue.go`
- `gqueue_t.go`

## 类型定义

### Queue

**类型**: struct

Queue is a concurrent-safe queue built on doubly linked list and channel.


**方法**:

- `func (q *Queue) Push(v any)`
  Push pushes the data `v` into the queue.
- `func (q *Queue) Pop() any`
  Pop pops an item from the queue in FIFO way.
- `func (q *Queue) Close()`
  Close closes the queue.
- `func (q *Queue) Len() length int64`
  Len returns the length of the queue.
- `func (q *Queue) Size() int64`
  Size is alias of Len.

### TQueue

**类型**: struct

TQueue is a concurrent-safe queue built on doubly linked list and channel.


## 函数

### New

```go
func New(limit ...int) *Queue
```

New returns an empty queue object.
Optional parameter `limit` is used to limit the size of the queue, which is unlimited in default.
When `limit` is given, the queue will be static and high performance which is comparable with stdlib channel.

### NewTQueue

```go
func NewTQueue(limit ...int) *
```

NewTQueue returns an empty queue object.
Optional parameter `limit` is used to limit the size of the queue, which is unlimited in default.
When `limit` is given, the queue will be static and high performance which is comparable with stdlib channel.

## 内部依赖

- `container/glist`
- `container/gtype`

