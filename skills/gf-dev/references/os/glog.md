# glog

> Package: `github.com/gogf/gf/v2/os/glog`

```go
import "github.com/gogf/gf/v2/os/glog"
```

## 概述

Package glog implements powerful and easy-to-use leveled logging functionality.

## 源文件

- `glog.go`
- `glog_api.go`
- `glog_chaining.go`
- `glog_config.go`
- `glog_instance.go`
- `glog_logger.go`
- `glog_logger_api.go`
- `glog_logger_chaining.go`
- `glog_logger_color.go`
- `glog_logger_config.go`
- `glog_logger_handler.go`
- `glog_logger_handler_json.go`
- `glog_logger_handler_structure.go`
- `glog_logger_level.go`
- `glog_logger_rotate.go`
- `glog_logger_writer.go`

## 类型定义

### ILogger

**类型**: interface

ILogger is the API interface for logger.


### Logger

**类型**: struct

Logger is the struct for logging management.


**方法**:

- `func (l *Logger) Clone() *Logger`
  Clone returns a new logger, which a `shallow copy` of the current logger.
- `func (l *Logger) PrintStack(ctx context.Context, skip ...int)`
  PrintStack prints the caller stack,
- `func (l *Logger) GetStack(skip ...int) string`
  GetStack returns the caller stack content,
- `func (l *Logger) Print(ctx context.Context, v ...any)`
  Print prints `v` with newline using fmt.Sprintln.
- `func (l *Logger) Printf(ctx context.Context, format string, v ...any)`
  Printf prints `v` with format `format` using fmt.Sprintf.
- `func (l *Logger) Fatal(ctx context.Context, v ...any)`
  Fatal prints the logging content with [FATA] header and newline, then exit th...
- `func (l *Logger) Fatalf(ctx context.Context, format string, v ...any)`
  Fatalf prints the logging content with [FATA] header, custom format and newli...
- `func (l *Logger) Panic(ctx context.Context, v ...any)`
  Panic prints the logging content with [PANI] header and newline, then panics.
- `func (l *Logger) Panicf(ctx context.Context, format string, v ...any)`
  Panicf prints the logging content with [PANI] header, custom format and newli...
- `func (l *Logger) Info(ctx context.Context, v ...any)`
  Info prints the logging content with [INFO] header and newline.
- `func (l *Logger) Infof(ctx context.Context, format string, v ...any)`
  Infof prints the logging content with [INFO] header, custom format and newline.
- `func (l *Logger) Debug(ctx context.Context, v ...any)`
  Debug prints the logging content with [DEBU] header and newline.
- `func (l *Logger) Debugf(ctx context.Context, format string, v ...any)`
  Debugf prints the logging content with [DEBU] header, custom format and newline.
- `func (l *Logger) Notice(ctx context.Context, v ...any)`
  Notice prints the logging content with [NOTI] header and newline.
- `func (l *Logger) Noticef(ctx context.Context, format string, v ...any)`
  Noticef prints the logging content with [NOTI] header, custom format and newl...
- `func (l *Logger) Warning(ctx context.Context, v ...any)`
  Warning prints the logging content with [WARN] header and newline.
- `func (l *Logger) Warningf(ctx context.Context, format string, v ...any)`
  Warningf prints the logging content with [WARN] header, custom format and new...
- `func (l *Logger) Error(ctx context.Context, v ...any)`
  Error prints the logging content with [ERRO] header and newline.
- `func (l *Logger) Errorf(ctx context.Context, format string, v ...any)`
  Errorf prints the logging content with [ERRO] header, custom format and newline.
- `func (l *Logger) Critical(ctx context.Context, v ...any)`
  Critical prints the logging content with [CRIT] header and newline.
- `func (l *Logger) Criticalf(ctx context.Context, format string, v ...any)`
  Criticalf prints the logging content with [CRIT] header, custom format and ne...
- `func (l *Logger) To(writer io.Writer) *Logger`
  To is a chaining function,
- `func (l *Logger) Path(path string) *Logger`
  Path is a chaining function,
- `func (l *Logger) Cat(category string) *Logger`
  Cat is a chaining function,
- `func (l *Logger) File(file string) *Logger`
  File is a chaining function,
