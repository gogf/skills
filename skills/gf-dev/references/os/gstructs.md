# gstructs

> Package: `github.com/gogf/gf/v2/os/gstructs`

```go
import "github.com/gogf/gf/v2/os/gstructs"
```

## 概述

Package gstructs provides functions for struct information retrieving.

## 源文件

- `gstructs.go`
- `gstructs_field.go`
- `gstructs_field_tag.go`
- `gstructs_tag.go`
- `gstructs_type.go`

## 类型定义

### Type

**类型**: struct

Type wraps reflect.Type for additional features.


**方法**:

- `func (t Type) Signature() string`
  Signature returns a unique string as this type.
- `func (t Type) FieldKeys() []string`
  FieldKeys returns the keys of current struct/map.

### Field

**类型**: struct

Field contains information of a struct field .


**方法**:

- `func (f *Field) Tag(key string) string`
  Tag returns the value associated with key in the tag string. If there is no
- `func (f *Field) TagLookup(key string) (value string, ok bool)`
  TagLookup returns the value associated with key in the tag string.
- `func (f *Field) IsEmbedded() bool`
  IsEmbedded returns true if the given field is an anonymous field (embedded)
- `func (f *Field) TagStr() string`
  TagStr returns the tag string of the field.
- `func (f *Field) TagMap() map[string]string`
  TagMap returns all the tag of the field along with its value string as map.
- `func (f *Field) IsExported() bool`
  IsExported returns true if the given field is exported.
- `func (f *Field) Name() string`
  Name returns the name of the given field.
- `func (f *Field) Type() Type`
  Type returns the type of the given field.
- `func (f *Field) Kind() reflect.Kind`
  Kind returns the reflect.Kind for Value of Field `f`.
- `func (f *Field) OriginalKind() reflect.Kind`
  OriginalKind retrieves and returns the original reflect.Kind for Value of Fie...
- `func (f *Field) OriginalValue() reflect.Value`
  OriginalValue retrieves and returns the original reflect.Value of Field `f`.
- `func (f *Field) IsEmpty() bool`
  IsEmpty checks and returns whether the value of this Field is empty.
- `func (f *Field) IsNil(traceSource ...bool) bool`
  IsNil checks and returns whether the value of this Field is nil.
- `func (f *Field) TagJsonName() string`
  TagJsonName returns the `json` tag name string of the field.
- `func (f *Field) TagDefault() string`
  TagDefault returns the most commonly used tag `default/d` value of the field.
- `func (f *Field) TagParam() string`
  TagParam returns the most commonly used tag `param/p` value of the field.
- `func (f *Field) TagValid() string`
  TagValid returns the most commonly used tag `valid/v` value of the field.
- `func (f *Field) TagDescription() string`
  TagDescription returns the most commonly used tag `description/des/dc` value ...
- `func (f *Field) TagSummary() string`
  TagSummary returns the most commonly used tag `summary/sum/sm` value of the f...
- `func (f *Field) TagAdditional() string`
  TagAdditional returns the most commonly used tag `additional/ad` value of the...
- `func (f *Field) TagExample() string`
  TagExample returns the most commonly used tag `example/eg` value of the field.
- `func (f *Field) TagIn() string`
  TagIn returns the most commonly used tag `in` value of the field.
- `func (f *Field) TagPriorityName() string`
  TagPriorityName checks and returns tag name that matches the name item in `gt...

### FieldsInput

**类型**: struct

FieldsInput is the input parameter struct type for function Fields.


### FieldMapInput

**类型**: struct

FieldMapInput is the input parameter struct type for function FieldMap.


### RecursiveOption

**类型**: type

## 函数

### Fields

```go
func Fields(in FieldsInput) ([]Field, error)
```

Fields retrieves and returns the fields of `pointer` as slice.

### FieldMap

```go
func FieldMap(in FieldMapInput) (map[string]Field, error)
```

FieldMap retrieves and returns struct field as map[name/tag]Field from `pointer`.

The parameter `pointer` should be type of struct/*struct.

The parameter `priority` specifies the priority tag array for retrieving from high to low.
If it's given `nil`, it returns map[name]Field, of which the `name` is attribute name.

The parameter `recursive` specifies whether retrieving the fields recursively if the attribute
is an embedded struct.

Note that it only retrieves the exported attributes with first letter upper-case from struct.

### StructType

```go
func StructType(object any) (*Type, error)
```

StructType retrieves and returns the struct Type of specified struct/*struct.
The parameter `object` should be either type of struct/*struct/[]struct/[]*struct.

### ParseTag

```go
func ParseTag(tag string) map[string]string
```

ParseTag parses tag string into map.
For example:
ParseTag(`v:"required" p:"id" d:"1"`) => map[v:required p:id d:1].

### TagFields

```go
func TagFields(pointer any, priority []string) ([]Field, error)
```

TagFields retrieves and returns struct tags as []Field from `pointer`.

The parameter `pointer` should be type of struct/*struct.

Note that,
1. It only retrieves the exported attributes with first letter upper-case from struct.
2. The parameter `priority` should be given, it only retrieves fields that has given tag.

### TagMapName

```go
func TagMapName(pointer any, priority []string) (map[string]string, error)
```

TagMapName retrieves and returns struct tags as map[tag]attribute from `pointer`.

The parameter `pointer` should be type of struct/*struct.

Note that,
1. It only retrieves the exported attributes with first letter upper-case from struct.
2. The parameter `priority` should be given, it only retrieves fields that has given tag.
3. If one field has no specified tag, it uses its field name as result map key.

### TagMapField

```go
func TagMapField(object any, priority []string) (map[string]Field, error)
```

TagMapField retrieves struct tags as map[tag]Field from `pointer`, and returns it.
The parameter `object` should be either type of struct/*struct/[]struct/[]*struct.

Note that,
1. It only retrieves the exported attributes with first letter upper-case from struct.
2. The parameter `priority` should be given, it only retrieves fields that has given tag.
3. If one field has no specified tag, it uses its field name as result map key.

## 内部依赖

- `errors/gcode`
- `errors/gerror`
- `internal/empty`
- `internal/utils`
- `util/gtag`

