# gmetric

> Package: `github.com/gogf/gf/v2/os/gmetric`

```go
import "github.com/gogf/gf/v2/os/gmetric"
```

## 概述

Package gmetric provides interface definitions and simple api for metric feature.

## 源文件

- `gmetric.go`
- `gmetric_attribute.go`
- `gmetric_attribute_map.go`
- `gmetric_global_attributes.go`
- `gmetric_meter.go`
- `gmetric_meter_callback.go`
- `gmetric_meter_counter.go`
- `gmetric_meter_histogram.go`
- `gmetric_meter_metric_info.go`
- `gmetric_meter_metric_instrument.go`
- `gmetric_meter_observable_counter.go`
- `gmetric_meter_observable_gauge.go`
- `gmetric_meter_observable_updown_counter.go`
- `gmetric_meter_updown_counter.go`
- `gmetric_metric.go`
- `gmetric_noop_counter_performer.go`
- `gmetric_noop_histogram_performer.go`
- `gmetric_noop_observable_counter_performer.go`
- `gmetric_noop_observable_gauge_performer.go`
- `gmetric_noop_observable_updown_counter_performer.go`
- `gmetric_noop_updown_counter_performer.go`
- `gmetric_provider.go`

## 类型定义

### MetricType

**类型**: type

MetricType is the type of metric.


### Provider

**类型**: interface

Provider manages all Metric exporting.
Be caution that the Histogram buckets could not be customized if the creation of the Histogram
is before the creation of Provider.


### MeterPerformer

**类型**: interface

MeterPerformer manages all Metric performers creating.


### MetricOption

**类型**: struct

MetricOption holds the basic options for creating a metric.


### Metric

**类型**: interface

Metric models a single sample value with its metadata being exported.


### MetricInfo

**类型**: interface

MetricInfo exports information of the Metric.


### InstrumentInfo

**类型**: interface

InstrumentInfo exports the instrument information of a metric.


### Counter

**类型**: interface

Counter is a Metric that represents a single numerical value that can ever
goes up.


### CounterPerformer

**类型**: interface

CounterPerformer performs operations for Counter metric.


### UpDownCounter

**类型**: interface

UpDownCounter is a Metric that represents a single numerical value that can ever
goes up or down.


### UpDownCounterPerformer

**类型**: interface

UpDownCounterPerformer performs operations for UpDownCounter metric.


### Histogram

**类型**: interface

Histogram counts individual observations from an event or sample stream in
configurable static buckets (or in dynamic sparse buckets as part of the
experimental Native Histograms, see below for more details). Similar to a
Summary, it also provides a sum of observations and an observation count.


### HistogramPerformer

**类型**: interface

HistogramPerformer performs operations for Histogram metric.


### ObservableCounter

**类型**: interface

ObservableCounter is an instrument used to asynchronously
record float64 measurements once per collection cycle. Observations are only
made within a callback for this instrument. The value observed is assumed
the to be the cumulative sum of the count.


### ObservableUpDownCounter

**类型**: interface

ObservableUpDownCounter is used to synchronously record float64 measurements during a computational
operation.


### ObservableGauge

**类型**: interface

ObservableGauge is an instrument used to asynchronously record
instantaneous float64 measurements once per collection cycle. Observations
are only made within a callback for this instrument.


### ObservableCounterPerformer

**类型**: type

### ObservableUpDownCounterPerformer

**类型**: type

### ObservableGaugePerformer

**类型**: type

### ObservableMetric

**类型**: interface

ObservableMetric is an instrument used to asynchronously record
instantaneous float64 measurements once per collection cycle.


### MetricInitializer

**类型**: interface

MetricInitializer manages the initialization for Metric.
It is called internally in metric interface implements.


### PerformerExporter

**类型**: interface

PerformerExporter exports internal Performer of Metric.
It is called internally in metric interface implements.


### MetricCallback

**类型**: type

MetricCallback is automatically called when metric reader starts reading the metric value.


