# gsession

> Package: `github.com/gogf/gf/v2/os/gsession`

```go
import "github.com/gogf/gf/v2/os/gsession"
```

## 概述

Package gsession implements manager and storage features for sessions.

## 源文件

- `gsession.go`
- `gsession_manager.go`
- `gsession_session.go`
- `gsession_storage.go`
- `gsession_storage_base.go`
- `gsession_storage_file.go`
- `gsession_storage_memory.go`
- `gsession_storage_redis.go`
- `gsession_storage_redis_hashtable.go`

## 类型定义

### Manager

**类型**: struct

Manager for sessions.


**方法**:

- `func (m *Manager) New(ctx context.Context, sessionId ...string) *Session`
  New creates or fetches the session for given session id.
- `func (m *Manager) SetStorage(storage Storage)`
  SetStorage sets the session storage for manager.
- `func (m *Manager) GetStorage() Storage`
  GetStorage returns the session storage of current manager.
- `func (m *Manager) SetTTL(ttl time.Duration)`
  SetTTL the TTL for the session manager.
- `func (m *Manager) GetTTL() time.Duration`
  GetTTL returns the TTL of the session manager.

### Session

**类型**: struct

Session struct for storing single session data, which is bound to a single request.
The Session struct is the interface with user, but the Storage is the underlying adapter designed interface
for functionality implements.


**方法**:

- `func (s *Session) Close() error`
  Close closes current session and updates its ttl in the session manager.
- `func (s *Session) Set(key string, value any) err error`
  Set sets key-value pair to this session.
- `func (s *Session) SetMap(data map[string]any) err error`
  SetMap batch sets the session using map.
- `func (s *Session) Remove(keys ...string) err error`
  Remove removes key along with its value from this session.
- `func (s *Session) RemoveAll() err error`
  RemoveAll deletes all key-value pairs from this session.
- `func (s *Session) Id() (id string, err error)`
  Id returns the session id for this session.
- `func (s *Session) SetId(id string) error`
  SetId sets custom session before session starts.
- `func (s *Session) SetIdFunc(f func) error`
  SetIdFunc sets custom session id creating function before session starts.
- `func (s *Session) Data() (sessionData map[string]any, err error)`
  Data returns all data as map.
- `func (s *Session) Size() (size int, err error)`
  Size returns the size of the session.
- `func (s *Session) Contains(key string) (ok bool, err error)`
  Contains checks whether key exist in the session.
- `func (s *Session) IsDirty() bool`
  IsDirty checks whether there's any data changes in the session.
- `func (s *Session) Get(key string, def ...any) (value *gvar.Var, err error)`
  Get retrieves session value with given key.
- `func (s *Session) MustId() string`
  MustId performs as function Id, but it panics if any error occurs.
- `func (s *Session) MustGet(key string, def ...any) *gvar.Var`
  MustGet performs as function Get, but it panics if any error occurs.
- `func (s *Session) MustSet(key string, value any)`
  MustSet performs as function Set, but it panics if any error occurs.
- `func (s *Session) MustSetMap(data map[string]any)`
  MustSetMap performs as function SetMap, but it panics if any error occurs.
- `func (s *Session) MustContains(key string) bool`
  MustContains performs as function Contains, but it panics if any error occurs.
- `func (s *Session) MustData() map[string]any`
  MustData performs as function Data, but it panics if any error occurs.
- `func (s *Session) MustSize() int`
  MustSize performs as function Size, but it panics if any error occurs.
- `func (s *Session) MustRemove(keys ...string)`
  MustRemove performs as function Remove, but it panics if any error occurs.
- `func (s *Session) RegenerateId(deleteOld bool) (newId string, err error)`
  RegenerateId regenerates a new session id for current session.
- `func (s *Session) MustRegenerateId(deleteOld bool) string`
  MustRegenerateId performs as function RegenerateId, but it panics if any erro...

### Storage

**类型**: interface

Storage is the interface definition for session storage.


### StorageBase

**类型**: struct

StorageBase is a base implement for Session Storage.


**方法**:

- `func (s *StorageBase) New(ctx context.Context, ttl time.Duration) (id string, err error)`
  New creates a session id.
- `func (s *StorageBase) Get(ctx context.Context, sessionId string, key string) (value any, err error)`
  Get retrieves certain session value with given key.
- `func (s *StorageBase) Data(ctx context.Context, sessionId string) (sessionData map[string]any, err error)`
  Data retrieves all key-value pairs as map from storage.
- `func (s *StorageBase) GetSize(ctx context.Context, sessionId string) (size int, err error)`
  GetSize retrieves the size of key-value pairs from storage.
- `func (s *StorageBase) Set(ctx context.Context, sessionId string, key string, value any, ttl time.Duration) error`
  Set sets key-value session pair to the storage.
- `func (s *StorageBase) SetMap(ctx context.Context, sessionId string, mapData map[string]any, ttl time.Duration) error`
  SetMap batch sets key-value session pairs with map to the storage.
- `func (s *StorageBase) Remove(ctx context.Context, sessionId string, key string) error`
  Remove deletes key with its value from storage.
- `func (s *StorageBase) RemoveAll(ctx context.Context, sessionId string) error`
  RemoveAll deletes session from storage.
