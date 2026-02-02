# glist

> Package: `github.com/gogf/gf/v2/container/glist`

```go
import "github.com/gogf/gf/v2/container/glist"
```

## 概述

Package glist provides most commonly used doubly linked list container which also supports concurrent-safe/unsafe switch feature.

## 源文件

- `glist.go`
- `glist_t.go`

## 类型定义

### List

**类型**: struct

**方法**:

- `func (l *List) PushFront(v any) e *Element`
  PushFront inserts a new element `e` with value `v` at the front of list `l` a...
- `func (l *List) PushBack(v any) e *Element`
  PushBack inserts a new element `e` with value `v` at the back of list `l` and...
- `func (l *List) PushFronts(values []any)`
  PushFronts inserts multiple new elements with values `values` at the front of...
- `func (l *List) PushBacks(values []any)`
  PushBacks inserts multiple new elements with values `values` at the back of l...
- `func (l *List) PopBack() value any`
  PopBack removes the element from back of `l` and returns the value of the ele...
- `func (l *List) PopFront() value any`
  PopFront removes the element from front of `l` and returns the value of the e...
- `func (l *List) PopBacks(max int) values []any`
  PopBacks removes `max` elements from back of `l`
- `func (l *List) PopFronts(max int) values []any`
  PopFronts removes `max` elements from front of `l`
- `func (l *List) PopBackAll() []any`
  PopBackAll removes all elements from back of `l`
- `func (l *List) PopFrontAll() []any`
  PopFrontAll removes all elements from front of `l`
- `func (l *List) FrontAll() values []any`
  FrontAll copies and returns values of all elements from front of `l` as slice.
- `func (l *List) BackAll() values []any`
  BackAll copies and returns values of all elements from back of `l` as slice.
- `func (l *List) FrontValue() value any`
  FrontValue returns value of the first element of `l` or nil if the list is em...
- `func (l *List) BackValue() value any`
  BackValue returns value of the last element of `l` or nil if the list is empty.
- `func (l *List) Front() e *Element`
  Front returns the first element of list `l` or nil if the list is empty.
- `func (l *List) Back() e *Element`
  Back returns the last element of list `l` or nil if the list is empty.
- `func (l *List) Len() length int`
  Len returns the number of elements of list `l`.
- `func (l *List) Size() int`
  Size is alias of Len.
- `func (l *List) MoveBefore(e *Element, p *Element)`
  MoveBefore moves element `e` to its new position before `p`.
- `func (l *List) MoveAfter(e *Element, p *Element)`
  MoveAfter moves element `e` to its new position after `p`.
- `func (l *List) MoveToFront(e *Element)`
  MoveToFront moves element `e` to the front of list `l`.
- `func (l *List) MoveToBack(e *Element)`
  MoveToBack moves element `e` to the back of list `l`.
- `func (l *List) PushBackList(other *List)`
  PushBackList inserts a copy of an other list at the back of list `l`.
- `func (l *List) PushFrontList(other *List)`
  PushFrontList inserts a copy of an other list at the front of list `l`.
- `func (l *List) InsertAfter(p *Element, v any) e *Element`
  InsertAfter inserts a new element `e` with value `v` immediately after `p` an...
- `func (l *List) InsertBefore(p *Element, v any) e *Element`
  InsertBefore inserts a new element `e` with value `v` immediately before `p` ...
- `func (l *List) Remove(e *Element) value any`
  Remove removes `e` from `l` if `e` is an element of list `l`.
- `func (l *List) Removes(es []*Element)`
  Removes removes multiple elements `es` from `l` if `es` are elements of list ...
- `func (l *List) RemoveAll()`
  RemoveAll removes all elements from list `l`.
- `func (l *List) Clear()`
  Clear is alias of RemoveAll.
- `func (l *List) RLockFunc(f func)`
  RLockFunc locks reading with given callback function `f` within RWMutex.RLock.
- `func (l *List) LockFunc(f func)`
  LockFunc locks writing with given callback function `f` within RWMutex.Lock.
- `func (l *List) Iterator(f func)`
  Iterator is alias of IteratorAsc.
- `func (l *List) IteratorAsc(f func)`
  IteratorAsc iterates the list readonly in ascending order with given callback...
- `func (l *List) IteratorDesc(f func)`
  IteratorDesc iterates the list readonly in descending order with given callba...
- `func (l *List) Join(glue string) string`
  Join joins list elements with a string `glue`.
- `func (l *List) String() string`
  String returns current list as a string.
- `func (l List) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (l *List) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (l *List) UnmarshalValue(value any) err error`
  UnmarshalValue is an interface implement which sets any type of value for list.
- `func (l *List) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### Element

**类型**: type

### TElement

**类型**: struct

TElement is an element of a linked list.


### TList

**类型**: struct

## 函数

### New

```go
func New(safe ...bool) *List
```

New creates and returns a new empty doubly linked list.

### NewFrom

```go
func NewFrom(array []any, safe ...bool) *List
```

NewFrom creates and returns a list from a copy of given slice `array`.
The parameter `safe` is used to specify whether using list in concurrent-safety,
which is false in default.

### NewT

```go
func NewT(safe ...bool) *
```

NewT creates and returns a new empty doubly linked list.

### NewTFrom

```go
func NewTFrom(array []T, safe ...bool) *
```

NewTFrom creates and returns a list from a copy of given slice `array`.
The parameter `safe` is used to specify whether using list in concurrent-safety,
which is false in default.

### MarshalJSON

```go
func (l ) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

## 内部依赖

- `internal/deepcopy`
- `internal/json`
- `internal/rwmutex`
- `util/gconv`

