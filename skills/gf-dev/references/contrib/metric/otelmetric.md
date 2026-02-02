# otelmetric

> Package: `github.com/gogf/gf/contrib/metric/otelmetric/v2`

```go
import "github.com/gogf/gf/contrib/metric/otelmetric/v2"
```

## 概述

Package otelmetric provides metric functionalities using OpenTelemetry metric.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/metric/otelmetric/v2`

## 源文件

- `otelmetric.go`
- `otelmetric_callback.go`
- `otelmetric_meter_counter_performer.go`
- `otelmetric_meter_histogram_performer.go`
- `otelmetric_meter_observable_counter_performer.go`
- `otelmetric_meter_observable_gauge_performer.go`
- `otelmetric_meter_observable_updown_counter_performer.go`
- `otelmetric_meter_performer.go`
- `otelmetric_meter_updown_counter_performer.go`
- `otelmetric_metric_callback.go`
- `otelmetric_option.go`
- `otelmetric_prometheus.go`
- `otelmetric_provider.go`
- `otelmetric_util.go`

## 类型定义

### Option

**类型**: interface

Option applies a configuration option value to a MeterProvider.


## 函数

### NewProvider

```go
func NewProvider(option ...Option) (gmetric.Provider, error)
```

NewProvider creates and returns a metrics provider.

### MustProvider

```go
func MustProvider(option ...Option) gmetric.Provider
```

MustProvider creates and returns a metrics provider.
It panics if any error occurs.

### WithBuiltInMetrics

```go
func WithBuiltInMetrics() Option
```

WithBuiltInMetrics enables builtin metrics.

### WithResource

```go
func WithResource(res *resource.Resource) Option
```

WithResource associates a Resource with a MeterProvider. This Resource
represents the entity producing telemetry and is associated with all Meters
the MeterProvider will create.

### WithReader

```go
func WithReader(reader metric.Reader) Option
```

WithReader associates Reader r with a MeterProvider.

By default, if this option is not used, the MeterProvider will perform no
operations; no data will be exported without a Reader.

### WithView

```go
func WithView(views ...metric.View) Option
```

WithView associates views a MeterProvider.

Views are appended to existing ones in a MeterProvider if this option is
used multiple times.

By default, if this option is not used, the MeterProvider will use the
default view.

### WithExemplarFilter

```go
func WithExemplarFilter(filter exemplar.Filter) Option
```

WithExemplarFilter configures the exemplar filter.

The exemplar filter determines which measurements are offered to the
exemplar reservoir, but the exemplar reservoir makes the final decision of
whether to store an exemplar.

By default, the [exemplar.SampledFilter]
is used. Exemplars can be entirely disabled by providing the
[exemplar.AlwaysOffFilter].

### PrometheusHandler

```go
func PrometheusHandler(r *ghttp.Request)
```

PrometheusHandler returns the http handler for prometheus metrics exporting.

### StartPrometheusMetricsServer

```go
func StartPrometheusMetricsServer(port int, path string)
```

StartPrometheusMetricsServer starts running a http server for metrics exporting.

## 内部依赖

- `container/gset`
- `encoding/gjson`
- `errors/gcode`
- `errors/gerror`
- `frame/g`
- `net/ghttp`
- `os/gmetric`
- `util/gconv`

## 外部依赖

- `github.com/prometheus/client_golang/prometheus`
- `github.com/prometheus/client_golang/prometheus/collectors`
- `github.com/prometheus/client_golang/prometheus/promhttp`
- `go.opentelemetry.io/contrib/instrumentation/runtime`
- `go.opentelemetry.io/otel`
- `go.opentelemetry.io/otel/attribute`
- `go.opentelemetry.io/otel/metric`
- `go.opentelemetry.io/otel/sdk/instrumentation`
- `go.opentelemetry.io/otel/sdk/metric`
- `go.opentelemetry.io/otel/sdk/metric/exemplar`
- `go.opentelemetry.io/otel/sdk/resource`

