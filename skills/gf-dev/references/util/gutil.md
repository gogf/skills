# gutil

> Package: `github.com/gogf/gf/v2/util/gutil`

```go
import "github.com/gogf/gf/v2/util/gutil"
```

## 概述

Package gutil provides utility functions.

## 源文件

- `gutil.go`
- `gutil_comparator.go`
- `gutil_copy.go`
- `gutil_default.go`
- `gutil_dump.go`
- `gutil_goroutine.go`
- `gutil_is.go`
- `gutil_list.go`
- `gutil_map.go`
- `gutil_reflect.go`
- `gutil_slice.go`
- `gutil_struct.go`
- `gutil_try_catch.go`

## 类型定义

### Comparator

**类型**: type

Comparator is a function that compare a and b, and returns the result as int.

Should return a number:

	negative , if a < b
	zero     , if a == b
	positive , if a > b


### DumpOption

**类型**: struct

DumpOption specifies the behavior of function Export.


### OriginValueAndKindOutput

**类型**: type

### OriginTypeAndKindOutput

**类型**: type

## 函数

### Keys

```go
func Keys(mapOrStruct any) keysOrAttrs []string
```

Keys retrieves and returns the keys from the given map or struct.

### Values

```go
func Values(mapOrStruct any) values []any
```

Values retrieves and returns the values from the given map or struct.

### ComparatorString

```go
func ComparatorString(a any, b any) int
```

ComparatorString provides a fast comparison on strings.

### ComparatorInt

```go
func ComparatorInt(a any, b any) int
```

ComparatorInt provides a basic comparison on int.

### ComparatorInt8

```go
func ComparatorInt8(a any, b any) int
```

ComparatorInt8 provides a basic comparison on int8.

### ComparatorInt16

```go
func ComparatorInt16(a any, b any) int
```

ComparatorInt16 provides a basic comparison on int16.

### ComparatorInt32

```go
func ComparatorInt32(a any, b any) int
```

ComparatorInt32 provides a basic comparison on int32.

### ComparatorInt64

```go
func ComparatorInt64(a any, b any) int
```

ComparatorInt64 provides a basic comparison on int64.

### ComparatorUint

```go
func ComparatorUint(a any, b any) int
```

ComparatorUint provides a basic comparison on uint.

### ComparatorUint8

```go
func ComparatorUint8(a any, b any) int
```

ComparatorUint8 provides a basic comparison on uint8.

### ComparatorUint16

```go
func ComparatorUint16(a any, b any) int
```

ComparatorUint16 provides a basic comparison on uint16.

### ComparatorUint32

```go
func ComparatorUint32(a any, b any) int
```

ComparatorUint32 provides a basic comparison on uint32.

### ComparatorUint64

```go
func ComparatorUint64(a any, b any) int
```

ComparatorUint64 provides a basic comparison on uint64.

### ComparatorFloat32

```go
func ComparatorFloat32(a any, b any) int
```

ComparatorFloat32 provides a basic comparison on float32.

### ComparatorFloat64

```go
func ComparatorFloat64(a any, b any) int
```

ComparatorFloat64 provides a basic comparison on float64.

### ComparatorByte

```go
func ComparatorByte(a any, b any) int
```

ComparatorByte provides a basic comparison on byte.

### ComparatorRune

```go
func ComparatorRune(a any, b any) int
```

ComparatorRune provides a basic comparison on rune.

### ComparatorTime

```go
func ComparatorTime(a any, b any) int
```

ComparatorTime provides a basic comparison on time.Time.

### ComparatorT

```go
func ComparatorT(a T, b T) int
```

ComparatorT provides a generic comparison for ordered types.

### ComparatorTStr

```go
func ComparatorTStr(a T, b T) int
```

### Copy

```go
func Copy(src any) dst any
```

Copy returns a deep copy of v.

Copy is unable to copy unexported fields in a struct (lowercase field names).
Unexported fields can't be reflected by the Go runtime and therefore
they can't perform any data copies.

### GetOrDefaultStr

```go
func GetOrDefaultStr(def string, param ...string) string
```

GetOrDefaultStr checks and returns value according whether parameter `param` available.
It returns `param[0]` if it is available, or else it returns `def`.

