# etcd

> Package: `github.com/gogf/gf/contrib/registry/etcd/v2`

```go
import "github.com/gogf/gf/contrib/registry/etcd/v2"
```

## 概述

Package etcd implements service Registry and Discovery using etcd.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/registry/etcd/v2`

## 源文件

- `etcd.go`
- `etcd_discovery.go`
- `etcd_registrar.go`
- `etcd_service.go`
- `etcd_watcher.go`

## 类型定义

### Registry

**类型**: struct

Registry implements gsvc.Registry interface.


**方法**:

- `func (r *Registry) Search(ctx context.Context, in gsvc.SearchInput) ([]gsvc.Service, error)`
  Search searches and returns services with specified condition.
- `func (r *Registry) Watch(ctx context.Context, key string) (gsvc.Watcher, error)`
  Watch watches specified condition changes.
- `func (r *Registry) Register(ctx context.Context, service gsvc.Service) (gsvc.Service, error)`
  Register registers `service` to Registry.
- `func (r *Registry) Deregister(ctx context.Context, service gsvc.Service) error`
  Deregister off-lines and removes `service` from the Registry.

### Option

**类型**: struct

Option is the option for the etcd registry.


### Service

**类型**: struct

Service wrapper.


**方法**:

- `func (s *Service) GetMetadata() gsvc.Metadata`
  GetMetadata returns the Metadata map of service.
- `func (s *Service) GetEndpoints() gsvc.Endpoints`
  GetEndpoints returns the Endpoints of service.
- `func (s *Service) GetValue() string`
  GetValue formats and returns the value of the service.

## 函数

### New

```go
func New(address string, option ...Option) *Registry
```

New creates and returns a new etcd registry.
Support Etcd Address format: ip:port,ip:port...,ip:port@username:password

### NewWithClient

```go
func NewWithClient(client *etcd3.Client, option ...Option) *Registry
```

NewWithClient creates and returns a new etcd registry with the given client.

### NewService

```go
func NewService(service gsvc.Service) *Service
```

NewService creates and returns local Service from gsvc.Service interface object.

## 内部依赖

- `container/gmap`
- `encoding/gjson`
- `errors/gcode`
- `errors/gerror`
- `frame/g`
- `net/gsvc`
- `os/glog`
- `text/gstr`
- `util/grand`

## 外部依赖

- `go.etcd.io/etcd/client/v3`
- `google.golang.org/grpc`

