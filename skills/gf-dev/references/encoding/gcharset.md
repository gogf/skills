# gcharset

> Package: `github.com/gogf/gf/v2/encoding/gcharset`

```go
import "github.com/gogf/gf/v2/encoding/gcharset"
```

## 概述

Package gcharset implements character-set conversion functionality.

## 源文件

- `gcharset.go`

## 函数

### Supported

```go
func Supported(charset string) bool
```

Supported returns whether charset `charset` is supported.

### Convert

```go
func Convert(dstCharset string, srcCharset string, src string) (dst string, err error)
```

Convert converts `src` charset encoding from `srcCharset` to `dstCharset`,
and returns the converted string.
It returns `src` as `dst` if it fails converting.

### ToUTF8

```go
func ToUTF8(srcCharset string, src string) (dst string, err error)
```

ToUTF8 converts `src` charset encoding from `srcCharset` to UTF-8 ,
and returns the converted string.

### UTF8To

```go
func UTF8To(dstCharset string, src string) (dst string, err error)
```

UTF8To converts `src` charset encoding from UTF-8 to `dstCharset`,
and returns the converted string.

## 内部依赖

- `errors/gcode`
- `errors/gerror`
- `internal/intlog`

## 外部依赖

- `golang.org/x/text/encoding`
- `golang.org/x/text/encoding/ianaindex`
- `golang.org/x/text/transform`