### GetOrDefaultAny

```go
func GetOrDefaultAny(def any, param ...any) any
```

GetOrDefaultAny checks and returns value according whether parameter `param` available.
It returns `param[0]` if it is available, or else it returns `def`.

### Dump

```go
func Dump(values ...any)
```

Dump prints variables `values` to stdout with more manually readable.

### DumpWithType

```go
func DumpWithType(values ...any)
```

DumpWithType acts like Dump, but with type information.
Also see Dump.

### DumpWithOption

```go
func DumpWithOption(value any, option DumpOption)
```

DumpWithOption returns variables `values` as a string with more manually readable.

### DumpTo

```go
func DumpTo(writer io.Writer, value any, option DumpOption)
```

DumpTo writes variables `values` as a string in to `writer` with more manually readable

### DumpJson

```go
func DumpJson(value any)
```

DumpJson pretty dumps json content to stdout.

### Go

```go
func Go(ctx context.Context, goroutineFunc func, recoverFunc func)
```

Go creates a new asynchronous goroutine function with specified recover function.

The parameter `recoverFunc` is called when any panic during executing of `goroutineFunc`.
If `recoverFunc` is given nil, it ignores the panic from `goroutineFunc` and no panic will
throw to parent goroutine.

But, note that, if `recoverFunc` also throws panic, such panic will be thrown to parent goroutine.

### IsEmpty

```go
func IsEmpty(value any) bool
```

IsEmpty checks given `value` empty or not.
It returns false if `value` is: integer(0), bool(false), slice/map(len=0), nil;
or else returns true.

### IsTypeOf

```go
func IsTypeOf(value any, valueInExpectType any) bool
```

IsTypeOf checks and returns whether the type of `value` and `valueInExpectType` equal.

### ListItemValues

```go
func ListItemValues(list any, key any, subKey ...any) values []any
```

ListItemValues retrieves and returns the elements of all item struct/map with key `key`.
Note that the parameter `list` should be type of slice which contains elements of map or struct,
or else it returns an empty slice.

The parameter `list` supports types like:
[]map[string]any
[]map[string]sub-map
[]struct
[]struct:sub-struct
Note that the sub-map/sub-struct makes sense only if the optional parameter `subKey` is given.

### ItemValue

```go
func ItemValue(item any, key any) (value any, found bool)
```

