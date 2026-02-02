# gtest

> Package: `github.com/gogf/gf/v2/test/gtest`

```go
import "github.com/gogf/gf/v2/test/gtest"
```

## 概述

Package gtest provides convenient test utilities for unit testing.

## 源文件

- `gtest.go`
- `gtest_t.go`
- `gtest_util.go`

## 类型定义

### T

**类型**: struct

T is the testing unit case management object.


**方法**:

- `func (t *T) Assert(value any, expect any)`
  Assert checks `value` and `expect` EQUAL.
- `func (t *T) AssertEQ(value any, expect any)`
  AssertEQ checks `value` and `expect` EQUAL, including their TYPES.
- `func (t *T) AssertNE(value any, expect any)`
  AssertNE checks `value` and `expect` NOT EQUAL.
- `func (t *T) AssertNQ(value any, expect any)`
  AssertNQ checks `value` and `expect` NOT EQUAL, including their TYPES.
- `func (t *T) AssertGT(value any, expect any)`
  AssertGT checks `value` is GREATER THAN `expect`.
- `func (t *T) AssertGE(value any, expect any)`
  AssertGE checks `value` is GREATER OR EQUAL THAN `expect`.
- `func (t *T) AssertLT(value any, expect any)`
  AssertLT checks `value` is LESS EQUAL THAN `expect`.
- `func (t *T) AssertLE(value any, expect any)`
  AssertLE checks `value` is LESS OR EQUAL THAN `expect`.
- `func (t *T) AssertIN(value any, expect any)`
  AssertIN checks `value` is IN `expect`.
- `func (t *T) AssertNI(value any, expect any)`
  AssertNI checks `value` is NOT IN `expect`.
- `func (t *T) AssertNil(value any)`
  AssertNil asserts `value` is nil.
- `func (t *T) Error(message ...any)`
  Error panics with given `message`.
- `func (t *T) Fatal(message ...any)`
  Fatal prints `message` to stderr and exit the process.

## 函数

### C

```go
func C(t *testing.T, f func)
```

C creates a unit testing case.
The parameter `t` is the pointer to testing.T of stdlib (*testing.T).
The parameter `f` is the closure function for unit testing case.

### Assert

```go
func Assert(value any, expect any)
```

Assert checks `value` and `expect` EQUAL.

### AssertEQ

```go
func AssertEQ(value any, expect any)
```

AssertEQ checks `value` and `expect` EQUAL, including their TYPES.

### AssertNE

```go
func AssertNE(value any, expect any)
```

AssertNE checks `value` and `expect` NOT EQUAL.

### AssertNQ

```go
func AssertNQ(value any, expect any)
```

AssertNQ checks `value` and `expect` NOT EQUAL, including their TYPES.

### AssertGT

```go
func AssertGT(value any, expect any)
```

AssertGT checks `value` is GREATER THAN `expect`.
Notice that, only string, integer and float types can be compared by AssertGT,
others are invalid.

### AssertGE

```go
func AssertGE(value any, expect any)
```

AssertGE checks `value` is GREATER OR EQUAL THAN `expect`.
Notice that, only string, integer and float types can be compared by AssertGTE,
others are invalid.

### AssertLT

```go
func AssertLT(value any, expect any)
```

AssertLT checks `value` is LESS EQUAL THAN `expect`.
Notice that, only string, integer and float types can be compared by AssertLT,
others are invalid.

### AssertLE

```go
func AssertLE(value any, expect any)
```

AssertLE checks `value` is LESS OR EQUAL THAN `expect`.
Notice that, only string, integer and float types can be compared by AssertLTE,
others are invalid.

### AssertIN

```go
func AssertIN(value any, expect any)
```

AssertIN checks `value` is IN `expect`.
The `expect` should be a slice,
but the `value` can be a slice or a basic type variable.
TODO: gconv.Strings(0) is not [0]

### AssertNI

```go
func AssertNI(value any, expect any)
```

AssertNI checks `value` is NOT IN `expect`.
The `expect` should be a slice,
but the `value` can be a slice or a basic type variable.

### Error

```go
func Error(message ...any)
```

Error panics with given `message`.

### Fatal

```go
func Fatal(message ...any)
```

Fatal prints `message` to stderr and exit the process.

### AssertNil

```go
func AssertNil(value any)
```

AssertNil asserts `value` is nil.

### DataPath

```go
func DataPath(names ...string) string
```

DataPath retrieves and returns the testdata path of current package,
which is used for unit testing cases only.
The optional parameter `names` specifies the sub-folders/sub-files,
which will be joined with current system separator and returned with the path.

### DataContent

```go
func DataContent(names ...string) string
```

DataContent retrieves and returns the file content for specified testdata path of current package

## 内部依赖

- `debug/gdebug`
- `internal/empty`
- `text/gstr`
- `util/gconv`

