# gcfg

> Package: `github.com/gogf/gf/v2/os/gcfg`

```go
import "github.com/gogf/gf/v2/os/gcfg"
```

## 概述

Package gcfg provides reading, caching and managing for configuration.

## 源文件

- `gcfg.go`
- `gcfg_adaper.go`
- `gcfg_adapter_content.go`
- `gcfg_adapter_content_ctx.go`
- `gcfg_adapter_file.go`
- `gcfg_adapter_file_content.go`
- `gcfg_adapter_file_ctx.go`
- `gcfg_adapter_file_path.go`
- `gcfg_ctx_keys.go`
- `gcfg_watcher_registry.go`

## 类型定义

### Config

**类型**: struct

Config is the configuration management object.


**方法**:

- `func (c *Config) SetAdapter(adapter Adapter)`
  SetAdapter sets the adapter of the current Config object.
- `func (c *Config) GetAdapter() Adapter`
  GetAdapter returns the adapter of the current Config object.
- `func (c *Config) Available(ctx context.Context, resource ...string) ok bool`
  Available checks and returns the configuration service is available.
- `func (c *Config) Get(ctx context.Context, pattern string, def ...any) (*gvar.Var, error)`
  Get retrieves and returns value by specified `pattern`.
- `func (c *Config) GetWithEnv(ctx context.Context, pattern string, def ...any) (*gvar.Var, error)`
  GetWithEnv returns the configuration value specified by pattern `pattern`.
- `func (c *Config) GetWithCmd(ctx context.Context, pattern string, def ...any) (*gvar.Var, error)`
  GetWithCmd returns the configuration value specified by pattern `pattern`.
- `func (c *Config) Data(ctx context.Context) (data map[string]any, err error)`
  Data retrieves and returns all configuration data as map type.
- `func (c *Config) MustGet(ctx context.Context, pattern string, def ...any) *gvar.Var`
  MustGet acts as function Get, but it panics if error occurs.
- `func (c *Config) MustGetWithEnv(ctx context.Context, pattern string, def ...any) *gvar.Var`
  MustGetWithEnv acts as function GetWithEnv, but it panics if error occurs.
- `func (c *Config) MustGetWithCmd(ctx context.Context, pattern string, def ...any) *gvar.Var`
  MustGetWithCmd acts as function GetWithCmd, but it panics if error occurs.
- `func (c *Config) MustData(ctx context.Context) map[string]any`
  MustData acts as function Data, but it panics if error occurs.

### Adapter

**类型**: interface

Adapter is the interface for configuration retrieving.


### WatcherAdapter

**类型**: interface

WatcherAdapter is the interface for configuration watcher.


### AdapterContent

**类型**: struct

AdapterContent implements interface Adapter using content.
The configuration content supports the coding types as package `gjson`.


**方法**:

- `func (a *AdapterContent) SetContent(content string) error`
  SetContent sets customized configuration content for specified `file`.
- `func (a *AdapterContent) Available(ctx context.Context, resource ...string) ok bool`
  Available checks and returns the backend configuration service is available.
- `func (a *AdapterContent) Get(ctx context.Context, pattern string) (value any, err error)`
  Get retrieves and returns value by specified `pattern` in current resource.
- `func (a *AdapterContent) Data(ctx context.Context) (data map[string]any, err error)`
  Data retrieves and returns all configuration data in current resource as map.
- `func (a *AdapterContent) AddWatcher(name string, fn func)`
  AddWatcher adds a watcher for the specified configuration file.
- `func (a *AdapterContent) RemoveWatcher(name string)`
  RemoveWatcher removes the watcher for the specified configuration file.
- `func (a *AdapterContent) GetWatcherNames() []string`
  GetWatcherNames returns all watcher names.

### AdapterContentCtx

**类型**: struct

AdapterContentCtx is the context for AdapterContent.


**方法**:

- `func (a *AdapterContentCtx) WithOperation(operation OperationType) *AdapterContentCtx`
  WithOperation sets the operation in the context and returns the updated Adapt...
- `func (a *AdapterContentCtx) WithContent(content string) *AdapterContentCtx`
  WithContent sets the content in the context and returns the updated AdapterCo...
- `func (a *AdapterContentCtx) GetOperation() OperationType`
  GetOperation retrieves the operation from the context.
- `func (a *AdapterContentCtx) GetContent() string`
  GetContent retrieves the content from the context.

