# gyaml

> Package: `github.com/gogf/gf/v2/encoding/gyaml`

```go
import "github.com/gogf/gf/v2/encoding/gyaml"
```

## 概述

Package gyaml provides accessing and converting for YAML content.

## 源文件

- `gyaml.go`

## 函数

### Encode

```go
func Encode(value any) (out []byte, err error)
```

Encode encodes `value` to an YAML format content as bytes.

### EncodeIndent

```go
func EncodeIndent(value any, indent string) (out []byte, err error)
```

EncodeIndent encodes `value` to an YAML format content with indent as bytes.

### Decode

```go
func Decode(content []byte) (map[string]any, error)
```

Decode parses `content` into and returns as map.

### DecodeTo

```go
func DecodeTo(value []byte, result any) err error
```

DecodeTo parses `content` into `result`.

### ToJson

```go
func ToJson(content []byte) (out []byte, err error)
```

ToJson converts `content` to JSON format content.

## 内部依赖

- `errors/gerror`
- `internal/json`
- `util/gconv`

## 外部依赖

- `gopkg.in/yaml.v3`

