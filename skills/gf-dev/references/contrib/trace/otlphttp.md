# otlphttp

> Package: `github.com/gogf/gf/contrib/trace/otlphttp/v2`

```go
import "github.com/gogf/gf/contrib/trace/otlphttp/v2"
```

## 概述

Package otlphttp provides gtrace.Tracer implementation using OpenTelemetry protocol.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/trace/otlphttp/v2`

## 源文件

- `otlphttp.go`

## 函数

### Init

```go
func Init(serviceName string, endpoint string, path string) (func, error)
```

Init initializes and registers `otlphttp` to global TracerProvider.

The output parameter `Shutdown` is used for waiting exported trace spans to be uploaded,
which is useful if your program is ending, and you do not want to lose recent spans.

## 内部依赖

- `frame/g`
- `net/gipv4`

## 外部依赖

- `go.opentelemetry.io/otel`
- `go.opentelemetry.io/otel/attribute`
- `go.opentelemetry.io/otel/exporters/otlp/otlptrace`
- `go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp`
- `go.opentelemetry.io/otel/propagation`
- `go.opentelemetry.io/otel/sdk/resource`
- `go.opentelemetry.io/otel/sdk/trace`
- `go.opentelemetry.io/otel/semconv/v1.24.0`

