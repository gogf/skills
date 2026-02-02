# apollo

> Package: `github.com/gogf/gf/contrib/config/apollo/v2`

```go
import "github.com/gogf/gf/contrib/config/apollo/v2"
```

## 概述

Package apollo implements gcfg.Adapter using apollo service.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/config/apollo/v2`

## 源文件

- `apollo.go`
- `apollo_adapter_ctx.go`

## 类型定义

### Config

**类型**: struct

Config is the configuration object for apollo client.


### Client

**类型**: struct

Client implements gcfg.Adapter implementing using apollo service.


**方法**:

- `func (c *Client) Available(ctx context.Context, resource ...string) ok bool`
  Available checks and returns the backend configuration service is available.
- `func (c *Client) Get(ctx context.Context, pattern string) (value any, err error)`
  Get retrieves and returns value by specified `pattern` in current resource.
- `func (c *Client) Data(ctx context.Context) (data map[string]any, err error)`
  Data retrieves and returns all configuration data in current resource as map.
- `func (c *Client) OnChange(event *storage.ChangeEvent)`
  OnChange is called when config changes.
- `func (c *Client) OnNewestChange(event *storage.FullChangeEvent)`
  OnNewestChange is called when any config changes.
- `func (c *Client) AddWatcher(name string, f func)`
  AddWatcher adds a watcher for the specified configuration file.
- `func (c *Client) RemoveWatcher(name string)`
  RemoveWatcher removes the watcher for the specified configuration file.
- `func (c *Client) GetWatcherNames() []string`
  GetWatcherNames returns all watcher names.

### ApolloAdapterCtx

**类型**: struct

ApolloAdapterCtx is the context adapter for Apollo configuration


**方法**:

- `func (a *ApolloAdapterCtx) WithOperation(operation gcfg.OperationType) *ApolloAdapterCtx`
  WithOperation sets the operation in the context
- `func (a *ApolloAdapterCtx) WithNamespace(namespace string) *ApolloAdapterCtx`
  WithNamespace sets the namespace in the context
- `func (a *ApolloAdapterCtx) WithAppId(appId string) *ApolloAdapterCtx`
  WithAppId sets the appId in the context
- `func (a *ApolloAdapterCtx) WithCluster(cluster string) *ApolloAdapterCtx`
  WithCluster sets the cluster in the context
- `func (a *ApolloAdapterCtx) WithContent(content *gjson.Json) *ApolloAdapterCtx`
  WithContent sets the content in the context
- `func (a *ApolloAdapterCtx) GetNamespace() string`
  GetNamespace retrieves the namespace from the context
- `func (a *ApolloAdapterCtx) GetAppId() string`
  GetAppId retrieves the appId from the context
- `func (a *ApolloAdapterCtx) GetCluster() string`
  GetCluster retrieves the cluster from the context
- `func (a *ApolloAdapterCtx) GetContent() *gjson.Json`
  GetContent retrieves the content from the context
- `func (a *ApolloAdapterCtx) GetOperation() gcfg.OperationType`
  GetOperation retrieves the operation from the context

## 函数

### New

```go
func New(ctx context.Context, config Config) (adapter gcfg.Adapter, err error)
```

New creates and returns gcfg.Adapter implementing using apollo service.

### NewAdapterCtxWithCtx

```go
func NewAdapterCtxWithCtx(ctx context.Context) *ApolloAdapterCtx
```

NewAdapterCtxWithCtx creates and returns a new ApolloAdapterCtx with the given context.

### NewAdapterCtx

```go
func NewAdapterCtx(ctx ...context.Context) *ApolloAdapterCtx
```

NewAdapterCtx creates and returns a new ApolloAdapterCtx.
If context is provided, it will be used; otherwise, a background context is created.

### GetAdapterCtx

```go
func GetAdapterCtx(ctx context.Context) *ApolloAdapterCtx
```

GetAdapterCtx creates a new ApolloAdapterCtx with the given context

## 内部依赖

- `encoding/gjson`
- `errors/gerror`
- `frame/g`
- `os/gcfg`
- `os/gctx`
- `text/gstr`
- `util/gconv`

## 外部依赖

- `github.com/apolloconfig/agollo/v4`
- `github.com/apolloconfig/agollo/v4/env/config`
- `github.com/apolloconfig/agollo/v4/storage`

