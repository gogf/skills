# gmap

> Package: `github.com/gogf/gf/v2/container/gmap`

```go
import "github.com/gogf/gf/v2/container/gmap"
```

## 概述

Package gmap provides most commonly used map container which also support concurrent-safe/unsafe switch feature.

## 源文件

- `gmap.go`
- `gmap_hash_any_any_map.go`
- `gmap_hash_int_any_map.go`
- `gmap_hash_int_int_map.go`
- `gmap_hash_int_str_map.go`
- `gmap_hash_k_v_map.go`
- `gmap_hash_str_any_map.go`
- `gmap_hash_str_int_map.go`
- `gmap_hash_str_str_map.go`
- `gmap_list_k_v_map.go`
- `gmap_list_map.go`
- `gmap_tree_k_v_map.go`
- `gmap_tree_map.go`

## 类型定义

### Map

**类型**: type

### HashMap

**类型**: type

### AnyAnyMap

**类型**: struct

AnyAnyMap wraps map type `map[any]any` and provides more map features.


**方法**:

- `func (m *AnyAnyMap) Iterator(f func)`
  Iterator iterates the hash map readonly with custom callback function `f`.
- `func (m *AnyAnyMap) Clone(safe ...bool) *AnyAnyMap`
  Clone returns a new hash map with copy of current map data.
- `func (m *AnyAnyMap) Map() map[any]any`
  Map returns the underlying data map.
- `func (m *AnyAnyMap) MapCopy() map[any]any`
  MapCopy returns a shallow copy of the underlying data of the hash map.
- `func (m *AnyAnyMap) MapStrAny() map[string]any`
  MapStrAny returns a copy of the underlying data of the map as map[string]any.
- `func (m *AnyAnyMap) FilterEmpty()`
  FilterEmpty deletes all key-value pair of which the value is empty.
- `func (m *AnyAnyMap) FilterNil()`
  FilterNil deletes all key-value pair of which the value is nil.
- `func (m *AnyAnyMap) Set(key any, value any)`
  Set sets key-value to the hash map.
- `func (m *AnyAnyMap) Sets(data map[any]any)`
  Sets batch sets key-values to the hash map.
- `func (m *AnyAnyMap) Search(key any) (value any, found bool)`
  Search searches the map with given `key`.
- `func (m *AnyAnyMap) Get(key any) value any`
  Get returns the value by given `key`.
- `func (m *AnyAnyMap) Pop() (key any, value any)`
  Pop retrieves and deletes an item from the map.
- `func (m *AnyAnyMap) Pops(size int) map[any]any`
  Pops retrieves and deletes `size` items from the map.
- `func (m *AnyAnyMap) GetOrSet(key any, value any) any`
  GetOrSet returns the value by key,
- `func (m *AnyAnyMap) GetOrSetFunc(key any, f func) any`
  GetOrSetFunc returns the value by key,
- `func (m *AnyAnyMap) GetOrSetFuncLock(key any, f func) any`
  GetOrSetFuncLock returns the value by key,
- `func (m *AnyAnyMap) GetVar(key any) *gvar.Var`
  GetVar returns a Var with the value by given `key`.
- `func (m *AnyAnyMap) GetVarOrSet(key any, value any) *gvar.Var`
  GetVarOrSet returns a Var with result from GetOrSet.
- `func (m *AnyAnyMap) GetVarOrSetFunc(key any, f func) *gvar.Var`
  GetVarOrSetFunc returns a Var with result from GetOrSetFunc.
- `func (m *AnyAnyMap) GetVarOrSetFuncLock(key any, f func) *gvar.Var`
  GetVarOrSetFuncLock returns a Var with result from GetOrSetFuncLock.
- `func (m *AnyAnyMap) SetIfNotExist(key any, value any) bool`
  SetIfNotExist sets `value` to the map if the `key` does not exist, and then r...
- `func (m *AnyAnyMap) SetIfNotExistFunc(key any, f func) bool`
  SetIfNotExistFunc sets value with return value of callback function `f`, and ...
- `func (m *AnyAnyMap) SetIfNotExistFuncLock(key any, f func) bool`
  SetIfNotExistFuncLock sets value with return value of callback function `f`, ...
- `func (m *AnyAnyMap) Remove(key any) value any`
  Remove deletes value from map by given `key`, and return this deleted value.
- `func (m *AnyAnyMap) Removes(keys []any)`
  Removes batch deletes values of the map by keys.
- `func (m *AnyAnyMap) Keys() []any`
  Keys returns all keys of the map as a slice.
- `func (m *AnyAnyMap) Values() []any`
  Values returns all values of the map as a slice.
- `func (m *AnyAnyMap) Contains(key any) bool`
  Contains checks whether a key exists.
- `func (m *AnyAnyMap) Size() int`
  Size returns the size of the map.
- `func (m *AnyAnyMap) IsEmpty() bool`
  IsEmpty checks whether the map is empty.
- `func (m *AnyAnyMap) Clear()`
  Clear deletes all data of the map, it will remake a new underlying data map.
- `func (m *AnyAnyMap) Replace(data map[any]any)`
  Replace the data of the map with given `data`.
- `func (m *AnyAnyMap) LockFunc(f func)`
  LockFunc locks writing with given callback function `f` within RWMutex.Lock.
- `func (m *AnyAnyMap) RLockFunc(f func)`
  RLockFunc locks reading with given callback function `f` within RWMutex.RLock.
- `func (m *AnyAnyMap) Flip()`
  Flip exchanges key-value of the map to value-key.
- `func (m *AnyAnyMap) Merge(other *AnyAnyMap)`
  Merge merges two hash maps.
- `func (m *AnyAnyMap) String() string`
  String returns the map as a string.
- `func (m AnyAnyMap) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (m *AnyAnyMap) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (m *AnyAnyMap) UnmarshalValue(value any) err error`
  UnmarshalValue is an interface implement which sets any type of value for map.
