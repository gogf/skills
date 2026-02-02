# gtype

> Package: `github.com/gogf/gf/v2/container/gtype`

```go
import "github.com/gogf/gf/v2/container/gtype"
```

## 概述

Package gtype provides high performance and concurrent-safe basic variable types.

## 源文件

- `gtype.go`
- `gtype_any.go`
- `gtype_bool.go`
- `gtype_byte.go`
- `gtype_bytes.go`
- `gtype_float32.go`
- `gtype_float64.go`
- `gtype_int.go`
- `gtype_int32.go`
- `gtype_int64.go`
- `gtype_interface.go`
- `gtype_string.go`
- `gtype_uint.go`
- `gtype_uint32.go`
- `gtype_uint64.go`

## 类型定义

### Any

**类型**: type

Any is a struct for concurrent-safe operation for type any.


### Bool

**类型**: struct

Bool is a struct for concurrent-safe operation for type bool.


**方法**:

- `func (v *Bool) Clone() *Bool`
  Clone clones and returns a new concurrent-safe object for bool type.
- `func (v *Bool) Set(value bool) old bool`
  Set atomically stores `value` into t.value and returns the previous value of ...
- `func (v *Bool) Val() bool`
  Val atomically loads and returns t.value.
- `func (v *Bool) Cas(old bool, new bool) swapped bool`
  Cas executes the compare-and-swap operation for value.
- `func (v *Bool) String() string`
  String implements String interface for string printing.
- `func (v Bool) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (v *Bool) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (v *Bool) UnmarshalValue(value any) error`
  UnmarshalValue is an interface implement which sets any type of value for `v`.
- `func (v *Bool) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### Byte

**类型**: struct

Byte is a struct for concurrent-safe operation for type byte.


**方法**:

- `func (v *Byte) Clone() *Byte`
  Clone clones and returns a new concurrent-safe object for byte type.
- `func (v *Byte) Set(value byte) old byte`
  Set atomically stores `value` into t.value and returns the previous value of ...
- `func (v *Byte) Val() byte`
  Val atomically loads and returns t.value.
- `func (v *Byte) Add(delta byte) new byte`
  Add atomically adds `delta` to t.value and returns the new value.
- `func (v *Byte) Cas(old byte, new byte) swapped bool`
  Cas executes the compare-and-swap operation for value.
- `func (v *Byte) String() string`
  String implements String interface for string printing.
- `func (v Byte) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (v *Byte) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (v *Byte) UnmarshalValue(value any) error`
  UnmarshalValue is an interface implement which sets any type of value for `v`.
- `func (v *Byte) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### Bytes

**类型**: struct

Bytes is a struct for concurrent-safe operation for type []byte.


**方法**:

- `func (v *Bytes) Clone() *Bytes`
  Clone clones and returns a new shallow copy object for []byte type.
- `func (v *Bytes) Set(value []byte) old []byte`
  Set atomically stores `value` into t.value and returns the previous value of ...
- `func (v *Bytes) Val() []byte`
  Val atomically loads and returns t.value.
- `func (v *Bytes) String() string`
  String implements String interface for string printing.
- `func (v Bytes) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (v *Bytes) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (v *Bytes) UnmarshalValue(value any) error`
  UnmarshalValue is an interface implement which sets any type of value for `v`.
- `func (v *Bytes) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### Float32

**类型**: struct

Float32 is a struct for concurrent-safe operation for type float32.


**方法**:

- `func (v *Float32) Clone() *Float32`
  Clone clones and returns a new concurrent-safe object for float32 type.
- `func (v *Float32) Set(value float32) old float32`
  Set atomically stores `value` into t.value and returns the previous value of ...
- `func (v *Float32) Val() float32`
  Val atomically loads and returns t.value.
- `func (v *Float32) Add(delta float32) new float32`
  Add atomically adds `delta` to t.value and returns the new value.
- `func (v *Float32) Cas(old float32, new float32) swapped bool`
  Cas executes the compare-and-swap operation for value.
- `func (v *Float32) String() string`
  String implements String interface for string printing.
- `func (v Float32) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (v *Float32) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (v *Float32) UnmarshalValue(value any) error`
  UnmarshalValue is an interface implement which sets any type of value for `v`.
- `func (v *Float32) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### Float64

**类型**: struct

Float64 is a struct for concurrent-safe operation for type float64.


**方法**:

- `func (v *Float64) Clone() *Float64`
  Clone clones and returns a new concurrent-safe object for float64 type.
- `func (v *Float64) Set(value float64) old float64`
  Set atomically stores `value` into t.value and returns the previous value of ...
- `func (v *Float64) Val() float64`
  Val atomically loads and returns t.value.
- `func (v *Float64) Add(delta float64) new float64`
  Add atomically adds `delta` to t.value and returns the new value.
- `func (v *Float64) Cas(old float64, new float64) swapped bool`
  Cas executes the compare-and-swap operation for value.
