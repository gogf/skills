# gcmd

> Package: `github.com/gogf/gf/v2/os/gcmd`

```go
import "github.com/gogf/gf/v2/os/gcmd"
```

## 概述

Package gcmd provides console operations, like options/arguments reading and command running.

## 源文件

- `gcmd.go`
- `gcmd_command.go`
- `gcmd_command_help.go`
- `gcmd_command_object.go`
- `gcmd_command_run.go`
- `gcmd_parser.go`
- `gcmd_scan.go`

## 类型定义

### Command

**类型**: struct

Command holds the info about an argument that can handle custom logic.


**方法**:

- `func (c *Command) AddCommand(commands ...*Command) error`
  AddCommand adds one or more sub-commands to current command.
- `func (c *Command) AddObject(objects ...any) error`
  AddObject adds one or more sub-commands to current command using struct object.
- `func (c *Command) Print()`
  Print prints help info to stdout for current command.
- `func (c *Command) PrintTo(writer io.Writer)`
  PrintTo prints help info to custom io.Writer.
- `func (c *Command) Run(ctx context.Context)`
  Run calls custom function in os.Args that bound to this command.
- `func (c *Command) RunWithValue(ctx context.Context) value any`
  RunWithValue calls custom function in os.Args that bound to this command with...
- `func (c *Command) RunWithError(ctx context.Context) err error`
  RunWithError calls custom function in os.Args that bound to this command with...
- `func (c *Command) RunWithValueError(ctx context.Context) (value any, err error)`
  RunWithValueError calls custom function in os.Args that bound to this command...
- `func (c *Command) RunWithSpecificArgs(ctx context.Context, args []string) (value any, err error)`
  RunWithSpecificArgs calls custom function in specific args that bound to this...

### Function

**类型**: type

Function is a custom command callback function that is bound to a certain argument.


### FuncWithValue

**类型**: type

FuncWithValue is similar like Func but with output parameters that can interact with command caller.


### Argument

**类型**: struct

Argument is the command value that are used by certain command.


### ParserOption

**类型**: struct

ParserOption manages the parsing options.


### Parser

**类型**: struct

Parser for arguments.


**方法**:

- `func (p *Parser) GetOpt(name string, def ...any) *gvar.Var`
  GetOpt returns the option value named `name` as gvar.Var.
- `func (p *Parser) GetOptAll() map[string]string`
  GetOptAll returns all parsed options.
- `func (p *Parser) GetArg(index int, def ...string) *gvar.Var`
  GetArg returns the argument at `index` as gvar.Var.
- `func (p *Parser) GetArgAll() []string`
  GetArgAll returns all parsed arguments.
- `func (p *Parser) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.

## 函数

### Init

```go
func Init(args ...string)
```

Init does custom initialization.

### GetOpt

```go
func GetOpt(name string, def ...string) *gvar.Var
```

GetOpt returns the option value named `name` as gvar.Var.

### GetOptAll

```go
func GetOptAll() map[string]string
```

GetOptAll returns all parsed options.

### GetArg

```go
func GetArg(index int, def ...string) *gvar.Var
```

GetArg returns the argument at `index` as gvar.Var.

### GetArgAll

```go
func GetArgAll() []string
```

GetArgAll returns all parsed arguments.

### GetOptWithEnv

```go
func GetOptWithEnv(key string, def ...any) *gvar.Var
```

GetOptWithEnv returns the command line argument of the specified `key`.
If the argument does not exist, then it returns the environment variable with specified `key`.
It returns the default value `def` if none of them exists.

Fetching Rules:
1. Command line arguments are in lowercase format, eg: gf.`package name`.<variable name>;
2. Environment arguments are in uppercase format, eg: GF_`package name`_<variable name>；

### BuildOptions

```go
func BuildOptions(m map[string]string, prefix ...string) string
```

BuildOptions builds the options as string.

### CommandFromCtx

```go
func CommandFromCtx(ctx context.Context) *Command
```

CommandFromCtx retrieves and returns Command from context.

### NewFromObject

```go
func NewFromObject(object any) (rootCmd *Command, err error)
```

NewFromObject creates and returns a root command object using given object.

### ParserFromCtx

```go
func ParserFromCtx(ctx context.Context) *Parser
```

ParserFromCtx retrieves and returns Parser from context.

### Parse

```go
func Parse(supportedOptions map[string]bool, option ...ParserOption) (*Parser, error)
```

Parse creates and returns a new Parser with os.Args and supported options.

Note that the parameter `supportedOptions` is as [option name: need argument], which means
the value item of `supportedOptions` indicates whether corresponding option name needs argument or not.

The optional parameter `strict` specifies whether stops parsing and returns error if invalid option passed.

### ParseArgs

```go
func ParseArgs(args []string, supportedOptions map[string]bool, option ...ParserOption) (*Parser, error)
```

ParseArgs creates and returns a new Parser with given arguments and supported options.

Note that the parameter `supportedOptions` is as [option name: need argument], which means
the value item of `supportedOptions` indicates whether corresponding option name needs argument or not.

The optional parameter `strict` specifies whether stops parsing and returns error if invalid option passed.

### Scan

```go
func Scan(info ...any) string
```

Scan prints `info` to stdout, reads and returns user input, which stops by '\n'.

### Scanf

```go
func Scanf(format string, info ...any) string
```

Scanf prints `info` to stdout with `format`, reads and returns user input, which stops by '\n'.

## 内部依赖

- `github.com/gogf/gf/v2`
- `container/gset`
- `container/gvar`
- `encoding/gjson`
- `errors/gcode`
- `errors/gerror`
- `internal/command`
- `internal/intlog`
- `internal/json`
- `internal/reflection`
- `internal/utils`
- `net/gtrace`
- `os/gcfg`
- `os/gctx`
- `os/genv`
- `os/glog`
- `os/gstructs`
- `text/gregex`
- `text/gstr`
- `util/gconv`
- `util/gmeta`
- `util/gtag`
- `util/gutil`
- `util/gvalid`

## 外部依赖

- `go.opentelemetry.io/otel`
- `go.opentelemetry.io/otel/propagation`
- `go.opentelemetry.io/otel/trace`

