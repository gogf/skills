# gclient

> Package: `github.com/gogf/gf/v2/net/gclient`

```go
import "github.com/gogf/gf/v2/net/gclient"
```

## 概述

Package gclient provides convenient http client functionalities.

## 源文件

- `gclient.go`
- `gclient_bytes.go`
- `gclient_chain.go`
- `gclient_config.go`
- `gclient_content.go`
- `gclient_discovery.go`
- `gclient_dump.go`
- `gclient_metrics.go`
- `gclient_middleware.go`
- `gclient_observability.go`
- `gclient_request.go`
- `gclient_request_obj.go`
- `gclient_response.go`
- `gclient_tracer.go`
- `gclient_tracer_metrics.go`
- `gclient_tracer_noop.go`
- `gclient_tracer_tracing.go`
- `gclient_var.go`
- `gclient_websocket.go`

## 类型定义

### Client

**类型**: struct

Client is the HTTP client for HTTP request management.


**方法**:

- `func (c *Client) Clone() *Client`
  Clone deeply clones current client and returns a new one.
- `func (c *Client) GetBytes(ctx context.Context, url string, data ...any) []byte`
  GetBytes sends a GET request, retrieves and returns the result content as bytes.
- `func (c *Client) PutBytes(ctx context.Context, url string, data ...any) []byte`
  PutBytes sends a PUT request, retrieves and returns the result content as bytes.
- `func (c *Client) PostBytes(ctx context.Context, url string, data ...any) []byte`
  PostBytes sends a POST request, retrieves and returns the result content as b...
- `func (c *Client) DeleteBytes(ctx context.Context, url string, data ...any) []byte`
  DeleteBytes sends a DELETE request, retrieves and returns the result content ...
- `func (c *Client) HeadBytes(ctx context.Context, url string, data ...any) []byte`
  HeadBytes sends a HEAD request, retrieves and returns the result content as b...
- `func (c *Client) PatchBytes(ctx context.Context, url string, data ...any) []byte`
  PatchBytes sends a PATCH request, retrieves and returns the result content as...
- `func (c *Client) ConnectBytes(ctx context.Context, url string, data ...any) []byte`
  ConnectBytes sends a CONNECT request, retrieves and returns the result conten...
- `func (c *Client) OptionsBytes(ctx context.Context, url string, data ...any) []byte`
  OptionsBytes sends an OPTIONS request, retrieves and returns the result conte...
- `func (c *Client) TraceBytes(ctx context.Context, url string, data ...any) []byte`
  TraceBytes sends a TRACE request, retrieves and returns the result content as...
- `func (c *Client) RequestBytes(ctx context.Context, method string, url string, data ...any) []byte`
  RequestBytes sends request using given HTTP method and data, retrieves return...
- `func (c *Client) Prefix(prefix string) *Client`
  Prefix is a chaining function,
- `func (c *Client) Header(m map[string]string) *Client`
  Header is a chaining function,
- `func (c *Client) HeaderRaw(headers string) *Client`
  HeaderRaw is a chaining function,
- `func (c *Client) Discovery(discovery gsvc.Discovery) *Client`
  Discovery is a chaining function, which sets the discovery for client.
- `func (c *Client) Cookie(m map[string]string) *Client`
  Cookie is a chaining function,
- `func (c *Client) ContentType(contentType string) *Client`
  ContentType is a chaining function,
- `func (c *Client) ContentJson() *Client`
  ContentJson is a chaining function,
- `func (c *Client) ContentXml() *Client`
  ContentXml is a chaining function,
- `func (c *Client) Timeout(t time.Duration) *Client`
  Timeout is a chaining function,
- `func (c *Client) BasicAuth(user string, pass string) *Client`
  BasicAuth is a chaining function,
- `func (c *Client) Retry(retryCount int, retryInterval time.Duration) *Client`
  Retry is a chaining function,
- `func (c *Client) Proxy(proxyURL string) *Client`
  Proxy is a chaining function,
- `func (c *Client) RedirectLimit(redirectLimit int) *Client`
  RedirectLimit is a chaining function,
- `func (c *Client) NoUrlEncode() *Client`
  NoUrlEncode sets the mark that do not encode the parameters before sending re...
- `func (c *Client) SetBrowserMode(enabled bool) *Client`
  SetBrowserMode enables browser mode of the client.
- `func (c *Client) SetHeader(key string, value string) *Client`
  SetHeader sets a custom HTTP header pair for the client.
- `func (c *Client) SetHeaderMap(m map[string]string) *Client`
  SetHeaderMap sets custom HTTP headers with map.
- `func (c *Client) SetAgent(agent string) *Client`
  SetAgent sets the User-Agent header for client.
- `func (c *Client) SetContentType(contentType string) *Client`
  SetContentType sets HTTP content type for the client.