- `func (m *AnyAnyMap) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.
- `func (m *AnyAnyMap) IsSubOf(other *AnyAnyMap) bool`
  IsSubOf checks whether the current map is a sub-map of `other`.
- `func (m *AnyAnyMap) Diff(other *AnyAnyMap) (addedKeys []any, removedKeys []any, updatedKeys []any)`
  Diff compares current map `m` with map `other` and returns their different keys.

### IntAnyMap

**类型**: struct

IntAnyMap implements map[int]any with RWMutex that has switch.


**方法**:

- `func (m *IntAnyMap) Iterator(f func)`
  Iterator iterates the hash map readonly with custom callback function `f`.
- `func (m *IntAnyMap) Clone(safe ...bool) *IntAnyMap`
  Clone returns a new hash map with copy of current map data.
- `func (m *IntAnyMap) Map() map[int]any`
  Map returns the underlying data map.
- `func (m *IntAnyMap) MapStrAny() map[string]any`
  MapStrAny returns a copy of the underlying data of the map as map[string]any.
- `func (m *IntAnyMap) MapCopy() map[int]any`
  MapCopy returns a copy of the underlying data of the hash map.
- `func (m *IntAnyMap) FilterEmpty()`
  FilterEmpty deletes all key-value pair of which the value is empty.
- `func (m *IntAnyMap) FilterNil()`
  FilterNil deletes all key-value pair of which the value is nil.
- `func (m *IntAnyMap) Set(key int, val any)`
  Set sets key-value to the hash map.
- `func (m *IntAnyMap) Sets(data map[int]any)`
  Sets batch sets key-values to the hash map.
- `func (m *IntAnyMap) Search(key int) (value any, found bool)`
  Search searches the map with given `key`.
- `func (m *IntAnyMap) Get(key int) value any`
  Get returns the value by given `key`.
- `func (m *IntAnyMap) Pop() (key int, value any)`
  Pop retrieves and deletes an item from the map.
- `func (m *IntAnyMap) Pops(size int) map[int]any`
  Pops retrieves and deletes `size` items from the map.
- `func (m *IntAnyMap) GetOrSet(key int, value any) any`
  GetOrSet returns the value by key,
- `func (m *IntAnyMap) GetOrSetFunc(key int, f func) any`
  GetOrSetFunc returns the value by key,
- `func (m *IntAnyMap) GetOrSetFuncLock(key int, f func) any`
  GetOrSetFuncLock returns the value by key,
- `func (m *IntAnyMap) GetVar(key int) *gvar.Var`
  GetVar returns a Var with the value by given `key`.
- `func (m *IntAnyMap) GetVarOrSet(key int, value any) *gvar.Var`
  GetVarOrSet returns a Var with result from GetVarOrSet.
- `func (m *IntAnyMap) GetVarOrSetFunc(key int, f func) *gvar.Var`
  GetVarOrSetFunc returns a Var with result from GetOrSetFunc.
- `func (m *IntAnyMap) GetVarOrSetFuncLock(key int, f func) *gvar.Var`
  GetVarOrSetFuncLock returns a Var with result from GetOrSetFuncLock.
- `func (m *IntAnyMap) SetIfNotExist(key int, value any) bool`
  SetIfNotExist sets `value` to the map if the `key` does not exist, and then r...
- `func (m *IntAnyMap) SetIfNotExistFunc(key int, f func) bool`
  SetIfNotExistFunc sets value with return value of callback function `f`, and ...
- `func (m *IntAnyMap) SetIfNotExistFuncLock(key int, f func) bool`
  SetIfNotExistFuncLock sets value with return value of callback function `f`, ...
- `func (m *IntAnyMap) Removes(keys []int)`
  Removes batch deletes values of the map by keys.
- `func (m *IntAnyMap) Remove(key int) value any`
  Remove deletes value from map by given `key`, and return this deleted value.
- `func (m *IntAnyMap) Keys() []int`
  Keys returns all keys of the map as a slice.
- `func (m *IntAnyMap) Values() []any`
  Values returns all values of the map as a slice.
- `func (m *IntAnyMap) Contains(key int) bool`
  Contains checks whether a key exists.
- `func (m *IntAnyMap) Size() int`
  Size returns the size of the map.
- `func (m *IntAnyMap) IsEmpty() bool`
  IsEmpty checks whether the map is empty.
- `func (m *IntAnyMap) Clear()`
  Clear deletes all data of the map, it will remake a new underlying data map.
- `func (m *IntAnyMap) Replace(data map[int]any)`
  Replace the data of the map with given `data`.
- `func (m *IntAnyMap) LockFunc(f func)`
  LockFunc locks writing with given callback function `f` within RWMutex.Lock.
- `func (m *IntAnyMap) RLockFunc(f func)`
  RLockFunc locks reading with given callback function `f` within RWMutex.RLock.
- `func (m *IntAnyMap) Flip()`
  Flip exchanges key-value of the map to value-key.
- `func (m *IntAnyMap) Merge(other *IntAnyMap)`
  Merge merges two hash maps.
- `func (m *IntAnyMap) String() string`
  String returns the map as a string.
- `func (m IntAnyMap) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (m *IntAnyMap) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (m *IntAnyMap) UnmarshalValue(value any) err error`
  UnmarshalValue is an interface implement which sets any type of value for map.
- `func (m *IntAnyMap) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.
- `func (m *IntAnyMap) IsSubOf(other *IntAnyMap) bool`
  IsSubOf checks whether the current map is a sub-map of `other`.
- `func (m *IntAnyMap) Diff(other *IntAnyMap) (addedKeys []int, removedKeys []int, updatedKeys []int)`
  Diff compares current map `m` with map `other` and returns their different keys.

### IntIntMap

**类型**: struct

IntIntMap implements map[int]int with RWMutex that has switch.


**方法**:

- `func (m *IntIntMap) Iterator(f func)`
  Iterator iterates the hash map readonly with custom callback function `f`.
- `func (m *IntIntMap) Clone(safe ...bool) *IntIntMap`
  Clone returns a new hash map with copy of current map data.
- `func (m *IntIntMap) Map() map[int]int`
  Map returns the underlying data map.
- `func (m *IntIntMap) MapStrAny() map[string]any`
  MapStrAny returns a copy of the underlying data of the map as map[string]any.
- `func (m *IntIntMap) MapCopy() map[int]int`
  MapCopy returns a copy of the underlying data of the hash map.
- `func (m *IntIntMap) FilterEmpty()`
  FilterEmpty deletes all key-value pair of which the value is empty.
