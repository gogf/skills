# consul

> Package: `github.com/gogf/gf/contrib/config/consul/v2`

```go
import "github.com/gogf/gf/contrib/config/consul/v2"
```

## 概述

Package consul implements gcfg.Adapter using consul service.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/config/consul/v2`

## 源文件

- `consul.go`
- `consul_adapter_ctx.go`

## 类型定义

### Config

**类型**: struct

Config is the configuration object for consul client.


### Client

**类型**: struct

Client implements gcfg.Adapter implementing using consul service.


**方法**:

- `func (c *Client) Available(ctx context.Context, resource ...string) ok bool`
  Available checks and returns the backend configuration service is available.
- `func (c *Client) Get(ctx context.Context, pattern string) (value any, err error)`
  Get retrieves and returns value by specified `pattern` in current resource.
- `func (c *Client) Data(ctx context.Context) (data map[string]any, err error)`
  Data retrieves and returns all configuration data in current resource as map.
- `func (c *Client) AddWatcher(name string, f func)`
  AddWatcher adds a watcher for the specified configuration file.
- `func (c *Client) RemoveWatcher(name string)`
  RemoveWatcher removes the watcher for the specified configuration file.
- `func (c *Client) GetWatcherNames() []string`
  GetWatcherNames returns all watcher names.

### ConsulAdapterCtx

**类型**: struct

ConsulAdapterCtx is the context adapter for Consul configuration


**方法**:

- `func (a *ConsulAdapterCtx) WithOperation(operation gcfg.OperationType) *ConsulAdapterCtx`
  WithOperation sets the operation in the context
- `func (a *ConsulAdapterCtx) WithPath(path string) *ConsulAdapterCtx`
  WithPath sets the path in the context
- `func (a *ConsulAdapterCtx) WithContent(content *gjson.Json) *ConsulAdapterCtx`
  WithContent sets the content in the context
- `func (a *ConsulAdapterCtx) GetContent() *gjson.Json`
  GetContent retrieves the content from the context
- `func (a *ConsulAdapterCtx) GetOperation() gcfg.OperationType`
  GetOperation retrieves the operation from the context
- `func (a *ConsulAdapterCtx) GetPath() string`
  GetPath retrieves the path from the context

## 函数

### New

```go
func New(ctx context.Context, config Config) (adapter gcfg.Adapter, err error)
```

New creates and returns gcfg.Adapter implementing using consul service.

### NewAdapterCtxWithCtx

```go
func NewAdapterCtxWithCtx(ctx context.Context) *ConsulAdapterCtx
```

NewAdapterCtxWithCtx creates and returns a new ConsulAdapterCtx with the given context.

### NewAdapterCtx

```go
func NewAdapterCtx(ctx ...context.Context) *ConsulAdapterCtx
```

NewAdapterCtx creates and returns a new ConsulAdapterCtx.
If context is provided, it will be used; otherwise, a background context is created.

### GetAdapterCtx

```go
func GetAdapterCtx(ctx context.Context) *ConsulAdapterCtx
```

GetAdapterCtx creates a new ConsulAdapterCtx with the given context

## 内部依赖

- `encoding/gjson`
- `errors/gerror`
- `frame/g`
- `os/gcfg`
- `os/gctx`
- `os/glog`

## 外部依赖

- `github.com/hashicorp/consul/api`
- `github.com/hashicorp/consul/api/watch`

