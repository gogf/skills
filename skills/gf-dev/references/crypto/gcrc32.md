# gcrc32

> Package: `github.com/gogf/gf/v2/crypto/gcrc32`

```go
import "github.com/gogf/gf/v2/crypto/gcrc32"
```

## 概述

Package gcrc32 provides useful API for CRC32 encryption algorithms.

## 源文件

- `gcrc32.go`

## 函数

### Encrypt

```go
func Encrypt(v any) uint32
```

Encrypt encrypts any type of variable using CRC32 algorithms.
It uses gconv package to convert `v` to its bytes type.

## 内部依赖

- `util/gconv`

