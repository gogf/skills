# gmode

> Package: `github.com/gogf/gf/v2/util/gmode`

```go
import "github.com/gogf/gf/v2/util/gmode"
```

## 概述

Package gmode provides release mode management for project.

## 源文件

- `gmode.go`

## 函数

### Set

```go
func Set(mode string)
```

Set sets the mode for current application.

### SetDevelop

```go
func SetDevelop()
```

SetDevelop sets current mode DEVELOP for current application.

### SetTesting

```go
func SetTesting()
```

SetTesting sets current mode TESTING for current application.

### SetStaging

```go
func SetStaging()
```

SetStaging sets current mode STAGING for current application.

### SetProduct

```go
func SetProduct()
```

SetProduct sets current mode PRODUCT for current application.

### Mode

```go
func Mode() string
```

Mode returns current application mode set.

### IsDevelop

```go
func IsDevelop() bool
```

IsDevelop checks and returns whether current application is running in DEVELOP mode.

### IsTesting

```go
func IsTesting() bool
```

IsTesting checks and returns whether current application is running in TESTING mode.

### IsStaging

```go
func IsStaging() bool
```

IsStaging checks and returns whether current application is running in STAGING mode.

### IsProduct

```go
func IsProduct() bool
```

IsProduct checks and returns whether current application is running in PRODUCT mode.

## 内部依赖

- `debug/gdebug`
- `internal/command`
- `os/gfile`

