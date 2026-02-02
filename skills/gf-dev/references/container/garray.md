# garray

> Package: `github.com/gogf/gf/v2/container/garray`

```go
import "github.com/gogf/gf/v2/container/garray"
```

## 概述

Package garray provides most commonly used array containers which also support concurrent-safe/unsafe switch feature.

## 源文件

- `garray.go`
- `garray_func.go`
- `garray_normal_any.go`
- `garray_normal_int.go`
- `garray_normal_str.go`
- `garray_normal_t.go`
- `garray_sorted_any.go`
- `garray_sorted_int.go`
- `garray_sorted_str.go`
- `garray_sorted_t.go`

## 类型定义

### Array

**类型**: struct

Array is a golang array with rich features.
It contains a concurrent-safe/unsafe switch, which should be set
when its initialization and cannot be changed then.


**方法**:

- `func (a *Array) At(index int) value any`
  At returns the value by the specified index.
- `func (a *Array) Get(index int) (value any, found bool)`
  Get returns the value by the specified index.
- `func (a *Array) Set(index int, value any) error`
  Set sets value to specified index.
- `func (a *Array) SetArray(array []any) *Array`
  SetArray sets the underlying slice array with the given `array`.
- `func (a *Array) Replace(array []any) *Array`
  Replace replaces the array items by given `array` from the beginning of array.
- `func (a *Array) Sum() sum int`
  Sum returns the sum of values in an array.
- `func (a *Array) SortFunc(less func) *Array`
  SortFunc sorts the array by custom function `less`.
- `func (a *Array) InsertBefore(index int, values ...any) error`
  InsertBefore inserts the `values` to the front of `index`.
- `func (a *Array) InsertAfter(index int, values ...any) error`
  InsertAfter inserts the `values` to the back of `index`.
- `func (a *Array) Remove(index int) (value any, found bool)`
  Remove removes an item by index.
- `func (a *Array) RemoveValue(value any) bool`
  RemoveValue removes an item by value.
- `func (a *Array) RemoveValues(values ...any)`
  RemoveValues removes multiple items by `values`.
- `func (a *Array) PushLeft(value ...any) *Array`
  PushLeft pushes one or multiple items to the beginning of array.
- `func (a *Array) PushRight(value ...any) *Array`
  PushRight pushes one or multiple items to the end of array.
- `func (a *Array) PopRand() (value any, found bool)`
  PopRand randomly pops and return an item out of array.
- `func (a *Array) PopRands(size int) []any`
  PopRands randomly pops and returns `size` items out of array.
- `func (a *Array) PopLeft() (value any, found bool)`
  PopLeft pops and returns an item from the beginning of array.
- `func (a *Array) PopRight() (value any, found bool)`
  PopRight pops and returns an item from the end of array.
- `func (a *Array) PopLefts(size int) []any`
  PopLefts pops and returns `size` items from the beginning of array.
- `func (a *Array) PopRights(size int) []any`
  PopRights pops and returns `size` items from the end of array.
- `func (a *Array) Range(start int, end ...int) []any`
  Range picks and returns items by range, like array[start:end].
- `func (a *Array) SubSlice(offset int, length ...int) []any`
  SubSlice returns a slice of elements from the array as specified
- `func (a *Array) Append(value ...any) *Array`
  Append is alias of PushRight, please See PushRight.
- `func (a *Array) Len() int`
  Len returns the length of array.
- `func (a *Array) Slice() []any`
  Slice returns the underlying data of array.
- `func (a *Array) Interfaces() []any`
  Interfaces returns current array as []any.
- `func (a *Array) Clone() newArray *Array`
  Clone returns a new array, which is a copy of current array.
- `func (a *Array) Clear() *Array`
  Clear deletes all items of current array.
- `func (a *Array) Contains(value any) bool`
  Contains checks whether a value exists in the array.
- `func (a *Array) Search(value any) int`
  Search searches array by `value`, returns the index of `value`,
- `func (a *Array) Unique() *Array`
  Unique uniques the array, clear repeated items.
- `func (a *Array) LockFunc(f func) *Array`
  LockFunc locks writing by callback function `f`.
- `func (a *Array) RLockFunc(f func) *Array`
  RLockFunc locks reading by callback function `f`.
- `func (a *Array) Merge(array any) *Array`
  Merge merges `array` into current array.
- `func (a *Array) Fill(startIndex int, num int, value any) error`
  Fill fills an array with num entries of the value `value`,
- `func (a *Array) Chunk(size int) [][]any`
  Chunk splits an array into multiple arrays,
- `func (a *Array) Pad(size int, val any) *Array`
  Pad pads array to the specified length with `value`.
- `func (a *Array) Rand() (value any, found bool)`
  Rand randomly returns one item from array(no deleting).
- `func (a *Array) Rands(size int) []any`
  Rands randomly returns `size` items from array(no deleting).
- `func (a *Array) Shuffle() *Array`
  Shuffle randomly shuffles the array.
- `func (a *Array) Reverse() *Array`
  Reverse makes array with elements in reverse order.
- `func (a *Array) Join(glue string) string`
  Join joins array elements with a string `glue`.
- `func (a *Array) CountValues() map[any]int`
  CountValues counts the number of occurrences of all values in the array.
- `func (a *Array) Iterator(f func)`
  Iterator is alias of IteratorAsc.
- `func (a *Array) IteratorAsc(f func)`
  IteratorAsc iterates the array readonly in ascending order with given callbac...
- `func (a *Array) IteratorDesc(f func)`
  IteratorDesc iterates the array readonly in descending order with given callb...
