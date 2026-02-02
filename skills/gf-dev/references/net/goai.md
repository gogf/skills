# goai

> Package: `github.com/gogf/gf/v2/net/goai`

```go
import "github.com/gogf/gf/v2/net/goai"
```

## 概述

Package goai implements and provides document generating for OpenApi specification.

## 源文件

- `goai.go`
- `goai_callback.go`
- `goai_components.go`
- `goai_config.go`
- `goai_example.go`
- `goai_external_docs.go`
- `goai_header.go`
- `goai_info.go`
- `goai_link.go`
- `goai_mediatype.go`
- `goai_operation.go`
- `goai_parameter.go`
- `goai_parameter_ref.go`
- `goai_path.go`
- `goai_requestbody.go`
- `goai_response.go`
- `goai_response_ref.go`
- `goai_security.go`
- `goai_server.go`
- `goai_shema.go`
- `goai_shema_ref.go`
- `goai_shemas.go`
- `goai_tags.go`
- `goai_xextensions.go`

## 类型定义

### OpenApiV3

**类型**: struct

OpenApiV3 is the structure defined from:
https://swagger.io/specification/
https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md


**方法**:

- `func (oai *OpenApiV3) Add(in AddInput) error`
  Add adds an instance of struct or a route function to OpenApiV3 definition im...
- `func (oai OpenApiV3) String() string`

### AddInput

**类型**: struct

AddInput is the structured parameter for function OpenApiV3.Add.


### Callback

**类型**: type

Callback is specified by OpenAPI/Swagger standard version 3.0.


### Callbacks

**类型**: type

### CallbackRef

**类型**: struct

**方法**:

- `func (r CallbackRef) MarshalJSON() ([]byte, error)`

### Components

**类型**: struct

Components is specified by OpenAPI/Swagger standard version 3.0.


### ParametersMap

**类型**: type

### RequestBodies

**类型**: type

### Config

**类型**: struct

Config provides extra configuration feature for OpenApiV3 implements.


### Example

**类型**: struct

Example is specified by OpenAPI/Swagger 3.0 standard.


### Examples

**类型**: type

### ExampleRef

**类型**: struct

**方法**:

- `func (r ExampleRef) MarshalJSON() ([]byte, error)`

### ExternalDocs

**类型**: struct

ExternalDocs is specified by OpenAPI/Swagger standard version 3.0.


**方法**:

- `func (ed *ExternalDocs) UnmarshalValue(value any) error`

### Header

**类型**: struct

Header is specified by OpenAPI/Swagger 3.0 standard.
See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#headerObject


### Headers

**类型**: type

### HeaderRef

**类型**: struct

**方法**:

- `func (r HeaderRef) MarshalJSON() ([]byte, error)`

### Info

**类型**: struct

Info is specified by OpenAPI/Swagger standard version 3.0.


### Contact

**类型**: struct

Contact is specified by OpenAPI/Swagger standard version 3.0.


### License

**类型**: struct

License is specified by OpenAPI/Swagger standard version 3.0.


### Link

**类型**: struct

Link is specified by OpenAPI/Swagger standard version 3.0.


### Links

**类型**: type

### LinkRef

**类型**: struct

**方法**:

- `func (r LinkRef) MarshalJSON() ([]byte, error)`

### MediaType

**类型**: struct

MediaType is specified by OpenAPI/Swagger 3.0 standard.


### Content

**类型**: type

Content is specified by OpenAPI/Swagger 3.0 standard.


### Encoding

**类型**: struct

Encoding is specified by OpenAPI/Swagger 3.0 standard.


### Operation

**类型**: struct

Operation represents "operation" specified by OpenAPI/Swagger 3.0 standard.


**方法**:

- `func (o Operation) MarshalJSON() ([]byte, error)`

### Parameter

**类型**: struct

Parameter is specified by OpenAPI/Swagger 3.0 standard.
See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#parameterObject


**方法**:

- `func (p Parameter) MarshalJSON() ([]byte, error)`

### Parameters

**类型**: type

Parameters is specified by OpenAPI/Swagger 3.0 standard.


### ParameterRef

