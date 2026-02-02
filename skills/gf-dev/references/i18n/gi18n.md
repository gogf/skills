# gi18n

> Package: `github.com/gogf/gf/v2/i18n/gi18n`

```go
import "github.com/gogf/gf/v2/i18n/gi18n"
```

## 概述

Package gi18n implements internationalization and localization.

## 源文件

- `gi18n.go`
- `gi18n_ctx.go`
- `gi18n_instance.go`
- `gi18n_manager.go`

## 类型定义

### Manager

**类型**: struct

Manager for i18n contents, it is concurrent safe, supporting hot reload.


**方法**:

- `func (m *Manager) SetPath(path string) error`
  SetPath sets the directory path storing i18n files.
- `func (m *Manager) SetLanguage(language string)`
  SetLanguage sets the language for translator.
- `func (m *Manager) SetDelimiters(left string, right string)`
  SetDelimiters sets the delimiters for translator.
- `func (m *Manager) T(ctx context.Context, content string) string`
  T is alias of Translate for convenience.
- `func (m *Manager) Tf(ctx context.Context, format string, values ...any) string`
  Tf is alias of TranslateFormat for convenience.
- `func (m *Manager) TranslateFormat(ctx context.Context, format string, values ...any) string`
  TranslateFormat translates, formats and returns the `format` with configured ...
- `func (m *Manager) Translate(ctx context.Context, content string) string`
  Translate translates `content` with configured language.
- `func (m *Manager) GetContent(ctx context.Context, key string) string`
  GetContent retrieves and returns the configured content for given key and spe...

### Options

**类型**: struct

Options is used for i18n object configuration.


## 函数

### SetPath

```go
func SetPath(path string) error
```

SetPath sets the directory path storing i18n files.

### SetLanguage

```go
func SetLanguage(language string)
```

SetLanguage sets the language for translator.

### SetDelimiters

```go
func SetDelimiters(left string, right string)
```

SetDelimiters sets the delimiters for translator.

### T

```go
func T(ctx context.Context, content string) string
```

T is alias of Translate for convenience.

### Tf

```go
func Tf(ctx context.Context, format string, values ...any) string
```

Tf is alias of TranslateFormat for convenience.

### TranslateFormat

```go
func TranslateFormat(ctx context.Context, format string, values ...any) string
```

TranslateFormat translates, formats and returns the `format` with configured language
and given `values`.

### Translate

```go
func Translate(ctx context.Context, content string) string
```

Translate translates `content` with configured language and returns the translated content.

### GetContent

```go
func GetContent(ctx context.Context, key string) string
```

GetContent retrieves and returns the configured content for given key and specified language.
It returns an empty string if not found.

### WithLanguage

```go
func WithLanguage(ctx context.Context, language string) context.Context
```

WithLanguage append language setting to the context and returns a new context.

### LanguageFromCtx

```go
func LanguageFromCtx(ctx context.Context) string
```

LanguageFromCtx retrieves and returns language name from context.
It returns an empty string if it is not set previously.

### Instance

```go
func Instance(name ...string) *Manager
```

Instance returns an instance of Resource.
The parameter `name` is the name for the instance.

### New

```go
func New(options ...Options) *Manager
```

New creates and returns a new i18n manager.
The optional parameter `option` specifies the custom options for i18n manager.
It uses a default one if it's not passed.

## 内部依赖

- `container/gmap`
- `encoding/gjson`
- `errors/gcode`
- `errors/gerror`
- `internal/intlog`
- `os/gctx`
- `os/gfile`
- `os/gfsnotify`
- `os/gres`
- `text/gregex`
- `util/gconv`

