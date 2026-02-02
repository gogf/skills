# gcache

> Package: `github.com/gogf/gf/v2/os/gcache`

```go
import "github.com/gogf/gf/v2/os/gcache"
```

## 概述

Package gcache provides kinds of cache management for process.

## 源文件

- `gcache.go`
- `gcache_adapter.go`
- `gcache_adapter_memory.go`
- `gcache_adapter_memory_data.go`
- `gcache_adapter_memory_expire_sets.go`
- `gcache_adapter_memory_expire_times.go`
- `gcache_adapter_memory_item.go`
- `gcache_adapter_memory_lru.go`
- `gcache_adapter_redis.go`
- `gcache_cache.go`
- `gcache_cache_must.go`

## 类型定义

### Func

**类型**: type

Func is the cache function that calculates and returns the value.


### Adapter

**类型**: interface

Adapter is the core adapter for cache features implements.

Note that the implementer itself should guarantee the concurrent safety of these functions.


### AdapterMemory

**类型**: struct

AdapterMemory is an adapter implements using memory.


**方法**:

- `func (c *AdapterMemory) Set(ctx context.Context, key any, value any, duration time.Duration) error`
  Set sets cache with `key`-`value` pair, which is expired after `duration`.
- `func (c *AdapterMemory) SetMap(ctx context.Context, data map[any]any, duration time.Duration) error`
  SetMap batch sets cache with key-value pairs by `data` map, which is expired ...
- `func (c *AdapterMemory) SetIfNotExist(ctx context.Context, key any, value any, duration time.Duration) (bool, error)`
  SetIfNotExist sets cache with `key`-`value` pair which is expired after `dura...
- `func (c *AdapterMemory) SetIfNotExistFunc(ctx context.Context, key any, f Func, duration time.Duration) (bool, error)`
  SetIfNotExistFunc sets `key` with result of function `f` and returns true
- `func (c *AdapterMemory) SetIfNotExistFuncLock(ctx context.Context, key any, f Func, duration time.Duration) (bool, error)`
  SetIfNotExistFuncLock sets `key` with result of function `f` and returns true
- `func (c *AdapterMemory) Get(ctx context.Context, key any) (*gvar.Var, error)`
  Get retrieves and returns the associated value of given `key`.
- `func (c *AdapterMemory) GetOrSet(ctx context.Context, key any, value any, duration time.Duration) (*gvar.Var, error)`
  GetOrSet retrieves and returns the value of `key`, or sets `key`-`value` pair...
- `func (c *AdapterMemory) GetOrSetFunc(ctx context.Context, key any, f Func, duration time.Duration) (*gvar.Var, error)`
  GetOrSetFunc retrieves and returns the value of `key`, or sets `key` with res...
- `func (c *AdapterMemory) GetOrSetFuncLock(ctx context.Context, key any, f Func, duration time.Duration) (*gvar.Var, error)`
  GetOrSetFuncLock retrieves and returns the value of `key`, or sets `key` with...
- `func (c *AdapterMemory) Contains(ctx context.Context, key any) (bool, error)`
  Contains checks and returns true if `key` exists in the cache, or else return...
- `func (c *AdapterMemory) GetExpire(ctx context.Context, key any) (time.Duration, error)`
  GetExpire retrieves and returns the expiration of `key` in the cache.
- `func (c *AdapterMemory) Remove(ctx context.Context, keys ...any) (*gvar.Var, error)`
  Remove deletes one or more keys from cache, and returns its value.
- `func (c *AdapterMemory) Update(ctx context.Context, key any, value any) (oldValue *gvar.Var, exist bool, err error)`
  Update updates the value of `key` without changing its expiration and returns...
- `func (c *AdapterMemory) UpdateExpire(ctx context.Context, key any, duration time.Duration) (oldDuration time.Duration, err error)`
  UpdateExpire updates the expiration of `key` and returns the old expiration d...
- `func (c *AdapterMemory) Size(ctx context.Context) (size int, err error)`
  Size returns the size of the cache.
- `func (c *AdapterMemory) Data(ctx context.Context) (map[any]any, error)`
  Data returns a copy of all key-value pairs in the cache as map type.
- `func (c *AdapterMemory) Keys(ctx context.Context) ([]any, error)`
  Keys returns all keys in the cache as slice.
- `func (c *AdapterMemory) Values(ctx context.Context) ([]any, error)`
  Values returns all values in the cache as slice.
- `func (c *AdapterMemory) Clear(ctx context.Context) error`
  Clear clears all data of the cache.
- `func (c *AdapterMemory) Close(ctx context.Context) error`
  Close closes the cache.

### AdapterRedis

**类型**: struct

AdapterRedis is the gcache adapter implements using Redis server.


**方法**:

- `func (c *AdapterRedis) Set(ctx context.Context, key any, value any, duration time.Duration) err error`
  Set sets cache with `key`-`value` pair, which is expired after `duration`.
- `func (c *AdapterRedis) SetMap(ctx context.Context, data map[any]any, duration time.Duration) error`
  SetMap batch sets cache with key-value pairs by `data` map, which is expired ...
- `func (c *AdapterRedis) SetIfNotExist(ctx context.Context, key any, value any, duration time.Duration) (bool, error)`
  SetIfNotExist sets cache with `key`-`value` pair which is expired after `dura...
