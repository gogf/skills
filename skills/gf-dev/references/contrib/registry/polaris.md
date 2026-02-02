# polaris

> Package: `github.com/gogf/gf/contrib/registry/polaris/v2`

```go
import "github.com/gogf/gf/contrib/registry/polaris/v2"
```

## 概述

Package polaris implements service Registry and Discovery using polaris.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/registry/polaris/v2`

## 源文件

- `polaris.go`
- `polaris_discovery.go`
- `polaris_registry.go`
- `polaris_service.go`
- `polaris_watcher.go`

## 类型定义

### Option

**类型**: type

Option The option is a polaris option.


### Registry

**类型**: struct

Registry is polaris registry.


**方法**:

- `func (r *Registry) Search(ctx context.Context, in gsvc.SearchInput) ([]gsvc.Service, error)`
  Search returns the service instances in memory according to the service name.
- `func (r *Registry) Watch(ctx context.Context, key string) (gsvc.Watcher, error)`
  Watch creates a watcher according to the service name.
- `func (r *Registry) Register(ctx context.Context, service gsvc.Service) (gsvc.Service, error)`
  Register the registration.
- `func (r *Registry) Deregister(ctx context.Context, service gsvc.Service) error`
  Deregister the registration.

### Service

**类型**: struct

Service for wrapping gsvc.Server and extends extra attributes for polaris purpose.


**方法**:

- `func (s *Service) GetKey() string`
  GetKey overwrites the GetKey function of gsvc.Service for replacing separator...
- `func (s *Service) GetPrefix() string`
  GetPrefix overwrites the GetPrefix function of gsvc.Service for replacing sep...

### Watcher

**类型**: struct

Watcher is a service watcher.


**方法**:

- `func (w *Watcher) Proceed() ([]gsvc.Service, error)`
  Proceed returns services in the following two cases:
- `func (w *Watcher) Close() error`
  Close the watcher.

## 函数

### WithNamespace

```go
func WithNamespace(namespace string) Option
```

WithNamespace with the Namespace option.

### WithServiceToken

```go
func WithServiceToken(serviceToken string) Option
```

WithServiceToken with ServiceToken option.

### WithProtocol

```go
func WithProtocol(protocol string) Option
```

WithProtocol with the Protocol option.

### WithWeight

```go
func WithWeight(weight int) Option
```

WithWeight with the Weight option.

### WithHealthy

```go
func WithHealthy(healthy bool) Option
```

WithHealthy with the Healthy option.

### WithIsolate

```go
func WithIsolate(isolate bool) Option
```

WithIsolate with the Isolate option.

### WithTTL

```go
func WithTTL(TTL int) Option
```

WithTTL with the TTL option.

### WithTimeout

```go
func WithTimeout(timeout time.Duration) Option
```

WithTimeout the Timeout option.

### WithRetryCount

```go
func WithRetryCount(retryCount int) Option
```

WithRetryCount with RetryCount option.

### WithLogger

```go
func WithLogger(logger glog.ILogger) Option
```

WithLogger with the Logger option.

### New

```go
func New(provider polaris.ProviderAPI, consumer polaris.ConsumerAPI, opts ...Option) gsvc.Registry
```

New create a new registry.

### NewWithConfig

```go
func NewWithConfig(conf config.Configuration, opts ...Option) gsvc.Registry
```

NewWithConfig new a registry with config.

## 内部依赖

- `container/gmap`
- `frame/g`
- `net/gsvc`
- `os/glog`
- `text/gstr`
- `util/gconv`

## 外部依赖

- `github.com/polarismesh/polaris-go`
- `github.com/polarismesh/polaris-go/pkg/config`
- `github.com/polarismesh/polaris-go/pkg/model`

