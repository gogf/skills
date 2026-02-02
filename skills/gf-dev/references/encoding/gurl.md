# gurl

> Package: `github.com/gogf/gf/v2/encoding/gurl`

```go
import "github.com/gogf/gf/v2/encoding/gurl"
```

## 概述

Package gurl provides useful API for URL handling.

## 源文件

- `url.go`

## 函数

### Encode

```go
func Encode(str string) string
```

Encode escapes the string so it can be safely placed
inside an URL query.

### Decode

```go
func Decode(str string) (string, error)
```

Decode does the inverse transformation of Encode,
converting each 3-byte encoded substring of the form "%AB" into the
hex-decoded byte 0xAB.
It returns an error if any % is not followed by two hexadecimal
digits.

### RawEncode

```go
func RawEncode(str string) string
```

RawEncode does encode the given string according
URL-encode according to RFC 3986.
See http://php.net/manual/en/function.rawurlencode.php.

### RawDecode

```go
func RawDecode(str string) (string, error)
```

RawDecode does decode the given string
Decode URL-encoded strings.
See http://php.net/manual/en/function.rawurldecode.php.

### BuildQuery

```go
func BuildQuery(queryData url.Values) string
```

BuildQuery Generate URL-encoded query string.
See http://php.net/manual/en/function.http-build-query.php.

### ParseURL

```go
func ParseURL(str string, component int) (map[string]string, error)
```

ParseURL Parse an URL and return its components.
-1: all; 1: scheme; 2: host; 4: port; 8: user; 16: pass; 32: path; 64: query; 128: fragment.
See http://php.net/manual/en/function.parse-url.php.

## 内部依赖

- `errors/gerror`