- `func (a *Array) String() string`
  String returns current array as a string, which implements like json.Marshal ...
- `func (a Array) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (a *Array) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (a *Array) UnmarshalValue(value any) error`
  UnmarshalValue is an interface implement which sets any type of value for array.
- `func (a *Array) Filter(filter func) *Array`
  Filter iterates array and filters elements using custom callback function.
- `func (a *Array) FilterNil() *Array`
  FilterNil removes all nil value of the array.
- `func (a *Array) FilterEmpty() *Array`
  FilterEmpty removes all empty value of the array.
- `func (a *Array) Walk(f func) *Array`
  Walk applies a user supplied function `f` to every item of array.
- `func (a *Array) IsEmpty() bool`
  IsEmpty checks whether the array is empty.
- `func (a *Array) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### IntArray

**类型**: struct

IntArray is a golang int array with rich features.
It contains a concurrent-safe/unsafe switch, which should be set
when its initialization and cannot be changed then.


**方法**:

- `func (a *IntArray) At(index int) value int`
  At returns the value by the specified index.
- `func (a *IntArray) Get(index int) (value int, found bool)`
  Get returns the value by the specified index.
- `func (a *IntArray) Set(index int, value int) error`
  Set sets value to specified index.
- `func (a *IntArray) SetArray(array []int) *IntArray`
  SetArray sets the underlying slice array with the given `array`.
- `func (a *IntArray) Replace(array []int) *IntArray`
  Replace replaces the array items by given `array` from the beginning of array.
- `func (a *IntArray) Sum() sum int`
  Sum returns the sum of values in an array.
- `func (a *IntArray) Sort(reverse ...bool) *IntArray`
  Sort sorts the array in increasing order.
- `func (a *IntArray) SortFunc(less func) *IntArray`
  SortFunc sorts the array by custom function `less`.
- `func (a *IntArray) InsertBefore(index int, values ...int) error`
  InsertBefore inserts the `values` to the front of `index`.
- `func (a *IntArray) InsertAfter(index int, values ...int) error`
  InsertAfter inserts the `value` to the back of `index`.
- `func (a *IntArray) Remove(index int) (value int, found bool)`
  Remove removes an item by index.
- `func (a *IntArray) RemoveValue(value int) bool`
  RemoveValue removes an item by value.
- `func (a *IntArray) RemoveValues(values ...int)`
  RemoveValues removes multiple items by `values`.
- `func (a *IntArray) PushLeft(value ...int) *IntArray`
  PushLeft pushes one or multiple items to the beginning of array.
- `func (a *IntArray) PushRight(value ...int) *IntArray`
  PushRight pushes one or multiple items to the end of array.
- `func (a *IntArray) PopLeft() (value int, found bool)`
  PopLeft pops and returns an item from the beginning of array.
- `func (a *IntArray) PopRight() (value int, found bool)`
  PopRight pops and returns an item from the end of array.
- `func (a *IntArray) PopRand() (value int, found bool)`
  PopRand randomly pops and return an item out of array.
- `func (a *IntArray) PopRands(size int) []int`
  PopRands randomly pops and returns `size` items out of array.
- `func (a *IntArray) PopLefts(size int) []int`
  PopLefts pops and returns `size` items from the beginning of array.
- `func (a *IntArray) PopRights(size int) []int`
  PopRights pops and returns `size` items from the end of array.
- `func (a *IntArray) Range(start int, end ...int) []int`
  Range picks and returns items by range, like array[start:end].
- `func (a *IntArray) SubSlice(offset int, length ...int) []int`
  SubSlice returns a slice of elements from the array as specified
- `func (a *IntArray) Append(value ...int) *IntArray`
  Append is alias of PushRight,please See PushRight.
- `func (a *IntArray) Len() int`
  Len returns the length of array.
- `func (a *IntArray) Slice() []int`
  Slice returns the underlying data of array.
- `func (a *IntArray) Interfaces() []any`
  Interfaces returns current array as []any.
- `func (a *IntArray) Clone() newArray *IntArray`
  Clone returns a new array, which is a copy of current array.
- `func (a *IntArray) Clear() *IntArray`
  Clear deletes all items of current array.
- `func (a *IntArray) Contains(value int) bool`
  Contains checks whether a value exists in the array.
- `func (a *IntArray) Search(value int) int`
  Search searches array by `value`, returns the index of `value`,
- `func (a *IntArray) Unique() *IntArray`
  Unique uniques the array, clear repeated items.
- `func (a *IntArray) LockFunc(f func) *IntArray`
  LockFunc locks writing by callback function `f`.
- `func (a *IntArray) RLockFunc(f func) *IntArray`
  RLockFunc locks reading by callback function `f`.
- `func (a *IntArray) Merge(array any) *IntArray`
  Merge merges `array` into current array.
- `func (a *IntArray) Fill(startIndex int, num int, value int) error`
  Fill fills an array with num entries of the value `value`,
- `func (a *IntArray) Chunk(size int) [][]int`
  Chunk splits an array into multiple arrays,
- `func (a *IntArray) Pad(size int, value int) *IntArray`
  Pad pads array to the specified length with `value`.
- `func (a *IntArray) Rand() (value int, found bool)`
  Rand randomly returns one item from array(no deleting).
- `func (a *IntArray) Rands(size int) []int`
  Rands randomly returns `size` items from array(no deleting).
- `func (a *IntArray) Shuffle() *IntArray`
  Shuffle randomly shuffles the array.
- `func (a *IntArray) Reverse() *IntArray`
  Reverse makes array with elements in reverse order.
- `func (a *IntArray) Join(glue string) string`
  Join joins array elements with a string `glue`.
- `func (a *IntArray) CountValues() map[int]int`
  CountValues counts the number of occurrences of all values in the array.
- `func (a *IntArray) Iterator(f func)`
  Iterator is alias of IteratorAsc.
- `func (a *IntArray) IteratorAsc(f func)`
  IteratorAsc iterates the array readonly in ascending order with given callbac...
- `func (a *IntArray) IteratorDesc(f func)`
  IteratorDesc iterates the array readonly in descending order with given callb...
- `func (a *IntArray) String() string`
  String returns current array as a string, which implements like json.Marshal ...
- `func (a IntArray) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (a *IntArray) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (a *IntArray) UnmarshalValue(value any) error`
  UnmarshalValue is an interface implement which sets any type of value for array.
- `func (a *IntArray) Filter(filter func) *IntArray`
  Filter iterates array and filters elements using custom callback function.
- `func (a *IntArray) FilterEmpty() *IntArray`
  FilterEmpty removes all zero value of the array.
- `func (a *IntArray) Walk(f func) *IntArray`
  Walk applies a user supplied function `f` to every item of array.
- `func (a *IntArray) IsEmpty() bool`
  IsEmpty checks whether the array is empty.
- `func (a *IntArray) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### StrArray

