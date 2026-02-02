# gfcmd

> Package: `github.com/gogf/gf/cmd/gf/v2/gfcmd`

```go
import "github.com/gogf/gf/cmd/gf/v2/gfcmd"
```

## 概述

Package gfcmd provides the management of CLI commands for `gf` tool.

## 所属模块

此包属于子模块: `github.com/gogf/gf/cmd/gf/v2`

## 源文件

- `gfcmd.go`

## 类型定义

### Command

**类型**: struct

Command manages the CLI command of `gf`.
This struct can be globally accessible and extended with custom struct.


**方法**:

- `func (c *Command) Run(ctx context.Context)`
  Run starts running the command according the command line arguments and options.

## 函数

### GetCommand

```go
func GetCommand(ctx context.Context) (*Command, error)
```

GetCommand retrieves and returns the root command of CLI `gf`.

## 内部依赖

- `internal/cmd`
- `internal/cmd/cmddep`
- `internal/packed`
- `internal/utility/allyes`
- `internal/utility/mlog`
- `errors/gcode`
- `errors/gerror`
- `frame/g`
- `os/gcfg`
- `os/gcmd`
- `os/gfile`
- `text/gstr`

