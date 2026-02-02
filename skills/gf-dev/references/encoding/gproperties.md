# gproperties

> Package: `github.com/gogf/gf/v2/encoding/gproperties`

```go
import "github.com/gogf/gf/v2/encoding/gproperties"
```

## 概述

Package gproperties provides accessing and converting for .properties content.

## 源文件

- `gproperties.go`

## 函数

### Decode

```go
func Decode(data []byte) (res map[string]any, err error)
```

Decode converts properties format to map.

### Encode

```go
func Encode(data map[string]any) (res []byte, err error)
```

Encode converts map to properties format.

### ToJson

```go
func ToJson(data []byte) (res []byte, err error)
```

ToJson convert .properties format to JSON.

## 内部依赖

- `errors/gerror`
- `internal/json`
- `util/gconv`

## 外部依赖

- `github.com/magiconair/properties`

