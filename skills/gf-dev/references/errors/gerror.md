# gerror

> Package: `github.com/gogf/gf/v2/errors/gerror`

```go
import "github.com/gogf/gf/v2/errors/gerror"
```

## 概述

Package gerror provides rich functionalities to manipulate errors.

## 源文件

- `gerror.go`
- `gerror_api.go`
- `gerror_api_code.go`
- `gerror_api_option.go`
- `gerror_api_stack.go`
- `gerror_error.go`
- `gerror_error_code.go`
- `gerror_error_format.go`
- `gerror_error_json.go`
- `gerror_error_stack.go`

## 类型定义

### IEqual

**类型**: interface

IEqual is the interface for Equal feature.


### ICode

**类型**: interface

ICode is the interface for Code feature.


### IStack

**类型**: interface

IStack is the interface for Stack feature.


### ICause

**类型**: interface

ICause is the interface for Cause feature.


### ICurrent

**类型**: interface

ICurrent is the interface for Current feature.


### IUnwrap

**类型**: interface

IUnwrap is the interface for Unwrap feature.


### ITextArgs

**类型**: interface

ITextArgs is the interface for Text and Args features.
This interface is mainly used for i18n features, that needs text and args separately.


### Option

**类型**: struct

Option is option for creating error.


### Error

**类型**: struct

Error is custom error for additional features.


**方法**:

- `func (err *Error) Error() string`
  Error implements the interface of Error, it returns all the error as string.
- `func (err *Error) Cause() error`
  Cause returns the root cause error.
- `func (err *Error) Current() error`
  Current creates and returns the current level error.
- `func (err *Error) Unwrap() error`
  Unwrap is alias of function `Next`.
- `func (err *Error) Equal(target error) bool`
  Equal reports whether current error `err` equals to error `target`.
- `func (err *Error) TextWithArgs() string`
  TextWithArgs returns the formatted error text with its arguments.
- `func (err *Error) Text() string`
  Text returns the error text of current error.
- `func (err *Error) Args() []any`
  Args returns the error arguments of current error.
- `func (err *Error) Code() gcode.Code`
  Code returns the error code.
- `func (err *Error) SetCode(code gcode.Code)`
  SetCode updates the internal code with given code.
- `func (err *Error) Format(s fmt.State, verb rune)`
  Format formats the frame according to the fmt.Formatter interface.
