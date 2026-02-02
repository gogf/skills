# gsel

> Package: `github.com/gogf/gf/v2/net/gsel`

```go
import "github.com/gogf/gf/v2/net/gsel"
```

## 概述

Package gsel provides selector definition and implements.

## 源文件

- `gsel.go`
- `gsel_builder.go`
- `gsel_builder_least_connection.go`
- `gsel_builder_random.go`
- `gsel_builder_round_robin.go`
- `gsel_builder_weight.go`
- `gsel_selector_least_connection.go`
- `gsel_selector_random.go`
- `gsel_selector_round_robin.go`
- `gsel_selector_weight.go`

## 类型定义

### Builder

**类型**: interface

Builder creates and returns selector in runtime.


### Selector

**类型**: interface

Selector for service balancer.


### Node

**类型**: interface

Node is node interface.


### Nodes

**类型**: type

Nodes contains multiple Node.


**方法**:

- `func (ns Nodes) String() string`
  String formats and returns Nodes as string.

### DoneFunc

**类型**: type

DoneFunc is callback function when RPC invoke done.


### DoneInfo

**类型**: struct

DoneInfo contains additional information for done.


### DoneInfoMD

**类型**: interface

DoneInfoMD is a mapping from metadata keys to value array.
Users should use the following two convenience functions New and Pairs to generate MD.


## 函数

### SetBuilder

```go
func SetBuilder(builder Builder)
```

SetBuilder sets the default builder for globally used purpose.

### GetBuilder

```go
func GetBuilder() Builder
```

GetBuilder returns the default builder for globally used purpose.

### NewBuilderLeastConnection

```go
func NewBuilderLeastConnection() Builder
```

### NewBuilderRandom

```go
func NewBuilderRandom() Builder
```

### NewBuilderRoundRobin

```go
func NewBuilderRoundRobin() Builder
```

### NewBuilderWeight

```go
func NewBuilderWeight() Builder
```

### NewSelectorLeastConnection

```go
func NewSelectorLeastConnection() Selector
```

### NewSelectorRandom

```go
func NewSelectorRandom() Selector
```

### NewSelectorRoundRobin

```go
func NewSelectorRoundRobin() Selector
```

### NewSelectorWeight

```go
func NewSelectorWeight() Selector
```

## 内部依赖

- `container/gtype`
- `internal/intlog`
- `net/gsvc`
- `util/grand`

