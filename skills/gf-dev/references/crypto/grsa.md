# grsa

> Package: `github.com/gogf/gf/v2/crypto/grsa`

```go
import "github.com/gogf/gf/v2/crypto/grsa"
```

## 概述

Package grsa provides useful API for RSA encryption/decryption algorithms.

## 源文件

- `grsa.go`

## 函数

### Encrypt

```go
func Encrypt(plainText []byte, publicKey []byte) ([]byte, error)
```

Encrypt encrypts data with public key using PKCS#1 v1.5 padding (auto-detect format).
The publicKey can be either PKCS#1 or PKCS#8 (PKIX) format.

Note: RSA encryption has a size limit based on key size.
For PKCS#1 v1.5 padding, max plaintext size = key_size_in_bytes - 11.
For example, a 2048-bit key can encrypt at most 245 bytes.

Security Warning: PKCS#1 v1.5 padding is vulnerable to padding oracle attacks.
For new applications, consider using EncryptOAEP instead.

### Decrypt

```go
func Decrypt(cipherText []byte, privateKey []byte) ([]byte, error)
```

Decrypt decrypts data with private key using PKCS#1 v1.5 padding (auto-detect format).
The privateKey can be either PKCS#1 or PKCS#8 format.

Security Warning: PKCS#1 v1.5 padding is vulnerable to padding oracle attacks.
For new applications, consider using DecryptOAEP instead.

### EncryptBase64

```go
func EncryptBase64(plainText []byte, publicKeyBase64 string) (string, error)
```

EncryptBase64 encrypts data with base64-encoded public key (auto-detect format)
and returns base64-encoded result.

### DecryptBase64

```go
func DecryptBase64(cipherTextBase64 string, privateKeyBase64 string) ([]byte, error)
```

DecryptBase64 decrypts base64-encoded data with base64-encoded private key (auto-detect format).

### EncryptPKIX

```go
func EncryptPKIX(plainText []byte, publicKey []byte) ([]byte, error)
```

EncryptPKIX encrypts data with public key in PKIX (X.509) format.
PKIX is the standard format for public keys, often referred to as "PKCS#8 public key".

Note: RSA encryption has a size limit based on key size.
For PKCS#1 v1.5 padding, max plaintext size = key_size_in_bytes - 11.

### EncryptPKCS8

```go
func EncryptPKCS8(plainText []byte, publicKey []byte) ([]byte, error)
```

EncryptPKCS8 is an alias for EncryptPKIX for backward compatibility.

Deprecated: Use EncryptPKIX instead. Public keys use PKIX format, not PKCS#8.

### EncryptPKCS1

```go
func EncryptPKCS1(plainText []byte, publicKey []byte) ([]byte, error)
```

EncryptPKCS1 encrypts data with public key in PKCS#1 format.

Note: RSA encryption has a size limit based on key size.
For PKCS#1 v1.5 padding, max plaintext size = key_size_in_bytes - 11.

### DecryptPKCS8

```go
func DecryptPKCS8(cipherText []byte, privateKey []byte) ([]byte, error)
```

DecryptPKCS8 decrypts data with private key by PKCS#8 format.

### DecryptPKCS1

```go
func DecryptPKCS1(cipherText []byte, privateKey []byte) ([]byte, error)
```

DecryptPKCS1 decrypts data with private key by PKCS#1 format.

### EncryptPKIXBase64

```go
func EncryptPKIXBase64(plainText []byte, publicKeyBase64 string) (string, error)
```

EncryptPKIXBase64 encrypts data with PKIX public key and returns base64-encoded result.

### EncryptPKCS8Base64

```go
func EncryptPKCS8Base64(plainText []byte, publicKeyBase64 string) (string, error)
```

EncryptPKCS8Base64 is an alias for EncryptPKIXBase64 for backward compatibility.

Deprecated: Use EncryptPKIXBase64 instead.

### EncryptPKCS1Base64

```go
func EncryptPKCS1Base64(plainText []byte, publicKeyBase64 string) (string, error)
```

EncryptPKCS1Base64 encrypts data with PKCS#1 public key and returns base64-encoded result.

### DecryptPKCS8Base64

