# gregex

> Package: `github.com/gogf/gf/v2/text/gregex`

```go
import "github.com/gogf/gf/v2/text/gregex"
```

## 概述

Package gregex provides high performance API for regular expression functionality.

## 源文件

- `gregex.go`
- `gregex_cache.go`

## 函数

### Quote

```go
func Quote(s string) string
```

Quote quotes `s` by replacing special chars in `s`
to match the rules of regular expression pattern.
And returns the copy.

Eg: Quote(`[foo]`) returns `\[foo\]`.

### Validate

```go
func Validate(pattern string) error
```

Validate checks whether given regular expression pattern `pattern` valid.

### IsMatch

```go
func IsMatch(pattern string, src []byte) bool
```

IsMatch checks whether given bytes `src` matches `pattern`.

### IsMatchString

```go
func IsMatchString(pattern string, src string) bool
```

IsMatchString checks whether given string `src` matches `pattern`.

### Match

```go
func Match(pattern string, src []byte) ([][]byte, error)
```

Match return bytes slice that matched `pattern`.

### MatchString

```go
func MatchString(pattern string, src string) ([]string, error)
```

MatchString return strings that matched `pattern`.

### MatchAll

```go
func MatchAll(pattern string, src []byte) ([][][]byte, error)
```

MatchAll return all bytes slices that matched `pattern`.

### MatchAllString

```go
func MatchAllString(pattern string, src string) ([][]string, error)
```

MatchAllString return all strings that matched `pattern`.

### Replace

```go
func Replace(pattern string, replace []byte, src []byte) ([]byte, error)
```

Replace replaces all matched `pattern` in bytes `src` with bytes `replace`.

### ReplaceString

```go
func ReplaceString(pattern string, replace string, src string) (string, error)
```

ReplaceString replace all matched `pattern` in string `src` with string `replace`.

### ReplaceFunc

```go
func ReplaceFunc(pattern string, src []byte, replaceFunc func) ([]byte, error)
```

ReplaceFunc replace all matched `pattern` in bytes `src`
with custom replacement function `replaceFunc`.

### ReplaceFuncMatch

```go
func ReplaceFuncMatch(pattern string, src []byte, replaceFunc func) ([]byte, error)
```

ReplaceFuncMatch replace all matched `pattern` in bytes `src`
with custom replacement function `replaceFunc`.
The parameter `match` type for `replaceFunc` is [][]byte,
which is the result contains all sub-patterns of `pattern` using Match function.

### ReplaceStringFunc

```go
func ReplaceStringFunc(pattern string, src string, replaceFunc func) (string, error)
```

ReplaceStringFunc replace all matched `pattern` in string `src`
with custom replacement function `replaceFunc`.

### ReplaceStringFuncMatch

```go
func ReplaceStringFuncMatch(pattern string, src string, replaceFunc func) (string, error)
```

ReplaceStringFuncMatch replace all matched `pattern` in string `src`
with custom replacement function `replaceFunc`.
The parameter `match` type for `replaceFunc` is []string,
which is the result contains all sub-patterns of `pattern` using MatchString function.

### Split

```go
func Split(pattern string, src string) []string
```

Split slices `src` into substrings separated by the expression and returns a slice of
the substrings between those expression matches.

## 内部依赖

- `errors/gerror`

