# gdes

> Package: `github.com/gogf/gf/v2/crypto/gdes`

```go
import "github.com/gogf/gf/v2/crypto/gdes"
```

## 概述

Package gdes provides useful API for DES encryption/decryption algorithms.

## 源文件

- `gdes.go`

## 函数

### EncryptECB

```go
func EncryptECB(plainText []byte, key []byte, padding int) ([]byte, error)
```

EncryptECB encrypts `plainText` using ECB mode.

### DecryptECB

```go
func DecryptECB(cipherText []byte, key []byte, padding int) ([]byte, error)
```

DecryptECB decrypts `cipherText` using ECB mode.

### EncryptECBTriple

```go
func EncryptECBTriple(plainText []byte, key []byte, padding int) ([]byte, error)
```

EncryptECBTriple encrypts `plainText` using TripleDES and ECB mode.
The length of the `key` should be either 16 or 24 bytes.

### DecryptECBTriple

```go
func DecryptECBTriple(cipherText []byte, key []byte, padding int) ([]byte, error)
```

DecryptECBTriple decrypts `cipherText` using TripleDES and ECB mode.
The length of the `key` should be either 16 or 24 bytes.

### EncryptCBC

```go
func EncryptCBC(plainText []byte, key []byte, iv []byte, padding int) ([]byte, error)
```

EncryptCBC encrypts `plainText` using CBC mode.

### DecryptCBC

```go
func DecryptCBC(cipherText []byte, key []byte, iv []byte, padding int) ([]byte, error)
```

DecryptCBC decrypts `cipherText` using CBC mode.

### EncryptCBCTriple

```go
func EncryptCBCTriple(plainText []byte, key []byte, iv []byte, padding int) ([]byte, error)
```

EncryptCBCTriple encrypts `plainText` using TripleDES and CBC mode.

### DecryptCBCTriple

```go
func DecryptCBCTriple(cipherText []byte, key []byte, iv []byte, padding int) ([]byte, error)
```

DecryptCBCTriple decrypts `cipherText` using TripleDES and CBC mode.

### PaddingPKCS5

```go
func PaddingPKCS5(text []byte, blockSize int) []byte
```

### UnPaddingPKCS5

```go
func UnPaddingPKCS5(text []byte) []byte
```

### Padding

```go
func Padding(text []byte, padding int) ([]byte, error)
```

### UnPadding

```go
func UnPadding(text []byte, padding int) ([]byte, error)
```

## 内部依赖

- `errors/gcode`
- `errors/gerror`