### AdapterFile

**类型**: struct

AdapterFile implements interface Adapter using file.


**方法**:

- `func (a *AdapterFile) SetViolenceCheck(check bool)`
  SetViolenceCheck sets whether to perform hierarchical conflict checking.
- `func (a *AdapterFile) SetFileName(fileNameOrPath string)`
  SetFileName sets the default configuration file name.
- `func (a *AdapterFile) GetFileName() string`
  GetFileName returns the default configuration file name.
- `func (a *AdapterFile) Get(ctx context.Context, pattern string) (value any, err error)`
  Get retrieves and returns value by specified `pattern`.
- `func (a *AdapterFile) Set(pattern string, value any) error`
  Set sets value with specified `pattern`.
- `func (a *AdapterFile) Data(ctx context.Context) (data map[string]any, err error)`
  Data retrieves and returns all configuration data as map type.
- `func (a *AdapterFile) MustGet(ctx context.Context, pattern string) *gvar.Var`
  MustGet acts as a function, but it panics if error occurs.
- `func (a *AdapterFile) Clear()`
  Clear removes all parsed configuration files content cache,
- `func (a *AdapterFile) Dump()`
  Dump prints current JSON object with more manually readable.
- `func (a *AdapterFile) Available(ctx context.Context, fileName ...string) bool`
  Available checks and returns whether configuration of given `file` is available.
- `func (a *AdapterFile) AddWatcher(name string, fn func)`
  AddWatcher adds a watcher for the specified configuration file.
- `func (a *AdapterFile) RemoveWatcher(name string)`
  RemoveWatcher removes the watcher for the specified configuration file.
- `func (a *AdapterFile) GetWatcherNames() []string`
  GetWatcherNames returns all watcher names.
- `func (a *AdapterFile) SetContent(content string, fileNameOrPath ...string)`
  SetContent sets customized configuration content for specified `file`.
- `func (a *AdapterFile) GetContent(fileNameOrPath ...string) string`
  GetContent returns customized configuration content for specified `file`.
- `func (a *AdapterFile) RemoveContent(fileNameOrPath ...string)`
  RemoveContent removes the global configuration with specified `file`.
- `func (a *AdapterFile) ClearContent()`
  ClearContent removes all global configuration contents.
- `func (a *AdapterFile) SetPath(directoryPath string) err error`
  SetPath sets the configuration `directory` path for file search.
- `func (a *AdapterFile) AddPath(directoryPaths ...string) err error`
  AddPath adds an absolute or relative `directory` path to the search paths.
- `func (a *AdapterFile) GetPaths() []string`
  GetPaths returns the searching directory path array of current configuration ...
- `func (a *AdapterFile) GetFilePath(fileNameOrPath ...string) (filePath string, err error)`
  GetFilePath returns the absolute configuration file path for the given filena...

### AdapterFileCtx

**类型**: struct

AdapterFileCtx is the context for AdapterFile.


**方法**:

- `func (a *AdapterFileCtx) WithFileName(fileName string) *AdapterFileCtx`
  WithFileName sets the file name in the context and returns the updated Adapte...
- `func (a *AdapterFileCtx) WithFilePath(filePath string) *AdapterFileCtx`
  WithFilePath sets the file path in the context and returns the updated Adapte...
- `func (a *AdapterFileCtx) WithFileType(fileType string) *AdapterFileCtx`
  WithFileType sets the file type in the context and returns the updated Adapte...
- `func (a *AdapterFileCtx) WithOperation(operation OperationType) *AdapterFileCtx`
  WithOperation sets the operation in the context and returns the updated Adapt...
- `func (a *AdapterFileCtx) WithKey(key string) *AdapterFileCtx`
  WithKey sets the key in the context and returns the updated AdapterFileCtx.
- `func (a *AdapterFileCtx) WithValue(value any) *AdapterFileCtx`
  WithValue sets the value in the context and returns the updated AdapterFileCtx.
- `func (a *AdapterFileCtx) WithContent(content any) *AdapterFileCtx`
  WithContent sets the content in the context and returns the updated AdapterFi...
- `func (a *AdapterFileCtx) GetFileName() string`
  GetFileName retrieves the file name from the context.
- `func (a *AdapterFileCtx) GetFilePath() string`
  GetFilePath retrieves the file path from the context.
