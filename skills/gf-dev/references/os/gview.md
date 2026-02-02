# gview

> Package: `github.com/gogf/gf/v2/os/gview`

```go
import "github.com/gogf/gf/v2/os/gview"
```

## 概述

Package gview implements a template engine based on text/template.

## 源文件

- `gview.go`
- `gview_buildin.go`
- `gview_config.go`
- `gview_error.go`
- `gview_i18n.go`
- `gview_instance.go`
- `gview_parse.go`

## 类型定义

### View

**类型**: struct

View object for template engine.


**方法**:

- `func (view *View) SetConfig(config Config) error`
  SetConfig sets the configuration for view.
- `func (view *View) SetConfigWithMap(m map[string]any) error`
  SetConfigWithMap set configurations with map for the view.
- `func (view *View) SetPath(path string) error`
  SetPath sets the template directory path for template file search.
- `func (view *View) AddPath(path string) error`
  AddPath adds an absolute or relative path to the search paths.
- `func (view *View) Assigns(data Params)`
  Assigns binds multiple global template variables to current view object.
- `func (view *View) Assign(key string, value any)`
  Assign binds a global template variable to current view object.
- `func (view *View) ClearAssigns()`
  ClearAssigns trunk all global template variables assignments.
- `func (view *View) SetDefaultFile(file string)`
  SetDefaultFile sets default template file for parsing.
- `func (view *View) GetDefaultFile() string`
  GetDefaultFile returns default template file for parsing.
- `func (view *View) SetDelimiters(left string, right string)`
  SetDelimiters sets customized delimiters for template parsing.
- `func (view *View) SetAutoEncode(enable bool)`
  SetAutoEncode enables/disables automatically html encoding feature.
- `func (view *View) BindFunc(name string, function any)`
  BindFunc registers customized global template function named `name`
- `func (view *View) BindFuncMap(funcMap FuncMap)`
  BindFuncMap registers customized global template functions by map to current ...
- `func (view *View) SetI18n(manager *gi18n.Manager)`
  SetI18n binds i18n manager to current view engine.
- `func (view *View) Parse(ctx context.Context, file string, params ...Params) (result string, err error)`
  Parse parses given template file `file` with given template variables `params`
- `func (view *View) ParseDefault(ctx context.Context, params ...Params) (result string, err error)`
  ParseDefault parses the default template file with params.
- `func (view *View) ParseContent(ctx context.Context, content string, params ...Params) (string, error)`
  ParseContent parses given template content `content` with template variables ...
- `func (view *View) ParseOption(ctx context.Context, option Option) (result string, err error)`
  ParseOption implements template parsing using Option.
- `func (view *View) ParseWithOptions(ctx context.Context, opts Options) (result string, err error)`
  ParseWithOptions implements template parsing using Option.

### Params

**类型**: type

### FuncMap

**类型**: type

### Config

**类型**: struct

Config is the configuration object for template engine.


### Option

**类型**: type

Option for template parsing.

Deprecated: use Options instead.


### Options

**类型**: struct

Options for template parsing.


## 函数

### ParseContent

```go
func ParseContent(ctx context.Context, content string, params ...Params) (string, error)
```

ParseContent parses the template content directly using the default view object
and returns the parsed content.

### New

```go
func New(path ...string) *View
```

New returns a new view object.
The parameter `path` specifies the template directory path to load template files.

### DefaultConfig

```go
func DefaultConfig() Config
```

DefaultConfig creates and returns a configuration object with default configurations.

### Instance

```go
func Instance(name ...string) *View
```

Instance returns an instance of View with default settings.
The parameter `name` is the name for the instance.

## 内部依赖

- `github.com/gogf/gf/v2`
- `container/garray`
- `container/gmap`
- `encoding/ghash`
- `encoding/ghtml`
- `encoding/gjson`
- `encoding/gurl`
- `errors/gcode`
- `errors/gerror`
- `i18n/gi18n`
- `internal/intlog`
- `os/gcmd`
- `os/gfile`
- `os/gfsnotify`
- `os/glog`
- `os/gmlock`
- `os/gres`
- `os/gspath`
- `os/gtime`
- `text/gstr`
- `util/gconv`
- `util/gmode`
- `util/gutil`