- `func (l *Logger) Level(level int) *Logger`
  Level is a chaining function,
- `func (l *Logger) LevelStr(levelStr string) *Logger`
  LevelStr is a chaining function,
- `func (l *Logger) Skip(skip int) *Logger`
  Skip is a chaining function,
- `func (l *Logger) Stack(enabled bool, skip ...int) *Logger`
  Stack is a chaining function,
- `func (l *Logger) StackWithFilter(filter string) *Logger`
  StackWithFilter is a chaining function,
- `func (l *Logger) Stdout(enabled ...bool) *Logger`
  Stdout is a chaining function,
- `func (l *Logger) Header(enabled ...bool) *Logger`
  Header is a chaining function,
- `func (l *Logger) Line(long ...bool) *Logger`
  Line is a chaining function,
- `func (l *Logger) Async(enabled ...bool) *Logger`
  Async is a chaining function,
- `func (l *Logger) GetConfig() Config`
  GetConfig returns the configuration of current Logger.
- `func (l *Logger) SetConfig(config Config) error`
  SetConfig set configurations for the logger.
- `func (l *Logger) SetConfigWithMap(m map[string]any) error`
  SetConfigWithMap set configurations with map for the logger.
- `func (l *Logger) SetDebug(debug bool)`
  SetDebug enables/disables the debug level for logger.
- `func (l *Logger) SetAsync(enabled bool)`
  SetAsync enables/disables async logging output feature.
- `func (l *Logger) SetFlags(flags int)`
  SetFlags sets extra flags for logging output features.
- `func (l *Logger) GetFlags() int`
  GetFlags returns the flags of logger.
- `func (l *Logger) SetStack(enabled bool)`
  SetStack enables/disables the stack feature in failure logging outputs.
- `func (l *Logger) SetStackSkip(skip int)`
  SetStackSkip sets the stack offset from the end point.
- `func (l *Logger) SetStackFilter(filter string)`
  SetStackFilter sets the stack filter from the end point.
- `func (l *Logger) SetCtxKeys(keys ...any)`
  SetCtxKeys sets the context keys for logger. The keys is used for retrieving ...
- `func (l *Logger) AppendCtxKeys(keys ...any)`
  AppendCtxKeys appends extra keys to logger.
- `func (l *Logger) GetCtxKeys() []any`
  GetCtxKeys retrieves and returns the context keys for logging.
- `func (l *Logger) SetWriter(writer io.Writer)`
  SetWriter sets the customized logging `writer` for logging.
- `func (l *Logger) GetWriter() io.Writer`
  GetWriter returns the customized writer object, which implements the io.Write...
- `func (l *Logger) SetPath(path string) error`
  SetPath sets the directory path for file logging.
- `func (l *Logger) GetPath() string`
  GetPath returns the logging directory path for file logging.
- `func (l *Logger) SetFile(pattern string)`
  SetFile sets the file name `pattern` for file logging.
- `func (l *Logger) SetTimeFormat(timeFormat string)`
  SetTimeFormat sets the time format for the logging time.
- `func (l *Logger) SetStdoutPrint(enabled bool)`
  SetStdoutPrint sets whether output the logging contents to stdout, which is t...
- `func (l *Logger) SetHeaderPrint(enabled bool)`
  SetHeaderPrint sets whether output header of the logging contents, which is t...
- `func (l *Logger) SetLevelPrint(enabled bool)`
  SetLevelPrint sets whether output level string of the logging contents, which...
- `func (l *Logger) SetPrefix(prefix string)`
  SetPrefix sets prefix string for every logging content.
- `func (l *Logger) SetHandlers(handlers ...Handler)`
  SetHandlers sets the logging handlers for current logger.
- `func (l *Logger) SetWriterColorEnable(enabled bool)`
  SetWriterColorEnable enables file/writer logging with color.
- `func (l *Logger) SetStdoutColorDisabled(disabled bool)`
  SetStdoutColorDisabled disables stdout logging with color.
- `func (l *Logger) SetLevel(level int)`
  SetLevel sets the logging level.
- `func (l *Logger) GetLevel() int`
  GetLevel returns the logging level value.
- `func (l *Logger) SetLevelStr(levelStr string) error`
  SetLevelStr sets the logging level by level string.
