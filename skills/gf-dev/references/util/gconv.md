# gconv

> Package: `github.com/gogf/gf/v2/util/gconv`

```go
import "github.com/gogf/gf/v2/util/gconv"
```

## 概述

Package gconv implements powerful and convenient converting functionality for any types of variables.

## 源文件

- `gconv.go`
- `gconv_basic.go`
- `gconv_convert.go`
- `gconv_float.go`
- `gconv_int.go`
- `gconv_map.go`
- `gconv_maps.go`
- `gconv_maptomap.go`
- `gconv_maptomaps.go`
- `gconv_ptr.go`
- `gconv_scan.go`
- `gconv_scan_list.go`
- `gconv_slice_any.go`
- `gconv_slice_bool.go`
- `gconv_slice_float.go`
- `gconv_slice_int.go`
- `gconv_slice_str.go`
- `gconv_slice_uint.go`
- `gconv_struct.go`
- `gconv_structs.go`
- `gconv_time.go`
- `gconv_uint.go`
- `gconv_unsafe.go`

## 类型定义

### Converter

**类型**: interface

Converter is the manager for type converting.


### ConverterForBasic

**类型**: interface

ConverterForBasic is the basic converting interface.


### ConverterForTime

**类型**: interface

ConverterForTime is the converting interface for time.


### ConverterForInt

**类型**: interface

ConverterForInt is the converting interface for integer.


### ConverterForUint

**类型**: interface

ConverterForUint is the converting interface for unsigned integer.


### ConverterForFloat

**类型**: interface

ConverterForFloat is the converting interface for float.


### ConverterForMap

**类型**: interface

ConverterForMap is the converting interface for map.


### ConverterForSlice

**类型**: interface

ConverterForSlice is the converting interface for slice.


### ConverterForStruct

**类型**: interface

ConverterForStruct is the converting interface for struct.


### ConverterForConvert

**类型**: interface

ConverterForConvert is the converting interface for custom converting.


### ConverterForRegister

**类型**: interface

ConverterForRegister is the converting interface for custom converter registration.


### AnyConvertFunc

**类型**: type

### MapOption

**类型**: type

### SliceOption

**类型**: type

### SliceMapOption

**类型**: type

### ScanOption

**类型**: type

### StructOption

**类型**: type

### StructsOption

**类型**: type

### ConvertOption

**类型**: type

### IUnmarshalValue

**类型**: type

IUnmarshalValue is the interface for custom defined types customizing value assignment.
Note that only pointer can implement interface IUnmarshalValue.


## 函数

### NewConverter

```go
func NewConverter() Converter
```

NewConverter creates and returns management object for type converting.

### RegisterConverter

```go
func RegisterConverter(fn any) err error
```

RegisterConverter registers custom converter.

Deprecated: use RegisterTypeConverterFunc instead for clear

### RegisterTypeConverterFunc

```go
func RegisterTypeConverterFunc(fn any) err error
```

RegisterTypeConverterFunc registers custom converter.

### RegisterAnyConverterFunc

```go
func RegisterAnyConverterFunc(f AnyConvertFunc, types ...reflect.Type)
```

RegisterAnyConverterFunc registers custom type converting function for specified type.

### Byte

```go
func Byte(anyInput any) byte
```

Byte converts `any` to byte.

### Bytes

```go
func Bytes(anyInput any) []byte
```

Bytes converts `any` to []byte.

### Rune

```go
func Rune(anyInput any) rune
```

Rune converts `any` to rune.

### Runes

```go
func Runes(anyInput any) []rune
```

Runes converts `any` to []rune.

### String

```go
func String(anyInput any) string
```

String converts `any` to string.
It's most commonly used converting function.

### Bool

```go
func Bool(anyInput any) bool
```

Bool converts `any` to bool.
It returns false if `any` is: false, "", 0, "false", "off", "no", empty slice/map.

### Convert

```go
func Convert(fromValue any, toTypeName string, extraParams ...any) any
```

Convert converts the variable `fromValue` to the type `toTypeName`, the type `toTypeName` is specified by string.

