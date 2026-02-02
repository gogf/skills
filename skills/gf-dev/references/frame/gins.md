# gins

> Package: `github.com/gogf/gf/v2/frame/gins`

```go
import "github.com/gogf/gf/v2/frame/gins"
```

## 概述

Package gins provides instances and core components management.

## 源文件

- `gins.go`
- `gins_config.go`
- `gins_database.go`
- `gins_httpclient.go`
- `gins_i18n.go`
- `gins_log.go`
- `gins_redis.go`
- `gins_resource.go`
- `gins_server.go`
- `gins_view.go`

## 函数

### Config

```go
func Config(name ...string) *gcfg.Config
```

Config returns an instance of View with default settings.
The parameter `name` is the name for the instance.

### Database

```go
func Database(name ...string) gdb.DB
```

Database returns an instance of database ORM object with specified configuration group name.
Note that it panics if any error occurs duration instance creating.

### HttpClient

```go
func HttpClient(name ...any) *gclient.Client
```

HttpClient returns an instance of http client with specified name.

### I18n

```go
func I18n(name ...string) *gi18n.Manager
```

I18n returns an instance of gi18n.Manager.
The parameter `name` is the name for the instance.

### Log

```go
func Log(name ...string) *glog.Logger
```

Log returns an instance of glog.Logger.
The parameter `name` is the name for the instance.
Note that it panics if any error occurs duration instance creating.

### Redis

```go
func Redis(name ...string) *gredis.Redis
```

Redis returns an instance of redis client with specified configuration group name.
Note that it panics if any error occurs duration instance creating.

### Resource

```go
func Resource(name ...string) *gres.Resource
```

Resource returns an instance of Resource.
The parameter `name` is the name for the instance.

### Server

```go
func Server(name ...any) *ghttp.Server
```

Server returns an instance of http server with specified name.
Note that it panics if any error occurs duration instance creating.

### View

```go
func View(name ...string) *gview.View
```

View returns an instance of View with default settings.
The parameter `name` is the name for the instance.
Note that it panics if any error occurs duration instance creating.

## 内部依赖

- `database/gdb`
- `database/gredis`
- `errors/gcode`
- `errors/gerror`
- `i18n/gi18n`
- `internal/consts`
- `internal/instance`
- `internal/intlog`
- `net/gclient`
- `net/ghttp`
- `os/gcfg`
- `os/glog`
- `os/gres`
- `os/gview`
- `util/gconv`
- `util/gutil`