- `func (m *IntIntMap) Set(key int, val int)`
  Set sets key-value to the hash map.
- `func (m *IntIntMap) Sets(data map[int]int)`
  Sets batch sets key-values to the hash map.
- `func (m *IntIntMap) Search(key int) (value int, found bool)`
  Search searches the map with given `key`.
- `func (m *IntIntMap) Get(key int) value int`
  Get returns the value by given `key`.
- `func (m *IntIntMap) Pop() (key int, value int)`
  Pop retrieves and deletes an item from the map.
- `func (m *IntIntMap) Pops(size int) map[int]int`
  Pops retrieves and deletes `size` items from the map.
- `func (m *IntIntMap) GetOrSet(key int, value int) int`
  GetOrSet returns the value by key,
- `func (m *IntIntMap) GetOrSetFunc(key int, f func) int`
  GetOrSetFunc returns the value by key,
- `func (m *IntIntMap) GetOrSetFuncLock(key int, f func) int`
  GetOrSetFuncLock returns the value by key,
- `func (m *IntIntMap) SetIfNotExist(key int, value int) bool`
  SetIfNotExist sets `value` to the map if the `key` does not exist, and then r...
- `func (m *IntIntMap) SetIfNotExistFunc(key int, f func) bool`
  SetIfNotExistFunc sets value with return value of callback function `f`, and ...
- `func (m *IntIntMap) SetIfNotExistFuncLock(key int, f func) bool`
  SetIfNotExistFuncLock sets value with return value of callback function `f`, ...
- `func (m *IntIntMap) Removes(keys []int)`
  Removes batch deletes values of the map by keys.
- `func (m *IntIntMap) Remove(key int) value int`
  Remove deletes value from map by given `key`, and return this deleted value.
- `func (m *IntIntMap) Keys() []int`
  Keys returns all keys of the map as a slice.
- `func (m *IntIntMap) Values() []int`
  Values returns all values of the map as a slice.
- `func (m *IntIntMap) Contains(key int) bool`
  Contains checks whether a key exists.
- `func (m *IntIntMap) Size() int`
  Size returns the size of the map.
- `func (m *IntIntMap) IsEmpty() bool`
  IsEmpty checks whether the map is empty.
- `func (m *IntIntMap) Clear()`
  Clear deletes all data of the map, it will remake a new underlying data map.
- `func (m *IntIntMap) Replace(data map[int]int)`
  Replace the data of the map with given `data`.
- `func (m *IntIntMap) LockFunc(f func)`
  LockFunc locks writing with given callback function `f` within RWMutex.Lock.
- `func (m *IntIntMap) RLockFunc(f func)`
  RLockFunc locks reading with given callback function `f` within RWMutex.RLock.
- `func (m *IntIntMap) Flip()`
  Flip exchanges key-value of the map to value-key.
- `func (m *IntIntMap) Merge(other *IntIntMap)`
  Merge merges two hash maps.
- `func (m *IntIntMap) String() string`
  String returns the map as a string.
- `func (m IntIntMap) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (m *IntIntMap) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (m *IntIntMap) UnmarshalValue(value any) err error`
  UnmarshalValue is an interface implement which sets any type of value for map.
- `func (m *IntIntMap) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.
- `func (m *IntIntMap) IsSubOf(other *IntIntMap) bool`
  IsSubOf checks whether the current map is a sub-map of `other`.
- `func (m *IntIntMap) Diff(other *IntIntMap) (addedKeys []int, removedKeys []int, updatedKeys []int)`
  Diff compares current map `m` with map `other` and returns their different keys.

### IntStrMap

**类型**: struct

IntStrMap implements map[int]string with RWMutex that has switch.


**方法**:

- `func (m *IntStrMap) Iterator(f func)`
  Iterator iterates the hash map readonly with custom callback function `f`.
- `func (m *IntStrMap) Clone(safe ...bool) *IntStrMap`
  Clone returns a new hash map with copy of current map data.
- `func (m *IntStrMap) Map() map[int]string`
  Map returns the underlying data map.
- `func (m *IntStrMap) MapStrAny() map[string]any`
  MapStrAny returns a copy of the underlying data of the map as map[string]any.
- `func (m *IntStrMap) MapCopy() map[int]string`
  MapCopy returns a copy of the underlying data of the hash map.
- `func (m *IntStrMap) FilterEmpty()`
  FilterEmpty deletes all key-value pair of which the value is empty.
- `func (m *IntStrMap) Set(key int, val string)`
  Set sets key-value to the hash map.
- `func (m *IntStrMap) Sets(data map[int]string)`
  Sets batch sets key-values to the hash map.
- `func (m *IntStrMap) Search(key int) (value string, found bool)`
  Search searches the map with given `key`.
- `func (m *IntStrMap) Get(key int) value string`
  Get returns the value by given `key`.
- `func (m *IntStrMap) Pop() (key int, value string)`
  Pop retrieves and deletes an item from the map.
- `func (m *IntStrMap) Pops(size int) map[int]string`
  Pops retrieves and deletes `size` items from the map.
- `func (m *IntStrMap) GetOrSet(key int, value string) string`
  GetOrSet returns the value by key,
- `func (m *IntStrMap) GetOrSetFunc(key int, f func) string`
  GetOrSetFunc returns the value by key,
- `func (m *IntStrMap) GetOrSetFuncLock(key int, f func) string`
  GetOrSetFuncLock returns the value by key,
- `func (m *IntStrMap) SetIfNotExist(key int, value string) bool`
  SetIfNotExist sets `value` to the map if the `key` does not exist, and then r...
- `func (m *IntStrMap) SetIfNotExistFunc(key int, f func) bool`
  SetIfNotExistFunc sets value with return value of callback function `f`, and ...
- `func (m *IntStrMap) SetIfNotExistFuncLock(key int, f func) bool`
  SetIfNotExistFuncLock sets value with return value of callback function `f`, ...
- `func (m *IntStrMap) Removes(keys []int)`
  Removes batch deletes values of the map by keys.
- `func (m *IntStrMap) Remove(key int) value string`
  Remove deletes value from map by given `key`, and return this deleted value.
- `func (m *IntStrMap) Keys() []int`
  Keys returns all keys of the map as a slice.
- `func (m *IntStrMap) Values() []string`
  Values returns all values of the map as a slice.
- `func (m *IntStrMap) Contains(key int) bool`
  Contains checks whether a key exists.
- `func (m *IntStrMap) Size() int`
  Size returns the size of the map.
- `func (m *IntStrMap) IsEmpty() bool`
  IsEmpty checks whether the map is empty.
- `func (m *IntStrMap) Clear()`
  Clear deletes all data of the map, it will remake a new underlying data map.
- `func (m *IntStrMap) Replace(data map[int]string)`
  Replace the data of the map with given `data`.
- `func (m *IntStrMap) LockFunc(f func)`
  LockFunc locks writing with given callback function `f` within RWMutex.Lock.
- `func (m *IntStrMap) RLockFunc(f func)`
  RLockFunc locks reading with given callback function `f` within RWMutex.RLock.
- `func (m *IntStrMap) Flip()`
  Flip exchanges key-value of the map to value-key.
- `func (m *IntStrMap) Merge(other *IntStrMap)`
  Merge merges two hash maps.
- `func (m *IntStrMap) String() string`
  String returns the map as a string.
- `func (m IntStrMap) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (m *IntStrMap) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (m *IntStrMap) UnmarshalValue(value any) err error`
  UnmarshalValue is an interface implement which sets any type of value for map.
- `func (m *IntStrMap) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.
- `func (m *IntStrMap) IsSubOf(other *IntStrMap) bool`
  IsSubOf checks whether the current map is a sub-map of `other`.
