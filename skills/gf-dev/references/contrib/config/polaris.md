# polaris

> Package: `github.com/gogf/gf/contrib/config/polaris/v2`

```go
import "github.com/gogf/gf/contrib/config/polaris/v2"
```

## 概述

Package polaris implements gcfg.Adapter using polaris service.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/config/polaris/v2`

## 源文件

- `polaris.go`
- `polaris_adapter_ctx.go`

## 类型定义

### Config

**类型**: struct

Config is the configuration for polaris.


### Client

**类型**: struct

Client implements gcfg.Adapter implementing using polaris service.


**方法**:

- `func (c *Client) LogDir(dir string) error`
  LogDir sets the log directory for polaris.
- `func (c *Client) Available(ctx context.Context, resource ...string) ok bool`
  Available checks and returns the backend configuration service is available.
- `func (c *Client) Get(ctx context.Context, pattern string) (value any, err error)`
  Get retrieves and return value by specified `pattern` in current resource.
- `func (c *Client) Data(ctx context.Context) (data map[string]any, err error)`
  Data retrieves and returns all configuration data in current resource as map.
- `func (c *Client) AddWatcher(name string, f func)`
  AddWatcher adds a watcher for the specified configuration file.
- `func (c *Client) RemoveWatcher(name string)`
  RemoveWatcher removes the watcher for the specified configuration file.
- `func (c *Client) GetWatcherNames() []string`
  GetWatcherNames returns all watcher names.

### PolarisAdapterCtx

**类型**: struct

PolarisAdapterCtx is the context adapter for polaris configuration


**方法**:

- `func (n *PolarisAdapterCtx) WithOperation(operation gcfg.OperationType) *PolarisAdapterCtx`
  WithOperation sets the operation in the context
- `func (n *PolarisAdapterCtx) WithNamespace(namespace string) *PolarisAdapterCtx`
  WithNamespace sets the namespace in the context
- `func (n *PolarisAdapterCtx) WithFileGroup(fileGroup string) *PolarisAdapterCtx`
  WithFileGroup sets the group in the context
- `func (n *PolarisAdapterCtx) WithFileName(fileName string) *PolarisAdapterCtx`
  WithFileName sets the fileName in the context
- `func (n *PolarisAdapterCtx) WithContent(content string) *PolarisAdapterCtx`
  WithContent sets the content in the context
- `func (n *PolarisAdapterCtx) GetNamespace() string`
  GetNamespace retrieves the namespace from the context
- `func (n *PolarisAdapterCtx) GetFileGroup() string`
  GetFileGroup retrieves the group from the context
- `func (n *PolarisAdapterCtx) GetFileName() string`
  GetFileName retrieves the fileName from the context
- `func (n *PolarisAdapterCtx) GetContent() string`
  GetContent retrieves the content from the context
- `func (n *PolarisAdapterCtx) GetOperation() gcfg.OperationType`
  GetOperation retrieves the operation from the context

## 函数

### New

```go
func New(ctx context.Context, config Config) (adapter gcfg.Adapter, err error)
```

New creates and returns gcfg.Adapter implementing using polaris service.

### NewAdapterCtxWithCtx

```go
func NewAdapterCtxWithCtx(ctx context.Context) *PolarisAdapterCtx
```

NewAdapterCtxWithCtx creates and returns a new PolarisAdapterCtx with the given context.

### NewAdapterCtx

```go
func NewAdapterCtx(ctx ...context.Context) *PolarisAdapterCtx
```

NewAdapterCtx creates and returns a new PolarisAdapterCtx.
If context is provided, it will be used; otherwise, a background context is created.

### GetAdapterCtx

```go
func GetAdapterCtx(ctx context.Context) *PolarisAdapterCtx
```

GetAdapterCtx creates a new PolarisAdapterCtx with the given context

## 内部依赖

- `encoding/gjson`
- `errors/gerror`
- `frame/g`
- `os/gcfg`
- `os/gctx`
- `text/gstr`

## 外部依赖

- `github.com/polarismesh/polaris-go`
- `github.com/polarismesh/polaris-go/api`
- `github.com/polarismesh/polaris-go/pkg/model`