### Callback

**类型**: type

Callback is a function registered with a Meter that makes observations for
the set of instruments it is registered with. The Observer parameter is used
to record measurement observations for these instruments.


### Observer

**类型**: interface

Observer sets the value for certain initialized Metric.


### MetricObserver

**类型**: interface

MetricObserver sets the value for bound Metric.


### Attributes

**类型**: type

Attributes is a slice of Attribute.


**方法**:

- `func (attrs Attributes) String() string`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (attrs Attributes) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.

### Attribute

**类型**: interface

Attribute is the key-value pair item for Metric.


### AttributeKey

**类型**: type

AttributeKey is the attribute key.


### Option

**类型**: struct

Option holds the option for perform a metric operation.


### AttributeMap

**类型**: type

AttributeMap contains the attribute key and value as map for easy filtering.


**方法**:

- `func (m AttributeMap) Sets(attrMap map[string]any)`
  Sets adds given attribute map to current map.
- `func (m AttributeMap) Pick(keys ...string) Attributes`
  Pick picks and returns attributes by given attribute keys.
- `func (m AttributeMap) PickEx(keys ...string) Attributes`
  PickEx picks and returns attributes of which the given attribute keys does no...

### SetGlobalAttributesOption

**类型**: struct

SetGlobalAttributesOption binds the global attributes to certain instrument.


### GetGlobalAttributesOption

**类型**: struct

GetGlobalAttributesOption binds the global attributes to certain instrument.


### MeterOption

**类型**: struct

MeterOption holds the creation option for a Meter.


### CallbackItem

**类型**: struct

CallbackItem is the global callback item registered.


### GlobalProvider

**类型**: interface

GlobalProvider hold the entry for creating Meter and Metric.
The GlobalProvider has only one function for Meter creating, which is designed for convenient usage.


### Meter

**类型**: interface

Meter hold the functions for kinds of Metric creating.


## 函数

### IsEnabled

```go
func IsEnabled() bool
```

IsEnabled returns whether the metrics feature is enabled.

### GetAllMetrics

```go
func GetAllMetrics() []Metric
```

GetAllMetrics returns all Metric that created by current package.

### CommonAttributes

```go
func CommonAttributes() Attributes
```

CommonAttributes returns the common used attributes for an instrument.

### NewAttribute

```go
func NewAttribute(key string, value any) Attribute
```

NewAttribute creates and returns an Attribute by given `key` and `value`.

### SetGlobalAttributes

```go
func SetGlobalAttributes(attrs Attributes, option SetGlobalAttributesOption)
```

SetGlobalAttributes appends global attributes according `SetGlobalAttributesOption`.
It appends global attributes to all metrics if given `SetGlobalAttributesOption` is empty.
It appends global attributes to certain instrument by given `SetGlobalAttributesOption`.

### GetGlobalAttributes

```go
func GetGlobalAttributes(option GetGlobalAttributesOption) Attributes
```

GetGlobalAttributes retrieves and returns the global attributes by `GetGlobalAttributesOption`.
It returns the global attributes if given `GetGlobalAttributesOption` is empty.
It returns global attributes of certain instrument if `GetGlobalAttributesOption` is not empty.

### GetRegisteredCallbacks

```go
func GetRegisteredCallbacks() []CallbackItem
```

GetRegisteredCallbacks retrieves and returns the registered global callbacks.
It truncates the callback slice is the callbacks are returned.

### GetGlobalProvider

```go
func GetGlobalProvider() GlobalProvider
```

GetGlobalProvider retrieves the GetGlobalProvider instance.

### SetGlobalProvider

```go
func SetGlobalProvider(provider Provider)
```

SetGlobalProvider registers `provider` as the global Provider,
which means the following metrics creating will be base on the global provider.

## 内部依赖

- `encoding/gjson`
- `errors/gcode`
- `errors/gerror`
- `internal/json`
- `os/gfile`
- `text/gregex`

