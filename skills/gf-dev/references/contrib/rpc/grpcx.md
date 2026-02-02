# grpcx

> Package: `github.com/gogf/gf/contrib/rpc/grpcx/v2`

```go
import "github.com/gogf/gf/contrib/rpc/grpcx/v2"
```

## 概述

Package grpcx provides grpc service functionalities.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/rpc/grpcx/v2`

## 源文件

- `grpcx.go`
- `grpcx_grpc_client.go`
- `grpcx_grpc_server.go`
- `grpcx_grpc_server_config.go`
- `grpcx_grpc_server_unary.go`
- `grpcx_interceptor_client.go`
- `grpcx_interceptor_server.go`
- `grpcx_registry_file.go`

## 类型定义

### GrpcServer

**类型**: struct

GrpcServer is the server for GRPC protocol.


**方法**:

- `func (s *GrpcServer) Service(services ...gsvc.Service)`
  Service binds service list to current server.
- `func (s *GrpcServer) Run()`
  Run starts the server in blocking way.
- `func (s *GrpcServer) Logger() *glog.Logger`
  Logger is alias of GetLogger.
- `func (s *GrpcServer) Start()`
  Start starts the server in no-blocking way.
- `func (s *GrpcServer) Wait()`
  Wait works with Start, which blocks current goroutine until the server stops.
- `func (s *GrpcServer) Stop()`
  Stop gracefully stops the server.
- `func (s *GrpcServer) GetConfig() *GrpcServerConfig`
  GetConfig returns the configuration of current Server.
- `func (s *GrpcServer) GetListenedAddress() string`
  GetListenedAddress retrieves and returns the address string which are listene...
- `func (s *GrpcServer) GetListenedPort() int`
  GetListenedPort retrieves and returns one port which is listened to by curren...
- `func (s *GrpcServer) UnaryLogger(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error)`
  UnaryLogger is the default unary interceptor for logging purpose.

### Service

**类型**: struct

Service implements gsvc.Service interface.


### GrpcServerConfig

**类型**: struct

GrpcServerConfig is the configuration for server.


**方法**:

- `func (c *GrpcServerConfig) SetWithMap(m g.Map) error`
  SetWithMap changes current configuration with map.
- `func (c *GrpcServerConfig) MustSetWithMap(m g.Map)`
  MustSetWithMap acts as SetWithMap but panics if error occurs.

## 内部依赖

- `github.com/gogf/gf/contrib/registry/file/v2`
- `internal/balancer`
- `internal/grpcctx`
- `internal/resolver`
- `internal/tracing`
- `errors/gcode`
- `errors/gerror`
- `frame/g`
- `net/gipv4`
- `net/gsel`
- `net/gsvc`
- `os/gctx`
- `os/gfile`
- `os/glog`
- `os/gproc`
- `text/gstr`
- `util/gconv`
- `util/gutil`

## 外部依赖

- `google.golang.org/grpc`
- `google.golang.org/grpc/codes`
- `google.golang.org/grpc/credentials/insecure`
- `google.golang.org/grpc/status`
- `google.golang.org/protobuf/proto`

