# gredis

> Package: `github.com/gogf/gf/v2/database/gredis`

```go
import "github.com/gogf/gf/v2/database/gredis"
```

## 概述

Package gredis provides convenient client for redis server.

## 源文件

- `gredis.go`
- `gredis_adapter.go`
- `gredis_config.go`
- `gredis_instance.go`
- `gredis_redis.go`
- `gredis_redis_group_generic.go`
- `gredis_redis_group_hash.go`
- `gredis_redis_group_list.go`
- `gredis_redis_group_pubsub.go`
- `gredis_redis_group_script.go`
- `gredis_redis_group_set.go`
- `gredis_redis_group_sorted_set.go`
- `gredis_redis_group_string.go`

## 类型定义

### AdapterFunc

**类型**: type

AdapterFunc is the function creating redis adapter.


### Adapter

**类型**: interface

Adapter is an interface for universal redis operations.


### AdapterGroup

**类型**: interface

AdapterGroup is an interface managing group operations for redis.


### RedisRawClient

**类型**: type

RedisRawClient is a type alias for any, representing the raw underlying redis client.
Implementations should return their concrete client type as this interface.


### AdapterOperation

**类型**: interface

AdapterOperation is the core operation functions for redis.
These functions can be easily overwritten by custom implements.


### Conn

**类型**: interface

Conn is an interface of a connection from universal redis client.


### ConnCommand

**类型**: interface

ConnCommand is an interface managing some operations bound to certain connection.


### Config

**类型**: struct

Config is redis configuration.


### Redis

**类型**: struct

Redis client.


**方法**:

- `func (r *Redis) SetAdapter(adapter Adapter)`
  SetAdapter changes the underlying adapter with custom adapter for current red...
- `func (r *Redis) GetAdapter() Adapter`
  GetAdapter returns the adapter that is set in current redis client.
- `func (r *Redis) Conn(ctx context.Context) (Conn, error)`
  Conn retrieves and returns a connection object for continuous operations.
- `func (r *Redis) Do(ctx context.Context, command string, args ...any) (*gvar.Var, error)`
  Do send a command to the server and returns the received reply.
- `func (r *Redis) MustConn(ctx context.Context) Conn`
  MustConn performs as function Conn, but it panics if any error occurs interna...
- `func (r *Redis) MustDo(ctx context.Context, command string, args ...any) *gvar.Var`
  MustDo performs as function Do, but it panics if any error occurs internally.
- `func (r *Redis) Close(ctx context.Context) error`
  Close closes current redis client, closes its connection pool and releases al...

### IGroupGeneric

**类型**: interface

IGroupGeneric manages generic redis operations.
Implements see redis.GroupGeneric.


### CopyOption

**类型**: struct

CopyOption provides options for function Copy.


### FlushOp

**类型**: type

### ExpireOption

**类型**: struct

ExpireOption provides options for function Expire.


### ScanOption

**类型**: struct

ScanOption provides options for function Scan.


**方法**:

- `func (scanOpt *ScanOption) ToUsedOption() doScanOption`
  ToUsedOption converts fields in ScanOption with zero values to nil. Only fiel...

### IGroupHash

**类型**: interface

IGroupHash manages redis hash operations.
Implements see redis.GroupHash.


### IGroupList

**类型**: interface

IGroupList manages redis list operations.
Implements see redis.GroupList.


### LInsertOp

**类型**: type

LInsertOp defines the operation name for function LInsert.


### IGroupPubSub

**类型**: interface

IGroupPubSub manages redis pub/sub operations.
Implements see redis.GroupPubSub.


### Message

**类型**: struct

Message received as result of a PUBLISH command issued by another client.


### Subscription

**类型**: struct

Subscription received after a successful subscription to channel.


**方法**:

- `func (m *Subscription) String() string`
  String converts current object to a readable string.

### IGroupScript

**类型**: interface

IGroupScript manages redis script operations.
Implements see redis.GroupScript.


### ScriptFlushOption

**类型**: struct

ScriptFlushOption provides options for function ScriptFlush.


### IGroupSet

**类型**: interface

IGroupSet manages redis set operations.
Implements see redis.GroupSet.


### IGroupSortedSet

**类型**: interface

IGroupSortedSet manages redis sorted set operations.
Implements see redis.GroupSortedSet.


### ZAddOption

**类型**: struct

ZAddOption provides options for function ZAdd.


### ZAddMember

**类型**: struct

ZAddMember is element struct for set.


### ZRangeOption

**类型**: struct

ZRangeOption provides extra option for ZRange function.


### ZRangeOptionLimit

**类型**: struct

ZRangeOptionLimit provides LIMIT argument for ZRange function.
The optional LIMIT argument can be used to obtain a sub-range from the matching elements
(similar to SELECT LIMIT offset, count in SQL). A negative `Count` returns all elements from the `Offset`.


### ZRevRangeOption

**类型**: struct

ZRevRangeOption provides options for function ZRevRange.


### IGroupString

**类型**: interface

IGroupString manages redis string operations.
Implements see redis.GroupString.


### TTLOption

**类型**: struct

TTLOption provides extra option for TTL related functions.


### SetOption

**类型**: struct

SetOption provides extra option for Set function.


### GetEXOption

**类型**: struct

GetEXOption provides extra option for GetEx function.


## 函数

### New

```go
func New(config ...*Config) (*Redis, error)
```

New creates and returns a redis client.
It creates a default redis adapter of go-redis.

### NewWithAdapter

```go
func NewWithAdapter(adapter Adapter) (*Redis, error)
```

NewWithAdapter creates and returns a redis client with given adapter.

### RegisterAdapterFunc

```go
func RegisterAdapterFunc(adapterFunc AdapterFunc)
```

RegisterAdapterFunc registers default function creating redis adapter.

### SetConfig

```go
func SetConfig(config *Config, name ...string)
```

SetConfig sets the global configuration for specified group.
If `name` is not passed, it sets configuration for the default group name.

### SetConfigByMap

```go
func SetConfigByMap(m map[string]any, name ...string) error
```

SetConfigByMap sets the global configuration for specified group with map.
If `name` is not passed, it sets configuration for the default group name.

### ConfigFromMap

```go
func ConfigFromMap(m map[string]any) (config *Config, err error)
```

ConfigFromMap parses and returns config from given map.

### GetConfig

```go
func GetConfig(name ...string) (config *Config, ok bool)
```

GetConfig returns the global configuration with specified group name.
If `name` is not passed, it returns configuration of the default group name.

### RemoveConfig

```go
func RemoveConfig(name ...string)
```

RemoveConfig removes the global configuration with specified group.
If `name` is not passed, it removes configuration of the default group name.

### ClearConfig

```go
func ClearConfig()
```

ClearConfig removes all configurations of redis.

### Instance

```go
func Instance(name ...string) *Redis
```

Instance returns an instance of redis client with specified group.
The `name` param is unnecessary, if `name` is not passed,
it returns a redis instance with default configuration group.

## 内部依赖

- `container/gmap`
- `container/gvar`
- `errors/gcode`
- `errors/gerror`
- `internal/intlog`
- `util/gconv`

