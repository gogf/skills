# gmeta

> Package: `github.com/gogf/gf/v2/util/gmeta`

```go
import "github.com/gogf/gf/v2/util/gmeta"
```

## 概述

Package gmeta provides embedded meta data feature for struct.

## 源文件

- `gmeta.go`

## 类型定义

### Meta

**类型**: struct

Meta is used as an embedded attribute for struct to enabled metadata feature.


## 函数

### Data

```go
func Data(object any) map[string]string
```

Data retrieves and returns all metadata from `object`.

### Get

```go
func Get(object any, key string) *gvar.Var
```

Get retrieves and returns specified metadata by `key` from `object`.

## 内部依赖

- `container/gvar`
- `os/gstructs`