**类型**: struct

StrArray is a golang string array with rich features.
It contains a concurrent-safe/unsafe switch, which should be set
when its initialization and cannot be changed then.


**方法**:

- `func (a *StrArray) At(index int) value string`
  At returns the value by the specified index.
- `func (a *StrArray) Get(index int) (value string, found bool)`
  Get returns the value by the specified index.
- `func (a *StrArray) Set(index int, value string) error`
  Set sets value to specified index.
- `func (a *StrArray) SetArray(array []string) *StrArray`
  SetArray sets the underlying slice array with the given `array`.
- `func (a *StrArray) Replace(array []string) *StrArray`
  Replace replaces the array items by given `array` from the beginning of array.
- `func (a *StrArray) Sum() sum int`
  Sum returns the sum of values in an array.
- `func (a *StrArray) Sort(reverse ...bool) *StrArray`
  Sort sorts the array in increasing order.
- `func (a *StrArray) SortFunc(less func) *StrArray`
  SortFunc sorts the array by custom function `less`.
- `func (a *StrArray) InsertBefore(index int, values ...string) error`
  InsertBefore inserts the `values` to the front of `index`.
- `func (a *StrArray) InsertAfter(index int, values ...string) error`
  InsertAfter inserts the `values` to the back of `index`.
- `func (a *StrArray) Remove(index int) (value string, found bool)`
  Remove removes an item by index.
- `func (a *StrArray) RemoveValue(value string) bool`
  RemoveValue removes an item by value.
- `func (a *StrArray) RemoveValues(values ...string)`
  RemoveValues removes multiple items by `values`.
- `func (a *StrArray) PushLeft(value ...string) *StrArray`
  PushLeft pushes one or multiple items to the beginning of array.
- `func (a *StrArray) PushRight(value ...string) *StrArray`
  PushRight pushes one or multiple items to the end of array.
- `func (a *StrArray) PopLeft() (value string, found bool)`
  PopLeft pops and returns an item from the beginning of array.
- `func (a *StrArray) PopRight() (value string, found bool)`
  PopRight pops and returns an item from the end of array.
- `func (a *StrArray) PopRand() (value string, found bool)`
  PopRand randomly pops and return an item out of array.
- `func (a *StrArray) PopRands(size int) []string`
  PopRands randomly pops and returns `size` items out of array.
- `func (a *StrArray) PopLefts(size int) []string`
  PopLefts pops and returns `size` items from the beginning of array.
- `func (a *StrArray) PopRights(size int) []string`
  PopRights pops and returns `size` items from the end of array.
- `func (a *StrArray) Range(start int, end ...int) []string`
  Range picks and returns items by range, like array[start:end].
- `func (a *StrArray) SubSlice(offset int, length ...int) []string`
  SubSlice returns a slice of elements from the array as specified
- `func (a *StrArray) Append(value ...string) *StrArray`
  Append is alias of PushRight,please See PushRight.
- `func (a *StrArray) Len() int`
  Len returns the length of array.
- `func (a *StrArray) Slice() []string`
  Slice returns the underlying data of array.
- `func (a *StrArray) Interfaces() []any`
  Interfaces returns current array as []any.
- `func (a *StrArray) Clone() newArray *StrArray`
  Clone returns a new array, which is a copy of current array.
- `func (a *StrArray) Clear() *StrArray`
  Clear deletes all items of current array.
- `func (a *StrArray) Contains(value string) bool`
  Contains checks whether a value exists in the array.
- `func (a *StrArray) ContainsI(value string) bool`
  ContainsI checks whether a value exists in the array with case-insensitively.
- `func (a *StrArray) Search(value string) int`
  Search searches array by `value`, returns the index of `value`,
- `func (a *StrArray) Unique() *StrArray`
  Unique uniques the array, clear repeated items.
- `func (a *StrArray) LockFunc(f func) *StrArray`
  LockFunc locks writing by callback function `f`.
- `func (a *StrArray) RLockFunc(f func) *StrArray`
  RLockFunc locks reading by callback function `f`.
- `func (a *StrArray) Merge(array any) *StrArray`
  Merge merges `array` into current array.
- `func (a *StrArray) Fill(startIndex int, num int, value string) error`
  Fill fills an array with num entries of the value `value`,
- `func (a *StrArray) Chunk(size int) [][]string`
  Chunk splits an array into multiple arrays,
