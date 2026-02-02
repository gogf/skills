# gset

> Package: `github.com/gogf/gf/v2/container/gset`

```go
import "github.com/gogf/gf/v2/container/gset"
```

## 概述

Package gset provides kinds of concurrent-safe/unsafe sets.

## 源文件

- `gset_any_set.go`
- `gset_int_set.go`
- `gset_str_set.go`
- `gset_t_set.go`

## 类型定义

### Set

**类型**: struct

Set is consisted of any items.


**方法**:

- `func (set *Set) Iterator(f func)`
  Iterator iterates the set readonly with given callback function `f`,
- `func (set *Set) Add(items ...any)`
  Add adds one or multiple items to the set.
- `func (set *Set) AddIfNotExist(item any) bool`
  AddIfNotExist checks whether item exists in the set,
- `func (set *Set) AddIfNotExistFunc(item any, f func) bool`
  AddIfNotExistFunc checks whether item exists in the set,
- `func (set *Set) AddIfNotExistFuncLock(item any, f func) bool`
  AddIfNotExistFuncLock checks whether item exists in the set,
- `func (set *Set) Contains(item any) bool`
  Contains checks whether the set contains `item`.
- `func (set *Set) Remove(item any)`
  Remove deletes `item` from set.
- `func (set *Set) Size() int`
  Size returns the size of the set.
- `func (set *Set) Clear()`
  Clear deletes all items of the set.
- `func (set *Set) Slice() []any`
  Slice returns all items of the set as slice.
- `func (set *Set) Join(glue string) string`
  Join joins items with a string `glue`.
- `func (set *Set) String() string`
  String returns items as a string, which implements like json.Marshal does.
- `func (set *Set) LockFunc(f func)`
  LockFunc locks writing with callback function `f`.
- `func (set *Set) RLockFunc(f func)`
  RLockFunc locks reading with callback function `f`.
- `func (set *Set) Equal(other *Set) bool`
  Equal checks whether the two sets equal.
- `func (set *Set) IsSubsetOf(other *Set) bool`
  IsSubsetOf checks whether the current set is a sub-set of `other`.
- `func (set *Set) Union(others ...*Set) newSet *Set`
  Union returns a new set which is the union of `set` and `others`.
- `func (set *Set) Diff(others ...*Set) newSet *Set`
  Diff returns a new set which is the difference set from `set` to `others`.
- `func (set *Set) Intersect(others ...*Set) newSet *Set`
  Intersect returns a new set which is the intersection from `set` to `others`.
- `func (set *Set) Complement(full *Set) newSet *Set`
  Complement returns a new set which is the complement from `set` to `full`.
- `func (set *Set) Merge(others ...*Set) *Set`
  Merge adds items from `others` sets into `set`.
- `func (set *Set) Sum() sum int`
  Sum sums items.
- `func (set *Set) Pop() any`
  Pop randomly pops an item from set.
- `func (set *Set) Pops(size int) []any`
  Pops randomly pops `size` items from set.
- `func (set *Set) Walk(f func) *Set`
  Walk applies a user supplied function `f` to every item of set.
- `func (set Set) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (set *Set) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (set *Set) UnmarshalValue(value any) err error`
  UnmarshalValue is an interface implement which sets any type of value for set.
