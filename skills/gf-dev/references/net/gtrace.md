# gtrace

> Package: `github.com/gogf/gf/v2/net/gtrace`

```go
import "github.com/gogf/gf/v2/net/gtrace"
```

## 概述

Package gtrace provides convenience wrapping functionality for tracing feature using OpenTelemetry.

## 源文件

- `gtrace.go`
- `gtrace_baggage.go`
- `gtrace_carrier.go`
- `gtrace_content.go`
- `gtrace_span.go`
- `gtrace_tracer.go`

## 类型定义

### Baggage

**类型**: struct

Baggage holds the data through all tracing spans.


**方法**:

- `func (b *Baggage) Ctx() context.Context`
  Ctx returns the context that Baggage holds.
- `func (b *Baggage) SetValue(key string, value any) context.Context`
  SetValue is a convenient function for adding one key-value pair to baggage.
- `func (b *Baggage) SetMap(data map[string]any) context.Context`
  SetMap is a convenient function for adding map key-value pairs to baggage.
- `func (b *Baggage) GetMap() *gmap.StrAnyMap`
  GetMap retrieves and returns the baggage values as map.
- `func (b *Baggage) GetVar(key string) *gvar.Var`
  GetVar retrieves value and returns a *gvar.Var for specified key from baggage.

### Carrier

**类型**: type

Carrier is the storage medium used by a TextMapPropagator.


**方法**:

- `func (c Carrier) Get(k string) string`
  Get returns the value associated with the passed key.
- `func (c Carrier) Set(k string, v string)`
  Set stores the key-value pair.
- `func (c Carrier) Keys() []string`
  Keys lists the keys stored in this carrier.
- `func (c Carrier) MustMarshal() []byte`
  MustMarshal .returns the JSON encoding of c
- `func (c Carrier) String() string`
  String converts and returns current Carrier as string.
- `func (c Carrier) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements interface UnmarshalJSON for package json.

### Span

**类型**: struct

Span warps trace.Span for compatibility and extension.


### Tracer

**类型**: struct

Tracer warps trace.Tracer for compatibility and extension.


## 函数

### IsUsingDefaultProvider

```go
func IsUsingDefaultProvider() bool
```

IsUsingDefaultProvider checks and return if currently using default trace provider.

### IsTracingInternal

```go
func IsTracingInternal() bool
```

IsTracingInternal returns whether tracing spans of internal components.

### MaxContentLogSize

```go
func MaxContentLogSize() int
```

MaxContentLogSize returns the max log size for request and response body, especially for HTTP/RPC request.

### CommonLabels

```go
func CommonLabels() []attribute.KeyValue
```

CommonLabels returns common used attribute labels:
ip.intranet, hostname.

### CheckSetDefaultTextMapPropagator

```go
func CheckSetDefaultTextMapPropagator()
```

CheckSetDefaultTextMapPropagator sets the default TextMapPropagator if it is not set previously.

### GetDefaultTextMapPropagator

```go
func GetDefaultTextMapPropagator() propagation.TextMapPropagator
```

GetDefaultTextMapPropagator returns the default propagator for context propagation between peers.

### GetTraceID

```go
func GetTraceID(ctx context.Context) string
```

GetTraceID retrieves and returns TraceId from context.
It returns an empty string is tracing feature is not activated.

### GetSpanID

```go
func GetSpanID(ctx context.Context) string
```

GetSpanID retrieves and returns SpanId from context.
It returns an empty string is tracing feature is not activated.

### SetBaggageValue

```go
func SetBaggageValue(ctx context.Context, key string, value any) context.Context
```

SetBaggageValue is a convenient function for adding one key-value pair to baggage.
Note that it uses attribute.Any to set the key-value pair.

### SetBaggageMap

```go
func SetBaggageMap(ctx context.Context, data map[string]any) context.Context
```

SetBaggageMap is a convenient function for adding map key-value pairs to baggage.
Note that it uses attribute.Any to set the key-value pair.

### GetBaggageMap

```go
func GetBaggageMap(ctx context.Context) *gmap.StrAnyMap
```

GetBaggageMap retrieves and returns the baggage values as map.

### GetBaggageVar

```go
func GetBaggageVar(ctx context.Context, key string) *gvar.Var
```

GetBaggageVar retrieves value and returns a *gvar.Var for specified key from baggage.

### WithUUID

```go
func WithUUID(ctx context.Context, uuid string) (context.Context, error)
```

WithUUID injects custom trace id with UUID into context to propagate.

### WithTraceID

```go
func WithTraceID(ctx context.Context, traceID string) (context.Context, error)
```

WithTraceID injects custom trace id into context to propagate.

### NewBaggage

```go
func NewBaggage(ctx context.Context) *Baggage
```

NewBaggage creates and returns a new Baggage object from given tracing context.

### NewCarrier

```go
func NewCarrier(data ...map[string]any) Carrier
```

NewCarrier creates and returns a Carrier.

### SafeContentForHttp

```go
func SafeContentForHttp(data []byte, header http.Header) (string, error)
```

SafeContentForHttp cuts and returns given content by `MaxContentLogSize`.
It appends string `...` to the tail of the result if the content size is greater than `MaxContentLogSize`.

### SafeContent

```go
func SafeContent(data []byte) string
```

SafeContent cuts and returns given content by `MaxContentLogSize`.
It appends string `...` to the tail of the result if the content size is greater than `MaxContentLogSize`.

### NewSpan

```go
func NewSpan(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, *Span)
```

NewSpan creates a span using default tracer.

### NewTracer

```go
func NewTracer(name ...string) *Tracer
```

NewTracer Tracer is a short function for retrieving Tracer.

## 内部依赖

- `container/gmap`
- `container/gvar`
- `encoding/gcompress`
- `errors/gcode`
- `errors/gerror`
- `internal/command`
- `internal/json`
- `net/gipv4`
- `net/gtrace/internal/provider`
- `text/gstr`
- `util/gconv`

## 外部依赖

- `go.opentelemetry.io/otel`
- `go.opentelemetry.io/otel/attribute`
- `go.opentelemetry.io/otel/baggage`
- `go.opentelemetry.io/otel/propagation`
- `go.opentelemetry.io/otel/semconv/v1.24.0`
- `go.opentelemetry.io/otel/trace`