**类型**: struct

**方法**:

- `func (r ParameterRef) MarshalJSON() ([]byte, error)`

### Path

**类型**: struct

Path is specified by OpenAPI/Swagger standard version 3.0.


**方法**:

- `func (p Path) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.

### Paths

**类型**: type

Paths are specified by OpenAPI/Swagger standard version 3.0.


### RequestBody

**类型**: struct

RequestBody is specified by OpenAPI/Swagger 3.0 standard.


### RequestBodyRef

**类型**: struct

**方法**:

- `func (r RequestBodyRef) MarshalJSON() ([]byte, error)`

### EnhancedStatusCode

**类型**: type

EnhancedStatusCode is http status for response.


### EnhancedStatusType

**类型**: struct

EnhancedStatusType is the structure for certain response status.
Currently, it only supports `Response` and `Examples`.
`Response` is the response structure
`Examples` is the examples for the response, map[string]any, []any are supported.


### IEnhanceResponseStatus

**类型**: interface

IEnhanceResponseStatus is used to enhance the documentation of the response.
Normal response structure could implement this interface to provide more information.


### Response

**类型**: struct

Response is specified by OpenAPI/Swagger 3.0 standard.


**方法**:

- `func (r Response) MarshalJSON() ([]byte, error)`

### ResponseRef

**类型**: struct

**方法**:

- `func (r ResponseRef) MarshalJSON() ([]byte, error)`

### Responses

**类型**: type

Responses is specified by OpenAPI/Swagger 3.0 standard.


### SecurityScheme

**类型**: struct

### SecuritySchemes

**类型**: type

### SecuritySchemeRef

**类型**: struct

**方法**:

- `func (r SecuritySchemeRef) MarshalJSON() ([]byte, error)`

### SecurityRequirements

**类型**: type

### SecurityRequirement

**类型**: type

### OAuthFlows

**类型**: struct

### OAuthFlow

**类型**: struct

### Server

**类型**: struct

Server is specified by OpenAPI/Swagger standard version 3.0.


### ServerVariable

**类型**: struct

ServerVariable is specified by OpenAPI/Swagger standard version 3.0.


### Servers

**类型**: type

Servers is specified by OpenAPI/Swagger standard version 3.0.


### Schema

**类型**: struct

Schema is specified by OpenAPI/Swagger 3.0 standard.


**方法**:

- `func (s *Schema) Clone() *Schema`
  Clone only clones necessary attributes.
- `func (s Schema) MarshalJSON() ([]byte, error)`

### Discriminator

**类型**: struct

Discriminator is specified by OpenAPI/Swagger standard version 3.0.


### SchemaRefs

**类型**: type

### SchemaRef

**类型**: struct

**方法**:

- `func (r SchemaRef) MarshalJSON() ([]byte, error)`

### Schemas

**类型**: struct

**方法**:

- `func (s *Schemas) Clone() *Schemas`
- `func (s *Schemas) Get(name string) *SchemaRef`
- `func (s *Schemas) Set(name string, ref SchemaRef)`
- `func (s *Schemas) Removes(names []any)`
- `func (s *Schemas) Map() map[string]SchemaRef`
- `func (s *Schemas) Iterator(f func)`
- `func (s Schemas) MarshalJSON() ([]byte, error)`

### Tags

**类型**: type

Tags is specified by OpenAPI/Swagger 3.0 standard.


### Tag

**类型**: struct

Tag is specified by OpenAPI/Swagger 3.0 standard.


### XExtensions

**类型**: type

XExtensions stores the `x-` custom extensions.


## 函数

### New

```go
func New() *OpenApiV3
```

New creates and returns an OpenApiV3 implements object.

## 内部依赖

- `container/garray`
- `container/gmap`
- `container/gset`
- `encoding/gjson`
- `errors/gcode`
- `errors/gerror`
- `internal/empty`
- `internal/intlog`
- `internal/json`
- `internal/utils`
- `os/gfile`
- `os/gres`
- `os/gstructs`
- `text/gstr`
- `util/gconv`
- `util/gmeta`
- `util/gtag`
- `util/gvalid`