- `func (a *AdapterFileCtx) GetFileType() string`
  GetFileType retrieves the file type from the context.
- `func (a *AdapterFileCtx) GetOperation() OperationType`
  GetOperation retrieves the operation from the context.
- `func (a *AdapterFileCtx) GetKey() string`
  GetKey retrieves the key from the context.
- `func (a *AdapterFileCtx) GetValue() *gvar.Var`
  GetValue retrieves the value from the context.
- `func (a *AdapterFileCtx) GetContent() *gvar.Var`
  GetContent retrieves the set content from the context.

### OperationType

**类型**: type

OperationType defines the type for configuration operation.


### WatcherRegistry

**类型**: struct

WatcherRegistry is a helper type for managing configuration watchers.
It provides a unified implementation of watcher management to avoid code duplication
across different adapter implementations.


**方法**:

- `func (r *WatcherRegistry) Add(name string, fn func)`
  Add adds a watcher with the specified name and callback function.
- `func (r *WatcherRegistry) Remove(name string)`
  Remove removes the watcher with the specified name.
- `func (r *WatcherRegistry) GetNames() []string`
  GetNames returns all watcher names.
- `func (r *WatcherRegistry) Notify(ctx context.Context)`
  Notify notifies all registered watchers by calling their callback functions.

## 函数

### New

```go
func New() (*Config, error)
```

New creates and returns a Config object with default adapter of AdapterFile.

### NewWithAdapter

```go
func NewWithAdapter(adapter Adapter) *Config
```

NewWithAdapter creates and returns a Config object with given adapter.

### Instance

```go
func Instance(name ...string) *Config
```

Instance returns an instance of Config with default settings.
The parameter `name` is the name for the instance. But very note that, if the file "name.toml"
exists in the configuration directory, it then sets it as the default configuration file. The
toml file type is the default configuration file type.

### NewAdapterContent

```go
func NewAdapterContent(content ...string) (*AdapterContent, error)
```

NewAdapterContent returns a new configuration management object using custom content.
The parameter `content` specifies the default configuration content for reading.

### NewAdapterContentCtxWithCtx

```go
func NewAdapterContentCtxWithCtx(ctx context.Context) *AdapterContentCtx
```

NewAdapterContentCtxWithCtx creates and returns a new AdapterContentCtx with the given context.

### NewAdapterContentCtx

```go
func NewAdapterContentCtx(ctx ...context.Context) *AdapterContentCtx
```

NewAdapterContentCtx creates and returns a new AdapterContentCtx.
If ctx is provided, it uses that context, otherwise it creates a background context.

### GetAdapterContentCtx

```go
func GetAdapterContentCtx(ctx context.Context) *AdapterContentCtx
```

GetAdapterContentCtx creates and returns an AdapterContentCtx with the given context.

### NewAdapterFile

```go
func NewAdapterFile(fileNameOrPath ...string) (*AdapterFile, error)
```

NewAdapterFile returns a new configuration management object.
The parameter `file` specifies the default configuration file name for reading.

### NewAdapterFileCtxWithCtx

```go
func NewAdapterFileCtxWithCtx(ctx context.Context) *AdapterFileCtx
```

NewAdapterFileCtxWithCtx creates and returns a new AdapterFileCtx with the given context.

### NewAdapterFileCtx

```go
func NewAdapterFileCtx(ctx ...context.Context) *AdapterFileCtx
```

NewAdapterFileCtx creates and returns a new AdapterFileCtx.
If ctx is provided, it uses that context, otherwise it creates a background context.

### GetAdapterFileCtx

```go
func GetAdapterFileCtx(ctx context.Context) *AdapterFileCtx
```

GetAdapterFileCtx creates and returns an AdapterFileCtx with the given context.

### NewWatcherRegistry

```go
func NewWatcherRegistry() *WatcherRegistry
```

NewWatcherRegistry creates and returns a new WatcherRegistry instance.

## 内部依赖

- `container/garray`
- `container/gmap`
- `container/gtype`
- `container/gvar`
- `encoding/gjson`
- `errors/gcode`
- `errors/gerror`
- `internal/command`
- `internal/intlog`
- `internal/utils`
- `os/gctx`
- `os/genv`
- `os/gfile`
- `os/gfsnotify`
- `os/gres`
- `os/gspath`
- `text/gstr`
- `util/gmode`
- `util/gutil`