- `func (m *IntStrMap) Diff(other *IntStrMap) (addedKeys []int, removedKeys []int, updatedKeys []int)`
  Diff compares current map `m` with map `other` and returns their different keys.

### NilChecker

**类型**: type

NilChecker is a function that checks whether the given value is nil.


### KVMap

**类型**: struct

KVMap wraps map type `map[K]V` and provides more map features.


### StrAnyMap

**类型**: struct

StrAnyMap implements map[string]any with RWMutex that has switch.


**方法**:

- `func (m *StrAnyMap) Iterator(f func)`
  Iterator iterates the hash map readonly with custom callback function `f`.
- `func (m *StrAnyMap) Clone(safe ...bool) *StrAnyMap`
  Clone returns a new hash map with copy of current map data.
- `func (m *StrAnyMap) Map() map[string]any`
  Map returns the underlying data map.
- `func (m *StrAnyMap) MapStrAny() map[string]any`
  MapStrAny returns a copy of the underlying data of the map as map[string]any.
- `func (m *StrAnyMap) MapCopy() map[string]any`
  MapCopy returns a copy of the underlying data of the hash map.
- `func (m *StrAnyMap) FilterEmpty()`
  FilterEmpty deletes all key-value pair of which the value is empty.
- `func (m *StrAnyMap) FilterNil()`
  FilterNil deletes all key-value pair of which the value is nil.
- `func (m *StrAnyMap) Set(key string, val any)`
  Set sets key-value to the hash map.
- `func (m *StrAnyMap) Sets(data map[string]any)`
  Sets batch sets key-values to the hash map.
- `func (m *StrAnyMap) Search(key string) (value any, found bool)`
  Search searches the map with given `key`.
- `func (m *StrAnyMap) Get(key string) value any`
  Get returns the value by given `key`.
- `func (m *StrAnyMap) Pop() (key string, value any)`
  Pop retrieves and deletes an item from the map.
- `func (m *StrAnyMap) Pops(size int) map[string]any`
  Pops retrieves and deletes `size` items from the map.
- `func (m *StrAnyMap) GetOrSet(key string, value any) any`
  GetOrSet returns the value by key,
- `func (m *StrAnyMap) GetOrSetFunc(key string, f func) any`
  GetOrSetFunc returns the value by key,
- `func (m *StrAnyMap) GetOrSetFuncLock(key string, f func) any`
  GetOrSetFuncLock returns the value by key,
- `func (m *StrAnyMap) GetVar(key string) *gvar.Var`
  GetVar returns a Var with the value by given `key`.
- `func (m *StrAnyMap) GetVarOrSet(key string, value any) *gvar.Var`
  GetVarOrSet returns a Var with result from GetVarOrSet.
- `func (m *StrAnyMap) GetVarOrSetFunc(key string, f func) *gvar.Var`
  GetVarOrSetFunc returns a Var with result from GetOrSetFunc.
- `func (m *StrAnyMap) GetVarOrSetFuncLock(key string, f func) *gvar.Var`
  GetVarOrSetFuncLock returns a Var with result from GetOrSetFuncLock.
- `func (m *StrAnyMap) SetIfNotExist(key string, value any) bool`
  SetIfNotExist sets `value` to the map if the `key` does not exist, and then r...
- `func (m *StrAnyMap) SetIfNotExistFunc(key string, f func) bool`
  SetIfNotExistFunc sets value with return value of callback function `f`, and ...
- `func (m *StrAnyMap) SetIfNotExistFuncLock(key string, f func) bool`
  SetIfNotExistFuncLock sets value with return value of callback function `f`, ...
- `func (m *StrAnyMap) Removes(keys []string)`
  Removes batch deletes values of the map by keys.
- `func (m *StrAnyMap) Remove(key string) value any`
  Remove deletes value from map by given `key`, and return this deleted value.
- `func (m *StrAnyMap) Keys() []string`
  Keys returns all keys of the map as a slice.
- `func (m *StrAnyMap) Values() []any`
  Values returns all values of the map as a slice.
- `func (m *StrAnyMap) Contains(key string) bool`
  Contains checks whether a key exists.
- `func (m *StrAnyMap) Size() int`
  Size returns the size of the map.
- `func (m *StrAnyMap) IsEmpty() bool`
  IsEmpty checks whether the map is empty.
- `func (m *StrAnyMap) Clear()`
  Clear deletes all data of the map, it will remake a new underlying data map.
- `func (m *StrAnyMap) Replace(data map[string]any)`
  Replace the data of the map with given `data`.
- `func (m *StrAnyMap) LockFunc(f func)`
  LockFunc locks writing with given callback function `f` within RWMutex.Lock.
- `func (m *StrAnyMap) RLockFunc(f func)`
  RLockFunc locks reading with given callback function `f` within RWMutex.RLock.
