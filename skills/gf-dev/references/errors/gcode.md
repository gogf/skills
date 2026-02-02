# gcode

> Package: `github.com/gogf/gf/v2/errors/gcode`

```go
import "github.com/gogf/gf/v2/errors/gcode"
```

## 概述

Package gcode provides universal error code definition and common error codes implements.

## 源文件

- `gcode.go`
- `gcode_local.go`

## 类型定义

### Code

**类型**: interface

Code is universal error code interface definition.


## 函数

### New

```go
func New(code int, message string, detail any) Code
```

New creates and returns an error code.
Note that it returns an interface object of Code.

### WithCode

```go
func WithCode(code Code, detail any) Code
```

WithCode creates and returns a new error code based on given Code.
The code and message is from given `code`, but the detail if from given `detail`.

