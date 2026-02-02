# gsha256

> Package: `github.com/gogf/gf/v2/crypto/gsha256`

```go
import "github.com/gogf/gf/v2/crypto/gsha256"
```

## 概述

Package gsha256 provides useful API for SHA256 encryption algorithms.

## 源文件

- `gsha256.go`

## 函数

### Encrypt

```go
func Encrypt(v any) string
```

Encrypt encrypts any type of variable using SHA256 algorithms.
It uses package gconv to convert `v` to its bytes type.

### EncryptFile

```go
func EncryptFile(path string) (encrypt string, err error)
```

EncryptFile encrypts file content of `path` using SHA256 algorithms.

### MustEncryptFile

```go
func MustEncryptFile(path string) string
```

MustEncryptFile encrypts file content of `path` using the SHA256 algorithm.
It panics if any error occurs.

## 内部依赖

- `errors/gerror`
- `util/gconv`