- `func (m *StrAnyMap) Flip()`
  Flip exchanges key-value of the map to value-key.
- `func (m *StrAnyMap) Merge(other *StrAnyMap)`
  Merge merges two hash maps.
- `func (m *StrAnyMap) String() string`
  String returns the map as a string.
- `func (m StrAnyMap) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (m *StrAnyMap) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (m *StrAnyMap) UnmarshalValue(value any) err error`
  UnmarshalValue is an interface implement which sets any type of value for map.
- `func (m *StrAnyMap) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.
- `func (m *StrAnyMap) IsSubOf(other *StrAnyMap) bool`
  IsSubOf checks whether the current map is a sub-map of `other`.
- `func (m *StrAnyMap) Diff(other *StrAnyMap) (addedKeys []string, removedKeys []string, updatedKeys []string)`
  Diff compares current map `m` with map `other` and returns their different keys.

### StrIntMap

**类型**: struct

StrIntMap implements map[string]int with RWMutex that has switch.


**方法**:

- `func (m *StrIntMap) Iterator(f func)`
  Iterator iterates the hash map readonly with custom callback function `f`.
- `func (m *StrIntMap) Clone(safe ...bool) *StrIntMap`
  Clone returns a new hash map with copy of current map data.
- `func (m *StrIntMap) Map() map[string]int`
  Map returns the underlying data map.
- `func (m *StrIntMap) MapStrAny() map[string]any`
  MapStrAny returns a copy of the underlying data of the map as map[string]any.
- `func (m *StrIntMap) MapCopy() map[string]int`
  MapCopy returns a copy of the underlying data of the hash map.
- `func (m *StrIntMap) FilterEmpty()`
  FilterEmpty deletes all key-value pair of which the value is empty.
- `func (m *StrIntMap) Set(key string, val int)`
  Set sets key-value to the hash map.
- `func (m *StrIntMap) Sets(data map[string]int)`
  Sets batch sets key-values to the hash map.
- `func (m *StrIntMap) Search(key string) (value int, found bool)`
  Search searches the map with given `key`.
- `func (m *StrIntMap) Get(key string) value int`
  Get returns the value by given `key`.
- `func (m *StrIntMap) Pop() (key string, value int)`
  Pop retrieves and deletes an item from the map.
- `func (m *StrIntMap) Pops(size int) map[string]int`
  Pops retrieves and deletes `size` items from the map.
- `func (m *StrIntMap) GetOrSet(key string, value int) int`
  GetOrSet returns the value by key,
- `func (m *StrIntMap) GetOrSetFunc(key string, f func) int`
  GetOrSetFunc returns the value by key,
- `func (m *StrIntMap) GetOrSetFuncLock(key string, f func) int`
  GetOrSetFuncLock returns the value by key,
- `func (m *StrIntMap) SetIfNotExist(key string, value int) bool`
  SetIfNotExist sets `value` to the map if the `key` does not exist, and then r...
- `func (m *StrIntMap) SetIfNotExistFunc(key string, f func) bool`
  SetIfNotExistFunc sets value with return value of callback function `f`, and ...
- `func (m *StrIntMap) SetIfNotExistFuncLock(key string, f func) bool`
  SetIfNotExistFuncLock sets value with return value of callback function `f`, ...
- `func (m *StrIntMap) Removes(keys []string)`
  Removes batch deletes values of the map by keys.
- `func (m *StrIntMap) Remove(key string) value int`
  Remove deletes value from map by given `key`, and return this deleted value.
- `func (m *StrIntMap) Keys() []string`
  Keys returns all keys of the map as a slice.
- `func (m *StrIntMap) Values() []int`
  Values returns all values of the map as a slice.
- `func (m *StrIntMap) Contains(key string) bool`
  Contains checks whether a key exists.
- `func (m *StrIntMap) Size() int`
  Size returns the size of the map.
- `func (m *StrIntMap) IsEmpty() bool`
  IsEmpty checks whether the map is empty.
- `func (m *StrIntMap) Clear()`
  Clear deletes all data of the map, it will remake a new underlying data map.
- `func (m *StrIntMap) Replace(data map[string]int)`
  Replace the data of the map with given `data`.
- `func (m *StrIntMap) LockFunc(f func)`
  LockFunc locks writing with given callback function `f` within RWMutex.Lock.
- `func (m *StrIntMap) RLockFunc(f func)`
  RLockFunc locks reading with given callback function `f` within RWMutex.RLock.
- `func (m *StrIntMap) Flip()`
  Flip exchanges key-value of the map to value-key.
- `func (m *StrIntMap) Merge(other *StrIntMap)`
  Merge merges two hash maps.
- `func (m *StrIntMap) String() string`
  String returns the map as a string.
- `func (m StrIntMap) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (m *StrIntMap) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (m *StrIntMap) UnmarshalValue(value any) err error`
  UnmarshalValue is an interface implement which sets any type of value for map.
- `func (m *StrIntMap) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.
- `func (m *StrIntMap) IsSubOf(other *StrIntMap) bool`
  IsSubOf checks whether the current map is a sub-map of `other`.
- `func (m *StrIntMap) Diff(other *StrIntMap) (addedKeys []string, removedKeys []string, updatedKeys []string)`
  Diff compares current map `m` with map `other` and returns their different keys.

### StrStrMap

**类型**: struct

StrStrMap implements map[string]string with RWMutex that has switch.


**方法**:

- `func (m *StrStrMap) Iterator(f func)`
  Iterator iterates the hash map readonly with custom callback function `f`.
- `func (m *StrStrMap) Clone(safe ...bool) *StrStrMap`
  Clone returns a new hash map with copy of current map data.
- `func (m *StrStrMap) Map() map[string]string`
  Map returns the underlying data map.
- `func (m *StrStrMap) MapStrAny() map[string]any`
  MapStrAny returns a copy of the underlying data of the map as map[string]any.
- `func (m *StrStrMap) MapCopy() map[string]string`
  MapCopy returns a copy of the underlying data of the hash map.
- `func (m *StrStrMap) FilterEmpty()`
  FilterEmpty deletes all key-value pair of which the value is empty.
- `func (m *StrStrMap) Set(key string, val string)`
  Set sets key-value to the hash map.
