# redis

> Package: `github.com/gogf/gf/contrib/nosql/redis/v2`

```go
import "github.com/gogf/gf/contrib/nosql/redis/v2"
```

## 概述

Package redis provides gredis.Adapter implements using go-redis.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/nosql/redis/v2`

## 源文件

- `redis.go`
- `redis_conn.go`
- `redis_func.go`
- `redis_group_generic.go`
- `redis_group_hash.go`
- `redis_group_list.go`
- `redis_group_pubsub.go`
- `redis_group_script.go`
- `redis_group_set.go`
- `redis_group_sorted_set.go`
- `redis_group_string.go`
- `redis_operation.go`

## 类型定义

### Redis

**类型**: struct

Redis is an implement of Adapter using go-redis.


**方法**:

- `func (r *Redis) GroupGeneric() gredis.IGroupGeneric`
  GroupGeneric creates and returns GroupGeneric.
- `func (r *Redis) GroupHash() gredis.IGroupHash`
  GroupHash creates and returns a redis group object for hash operations.
- `func (r *Redis) GroupList() gredis.IGroupList`
  GroupList creates and returns a redis group object for list operations.
- `func (r *Redis) GroupPubSub() gredis.IGroupPubSub`
  GroupPubSub creates and returns GroupPubSub.
- `func (r *Redis) GroupScript() gredis.IGroupScript`
  GroupScript creates and returns GroupScript.
- `func (r *Redis) GroupSet() gredis.IGroupSet`
  GroupSet creates and returns GroupSet.
- `func (r *Redis) GroupSortedSet() gredis.IGroupSortedSet`
  GroupSortedSet creates and returns GroupSortedSet.
- `func (r *Redis) GroupString() gredis.IGroupString`
  GroupString is the redis group object for string operations.
- `func (r *Redis) Do(ctx context.Context, command string, args ...any) (*gvar.Var, error)`
  Do send a command to the server and returns the received reply.
- `func (r *Redis) Close(ctx context.Context) err error`
  Close closes the redis connection pool, which will release all connections re...
- `func (r *Redis) Conn(ctx context.Context) (gredis.Conn, error)`
  Conn retrieves and returns a connection object for continuous operations.
- `func (r *Redis) Client() gredis.RedisRawClient`
  Client returns the underlying redis client instance.

### Conn

**类型**: struct

Conn manages the connection operations.


**方法**:

- `func (c *Conn) Do(ctx context.Context, command string, args ...any) (reply *gvar.Var, err error)`
  Do send a command to the server and returns the received reply.
- `func (c *Conn) Receive(ctx context.Context) (*gvar.Var, error)`
  Receive receives a single reply as gvar.Var from the Redis server.
- `func (c *Conn) Close(ctx context.Context) err error`
  Close closes current PubSub or puts the connection back to connection pool.
- `func (c *Conn) Subscribe(ctx context.Context, channel string, channels ...string) ([]*gredis.Subscription, error)`
  Subscribe subscribes the client to the specified channels.
- `func (c *Conn) PSubscribe(ctx context.Context, pattern string, patterns ...string) ([]*gredis.Subscription, error)`
  PSubscribe subscribes the client to the given patterns.
- `func (c *Conn) ReceiveMessage(ctx context.Context) (*gredis.Message, error)`
  ReceiveMessage receives a single message of subscription from the Redis server.

### GroupGeneric

**类型**: struct

GroupGeneric provides generic functions of redis.


**方法**:

- `func (r GroupGeneric) Copy(ctx context.Context, source string, destination string, option ...gredis.CopyOption) (int64, error)`
  Copy copies the value stored at the source key to the destination key.
- `func (r GroupGeneric) Exists(ctx context.Context, keys ...string) (int64, error)`
  Exists returns if key exists.
- `func (r GroupGeneric) Type(ctx context.Context, key string) (string, error)`
  Type returns the string representation of the type of the value stored at key.
- `func (r GroupGeneric) Unlink(ctx context.Context, keys ...string) (int64, error)`
  Unlink is very similar to DEL: it removes the specified keys. Just like DEL a...
- `func (r GroupGeneric) Rename(ctx context.Context, key string, newKey string) error`
  Rename renames key to newKey. It returns an error when key does not exist.
- `func (r GroupGeneric) RenameNX(ctx context.Context, key string, newKey string) (int64, error)`
  RenameNX renames key to newKey if newKey does not yet exist.
- `func (r GroupGeneric) Move(ctx context.Context, key string, db int) (int64, error)`
  Move moves key from the currently selected database (see SELECT) to the speci...
- `func (r GroupGeneric) Del(ctx context.Context, keys ...string) (int64, error)`
  Del removes the specified keys.
- `func (r GroupGeneric) RandomKey(ctx context.Context) (string, error)`
  RandomKey return a random key from the currently selected database.
- `func (r GroupGeneric) DBSize(ctx context.Context) (int64, error)`
  DBSize return the number of keys in the currently-selected database.
- `func (r GroupGeneric) Keys(ctx context.Context, pattern string) ([]string, error)`
  Keys return all keys matching pattern.
- `func (r GroupGeneric) Scan(ctx context.Context, cursor uint64, option ...gredis.ScanOption) (uint64, []string, error)`
  Scan executes a single iteration of the SCAN command, returning a subset of k...
- `func (r GroupGeneric) FlushDB(ctx context.Context, option ...gredis.FlushOp) error`
  FlushDB delete all the keys of the currently selected DB. This command never ...
- `func (r GroupGeneric) FlushAll(ctx context.Context, option ...gredis.FlushOp) error`
  FlushAll delete all the keys of all the existing databases, not just the curr...
- `func (r GroupGeneric) Expire(ctx context.Context, key string, seconds int64, option ...gredis.ExpireOption) (int64, error)`
  Expire sets a timeout on key.
- `func (r GroupGeneric) ExpireAt(ctx context.Context, key string, time time.Time, option ...gredis.ExpireOption) (int64, error)`
  ExpireAt has the same effect and semantic as EXPIRE, but instead of specifyin...
- `func (r GroupGeneric) ExpireTime(ctx context.Context, key string) (*gvar.Var, error)`
  ExpireTime returns the absolute time at which the given key will expire.
- `func (r GroupGeneric) TTL(ctx context.Context, key string) (int64, error)`
  TTL returns the remaining time to live of a key that has a timeout.
- `func (r GroupGeneric) Persist(ctx context.Context, key string) (int64, error)`
  Persist removes the existing timeout on key, turning the key from volatile (a...
- `func (r GroupGeneric) PExpire(ctx context.Context, key string, milliseconds int64, option ...gredis.ExpireOption) (int64, error)`
  PExpire works exactly like EXPIRE but the time to live of the key is specifie...
- `func (r GroupGeneric) PExpireAt(ctx context.Context, key string, time time.Time, option ...gredis.ExpireOption) (int64, error)`
  PExpireAt has the same effect and semantic as ExpireAt, but the Unix time at ...
- `func (r GroupGeneric) PExpireTime(ctx context.Context, key string) (*gvar.Var, error)`
  PExpireTime returns the expiration time of given `key`.
- `func (r GroupGeneric) PTTL(ctx context.Context, key string) (int64, error)`
  PTTL like TTL this command returns the remaining time to live of a key that h...

### GroupHash

**类型**: struct

GroupHash is the redis group object for hash operations.


**方法**:

- `func (r GroupHash) HSet(ctx context.Context, key string, fields map[string]any) (int64, error)`
  HSet sets field in the hash stored at key to value.
- `func (r GroupHash) HSetNX(ctx context.Context, key string, field string, value any) (int64, error)`
  HSetNX sets field in the hash stored at key to value, only if field does not ...
- `func (r GroupHash) HGet(ctx context.Context, key string, field string) (*gvar.Var, error)`
  HGet returns the value associated with field in the hash stored at key.
- `func (r GroupHash) HStrLen(ctx context.Context, key string, field string) (int64, error)`
  HStrLen Returns the string length of the value associated with field in the h...
- `func (r GroupHash) HExists(ctx context.Context, key string, field string) (int64, error)`
  HExists returns if field is an existing field in the hash stored at key.
- `func (r GroupHash) HDel(ctx context.Context, key string, fields ...string) (int64, error)`
  HDel removes the specified fields from the hash stored at key.
- `func (r GroupHash) HLen(ctx context.Context, key string) (int64, error)`
  HLen returns the number of fields contained in the hash stored at key.
- `func (r GroupHash) HIncrBy(ctx context.Context, key string, field string, increment int64) (int64, error)`
  HIncrBy increments the number stored at field in the hash stored at key by in...
- `func (r GroupHash) HIncrByFloat(ctx context.Context, key string, field string, increment float64) (float64, error)`
  HIncrByFloat increments the specified field of a hash stored at key, and repr...
- `func (r GroupHash) HMSet(ctx context.Context, key string, fields map[string]any) error`
  HMSet sets the specified fields to their respective values in the hash stored...
- `func (r GroupHash) HMGet(ctx context.Context, key string, fields ...string) (gvar.Vars, error)`
  HMGet return  the values associated with the specified fields in the hash sto...
- `func (r GroupHash) HKeys(ctx context.Context, key string) ([]string, error)`
  HKeys returns all field names in the hash stored at key.
- `func (r GroupHash) HVals(ctx context.Context, key string) (gvar.Vars, error)`
  HVals return all values in the hash stored at key.
- `func (r GroupHash) HGetAll(ctx context.Context, key string) (*gvar.Var, error)`
  HGetAll returns all fields and values of the hash stored at key.

### GroupList

**类型**: struct

GroupList is the redis group list object.


**方法**:

- `func (r GroupList) LPush(ctx context.Context, key string, values ...any) (int64, error)`
  LPush inserts all the specified values at the head of the list stored at key
- `func (r GroupList) LPushX(ctx context.Context, key string, element any, elements ...any) (int64, error)`
  LPushX insert value at the head of the list stored at key, only if key exists...
- `func (r GroupList) RPush(ctx context.Context, key string, values ...any) (int64, error)`
  RPush inserts all the specified values at the tail of the list stored at key.
- `func (r GroupList) RPushX(ctx context.Context, key string, value any) (int64, error)`
  RPushX inserts value at the tail of the list stored at key, only if key exist...
- `func (r GroupList) LPop(ctx context.Context, key string, count ...int) (*gvar.Var, error)`
  LPop remove and returns the first element of the list stored at key.
- `func (r GroupList) RPop(ctx context.Context, key string, count ...int) (*gvar.Var, error)`
  RPop remove and returns the last element of the list stored at key.
- `func (r GroupList) LRem(ctx context.Context, key string, count int64, value any) (int64, error)`
  LRem removes the first count occurrences of elements equal to value from the ...
- `func (r GroupList) LLen(ctx context.Context, key string) (int64, error)`
  LLen returns the length of the list stored at key.
- `func (r GroupList) LIndex(ctx context.Context, key string, index int64) (*gvar.Var, error)`
  LIndex return the element at index in the list stored at key.
- `func (r GroupList) LInsert(ctx context.Context, key string, op gredis.LInsertOp, pivot any, value any) (int64, error)`
  LInsert inserts element in the list stored at key either before or after the ...
- `func (r GroupList) LSet(ctx context.Context, key string, index int64, value any) (*gvar.Var, error)`
  LSet sets the list element at index to element.
- `func (r GroupList) LRange(ctx context.Context, key string, start int64, stop int64) (gvar.Vars, error)`
  LRange returns the specified elements of the list stored at key.
- `func (r GroupList) LTrim(ctx context.Context, key string, start int64, stop int64) error`
  LTrim trims an existing list so that it will contain only the specified range...
- `func (r GroupList) BLPop(ctx context.Context, timeout int64, keys ...string) (gvar.Vars, error)`
  BLPop is a blocking list pop primitive.
- `func (r GroupList) BRPop(ctx context.Context, timeout int64, keys ...string) (gvar.Vars, error)`
  BRPop is a blocking list pop primitive.
- `func (r GroupList) RPopLPush(ctx context.Context, source string, destination string) (*gvar.Var, error)`
  RPopLPush remove the last element in list source, appends it to the front of ...
- `func (r GroupList) BRPopLPush(ctx context.Context, source string, destination string, timeout int64) (*gvar.Var, error)`
  BRPopLPush is the blocking variant of RPopLPush.

### GroupPubSub

**类型**: struct

GroupPubSub provides pub/sub functions for redis.


**方法**:

- `func (r GroupPubSub) Publish(ctx context.Context, channel string, message any) (int64, error)`
  Publish posts a message to the given channel.
- `func (r GroupPubSub) Subscribe(ctx context.Context, channel string, channels ...string) (gredis.Conn, []*gredis.Subscription, error)`
  Subscribe subscribes the client to the specified channels.
- `func (r GroupPubSub) PSubscribe(ctx context.Context, pattern string, patterns ...string) (gredis.Conn, []*gredis.Subscription, error)`
  PSubscribe subscribes the client to the given patterns.

### GroupScript

**类型**: struct

GroupScript provides script functions for redis.


**方法**:

- `func (r GroupScript) Eval(ctx context.Context, script string, numKeys int64, keys []string, args []any) (*gvar.Var, error)`
  Eval invokes the execution of a server-side Lua script.
- `func (r GroupScript) EvalSha(ctx context.Context, sha1 string, numKeys int64, keys []string, args []any) (*gvar.Var, error)`
  EvalSha evaluates a script from the server's cache by its SHA1 digest.
- `func (r GroupScript) ScriptLoad(ctx context.Context, script string) (string, error)`
  ScriptLoad loads a script into the scripts cache, without executing it.
- `func (r GroupScript) ScriptExists(ctx context.Context, sha1 string, sha1s ...string) (map[string]bool, error)`
  ScriptExists returns information about the existence of the scripts in the sc...
- `func (r GroupScript) ScriptFlush(ctx context.Context, option ...gredis.ScriptFlushOption) error`
  ScriptFlush flush the Lua scripts cache.
- `func (r GroupScript) ScriptKill(ctx context.Context) error`
  ScriptKill kills the currently executing EVAL script, assuming no write opera...

### GroupSet

**类型**: struct

GroupSet provides set functions for redis.


**方法**:

- `func (r GroupSet) SAdd(ctx context.Context, key string, member any, members ...any) (int64, error)`
  SAdd adds the specified members to the set stored at key.
- `func (r GroupSet) SIsMember(ctx context.Context, key string, member any) (int64, error)`
  SIsMember returns if member is a member of the set stored at key.
- `func (r GroupSet) SPop(ctx context.Context, key string, count ...int) (*gvar.Var, error)`
  SPop removes and returns one or more random members from the set value store ...
- `func (r GroupSet) SRandMember(ctx context.Context, key string, count ...int) (*gvar.Var, error)`
  SRandMember called with just the key argument, return a random element from t...
- `func (r GroupSet) SRem(ctx context.Context, key string, member any, members ...any) (int64, error)`
  SRem removes the specified members from the set stored at key.
- `func (r GroupSet) SMove(ctx context.Context, source string, destination string, member any) (int64, error)`
  SMove moves member from the set at source to the set at destination.
- `func (r GroupSet) SCard(ctx context.Context, key string) (int64, error)`
  SCard returns the set cardinality (number of elements) of the set stored at key.
- `func (r GroupSet) SMembers(ctx context.Context, key string) (gvar.Vars, error)`
  SMembers returns all the members of the set value stored at key.
- `func (r GroupSet) SMIsMember(ctx context.Context, key any, member any, members ...any) ([]int, error)`
  SMIsMember returns whether each member is a member of the set stored at key.
- `func (r GroupSet) SInter(ctx context.Context, key string, keys ...string) (gvar.Vars, error)`
  SInter returns the members of the set resulting from the intersection of all ...
- `func (r GroupSet) SInterStore(ctx context.Context, destination string, key string, keys ...string) (int64, error)`
  SInterStore is equal to SInter, but instead of returning the resulting set, i...
- `func (r GroupSet) SUnion(ctx context.Context, key string, keys ...string) (gvar.Vars, error)`
  SUnion returns the members of the set resulting from the union of all the giv...
- `func (r GroupSet) SUnionStore(ctx context.Context, destination string, key string, keys ...string) (int64, error)`
  SUnionStore is equal to SUnion, but instead of returning the resulting set, i...
- `func (r GroupSet) SDiff(ctx context.Context, key string, keys ...string) (gvar.Vars, error)`
  SDiff returns the members of the set resulting from the difference between th...
- `func (r GroupSet) SDiffStore(ctx context.Context, destination string, key string, keys ...string) (int64, error)`
  SDiffStore is equal to SDiff, but instead of returning the resulting set, it ...

### GroupSortedSet

**类型**: struct

GroupSortedSet provides sorted set functions for redis.


**方法**:

- `func (r GroupSortedSet) ZAdd(ctx context.Context, key string, option *gredis.ZAddOption, member gredis.ZAddMember, members ...gredis.ZAddMember) (*gvar.Var, error)`
  ZAdd adds all the specified members with the specified scores to the sorted s...
- `func (r GroupSortedSet) ZScore(ctx context.Context, key string, member any) (float64, error)`
  ZScore Returns the score of member in the sorted set at key.
- `func (r GroupSortedSet) ZIncrBy(ctx context.Context, key string, increment float64, member any) (float64, error)`
  ZIncrBy increments the score of member in the sorted set stored at key by inc...
- `func (r GroupSortedSet) ZCard(ctx context.Context, key string) (int64, error)`
  ZCard returns the sorted set cardinality (number of elements) of the sorted s...
- `func (r GroupSortedSet) ZCount(ctx context.Context, key string, min string, max string) (int64, error)`
  ZCount returns the number of elements in the sorted set at key with a score b...
- `func (r GroupSortedSet) ZRange(ctx context.Context, key string, start int64, stop int64, option ...gredis.ZRangeOption) (gvar.Vars, error)`
  ZRange return the specified range of elements in the sorted set stored at <key>.
- `func (r GroupSortedSet) ZRevRange(ctx context.Context, key string, start int64, stop int64, option ...gredis.ZRevRangeOption) (*gvar.Var, error)`
  ZRevRange returns the specified range of elements in the sorted set stored at...
- `func (r GroupSortedSet) ZRank(ctx context.Context, key string, member any) (int64, error)`
  ZRank returns the rank of member in the sorted set stored at key, with the sc...
- `func (r GroupSortedSet) ZRevRank(ctx context.Context, key string, member any) (int64, error)`
  ZRevRank returns the rank of member in the sorted set stored at key, with the...
- `func (r GroupSortedSet) ZRem(ctx context.Context, key string, member any, members ...any) (int64, error)`
  ZRem remove the specified members from the sorted set stored at key.
- `func (r GroupSortedSet) ZRemRangeByRank(ctx context.Context, key string, start int64, stop int64) (int64, error)`
  ZRemRangeByRank removes all elements in the sorted set stored at key with ran...
- `func (r GroupSortedSet) ZRemRangeByScore(ctx context.Context, key string, min string, max string) (int64, error)`
  ZRemRangeByScore removes all elements in the sorted set stored at key with a ...
- `func (r GroupSortedSet) ZRemRangeByLex(ctx context.Context, key string, min string, max string) (int64, error)`
  ZRemRangeByLex removes all elements in the sorted set stored at key between the
- `func (r GroupSortedSet) ZLexCount(ctx context.Context, key string, min string, max string) (int64, error)`
  ZLexCount all the elements in a sorted set are inserted with the same score,

### GroupString

**类型**: struct

GroupString is the function group manager for string operations.


**方法**:

- `func (r GroupString) Set(ctx context.Context, key string, value any, option ...gredis.SetOption) (*gvar.Var, error)`
  Set key to hold the string value. If key already holds a value, it is overwri...
- `func (r GroupString) SetNX(ctx context.Context, key string, value any) (bool, error)`
  SetNX sets key to hold string value if key does not exist.
- `func (r GroupString) SetEX(ctx context.Context, key string, value any, ttlInSeconds int64) error`
  SetEX sets key to hold the string value and set key to timeout after a given ...
- `func (r GroupString) Get(ctx context.Context, key string) (*gvar.Var, error)`
  Get the value of key. If the key does not exist the special value nil is retu...
- `func (r GroupString) GetDel(ctx context.Context, key string) (*gvar.Var, error)`
  GetDel gets the value of key and delete the key.
- `func (r GroupString) GetEX(ctx context.Context, key string, option ...gredis.GetEXOption) (*gvar.Var, error)`
  GetEX is similar to GET, but is a write command with additional options.
- `func (r GroupString) GetSet(ctx context.Context, key string, value any) (*gvar.Var, error)`
  GetSet atomically sets key to value and returns the old value stored at key.
- `func (r GroupString) StrLen(ctx context.Context, key string) (int64, error)`
  StrLen returns the length of the string value stored at key.
- `func (r GroupString) Append(ctx context.Context, key string, value string) (int64, error)`
  Append appends the value at the end of the string, if key already exists and ...
- `func (r GroupString) SetRange(ctx context.Context, key string, offset int64, value string) (int64, error)`
  SetRange overwrites part of the string stored at key, starting at the specifi...
- `func (r GroupString) GetRange(ctx context.Context, key string, start int64, end int64) (string, error)`
  GetRange returns the substring of the string value stored at key,
- `func (r GroupString) Incr(ctx context.Context, key string) (int64, error)`
  Incr increments the number stored at key by one.
- `func (r GroupString) IncrBy(ctx context.Context, key string, increment int64) (int64, error)`
  IncrBy increments the number stored at key by increment. If the key does not ...
- `func (r GroupString) IncrByFloat(ctx context.Context, key string, increment float64) (float64, error)`
  IncrByFloat increments the string representing a floating point number stored...
- `func (r GroupString) Decr(ctx context.Context, key string) (int64, error)`
  Decr decrements the number stored at key by one.
- `func (r GroupString) DecrBy(ctx context.Context, key string, decrement int64) (int64, error)`
  DecrBy decrements the number stored at key by decrement.
- `func (r GroupString) MSet(ctx context.Context, keyValueMap map[string]any) error`
  MSet sets the given keys to their respective values.
- `func (r GroupString) MSetNX(ctx context.Context, keyValueMap map[string]any) (bool, error)`
  MSetNX sets the given keys to their respective values.
- `func (r GroupString) MGet(ctx context.Context, keys ...string) (map[string]*gvar.Var, error)`
  MGet returns the values of all specified keys.

## 函数

### New

```go
func New(config *gredis.Config) *Redis
```

New creates and returns a redis adapter using go-redis.

## 内部依赖

- `github.com/gogf/gf/v2`
- `container/gvar`
- `database/gredis`
- `encoding/gjson`
- `errors/gerror`
- `net/gtrace`
- `os/gstructs`
- `os/gtime`
- `text/gstr`
- `util/gconv`
- `util/gutil`

## 外部依赖

- `github.com/redis/go-redis/v9`
- `go.opentelemetry.io/otel`
- `go.opentelemetry.io/otel/attribute`
- `go.opentelemetry.io/otel/codes`
- `go.opentelemetry.io/otel/trace`

