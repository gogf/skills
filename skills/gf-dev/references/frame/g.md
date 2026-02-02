# g

> Package: `github.com/gogf/gf/v2/frame/g`

```go
import "github.com/gogf/gf/v2/frame/g"
```

## 概述

Package g provides commonly used type/function defines and coupled calling for creating commonly-used objects.

## 源文件

- `g.go`
- `g_func.go`
- `g_object.go`
- `g_setting.go`

## 类型定义

### Var

**类型**: type

### Ctx

**类型**: type

### Meta

**类型**: type

### Map

**类型**: type

### MapAnyAny

**类型**: type

### MapAnyStr

**类型**: type

### MapAnyInt

**类型**: type

### MapStrAny

**类型**: type

### MapStrStr

**类型**: type

### MapStrInt

**类型**: type

### MapIntAny

**类型**: type

### MapIntStr

**类型**: type

### MapIntInt

**类型**: type

### MapAnyBool

**类型**: type

### MapStrBool

**类型**: type

### MapIntBool

**类型**: type

### List

**类型**: type

### ListAnyAny

**类型**: type

### ListAnyStr

**类型**: type

### ListAnyInt

**类型**: type

### ListStrAny

**类型**: type

### ListStrStr

**类型**: type

### ListStrInt

**类型**: type

### ListIntAny

**类型**: type

### ListIntStr

**类型**: type

### ListIntInt

**类型**: type

### ListAnyBool

**类型**: type

### ListStrBool

**类型**: type

### ListIntBool

**类型**: type

### Slice

**类型**: type

### SliceAny

**类型**: type

### SliceStr

**类型**: type

### SliceInt

**类型**: type

### Array

**类型**: type

### ArrayAny

**类型**: type

### ArrayStr

**类型**: type

### ArrayInt

**类型**: type

## 函数

### Go

```go
func Go(ctx context.Context, goroutineFunc func, recoverFunc func)
```

Go creates a new asynchronous goroutine function with specified recover function.

The parameter `recoverFunc` is called when any panic during executing of `goroutineFunc`.
If `recoverFunc` is given nil, it ignores the panic from `goroutineFunc` and no panic will
throw to parent goroutine.

But, note that, if `recoverFunc` also throws panic, such panic will be thrown to parent goroutine.

### NewVar

```go
func NewVar(i any, safe ...bool) *Var
```

NewVar returns a gvar.Var.

### Wait

```go
func Wait()
```

Wait is an alias of ghttp.Wait, which blocks until all the web servers shutdown.
It's commonly used in multiple servers' situation.

### Listen

```go
func Listen()
```

Listen is an alias of gproc.Listen, which handles the signals received and automatically
calls registered signal handler functions.
It blocks until shutdown signals received and all registered shutdown handlers done.

### Dump

```go
func Dump(values ...any)
```

Dump dumps a variable to stdout with more manually readable.

### DumpTo

```go
func DumpTo(writer io.Writer, value any, option gutil.DumpOption)
```

DumpTo writes variables `values` as a string in to `writer` with more manually readable

### DumpWithType

```go
func DumpWithType(values ...any)
```

DumpWithType acts like Dump, but with type information.
Also see Dump.

### DumpWithOption

```go
func DumpWithOption(value any, option gutil.DumpOption)
```

DumpWithOption returns variables `values` as a string with more manually readable.

### DumpJson

```go
func DumpJson(value any)
```

DumpJson pretty dumps json content to stdout.

### Throw

```go
func Throw(exception any)
```

Throw throws an exception, which can be caught by TryCatch function.

### Try

```go
func Try(ctx context.Context, try func) err error
```

Try implements try... logistics using internal panic...recover.
It returns error if any exception occurs, or else it returns nil.

### TryCatch

```go
func TryCatch(ctx context.Context, try func, catch func)
```

TryCatch implements try...catch... logistics using internal panic...recover.
It automatically calls function `catch` if any exception occurs and passes the exception as an error.

But, note that, if function `catch` also throws panic, the current goroutine will panic.

### IsNil

```go
func IsNil(value any, traceSource ...bool) bool
```

IsNil checks whether given `value` is nil.
Parameter `traceSource` is used for tracing to the source variable if given `value` is type
of pointer that also points to a pointer. It returns nil if the source is nil when `traceSource`
is true.
Note that it might use reflect feature which affects performance a little.

### IsEmpty

```go
func IsEmpty(value any, traceSource ...bool) bool
```

IsEmpty checks whether given `value` empty.
It returns true if `value` is in: 0, nil, false, "", len(slice/map/chan) == 0.
Or else it returns true.

The parameter `traceSource` is used for tracing to the source variable if given `value` is type of pointer
that also points to a pointer. It returns true if the source is empty when `traceSource` is true.
Note that it might use reflect feature which affects performance a little.

### RequestFromCtx

```go
func RequestFromCtx(ctx context.Context) *ghttp.Request
```

RequestFromCtx retrieves and returns the Request object from context.

### Client

```go
func Client() *gclient.Client
```

Client is a convenience function, which creates and returns a new HTTP client.

### Server

```go
func Server(name ...any) *ghttp.Server
```

Server returns an instance of http server with specified name.

### TCPServer

```go
func TCPServer(name ...any) *gtcp.Server
```

TCPServer returns an instance of tcp server with specified name.

### UDPServer

```go
func UDPServer(name ...any) *gudp.Server
```

UDPServer returns an instance of udp server with specified name.

### View

```go
func View(name ...string) *gview.View
```

View returns an instance of template engine object with specified name.

### Config

```go
func Config(name ...string) *gcfg.Config
```

Config returns an instance of config object with specified name.

### Cfg

```go
func Cfg(name ...string) *gcfg.Config
```

Cfg is alias of Config.
See Config.

### Resource

```go
func Resource(name ...string) *gres.Resource
```

Resource returns an instance of Resource.
The parameter `name` is the name for the instance.

### I18n

```go
func I18n(name ...string) *gi18n.Manager
```

I18n returns an instance of gi18n.Manager.
The parameter `name` is the name for the instance.

### Res

```go
func Res(name ...string) *gres.Resource
```

Res is alias of Resource.
See Resource.

### Log

```go
func Log(name ...string) *glog.Logger
```

Log returns an instance of glog.Logger.
The parameter `name` is the name for the instance.

### DB

```go
func DB(name ...string) gdb.DB
```

DB returns an instance of database ORM object with specified configuration group name.

### Model

```go
func Model(tableNameOrStruct ...any) *gdb.Model
```

Model creates and returns a model based on configuration of default database group.

### ModelRaw

```go
func ModelRaw(rawSql string, args ...any) *gdb.Model
```

ModelRaw creates and returns a model based on a raw sql not a table.

### Redis

```go
func Redis(name ...string) *gredis.Redis
```

Redis returns an instance of redis client with specified configuration group name.

### Validator

```go
func Validator() *gvalid.Validator
```

Validator is a convenience function, which creates and returns a new validation manager object.

### SetDebug

```go
func SetDebug(enabled bool)
```

SetDebug enables/disables the GoFrame internal logging manually.
Note that this function is not concurrent safe, be aware of the DATA RACE,
which means you should call this function in your boot but not the runtime.

## 内部依赖

- `container/gvar`
- `database/gdb`
- `database/gredis`
- `frame/gins`
- `i18n/gi18n`
- `internal/empty`
- `internal/utils`
- `net/gclient`
- `net/ghttp`
- `net/gtcp`
- `net/gudp`
- `os/gcfg`
- `os/glog`
- `os/gproc`
- `os/gres`
- `os/gview`
- `util/gmeta`
- `util/gutil`
- `util/gvalid`