- `func (a *StrArray) Pad(size int, value string) *StrArray`
  Pad pads array to the specified length with `value`.
- `func (a *StrArray) Rand() (value string, found bool)`
  Rand randomly returns one item from array(no deleting).
- `func (a *StrArray) Rands(size int) []string`
  Rands randomly returns `size` items from array(no deleting).
- `func (a *StrArray) Shuffle() *StrArray`
  Shuffle randomly shuffles the array.
- `func (a *StrArray) Reverse() *StrArray`
  Reverse makes array with elements in reverse order.
- `func (a *StrArray) Join(glue string) string`
  Join joins array elements with a string `glue`.
- `func (a *StrArray) CountValues() map[string]int`
  CountValues counts the number of occurrences of all values in the array.
- `func (a *StrArray) Iterator(f func)`
  Iterator is alias of IteratorAsc.
- `func (a *StrArray) IteratorAsc(f func)`
  IteratorAsc iterates the array readonly in ascending order with given callbac...
- `func (a *StrArray) IteratorDesc(f func)`
  IteratorDesc iterates the array readonly in descending order with given callb...
- `func (a *StrArray) String() string`
  String returns current array as a string, which implements like json.Marshal ...
- `func (a StrArray) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (a *StrArray) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (a *StrArray) UnmarshalValue(value any) error`
  UnmarshalValue is an interface implement which sets any type of value for array.
- `func (a *StrArray) Filter(filter func) *StrArray`
  Filter iterates array and filters elements using custom callback function.
- `func (a *StrArray) FilterEmpty() *StrArray`
  FilterEmpty removes all empty string value of the array.
- `func (a *StrArray) Walk(f func) *StrArray`
  Walk applies a user supplied function `f` to every item of array.
- `func (a *StrArray) IsEmpty() bool`
  IsEmpty checks whether the array is empty.
- `func (a *StrArray) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### TArray

**类型**: struct

TArray is a golang array with rich features.
It contains a concurrent-safe/unsafe switch, which should be set
when its initialization and cannot be changed then.


### SortedArray

**类型**: struct

SortedArray is a golang sorted array with rich features.
It is using increasing order in default, which can be changed by
setting it a custom comparator.
It contains a concurrent-safe/unsafe switch, which should be set
when its initialization and cannot be changed then.


**方法**:

- `func (a *SortedArray) At(index int) value any`
  At returns the value by the specified index.
- `func (a *SortedArray) SetArray(array []any) *SortedArray`
  SetArray sets the underlying slice array with the given `array`.
- `func (a *SortedArray) SetComparator(comparator func)`
  SetComparator sets/changes the comparator for sorting.
- `func (a *SortedArray) Sort() *SortedArray`
  Sort sorts the array in increasing order.
- `func (a *SortedArray) Add(values ...any) *SortedArray`
  Add adds one or multiple values to sorted array, the array always keeps sorted.
- `func (a *SortedArray) Append(values ...any) *SortedArray`
  Append adds one or multiple values to sorted array, the array always keeps so...
- `func (a *SortedArray) Get(index int) (value any, found bool)`
  Get returns the value by the specified index.
- `func (a *SortedArray) Remove(index int) (value any, found bool)`
  Remove removes an item by index.
- `func (a *SortedArray) RemoveValue(value any) bool`
  RemoveValue removes an item by value.
- `func (a *SortedArray) RemoveValues(values ...any)`
  RemoveValues removes an item by `values`.
- `func (a *SortedArray) PopLeft() (value any, found bool)`
  PopLeft pops and returns an item from the beginning of array.
- `func (a *SortedArray) PopRight() (value any, found bool)`
  PopRight pops and returns an item from the end of array.
- `func (a *SortedArray) PopRand() (value any, found bool)`
  PopRand randomly pops and return an item out of array.
- `func (a *SortedArray) PopRands(size int) []any`
  PopRands randomly pops and returns `size` items out of array.
- `func (a *SortedArray) PopLefts(size int) []any`
  PopLefts pops and returns `size` items from the beginning of array.
- `func (a *SortedArray) PopRights(size int) []any`
  PopRights pops and returns `size` items from the end of array.
- `func (a *SortedArray) Range(start int, end ...int) []any`
  Range picks and returns items by range, like array[start:end].
- `func (a *SortedArray) SubSlice(offset int, length ...int) []any`
  SubSlice returns a slice of elements from the array as specified
- `func (a *SortedArray) Sum() sum int`
  Sum returns the sum of values in an array.
- `func (a *SortedArray) Len() int`
  Len returns the length of array.
- `func (a *SortedArray) Slice() []any`
  Slice returns the underlying data of array.
- `func (a *SortedArray) Interfaces() []any`
  Interfaces returns current array as []any.
- `func (a *SortedArray) Contains(value any) bool`
  Contains checks whether a value exists in the array.
- `func (a *SortedArray) Search(value any) index int`
  Search searches array by `value`, returns the index of `value`,
- `func (a *SortedArray) SetUnique(unique bool) *SortedArray`
  SetUnique sets unique mark to the array,
- `func (a *SortedArray) Unique() *SortedArray`
  Unique uniques the array, clear repeated items.
- `func (a *SortedArray) Clone() newArray *SortedArray`
  Clone returns a new array, which is a copy of current array.
- `func (a *SortedArray) Clear() *SortedArray`
  Clear deletes all items of current array.
- `func (a *SortedArray) LockFunc(f func) *SortedArray`
  LockFunc locks writing by callback function `f`.
- `func (a *SortedArray) RLockFunc(f func) *SortedArray`
  RLockFunc locks reading by callback function `f`.
- `func (a *SortedArray) Merge(array any) *SortedArray`
  Merge merges `array` into current array.
- `func (a *SortedArray) Chunk(size int) [][]any`
  Chunk splits an array into multiple arrays,
- `func (a *SortedArray) Rand() (value any, found bool)`
  Rand randomly returns one item from array(no deleting).
- `func (a *SortedArray) Rands(size int) []any`
  Rands randomly returns `size` items from array(no deleting).
- `func (a *SortedArray) Join(glue string) string`
  Join joins array elements with a string `glue`.
- `func (a *SortedArray) CountValues() map[any]int`
  CountValues counts the number of occurrences of all values in the array.
- `func (a *SortedArray) Iterator(f func)`
  Iterator is alias of IteratorAsc.
- `func (a *SortedArray) IteratorAsc(f func)`
  IteratorAsc iterates the array readonly in ascending order with given callbac...
- `func (a *SortedArray) IteratorDesc(f func)`
  IteratorDesc iterates the array readonly in descending order with given callb...
- `func (a *SortedArray) String() string`
  String returns current array as a string, which implements like json.Marshal ...
- `func (a SortedArray) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (a *SortedArray) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (a *SortedArray) UnmarshalValue(value any) err error`
  UnmarshalValue is an interface implement which sets any type of value for array.
- `func (a *SortedArray) FilterNil() *SortedArray`
  FilterNil removes all nil value of the array.
- `func (a *SortedArray) Filter(filter func) *SortedArray`
  Filter iterates array and filters elements using custom callback function.
- `func (a *SortedArray) FilterEmpty() *SortedArray`
  FilterEmpty removes all empty value of the array.
- `func (a *SortedArray) Walk(f func) *SortedArray`
  Walk applies a user supplied function `f` to every item of array.
- `func (a *SortedArray) IsEmpty() bool`
  IsEmpty checks whether the array is empty.
- `func (a *SortedArray) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### SortedIntArray

