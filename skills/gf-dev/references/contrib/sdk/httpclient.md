# httpclient

> Package: `github.com/gogf/gf/contrib/sdk/httpclient/v2`

```go
import "github.com/gogf/gf/contrib/sdk/httpclient/v2"
```

## 概述

Package httpclient provides http client used for SDK.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/sdk/httpclient/v2`

## 源文件

- `httpclient.go`
- `httpclient_config.go`
- `httpclient_handler.go`

## 类型定义

### Client

**类型**: struct

Client is a http client for SDK.


**方法**:

- `func (c *Client) Request(ctx context.Context, req any, res any) error`
  Request sends request to service by struct object `req`, and receives respons...
- `func (c *Client) Get(ctx context.Context, path string, in any, out any) error`
  Get sends a request using GET method.

### Config

**类型**: struct

Config is the configuration struct for SDK client.


### Handler

**类型**: interface

Handler is the interface for http response handling.


### DefaultHandler

**类型**: struct

DefaultHandler handle ghttp.DefaultHandlerResponse of json format.


**方法**:

- `func (h DefaultHandler) HandleResponse(ctx context.Context, res *gclient.Response, out any) error`

## 函数

### New

```go
func New(config Config) *Client
```

New creates and returns a http client for SDK.

### NewDefaultHandler

```go
func NewDefaultHandler(logger *glog.Logger, rawRump bool) *DefaultHandler
```

## 内部依赖

- `encoding/gurl`
- `errors/gcode`
- `errors/gerror`
- `frame/g`
- `net/gclient`
- `net/ghttp`
- `os/glog`
- `text/gregex`
- `text/gstr`
- `util/gconv`
- `util/gmeta`
- `util/gtag`