- `func (c *AdapterRedis) SetIfNotExistFunc(ctx context.Context, key any, f Func, duration time.Duration) (ok bool, err error)`
  SetIfNotExistFunc sets `key` with result of function `f` and returns true
- `func (c *AdapterRedis) SetIfNotExistFuncLock(ctx context.Context, key any, f Func, duration time.Duration) (ok bool, err error)`
  SetIfNotExistFuncLock sets `key` with result of function `f` and returns true
- `func (c *AdapterRedis) Get(ctx context.Context, key any) (*gvar.Var, error)`
  Get retrieves and returns the associated value of given <key>.
- `func (c *AdapterRedis) GetOrSet(ctx context.Context, key any, value any, duration time.Duration) (result *gvar.Var, err error)`
  GetOrSet retrieves and returns the value of `key`, or sets `key`-`value` pair...
- `func (c *AdapterRedis) GetOrSetFunc(ctx context.Context, key any, f Func, duration time.Duration) (result *gvar.Var, err error)`
  GetOrSetFunc retrieves and returns the value of `key`, or sets `key` with res...
- `func (c *AdapterRedis) GetOrSetFuncLock(ctx context.Context, key any, f Func, duration time.Duration) (result *gvar.Var, err error)`
  GetOrSetFuncLock retrieves and returns the value of `key`, or sets `key` with...
- `func (c *AdapterRedis) Contains(ctx context.Context, key any) (bool, error)`
  Contains checks and returns true if `key` exists in the cache, or else return...
- `func (c *AdapterRedis) Size(ctx context.Context) (size int, err error)`
  Size returns the number of items in the cache.
- `func (c *AdapterRedis) Data(ctx context.Context) (map[any]any, error)`
  Data returns a copy of all key-value pairs in the cache as map type.
- `func (c *AdapterRedis) Keys(ctx context.Context) ([]any, error)`
  Keys returns all keys in the cache as slice.
- `func (c *AdapterRedis) Values(ctx context.Context) ([]any, error)`
  Values returns all values in the cache as slice.
- `func (c *AdapterRedis) Update(ctx context.Context, key any, value any) (oldValue *gvar.Var, exist bool, err error)`
  Update updates the value of `key` without changing its expiration and returns...
- `func (c *AdapterRedis) UpdateExpire(ctx context.Context, key any, duration time.Duration) (oldDuration time.Duration, err error)`
  UpdateExpire updates the expiration of `key` and returns the old expiration d...
