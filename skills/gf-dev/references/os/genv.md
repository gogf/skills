# genv

> Package: `github.com/gogf/gf/v2/os/genv`

```go
import "github.com/gogf/gf/v2/os/genv"
```

## 概述

Package genv provides operations for environment variables of system.

## 源文件

- `genv.go`
- `genv_must.go`

## 函数

### All

```go
func All() []string
```

All returns a copy of strings representing the environment,
in the form "key=value".

### Map

```go
func Map() map[string]string
```

Map returns a copy of strings representing the environment as a map.

### Get

```go
func Get(key string, def ...any) *gvar.Var
```

Get creates and returns a Var with the value of the environment variable
named by the `key`. It uses the given `def` if the variable does not exist
in the environment.

### Set

```go
func Set(key string, value string) err error
```

Set sets the value of the environment variable named by the `key`.
It returns an error, if any.

### SetMap

```go
func SetMap(m map[string]string) err error
```

SetMap sets the environment variables using map.

### Contains

```go
func Contains(key string) bool
```

Contains checks whether the environment variable named `key` exists.

### Remove

```go
func Remove(key ...string) err error
```

Remove deletes one or more environment variables.

### GetWithCmd

```go
func GetWithCmd(key string, def ...any) *gvar.Var
```

GetWithCmd returns the environment value specified `key`.
If the environment value does not exist, then it retrieves and returns the value from command line options.
It returns the default value `def` if none of them exists.

Fetching Rules:
1. Environment arguments are in uppercase format, eg: GF_<package name>_<variable name>；
2. Command line arguments are in lowercase format, eg: gf.<package name>.<variable name>;

### Build

```go
func Build(m map[string]string) []string
```

Build builds a map to an environment variable slice.

### MapFromEnv

```go
func MapFromEnv(envs []string) map[string]string
```

MapFromEnv converts environment variables from slice to map.

### MapToEnv

```go
func MapToEnv(m map[string]string) []string
```

MapToEnv converts environment variables from map to slice.

### Filter

```go
func Filter(envs []string) []string
```

Filter filters repeated items from given environment variables.

### MustSet

```go
func MustSet(key string, value string)
```

MustSet performs as Set, but it panics if any error occurs.

### MustRemove

```go
func MustRemove(key ...string)
```

MustRemove performs as Remove, but it panics if any error occurs.

## 内部依赖

- `container/gvar`
- `errors/gerror`
- `internal/command`
- `internal/utils`