```go
func DecryptPKCS8Base64(cipherTextBase64 string, privateKeyBase64 string) ([]byte, error)
```

DecryptPKCS8Base64 decrypts data with private key by PKCS#8 format and decode base64 input.

### DecryptPKCS1Base64

```go
func DecryptPKCS1Base64(cipherTextBase64 string, privateKeyBase64 string) ([]byte, error)
```

DecryptPKCS1Base64 decrypts base64-encoded data with PKCS#1 private key.

### GetPrivateKeyType

```go
func GetPrivateKeyType(privateKey []byte) (string, error)
```

GetPrivateKeyType detects the type of private key (PKCS#1 or PKCS#8).
It attempts to parse the key in both formats to determine the actual type.

### GetPrivateKeyTypeBase64

```go
func GetPrivateKeyTypeBase64(privateKeyBase64 string) (string, error)
```

GetPrivateKeyTypeBase64 detects the type of base64 encoded private key (PKCS#1 or PKCS#8).

### GenerateKeyPair

```go
func GenerateKeyPair(bits int) (privateKey []byte, publicKey []byte, err error)
```

GenerateKeyPair generates a new RSA key pair with the given bits.

### GenerateKeyPairPKCS8

```go
func GenerateKeyPairPKCS8(bits int) (privateKey []byte, publicKey []byte, err error)
```

GenerateKeyPairPKCS8 generates a new RSA key pair with the given bits in PKCS#8 format.

### GenerateDefaultKeyPair

```go
func GenerateDefaultKeyPair() (privateKey []byte, publicKey []byte, err error)
```

GenerateDefaultKeyPair generates a new RSA key pair with default bits (2048).

### ExtractPKCS1PublicKey

```go
func ExtractPKCS1PublicKey(privateKey []byte) ([]byte, error)
```

ExtractPKCS1PublicKey extracts PKCS#1 public key from private key.

### EncryptOAEP

```go
func EncryptOAEP(plainText []byte, publicKey []byte) ([]byte, error)
```

EncryptOAEP encrypts data with public key using OAEP padding (auto-detect format).
The publicKey can be either PKCS#1 or PKCS#8 (PKIX) format.
Uses SHA-256 as the hash function by default.

OAEP (Optimal Asymmetric Encryption Padding) is more secure than PKCS#1 v1.5
and is recommended for new applications.

Note: For OAEP with SHA-256, max plaintext size = key_size_in_bytes - 2*32 - 2.
For a 2048-bit key, this is 190 bytes.

### EncryptOAEPWithHash

```go
func EncryptOAEPWithHash(plainText []byte, publicKey []byte, label []byte, hash hash.Hash) ([]byte, error)
```

EncryptOAEPWithHash encrypts data with public key using OAEP padding with custom hash.
The publicKey can be either PKCS#1 or PKCS#8 (PKIX) format.
The label parameter can be nil for most use cases.
The hash parameter specifies the hash function to use (e.g., sha256.New()).

### DecryptOAEP

```go
func DecryptOAEP(cipherText []byte, privateKey []byte) ([]byte, error)
```

DecryptOAEP decrypts data with private key using OAEP padding (auto-detect format).
The privateKey can be either PKCS#1 or PKCS#8 format.
Uses SHA-256 as the hash function by default.

### DecryptOAEPWithHash

```go
func DecryptOAEPWithHash(cipherText []byte, privateKey []byte, label []byte, hash hash.Hash) ([]byte, error)
```

DecryptOAEPWithHash decrypts data with private key using OAEP padding with custom hash.
The privateKey can be either PKCS#1 or PKCS#8 format.
The label parameter must match the label used during encryption (nil if not used).
The hash parameter must match the hash function used during encryption.

### EncryptOAEPBase64

```go
func EncryptOAEPBase64(plainText []byte, publicKeyBase64 string) (string, error)
```

EncryptOAEPBase64 encrypts data with public key using OAEP padding
and returns base64-encoded result.

### DecryptOAEPBase64

```go
func DecryptOAEPBase64(cipherTextBase64 string, privateKeyBase64 string) ([]byte, error)
```

DecryptOAEPBase64 decrypts base64-encoded data with private key using OAEP padding.

## 内部依赖

- `errors/gcode`
- `errors/gerror`

