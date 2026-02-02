# gbase64

> Package: `github.com/gogf/gf/v2/encoding/gbase64`

```go
import "github.com/gogf/gf/v2/encoding/gbase64"
```

## 概述

Package gbase64 provides useful API for BASE64 encoding/decoding algorithm.

## 源文件

- `gbase64.go`

## 函数

### Encode

```go
func Encode(src []byte) []byte
```

Encode encodes bytes with BASE64 algorithm.

### EncodeString

```go
func EncodeString(src string) string
```

EncodeString encodes string with BASE64 algorithm.

### EncodeToString

```go
func EncodeToString(src []byte) string
```

EncodeToString encodes bytes to string with BASE64 algorithm.

### EncodeFile

```go
func EncodeFile(path string) ([]byte, error)
```

EncodeFile encodes file content of `path` using BASE64 algorithms.

### MustEncodeFile

```go
func MustEncodeFile(path string) []byte
```

MustEncodeFile encodes file content of `path` using BASE64 algorithms.
It panics if any error occurs.

### EncodeFileToString

```go
func EncodeFileToString(path string) (string, error)
```

EncodeFileToString encodes file content of `path` to string using BASE64 algorithms.

### MustEncodeFileToString

```go
func MustEncodeFileToString(path string) string
```

MustEncodeFileToString encodes file content of `path` to string using BASE64 algorithms.
It panics if any error occurs.

### Decode

```go
func Decode(data []byte) ([]byte, error)
```

Decode decodes bytes with BASE64 algorithm.

### MustDecode

```go
func MustDecode(data []byte) []byte
```

MustDecode decodes bytes with BASE64 algorithm.
It panics if any error occurs.

### DecodeString

```go
func DecodeString(data string) ([]byte, error)
```

DecodeString decodes string with BASE64 algorithm.

### MustDecodeString

```go
func MustDecodeString(data string) []byte
```

MustDecodeString decodes string with BASE64 algorithm.
It panics if any error occurs.

### DecodeToString

```go
func DecodeToString(data string) (string, error)
```

DecodeToString decodes string with BASE64 algorithm.

### MustDecodeToString

```go
func MustDecodeToString(data string) string
```

MustDecodeToString decodes string with BASE64 algorithm.
It panics if any error occurs.

## 内部依赖

- `errors/gerror`