- `func (c *Client) SetHeaderRaw(headers string) *Client`
  SetHeaderRaw sets custom HTTP header using raw string.
- `func (c *Client) SetCookie(key string, value string) *Client`
  SetCookie sets a cookie pair for the client.
- `func (c *Client) SetCookieMap(m map[string]string) *Client`
  SetCookieMap sets cookie items with map.
- `func (c *Client) SetPrefix(prefix string) *Client`
  SetPrefix sets the request server URL prefix.
- `func (c *Client) SetTimeout(t time.Duration) *Client`
  SetTimeout sets the request timeout for the client.
- `func (c *Client) SetBasicAuth(user string, pass string) *Client`
  SetBasicAuth sets HTTP basic authentication information for the client.
- `func (c *Client) SetRetry(retryCount int, retryInterval time.Duration) *Client`
  SetRetry sets retry count and interval.
- `func (c *Client) SetRedirectLimit(redirectLimit int) *Client`
  SetRedirectLimit limits the number of jumps.
- `func (c *Client) SetNoUrlEncode(noUrlEncode bool) *Client`
  SetNoUrlEncode sets the mark that do not encode the parameters before sending...
- `func (c *Client) SetProxy(proxyURL string)`
  SetProxy set proxy for the client.
- `func (c *Client) SetTLSKeyCrt(crtFile string, keyFile string) error`
  SetTLSKeyCrt sets the certificate and key file for TLS configuration of client.
- `func (c *Client) SetTLSConfig(tlsConfig *tls.Config) error`
  SetTLSConfig sets the TLS configuration of client.
- `func (c *Client) SetBuilder(builder gsel.Builder)`
  SetBuilder sets the load balance builder for client.
- `func (c *Client) SetDiscovery(discovery gsvc.Discovery)`
  SetDiscovery sets the load balance builder for client.
- `func (c *Client) GetContent(ctx context.Context, url string, data ...any) string`
  GetContent is a convenience method for sending GET request, which retrieves a...
- `func (c *Client) PutContent(ctx context.Context, url string, data ...any) string`
  PutContent is a convenience method for sending PUT request, which retrieves a...
- `func (c *Client) PostContent(ctx context.Context, url string, data ...any) string`
  PostContent is a convenience method for sending POST request, which retrieves...
- `func (c *Client) DeleteContent(ctx context.Context, url string, data ...any) string`
  DeleteContent is a convenience method for sending DELETE request, which retri...
- `func (c *Client) HeadContent(ctx context.Context, url string, data ...any) string`
  HeadContent is a convenience method for sending HEAD request, which retrieves...
- `func (c *Client) PatchContent(ctx context.Context, url string, data ...any) string`
  PatchContent is a convenience method for sending PATCH request, which retriev...
- `func (c *Client) ConnectContent(ctx context.Context, url string, data ...any) string`
  ConnectContent is a convenience method for sending CONNECT request, which ret...
- `func (c *Client) OptionsContent(ctx context.Context, url string, data ...any) string`
  OptionsContent is a convenience method for sending OPTIONS request, which ret...
- `func (c *Client) TraceContent(ctx context.Context, url string, data ...any) string`
  TraceContent is a convenience method for sending TRACE request, which retriev...
- `func (c *Client) RequestContent(ctx context.Context, method string, url string, data ...any) string`
  RequestContent is a convenience method for sending custom http method request...
- `func (c *Client) Use(handlers ...HandlerFunc) *Client`
  Use adds one or more middleware handlers to client.
- `func (c *Client) Next(req *http.Request) (*Response, error)`
  Next calls the next middleware.
- `func (c *Client) Get(ctx context.Context, url string, data ...any) (*Response, error)`
  Get send GET request and returns the response object.
- `func (c *Client) Put(ctx context.Context, url string, data ...any) (*Response, error)`
  Put send PUT request and returns the response object.
- `func (c *Client) Post(ctx context.Context, url string, data ...any) (*Response, error)`
  Post sends request using HTTP method POST and returns the response object.
- `func (c *Client) Delete(ctx context.Context, url string, data ...any) (*Response, error)`
  Delete send DELETE request and returns the response object.
- `func (c *Client) Head(ctx context.Context, url string, data ...any) (*Response, error)`
  Head send HEAD request and returns the response object.
- `func (c *Client) Patch(ctx context.Context, url string, data ...any) (*Response, error)`
  Patch send PATCH request and returns the response object.
- `func (c *Client) Connect(ctx context.Context, url string, data ...any) (*Response, error)`
  Connect send CONNECT request and returns the response object.
- `func (c *Client) Options(ctx context.Context, url string, data ...any) (*Response, error)`
  Options send OPTIONS request and returns the response object.