**类型**: struct

SortedIntArray is a golang sorted int array with rich features.
It is using increasing order in default, which can be changed by
setting it a custom comparator.
It contains a concurrent-safe/unsafe switch, which should be set
when its initialization and cannot be changed then.


**方法**:

- `func (a *SortedIntArray) At(index int) value int`
  At returns the value by the specified index.
- `func (a *SortedIntArray) SetArray(array []int) *SortedIntArray`
  SetArray sets the underlying slice array with the given `array`.
- `func (a *SortedIntArray) Sort() *SortedIntArray`
  Sort sorts the array in increasing order.
- `func (a *SortedIntArray) Add(values ...int) *SortedIntArray`
  Add adds one or multiple values to sorted array, the array always keeps sorted.
- `func (a *SortedIntArray) Append(values ...int) *SortedIntArray`
  Append adds one or multiple values to sorted array, the array always keeps so...
- `func (a *SortedIntArray) Get(index int) (value int, found bool)`
  Get returns the value by the specified index.
- `func (a *SortedIntArray) Remove(index int) (value int, found bool)`
  Remove removes an item by index.
- `func (a *SortedIntArray) RemoveValue(value int) bool`
  RemoveValue removes an item by value.
- `func (a *SortedIntArray) RemoveValues(values ...int)`
  RemoveValues removes an item by `values`.
- `func (a *SortedIntArray) PopLeft() (value int, found bool)`
  PopLeft pops and returns an item from the beginning of array.
- `func (a *SortedIntArray) PopRight() (value int, found bool)`
  PopRight pops and returns an item from the end of array.
- `func (a *SortedIntArray) PopRand() (value int, found bool)`
  PopRand randomly pops and return an item out of array.
- `func (a *SortedIntArray) PopRands(size int) []int`
  PopRands randomly pops and returns `size` items out of array.
- `func (a *SortedIntArray) PopLefts(size int) []int`
  PopLefts pops and returns `size` items from the beginning of array.
- `func (a *SortedIntArray) PopRights(size int) []int`
  PopRights pops and returns `size` items from the end of array.
- `func (a *SortedIntArray) Range(start int, end ...int) []int`
  Range picks and returns items by range, like array[start:end].
- `func (a *SortedIntArray) SubSlice(offset int, length ...int) []int`
  SubSlice returns a slice of elements from the array as specified
- `func (a *SortedIntArray) Len() int`
  Len returns the length of array.
- `func (a *SortedIntArray) Sum() sum int`
  Sum returns the sum of values in an array.
- `func (a *SortedIntArray) Slice() []int`
  Slice returns the underlying data of array.
- `func (a *SortedIntArray) Interfaces() []any`
  Interfaces returns current array as []any.
- `func (a *SortedIntArray) Contains(value int) bool`
  Contains checks whether a value exists in the array.
- `func (a *SortedIntArray) Search(value int) index int`
  Search searches array by `value`, returns the index of `value`,
- `func (a *SortedIntArray) SetUnique(unique bool) *SortedIntArray`
  SetUnique sets unique mark to the array,
- `func (a *SortedIntArray) Unique() *SortedIntArray`
  Unique uniques the array, clear repeated items.
- `func (a *SortedIntArray) Clone() newArray *SortedIntArray`
  Clone returns a new array, which is a copy of current array.
- `func (a *SortedIntArray) Clear() *SortedIntArray`
  Clear deletes all items of current array.
- `func (a *SortedIntArray) LockFunc(f func) *SortedIntArray`
  LockFunc locks writing by callback function `f`.
- `func (a *SortedIntArray) RLockFunc(f func) *SortedIntArray`
  RLockFunc locks reading by callback function `f`.
