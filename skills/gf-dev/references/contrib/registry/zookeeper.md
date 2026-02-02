# zookeeper

> Package: `github.com/gogf/gf/contrib/registry/zookeeper/v2`

```go
import "github.com/gogf/gf/contrib/registry/zookeeper/v2"
```

## 概述

Package zookeeper implements service Registry and Discovery using zookeeper.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/registry/zookeeper/v2`

## 源文件

- `zookeeper.go`
- `zookeeper_discovery.go`
- `zookeeper_func.go`
- `zookeeper_registrar.go`
- `zookeeper_watcher.go`

## 类型定义

### Content

**类型**: struct

Content for custom service Marshal/Unmarshal.


### Option

**类型**: type

Option is etcd registry option.


### Registry

**类型**: struct

Registry is consul registry


**方法**:

- `func (r *Registry) Search(_ context.Context, in gsvc.SearchInput) ([]gsvc.Service, error)`
  Search searches and returns services with specified condition.
- `func (r *Registry) Watch(ctx context.Context, key string) (gsvc.Watcher, error)`
  Watch watches specified condition changes.
- `func (r *Registry) Register(_ context.Context, service gsvc.Service) (gsvc.Service, error)`
  Register registers `service` to Registry.
- `func (r *Registry) Deregister(ctx context.Context, service gsvc.Service) error`
  Deregister off-lines and removes `service` from the Registry.

## 函数

### WithRootPath

```go
func WithRootPath(path string) Option
```

WithRootPath with registry root path.

### WithDigestACL

```go
func WithDigestACL(user string, password string) Option
```

WithDigestACL with registry password.

### New

```go
func New(address []string, opts ...Option) *Registry
```

## 内部依赖

- `container/gmap`
- `errors/gerror`
- `net/gsvc`
- `text/gstr`

## 外部依赖

- `github.com/go-zookeeper/zk`
- `golang.org/x/sync/singleflight`