- `func (v *Float64) String() string`
  String implements String interface for string printing.
- `func (v Float64) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (v *Float64) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (v *Float64) UnmarshalValue(value any) error`
  UnmarshalValue is an interface implement which sets any type of value for `v`.
- `func (v *Float64) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### Int

**类型**: struct

Int is a struct for concurrent-safe operation for type int.


**方法**:

- `func (v *Int) Clone() *Int`
  Clone clones and returns a new concurrent-safe object for int type.
- `func (v *Int) Set(value int) old int`
  Set atomically stores `value` into t.value and returns the previous value of ...
- `func (v *Int) Val() int`
  Val atomically loads and returns t.value.
- `func (v *Int) Add(delta int) new int`
  Add atomically adds `delta` to t.value and returns the new value.
- `func (v *Int) Cas(old int, new int) swapped bool`
  Cas executes the compare-and-swap operation for value.
- `func (v *Int) String() string`
  String implements String interface for string printing.
- `func (v Int) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (v *Int) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (v *Int) UnmarshalValue(value any) error`
  UnmarshalValue is an interface implement which sets any type of value for `v`.
- `func (v *Int) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### Int32

**类型**: struct

Int32 is a struct for concurrent-safe operation for type int32.


**方法**:

- `func (v *Int32) Clone() *Int32`
  Clone clones and returns a new concurrent-safe object for int32 type.
- `func (v *Int32) Set(value int32) old int32`
  Set atomically stores `value` into t.value and returns the previous value of ...
- `func (v *Int32) Val() int32`
  Val atomically loads and returns t.value.
- `func (v *Int32) Add(delta int32) new int32`
  Add atomically adds `delta` to t.value and returns the new value.
- `func (v *Int32) Cas(old int32, new int32) swapped bool`
  Cas executes the compare-and-swap operation for value.
- `func (v *Int32) String() string`
  String implements String interface for string printing.
- `func (v Int32) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (v *Int32) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (v *Int32) UnmarshalValue(value any) error`
  UnmarshalValue is an interface implement which sets any type of value for `v`.
- `func (v *Int32) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### Int64

**类型**: struct

Int64 is a struct for concurrent-safe operation for type int64.


**方法**:

- `func (v *Int64) Clone() *Int64`
  Clone clones and returns a new concurrent-safe object for int64 type.
- `func (v *Int64) Set(value int64) old int64`
  Set atomically stores `value` into t.value and returns the previous value of ...
- `func (v *Int64) Val() int64`
  Val atomically loads and returns t.value.
- `func (v *Int64) Add(delta int64) new int64`
  Add atomically adds `delta` to t.value and returns the new value.
- `func (v *Int64) Cas(old int64, new int64) swapped bool`
  Cas executes the compare-and-swap operation for value.
- `func (v *Int64) String() string`
  String implements String interface for string printing.
- `func (v Int64) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (v *Int64) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (v *Int64) UnmarshalValue(value any) error`
  UnmarshalValue is an interface implement which sets any type of value for `v`.
- `func (v *Int64) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### Interface

**类型**: struct

Interface is a struct for concurrent-safe operation for type any.


**方法**:

- `func (v *Interface) Clone() *Interface`
  Clone clones and returns a new concurrent-safe object for any type.
- `func (v *Interface) Set(value any) old any`
  Set atomically stores `value` into t.value and returns the previous value of ...
- `func (v *Interface) Val() any`
  Val atomically loads and returns t.value.
- `func (v *Interface) String() string`
  String implements String interface for string printing.
- `func (v Interface) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (v *Interface) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (v *Interface) UnmarshalValue(value any) error`
  UnmarshalValue is an interface implement which sets any type of value for `v`.
- `func (v *Interface) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### String

**类型**: struct

String is a struct for concurrent-safe operation for type string.


**方法**:

- `func (v *String) Clone() *String`
  Clone clones and returns a new concurrent-safe object for string type.
- `func (v *String) Set(value string) old string`
  Set atomically stores `value` into t.value and returns the previous value of ...
- `func (v *String) Val() string`
  Val atomically loads and returns t.value.
- `func (v *String) String() string`
  String implements String interface for string printing.
- `func (v String) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (v *String) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (v *String) UnmarshalValue(value any) error`
  UnmarshalValue is an interface implement which sets any type of value for `v`.
- `func (v *String) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### Uint

**类型**: struct

Uint is a struct for concurrent-safe operation for type uint.


**方法**:

- `func (v *Uint) Clone() *Uint`
  Clone clones and returns a new concurrent-safe object for uint type.
- `func (v *Uint) Set(value uint) old uint`
  Set atomically stores `value` into t.value and returns the previous value of ...
- `func (v *Uint) Val() uint`
  Val atomically loads and returns t.value.
- `func (v *Uint) Add(delta uint) new uint`
  Add atomically adds `delta` to t.value and returns the new value.