- `func (set *Set) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### IntSet

**类型**: struct

IntSet is consisted of int items.


**方法**:

- `func (set *IntSet) Iterator(f func)`
  Iterator iterates the set readonly with given callback function `f`,
- `func (set *IntSet) Add(item ...int)`
  Add adds one or multiple items to the set.
- `func (set *IntSet) AddIfNotExist(item int) bool`
  AddIfNotExist checks whether item exists in the set,
- `func (set *IntSet) AddIfNotExistFunc(item int, f func) bool`
  AddIfNotExistFunc checks whether item exists in the set,
- `func (set *IntSet) AddIfNotExistFuncLock(item int, f func) bool`
  AddIfNotExistFuncLock checks whether item exists in the set,
- `func (set *IntSet) Contains(item int) bool`
  Contains checks whether the set contains `item`.
- `func (set *IntSet) Remove(item int)`
  Remove deletes `item` from set.
- `func (set *IntSet) Size() int`
  Size returns the size of the set.
- `func (set *IntSet) Clear()`
  Clear deletes all items of the set.
- `func (set *IntSet) Slice() []int`
  Slice returns the an of items of the set as slice.
- `func (set *IntSet) Join(glue string) string`
  Join joins items with a string `glue`.
- `func (set *IntSet) String() string`
  String returns items as a string, which implements like json.Marshal does.
- `func (set *IntSet) LockFunc(f func)`
  LockFunc locks writing with callback function `f`.
- `func (set *IntSet) RLockFunc(f func)`
  RLockFunc locks reading with callback function `f`.
- `func (set *IntSet) Equal(other *IntSet) bool`
  Equal checks whether the two sets equal.
- `func (set *IntSet) IsSubsetOf(other *IntSet) bool`
  IsSubsetOf checks whether the current set is a sub-set of `other`.
- `func (set *IntSet) Union(others ...*IntSet) newSet *IntSet`
  Union returns a new set which is the union of `set` and `other`.
- `func (set *IntSet) Diff(others ...*IntSet) newSet *IntSet`
  Diff returns a new set which is the difference set from `set` to `other`.
- `func (set *IntSet) Intersect(others ...*IntSet) newSet *IntSet`
  Intersect returns a new set which is the intersection from `set` to `other`.
- `func (set *IntSet) Complement(full *IntSet) newSet *IntSet`
  Complement returns a new set which is the complement from `set` to `full`.
- `func (set *IntSet) Merge(others ...*IntSet) *IntSet`
  Merge adds items from `others` sets into `set`.
- `func (set *IntSet) Sum() sum int`
  Sum sums items.
- `func (set *IntSet) Pop() int`
  Pop randomly pops an item from set.
- `func (set *IntSet) Pops(size int) []int`
  Pops randomly pops `size` items from set.
- `func (set *IntSet) Walk(f func) *IntSet`
  Walk applies a user supplied function `f` to every item of set.
- `func (set IntSet) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (set *IntSet) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (set *IntSet) UnmarshalValue(value any) err error`
  UnmarshalValue is an interface implement which sets any type of value for set.
- `func (set *IntSet) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### StrSet

**类型**: struct

StrSet is consisted of string items.


**方法**:

- `func (set *StrSet) Iterator(f func)`
  Iterator iterates the set readonly with given callback function `f`,
- `func (set *StrSet) Add(item ...string)`
  Add adds one or multiple items to the set.
- `func (set *StrSet) AddIfNotExist(item string) bool`
  AddIfNotExist checks whether item exists in the set,
- `func (set *StrSet) AddIfNotExistFunc(item string, f func) bool`
  AddIfNotExistFunc checks whether item exists in the set,
- `func (set *StrSet) AddIfNotExistFuncLock(item string, f func) bool`
  AddIfNotExistFuncLock checks whether item exists in the set,
- `func (set *StrSet) Contains(item string) bool`
  Contains checks whether the set contains `item`.
- `func (set *StrSet) ContainsI(item string) bool`
  ContainsI checks whether a value exists in the set with case-insensitively.
- `func (set *StrSet) Remove(item string)`
  Remove deletes `item` from set.
- `func (set *StrSet) Size() int`
  Size returns the size of the set.
- `func (set *StrSet) Clear()`
  Clear deletes all items of the set.
- `func (set *StrSet) Slice() []string`
  Slice returns the an of items of the set as slice.
- `func (set *StrSet) Join(glue string) string`
  Join joins items with a string `glue`.
- `func (set *StrSet) String() string`
  String returns items as a string, which implements like json.Marshal does.
- `func (set *StrSet) LockFunc(f func)`
  LockFunc locks writing with callback function `f`.
- `func (set *StrSet) RLockFunc(f func)`
  RLockFunc locks reading with callback function `f`.
- `func (set *StrSet) Equal(other *StrSet) bool`
  Equal checks whether the two sets equal.
