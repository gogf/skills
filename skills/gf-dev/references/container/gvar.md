# gvar

> Package: `github.com/gogf/gf/v2/container/gvar`

```go
import "github.com/gogf/gf/v2/container/gvar"
```

## 概述

Package gvar provides an universal variable type, like runtime generics.

## 源文件

- `gvar.go`
- `gvar_basic.go`
- `gvar_copy.go`
- `gvar_is.go`
- `gvar_list.go`
- `gvar_map.go`
- `gvar_scan.go`
- `gvar_set.go`
- `gvar_slice.go`
- `gvar_struct.go`
- `gvar_time.go`
- `gvar_vars.go`

## 类型定义

### Var

**类型**: struct

Var is an universal variable type implementer.


**方法**:

- `func (v *Var) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (v *Var) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (v *Var) UnmarshalValue(value any) error`
  UnmarshalValue is an interface implement which sets any type of value for Var.
- `func (v *Var) Val() any`
  Val returns the current value of `v`.
- `func (v *Var) Interface() any`
  Interface is alias of Val.
- `func (v *Var) Bytes() []byte`
  Bytes converts and returns `v` as []byte.
- `func (v *Var) String() string`
  String converts and returns `v` as string.
- `func (v *Var) Bool() bool`
  Bool converts and returns `v` as bool.
- `func (v *Var) Int() int`
  Int converts and returns `v` as int.
- `func (v *Var) Int8() int8`
  Int8 converts and returns `v` as int8.
- `func (v *Var) Int16() int16`
  Int16 converts and returns `v` as int16.
- `func (v *Var) Int32() int32`
  Int32 converts and returns `v` as int32.
- `func (v *Var) Int64() int64`
  Int64 converts and returns `v` as int64.
- `func (v *Var) Uint() uint`
  Uint converts and returns `v` as uint.
- `func (v *Var) Uint8() uint8`
  Uint8 converts and returns `v` as uint8.
- `func (v *Var) Uint16() uint16`
  Uint16 converts and returns `v` as uint16.
- `func (v *Var) Uint32() uint32`
  Uint32 converts and returns `v` as uint32.
- `func (v *Var) Uint64() uint64`
  Uint64 converts and returns `v` as uint64.
- `func (v *Var) Float32() float32`
  Float32 converts and returns `v` as float32.
- `func (v *Var) Float64() float64`
  Float64 converts and returns `v` as float64.
- `func (v *Var) Copy() *Var`
  Copy does a deep copy of current Var and returns a pointer to this Var.
- `func (v *Var) Clone() *Var`
  Clone does a shallow copy of current Var and returns a pointer to this Var.