- `func (c *Client) Trace(ctx context.Context, url string, data ...any) (*Response, error)`
  Trace send TRACE request and returns the response object.
- `func (c *Client) PostForm(ctx context.Context, url string, data map[string]string) (resp *Response, err error)`
  PostForm is different from net/http.PostForm.
- `func (c *Client) DoRequest(ctx context.Context, method string, url string, data ...any) (resp *Response, err error)`
  DoRequest sends request with given HTTP method and data and returns the respo...
- `func (c *Client) DoRequestObj(ctx context.Context, req any, res any) error`
  DoRequestObj does HTTP request using standard request/response object.
- `func (c *Client) GetVar(ctx context.Context, url string, data ...any) *gvar.Var`
  GetVar sends a GET request, retrieves and converts the result content to *gva...
- `func (c *Client) PutVar(ctx context.Context, url string, data ...any) *gvar.Var`
  PutVar sends a PUT request, retrieves and converts the result content to *gva...
- `func (c *Client) PostVar(ctx context.Context, url string, data ...any) *gvar.Var`
  PostVar sends a POST request, retrieves and converts the result content to *g...
- `func (c *Client) DeleteVar(ctx context.Context, url string, data ...any) *gvar.Var`
  DeleteVar sends a DELETE request, retrieves and converts the result content t...
- `func (c *Client) HeadVar(ctx context.Context, url string, data ...any) *gvar.Var`
  HeadVar sends a HEAD request, retrieves and converts the result content to *g...
- `func (c *Client) PatchVar(ctx context.Context, url string, data ...any) *gvar.Var`
  PatchVar sends a PATCH request, retrieves and converts the result content to ...
- `func (c *Client) ConnectVar(ctx context.Context, url string, data ...any) *gvar.Var`
  ConnectVar sends a CONNECT request, retrieves and converts the result content...
- `func (c *Client) OptionsVar(ctx context.Context, url string, data ...any) *gvar.Var`
  OptionsVar sends an OPTIONS request, retrieves and converts the result conten...
- `func (c *Client) TraceVar(ctx context.Context, url string, data ...any) *gvar.Var`
  TraceVar sends a TRACE request, retrieves and converts the result content to ...
- `func (c *Client) RequestVar(ctx context.Context, method string, url string, data ...any) *gvar.Var`
  RequestVar sends request using given HTTP method and data, retrieves converts...

### HandlerFunc

**类型**: type

HandlerFunc middleware handler func


### Response

**类型**: struct

Response is the struct for client request response.


**方法**:

- `func (r *Response) GetCookie(key string) string`
  GetCookie retrieves and returns the cookie value of specified `key`.
- `func (r *Response) GetCookieMap() map[string]string`
  GetCookieMap retrieves and returns a copy of current cookie values map.
- `func (r *Response) ReadAll() []byte`
  ReadAll retrieves and returns the response content as []byte.
- `func (r *Response) ReadAllString() string`
  ReadAllString retrieves and returns the response content as string.
- `func (r *Response) SetBodyContent(content []byte)`
  SetBodyContent overwrites response content with custom one.
- `func (r *Response) Close() error`
  Close closes the response when it will never be used.

### WebSocketClient

**类型**: struct

WebSocketClient wraps the underlying websocket client connection
and provides convenient functions.

Deprecated: please use third-party library for websocket client instead.


## 函数

### New

```go
func New() *Client
```

New creates and returns a new HTTP client object.

### LoadKeyCrt

```go
func LoadKeyCrt(crtFile string, keyFile string) (*tls.Config, error)
```

LoadKeyCrt creates and returns a TLS configuration object with given certificate and key files.

### NewWebSocket

```go
func NewWebSocket() *WebSocketClient
```

NewWebSocket creates and returns a new WebSocketClient object.

Deprecated: please use third-party library for websocket client instead.

## 内部依赖

- `github.com/gogf/gf/v2`
- `container/gmap`
- `container/gvar`
- `encoding/gjson`
- `encoding/gurl`
- `errors/gcode`
- `errors/gerror`
- `internal/httputil`
- `internal/intlog`
- `internal/json`
- `internal/utils`
- `net/gsel`
- `net/gsvc`
- `net/gtrace`
- `os/gctx`
- `os/gfile`
- `os/gmetric`
- `os/gtime`
- `text/gregex`
- `text/gstr`
- `util/gconv`
- `util/gmeta`
- `util/gtag`
- `util/gutil`

## 外部依赖

- `github.com/gorilla/websocket`
- `go.opentelemetry.io/otel`
- `go.opentelemetry.io/otel/attribute`
- `go.opentelemetry.io/otel/codes`
- `go.opentelemetry.io/otel/propagation`
- `go.opentelemetry.io/otel/trace`
- `golang.org/x/net/proxy`