- `func (m *StrStrMap) Sets(data map[string]string)`
  Sets batch sets key-values to the hash map.
- `func (m *StrStrMap) Search(key string) (value string, found bool)`
  Search searches the map with given `key`.
- `func (m *StrStrMap) Get(key string) value string`
  Get returns the value by given `key`.
- `func (m *StrStrMap) Pop() (key string, value string)`
  Pop retrieves and deletes an item from the map.
- `func (m *StrStrMap) Pops(size int) map[string]string`
  Pops retrieves and deletes `size` items from the map.
- `func (m *StrStrMap) GetOrSet(key string, value string) string`
  GetOrSet returns the value by key,
- `func (m *StrStrMap) GetOrSetFunc(key string, f func) string`
  GetOrSetFunc returns the value by key,
- `func (m *StrStrMap) GetOrSetFuncLock(key string, f func) string`
  GetOrSetFuncLock returns the value by key,
- `func (m *StrStrMap) SetIfNotExist(key string, value string) bool`
  SetIfNotExist sets `value` to the map if the `key` does not exist, and then r...
- `func (m *StrStrMap) SetIfNotExistFunc(key string, f func) bool`
  SetIfNotExistFunc sets value with return value of callback function `f`, and ...
- `func (m *StrStrMap) SetIfNotExistFuncLock(key string, f func) bool`
  SetIfNotExistFuncLock sets value with return value of callback function `f`, ...
- `func (m *StrStrMap) Removes(keys []string)`
  Removes batch deletes values of the map by keys.
- `func (m *StrStrMap) Remove(key string) value string`
  Remove deletes value from map by given `key`, and return this deleted value.
- `func (m *StrStrMap) Keys() []string`
  Keys returns all keys of the map as a slice.
- `func (m *StrStrMap) Values() []string`
  Values returns all values of the map as a slice.
- `func (m *StrStrMap) Contains(key string) bool`
  Contains checks whether a key exists.
- `func (m *StrStrMap) Size() int`
  Size returns the size of the map.
- `func (m *StrStrMap) IsEmpty() bool`
  IsEmpty checks whether the map is empty.
- `func (m *StrStrMap) Clear()`
  Clear deletes all data of the map, it will remake a new underlying data map.
- `func (m *StrStrMap) Replace(data map[string]string)`
  Replace the data of the map with given `data`.
- `func (m *StrStrMap) LockFunc(f func)`
  LockFunc locks writing with given callback function `f` within RWMutex.Lock.
- `func (m *StrStrMap) RLockFunc(f func)`
  RLockFunc locks reading with given callback function `f` within RWMutex.RLock.
- `func (m *StrStrMap) Flip()`
  Flip exchanges key-value of the map to value-key.
- `func (m *StrStrMap) Merge(other *StrStrMap)`
  Merge merges two hash maps.
- `func (m *StrStrMap) String() string`
  String returns the map as a string.
- `func (m StrStrMap) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (m *StrStrMap) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (m *StrStrMap) UnmarshalValue(value any) err error`
  UnmarshalValue is an interface implement which sets any type of value for map.