- `func (v *Var) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.
- `func (v *Var) IsNil() bool`
  IsNil checks whether `v` is nil.
- `func (v *Var) IsEmpty() bool`
  IsEmpty checks whether `v` is empty.
- `func (v *Var) IsInt() bool`
  IsInt checks whether `v` is type of int.
- `func (v *Var) IsUint() bool`
  IsUint checks whether `v` is type of uint.
- `func (v *Var) IsFloat() bool`
  IsFloat checks whether `v` is type of float.
- `func (v *Var) IsSlice() bool`
  IsSlice checks whether `v` is type of slice.
- `func (v *Var) IsMap() bool`
  IsMap checks whether `v` is type of map.
- `func (v *Var) IsStruct() bool`
  IsStruct checks whether `v` is type of struct.
- `func (v *Var) ListItemValues(key any) values []any`
  ListItemValues retrieves and returns the elements of all item struct/map with...
- `func (v *Var) ListItemValuesUnique(key string) []any`
  ListItemValuesUnique retrieves and returns the unique elements of all struct/...
- `func (v *Var) Map(option ...MapOption) map[string]any`
  Map converts and returns `v` as map[string]any.
- `func (v *Var) MapStrAny(option ...MapOption) map[string]any`
  MapStrAny is like function Map, but implements the interface of MapStrAny.
- `func (v *Var) MapStrStr(option ...MapOption) map[string]string`
  MapStrStr converts and returns `v` as map[string]string.
- `func (v *Var) MapStrVar(option ...MapOption) map[string]*Var`
  MapStrVar converts and returns `v` as map[string]Var.
- `func (v *Var) MapDeep(tags ...string) map[string]any`
  MapDeep converts and returns `v` as map[string]any recursively.
- `func (v *Var) MapStrStrDeep(tags ...string) map[string]string`
  MapStrStrDeep converts and returns `v` as map[string]string recursively.
- `func (v *Var) MapStrVarDeep(tags ...string) map[string]*Var`
  MapStrVarDeep converts and returns `v` as map[string]*Var recursively.
- `func (v *Var) Maps(option ...MapOption) []map[string]any`
  Maps converts and returns `v` as map[string]string.
- `func (v *Var) MapsDeep(tags ...string) []map[string]any`
  MapsDeep converts `value` to []map[string]any recursively.
- `func (v *Var) MapToMap(pointer any, mapping ...map[string]string) err error`
  MapToMap converts any map type variable `params` to another map type variable...
- `func (v *Var) MapToMaps(pointer any, mapping ...map[string]string) err error`
  MapToMaps converts any map type variable `params` to another map type variabl...
- `func (v *Var) MapToMapsDeep(pointer any, mapping ...map[string]string) err error`
  MapToMapsDeep converts any map type variable `params` to another map type var...
- `func (v *Var) Scan(pointer any, mapping ...map[string]string) error`
  Scan automatically checks the type of `pointer` and converts value of Var to ...
- `func (v *Var) Set(value any) old any`
  Set sets `value` to `v`, and returns the old value.
- `func (v *Var) Bools() []bool`
  Bools converts and returns `v` as []bool.
- `func (v *Var) Ints() []int`
  Ints converts and returns `v` as []int.
- `func (v *Var) Int64s() []int64`
  Int64s converts and returns `v` as []int64.
- `func (v *Var) Uints() []uint`
  Uints converts and returns `v` as []uint.
- `func (v *Var) Uint64s() []uint64`
  Uint64s converts and returns `v` as []uint64.
- `func (v *Var) Floats() []float64`
  Floats is alias of Float64s.
- `func (v *Var) Float32s() []float32`
  Float32s converts and returns `v` as []float32.
- `func (v *Var) Float64s() []float64`
  Float64s converts and returns `v` as []float64.
- `func (v *Var) Strings() []string`
  Strings converts and returns `v` as []string.
- `func (v *Var) Interfaces() []any`
  Interfaces converts and returns `v` as []interfaces{}.
- `func (v *Var) Slice() []any`
  Slice is alias of Interfaces.
- `func (v *Var) Array() []any`
  Array is alias of Interfaces.
- `func (v *Var) Vars() []*Var`
  Vars converts and returns `v` as []Var.
- `func (v *Var) Struct(pointer any, mapping ...map[string]string) error`
  Struct maps value of `v` to `pointer`.
- `func (v *Var) Structs(pointer any, mapping ...map[string]string) error`
  Structs converts and returns `v` as given struct slice.
- `func (v *Var) Time(format ...string) time.Time`
  Time converts and returns `v` as time.Time.
- `func (v *Var) Duration() time.Duration`
  Duration converts and returns `v` as time.Duration.
- `func (v *Var) GTime(format ...string) *gtime.Time`
  GTime converts and returns `v` as *gtime.Time.

### MapOption

**类型**: type

MapOption specifies the option for map converting.


### Vars

**类型**: type

Vars is a slice of *Var.


**方法**:

- `func (vs Vars) Strings() s []string`
  Strings converts and returns `vs` as []string.
- `func (vs Vars) Bools() s []bool`
  Bools converts and returns `vs` as []bool.
- `func (vs Vars) Interfaces() s []any`
  Interfaces converts and returns `vs` as []any.
- `func (vs Vars) Float32s() s []float32`
  Float32s converts and returns `vs` as []float32.
- `func (vs Vars) Float64s() s []float64`
  Float64s converts and returns `vs` as []float64.
- `func (vs Vars) Ints() s []int`
  Ints converts and returns `vs` as []Int.
- `func (vs Vars) Int8s() s []int8`
  Int8s converts and returns `vs` as []int8.
- `func (vs Vars) Int16s() s []int16`
  Int16s converts and returns `vs` as []int16.
- `func (vs Vars) Int32s() s []int32`
  Int32s converts and returns `vs` as []int32.
- `func (vs Vars) Int64s() s []int64`
  Int64s converts and returns `vs` as []int64.
- `func (vs Vars) Uints() s []uint`
  Uints converts and returns `vs` as []uint.
- `func (vs Vars) Uint8s() s []uint8`
  Uint8s converts and returns `vs` as []uint8.
- `func (vs Vars) Uint16s() s []uint16`
  Uint16s converts and returns `vs` as []uint16.
- `func (vs Vars) Uint32s() s []uint32`
  Uint32s converts and returns `vs` as []uint32.
- `func (vs Vars) Uint64s() s []uint64`
  Uint64s converts and returns `vs` as []uint64.
- `func (vs Vars) Scan(pointer any, mapping ...map[string]string) error`
  Scan converts `vs` to []struct/[]*struct.

## 函数

### New

```go
func New(value any, safe ...bool) *Var
```

New creates and returns a new Var with given `value`.
The optional parameter `safe` specifies whether Var is used in concurrent-safety,
which is false in default.

## 内部依赖

- `container/gtype`
- `internal/deepcopy`
- `internal/json`
- `internal/utils`
- `os/gtime`
- `util/gconv`
- `util/gutil`

