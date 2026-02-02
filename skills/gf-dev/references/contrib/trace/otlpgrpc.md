# otlpgrpc

> Package: `github.com/gogf/gf/contrib/trace/otlpgrpc/v2`

```go
import "github.com/gogf/gf/contrib/trace/otlpgrpc/v2"
```

## 概述

Package otlpgrpc provides gtrace.Tracer implementation using OpenTelemetry protocol.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/trace/otlpgrpc/v2`

## 源文件

- `otlpgrpc.go`

## 函数

### Init

```go
func Init(serviceName string, endpoint string, traceToken string) (func, error)
```

Init initializes and registers `otlpgrpc` to global TracerProvider.

The output parameter `Shutdown` is used for waiting exported trace spans to be uploaded,
which is useful if your program is ending, and you do not want to lose recent spans.

## 内部依赖

- `frame/g`
- `net/gipv4`

## 外部依赖

- `go.opentelemetry.io/otel`
- `go.opentelemetry.io/otel/attribute`
- `go.opentelemetry.io/otel/exporters/otlp/otlptrace`
- `go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc`
- `go.opentelemetry.io/otel/propagation`
- `go.opentelemetry.io/otel/sdk/resource`
- `go.opentelemetry.io/otel/sdk/trace`
- `go.opentelemetry.io/otel/semconv/v1.24.0`
- `google.golang.org/grpc/encoding/gzip`

