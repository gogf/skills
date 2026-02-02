# gbinary

> Package: `github.com/gogf/gf/v2/encoding/gbinary`

```go
import "github.com/gogf/gf/v2/encoding/gbinary"
```

## 概述

Package gbinary provides useful API for handling binary/bytes data.

## 源文件

- `gbinary.go`
- `gbinary_be.go`
- `gbinary_bit.go`
- `gbinary_func.go`
- `gbinary_le.go`

## 类型定义

### Bit

**类型**: type

Bit Binary bit (0 | 1)


## 函数

### Encode

```go
func Encode(values ...any) []byte
```

### EncodeByLength

```go
func EncodeByLength(length int, values ...any) []byte
```

### Decode

```go
func Decode(b []byte, values ...any) error
```

### EncodeString

```go
func EncodeString(s string) []byte
```

### DecodeToString

```go
func DecodeToString(b []byte) string
```

### EncodeBool

```go
func EncodeBool(b bool) []byte
```

### EncodeInt

```go
func EncodeInt(i int) []byte
```

### EncodeUint

```go
func EncodeUint(i uint) []byte
```

### EncodeInt8

```go
func EncodeInt8(i int8) []byte
```

### EncodeUint8

```go
func EncodeUint8(i uint8) []byte
```

### EncodeInt16

```go
func EncodeInt16(i int16) []byte
```

### EncodeUint16

```go
func EncodeUint16(i uint16) []byte
```

### EncodeInt32

```go
func EncodeInt32(i int32) []byte
```

### EncodeUint32

```go
func EncodeUint32(i uint32) []byte
```

### EncodeInt64

```go
func EncodeInt64(i int64) []byte
```

### EncodeUint64

```go
func EncodeUint64(i uint64) []byte
```

### EncodeFloat32

```go
func EncodeFloat32(f float32) []byte
```

### EncodeFloat64

```go
func EncodeFloat64(f float64) []byte
```

### DecodeToInt

```go
func DecodeToInt(b []byte) int
```

### DecodeToUint

```go
func DecodeToUint(b []byte) uint
```

### DecodeToBool

```go
func DecodeToBool(b []byte) bool
```

### DecodeToInt8

```go
func DecodeToInt8(b []byte) int8
```

### DecodeToUint8

```go
func DecodeToUint8(b []byte) uint8
```

### DecodeToInt16

```go
func DecodeToInt16(b []byte) int16
```

### DecodeToUint16

```go
func DecodeToUint16(b []byte) uint16
```

### DecodeToInt32

```go
func DecodeToInt32(b []byte) int32
```

### DecodeToUint32

```go
func DecodeToUint32(b []byte) uint32
```

### DecodeToInt64

```go
func DecodeToInt64(b []byte) int64
```

### DecodeToUint64

```go
func DecodeToUint64(b []byte) uint64
```

### DecodeToFloat32

```go
func DecodeToFloat32(b []byte) float32
```

### DecodeToFloat64

```go
func DecodeToFloat64(b []byte) float64
```

### BeEncode

```go
func BeEncode(values ...any) []byte
```

BeEncode encodes one or multiple `values` into bytes using BigEndian.
It uses type asserting checking the type of each value of `values` and internally
calls corresponding converting function do the bytes converting.

It supports common variable type asserting, and finally it uses fmt.Sprintf converting
value to string and then to bytes.

### BeEncodeByLength

```go
func BeEncodeByLength(length int, values ...any) []byte
```

### BeDecode

```go
func BeDecode(b []byte, values ...any) error
```

### BeEncodeString

```go
func BeEncodeString(s string) []byte
```

### BeDecodeToString

```go
func BeDecodeToString(b []byte) string
```

### BeEncodeBool

```go
func BeEncodeBool(b bool) []byte
```

### BeEncodeInt

```go
func BeEncodeInt(i int) []byte
```

### BeEncodeUint

```go
func BeEncodeUint(i uint) []byte
```

### BeEncodeInt8

```go
func BeEncodeInt8(i int8) []byte
```

### BeEncodeUint8

```go
func BeEncodeUint8(i uint8) []byte
```

### BeEncodeInt16

```go
func BeEncodeInt16(i int16) []byte
```

### BeEncodeUint16

```go
func BeEncodeUint16(i uint16) []byte
```

### BeEncodeInt32

```go
func BeEncodeInt32(i int32) []byte
```

### BeEncodeUint32

```go
func BeEncodeUint32(i uint32) []byte
```

### BeEncodeInt64

```go
func BeEncodeInt64(i int64) []byte
```

### BeEncodeUint64

```go
func BeEncodeUint64(i uint64) []byte
```

### BeEncodeFloat32

```go
func BeEncodeFloat32(f float32) []byte
```

### BeEncodeFloat64

```go
func BeEncodeFloat64(f float64) []byte
```

### BeDecodeToInt

```go
func BeDecodeToInt(b []byte) int
```

### BeDecodeToUint

```go
func BeDecodeToUint(b []byte) uint
```

### BeDecodeToBool

```go
func BeDecodeToBool(b []byte) bool
```

### BeDecodeToInt8

```go
func BeDecodeToInt8(b []byte) int8
```

### BeDecodeToUint8

```go
func BeDecodeToUint8(b []byte) uint8
```

### BeDecodeToInt16

```go
func BeDecodeToInt16(b []byte) int16
```

### BeDecodeToUint16

