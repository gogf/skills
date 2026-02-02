# gmd5

> Package: `github.com/gogf/gf/v2/crypto/gmd5`

```go
import "github.com/gogf/gf/v2/crypto/gmd5"
```

## 概述

Package gmd5 provides useful API for MD5 encryption algorithms.

## 源文件

- `gmd5.go`

## 函数

### Encrypt

```go
func Encrypt(data any) (encrypt string, err error)
```

Encrypt encrypts any type of variable using MD5 algorithms.
It uses gconv package to convert `v` to its bytes type.

### MustEncrypt

```go
func MustEncrypt(data any) string
```

MustEncrypt encrypts any type of variable using MD5 algorithms.
It uses gconv package to convert `v` to its bytes type.
It panics if any error occurs.

### EncryptBytes

```go
func EncryptBytes(data []byte) (encrypt string, err error)
```

EncryptBytes encrypts `data` using MD5 algorithms.

### MustEncryptBytes

```go
func MustEncryptBytes(data []byte) string
```

MustEncryptBytes encrypts `data` using MD5 algorithms.
It panics if any error occurs.

### EncryptString

```go
func EncryptString(data string) (encrypt string, err error)
```

EncryptString encrypts string `data` using MD5 algorithms.

### MustEncryptString

```go
func MustEncryptString(data string) string
```

MustEncryptString encrypts string `data` using MD5 algorithms.
It panics if any error occurs.

### EncryptFile

```go
func EncryptFile(path string) (encrypt string, err error)
```

EncryptFile encrypts file content of `path` using MD5 algorithms.

### MustEncryptFile

```go
func MustEncryptFile(path string) string
```

MustEncryptFile encrypts file content of `path` using MD5 algorithms.
It panics if any error occurs.

## 内部依赖

- `errors/gerror`
- `util/gconv`