- `func (c *AdapterRedis) GetExpire(ctx context.Context, key any) (time.Duration, error)`
  GetExpire retrieves and returns the expiration of `key` in the cache.
- `func (c *AdapterRedis) Remove(ctx context.Context, keys ...any) (lastValue *gvar.Var, err error)`
  Remove deletes the one or more keys from cache, and returns its value.
- `func (c *AdapterRedis) Clear(ctx context.Context) err error`
  Clear clears all data of the cache.
- `func (c *AdapterRedis) Close(ctx context.Context) error`
  Close closes the cache.

### Cache

**类型**: struct

Cache struct.


**方法**:

- `func (c *Cache) SetAdapter(adapter Adapter)`
  SetAdapter changes the adapter for this cache.
- `func (c *Cache) GetAdapter() Adapter`
  GetAdapter returns the adapter that is set in current Cache.
- `func (c *Cache) Removes(ctx context.Context, keys []any) error`
  Removes deletes `keys` in the cache.
- `func (c *Cache) KeyStrings(ctx context.Context) ([]string, error)`
  KeyStrings returns all keys in the cache as string slice.
- `func (c *Cache) MustGet(ctx context.Context, key any) *gvar.Var`
  MustGet acts like Get, but it panics if any error occurs.
- `func (c *Cache) MustGetOrSet(ctx context.Context, key any, value any, duration time.Duration) *gvar.Var`
  MustGetOrSet acts like GetOrSet, but it panics if any error occurs.
- `func (c *Cache) MustGetOrSetFunc(ctx context.Context, key any, f Func, duration time.Duration) *gvar.Var`
  MustGetOrSetFunc acts like GetOrSetFunc, but it panics if any error occurs.
- `func (c *Cache) MustGetOrSetFuncLock(ctx context.Context, key any, f Func, duration time.Duration) *gvar.Var`
  MustGetOrSetFuncLock acts like GetOrSetFuncLock, but it panics if any error o...
- `func (c *Cache) MustContains(ctx context.Context, key any) bool`
  MustContains acts like Contains, but it panics if any error occurs.
- `func (c *Cache) MustGetExpire(ctx context.Context, key any) time.Duration`
  MustGetExpire acts like GetExpire, but it panics if any error occurs.
- `func (c *Cache) MustSize(ctx context.Context) int`
  MustSize acts like Size, but it panics if any error occurs.
- `func (c *Cache) MustData(ctx context.Context) map[any]any`
  MustData acts like Data, but it panics if any error occurs.
- `func (c *Cache) MustKeys(ctx context.Context) []any`
  MustKeys acts like Keys, but it panics if any error occurs.
- `func (c *Cache) MustKeyStrings(ctx context.Context) []string`
  MustKeyStrings acts like KeyStrings, but it panics if any error occurs.
- `func (c *Cache) MustValues(ctx context.Context) []any`
  MustValues acts like Values, but it panics if any error occurs.

## 函数

### Set

```go
func Set(ctx context.Context, key any, value any, duration time.Duration) error
```

Set sets cache with `key`-`value` pair, which is expired after `duration`.

It does not expire if `duration` == 0.
It deletes the keys of `data` if `duration` < 0 or given `value` is nil.

### SetMap

```go
func SetMap(ctx context.Context, data map[any]any, duration time.Duration) error
```

SetMap batch sets cache with key-value pairs by `data` map, which is expired after `duration`.

It does not expire if `duration` == 0.
It deletes the keys of `data` if `duration` < 0 or given `value` is nil.

### SetIfNotExist

```go
func SetIfNotExist(ctx context.Context, key any, value any, duration time.Duration) (bool, error)
```

SetIfNotExist sets cache with `key`-`value` pair which is expired after `duration`
if `key` does not exist in the cache. It returns true the `key` does not exist in the
cache, and it sets `value` successfully to the cache, or else it returns false.

It does not expire if `duration` == 0.
It deletes the `key` if `duration` < 0 or given `value` is nil.

