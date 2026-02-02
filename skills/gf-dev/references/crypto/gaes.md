# gaes

> Package: `github.com/gogf/gf/v2/crypto/gaes`

```go
import "github.com/gogf/gf/v2/crypto/gaes"
```

## 概述

Package gaes provides useful API for AES encryption/decryption algorithms.

## 源文件

- `gaes.go`

## 函数

### Encrypt

```go
func Encrypt(plainText []byte, key []byte, iv ...[]byte) ([]byte, error)
```

Encrypt is alias of EncryptCBC.

### Decrypt

```go
func Decrypt(cipherText []byte, key []byte, iv ...[]byte) ([]byte, error)
```

Decrypt is alias of DecryptCBC.

### EncryptCBC

```go
func EncryptCBC(plainText []byte, key []byte, iv ...[]byte) ([]byte, error)
```

EncryptCBC encrypts `plainText` using CBC mode.
Note that the key must be 16/24/32 bit length.
The parameter `iv` initialization vector is unnecessary.

### DecryptCBC

```go
func DecryptCBC(cipherText []byte, key []byte, iv ...[]byte) ([]byte, error)
```

DecryptCBC decrypts `cipherText` using CBC mode.
Note that the key must be 16/24/32 bit length.
The parameter `iv` initialization vector is unnecessary.

### PKCS5Padding

```go
func PKCS5Padding(src []byte, blockSize ...int) []byte
```

PKCS5Padding applies PKCS#5 padding to the source byte slice to match the given block size.

If the block size is not provided, it defaults to 8.

### PKCS5UnPadding

```go
func PKCS5UnPadding(src []byte, blockSize ...int) ([]byte, error)
```

PKCS5UnPadding removes PKCS#5 padding from the source byte slice based on the given block size.

If the block size is not provided, it defaults to 8.

### PKCS7Padding

```go
func PKCS7Padding(src []byte, blockSize int) []byte
```

PKCS7Padding applies PKCS#7 padding to the source byte slice to match the given block size.

### PKCS7UnPadding

```go
func PKCS7UnPadding(src []byte, blockSize int) ([]byte, error)
```

PKCS7UnPadding removes PKCS#7 padding from the source byte slice based on the given block size.

### EncryptCFB

```go
func EncryptCFB(plainText []byte, key []byte, padding *int, iv ...[]byte) ([]byte, error)
```

EncryptCFB encrypts `plainText` using CFB mode.
Note that the key must be 16/24/32 bit length.
The parameter `iv` initialization vector is unnecessary.

### DecryptCFB

```go
func DecryptCFB(cipherText []byte, key []byte, unPadding int, iv ...[]byte) ([]byte, error)
```

DecryptCFB decrypts `plainText` using CFB mode.
Note that the key must be 16/24/32 bit length.
The parameter `iv` initialization vector is unnecessary.

### ZeroPadding

```go
func ZeroPadding(cipherText []byte, blockSize int) ([]byte, int)
```

### ZeroUnPadding

```go
func ZeroUnPadding(plaintext []byte, unPadding int) []byte
```

## 内部依赖

- `errors/gcode`
- `errors/gerror`

