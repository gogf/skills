# gjson

> Package: `github.com/gogf/gf/v2/encoding/gjson`

```go
import "github.com/gogf/gf/v2/encoding/gjson"
```

## 概述

Package gjson provides convenient API for JSON/XML/INI/YAML/TOML data handling.

## 源文件

- `gjson.go`
- `gjson_api.go`
- `gjson_api_config.go`
- `gjson_api_encoding.go`
- `gjson_api_new_load.go`
- `gjson_api_new_load_content.go`
- `gjson_api_new_load_path.go`
- `gjson_implements.go`
- `gjson_stdlib_json_util.go`

## 类型定义

### ContentType

**类型**: type

### Json

**类型**: struct

Json is the customized JSON struct.


**方法**:

- `func (j *Json) Interface() any`
  Interface returns the json value.
- `func (j *Json) Var() *gvar.Var`
  Var returns the json value as *gvar.Var.
- `func (j *Json) IsNil() bool`
  IsNil checks whether the value pointed by `j` is nil.
- `func (j *Json) Get(pattern string, def ...any) *gvar.Var`
  Get retrieves and returns value by specified `pattern`.
- `func (j *Json) GetJson(pattern string, def ...any) *Json`
  GetJson gets the value by specified `pattern`,
- `func (j *Json) GetJsons(pattern string, def ...any) []*Json`
  GetJsons gets the value by specified `pattern`,
- `func (j *Json) GetJsonMap(pattern string, def ...any) map[string]*Json`
  GetJsonMap gets the value by specified `pattern`,
- `func (j *Json) Set(pattern string, value any) error`
  Set sets value with specified `pattern`.
- `func (j *Json) MustSet(pattern string, value any)`
  MustSet performs as Set, but it panics if any error occurs.
- `func (j *Json) Remove(pattern string) error`
  Remove deletes value with specified `pattern`.
- `func (j *Json) MustRemove(pattern string)`
  MustRemove performs as Remove, but it panics if any error occurs.
- `func (j *Json) Contains(pattern string) bool`
  Contains checks whether the value by specified `pattern` exist.
- `func (j *Json) Len(pattern string) int`
  Len returns the length/size of the value by specified `pattern`.
- `func (j *Json) Append(pattern string, value any) error`
  Append appends value to the value by specified `pattern`.
- `func (j *Json) MustAppend(pattern string, value any)`
  MustAppend performs as Append, but it panics if any error occurs.
- `func (j *Json) Map() map[string]any`
  Map converts current Json object to map[string]any.
- `func (j *Json) Array() []any`
  Array converts current Json object to []any.
- `func (j *Json) Scan(pointer any, mapping ...map[string]string) error`
  Scan automatically calls Struct or Structs function according to the type of ...
- `func (j *Json) Dump()`
  Dump prints current Json object with more manually readable.
- `func (j *Json) SetSplitChar(char byte)`
  SetSplitChar sets the separator char for hierarchical data access.
- `func (j *Json) SetViolenceCheck(enabled bool)`
  SetViolenceCheck enables/disables violence check for hierarchical data access.
- `func (j *Json) ToJson() ([]byte, error)`
- `func (j *Json) ToJsonString() (string, error)`
- `func (j *Json) ToJsonIndent() ([]byte, error)`
- `func (j *Json) ToJsonIndentString() (string, error)`
- `func (j *Json) MustToJson() []byte`
- `func (j *Json) MustToJsonString() string`
- `func (j *Json) MustToJsonIndent() []byte`
- `func (j *Json) MustToJsonIndentString() string`
- `func (j *Json) ToXml(rootTag ...string) ([]byte, error)`
- `func (j *Json) ToXmlString(rootTag ...string) (string, error)`
- `func (j *Json) ToXmlIndent(rootTag ...string) ([]byte, error)`
- `func (j *Json) ToXmlIndentString(rootTag ...string) (string, error)`
- `func (j *Json) MustToXml(rootTag ...string) []byte`
- `func (j *Json) MustToXmlString(rootTag ...string) string`
- `func (j *Json) MustToXmlIndent(rootTag ...string) []byte`
- `func (j *Json) MustToXmlIndentString(rootTag ...string) string`
- `func (j *Json) ToYaml() ([]byte, error)`
- `func (j *Json) ToYamlIndent(indent string) ([]byte, error)`
- `func (j *Json) ToYamlString() (string, error)`
- `func (j *Json) MustToYaml() []byte`
- `func (j *Json) MustToYamlString() string`
- `func (j *Json) ToToml() ([]byte, error)`
- `func (j *Json) ToTomlString() (string, error)`
- `func (j *Json) MustToToml() []byte`
- `func (j *Json) MustToTomlString() string`
- `func (j *Json) ToIni() ([]byte, error)`
  ToIni json to ini
- `func (j *Json) ToIniString() (string, error)`
  ToIniString ini to string
- `func (j *Json) MustToIni() []byte`
- `func (j *Json) MustToIniString() string`
  MustToIniString .
- `func (j *Json) ToProperties() ([]byte, error)`
  ========================================================================
- `func (j *Json) ToPropertiesString() (string, error)`
  ToPropertiesString properties to string
- `func (j *Json) MustToProperties() []byte`
- `func (j *Json) MustToPropertiesString() string`
  MustToPropertiesString
- `func (j Json) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (j *Json) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (j *Json) UnmarshalValue(value any) error`
  UnmarshalValue is an interface implement which sets any type of value for Json.
- `func (j *Json) MapStrAny() map[string]any`
  MapStrAny implements interface function MapStrAny().