### SetIfNotExistFunc

```go
func SetIfNotExistFunc(ctx context.Context, key any, f Func, duration time.Duration) (bool, error)
```

SetIfNotExistFunc sets `key` with result of function `f` and returns true
if `key` does not exist in the cache, or else it does nothing and returns false if `key` already exists.

The parameter `value` can be type of `func() any`, but it does nothing if its
result is nil.

It does not expire if `duration` == 0.
It deletes the `key` if `duration` < 0 or given `value` is nil.

### SetIfNotExistFuncLock

```go
func SetIfNotExistFuncLock(ctx context.Context, key any, f Func, duration time.Duration) (bool, error)
```

SetIfNotExistFuncLock sets `key` with result of function `f` and returns true
if `key` does not exist in the cache, or else it does nothing and returns false if `key` already exists.

It does not expire if `duration` == 0.
It deletes the `key` if `duration` < 0 or given `value` is nil.

Note that it differs from function `SetIfNotExistFunc` is that the function `f` is executed within
writing mutex lock for concurrent safety purpose.

### Get

```go
func Get(ctx context.Context, key any) (*gvar.Var, error)
```

Get retrieves and returns the associated value of given `key`.
It returns nil if it does not exist, or its value is nil, or it's expired.
If you would like to check if the `key` exists in the cache, it's better using function Contains.

### GetOrSet

```go
func GetOrSet(ctx context.Context, key any, value any, duration time.Duration) (*gvar.Var, error)
```

GetOrSet retrieves and returns the value of `key`, or sets `key`-`value` pair and
returns `value` if `key` does not exist in the cache. The key-value pair expires
after `duration`.

It does not expire if `duration` == 0.
It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing
if `value` is a function and the function result is nil.

### GetOrSetFunc

```go
func GetOrSetFunc(ctx context.Context, key any, f Func, duration time.Duration) (*gvar.Var, error)
```

GetOrSetFunc retrieves and returns the value of `key`, or sets `key` with result of
function `f` and returns its result if `key` does not exist in the cache. The key-value
pair expires after `duration`.

It does not expire if `duration` == 0.
It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing
if `value` is a function and the function result is nil.

### GetOrSetFuncLock

```go
func GetOrSetFuncLock(ctx context.Context, key any, f Func, duration time.Duration) (*gvar.Var, error)
```

GetOrSetFuncLock retrieves and returns the value of `key`, or sets `key` with result of
function `f` and returns its result if `key` does not exist in the cache. The key-value
pair expires after `duration`.

It does not expire if `duration` == 0.
It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing
if `value` is a function and the function result is nil.

Note that it differs from function `GetOrSetFunc` is that the function `f` is executed within
writing mutex lock for concurrent safety purpose.

### Contains

```go
func Contains(ctx context.Context, key any) (bool, error)
```

Contains checks and returns true if `key` exists in the cache, or else returns false.

### GetExpire

```go
func GetExpire(ctx context.Context, key any) (time.Duration, error)
```

GetExpire retrieves and returns the expiration of `key` in the cache.

Note that,
It returns 0 if the `key` does not expire.
It returns -1 if the `key` does not exist in the cache.

### Remove

```go
func Remove(ctx context.Context, keys ...any) (value *gvar.Var, err error)
```

Remove deletes one or more keys from cache, and returns its value.
If multiple keys are given, it returns the value of the last deleted item.

### Removes

```go
func Removes(ctx context.Context, keys []any) error
```

Removes deletes `keys` in the cache.

### Update

```go
func Update(ctx context.Context, key any, value any) (oldValue *gvar.Var, exist bool, err error)
```

Update updates the value of `key` without changing its expiration and returns the old value.
The returned value `exist` is false if the `key` does not exist in the cache.

It deletes the `key` if given `value` is nil.
It does nothing if `key` does not exist in the cache.

### UpdateExpire