- `func (a *SortedIntArray) Merge(array any) *SortedIntArray`
  Merge merges `array` into current array.
- `func (a *SortedIntArray) Chunk(size int) [][]int`
  Chunk splits an array into multiple arrays,
- `func (a *SortedIntArray) Rand() (value int, found bool)`
  Rand randomly returns one item from array(no deleting).
- `func (a *SortedIntArray) Rands(size int) []int`
  Rands randomly returns `size` items from array(no deleting).
- `func (a *SortedIntArray) Join(glue string) string`
  Join joins array elements with a string `glue`.
- `func (a *SortedIntArray) CountValues() map[int]int`
  CountValues counts the number of occurrences of all values in the array.
- `func (a *SortedIntArray) Iterator(f func)`
  Iterator is alias of IteratorAsc.
- `func (a *SortedIntArray) IteratorAsc(f func)`
  IteratorAsc iterates the array readonly in ascending order with given callbac...
- `func (a *SortedIntArray) IteratorDesc(f func)`
  IteratorDesc iterates the array readonly in descending order with given callb...
- `func (a *SortedIntArray) String() string`
  String returns current array as a string, which implements like json.Marshal ...
- `func (a SortedIntArray) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (a *SortedIntArray) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (a *SortedIntArray) UnmarshalValue(value any) err error`
  UnmarshalValue is an interface implement which sets any type of value for array.
- `func (a *SortedIntArray) Filter(filter func) *SortedIntArray`
  Filter iterates array and filters elements using custom callback function.
- `func (a *SortedIntArray) FilterEmpty() *SortedIntArray`
  FilterEmpty removes all zero value of the array.
- `func (a *SortedIntArray) Walk(f func) *SortedIntArray`
  Walk applies a user supplied function `f` to every item of array.
- `func (a *SortedIntArray) IsEmpty() bool`
  IsEmpty checks whether the array is empty.
- `func (a *SortedIntArray) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### SortedStrArray

**类型**: struct

SortedStrArray is a golang sorted string array with rich features.
It is using increasing order in default, which can be changed by
setting it a custom comparator.
It contains a concurrent-safe/unsafe switch, which should be set
when its initialization and cannot be changed then.


**方法**:

- `func (a *SortedStrArray) SetArray(array []string) *SortedStrArray`
  SetArray sets the underlying slice array with the given `array`.
- `func (a *SortedStrArray) At(index int) value string`
  At returns the value by the specified index.
- `func (a *SortedStrArray) Sort() *SortedStrArray`
  Sort sorts the array in increasing order.
- `func (a *SortedStrArray) Add(values ...string) *SortedStrArray`
  Add adds one or multiple values to sorted array, the array always keeps sorted.
- `func (a *SortedStrArray) Append(values ...string) *SortedStrArray`
  Append adds one or multiple values to sorted array, the array always keeps so...
- `func (a *SortedStrArray) Get(index int) (value string, found bool)`
  Get returns the value by the specified index.
- `func (a *SortedStrArray) Remove(index int) (value string, found bool)`
  Remove removes an item by index.
- `func (a *SortedStrArray) RemoveValue(value string) bool`
  RemoveValue removes an item by value.
- `func (a *SortedStrArray) RemoveValues(values ...string)`
  RemoveValues removes an item by `values`.
- `func (a *SortedStrArray) PopLeft() (value string, found bool)`
  PopLeft pops and returns an item from the beginning of array.
- `func (a *SortedStrArray) PopRight() (value string, found bool)`
  PopRight pops and returns an item from the end of array.
- `func (a *SortedStrArray) PopRand() (value string, found bool)`
  PopRand randomly pops and return an item out of array.
- `func (a *SortedStrArray) PopRands(size int) []string`
  PopRands randomly pops and returns `size` items out of array.
- `func (a *SortedStrArray) PopLefts(size int) []string`
  PopLefts pops and returns `size` items from the beginning of array.
- `func (a *SortedStrArray) PopRights(size int) []string`
  PopRights pops and returns `size` items from the end of array.
- `func (a *SortedStrArray) Range(start int, end ...int) []string`
  Range picks and returns items by range, like array[start:end].
- `func (a *SortedStrArray) SubSlice(offset int, length ...int) []string`
  SubSlice returns a slice of elements from the array as specified
- `func (a *SortedStrArray) Sum() sum int`
  Sum returns the sum of values in an array.
- `func (a *SortedStrArray) Len() int`
  Len returns the length of array.
- `func (a *SortedStrArray) Slice() []string`
  Slice returns the underlying data of array.
- `func (a *SortedStrArray) Interfaces() []any`
  Interfaces returns current array as []any.
- `func (a *SortedStrArray) Contains(value string) bool`
  Contains checks whether a value exists in the array.
- `func (a *SortedStrArray) ContainsI(value string) bool`
  ContainsI checks whether a value exists in the array with case-insensitively.
- `func (a *SortedStrArray) Search(value string) index int`
  Search searches array by `value`, returns the index of `value`,
- `func (a *SortedStrArray) SetUnique(unique bool) *SortedStrArray`
  SetUnique sets unique mark to the array,
- `func (a *SortedStrArray) Unique() *SortedStrArray`
  Unique uniques the array, clear repeated items.
- `func (a *SortedStrArray) Clone() newArray *SortedStrArray`
  Clone returns a new array, which is a copy of current array.
- `func (a *SortedStrArray) Clear() *SortedStrArray`
  Clear deletes all items of current array.
- `func (a *SortedStrArray) LockFunc(f func) *SortedStrArray`
  LockFunc locks writing by callback function `f`.
- `func (a *SortedStrArray) RLockFunc(f func) *SortedStrArray`
  RLockFunc locks reading by callback function `f`.
- `func (a *SortedStrArray) Merge(array any) *SortedStrArray`
  Merge merges `array` into current array.
- `func (a *SortedStrArray) Chunk(size int) [][]string`
  Chunk splits an array into multiple arrays,
- `func (a *SortedStrArray) Rand() (value string, found bool)`
  Rand randomly returns one item from array(no deleting).
- `func (a *SortedStrArray) Rands(size int) []string`
  Rands randomly returns `size` items from array(no deleting).
- `func (a *SortedStrArray) Join(glue string) string`
  Join joins array elements with a string `glue`.
- `func (a *SortedStrArray) CountValues() map[string]int`
  CountValues counts the number of occurrences of all values in the array.
- `func (a *SortedStrArray) Iterator(f func)`
  Iterator is alias of IteratorAsc.
- `func (a *SortedStrArray) IteratorAsc(f func)`
  IteratorAsc iterates the array readonly in ascending order with given callbac...
- `func (a *SortedStrArray) IteratorDesc(f func)`
  IteratorDesc iterates the array readonly in descending order with given callb...
- `func (a *SortedStrArray) String() string`
  String returns current array as a string, which implements like json.Marshal ...
- `func (a SortedStrArray) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (a *SortedStrArray) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (a *SortedStrArray) UnmarshalValue(value any) err error`
  UnmarshalValue is an interface implement which sets any type of value for array.
- `func (a *SortedStrArray) Filter(filter func) *SortedStrArray`
  Filter iterates array and filters elements using custom callback function.
- `func (a *SortedStrArray) FilterEmpty() *SortedStrArray`
  FilterEmpty removes all empty string value of the array.
- `func (a *SortedStrArray) Walk(f func) *SortedStrArray`
  Walk applies a user supplied function `f` to every item of array.
- `func (a *SortedStrArray) IsEmpty() bool`
  IsEmpty checks whether the array is empty.
- `func (a *SortedStrArray) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### SortedTArray

