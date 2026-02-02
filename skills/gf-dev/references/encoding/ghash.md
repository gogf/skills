# ghash

> Package: `github.com/gogf/gf/v2/encoding/ghash`

```go
import "github.com/gogf/gf/v2/encoding/ghash"
```

## 概述

Package ghash provides some classic hash functions(uint32/uint64) in go.

## 源文件

- `ghash.go`
- `ghash_ap.go`
- `ghash_bkdr.go`
- `ghash_djb.go`
- `ghash_elf.go`
- `ghash_jshash.go`
- `ghash_pjw.go`
- `ghash_rs.go`
- `ghash_sdbm.go`

## 函数

### AP

```go
func AP(str []byte) uint32
```

AP implements the classic AP hash algorithm for 32 bits.

### AP64

```go
func AP64(str []byte) uint64
```

AP64 implements the classic AP hash algorithm for 64 bits.

### BKDR

```go
func BKDR(str []byte) uint32
```

BKDR implements the classic BKDR hash algorithm for 32 bits.

### BKDR64

```go
func BKDR64(str []byte) uint64
```

BKDR64 implements the classic BKDR hash algorithm for 64 bits.

### DJB

```go
func DJB(str []byte) uint32
```

DJB implements the classic DJB hash algorithm for 32 bits.

### DJB64

```go
func DJB64(str []byte) uint64
```

DJB64 implements the classic DJB hash algorithm for 64 bits.

### ELF

```go
func ELF(str []byte) uint32
```

ELF implements the classic ELF hash algorithm for 32 bits.

### ELF64

```go
func ELF64(str []byte) uint64
```

ELF64 implements the classic ELF hash algorithm for 64 bits.

### JS

```go
func JS(str []byte) uint32
```

JS implements the classic JS hash algorithm for 32 bits.

### JS64

```go
func JS64(str []byte) uint64
```

JS64 implements the classic JS hash algorithm for 64 bits.

### PJW

```go
func PJW(str []byte) uint32
```

PJW implements the classic PJW hash algorithm for 32 bits.

### PJW64

```go
func PJW64(str []byte) uint64
```

PJW64 implements the classic PJW hash algorithm for 64 bits.

### RS

```go
func RS(str []byte) uint32
```

RS implements the classic RS hash algorithm for 32 bits.

### RS64

```go
func RS64(str []byte) uint64
```

RS64 implements the classic RS hash algorithm for 64 bits.

### SDBM

```go
func SDBM(str []byte) uint32
```

SDBM implements the classic SDBM hash algorithm for 32 bits.

### SDBM64

```go
func SDBM64(str []byte) uint64
```

SDBM64 implements the classic SDBM hash algorithm for 64 bits.