The optional parameter `extraParams` is used for additional necessary parameter for this conversion.
It supports common basic types conversion as its conversion based on type name string.

### ConvertWithRefer

```go
func ConvertWithRefer(fromValue any, referValue any, extraParams ...any) any
```

ConvertWithRefer converts the variable `fromValue` to the type referred by value `referValue`.

The optional parameter `extraParams` is used for additional necessary parameter for this conversion.
It supports common basic types conversion as its conversion based on type name string.

### Float32

```go
func Float32(anyInput any) float32
```

Float32 converts `any` to float32.

### Float64

```go
func Float64(anyInput any) float64
```

Float64 converts `any` to float64.

### Int

```go
func Int(anyInput any) int
```

Int converts `any` to int.

### Int8

```go
func Int8(anyInput any) int8
```

Int8 converts `any` to int8.

### Int16

```go
func Int16(anyInput any) int16
```

Int16 converts `any` to int16.

### Int32

```go
func Int32(anyInput any) int32
```

Int32 converts `any` to int32.

### Int64

```go
func Int64(anyInput any) int64
```

Int64 converts `any` to int64.

### Map

```go
func Map(value any, option ...MapOption) map[string]any
```

Map converts any variable `value` to map[string]any. If the parameter `value` is not a
map/struct/*struct type, then the conversion will fail and returns nil.

If `value` is a struct/*struct object, the second parameter `priorityTagAndFieldName` specifies the most priority
priorityTagAndFieldName that will be detected, otherwise it detects the priorityTagAndFieldName in order of:
gconv, json, field name.

### MapDeep

```go
func MapDeep(value any, tags ...string) map[string]any
```

MapDeep does Map function recursively, which means if the attribute of `value`
is also a struct/*struct, calls Map function on this attribute converting it to
a map[string]any type variable.

Deprecated: used Map instead.

### MapStrStr

```go
func MapStrStr(value any, option ...MapOption) map[string]string
```

MapStrStr converts `value` to map[string]string.
Note that there might be data copy for this map type converting.

### MapStrStrDeep

```go
func MapStrStrDeep(value any, tags ...string) map[string]string
```

MapStrStrDeep converts `value` to map[string]string recursively.
Note that there might be data copy for this map type converting.

Deprecated: used MapStrStr instead.

### SliceMap

```go
func SliceMap(anyInput any, option ...MapOption) []map[string]any
```

SliceMap is alias of Maps.

### SliceMapDeep

```go
func SliceMapDeep(anyInput any) []map[string]any
```

SliceMapDeep is alias of MapsDeep.

Deprecated: used SliceMap instead.

### Maps

```go
func Maps(value any, option ...MapOption) []map[string]any
```

Maps converts `value` to []map[string]any.
Note that it automatically checks and converts json string to []map if `value` is string/[]byte.

### MapsDeep

```go
func MapsDeep(value any, tags ...string) []map[string]any
```

MapsDeep converts `value` to []map[string]any recursively.

TODO completely implement the recursive converting for all types.

Deprecated: used Maps instead.

### MapToMap

```go
func MapToMap(params any, pointer any, mapping ...map[string]string) error
```

MapToMap converts any map type variable `params` to another map type variable `pointer`
using reflect.
See doMapToMap.

### MapToMaps

```go
func MapToMaps(params any, pointer any, mapping ...map[string]string) error
```

MapToMaps converts any slice type variable `params` to another map slice type variable `pointer`.
See doMapToMaps.

### PtrAny

```go
func PtrAny(anyInput any) *any
```

PtrAny creates and returns an any pointer variable to this value.

### PtrString

```go
func PtrString(anyInput any) *string
```

PtrString creates and returns a string pointer variable to this value.

### PtrBool

```go
func PtrBool(anyInput any) *bool
```

PtrBool creates and returns a bool pointer variable to this value.

### PtrInt

```go
func PtrInt(anyInput any) *int
```

PtrInt creates and returns an int pointer variable to this value.

### PtrInt8

```go
func PtrInt8(anyInput any) *int8
```

PtrInt8 creates and returns an int8 pointer variable to this value.

### PtrInt16

```go
func PtrInt16(anyInput any) *int16
```

PtrInt16 creates and returns an int16 pointer variable to this value.

### PtrInt32

```go
func PtrInt32(anyInput any) *int32
```

PtrInt32 creates and returns an int32 pointer variable to this value.

### PtrInt64

```go
func PtrInt64(anyInput any) *int64
```

PtrInt64 creates and returns an int64 pointer variable to this value.

### PtrUint

```go
func PtrUint(anyInput any) *uint
```

PtrUint creates and returns an uint pointer variable to this value.

### PtrUint8

```go
func PtrUint8(anyInput any) *uint8
```

PtrUint8 creates and returns an uint8 pointer variable to this value.

### PtrUint16

```go
func PtrUint16(anyInput any) *uint16
```

PtrUint16 creates and returns an uint16 pointer variable to this value.

### PtrUint32

```go
func PtrUint32(anyInput any) *uint32
```

PtrUint32 creates and returns an uint32 pointer variable to this value.

### PtrUint64

```go
func PtrUint64(anyInput any) *uint64
```

PtrUint64 creates and returns an uint64 pointer variable to this value.

### PtrFloat32

```go
func PtrFloat32(anyInput any) *float32
```

PtrFloat32 creates and returns a float32 pointer variable to this value.

### PtrFloat64

```go
func PtrFloat64(anyInput any) *float64
```

PtrFloat64 creates and returns a float64 pointer variable to this value.

### Scan

```go
func Scan(srcValue any, dstPointer any, paramKeyToAttrMap ...map[string]string) err error
```

Scan automatically checks the type of `pointer` and converts `params` to `pointer`.
It supports various types of parameter conversions, including:
1. Basic types (int, string, float, etc.)
2. Pointer types
3. Slice types
4. Map types
5. Struct types

The `paramKeyToAttrMap` parameter is used for mapping between attribute names and parameter keys.
TODO: change `paramKeyToAttrMap` to `ScanOption` to be more scalable; add `DeepCopy` option for `ScanOption`.

### ScanWithOptions

```go
func ScanWithOptions(srcValue any, dstPointer any, option ...ScanOption) err error
```

ScanWithOptions automatically checks the type of `dstPointer` and converts `srcValue` to `dstPointer`.
It is the same as Scan function, but accepts one or more ScanOption values for additional conversion control.

When using ScanWithOptions, the term "omit" means that the assignment from the source to the destination
is skipped, so the existing value in the destination field is preserved.

  - option.OmitEmpty, when set to true, skips assignment of empty source values (for example: empty strings,
    zero numeric values, zero time values, empty slices or maps), preserving any existing non-empty values
    in the destination.

  - option.OmitNil, when set to true, skips assignment of nil source values, preserving the existing values
    in the destination when the source contains nil.

Example:

	type User struct {
	    Name  string
	    Email string
	}

	dst := &User{Name: "Alice", Email: "alice@example.com"}
	src := map[string]any{
	    "Name":  "",
	    "Email": nil,
	}

	// With OmitEmpty and OmitNil, empty and nil values in src will not overwrite dst.
	err := ScanWithOptions(src, dst, ScanOption{OmitEmpty: true, OmitNil: true})

### ScanList

```go
func ScanList(structSlice any, structSlicePointer any, bindToAttrName string, relationAttrNameAndFields ...string) err error
```

ScanList converts `structSlice` to struct slice which contains other complex struct attributes.
Note that the parameter `structSlicePointer` should be type of *[]struct/*[]*struct.

Usage example 1: Normal attribute struct relation:

	type EntityUser struct {
	    Uid  int
	    Name string
	}

	type EntityUserDetail struct {
	    Uid     int
	    Address string
	}

	type EntityUserScores struct {
	    Id     int
	    Uid    int
	    Score  int
	    Course string
	}

	type Entity struct {
	    User       *EntityUser
	    UserDetail *EntityUserDetail
	    UserScores []*EntityUserScores
	}

var users []*Entity
var userRecords   = EntityUser{Uid: 1, Name:"john"}
var detailRecords = EntityUser{Uid: 1, Address: "chengdu"}
var scoresRecords = EntityUser{Id: 1, Uid: 1, Score: 100, Course: "math"}
ScanList(userRecords, &users, "User")
ScanList(userRecords, &users, "User", "uid")
ScanList(detailRecords, &users, "UserDetail", "User", "uid:Uid")
ScanList(scoresRecords, &users, "UserScores", "User", "uid:Uid")
ScanList(scoresRecords, &users, "UserScores", "User", "uid")

Usage example 2: Embedded attribute struct relation:

	type EntityUser struct {
		   Uid  int
		   Name string
	}

	type EntityUserDetail struct {
		   Uid     int
		   Address string
	}

	type EntityUserScores struct {
		   Id    int
		   Uid   int
		   Score int
	}

	type Entity struct {
		   EntityUser
		   UserDetail EntityUserDetail
		   UserScores []EntityUserScores
	}

var userRecords   = EntityUser{Uid: 1, Name:"john"}
var detailRecords = EntityUser{Uid: 1, Address: "chengdu"}
var scoresRecords = EntityUser{Id: 1, Uid: 1, Score: 100, Course: "math"}
ScanList(userRecords, &users)
ScanList(detailRecords, &users, "UserDetail", "uid")
ScanList(scoresRecords, &users, "UserScores", "uid")

The parameters "User/UserDetail/UserScores" in the example codes specify the target attribute struct
that current result will be bound to.

The "uid" in the example codes is the table field name of the result, and the "Uid" is the relational
struct attribute name - not the attribute name of the bound to target. In the example codes, it's attribute
name "Uid" of "User" of entity "Entity". It automatically calculates the HasOne/HasMany relationship with
given `relation` parameter.

See the example or unit testing cases for clear understanding for this function.

### SliceAny

```go
func SliceAny(anyInput any) []any
```

SliceAny is alias of Interfaces.

### Interfaces

```go
func Interfaces(anyInput any) []any
```

Interfaces converts `any` to []any.

### SliceBool

```go
func SliceBool(anyInput any) []bool
```

SliceBool is alias of Bools.

### Bools

```go
func Bools(anyInput any) []bool
```

Bools converts `any` to []bool.

### SliceFloat

```go
func SliceFloat(anyInput any) []float64
```

SliceFloat is alias of Floats.

### SliceFloat32

```go
func SliceFloat32(anyInput any) []float32
```

SliceFloat32 is alias of Float32s.

### SliceFloat64

```go
func SliceFloat64(anyInput any) []float64
```

SliceFloat64 is alias of Float64s.

### Floats

```go
func Floats(anyInput any) []float64
```

Floats converts `any` to []float64.

### Float32s

```go
func Float32s(anyInput any) []float32
```

Float32s converts `any` to []float32.

### Float64s

```go
func Float64s(anyInput any) []float64
```

Float64s converts `any` to []float64.

### SliceInt

```go
func SliceInt(anyInput any) []int
```

SliceInt is alias of Ints.

### SliceInt32

```go
func SliceInt32(anyInput any) []int32
```

SliceInt32 is alias of Int32s.

### SliceInt64

```go
func SliceInt64(anyInput any) []int64
```

SliceInt64 is alias of Int64s.

### Ints

```go
func Ints(anyInput any) []int
```

Ints converts `any` to []int.

### Int32s

```go
func Int32s(anyInput any) []int32
```

Int32s converts `any` to []int32.

### Int64s

```go
func Int64s(anyInput any) []int64
```

Int64s converts `any` to []int64.

### SliceStr

```go
func SliceStr(anyInput any) []string
```

SliceStr is alias of Strings.

### Strings

```go
func Strings(anyInput any) []string
```

Strings converts `any` to []string.

### SliceUint

```go
func SliceUint(anyInput any) []uint
```

SliceUint is alias of Uints.

### SliceUint32

```go
func SliceUint32(anyInput any) []uint32
```

SliceUint32 is alias of Uint32s.

### SliceUint64

```go
func SliceUint64(anyInput any) []uint64
```

SliceUint64 is alias of Uint64s.

### Uints

```go
func Uints(anyInput any) []uint
```

Uints converts `any` to []uint.

### Uint32s

```go
func Uint32s(anyInput any) []uint32
```

Uint32s converts `any` to []uint32.

### Uint64s

```go
func Uint64s(anyInput any) []uint64
```

Uint64s converts `any` to []uint64.

### Struct

```go
func Struct(params any, pointer any, paramKeyToAttrMap ...map[string]string) err error
```

Struct maps the params key-value pairs to the corresponding struct object's attributes.
The third parameter `mapping` is unnecessary, indicating the mapping rules between the
custom key name and the attribute name(case-sensitive).

Note:
 1. The `params` can be any type of map/struct, usually a map.
 2. The `pointer` should be type of *struct/**struct, which is a pointer to struct object
    or struct pointer.
 3. Only the public attributes of struct object can be mapped.
 4. If `params` is a map, the key of the map `params` can be lowercase.
    It will automatically convert the first letter of the key to uppercase
    in mapping procedure to do the matching.
    It ignores the map key, if it does not match.

### StructTag

```go
func StructTag(params any, pointer any, priorityTag string) err error
```

StructTag acts as Struct but also with support for priority tag feature, which retrieves the
specified priorityTagAndFieldName for `params` key-value items to struct attribute names mapping.
The parameter `priorityTag` supports multiple priorityTagAndFieldName that can be joined with char ','.

### Structs

```go
func Structs(params any, pointer any, paramKeyToAttrMap ...map[string]string) err error
```

Structs converts any slice to given struct slice.
Also see Scan, Struct.

### SliceStruct

```go
func SliceStruct(params any, pointer any, mapping ...map[string]string) err error
```

SliceStruct is alias of Structs.

### StructsTag

```go
func StructsTag(params any, pointer any, priorityTag string) err error
```

StructsTag acts as Structs but also with support for priority tag feature, which retrieves the
specified priorityTagAndFieldName for `params` key-value items to struct attribute names mapping.
The parameter `priorityTag` supports multiple priorityTagAndFieldName that can be joined with char ','.

### Time

```go
func Time(anyInput any, format ...string) time.Time
```

Time converts `any` to time.Time.

### Duration

```go
func Duration(anyInput any) time.Duration
```

Duration converts `any` to time.Duration.
If `any` is string, then it uses time.ParseDuration to convert it.
If `any` is numeric, then it converts `any` as nanoseconds.

### GTime

```go
func GTime(anyInput any, format ...string) *gtime.Time
```

GTime converts `any` to *gtime.Time.
The parameter `format` can be used to specify the format of `any`.
It returns the converted value that matched the first format of the formats slice.
If no `format` given, it converts `any` using gtime.NewFromTimeStamp if `any` is numeric,
or using gtime.StrToTime if `any` is string.

### Uint

```go
func Uint(anyInput any) uint
```

Uint converts `any` to uint.

### Uint8

```go
func Uint8(anyInput any) uint8
```

Uint8 converts `any` to uint8.

### Uint16

```go
func Uint16(anyInput any) uint16
```

Uint16 converts `any` to uint16.

### Uint32

```go
func Uint32(anyInput any) uint32
```

Uint32 converts `any` to uint32.

### Uint64

```go
func Uint64(anyInput any) uint64
```

Uint64 converts `any` to uint64.

### UnsafeStrToBytes

```go
func UnsafeStrToBytes(s string) []byte
```

UnsafeStrToBytes converts string to []byte without memory copy.
Note that, if you completely sure you will never use `s` variable in the feature,
you can use this unsafe function to implement type conversion in high performance.

### UnsafeBytesToStr

```go
func UnsafeBytesToStr(b []byte) string
```

UnsafeBytesToStr converts []byte to string without memory copy.
Note that, if you completely sure you will never use `b` variable in the feature,
you can use this unsafe function to implement type conversion in high performance.

## 内部依赖

- `errors/gcode`
- `errors/gerror`
- `internal/json`
- `internal/utils`
- `os/gstructs`
- `os/gtime`
- `util/gconv/internal/converter`
- `util/gconv/internal/localinterface`
- `util/gconv/internal/structcache`