- `func (err *Error) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface json.Marshaler for Error.
- `func (err *Error) Stack() string`
  Stack returns the error stack information as string.

## 函数

### New

```go
func New(text string) error
```

New creates and returns an error which is formatted from given text.

### Newf

```go
func Newf(format string, args ...any) error
```

Newf returns an error that formats as the given format and args.

### NewSkip

```go
func NewSkip(skip int, text string) error
```

NewSkip creates and returns an error which is formatted from given text.
The parameter `skip` specifies the stack callers skipped amount.

### NewSkipf

```go
func NewSkipf(skip int, format string, args ...any) error
```

NewSkipf returns an error that formats as the given format and args.
The parameter `skip` specifies the stack callers skipped amount.

### Wrap

```go
func Wrap(err error, text string) error
```

Wrap wraps error with text. It returns nil if given err is nil.
Note that it does not lose the error code of wrapped error, as it inherits the error code from it.

### Wrapf

```go
func Wrapf(err error, format string, args ...any) error
```

Wrapf returns an error annotating err with a stack trace at the point Wrapf is called, and the format specifier.
It returns nil if given `err` is nil.
Note that it does not lose the error code of wrapped error, as it inherits the error code from it.

### WrapSkip

```go
func WrapSkip(skip int, err error, text string) error
```

WrapSkip wraps error with text. It returns nil if given err is nil.
The parameter `skip` specifies the stack callers skipped amount.
Note that it does not lose the error code of wrapped error, as it inherits the error code from it.

### WrapSkipf

```go
func WrapSkipf(skip int, err error, format string, args ...any) error
```

WrapSkipf wraps error with text that is formatted with given format and args. It returns nil if given err is nil.
The parameter `skip` specifies the stack callers skipped amount.
Note that it does not lose the error code of wrapped error, as it inherits the error code from it.

### NewCode

```go
func NewCode(code gcode.Code, text ...string) error
```

NewCode creates and returns an error that has error code and given text.

### NewCodef

```go
func NewCodef(code gcode.Code, format string, args ...any) error
```

NewCodef returns an error that has error code and formats as the given format and args.

### NewCodeSkip

```go
func NewCodeSkip(code gcode.Code, skip int, text ...string) error
```

NewCodeSkip creates and returns an error which has error code and is formatted from given text.
The parameter `skip` specifies the stack callers skipped amount.

### NewCodeSkipf

```go
func NewCodeSkipf(code gcode.Code, skip int, format string, args ...any) error
```

NewCodeSkipf returns an error that has error code and formats as the given format and args.
The parameter `skip` specifies the stack callers skipped amount.

### WrapCode

```go
func WrapCode(code gcode.Code, err error, text ...string) error
```

WrapCode wraps error with code and text.
It returns nil if given err is nil.

### WrapCodef

```go
func WrapCodef(code gcode.Code, err error, format string, args ...any) error
```

WrapCodef wraps error with code and format specifier.
It returns nil if given `err` is nil.

### WrapCodeSkip

```go
func WrapCodeSkip(code gcode.Code, skip int, err error, text ...string) error
```

WrapCodeSkip wraps error with code and text.
It returns nil if given err is nil.
The parameter `skip` specifies the stack callers skipped amount.

### WrapCodeSkipf

```go
func WrapCodeSkipf(code gcode.Code, skip int, err error, format string, args ...any) error
```

WrapCodeSkipf wraps error with code and text that is formatted with given format and args.
It returns nil if given err is nil.
The parameter `skip` specifies the stack callers skipped amount.

### Code

```go
func Code(err error) gcode.Code
```

Code returns the error code of `current error`.
It returns `CodeNil` if it has no error code neither it does not implement interface Code.

### HasCode

```go
func HasCode(err error, code gcode.Code) bool
```

HasCode checks and reports whether `err` has `code` in its chaining errors.

### NewWithOption

```go
func NewWithOption(option Option) error
```

NewWithOption creates and returns a custom error with Option.
It is the senior usage for creating error, which is often used internally in framework.

### NewOption

```go
func NewOption(option Option) error
```

NewOption creates and returns a custom error with Option.

Deprecated: use NewWithOption instead.

### Cause

```go
func Cause(err error) error
```

Cause returns the root cause error of `err`.

### Stack

```go
func Stack(err error) string
```

Stack returns the stack callers as string.
It returns the error string directly if the `err` does not support stacks.

### Current

```go
func Current(err error) error
```

Current creates and returns the current level error.
It returns nil if current level error is nil.

### Unwrap

```go
func Unwrap(err error) error
```

Unwrap returns the next level error.
It returns nil if current level error or the next level error is nil.

### HasStack

```go
func HasStack(err error) bool
```

HasStack checks and reports whether `err` implemented interface `gerror.IStack`.

### Equal

```go
func Equal(err error, target error) bool
```

Equal reports whether current error `err` equals to error `target`.
Please note that, in default comparison logic for `Error`,
the errors are considered the same if both the `code` and `text` of them are the same.

### Is

```go
func Is(err error, target error) bool
```

Is reports whether current error `err` has error `target` in its chaining errors.
There's similar function HasError which is designed and implemented early before errors.Is of go stdlib.
It is now alias of errors.Is of go stdlib, to guarantee the same performance as go stdlib.

### As

```go
func As(err error, target any) bool
```

As finds the first error in err's chain that matches target, and if so, sets
target to that error value and returns true.

The chain consists of err itself followed by the sequence of errors obtained by
repeatedly calling Unwrap.

An error matches target if the error's concrete value is assignable to the value
pointed to by target, or if the error has a method As(any) bool such that
As(target) returns true. In the latter case, the As method is responsible for
setting target.

As will panic if target is not a non-nil pointer to either a type that implements
error, or to any interface type. As returns false if err is nil.

### HasError

```go
func HasError(err error, target error) bool
```

HasError performs as Is.
This function is designed and implemented early before errors.Is of go stdlib.

Deprecated: use Is instead.

## 内部依赖

- `errors/gcode`
- `internal/consts`
- `internal/errors`