**类型**: struct

SortedTArray is a golang sorted array with rich features.
It is using increasing order in default, which can be changed by
setting it a custom comparator.
It contains a concurrent-safe/unsafe switch, which should be set
when its initialization and cannot be changed then.


## 函数

### New

```go
func New(safe ...bool) *Array
```

New creates and returns an empty array.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewArray

```go
func NewArray(safe ...bool) *Array
```

NewArray is alias of New, please see New.

### NewArraySize

```go
func NewArraySize(size int, cap int, safe ...bool) *Array
```

NewArraySize create and returns an array with given size and cap.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewArrayRange

```go
func NewArrayRange(start int, end int, step int, safe ...bool) *Array
```

NewArrayRange creates and returns an array by a range from `start` to `end`
with step value `step`.

### NewFrom

```go
func NewFrom(array []any, safe ...bool) *Array
```

NewFrom is alias of NewArrayFrom.
See NewArrayFrom.

### NewFromCopy

```go
func NewFromCopy(array []any, safe ...bool) *Array
```

NewFromCopy is alias of NewArrayFromCopy.
See NewArrayFromCopy.

### NewArrayFrom

```go
func NewArrayFrom(array []any, safe ...bool) *Array
```

NewArrayFrom creates and returns an array with given slice `array`.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewArrayFromCopy

```go
func NewArrayFromCopy(array []any, safe ...bool) *Array
```

NewArrayFromCopy creates and returns an array from a copy of given slice `array`.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewIntArray

```go
func NewIntArray(safe ...bool) *IntArray
```

NewIntArray creates and returns an empty array.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewIntArraySize

```go
func NewIntArraySize(size int, cap int, safe ...bool) *IntArray
```

NewIntArraySize create and returns an array with given size and cap.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewIntArrayRange

```go
func NewIntArrayRange(start int, end int, step int, safe ...bool) *IntArray
```

NewIntArrayRange creates and returns an array by a range from `start` to `end`
with step value `step`.

### NewIntArrayFrom

```go
func NewIntArrayFrom(array []int, safe ...bool) *IntArray
```

NewIntArrayFrom creates and returns an array with given slice `array`.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewIntArrayFromCopy

```go
func NewIntArrayFromCopy(array []int, safe ...bool) *IntArray
```

NewIntArrayFromCopy creates and returns an array from a copy of given slice `array`.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewStrArray

```go
func NewStrArray(safe ...bool) *StrArray
```

NewStrArray creates and returns an empty array.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewStrArraySize

```go
func NewStrArraySize(size int, cap int, safe ...bool) *StrArray
```

NewStrArraySize create and returns an array with given size and cap.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewStrArrayFrom

```go
func NewStrArrayFrom(array []string, safe ...bool) *StrArray
```

NewStrArrayFrom creates and returns an array with given slice `array`.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewStrArrayFromCopy

```go
func NewStrArrayFromCopy(array []string, safe ...bool) *StrArray
```

NewStrArrayFromCopy creates and returns an array from a copy of given slice `array`.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewTArray

```go
func NewTArray(safe ...bool) *
```

NewTArray creates and returns an empty array.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewTArraySize

```go
func NewTArraySize(size int, cap int, safe ...bool) *
```

NewTArraySize create and returns an array with given size and cap.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewTArrayFrom

```go
func NewTArrayFrom(array []T, safe ...bool) *
```

NewTArrayFrom creates and returns an array with given slice `array`.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewTArrayFromCopy

```go
func NewTArrayFromCopy(array []T, safe ...bool) *
```