- `func (l *Logger) SetLevelPrefix(level int, prefix string)`
  SetLevelPrefix sets the prefix string for specified level.
- `func (l *Logger) SetLevelPrefixes(prefixes map[int]string)`
  SetLevelPrefixes sets the level to prefix string mapping for the logger.
- `func (l *Logger) GetLevelPrefix(level int) string`
  GetLevelPrefix returns the prefix string for specified level.
- `func (l *Logger) Write(p []byte) (n int, err error)`
  Write implements the io.Writer interface.

### Config

**类型**: struct

Config is the configuration object for logger.


### Handler

**类型**: type

Handler is function handler for custom logging content outputs.


### HandlerInput

**类型**: struct

HandlerInput is the input parameter struct for logging Handler.

The logging content is consisted in:
TimeFormat [LevelFormat] {TraceId} {CtxStr} Prefix CallerFunc CallerPath Content Values Stack

The header in the logging content is:
TimeFormat [LevelFormat] {TraceId} {CtxStr} Prefix CallerFunc CallerPath


**方法**:

- `func (in *HandlerInput) Next(ctx context.Context)`
  Next calls the next logging handler in middleware way.
- `func (in *HandlerInput) String(withColor ...bool) string`
  String returns the logging content formatted by default logging handler.
- `func (in *HandlerInput) ValuesContent() string`
  ValuesContent converts and returns values as string content.

### HandlerOutputJson

**类型**: struct

HandlerOutputJson is the structure outputting logging content as single json.


## 函数

### DefaultLogger

```go
func DefaultLogger() *Logger
```

DefaultLogger returns the default logger.

### SetDefaultLogger

```go
func SetDefaultLogger(l *Logger)
```

SetDefaultLogger sets the default logger for package glog.
Note that there might be concurrent safety issue if calls this function
in different goroutines.

### Print

```go
func Print(ctx context.Context, v ...any)
```

Print prints `v` with newline using fmt.Sprintln.
The parameter `v` can be multiple variables.

### Printf

```go
func Printf(ctx context.Context, format string, v ...any)
```

Printf prints `v` with format `format` using fmt.Sprintf.
The parameter `v` can be multiple variables.

### Fatal

```go
func Fatal(ctx context.Context, v ...any)
```

Fatal prints the logging content with [FATA] header and newline, then exit the current process.

### Fatalf

```go
func Fatalf(ctx context.Context, format string, v ...any)
```

Fatalf prints the logging content with [FATA] header, custom format and newline, then exit the current process.

### Panic

```go
func Panic(ctx context.Context, v ...any)
```

Panic prints the logging content with [PANI] header and newline, then panics.

### Panicf

```go
func Panicf(ctx context.Context, format string, v ...any)
```

Panicf prints the logging content with [PANI] header, custom format and newline, then panics.

### Info

```go
func Info(ctx context.Context, v ...any)
```

Info prints the logging content with [INFO] header and newline.

### Infof

```go
func Infof(ctx context.Context, format string, v ...any)
```

Infof prints the logging content with [INFO] header, custom format and newline.

### Debug

```go
func Debug(ctx context.Context, v ...any)
```

Debug prints the logging content with [DEBU] header and newline.

### Debugf

```go
func Debugf(ctx context.Context, format string, v ...any)
```

Debugf prints the logging content with [DEBU] header, custom format and newline.

### Notice

```go
func Notice(ctx context.Context, v ...any)
```

Notice prints the logging content with [NOTI] header and newline.
It also prints caller stack info if stack feature is enabled.

### Noticef

```go
func Noticef(ctx context.Context, format string, v ...any)
```

Noticef prints the logging content with [NOTI] header, custom format and newline.
It also prints caller stack info if stack feature is enabled.

### Warning

```go
func Warning(ctx context.Context, v ...any)
```

Warning prints the logging content with [WARN] header and newline.
It also prints caller stack info if stack feature is enabled.

### Warningf

```go
func Warningf(ctx context.Context, format string, v ...any)
```

Warningf prints the logging content with [WARN] header, custom format and newline.
It also prints caller stack info if stack feature is enabled.

### Error

```go
func Error(ctx context.Context, v ...any)
```

Error prints the logging content with [ERRO] header and newline.
It also prints caller stack info if stack feature is enabled.

