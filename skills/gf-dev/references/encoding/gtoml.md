# gtoml

> Package: `github.com/gogf/gf/v2/encoding/gtoml`

```go
import "github.com/gogf/gf/v2/encoding/gtoml"
```

## 概述

Package gtoml provides accessing and converting for TOML content.

## 源文件

- `gtoml.go`

## 函数

### Encode

```go
func Encode(v any) ([]byte, error)
```

### Decode

```go
func Decode(v []byte) (any, error)
```

### DecodeTo

```go
func DecodeTo(v []byte, result any) err error
```

### ToJson

```go
func ToJson(v []byte) ([]byte, error)
```

## 内部依赖

- `errors/gerror`
- `internal/json`

## 外部依赖

- `github.com/BurntSushi/toml`

