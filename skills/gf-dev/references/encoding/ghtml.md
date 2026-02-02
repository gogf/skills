# ghtml

> Package: `github.com/gogf/gf/v2/encoding/ghtml`

```go
import "github.com/gogf/gf/v2/encoding/ghtml"
```

## 概述

Package ghtml provides useful API for HTML content handling.

## 源文件

- `ghtml.go`

## 函数

### StripTags

```go
func StripTags(s string) string
```

StripTags strips HTML tags from content, and returns only text.
Referer: http://php.net/manual/zh/function.strip-tags.php

### Entities

```go
func Entities(s string) string
```

Entities encodes all HTML chars for content.
Referer: http://php.net/manual/zh/function.htmlentities.php

### EntitiesDecode

```go
func EntitiesDecode(s string) string
```

EntitiesDecode decodes all HTML chars for content.
Referer: http://php.net/manual/zh/function.html-entity-decode.php

### SpecialChars

```go
func SpecialChars(s string) string
```

SpecialChars encodes some special chars for content, these special chars are:
"&", "<", ">", `"`, "'".
Referer: http://php.net/manual/zh/function.htmlspecialchars.php

### SpecialCharsDecode

```go
func SpecialCharsDecode(s string) string
```

SpecialCharsDecode decodes some special chars for content, these special chars are:
"&", "<", ">", `"`, "'".
Referer: http://php.net/manual/zh/function.htmlspecialchars-decode.php

### SpecialCharsMapOrStruct

```go
func SpecialCharsMapOrStruct(mapOrStruct any) error
```

SpecialCharsMapOrStruct automatically encodes string values/attributes for map/struct.

Note that, if operation on struct, the given parameter `mapOrStruct` should be type of pointer to struct.

For example:
var m = map{}
var s = struct{}{}
OK: SpecialCharsMapOrStruct(m)
OK: SpecialCharsMapOrStruct(&s)
Error: SpecialCharsMapOrStruct(s)

## 内部依赖

- `errors/gcode`
- `errors/gerror`

## 外部依赖

- `github.com/grokify/html-strip-tags-go`