- `func (j *Json) Interfaces() []any`
  Interfaces implements interface function Interfaces().
- `func (j *Json) String() string`
  String returns current Json object as string.

### Options

**类型**: struct

Options for Json object creating/loading.


## 函数

### New

```go
func New(data any, safe ...bool) *Json
```

New creates a Json object with any variable type of `data`, but `data` should be a map
or slice for data access reason, or it will make no sense.

The parameter `safe` specifies whether using this Json object in concurrent-safe context,
which is false in default.

### NewWithTag

```go
func NewWithTag(data any, tags string, safe ...bool) *Json
```

NewWithTag creates a Json object with any variable type of `data`, but `data` should be a map
or slice for data access reason, or it will make no sense.

The parameter `tags` specifies priority tags for struct conversion to map, multiple tags joined
with char ','.

The parameter `safe` specifies whether using this Json object in concurrent-safe context, which
is false in default.

### NewWithOptions

```go
func NewWithOptions(data any, options Options) *Json
```

NewWithOptions creates a Json object with any variable type of `data`, but `data` should be a map
or slice for data access reason, or it will make no sense.

### LoadWithOptions

```go
func LoadWithOptions(data []byte, options Options) (*Json, error)
```

LoadWithOptions creates a Json object from given JSON format content and options.

### LoadJson

```go
func LoadJson(data []byte, safe ...bool) (*Json, error)
```

LoadJson creates a Json object from given JSON format content.

### LoadXml

```go
func LoadXml(data []byte, safe ...bool) (*Json, error)
```

LoadXml creates a Json object from given XML format content.

### LoadIni

```go
func LoadIni(data []byte, safe ...bool) (*Json, error)
```

LoadIni creates a Json object from given INI format content.

### LoadYaml

```go
func LoadYaml(data []byte, safe ...bool) (*Json, error)
```

LoadYaml creates a Json object from given YAML format content.

### LoadToml

```go
func LoadToml(data []byte, safe ...bool) (*Json, error)
```

LoadToml creates a Json object from given TOML format content.

### LoadProperties

```go
func LoadProperties(data []byte, safe ...bool) (*Json, error)
```

LoadProperties creates a Json object from given TOML format content.

### LoadContent

```go
func LoadContent(data []byte, safe ...bool) (*Json, error)
```

LoadContent creates a Json object from given content, it checks the data type of `content`
automatically, supporting data content type as follows:
JSON, XML, INI, YAML and TOML.

### LoadContentType

```go
func LoadContentType(dataType ContentType, data []byte, safe ...bool) (*Json, error)
```

LoadContentType creates a Json object from given type and content,
supporting data content type as follows:
JSON, XML, INI, YAML and TOML.

### IsValidDataType

```go
func IsValidDataType(dataType ContentType) bool
```

IsValidDataType checks and returns whether given `dataType` a valid data type for loading.

### Load

```go
func Load(path string, safe ...bool) (*Json, error)
```

Load loads content from specified file `path`, and creates a Json object from its content.

Deprecated: use LoadPath instead.

### LoadPath

```go
func LoadPath(path string, options Options) (*Json, error)
```

LoadPath loads content from specified file `path`, and creates a Json object from its content.

### Valid

```go
func Valid(data any) bool
```

Valid checks whether `data` is a valid JSON data type.
The parameter `data` specifies the json format data, which can be either
bytes or string type.

### Marshal

```go
func Marshal(v any) (marshaledBytes []byte, err error)
```

Marshal is alias of Encode in order to fit the habit of json.Marshal/Unmarshal functions.

### MarshalIndent

```go
func MarshalIndent(v any, prefix string, indent string) (marshaledBytes []byte, err error)
```

MarshalIndent is alias of json.MarshalIndent in order to fit the habit of json.MarshalIndent function.

### Unmarshal

```go
func Unmarshal(data []byte, v any) err error
```

Unmarshal is alias of DecodeTo in order to fit the habit of json.Marshal/Unmarshal functions.

### Encode

```go
func Encode(value any) ([]byte, error)
```

Encode encodes any golang variable `value` to JSON bytes.

### MustEncode

```go
func MustEncode(value any) []byte
```

MustEncode performs as Encode, but it panics if any error occurs.

### EncodeString

```go
func EncodeString(value any) (string, error)
```

EncodeString encodes any golang variable `value` to JSON string.

### MustEncodeString

```go
func MustEncodeString(value any) string
```

MustEncodeString encodes any golang variable `value` to JSON string.
It panics if any error occurs.

### Decode

```go
func Decode(data any, options ...Options) (any, error)
```

Decode decodes json format `data` to golang variable.
The parameter `data` can be either bytes or string type.

### DecodeTo

```go
func DecodeTo(data any, v any, options ...Options) err error
```

DecodeTo decodes json format `data` to specified golang variable `v`.
The parameter `data` can be either bytes or string type.
The parameter `v` should be a pointer type.

### DecodeToJson

```go
func DecodeToJson(data any, options ...Options) (*Json, error)
```

DecodeToJson codes json format `data` to a Json object.
The parameter `data` can be either bytes or string type.

## 内部依赖

- `container/gvar`
- `encoding/gini`
- `encoding/gproperties`
- `encoding/gtoml`
- `encoding/gxml`
- `encoding/gyaml`
- `errors/gcode`
- `errors/gerror`
- `internal/json`
- `internal/reflection`
- `internal/rwmutex`
- `internal/utils`
- `os/gfile`
- `text/gregex`
- `text/gstr`
- `util/gconv`
- `util/gutil`

