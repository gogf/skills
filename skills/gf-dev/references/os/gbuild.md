# gbuild

> Package: `github.com/gogf/gf/v2/os/gbuild`

```go
import "github.com/gogf/gf/v2/os/gbuild"
```

## 概述

Package gbuild manages the build-in variables from "gf build".

## 源文件

- `gbuild.go`

## 类型定义

### BuildInfo

**类型**: struct

BuildInfo maintains the built info of the current binary.


## 函数

### Info

```go
func Info() BuildInfo
```

Info returns the basic built information of the binary as map.
Note that it should be used with gf-cli tool "gf build",
which automatically injects necessary information into the binary.

### Get

```go
func Get(name string, def ...any) *gvar.Var
```

Get retrieves and returns the build-in binary variable with given name.

### Data

```go
func Data() map[string]any
```

Data returns the custom build-in variables as the map.

## 内部依赖

- `github.com/gogf/gf/v2`
- `container/gvar`
- `encoding/gbase64`
- `internal/intlog`
- `internal/json`