- `func (v *Uint) Cas(old uint, new uint) swapped bool`
  Cas executes the compare-and-swap operation for value.
- `func (v *Uint) String() string`
  String implements String interface for string printing.
- `func (v Uint) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (v *Uint) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (v *Uint) UnmarshalValue(value any) error`
  UnmarshalValue is an interface implement which sets any type of value for `v`.
- `func (v *Uint) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### Uint32

**类型**: struct

Uint32 is a struct for concurrent-safe operation for type uint32.


**方法**:

- `func (v *Uint32) Clone() *Uint32`
  Clone clones and returns a new concurrent-safe object for uint32 type.
- `func (v *Uint32) Set(value uint32) old uint32`
  Set atomically stores `value` into t.value and returns the previous value of ...
- `func (v *Uint32) Val() uint32`
  Val atomically loads and returns t.value.
- `func (v *Uint32) Add(delta uint32) new uint32`
  Add atomically adds `delta` to t.value and returns the new value.
- `func (v *Uint32) Cas(old uint32, new uint32) swapped bool`
  Cas executes the compare-and-swap operation for value.
- `func (v *Uint32) String() string`
  String implements String interface for string printing.
- `func (v Uint32) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (v *Uint32) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (v *Uint32) UnmarshalValue(value any) error`
  UnmarshalValue is an interface implement which sets any type of value for `v`.
- `func (v *Uint32) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

### Uint64

**类型**: struct

Uint64 is a struct for concurrent-safe operation for type uint64.


**方法**:

- `func (v *Uint64) Clone() *Uint64`
  Clone clones and returns a new concurrent-safe object for uint64 type.
- `func (v *Uint64) Set(value uint64) old uint64`
  Set atomically stores `value` into t.value and returns the previous value of ...
- `func (v *Uint64) Val() uint64`
  Val atomically loads and returns t.value.
- `func (v *Uint64) Add(delta uint64) new uint64`
  Add atomically adds `delta` to t.value and returns the new value.
- `func (v *Uint64) Cas(old uint64, new uint64) swapped bool`
  Cas executes the compare-and-swap operation for value.
- `func (v *Uint64) String() string`
  String implements String interface for string printing.
- `func (v Uint64) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (v *Uint64) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (v *Uint64) UnmarshalValue(value any) error`
  UnmarshalValue is an interface implement which sets any type of value for `v`.
- `func (v *Uint64) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.

## 函数

### New

```go
func New(value ...any) *Any
```

New is alias of NewAny.
See NewAny, NewInterface.

### NewAny

```go
func NewAny(value ...any) *Any
```

NewAny creates and returns a concurrent-safe object for any type,
with given initial value `value`.

### NewBool

```go
func NewBool(value ...bool) *Bool
```

NewBool creates and returns a concurrent-safe object for bool type,
with given initial value `value`.

### NewByte

```go
func NewByte(value ...byte) *Byte
```

NewByte creates and returns a concurrent-safe object for byte type,
with given initial value `value`.

### NewBytes

```go
func NewBytes(value ...[]byte) *Bytes
```

NewBytes creates and returns a concurrent-safe object for []byte type,
with given initial value `value`.

### NewFloat32

```go
func NewFloat32(value ...float32) *Float32
```

NewFloat32 creates and returns a concurrent-safe object for float32 type,
with given initial value `value`.

### NewFloat64

```go
func NewFloat64(value ...float64) *Float64
```

NewFloat64 creates and returns a concurrent-safe object for float64 type,
with given initial value `value`.

### NewInt

```go
func NewInt(value ...int) *Int
```

NewInt creates and returns a concurrent-safe object for int type,
with given initial value `value`.

### NewInt32

```go
func NewInt32(value ...int32) *Int32
```

NewInt32 creates and returns a concurrent-safe object for int32 type,
with given initial value `value`.

### NewInt64

```go
func NewInt64(value ...int64) *Int64
```

NewInt64 creates and returns a concurrent-safe object for int64 type,
with given initial value `value`.

### NewInterface

```go
func NewInterface(value ...any) *Interface
```

NewInterface creates and returns a concurrent-safe object for any type,
with given initial value `value`.

### NewString

```go
func NewString(value ...string) *String
```

NewString creates and returns a concurrent-safe object for string type,
with given initial value `value`.

### NewUint

```go
func NewUint(value ...uint) *Uint
```

NewUint creates and returns a concurrent-safe object for uint type,
with given initial value `value`.

### NewUint32

```go
func NewUint32(value ...uint32) *Uint32
```

NewUint32 creates and returns a concurrent-safe object for uint32 type,
with given initial value `value`.

### NewUint64

```go
func NewUint64(value ...uint64) *Uint64
```

NewUint64 creates and returns a concurrent-safe object for uint64 type,
with given initial value `value`.

## 内部依赖

- `errors/gerror`
- `internal/deepcopy`
- `internal/json`
- `util/gconv`

