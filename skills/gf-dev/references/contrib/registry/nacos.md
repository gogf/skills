# nacos

> Package: `github.com/gogf/gf/contrib/registry/nacos/v2`

```go
import "github.com/gogf/gf/contrib/registry/nacos/v2"
```

## 概述

Package nacos implements service Registry and Discovery using nacos.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/registry/nacos/v2`

## 源文件

- `nacos.go`
- `nacos_discovery.go`
- `nacos_register.go`
- `nacos_service.go`
- `nacos_watcher.go`

## 类型定义

### Registry

**类型**: struct

Registry is nacos registry.


**方法**:

- `func (reg *Registry) SetClusterName(clusterName string) *Registry`
  SetClusterName can set the clusterName. The default is 'DEFAULT'
- `func (reg *Registry) SetGroupName(groupName string) *Registry`
  SetGroupName can set the groupName. The default is 'DEFAULT_GROUP'
- `func (reg *Registry) SetDefaultEndpoint(endpoint string) *Registry`
  SetDefaultEndpoint sets the default endpoint for service registration.
- `func (reg *Registry) SetDefaultMetadata(metadata map[string]string) *Registry`
  SetDefaultMetadata sets the default metadata for service registration.
- `func (reg *Registry) Search(ctx context.Context, in gsvc.SearchInput) (result []gsvc.Service, err error)`
  Search searches and returns services with specified condition.
- `func (reg *Registry) Watch(ctx context.Context, key string) (watcher gsvc.Watcher, err error)`
  Watch watches specified condition changes.
- `func (reg *Registry) Register(ctx context.Context, service gsvc.Service) (registered gsvc.Service, err error)`
  Register registers `service` to Registry.
- `func (reg *Registry) Deregister(ctx context.Context, service gsvc.Service) err error`
  Deregister off-lines and removes `service` from the Registry.

### Config

**类型**: struct

Config is the configuration object for nacos client.


### Watcher

**类型**: struct

Watcher used to mange service event such as update.


**方法**:

- `func (w *Watcher) Proceed() (services []gsvc.Service, err error)`
  Proceed proceeds watch in blocking way.
- `func (w *Watcher) Close() err error`
  Close closes the watcher.
- `func (w *Watcher) SetCloseFunc(close func)`
  SetCloseFunc set the close callback function
- `func (w *Watcher) Push(services []model.Instance, err error)`
  Push add the services watchevent to event queue

## 函数

### New

```go
func New(address string, opts ...constant.ClientOption) reg *Registry
```

New new a registry with address and opts

### NewWithConfig

```go
func NewWithConfig(ctx context.Context, config Config) (reg *Registry, err error)
```

New creates and returns registry with Config.

### NewWithClient

```go
func NewWithClient(client naming_client.INamingClient) *Registry
```

NewWithClient new the instance with INamingClient

### NewServiceFromInstance

```go
func NewServiceFromInstance(instance []model.Instance) gsvc.Service
```

NewServiceFromInstance new one service from instance

### NewServicesFromInstances

```go
func NewServicesFromInstances(instances []model.Instance) []gsvc.Service
```

NewServicesFromInstances new some services from some instances

## 内部依赖

- `container/gmap`
- `errors/gcode`
- `errors/gerror`
- `frame/g`
- `net/gsvc`
- `os/gctx`
- `text/gstr`
- `util/gconv`

## 外部依赖

- `github.com/nacos-group/nacos-sdk-go/v2/clients`
- `github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client`
- `github.com/nacos-group/nacos-sdk-go/v2/common/constant`
- `github.com/nacos-group/nacos-sdk-go/v2/model`
- `github.com/nacos-group/nacos-sdk-go/v2/vo`

