# gvalid

> Package: `github.com/gogf/gf/v2/util/gvalid`

```go
import "github.com/gogf/gf/v2/util/gvalid"
```

## 概述

Package gvalid implements powerful and useful data/form validation functionality.

## 源文件

- `gvalid.go`
- `gvalid_error.go`
- `gvalid_register.go`
- `gvalid_validator.go`
- `gvalid_validator_check_map.go`
- `gvalid_validator_check_struct.go`
- `gvalid_validator_check_value.go`
- `gvalid_validator_message.go`

## 类型定义

### CustomMsg

**类型**: type

CustomMsg is the custom error message type,
like: map[field] => string|map[rule]string


### Error

**类型**: interface

Error is the validation error for validation result.


### RuleFunc

**类型**: type

RuleFunc is the custom function for data validation.


### RuleFuncInput

**类型**: struct

RuleFuncInput holds the input parameters that passed to custom rule function RuleFunc.


### Validator

**类型**: struct

Validator is the validation manager for chaining operations.


**方法**:

- `func (v *Validator) Run(ctx context.Context) Error`
  Run starts validating the given data with rules and messages.
- `func (v *Validator) Clone() *Validator`
  Clone creates and returns a new Validator, which is a shallow copy of the cur...
- `func (v *Validator) I18n(i18nManager *gi18n.Manager) *Validator`
  I18n sets the i18n manager for the validator.
- `func (v *Validator) Bail() *Validator`
  Bail sets the mark for stopping validation after the first validation error.
- `func (v *Validator) Foreach() *Validator`
  Foreach tells the next validation using current value as an array and validat...
- `func (v *Validator) Ci() *Validator`
  Ci sets the mark for Case-Insensitive for those rules that need value compari...
- `func (v *Validator) Data(data any) *Validator`
  Data is a chaining operation function, which sets validation data for current...
- `func (v *Validator) Assoc(assoc any) *Validator`
  Assoc is a chaining operation function, which sets associated validation data...
- `func (v *Validator) Rules(rules any) *Validator`
  Rules is a chaining operation function, which sets custom validation rules fo...
- `func (v *Validator) Messages(messages any) *Validator`
  Messages is a chaining operation function, which sets custom error messages f...
- `func (v *Validator) RuleFunc(rule string, f RuleFunc) *Validator`
  RuleFunc registers one custom rule function to current Validator.
- `func (v *Validator) RuleFuncMap(m map[string]RuleFunc) *Validator`
  RuleFuncMap registers multiple custom rule functions to current Validator.

## 函数

### ParseTagValue

```go
func ParseTagValue(tag string) (field string, rule string, msg string)
```

ParseTagValue parses one sequence tag to field, rule and error message.
The sequence tag is like: [alias@]rule[...#msg...]

### GetTags

```go
func GetTags() []string
```

GetTags returns the validation tags.

### RegisterRule

```go
func RegisterRule(rule string, f RuleFunc)
```

RegisterRule registers custom validation rule and function for package.

### RegisterRuleByMap

```go
func RegisterRuleByMap(m map[string]RuleFunc)
```

RegisterRuleByMap registers custom validation rules using map for package.

### GetRegisteredRuleMap

```go
func GetRegisteredRuleMap() map[string]RuleFunc
```

GetRegisteredRuleMap returns all the custom registered rules and associated functions.

### DeleteRule

```go
func DeleteRule(rules ...string)
```

DeleteRule deletes custom defined validation one or more rules and associated functions from global package.

### New

```go
func New() *Validator
```

New creates and returns a new Validator.

## 内部依赖

- `container/gset`
- `container/gvar`
- `encoding/gjson`
- `errors/gcode`
- `errors/gerror`
- `i18n/gi18n`
- `internal/empty`
- `internal/intlog`
- `internal/reflection`
- `internal/utils`
- `os/gstructs`
- `text/gregex`
- `text/gstr`
- `util/gconv`
- `util/gmeta`
- `util/gtag`
- `util/gutil`
- `util/gvalid/internal/builtin`

