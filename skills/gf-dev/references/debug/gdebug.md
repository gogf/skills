# gdebug

> Package: `github.com/gogf/gf/v2/debug/gdebug`

```go
import "github.com/gogf/gf/v2/debug/gdebug"
```

## 概述

Package gdebug contains facilities for programs to debug themselves while they are running.

## 源文件

- `gdebug.go`
- `gdebug_caller.go`
- `gdebug_grid.go`
- `gdebug_stack.go`
- `gdebug_version.go`

## 函数

### Caller

```go
func Caller(skip ...int) (function string, path string, line int)
```

Caller returns the function name and the absolute file path along with its line
number of the caller.

### CallerWithFilter

```go
func CallerWithFilter(filters []string, skip ...int) (function string, path string, line int)
```

CallerWithFilter returns the function name and the absolute file path along with
its line number of the caller.

The parameter `filters` is used to filter the path of the caller.

### CallerPackage

```go
func CallerPackage() string
```

CallerPackage returns the package name of the caller.

### CallerFunction

```go
func CallerFunction() string
```

CallerFunction returns the function name of the caller.

### CallerFilePath

```go
func CallerFilePath() string
```

CallerFilePath returns the file path of the caller.

### CallerDirectory

```go
func CallerDirectory() string
```

CallerDirectory returns the directory of the caller.

### CallerFileLine

```go
func CallerFileLine() string
```

CallerFileLine returns the file path along with the line number of the caller.

### CallerFileLineShort

```go
func CallerFileLineShort() string
```

CallerFileLineShort returns the file name along with the line number of the caller.

### FuncPath

```go
func FuncPath(f any) string
```

FuncPath returns the complete function path of given `f`.

### FuncName

```go
func FuncName(f any) string
```

FuncName returns the function name of given `f`.

### GoroutineID

```go
func GoroutineID() int
```

GoroutineID retrieves and returns the current goroutine id from stack information.
Be very aware that, it is with low performance as it uses runtime.Stack function.
It is commonly used for debugging purpose.

### PrintStack

```go
func PrintStack(skip ...int)
```

PrintStack prints to standard error the stack trace returned by runtime.Stack.

### Stack

```go
func Stack(skip ...int) string
```

Stack returns a formatted stack trace of the goroutine that calls it.
It calls runtime.Stack with a large enough buffer to capture the entire trace.

### StackWithFilter

```go
func StackWithFilter(filters []string, skip ...int) string
```

StackWithFilter returns a formatted stack trace of the goroutine that calls it.
It calls runtime.Stack with a large enough buffer to capture the entire trace.

The parameter `filter` is used to filter the path of the caller.

### StackWithFilters

```go
func StackWithFilters(filters []string, skip ...int) string
```

StackWithFilters returns a formatted stack trace of the goroutine that calls it.
It calls runtime.Stack with a large enough buffer to capture the entire trace.

The parameter `filters` is a slice of strings, which are used to filter the path of the
caller.

TODO Improve the performance using debug.Stack.

### BinVersion

```go
func BinVersion() string
```

BinVersion returns the version of current running binary.
It uses ghash.BKDRHash+BASE36 algorithm to calculate the unique version of the binary.

### BinVersionMd5

```go
func BinVersionMd5() string
```

BinVersionMd5 returns the version of current running binary.
It uses MD5 algorithm to calculate the unique version of the binary.

## 内部依赖

- `encoding/ghash`
- `errors/gerror`