- `func (s *StorageBase) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*gmap.StrAnyMap, error)`
  GetSession returns the session data as *gmap.StrAnyMap for given session id f...
- `func (s *StorageBase) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error`
  SetSession updates the data map for specified session id.
- `func (s *StorageBase) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error`
  UpdateTTL updates the TTL for specified session id.

### StorageFile

**类型**: struct

StorageFile implements the Session Storage interface with file system.


**方法**:

- `func (s *StorageFile) SetCryptoKey(key []byte)`
  SetCryptoKey sets the crypto key for session storage.
- `func (s *StorageFile) SetCryptoEnabled(enabled bool)`
  SetCryptoEnabled enables/disables the crypto feature for session storage.
- `func (s *StorageFile) RemoveAll(ctx context.Context, sessionId string) error`
  RemoveAll deletes all key-value pairs from storage.
- `func (s *StorageFile) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (sessionData *gmap.StrAnyMap, err error)`
  GetSession returns the session data as *gmap.StrAnyMap for given session id f...
- `func (s *StorageFile) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error`
  SetSession updates the data map for specified session id.
- `func (s *StorageFile) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error`
  UpdateTTL updates the TTL for specified session id.

### StorageMemory

**类型**: struct

StorageMemory implements the Session Storage interface with memory.


**方法**:

- `func (s *StorageMemory) RemoveAll(ctx context.Context, sessionId string) error`
  RemoveAll deletes session from storage.
- `func (s *StorageMemory) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*gmap.StrAnyMap, error)`
  GetSession returns the session data as *gmap.StrAnyMap for given session id f...
- `func (s *StorageMemory) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error`
  SetSession updates the data map for specified session id.
- `func (s *StorageMemory) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error`
  UpdateTTL updates the TTL for specified session id.

### StorageRedis

**类型**: struct

StorageRedis implements the Session Storage interface with redis.


**方法**:

- `func (s *StorageRedis) RemoveAll(ctx context.Context, sessionId string) error`
  RemoveAll deletes all key-value pairs from storage.
- `func (s *StorageRedis) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*gmap.StrAnyMap, error)`
  GetSession returns the session data as *gmap.StrAnyMap for given session id f...
- `func (s *StorageRedis) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error`
  SetSession updates the data map for specified session id.
- `func (s *StorageRedis) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error`
  UpdateTTL updates the TTL for specified session id.

### StorageRedisHashTable

**类型**: struct

StorageRedisHashTable implements the Session Storage interface with redis hash table.


**方法**:

- `func (s *StorageRedisHashTable) Get(ctx context.Context, sessionId string, key string) (value any, err error)`
  Get retrieves session value with given key.
- `func (s *StorageRedisHashTable) Data(ctx context.Context, sessionId string) (data map[string]any, err error)`
  Data retrieves all key-value pairs as map from storage.
- `func (s *StorageRedisHashTable) GetSize(ctx context.Context, sessionId string) (size int, err error)`
  GetSize retrieves the size of key-value pairs from storage.
- `func (s *StorageRedisHashTable) Set(ctx context.Context, sessionId string, key string, value any, ttl time.Duration) error`
  Set sets key-value session pair to the storage.
- `func (s *StorageRedisHashTable) SetMap(ctx context.Context, sessionId string, data map[string]any, ttl time.Duration) error`
  SetMap batch sets key-value session pairs with map to the storage.
- `func (s *StorageRedisHashTable) Remove(ctx context.Context, sessionId string, key string) error`
  Remove deletes key with its value from storage.
- `func (s *StorageRedisHashTable) RemoveAll(ctx context.Context, sessionId string) error`
  RemoveAll deletes all key-value pairs from storage.
- `func (s *StorageRedisHashTable) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*gmap.StrAnyMap, error)`
  GetSession returns the session data as *gmap.StrAnyMap for given session id f...
- `func (s *StorageRedisHashTable) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error`
  SetSession updates the data map for specified session id.
- `func (s *StorageRedisHashTable) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error`
  UpdateTTL updates the TTL for specified session id.

## 函数

### NewSessionId

```go
func NewSessionId() string
```

NewSessionId creates and returns a new and unique session id string,
which is in 32 bytes.

### New

```go
func New(ttl time.Duration, storage ...Storage) *Manager
```

New creates and returns a new session manager.

### NewStorageFile

```go
func NewStorageFile(path string, ttl time.Duration) *StorageFile
```

NewStorageFile creates and returns a file storage object for session.

### NewStorageMemory

```go
func NewStorageMemory() *StorageMemory
```

NewStorageMemory creates and returns a file storage object for session.

### NewStorageRedis

```go
func NewStorageRedis(redis *gredis.Redis, prefix ...string) *StorageRedis
```

NewStorageRedis creates and returns a redis storage object for session.

### NewStorageRedisHashTable

```go
func NewStorageRedisHashTable(redis *gredis.Redis, prefix ...string) *StorageRedisHashTable
```

NewStorageRedisHashTable creates and returns a redis hash table storage object for session.

## 内部依赖

- `container/gmap`
- `container/gset`
- `container/gvar`
- `crypto/gaes`
- `database/gredis`
- `encoding/gbinary`
- `errors/gcode`
- `errors/gerror`
- `internal/intlog`
- `internal/json`
- `os/gcache`
- `os/gfile`
- `os/gtime`
- `os/gtimer`
- `util/guid`

