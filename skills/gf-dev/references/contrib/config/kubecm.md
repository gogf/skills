# kubecm

> Package: `github.com/gogf/gf/contrib/config/kubecm/v2`

```go
import "github.com/gogf/gf/contrib/config/kubecm/v2"
```

## 概述

Package kubecm implements gcfg.Adapter using kubernetes configmap.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/config/kubecm/v2`

## 源文件

- `kubecm.go`
- `kubecm_adapter_ctx.go`
- `kubecm_kube.go`

## 类型定义

### Client

**类型**: struct

Client implements gcfg.Adapter.


**方法**:

- `func (c *Client) Available(ctx context.Context, configMap ...string) ok bool`
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

### Config

**类型**: struct

Config for Client.


### KubecmAdapterCtx

**类型**: struct

KubecmAdapterCtx is the context adapter for kubecm configuration


**方法**:

- `func (a *KubecmAdapterCtx) WithOperation(operation gcfg.OperationType) *KubecmAdapterCtx`
  WithOperation sets the operation in the context
- `func (a *KubecmAdapterCtx) WithNamespace(namespace string) *KubecmAdapterCtx`
  WithNamespace sets the namespace in the context
- `func (a *KubecmAdapterCtx) WithConfigMap(configMap string) *KubecmAdapterCtx`
  WithConfigMap sets the configmap in the context
- `func (a *KubecmAdapterCtx) WithDataItem(dataItem string) *KubecmAdapterCtx`
  WithDataItem sets the dataitem in the context
- `func (a *KubecmAdapterCtx) WithContent(content *gjson.Json) *KubecmAdapterCtx`
  WithContent sets the content in the context
- `func (a *KubecmAdapterCtx) GetOperation() gcfg.OperationType`
  GetOperation retrieves the operation from the context
- `func (a *KubecmAdapterCtx) GetNamespace() string`
  GetNamespace retrieves the namespace from the context
- `func (a *KubecmAdapterCtx) GetConfigMap() string`
  GetConfigMap retrieves the configmap from the context
- `func (a *KubecmAdapterCtx) GetDataItem() string`
  GetDataItem retrieves the dataitem from the context
- `func (a *KubecmAdapterCtx) GetContent() *gjson.Json`
  GetContent retrieves the content from the context

## 函数

### New

```go
func New(ctx context.Context, config Config) (adapter gcfg.Adapter, err error)
```

New creates and returns gcfg.Adapter implementing using kubernetes configmap.

### NewKubecmAdapterCtx

```go
func NewKubecmAdapterCtx(ctx context.Context) *KubecmAdapterCtx
```

NewKubecmAdapterCtx creates and returns a new KubecmAdapterCtx with the given context.

### NewAdapterCtx

```go
func NewAdapterCtx(ctx ...context.Context) *KubecmAdapterCtx
```

NewAdapterCtx creates and returns a new KubecmAdapterCtx.
If context is provided, it will be used; otherwise, a background context is created.

### GetAdapterCtx

```go
func GetAdapterCtx(ctx context.Context) *KubecmAdapterCtx
```

GetAdapterCtx creates a new KubecmAdapterCtx with the given context

### Namespace

```go
func Namespace() string
```

Namespace retrieves and returns the namespace of current pod.
Note that this function should be called in kubernetes pod.

### NewDefaultKubeClient

```go
func NewDefaultKubeClient(ctx context.Context) (*kubernetes.Clientset, error)
```

NewDefaultKubeClient creates and returns a default kubernetes client.
It is commonly used when the service is running inside kubernetes cluster.

### NewKubeClientFromPath

```go
func NewKubeClientFromPath(ctx context.Context, kubeConfigFilePath string) (*kubernetes.Clientset, error)
```

NewKubeClientFromPath creates and returns a kubernetes REST client by given `kubeConfigFilePath`.

### NewKubeClientFromConfig

```go
func NewKubeClientFromConfig(ctx context.Context, config *rest.Config) (*kubernetes.Clientset, error)
```

NewKubeClientFromConfig creates and returns client by given `rest.Config`.

### NewDefaultKubeConfig

```go
func NewDefaultKubeConfig(ctx context.Context) (*rest.Config, error)
```

NewDefaultKubeConfig creates and returns a default kubernetes config.
It is commonly used when the service is running inside kubernetes cluster.

### NewKubeConfigFromPath

```go
func NewKubeConfigFromPath(ctx context.Context, kubeConfigFilePath string) (*rest.Config, error)
```

NewKubeConfigFromPath creates and returns rest.Config object from given `kubeConfigFilePath`.

## 内部依赖

- `encoding/gjson`
- `errors/gerror`
- `frame/g`
- `os/gcfg`
- `os/gctx`
- `os/gfile`
- `util/gutil`

## 外部依赖

- `k8s.io/apimachinery/pkg/apis/meta/v1`
- `k8s.io/apimachinery/pkg/watch`
- `k8s.io/client-go/kubernetes`
- `k8s.io/client-go/rest`
- `k8s.io/client-go/tools/clientcmd`