- `func (set *StrSet) IsSubsetOf(other *StrSet) bool`
  IsSubsetOf checks whether the current set is a sub-set of `other`.
- `func (set *StrSet) Union(others ...*StrSet) newSet *StrSet`
  Union returns a new set which is the union of `set` and `other`.
- `func (set *StrSet) Diff(others ...*StrSet) newSet *StrSet`
  Diff returns a new set which is the difference set from `set` to `other`.
- `func (set *StrSet) Intersect(others ...*StrSet) newSet *StrSet`
  Intersect returns a new set which is the intersection from `set` to `other`.
- `func (set *StrSet) Complement(full *StrSet) newSet *StrSet`
  Complement returns a new set which is the complement from `set` to `full`.
- `func (set *StrSet) Merge(others ...*StrSet) *StrSet`
  Merge adds items from `others` sets into `set`.
- `func (set *StrSet) Sum() sum int`
  Sum sums items.
- `func (set *StrSet) Pop() string`
  Pop randomly pops an item from set.
- `func (set *StrSet) Pops(size int) []string`
  Pops randomly pops `size` items from set.
- `func (set *StrSet) Walk(f func) *StrSet`
  Walk applies a user supplied function `f` to every item of set.
- `func (set StrSet) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (set *StrSet) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (set *StrSet) UnmarshalValue(value any) err error`
  UnmarshalValue is an interface implement which sets any type of value for set.
- `func (set *StrSet) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### NilChecker

**类型**: type

NilChecker is a function that checks whether the given value is nil.


### TSet

**类型**: struct

TSet[T] is consisted of any items.


## 函数

### New

```go
func New(safe ...bool) *Set
```

New create and returns a new set, which contains un-repeated items.
The parameter `safe` is used to specify whether using set in concurrent-safety,
which is false in default.

### NewSet

```go
func NewSet(safe ...bool) *Set
```

NewSet create and returns a new set, which contains un-repeated items.
Also see New.

### NewFrom

```go
func NewFrom(items any, safe ...bool) *Set
```

NewFrom returns a new set from `items`.
Parameter `items` can be either a variable of any type, or a slice.

### NewIntSet

```go
func NewIntSet(safe ...bool) *IntSet
```

NewIntSet create and returns a new set, which contains un-repeated items.
The parameter `safe` is used to specify whether using set in concurrent-safety,
which is false in default.

### NewIntSetFrom

```go
func NewIntSetFrom(items []int, safe ...bool) *IntSet
```

NewIntSetFrom returns a new set from `items`.

### NewStrSet

```go
func NewStrSet(safe ...bool) *StrSet
```

NewStrSet create and returns a new set, which contains un-repeated items.
The parameter `safe` is used to specify whether using set in concurrent-safety,
which is false in default.

### NewStrSetFrom

```go
func NewStrSetFrom(items []string, safe ...bool) *StrSet
```

NewStrSetFrom returns a new set from `items`.

### NewTSet

```go
func NewTSet(safe ...bool) *
```

NewTSet creates and returns a new set, which contains un-repeated items.
Also see New.

### NewTSetWithChecker

```go
func NewTSetWithChecker(checker , safe ...bool) *
```

NewTSetWithChecker creates and returns a new set with a custom nil checker.
The parameter `nilChecker` is a function used to determine if a value is nil.
The parameter `safe` is used to specify whether using set in concurrent-safety mode.

### NewTSetFrom

```go
func NewTSetFrom(items []T, safe ...bool) *
```

NewTSetFrom returns a new set from `items`.
`items` - A slice of type T.

### NewTSetWithCheckerFrom

```go
func NewTSetWithCheckerFrom(items []T, checker , safe ...bool) *
```

NewTSetWithCheckerFrom returns a new set from `items` with a custom nil checker.
The parameter `items` is a slice of elements to be added to the set.
The parameter `checker` is a function used to determine if a value is nil.
The parameter `safe` is used to specify whether using set in concurrent-safety mode.

### MarshalJSON

```go
func (set ) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

## 内部依赖

- `internal/json`
- `internal/rwmutex`
- `text/gstr`
- `util/gconv`