- `func (m *StrStrMap) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.
- `func (m *StrStrMap) IsSubOf(other *StrStrMap) bool`
  IsSubOf checks whether the current map is a sub-map of `other`.
- `func (m *StrStrMap) Diff(other *StrStrMap) (addedKeys []string, removedKeys []string, updatedKeys []string)`
  Diff compares current map `m` with map `other` and returns their different keys.

### ListKVMap

**类型**: struct

ListKVMap is a map that preserves insertion-order.

It is backed by a hash table to store values and doubly-linked list to store ordering.

Thread-safety is optional and controlled by the `safe` parameter during initialization.

Reference: http://en.wikipedia.org/wiki/Associative_array


### ListMap

**类型**: struct

ListMap is a map that preserves insertion-order.

It is backed by a hash table to store values and doubly-linked list to store ordering.

Structure is not thread safe.

Reference: http://en.wikipedia.org/wiki/Associative_array


**方法**:

- `func (m *ListMap) Iterator(f func)`
  Iterator is alias of IteratorAsc.
- `func (m *ListMap) IteratorAsc(f func)`
  IteratorAsc iterates the map readonly in ascending order with given callback ...
- `func (m *ListMap) IteratorDesc(f func)`
  IteratorDesc iterates the map readonly in descending order with given callbac...
- `func (m *ListMap) Clone(safe ...bool) *ListMap`
  Clone returns a new link map with copy of current map data.
- `func (m *ListMap) Clear()`
  Clear deletes all data of the map, it will remake a new underlying data map.
- `func (m *ListMap) Replace(data map[any]any)`
  Replace the data of the map with given `data`.
- `func (m *ListMap) Map() map[any]any`
  Map returns a copy of the underlying data of the map.
- `func (m *ListMap) MapStrAny() map[string]any`
  MapStrAny returns a copy of the underlying data of the map as map[string]any.
- `func (m *ListMap) FilterEmpty()`
  FilterEmpty deletes all key-value pair of which the value is empty.
- `func (m *ListMap) Set(key any, value any)`
  Set sets key-value to the map.
- `func (m *ListMap) Sets(data map[any]any)`
  Sets batch sets key-values to the map.
- `func (m *ListMap) Search(key any) (value any, found bool)`
  Search searches the map with given `key`.
- `func (m *ListMap) Get(key any) value any`
  Get returns the value by given `key`.
- `func (m *ListMap) Pop() (key any, value any)`
  Pop retrieves and deletes an item from the map.
- `func (m *ListMap) Pops(size int) map[any]any`
  Pops retrieves and deletes `size` items from the map.
- `func (m *ListMap) GetOrSet(key any, value any) any`
  GetOrSet returns the value by key,
- `func (m *ListMap) GetOrSetFunc(key any, f func) any`
  GetOrSetFunc returns the value by key,
- `func (m *ListMap) GetOrSetFuncLock(key any, f func) any`
  GetOrSetFuncLock returns the value by key,
- `func (m *ListMap) GetVar(key any) *gvar.Var`
  GetVar returns a Var with the value by given `key`.
- `func (m *ListMap) GetVarOrSet(key any, value any) *gvar.Var`
  GetVarOrSet returns a Var with result from GetVarOrSet.
- `func (m *ListMap) GetVarOrSetFunc(key any, f func) *gvar.Var`
  GetVarOrSetFunc returns a Var with result from GetOrSetFunc.
- `func (m *ListMap) GetVarOrSetFuncLock(key any, f func) *gvar.Var`
  GetVarOrSetFuncLock returns a Var with result from GetOrSetFuncLock.
- `func (m *ListMap) SetIfNotExist(key any, value any) bool`
  SetIfNotExist sets `value` to the map if the `key` does not exist, and then r...
- `func (m *ListMap) SetIfNotExistFunc(key any, f func) bool`
  SetIfNotExistFunc sets value with return value of callback function `f`, and ...
- `func (m *ListMap) SetIfNotExistFuncLock(key any, f func) bool`
  SetIfNotExistFuncLock sets value with return value of callback function `f`, ...
- `func (m *ListMap) Remove(key any) value any`
  Remove deletes value from map by given `key`, and return this deleted value.
- `func (m *ListMap) Removes(keys []any)`
  Removes batch deletes values of the map by keys.
- `func (m *ListMap) Keys() []any`
  Keys returns all keys of the map as a slice in ascending order.
- `func (m *ListMap) Values() []any`
  Values returns all values of the map as a slice.
- `func (m *ListMap) Contains(key any) ok bool`
  Contains checks whether a key exists.
- `func (m *ListMap) Size() size int`
  Size returns the size of the map.
- `func (m *ListMap) IsEmpty() bool`
  IsEmpty checks whether the map is empty.
- `func (m *ListMap) Flip()`
  Flip exchanges key-value of the map to value-key.
- `func (m *ListMap) Merge(other *ListMap)`
  Merge merges two link maps.
- `func (m *ListMap) String() string`
  String returns the map as a string.
- `func (m ListMap) MarshalJSON() (jsonBytes []byte, err error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (m *ListMap) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (m *ListMap) UnmarshalValue(value any) err error`
  UnmarshalValue is an interface implement which sets any type of value for map.
- `func (m *ListMap) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### TreeKVMap

**类型**: type

TreeKVMap based on red-black tree, alias of RedBlackKVTree.


### TreeMap

**类型**: type

TreeMap based on red-black tree, alias of RedBlackTree.


## 函数

### New

```go
func New(safe ...bool) *Map
```

New creates and returns an empty hash map.
The parameter `safe` is used to specify whether using map in concurrent-safety,
which is false in default.

### NewFrom

```go
func NewFrom(data map[any]any, safe ...bool) *Map
```

NewFrom creates and returns a hash map from given map `data`.
Note that, the param `data` map will be set as the underlying data map(no deep copy),
there might be some concurrent-safe issues when changing the map outside.
The parameter `safe` is used to specify whether using tree in concurrent-safety,
which is false in default.

### NewHashMap

```go
func NewHashMap(safe ...bool) *Map
```

NewHashMap creates and returns an empty hash map.
The parameter `safe` is used to specify whether using map in concurrent-safety,
which is false in default.

### NewHashMapFrom

```go
func NewHashMapFrom(data map[any]any, safe ...bool) *Map
```

NewHashMapFrom creates and returns a hash map from given map `data`.
Note that, the param `data` map will be set as the underlying data map(no deep copy),
there might be some concurrent-safe issues when changing the map outside.
The parameter `safe` is used to specify whether using tree in concurrent-safety,
which is false in default.

### NewAnyAnyMap

```go
func NewAnyAnyMap(safe ...bool) *AnyAnyMap
```

NewAnyAnyMap creates and returns an empty hash map.
The parameter `safe` is used to specify whether using map in concurrent-safety,
which is false in default.

### NewAnyAnyMapFrom

```go
func NewAnyAnyMapFrom(data map[any]any, safe ...bool) *AnyAnyMap
```

NewAnyAnyMapFrom creates and returns a hash map from given map `data`.
Note that, the param `data` map will be set as the underlying data map(no deep copy),
there might be some concurrent-safe issues when changing the map outside.

### NewIntAnyMap

```go
func NewIntAnyMap(safe ...bool) *IntAnyMap
```

NewIntAnyMap returns an empty IntAnyMap object.
The parameter `safe` is used to specify whether using map in concurrent-safety,
which is false in default.

### NewIntAnyMapFrom

```go
func NewIntAnyMapFrom(data map[int]any, safe ...bool) *IntAnyMap
```

NewIntAnyMapFrom creates and returns a hash map from given map `data`.
Note that, the param `data` map will be set as the underlying data map(no deep copy),
there might be some concurrent-safe issues when changing the map outside.

### NewIntIntMap

```go
func NewIntIntMap(safe ...bool) *IntIntMap
```

NewIntIntMap returns an empty IntIntMap object.
The parameter `safe` is used to specify whether using map in concurrent-safety,
which is false in default.

### NewIntIntMapFrom

```go
func NewIntIntMapFrom(data map[int]int, safe ...bool) *IntIntMap
```

NewIntIntMapFrom creates and returns a hash map from given map `data`.
Note that, the param `data` map will be set as the underlying data map(no deep copy),
there might be some concurrent-safe issues when changing the map outside.

### NewIntStrMap

```go
func NewIntStrMap(safe ...bool) *IntStrMap
```

NewIntStrMap returns an empty IntStrMap object.
The parameter `safe` is used to specify whether using map in concurrent-safety,
which is false in default.

### NewIntStrMapFrom

```go
func NewIntStrMapFrom(data map[int]string, safe ...bool) *IntStrMap
```

NewIntStrMapFrom creates and returns a hash map from given map `data`.
Note that, the param `data` map will be set as the underlying data map(no deep copy),
there might be some concurrent-safe issues when changing the map outside.

### NewKVMap

```go
func NewKVMap(safe ...bool) *
```

NewKVMap creates and returns an empty hash map.
The parameter `safe` is used to specify whether to use the map in concurrent-safety mode, which is false by default.

### NewKVMapWithChecker

```go
func NewKVMapWithChecker(checker , safe ...bool) *
```

NewKVMapWithChecker creates and returns an empty hash map with a custom nil checker.
The parameter `checker` is a function used to determine if a value is nil.
The parameter `safe` is used to specify whether to use the map in concurrent-safety mode, which is false by default.

### NewKVMapFrom

```go
func NewKVMapFrom(data map[K]V, safe ...bool) *
```

NewKVMapFrom creates and returns a hash map from given map `data`.
Note that, the param `data` map will be set as the underlying data map (no deep copy),
there might be some concurrent-safe issues when changing the map outside.

### NewKVMapWithCheckerFrom

```go
func NewKVMapWithCheckerFrom(data map[K]V, checker , safe ...bool) *
```

NewKVMapWithCheckerFrom creates and returns a hash map from given map `data` with a custom nil checker.
Note that, the param `data` map will be set as the underlying data map (no deep copy),
and there might be some concurrent-safe issues when changing the map outside.
The parameter `checker` is a function used to determine if a value is nil.
The parameter `safe` is used to specify whether to use the map in concurrent-safety mode, which is false by default.

### MarshalJSON

```go
func (m ) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

### NewStrAnyMap

```go
func NewStrAnyMap(safe ...bool) *StrAnyMap
```

NewStrAnyMap returns an empty StrAnyMap object.
The parameter `safe` is used to specify whether using map in concurrent-safety,
which is false in default.

### NewStrAnyMapFrom

```go
func NewStrAnyMapFrom(data map[string]any, safe ...bool) *StrAnyMap
```

NewStrAnyMapFrom creates and returns a hash map from given map `data`.
Note that, the param `data` map will be set as the underlying data map(no deep copy),
there might be some concurrent-safe issues when changing the map outside.

### NewStrIntMap

```go
func NewStrIntMap(safe ...bool) *StrIntMap
```

NewStrIntMap returns an empty StrIntMap object.
The parameter `safe` is used to specify whether using map in concurrent-safety,
which is false in default.

### NewStrIntMapFrom

```go
func NewStrIntMapFrom(data map[string]int, safe ...bool) *StrIntMap
```

NewStrIntMapFrom creates and returns a hash map from given map `data`.
Note that, the param `data` map will be set as the underlying data map(no deep copy),
there might be some concurrent-safe issues when changing the map outside.

### NewStrStrMap

```go
func NewStrStrMap(safe ...bool) *StrStrMap
```

NewStrStrMap returns an empty StrStrMap object.
The parameter `safe` is used to specify whether using map in concurrent-safety,
which is false in default.

### NewStrStrMapFrom

```go
func NewStrStrMapFrom(data map[string]string, safe ...bool) *StrStrMap
```

NewStrStrMapFrom creates and returns a hash map from given map `data`.
Note that, the param `data` map will be set as the underlying data map(no deep copy),
there might be some concurrent-safe issues when changing the map outside.

### NewListKVMap

```go
func NewListKVMap(safe ...bool) *
```

NewListKVMap returns an empty link map.
ListKVMap is backed by a hash table to store values and doubly-linked list to store ordering.
The parameter `safe` is used to specify whether using map in concurrent-safety,
which is false in default.

### NewListKVMapWithChecker

```go
func NewListKVMapWithChecker(checker , safe ...bool) *
```

NewListKVMapWithChecker creates and returns a new ListKVMap instance with a custom nil checker.
The parameter `checker` is a function used to determine if a value is nil.
The parameter `safe` is used to specify whether using map in concurrent-safety,
which is false by default.

### NewListKVMapFrom

```go
func NewListKVMapFrom(data map[K]V, safe ...bool) *
```

NewListKVMapFrom returns a link map from given map `data`.
Note that, the param `data` map will be copied to the underlying data structure,
so changes to the original map will not affect the link map.

### NewListKVMapWithCheckerFrom

```go
func NewListKVMapWithCheckerFrom(data map[K]V, nilChecker , safe ...bool) *
```

NewListKVMapWithCheckerFrom returns a link map from given map `data` with a custom nil checker.
Note that, the param `data` map will be copied to the underlying data structure,
so changes to the original map will not affect the link map.
The parameter `checker` is a function used to determine if a value is nil.
The parameter `safe` is used to specify whether using map in concurrent-safety,
which is false by default.

### MarshalJSON

```go
func (m ) MarshalJSON() (jsonBytes []byte, err error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

### NewListMap

```go
func NewListMap(safe ...bool) *ListMap
```

NewListMap returns an empty link map.
ListMap is backed by a hash table to store values and doubly-linked list to store ordering.
The parameter `safe` is used to specify whether using map in concurrent-safety,
which is false in default.

### NewListMapFrom

```go
func NewListMapFrom(data map[any]any, safe ...bool) *ListMap
```

NewListMapFrom returns a link map from given map `data`.
Note that, the param `data` map will be set as the underlying data map(no deep copy),
there might be some concurrent-safe issues when changing the map outside.

### NewTreeKVMap

```go
func NewTreeKVMap(comparator func, safe ...bool) *
```

NewTreeKVMap instantiates a tree map with the custom comparator.
The parameter `safe` is used to specify whether using tree in concurrent-safety,
which is false in default.

### NewTreeKVMapFrom

```go
func NewTreeKVMapFrom(comparator func, data map[K]V, safe ...bool) *
```

NewTreeKVMapFrom instantiates a tree map with the custom comparator and `data` map.
Note that, the param `data` map will be set as the underlying data map(no deep copy),
there might be some concurrent-safe issues when changing the map outside.
The parameter `safe` is used to specify whether using tree in concurrent-safety,
which is false in default.

### NewTreeMap

```go
func NewTreeMap(comparator func, safe ...bool) *TreeMap
```

NewTreeMap instantiates a tree map with the custom comparator.
The parameter `safe` is used to specify whether using tree in concurrent-safety,
which is false in default.

### NewTreeMapFrom

```go
func NewTreeMapFrom(comparator func, data map[any]any, safe ...bool) *TreeMap
```

NewTreeMapFrom instantiates a tree map with the custom comparator and `data` map.
Note that, the param `data` map will be set as the underlying data map(no deep copy),
there might be some concurrent-safe issues when changing the map outside.
The parameter `safe` is used to specify whether using tree in concurrent-safety,
which is false in default.

## 内部依赖

- `container/glist`
- `container/gtree`
- `container/gvar`
- `internal/deepcopy`
- `internal/empty`
- `internal/json`
- `internal/rwmutex`
- `util/gconv`