```go
func BeDecodeToUint16(b []byte) uint16
```

### BeDecodeToInt32

```go
func BeDecodeToInt32(b []byte) int32
```

### BeDecodeToUint32

```go
func BeDecodeToUint32(b []byte) uint32
```

### BeDecodeToInt64

```go
func BeDecodeToInt64(b []byte) int64
```

### BeDecodeToUint64

```go
func BeDecodeToUint64(b []byte) uint64
```

### BeDecodeToFloat32

```go
func BeDecodeToFloat32(b []byte) float32
```

### BeDecodeToFloat64

```go
func BeDecodeToFloat64(b []byte) float64
```

### BeFillUpSize

```go
func BeFillUpSize(b []byte, l int) []byte
```

BeFillUpSize fills up the bytes `b` to given length `l` using big BigEndian.

Note that it creates a new bytes slice by copying the original one to avoid changing
the original parameter bytes.

### EncodeBits

```go
func EncodeBits(bits []Bit, i int, l int) []Bit
```

EncodeBits does encode bits return bits Default coding

### EncodeBitsWithUint

```go
func EncodeBitsWithUint(bits []Bit, ui uint, l int) []Bit
```

EncodeBitsWithUint . Merge ui bitwise into the bits array and occupy the length bits
(Note: binary 0 | 1 digits are stored in the uis array)

### EncodeBitsToBytes

```go
func EncodeBitsToBytes(bits []Bit) []byte
```

EncodeBitsToBytes . does encode bits to bytes
Convert bits to [] byte, encode from left to right, and add less than 1 byte from 0 to the end.

### DecodeBits

```go
func DecodeBits(bits []Bit) int
```

DecodeBits .does decode bits to int
Resolve to int

### DecodeBitsToUint

```go
func DecodeBitsToUint(bits []Bit) uint
```

DecodeBitsToUint .Resolve to uint

### DecodeBytesToBits

```go
func DecodeBytesToBits(bs []byte) []Bit
```

DecodeBytesToBits .Parsing [] byte into character array [] uint8

### LeEncode

```go
func LeEncode(values ...any) []byte
```

LeEncode encodes one or multiple `values` into bytes using LittleEndian.
It uses type asserting checking the type of each value of `values` and internally
calls corresponding converting function do the bytes converting.

It supports common variable type asserting, and finally it uses fmt.Sprintf converting
value to string and then to bytes.

### LeEncodeByLength

```go
func LeEncodeByLength(length int, values ...any) []byte
```

### LeDecode

```go
func LeDecode(b []byte, values ...any) error
```

### LeEncodeString

```go
func LeEncodeString(s string) []byte
```

### LeDecodeToString

```go
func LeDecodeToString(b []byte) string
```

### LeEncodeBool

```go
func LeEncodeBool(b bool) []byte
```

### LeEncodeInt

```go
func LeEncodeInt(i int) []byte
```

### LeEncodeUint

```go
func LeEncodeUint(i uint) []byte
```

### LeEncodeInt8

```go
func LeEncodeInt8(i int8) []byte
```

### LeEncodeUint8

```go
func LeEncodeUint8(i uint8) []byte
```

### LeEncodeInt16

```go
func LeEncodeInt16(i int16) []byte
```

### LeEncodeUint16

```go
func LeEncodeUint16(i uint16) []byte
```

### LeEncodeInt32

```go
func LeEncodeInt32(i int32) []byte
```

### LeEncodeUint32

```go
func LeEncodeUint32(i uint32) []byte
```

### LeEncodeInt64

```go
func LeEncodeInt64(i int64) []byte
```

### LeEncodeUint64

```go
func LeEncodeUint64(i uint64) []byte
```

### LeEncodeFloat32

```go
func LeEncodeFloat32(f float32) []byte
```

### LeEncodeFloat64

```go
func LeEncodeFloat64(f float64) []byte
```

### LeDecodeToInt

```go
func LeDecodeToInt(b []byte) int
```

### LeDecodeToUint

```go
func LeDecodeToUint(b []byte) uint
```

### LeDecodeToBool

```go
func LeDecodeToBool(b []byte) bool
```

### LeDecodeToInt8

```go
func LeDecodeToInt8(b []byte) int8
```

### LeDecodeToUint8

```go
func LeDecodeToUint8(b []byte) uint8
```

### LeDecodeToInt16

```go
func LeDecodeToInt16(b []byte) int16
```

### LeDecodeToUint16

```go
func LeDecodeToUint16(b []byte) uint16
```

### LeDecodeToInt32

```go
func LeDecodeToInt32(b []byte) int32
```

### LeDecodeToUint32

```go
func LeDecodeToUint32(b []byte) uint32
```

### LeDecodeToInt64

```go
func LeDecodeToInt64(b []byte) int64
```

### LeDecodeToUint64

```go
func LeDecodeToUint64(b []byte) uint64
```

### LeDecodeToFloat32

```go
func LeDecodeToFloat32(b []byte) float32
```

### LeDecodeToFloat64

```go
func LeDecodeToFloat64(b []byte) float64
```

### LeFillUpSize

```go
func LeFillUpSize(b []byte, l int) []byte
```

LeFillUpSize fills up the bytes `b` to given length `l` using LittleEndian.

Note that it creates a new bytes slice by copying the original one to avoid changing
the original parameter bytes.

## 内部依赖

- `errors/gerror`
- `internal/intlog`