ItemValue retrieves and returns its value of which name/attribute specified by `key`.
The parameter `item` can be type of map/*map/struct/*struct.

### ListItemValuesUnique

```go
func ListItemValuesUnique(list any, key string, subKey ...any) []any
```

ListItemValuesUnique retrieves and returns the unique elements of all struct/map with key `key`.
Note that the parameter `list` should be type of slice which contains elements of map or struct,
or else it returns an empty slice.

### ListToMapByKey

```go
func ListToMapByKey(list []map[string]any, key string) map[string]any
```

ListToMapByKey converts `list` to a map[string]any of which key is specified by `key`.
Note that the item value may be type of slice.

### MapCopy

```go
func MapCopy(data map[string]any) copy map[string]any
```

MapCopy does a shallow copy from map `data` to `copy` for most commonly used map type
map[string]any.

### MapContains

```go
func MapContains(data map[string]any, key string) ok bool
```

MapContains checks whether map `data` contains `key`.

### MapDelete

```go
func MapDelete(data map[string]any, keys ...string)
```

MapDelete deletes all `keys` from map `data`.

### MapMerge

```go
func MapMerge(dstMap map[string]any, srcMaps ...map[string]any)
```

MapMerge merges all map from `src` to map `dst`.

### MapMergeCopy

```go
func MapMergeCopy(maps ...map[string]any) copy map[string]any
```

MapMergeCopy creates and returns a new map which merges all map from `src`.

### MapPossibleItemByKey

```go
func MapPossibleItemByKey(data map[string]any, key string) (foundKey string, foundValue any)
```

MapPossibleItemByKey tries to find the possible key-value pair for given key ignoring cases and symbols.

Note that this function might be of low performance.

### MapContainsPossibleKey

```go
func MapContainsPossibleKey(data map[string]any, key string) bool
```

MapContainsPossibleKey checks if the given `key` is contained in given map `data`.
It checks the key ignoring cases and symbols.

Note that this function might be of low performance.

### MapOmitEmpty

```go
func MapOmitEmpty(data map[string]any)
```

MapOmitEmpty deletes all empty values from given map.

### MapToSlice

```go
func MapToSlice(data any) []any
```

MapToSlice converts map to slice of which all keys and values are its items.
Eg: {"K1": "v1", "K2": "v2"} => ["K1", "v1", "K2", "v2"]

### OriginValueAndKind

```go
func OriginValueAndKind(value any) out OriginValueAndKindOutput
```

OriginValueAndKind retrieves and returns the original reflect value and kind.

### OriginTypeAndKind

```go
func OriginTypeAndKind(value any) out OriginTypeAndKindOutput
```

OriginTypeAndKind retrieves and returns the original reflect type and kind.

### SliceCopy

```go
func SliceCopy(slice []any) []any
```

SliceCopy does a shallow copy of slice `data` for most commonly used slice type
[]any.

### SliceInsertBefore

```go
func SliceInsertBefore(slice []any, index int, values ...any) newSlice []any
```

SliceInsertBefore inserts the `values` to the front of `index` and returns a new slice.

### SliceInsertAfter

```go
func SliceInsertAfter(slice []any, index int, values ...any) newSlice []any
```

SliceInsertAfter inserts the `values` to the back of `index` and returns a new slice.

### SliceDelete

```go
func SliceDelete(slice []any, index int) newSlice []any
```

SliceDelete deletes an element at `index` and returns the new slice.
It does nothing if the given `index` is invalid.

### SliceToMap

```go
func SliceToMap(slice any) map[string]any
```

SliceToMap converts slice type variable `slice` to `map[string]any`.
Note that if the length of `slice` is not an even number, it returns nil.
Eg:
["K1", "v1", "K2", "v2"] => {"K1": "v1", "K2": "v2"}
["K1", "v1", "K2"]       => nil

### SliceToMapWithColumnAsKey

```go
func SliceToMapWithColumnAsKey(slice any, key any) map[any]any
```

SliceToMapWithColumnAsKey converts slice type variable `slice` to `map[any]any`
The value of specified column use as the key for returned map.
Eg:
SliceToMapWithColumnAsKey([{"K1": "v1", "K2": 1}, {"K1": "v2", "K2": 2}], "K1") => {"v1": {"K1": "v1", "K2": 1}, "v2": {"K1": "v2", "K2": 2}}
SliceToMapWithColumnAsKey([{"K1": "v1", "K2": 1}, {"K1": "v2", "K2": 2}], "K2") => {1: {"K1": "v1", "K2": 1}, 2: {"K1": "v2", "K2": 2}}

### StructToSlice

```go
func StructToSlice(data any) []any
```

StructToSlice converts struct to slice of which all keys and values are its items.
Eg: {"K1": "v1", "K2": "v2"} => ["K1", "v1", "K2", "v2"]

### FillStructWithDefault

```go
func FillStructWithDefault(structPtr any) error
```

FillStructWithDefault fills  attributes of pointed struct with tag value from `default/d` tag .
The parameter `structPtr` should be either type of *struct/[]*struct.

### Throw

```go
func Throw(exception any)
```

Throw throws out an exception, which can be caught be TryCatch or recover.

### Try

```go
func Try(ctx context.Context, try func) err error
```

Try implements try... logistics using internal panic...recover.
It returns error if any exception occurs, or else it returns nil.

### TryCatch

```go
func TryCatch(ctx context.Context, try func, catch func)
```

TryCatch implements `try...catch..`. logistics using internal `panic...recover`.
It automatically calls function `catch` if any exception occurs and passes the exception as an error.
If `catch` is given nil, it ignores the panic from `try` and no panic will throw to parent goroutine.

But, note that, if function `catch` also throws panic, the current goroutine will panic.

## 内部依赖

- `errors/gcode`
- `errors/gerror`
- `internal/deepcopy`
- `internal/empty`
- `internal/reflection`
- `internal/utils`
- `os/gstructs`
- `text/gstr`
- `util/gconv`

