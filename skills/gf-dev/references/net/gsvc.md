# gsvc

> Package: `github.com/gogf/gf/v2/net/gsvc`

```go
import "github.com/gogf/gf/v2/net/gsvc"
```

## 概述

Package gsvc provides service registry and discovery definition.

## 源文件

- `gsvc.go`
- `gsvc_discovery.go`
- `gsvc_endpoint.go`
- `gsvc_endpoints.go`
- `gsvc_metadata.go`
- `gsvc_registry.go`
- `gsvc_service.go`

## 类型定义

### Registry

**类型**: interface

Registry interface for service.


### Registrar

**类型**: interface

Registrar interface for service registrar.


### Discovery

**类型**: interface

Discovery interface for service discovery.


### Watcher

**类型**: interface

Watcher interface for service.


### Service

**类型**: interface

Service interface for service definition.


### Endpoint

**类型**: interface

Endpoint interface for service.


### Endpoints

**类型**: type

Endpoints are composed by multiple Endpoint.


**方法**:

- `func (es Endpoints) String() string`
  String formats and returns the Endpoints as a string like:

### Metadata

**类型**: type

Metadata stores custom key-value pairs.


**方法**:

- `func (m Metadata) Set(key string, value any)`
  Set sets key-value pair into metadata.
- `func (m Metadata) Sets(kvs map[string]any)`
  Sets sets key-value pairs into metadata.
- `func (m Metadata) Get(key string) *gvar.Var`
  Get retrieves and returns value of specified key as gvar.
- `func (m Metadata) IsEmpty() bool`
  IsEmpty checks and returns whether current Metadata is empty.

### SearchInput

**类型**: struct

SearchInput is the input for service searching.


### ServiceWatch

**类型**: type

ServiceWatch is used to watch the service status.


### LocalEndpoint

**类型**: struct

LocalEndpoint implements interface Endpoint.


**方法**:

- `func (e *LocalEndpoint) Host() string`
  Host returns the IPv4/IPv6 address of a service.
- `func (e *LocalEndpoint) Port() int`
  Port returns the port of a service.
- `func (e *LocalEndpoint) String() string`
  String formats and returns the Endpoint as a string, like: 192.168.1.100:80.

### LocalService

**类型**: struct

LocalService provides a default implements for interface Service.


**方法**:

- `func (s *LocalService) GetName() string`
  GetName returns the name of the service.
- `func (s *LocalService) GetVersion() string`
  GetVersion returns the version of the service.
- `func (s *LocalService) GetKey() string`
  GetKey formats and returns a unique key string for service.
- `func (s *LocalService) GetValue() string`
  GetValue formats and returns the value of the service.
- `func (s *LocalService) GetPrefix() string`
  GetPrefix formats and returns the key prefix string.
- `func (s *LocalService) GetMetadata() Metadata`
  GetMetadata returns the Metadata map of service.
- `func (s *LocalService) GetEndpoints() Endpoints`
  GetEndpoints returns the Endpoints of service.

## 函数

### SetRegistry

```go
func SetRegistry(registry Registry)
```

SetRegistry sets the default Registry implements as your own implemented interface.

### GetRegistry

```go
func GetRegistry() Registry
```

GetRegistry returns the default Registry that is previously set.
It returns nil if no Registry is set.

### Get

```go
func Get(ctx context.Context, name string) (service Service, err error)
```

Get retrieves and returns the service by service name.

### GetWithDiscovery

```go
func GetWithDiscovery(ctx context.Context, discovery Discovery, name string) (service Service, err error)
```

GetWithDiscovery retrieves and returns the service by service name in `discovery`.

### GetAndWatch

```go
func GetAndWatch(ctx context.Context, name string, watch ServiceWatch) (service Service, err error)
```

GetAndWatch is used to getting the service with custom watch callback function.

### GetAndWatchWithDiscovery

```go
func GetAndWatchWithDiscovery(ctx context.Context, discovery Discovery, name string, watch ServiceWatch) (service Service, err error)
```

GetAndWatchWithDiscovery is used to getting the service with custom watch callback function in `discovery`.

### Search

```go
func Search(ctx context.Context, in SearchInput) ([]Service, error)
```

Search searches and returns services with specified condition.

### Watch

```go
func Watch(ctx context.Context, key string) (Watcher, error)
```

Watch watches specified condition changes.

### NewEndpoint

```go
func NewEndpoint(address string) Endpoint
```

NewEndpoint creates and returns an Endpoint from address string of pattern "host:port",
eg: "192.168.1.100:80".

### NewEndpoints

```go
func NewEndpoints(addresses string) Endpoints
```

NewEndpoints creates and returns Endpoints from multiple addresses like:
"192.168.1.100:80,192.168.1.101:80".

### Register

```go
func Register(ctx context.Context, service Service) (Service, error)
```

Register registers `service` to default registry..

### Deregister

```go
func Deregister(ctx context.Context, service Service) error
```

Deregister removes `service` from default registry.

### NewServiceWithName

```go
func NewServiceWithName(name string) Service
```

NewServiceWithName creates and returns a default implements for interface Service by service name.

### NewServiceWithKV

```go
func NewServiceWithKV(key string, value string) (Service, error)
```

NewServiceWithKV creates and returns a default implements for interface Service by key-value pair string.

## 内部依赖

- `container/gmap`
- `container/gvar`
- `encoding/gjson`
- `errors/gcode`
- `errors/gerror`
- `internal/intlog`
- `os/gcmd`
- `text/gstr`
- `util/gconv`
- `util/gutil`