### Errorf

```go
func Errorf(ctx context.Context, format string, v ...any)
```

Errorf prints the logging content with [ERRO] header, custom format and newline.
It also prints caller stack info if stack feature is enabled.

### Critical

```go
func Critical(ctx context.Context, v ...any)
```

Critical prints the logging content with [CRIT] header and newline.
It also prints caller stack info if stack feature is enabled.

### Criticalf

```go
func Criticalf(ctx context.Context, format string, v ...any)
```

Criticalf prints the logging content with [CRIT] header, custom format and newline.
It also prints caller stack info if stack feature is enabled.

### Expose

```go
func Expose() *Logger
```

Expose returns the default logger of package glog.

### To

```go
func To(writer io.Writer) *Logger
```

To is a chaining function,
which redirects current logging content output to the sepecified `writer`.

### Path

```go
func Path(path string) *Logger
```

Path is a chaining function,
which sets the directory path to `path` for current logging content output.

### Cat

```go
func Cat(category string) *Logger
```

Cat is a chaining function,
which sets the category to `category` for current logging content output.

### File

```go
func File(pattern string) *Logger
```

File is a chaining function,
which sets file name `pattern` for the current logging content output.

### Level

```go
func Level(level int) *Logger
```

Level is a chaining function,
which sets logging level for the current logging content output.

### LevelStr

```go
func LevelStr(levelStr string) *Logger
```

LevelStr is a chaining function,
which sets logging level for the current logging content output using level string.

### Skip

```go
func Skip(skip int) *Logger
```

Skip is a chaining function,
which sets stack skip for the current logging content output.
It also affects the caller file path checks when line number printing enabled.

### Stack

```go
func Stack(enabled bool, skip ...int) *Logger
```

Stack is a chaining function,
which sets stack options for the current logging content output .

### StackWithFilter

```go
func StackWithFilter(filter string) *Logger
```

StackWithFilter is a chaining function,
which sets stack filter for the current logging content output .

### Stdout

```go
func Stdout(enabled ...bool) *Logger
```

Stdout is a chaining function,
which enables/disables stdout for the current logging content output.
It's enabled in default.

### Header

```go
func Header(enabled ...bool) *Logger
```

Header is a chaining function,
which enables/disables log header for the current logging content output.
It's enabled in default.

### Line

```go
func Line(long ...bool) *Logger
```

Line is a chaining function,
which enables/disables printing its caller file along with its line number.
The parameter `long` specified whether print the long absolute file path, eg: /a/b/c/d.go:23.

### Async

```go
func Async(enabled ...bool) *Logger
```

Async is a chaining function,
which enables/disables async logging output feature.

### SetConfig

```go
func SetConfig(config Config) error
```

SetConfig set configurations for the defaultLogger.

### SetConfigWithMap

```go
func SetConfigWithMap(m map[string]any) error
```

SetConfigWithMap set configurations with map for the defaultLogger.

### SetPath

```go
func SetPath(path string) error
```

SetPath sets the directory path for file logging.

### GetPath

```go
func GetPath() string
```

GetPath returns the logging directory path for file logging.
It returns empty string if no directory path set.

### SetFile

```go
func SetFile(pattern string)
```

SetFile sets the file name `pattern` for file logging.
Datetime pattern can be used in `pattern`, eg: access-{Ymd}.log.
The default file name pattern is: Y-m-d.log, eg: 2018-01-01.log

### SetLevel

```go
func SetLevel(level int)
```

SetLevel sets the default logging level.

### GetLevel

```go
func GetLevel() int
```

GetLevel returns the default logging level value.

### SetWriter

```go
func SetWriter(writer io.Writer)
```

SetWriter sets the customized logging `writer` for logging.
The `writer` object should implements the io.Writer interface.
Developer can use customized logging `writer` to redirect logging output to another service,
eg: kafka, mysql, mongodb, etc.

### GetWriter

```go
func GetWriter() io.Writer
```

GetWriter returns the customized writer object, which implements the io.Writer interface.
It returns nil if no customized writer set.

### SetDebug

```go
func SetDebug(debug bool)
```

SetDebug enables/disables the debug level for default defaultLogger.
The debug level is enabled in default.

### SetAsync

```go
func SetAsync(enabled bool)
```

