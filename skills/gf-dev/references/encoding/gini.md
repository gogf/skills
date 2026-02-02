# gini

> Package: `github.com/gogf/gf/v2/encoding/gini`

```go
import "github.com/gogf/gf/v2/encoding/gini"
```

## 概述

Package gini provides accessing and converting for INI content.

## 源文件

- `gini.go`

## 函数

### Decode

```go
func Decode(data []byte) (res map[string]any, err error)
```

Decode converts INI format to map.

### Encode

```go
func Encode(data map[string]any) (res []byte, err error)
```

Encode converts map to INI format.

### ToJson

```go
func ToJson(data []byte) (res []byte, err error)
```

ToJson convert INI format to JSON.

## 内部依赖

- `errors/gcode`
- `errors/gerror`
- `internal/json`

