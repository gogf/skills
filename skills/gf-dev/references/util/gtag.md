# gtag

> Package: `github.com/gogf/gf/v2/util/gtag`

```go
import "github.com/gogf/gf/v2/util/gtag"
```

## 概述

Package gtag providing tag content storing for struct.

## 源文件

- `gtag.go`
- `gtag_enums.go`
- `gtag_func.go`

## 函数

### SetGlobalEnums

```go
func SetGlobalEnums(enumsJSON string) error
```

SetGlobalEnums sets the global enums into package.
Note that this operation is not concurrent safety.

### GetGlobalEnums

```go
func GetGlobalEnums() (string, error)
```

GetGlobalEnums retrieves and returns the global enums.

### GetEnumsByType

```go
func GetEnumsByType(typeName string) string
```

GetEnumsByType retrieves and returns the stored enums json by type name.
The type name is like: github.com/gogf/gf/v2/encoding/gjson.ContentType

### Set

```go
func Set(name string, value string)
```

Set sets tag content for specified name.
Note that it panics if `name` already exists.

### SetOver

```go
func SetOver(name string, value string)
```

SetOver performs as Set, but it overwrites the old value if `name` already exists.

### Sets

```go
func Sets(m map[string]string)
```

Sets sets multiple tag content by map.

### SetsOver

```go
func SetsOver(m map[string]string)
```

SetsOver performs as Sets, but it overwrites the old value if `name` already exists.

### Get

```go
func Get(name string) string
```

Get retrieves and returns the stored tag content for specified name.

### Parse

```go
func Parse(content string) string
```

Parse parses and returns the content by replacing all tag name variable to
its content for given `content`.
Eg:
gtag.Set("demo", "content")
Parse(`This is {demo}`) -> `This is content`.

## 内部依赖

- `errors/gerror`
- `internal/json`