SetAsync enables/disables async logging output feature for default defaultLogger.

### SetStdoutPrint

```go
func SetStdoutPrint(enabled bool)
```

SetStdoutPrint sets whether ouptput the logging contents to stdout, which is true in default.

### SetHeaderPrint

```go
func SetHeaderPrint(enabled bool)
```

SetHeaderPrint sets whether output header of the logging contents, which is true in default.

### SetPrefix

```go
func SetPrefix(prefix string)
```

SetPrefix sets prefix string for every logging content.
Prefix is part of header, which means if header output is shut, no prefix will be output.

### SetFlags

```go
func SetFlags(flags int)
```

SetFlags sets extra flags for logging output features.

### GetFlags

```go
func GetFlags() int
```

GetFlags returns the flags of defaultLogger.

### SetCtxKeys

```go
func SetCtxKeys(keys ...any)
```

SetCtxKeys sets the context keys for defaultLogger. The keys is used for retrieving values
from context and printing them to logging content.

Note that multiple calls of this function will overwrite the previous set context keys.

### GetCtxKeys

```go
func GetCtxKeys() []any
```

GetCtxKeys retrieves and returns the context keys for logging.

### PrintStack

```go
func PrintStack(ctx context.Context, skip ...int)
```

PrintStack prints the caller stack,
the optional parameter `skip` specify the skipped stack offset from the end point.

### GetStack

```go
func GetStack(skip ...int) string
```

GetStack returns the caller stack content,
the optional parameter `skip` specify the skipped stack offset from the end point.

### SetStack

```go
func SetStack(enabled bool)
```

SetStack enables/disables the stack feature in failure logging outputs.

### SetLevelStr

```go
func SetLevelStr(levelStr string) error
```

SetLevelStr sets the logging level by level string.

### SetLevelPrefix

```go
func SetLevelPrefix(level int, prefix string)
```

SetLevelPrefix sets the prefix string for specified level.

### SetLevelPrefixes

```go
func SetLevelPrefixes(prefixes map[int]string)
```

SetLevelPrefixes sets the level to prefix string mapping for the defaultLogger.

### GetLevelPrefix

```go
func GetLevelPrefix(level int) string
```

GetLevelPrefix returns the prefix string for specified level.

### SetHandlers

```go
func SetHandlers(handlers ...Handler)
```

SetHandlers sets the logging handlers for default defaultLogger.

### SetWriterColorEnable

```go
func SetWriterColorEnable(enabled bool)
```

SetWriterColorEnable sets the file logging with color

### Instance

```go
func Instance(name ...string) *Logger
```

Instance returns an instance of Logger with default settings.
The parameter `name` is the name for the instance.

### New

```go
func New() *Logger
```

New creates and returns a custom logger.

### NewWithWriter

```go
func NewWithWriter(writer io.Writer) *Logger
```

NewWithWriter creates and returns a custom logger with io.Writer.

### DefaultConfig

```go
func DefaultConfig() Config
```

DefaultConfig returns the default configuration for logger.

### SetDefaultHandler

```go
func SetDefaultHandler(handler Handler)
```

SetDefaultHandler sets default handler for package.

### GetDefaultHandler

```go
func GetDefaultHandler() Handler
```

GetDefaultHandler returns the default handler of package.

### HandlerJson

```go
func HandlerJson(ctx context.Context, in *HandlerInput)
```

HandlerJson is a handler for output logging content as a single json string.

### HandlerStructure

```go
func HandlerStructure(ctx context.Context, in *HandlerInput)
```

HandlerStructure is a handler for output logging content as a structured string.

## 内部依赖

- `container/garray`
- `container/gmap`
- `container/gtype`
- `debug/gdebug`
- `encoding/gcompress`
- `errors/gcode`
- `errors/gerror`
- `internal/command`
- `internal/consts`
- `internal/errors`
- `internal/intlog`
- `internal/json`
- `os/gctx`
- `os/gfile`
- `os/gfpool`
- `os/gmlock`
- `os/grpool`
- `os/gtime`
- `os/gtimer`
- `text/gregex`
- `util/gconv`
- `util/gutil`

## 外部依赖

- `github.com/fatih/color`
- `go.opentelemetry.io/otel/trace`