NewTArrayFromCopy creates and returns an array from a copy of given slice `array`.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### MarshalJSON

```go
func (a ) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.
Note that do not use pointer as its receiver here.

### NewSortedArray

```go
func NewSortedArray(comparator func, safe ...bool) *SortedArray
```

NewSortedArray creates and returns an empty sorted array.
The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.
The parameter `comparator` used to compare values to sort in array,
if it returns value < 0, means `a` < `b`; the `a` will be inserted before `b`;
if it returns value = 0, means `a` = `b`; the `a` will be replaced by     `b`;
if it returns value > 0, means `a` > `b`; the `a` will be inserted after  `b`;

### NewSortedArraySize

```go
func NewSortedArraySize(cap int, comparator func, safe ...bool) *SortedArray
```

NewSortedArraySize create and returns an sorted array with given size and cap.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewSortedArrayRange

```go
func NewSortedArrayRange(start int, end int, step int, comparator func, safe ...bool) *SortedArray
```

NewSortedArrayRange creates and returns an array by a range from `start` to `end`
with step value `step`.

### NewSortedArrayFrom

```go
func NewSortedArrayFrom(array []any, comparator func, safe ...bool) *SortedArray
```

NewSortedArrayFrom creates and returns an sorted array with given slice `array`.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewSortedArrayFromCopy

```go
func NewSortedArrayFromCopy(array []any, comparator func, safe ...bool) *SortedArray
```

NewSortedArrayFromCopy creates and returns an sorted array from a copy of given slice `array`.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewSortedIntArray

```go
func NewSortedIntArray(safe ...bool) *SortedIntArray
```

NewSortedIntArray creates and returns an empty sorted array.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewSortedIntArrayComparator

```go
func NewSortedIntArrayComparator(comparator func, safe ...bool) *SortedIntArray
```

NewSortedIntArrayComparator creates and returns an empty sorted array with specified comparator.
The parameter `safe` is used to specify whether using array in concurrent-safety which is false in default.

### NewSortedIntArraySize

```go
func NewSortedIntArraySize(cap int, safe ...bool) *SortedIntArray
```

NewSortedIntArraySize create and returns an sorted array with given size and cap.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewSortedIntArrayRange

```go
func NewSortedIntArrayRange(start int, end int, step int, safe ...bool) *SortedIntArray
```

NewSortedIntArrayRange creates and returns an array by a range from `start` to `end`
with step value `step`.

### NewSortedIntArrayFrom

```go
func NewSortedIntArrayFrom(array []int, safe ...bool) *SortedIntArray
```

NewSortedIntArrayFrom creates and returns an sorted array with given slice `array`.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewSortedIntArrayFromCopy

```go
func NewSortedIntArrayFromCopy(array []int, safe ...bool) *SortedIntArray
```

NewSortedIntArrayFromCopy creates and returns an sorted array from a copy of given slice `array`.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewSortedStrArray

```go
func NewSortedStrArray(safe ...bool) *SortedStrArray
```

NewSortedStrArray creates and returns an empty sorted array.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewSortedStrArrayComparator

```go
func NewSortedStrArrayComparator(comparator func, safe ...bool) *SortedStrArray
```

NewSortedStrArrayComparator creates and returns an empty sorted array with specified comparator.
The parameter `safe` is used to specify whether using array in concurrent-safety which is false in default.

### NewSortedStrArraySize

```go
func NewSortedStrArraySize(cap int, safe ...bool) *SortedStrArray
```

NewSortedStrArraySize create and returns an sorted array with given size and cap.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewSortedStrArrayFrom

```go
func NewSortedStrArrayFrom(array []string, safe ...bool) *SortedStrArray
```

NewSortedStrArrayFrom creates and returns an sorted array with given slice `array`.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewSortedStrArrayFromCopy

```go
func NewSortedStrArrayFromCopy(array []string, safe ...bool) *SortedStrArray
```

NewSortedStrArrayFromCopy creates and returns an sorted array from a copy of given slice `array`.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewSortedTArray

```go
func NewSortedTArray(comparator func, safe ...bool) *
```

NewSortedTArray creates and returns an empty sorted array.
The parameter `safe` is used to specify whether using array in concurrent-safety, which is false in default.
The parameter `comparator` used to compare values to sort in array,
if it returns value < 0, means `a` < `b`; the `a` will be inserted before `b`;
if it returns value = 0, means `a` = `b`; the `a` will be replaced by     `b`;
if it returns value > 0, means `a` > `b`; the `a` will be inserted after  `b`;

### NewSortedTArraySize

```go
func NewSortedTArraySize(cap int, comparator func, safe ...bool) *
```

NewSortedTArraySize create and returns an sorted array with given size and cap.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewSortedTArrayFrom

```go
func NewSortedTArrayFrom(array []T, comparator func, safe ...bool) *
```

NewSortedTArrayFrom creates and returns an sorted array with given slice `array`.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### NewSortedTArrayFromCopy

```go
func NewSortedTArrayFromCopy(array []T, comparator func, safe ...bool) *
```

NewSortedTArrayFromCopy creates and returns an sorted array from a copy of given slice `array`.
The parameter `safe` is used to specify whether using array in concurrent-safety,
which is false in default.

### MarshalJSON

```go
func (a ) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.
Note that do not use pointer as its receiver here.

## 内部依赖

- `errors/gcode`
- `errors/gerror`
- `internal/deepcopy`
- `internal/empty`
- `internal/json`
- `internal/rwmutex`
- `text/gstr`
- `util/gconv`
- `util/grand`
- `util/gutil`

