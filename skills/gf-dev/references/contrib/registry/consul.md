# consul

> Package: `github.com/gogf/gf/contrib/registry/consul/v2`

```go
import "github.com/gogf/gf/contrib/registry/consul/v2"
```

## 概述

Package consul implements service Registry and Discovery using consul.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/registry/consul/v2`

## 源文件

- `consul.go`
- `consul_discovery.go`
- `consul_watcher.go`

## 类型定义

### Registry

**类型**: struct

Registry implements gsvc.Registry interface using consul.


**方法**:

- `func (r *Registry) Register(ctx context.Context, service gsvc.Service) (gsvc.Service, error)`
  Register registers a service to consul.
- `func (r *Registry) Deregister(ctx context.Context, service gsvc.Service) error`
  Deregister deregisters a service from consul.
- `func (r *Registry) GetAddress() string`
  GetAddress returns the consul address
- `func (r *Registry) Watch(ctx context.Context, key string) (gsvc.Watcher, error)`
  Watch creates and returns a watcher for specified service.
- `func (r *Registry) Search(ctx context.Context, in gsvc.SearchInput) ([]gsvc.Service, error)`
  Search searches and returns services with specified condition.

### Option

**类型**: type

Option is the configuration option type for registry.


### Watcher

**类型**: struct

Watcher watches the service changes.


**方法**:

- `func (w *Watcher) Proceed() ([]gsvc.Service, error)`
  Proceed returns current services and waits for the next service change.
- `func (w *Watcher) Close() error`
  Close closes the watcher.
- `func (w *Watcher) Services() ([]gsvc.Service, error)`
  Services returns current services from the watcher.

## 函数

### WithAddress

```go
func WithAddress(address string) Option
```

WithAddress sets the address for consul client.

### WithToken

```go
func WithToken(token string) Option
```

WithToken sets the ACL token for consul client.

### New

```go
func New(opts ...Option) (gsvc.Registry, error)
```

New creates and returns a new Registry.

## 内部依赖

- `errors/gerror`
- `net/gsvc`

## 外部依赖

- `github.com/hashicorp/consul/api`
- `github.com/hashicorp/consul/api/watch`