```go
func UpdateExpire(ctx context.Context, key any, duration time.Duration) (oldDuration time.Duration, err error)
```

UpdateExpire updates the expiration of `key` and returns the old expiration duration value.

It returns -1 and does nothing if the `key` does not exist in the cache.
It deletes the `key` if `duration` < 0.

### Size

```go
func Size(ctx context.Context) (int, error)
```

Size returns the number of items in the cache.

### Data

```go
func Data(ctx context.Context) (map[any]any, error)
```

Data returns a copy of all key-value pairs in the cache as map type.
Note that this function may lead lots of memory usage, you can implement this function
if necessary.

### Keys

```go
func Keys(ctx context.Context) ([]any, error)
```

Keys returns all keys in the cache as slice.

### KeyStrings

```go
func KeyStrings(ctx context.Context) ([]string, error)
```

KeyStrings returns all keys in the cache as string slice.

### Values

```go
func Values(ctx context.Context) ([]any, error)
```

Values returns all values in the cache as slice.

### MustGet

```go
func MustGet(ctx context.Context, key any) *gvar.Var
```

MustGet acts like Get, but it panics if any error occurs.

### MustGetOrSet

```go
func MustGetOrSet(ctx context.Context, key any, value any, duration time.Duration) *gvar.Var
```

MustGetOrSet acts like GetOrSet, but it panics if any error occurs.

### MustGetOrSetFunc

```go
func MustGetOrSetFunc(ctx context.Context, key any, f Func, duration time.Duration) *gvar.Var
```

MustGetOrSetFunc acts like GetOrSetFunc, but it panics if any error occurs.

### MustGetOrSetFuncLock

```go
func MustGetOrSetFuncLock(ctx context.Context, key any, f Func, duration time.Duration) *gvar.Var
```

MustGetOrSetFuncLock acts like GetOrSetFuncLock, but it panics if any error occurs.

### MustContains

```go
func MustContains(ctx context.Context, key any) bool
```

MustContains acts like Contains, but it panics if any error occurs.

### MustGetExpire

```go
func MustGetExpire(ctx context.Context, key any) time.Duration
```

MustGetExpire acts like GetExpire, but it panics if any error occurs.

### MustSize

```go
func MustSize(ctx context.Context) int
```

MustSize acts like Size, but it panics if any error occurs.

### MustData

```go
func MustData(ctx context.Context) map[any]any
```

MustData acts like Data, but it panics if any error occurs.

### MustKeys

```go
func MustKeys(ctx context.Context) []any
```

MustKeys acts like Keys, but it panics if any error occurs.

### MustKeyStrings

```go
func MustKeyStrings(ctx context.Context) []string
```

MustKeyStrings acts like KeyStrings, but it panics if any error occurs.

### MustValues

```go
func MustValues(ctx context.Context) []any
```

MustValues acts like Values, but it panics if any error occurs.

### NewAdapterMemory

```go
func NewAdapterMemory() *AdapterMemory
```

NewAdapterMemory creates and returns a new adapter_memory cache object.

### NewAdapterMemoryLru

```go
func NewAdapterMemoryLru(cap int) *AdapterMemory
```

NewAdapterMemoryLru creates and returns a new adapter_memory cache object with LRU.

### NewAdapterRedis

```go
func NewAdapterRedis(redis *gredis.Redis) *AdapterRedis
```

NewAdapterRedis creates and returns a new Redis cache adapter.

### New

```go
func New(lruCap ...int) *Cache
```

New creates and returns a new cache object using default memory adapter.
Note that the LRU feature is only available using memory adapter.

### NewWithAdapter

```go
func NewWithAdapter(adapter Adapter) *Cache
```

NewWithAdapter creates and returns a Cache object with given Adapter implements.

## 内部依赖

- `container/glist`
- `container/gmap`
- `container/gset`
- `container/gtype`
- `container/gvar`
- `database/gredis`
- `os/gtime`
- `os/gtimer`
- `util/gconv`

