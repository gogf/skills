# file

> Package: `github.com/gogf/gf/contrib/registry/file/v2`

```go
import "github.com/gogf/gf/contrib/registry/file/v2"
```

## 概述

Package file implements service Registry and Discovery using file.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/registry/file/v2`

## 源文件

- `file.go`
- `file_discovery.go`
- `file_registrar.go`
- `file_service.go`
- `file_watcher.go`

## 类型定义

### Registry

**类型**: struct

Registry implements interface Registry using file.
This implement is usually for testing only.


**方法**:

- `func (r *Registry) Search(ctx context.Context, in gsvc.SearchInput) (result []gsvc.Service, err error)`
  Search searches and returns services with specified condition.
- `func (r *Registry) Watch(ctx context.Context, key string) (watcher gsvc.Watcher, err error)`
  Watch watches specified condition changes.
- `func (r *Registry) Register(ctx context.Context, service gsvc.Service) (registered gsvc.Service, err error)`
  Register registers `service` to Registry.
- `func (r *Registry) Deregister(ctx context.Context, service gsvc.Service) error`
  Deregister off-lines and removes `service` from the Registry.

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

### Watcher

**类型**: struct

Watcher for file changes watch.


**方法**:

- `func (w *Watcher) Proceed() (services []gsvc.Service, err error)`
  Proceed proceeds watch in blocking way.
- `func (w *Watcher) Close() error`
  Close closes the watcher.

## 函数

### New

```go
func New(path string) gsvc.Registry
```

New creates and returns a gsvc.Registry implements using file.

### NewService

```go
func NewService(service gsvc.Service) *Service
```

NewService creates and returns local Service from gsvc.Service interface object.

## 内部依赖

- `container/gmap`
- `container/gtype`
- `encoding/gjson`
- `errors/gerror`
- `frame/g`
- `net/gsvc`
- `os/gfile`
- `os/gfsnotify`
- `os/gtime`
- `os/gtimer`
- `text/gstr`

