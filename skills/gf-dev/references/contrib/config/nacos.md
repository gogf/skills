# nacos

> Package: `github.com/gogf/gf/contrib/config/nacos/v2`

```go
import "github.com/gogf/gf/contrib/config/nacos/v2"
```

## 概述

Package nacos implements gcfg.Adapter using nacos service.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/config/nacos/v2`

## 源文件

- `nacos.go`
- `nacos_adapter_ctx.go`

## 类型定义

### Config

**类型**: struct

Config is the configuration object for nacos client.


### Client

**类型**: struct

Client implements gcfg.Adapter implementing using nacos service.


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

### NacosAdapterCtx

**类型**: struct

NacosAdapterCtx is the context adapter for Nacos configuration


**方法**:

- `func (n *NacosAdapterCtx) WithOperation(operation gcfg.OperationType) *NacosAdapterCtx`
  WithOperation sets the operation in the context
- `func (n *NacosAdapterCtx) WithNamespace(namespace string) *NacosAdapterCtx`
  WithNamespace sets the namespace in the context
- `func (n *NacosAdapterCtx) WithGroup(group string) *NacosAdapterCtx`
  WithGroup sets the group in the context
- `func (n *NacosAdapterCtx) WithDataId(dataId string) *NacosAdapterCtx`
  WithDataId sets the dataId in the context
- `func (n *NacosAdapterCtx) WithContent(content string) *NacosAdapterCtx`
  WithContent sets the content in the context
- `func (n *NacosAdapterCtx) GetNamespace() string`
  GetNamespace retrieves the namespace from the context
- `func (n *NacosAdapterCtx) GetGroup() string`
  GetGroup retrieves the group from the context
- `func (n *NacosAdapterCtx) GetDataId() string`
  GetDataId retrieves the dataId from the context
- `func (n *NacosAdapterCtx) GetContent() string`
  GetContent retrieves the content from the context
- `func (n *NacosAdapterCtx) GetOperation() gcfg.OperationType`
  GetOperation retrieves the operation from the context

## 函数

### New

```go
func New(ctx context.Context, config Config) (adapter gcfg.Adapter, err error)
```

New creates and returns gcfg.Adapter implementing using nacos service.

### NewAdapterCtxWithCtx

```go
func NewAdapterCtxWithCtx(ctx context.Context) *NacosAdapterCtx
```

NewAdapterCtxWithCtx creates and returns a new NacosAdapterCtx with the given context.

### NewAdapterCtx

```go
func NewAdapterCtx(ctx ...context.Context) *NacosAdapterCtx
```

NewAdapterCtx creates and returns a new NacosAdapterCtx.
If context is provided, it will be used; otherwise, a background context is created.

### GetAdapterCtx

```go
func GetAdapterCtx(ctx context.Context) *NacosAdapterCtx
```

GetAdapterCtx creates a new NacosAdapterCtx with the given context

## 内部依赖

- `encoding/gjson`
- `errors/gerror`
- `frame/g`
- `os/gcfg`
- `os/gctx`

## 外部依赖

- `github.com/nacos-group/nacos-sdk-go/v2/clients`
- `github.com/nacos-group/nacos-sdk-go/v2/clients/config_client`
- `github.com/nacos-group/nacos-sdk-go/v2/common/constant`
- `github.com/nacos-group/nacos-sdk-go/v2/vo`

