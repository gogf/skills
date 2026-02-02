# ghttp

> Package: `github.com/gogf/gf/v2/net/ghttp`

```go
import "github.com/gogf/gf/v2/net/ghttp"
```

## 概述

Package ghttp provides powerful http server and simple client implements.

## 源文件

- `ghttp.go`
- `ghttp_func.go`
- `ghttp_middleware_cors.go`
- `ghttp_middleware_gzip.go`
- `ghttp_middleware_handler_response.go`
- `ghttp_middleware_json_body.go`
- `ghttp_middleware_never_done_ctx.go`
- `ghttp_middleware_tracing.go`
- `ghttp_request.go`
- `ghttp_request_auth.go`
- `ghttp_request_middleware.go`
- `ghttp_request_param.go`
- `ghttp_request_param_ctx.go`
- `ghttp_request_param_file.go`
- `ghttp_request_param_form.go`
- `ghttp_request_param_handler.go`
- `ghttp_request_param_page.go`
- `ghttp_request_param_param.go`
- `ghttp_request_param_query.go`
- `ghttp_request_param_request.go`
- `ghttp_request_param_router.go`
- `ghttp_request_view.go`
- `ghttp_response.go`
- `ghttp_response_cors.go`
- `ghttp_response_view.go`
- `ghttp_response_write.go`
- `ghttp_server.go`
- `ghttp_server_admin.go`
- `ghttp_server_admin_process.go`
- `ghttp_server_admin_unix.go`
- `ghttp_server_config.go`
- `ghttp_server_config_api.go`
- `ghttp_server_config_cookie.go`
- `ghttp_server_config_logging.go`
- `ghttp_server_config_mess.go`
- `ghttp_server_config_route.go`
- `ghttp_server_config_session.go`
- `ghttp_server_config_static.go`
- `ghttp_server_cookie.go`
- `ghttp_server_domain.go`
- `ghttp_server_error_logger.go`
- `ghttp_server_graceful.go`
- `ghttp_server_handler.go`
- `ghttp_server_log.go`
- `ghttp_server_metric.go`
- `ghttp_server_openapi.go`
- `ghttp_server_plugin.go`
- `ghttp_server_pprof.go`
- `ghttp_server_registry.go`
- `ghttp_server_router.go`
- `ghttp_server_router_group.go`
- `ghttp_server_router_hook.go`
- `ghttp_server_router_middleware.go`
- `ghttp_server_router_serve.go`
- `ghttp_server_service_handler.go`
- `ghttp_server_service_object.go`
- `ghttp_server_session.go`
- `ghttp_server_status.go`
- `ghttp_server_swagger.go`
- `ghttp_server_util.go`
- `ghttp_server_websocket.go`

## 类型定义

### Server

**类型**: struct

**方法**:

- `func (s *Server) Start() error`
  Start starts listening on configured port.
- `func (s *Server) GetOpenApi() *goai.OpenApiV3`
  GetOpenApi returns the OpenApi specification management object of the current...
- `func (s *Server) GetRoutes() []RouterItem`
  GetRoutes retrieves and returns the router array.
- `func (s *Server) Run()`
  Run starts server listening in blocking way.
- `func (s *Server) Status() ServerStatus`
  Status retrieves and returns the server status.
- `func (s *Server) GetListenedPort() int`
  GetListenedPort returns a port currently listened to by the server.
- `func (s *Server) GetListenedHTTPSPort() int`
  GetListenedHTTPSPort retrieves and returns one port which is listened using T...
- `func (s *Server) GetListenedPorts() []int`
  GetListenedPorts retrieves and returns the ports which are listened by curren...
- `func (s *Server) GetListenedAddress() string`
  GetListenedAddress retrieves and returns the address string which are listene...
- `func (s *Server) EnableAdmin(pattern ...string)`
  EnableAdmin enables the administration feature for the process.
- `func (s *Server) Shutdown() error`
  Shutdown shuts the current server.
- `func (s *Server) SetConfigWithMap(m map[string]any) error`
  SetConfigWithMap sets the configuration for the server using map.
- `func (s *Server) SetConfig(c ServerConfig) error`
  SetConfig sets the configuration for the server.
- `func (s *Server) SetAddr(address string)`
  SetAddr sets the listening address for the server.
- `func (s *Server) SetPort(port ...int)`
  SetPort sets the listening ports for the server.
- `func (s *Server) SetHTTPSAddr(address string)`
  SetHTTPSAddr sets the HTTPS listening ports for the server.
- `func (s *Server) SetHTTPSPort(port ...int)`
  SetHTTPSPort sets the HTTPS listening ports for the server.
- `func (s *Server) SetListener(listeners ...net.Listener) error`
  SetListener set the custom listener for the server.
- `func (s *Server) EnableHTTPS(certFile string, keyFile string, tlsConfig ...*tls.Config)`
  EnableHTTPS enables HTTPS with given certification and key files for the server.
- `func (s *Server) SetTLSConfig(tlsConfig *tls.Config)`
  SetTLSConfig sets custom TLS configuration and enables HTTPS feature for the ...
- `func (s *Server) SetReadTimeout(t time.Duration)`
  SetReadTimeout sets the ReadTimeout for the server.
- `func (s *Server) SetWriteTimeout(t time.Duration)`
  SetWriteTimeout sets the WriteTimeout for the server.
- `func (s *Server) SetIdleTimeout(t time.Duration)`
  SetIdleTimeout sets the IdleTimeout for the server.
- `func (s *Server) SetMaxHeaderBytes(b int)`
  SetMaxHeaderBytes sets the MaxHeaderBytes for the server.
- `func (s *Server) SetServerAgent(agent string)`
  SetServerAgent sets the ServerAgent for the server.
- `func (s *Server) SetKeepAlive(enabled bool)`
  SetKeepAlive sets the KeepAlive for the server.
- `func (s *Server) SetView(view *gview.View)`
  SetView sets the View for the server.
- `func (s *Server) GetName() string`
  GetName returns the name of the server.
- `func (s *Server) SetName(name string)`
  SetName sets the name for the server.
- `func (s *Server) SetEndpoints(endpoints []string)`
  SetEndpoints sets the Endpoints for the server.
- `func (s *Server) SetHandler(h func)`
  SetHandler sets the request handler for server.
- `func (s *Server) GetHandler() func`
  GetHandler returns the request handler of the server.
- `func (s *Server) SetRegistrar(registrar gsvc.Registrar)`
  SetRegistrar sets the Registrar for server.
- `func (s *Server) GetRegistrar() gsvc.Registrar`
  GetRegistrar returns the Registrar of server.
- `func (s *Server) SetSwaggerPath(path string)`
  SetSwaggerPath sets the SwaggerPath for server.
- `func (s *Server) SetSwaggerUITemplate(swaggerUITemplate string)`
  SetSwaggerUITemplate sets the Swagger template for server.
- `func (s *Server) SetOpenApiPath(path string)`
  SetOpenApiPath sets the OpenApiPath for server.
- `func (s *Server) GetOpenApiPath() string`
  GetOpenApiPath returns the `OpenApiPath` configuration of the server.
- `func (s *Server) SetCookieMaxAge(ttl time.Duration)`
  SetCookieMaxAge sets the CookieMaxAge for server.
- `func (s *Server) SetCookiePath(path string)`
  SetCookiePath sets the CookiePath for server.
- `func (s *Server) SetCookieDomain(domain string)`
  SetCookieDomain sets the CookieDomain for server.
- `func (s *Server) GetCookieMaxAge() time.Duration`
  GetCookieMaxAge returns the CookieMaxAge of the server.
- `func (s *Server) GetCookiePath() string`
  GetCookiePath returns the CookiePath of server.
- `func (s *Server) GetCookieDomain() string`
  GetCookieDomain returns CookieDomain of server.
- `func (s *Server) GetCookieSameSite() http.SameSite`
  GetCookieSameSite return CookieSameSite of server.
- `func (s *Server) GetCookieSecure() bool`
- `func (s *Server) GetCookieHttpOnly() bool`
- `func (s *Server) SetLogPath(path string) error`
  SetLogPath sets the log path for server.
- `func (s *Server) SetLogger(logger *glog.Logger)`
  SetLogger sets the logger for logging responsibility.
- `func (s *Server) Logger() *glog.Logger`
  Logger is alias of GetLogger.
- `func (s *Server) SetLogLevel(level string)`
  SetLogLevel sets logging level by level string.
- `func (s *Server) SetLogStdout(enabled bool)`
  SetLogStdout sets whether output the logging content to stdout.
- `func (s *Server) SetAccessLogEnabled(enabled bool)`
  SetAccessLogEnabled enables/disables the access log.
- `func (s *Server) SetErrorLogEnabled(enabled bool)`
  SetErrorLogEnabled enables/disables the error log.
- `func (s *Server) SetErrorStack(enabled bool)`
  SetErrorStack enables/disables the error stack feature.
- `func (s *Server) GetLogPath() string`
  GetLogPath returns the log path.
- `func (s *Server) IsAccessLogEnabled() bool`
  IsAccessLogEnabled checks whether the access log enabled.
- `func (s *Server) IsErrorLogEnabled() bool`
  IsErrorLogEnabled checks whether the error log enabled.
- `func (s *Server) SetNameToUriType(t int)`
  SetNameToUriType sets the NameToUriType for server.
- `func (s *Server) SetDumpRouterMap(enabled bool)`
  SetDumpRouterMap sets the DumpRouterMap for server.
- `func (s *Server) SetClientMaxBodySize(maxSize int64)`
  SetClientMaxBodySize sets the ClientMaxBodySize for server.
- `func (s *Server) SetFormParsingMemory(maxMemory int64)`
  SetFormParsingMemory sets the FormParsingMemory for server.
- `func (s *Server) SetGraceful(graceful bool)`
  SetGraceful sets the Graceful for server.
- `func (s *Server) GetGraceful() bool`
  GetGraceful returns the Graceful for server.
- `func (s *Server) SetGracefulTimeout(gracefulTimeout int)`
  SetGracefulTimeout sets the GracefulTimeout for server.
- `func (s *Server) GetGracefulTimeout() int`
  GetGracefulTimeout returns the GracefulTimeout for server.
- `func (s *Server) SetGracefulShutdownTimeout(gracefulShutdownTimeout int)`
  SetGracefulShutdownTimeout sets the GracefulShutdownTimeout for server.
- `func (s *Server) GetGracefulShutdownTimeout() int`
  GetGracefulShutdownTimeout returns the GracefulShutdownTimeout for server.
- `func (s *Server) SetRewrite(uri string, rewrite string)`
  SetRewrite sets rewrites for static URI for server.
- `func (s *Server) SetRewriteMap(rewrites map[string]string)`
  SetRewriteMap sets the rewritten map for server.
- `func (s *Server) SetRouteOverWrite(enabled bool)`
  SetRouteOverWrite sets the RouteOverWrite for server.
- `func (s *Server) SetSessionMaxAge(ttl time.Duration)`
  SetSessionMaxAge sets the SessionMaxAge for server.
- `func (s *Server) SetSessionIdName(name string)`
  SetSessionIdName sets the SessionIdName for server.
- `func (s *Server) SetSessionStorage(storage gsession.Storage)`
  SetSessionStorage sets the SessionStorage for server.
- `func (s *Server) SetSessionCookieOutput(enabled bool)`
  SetSessionCookieOutput sets the SetSessionCookieOutput for server.
- `func (s *Server) SetSessionCookieMaxAge(maxAge time.Duration)`
  SetSessionCookieMaxAge sets the SessionCookieMaxAge for server.
- `func (s *Server) GetSessionMaxAge() time.Duration`
  GetSessionMaxAge returns the SessionMaxAge of server.
- `func (s *Server) GetSessionIdName() string`
  GetSessionIdName returns the SessionIdName of server.
- `func (s *Server) GetSessionCookieMaxAge() time.Duration`
  GetSessionCookieMaxAge returns the SessionCookieMaxAge of server.
- `func (s *Server) SetIndexFiles(indexFiles []string)`
  SetIndexFiles sets the index files for server.
- `func (s *Server) GetIndexFiles() []string`
  GetIndexFiles retrieves and returns the index files from the server.
- `func (s *Server) SetIndexFolder(enabled bool)`
  SetIndexFolder enables/disables listing the sub-files if requesting a directory.
- `func (s *Server) SetFileServerEnabled(enabled bool)`
  SetFileServerEnabled enables/disables the static file service.
- `func (s *Server) SetServerRoot(root string)`
  SetServerRoot sets the document root for static service.
- `func (s *Server) AddSearchPath(path string)`
  AddSearchPath add searching directory path for static file service.
- `func (s *Server) AddStaticPath(prefix string, path string)`
  AddStaticPath sets the uri to static directory path mapping for static file s...
- `func (s *Server) Domain(domains string) *Domain`
  Domain creates and returns a domain object for management for one or more dom...
- `func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request)`
  ServeHTTP is the default handler for http request.
- `func (s *Server) Plugin(plugin ...Plugin)`
  Plugin adds plugin to the server.
- `func (s *Server) EnablePProf(pattern ...string)`
  EnablePProf enables PProf feature for server.
- `func (s *Server) Group(prefix string, groups ...func) *RouterGroup`
  Group creates and returns a RouterGroup object.
- `func (s *Server) BindHookHandler(pattern string, hook HookName, handler HandlerFunc)`
  BindHookHandler registers handler for specified hook.
- `func (s *Server) BindHookHandlerByMap(pattern string, hookMap map[HookName]HandlerFunc)`
  BindHookHandlerByMap registers handler for specified hook.
- `func (s *Server) BindMiddleware(pattern string, handlers ...HandlerFunc)`
  BindMiddleware registers one or more global middleware to the server.
- `func (s *Server) BindMiddlewareDefault(handlers ...HandlerFunc)`
  BindMiddlewareDefault registers one or more global middleware to the server u...
- `func (s *Server) Use(handlers ...HandlerFunc)`
  Use is the alias of BindMiddlewareDefault.
- `func (s *Server) BindHandler(pattern string, handler any)`
  BindHandler registers a handler function to server with a given pattern.
- `func (s *Server) BindObject(pattern string, object any, method ...string)`
  BindObject registers object to server routes with a given pattern.
- `func (s *Server) BindObjectMethod(pattern string, object any, method string)`
  BindObjectMethod registers specified method of the object to server routes wi...
- `func (s *Server) BindObjectRest(pattern string, object any)`
  BindObjectRest registers object in REST API styles to server with a specified...
- `func (s *Server) BindStatusHandler(status int, handler HandlerFunc)`
  BindStatusHandler registers handler for given status code.
- `func (s *Server) BindStatusHandlerByMap(handlerMap map[int]HandlerFunc)`
  BindStatusHandlerByMap registers handler for given status code using map.

### Router

**类型**: struct

### RouterItem

**类型**: struct

### HandlerFunc

**类型**: type

### HandlerItem

**类型**: struct

**方法**:

- `func (h *HandlerItem) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (h *HandlerItem) GetMetaTag(key string) string`
  GetMetaTag retrieves and returns the metadata value associated with the given...

### HandlerItemParsed

**类型**: struct

**方法**:

- `func (h *HandlerItemParsed) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (h *HandlerItemParsed) GetMetaTag(key string) string`
  GetMetaTag retrieves and returns the metadata value associated with the given...

### ServerStatus

**类型**: type

### HookName

**类型**: type

### HandlerType

**类型**: type

### DefaultHandlerResponse

**类型**: struct

DefaultHandlerResponse is the default implementation of HandlerResponse.


### Request

**类型**: struct

Request is the context object for a request.


**方法**:

- `func (r *Request) WebSocket() (*WebSocket, error)`
  WebSocket upgrades current request as a websocket request.
- `func (r *Request) Exit()`
  Exit exits executing of current HTTP handler.
- `func (r *Request) ExitAll()`
  ExitAll exits executing of current and following HTTP handlers.
- `func (r *Request) ExitHook()`
  ExitHook exits executing of current and following HTTP HOOK handlers.
- `func (r *Request) IsExited() bool`
  IsExited checks and returns whether current request is exited.
- `func (r *Request) GetHeader(key string, def ...string) string`
  GetHeader retrieves and returns the header value with given `key`.
- `func (r *Request) GetHost() string`
  GetHost returns current request host name, which might be a domain or an IP w...
- `func (r *Request) IsFileRequest() bool`
  IsFileRequest checks and returns whether current request is serving file.
- `func (r *Request) IsAjaxRequest() bool`
  IsAjaxRequest checks and returns whether current request is an AJAX request.
- `func (r *Request) GetClientIp() string`
  GetClientIp returns the client ip of this request without port.
- `func (r *Request) GetRemoteIp() string`
  GetRemoteIp returns the ip from RemoteAddr.
- `func (r *Request) GetSchema() string`
  GetSchema returns the schema of this request.
- `func (r *Request) GetUrl() string`
  GetUrl returns current URL of this request.
- `func (r *Request) GetSessionId() string`
  GetSessionId retrieves and returns session id from cookie or header.
- `func (r *Request) GetReferer() string`
  GetReferer returns referer of this request.
- `func (r *Request) GetError() error`
  GetError returns the error occurs in the procedure of the request.
- `func (r *Request) SetError(err error)`
  SetError sets custom error for current request.
- `func (r *Request) ReloadParam()`
  ReloadParam is used for modifying request parameter.
- `func (r *Request) BasicAuth(user string, pass string, tips ...string) bool`
  BasicAuth enables the http basic authentication feature with a given passport...
- `func (r *Request) Parse(pointer any) error`
  Parse is the most commonly used function, which converts request parameters t...
- `func (r *Request) ParseQuery(pointer any) error`
  ParseQuery performs like function Parse, but only parses the query parameters.
- `func (r *Request) ParseForm(pointer any) error`
  ParseForm performs like function Parse, but only parses the form parameters o...
- `func (r *Request) Get(key string, def ...any) *gvar.Var`
  Get is alias of GetRequest, which is one of the most commonly used functions for
- `func (r *Request) GetBody() []byte`
  GetBody retrieves and returns request body content as bytes.
- `func (r *Request) MakeBodyRepeatableRead(repeatableRead bool) []byte`
  MakeBodyRepeatableRead marks the request body could be repeatedly readable or...
- `func (r *Request) GetBodyString() string`
  GetBodyString retrieves and returns request body content as string.
- `func (r *Request) GetJson() (*gjson.Json, error)`
  GetJson parses current request content as JSON format, and returns the JSON o...
- `func (r *Request) GetMap(def ...map[string]any) map[string]any`
  GetMap is an alias and convenient function for GetRequestMap.
- `func (r *Request) GetMapStrStr(def ...map[string]any) map[string]string`
  GetMapStrStr is an alias and convenient function for GetRequestMapStrStr.
- `func (r *Request) GetStruct(pointer any, mapping ...map[string]string) error`
  GetStruct is an alias and convenient function for GetRequestStruct.
- `func (r *Request) GetMultipartForm() *multipart.Form`
  GetMultipartForm parses and returns the form as multipart forms.
- `func (r *Request) GetMultipartFiles(name string) []*multipart.FileHeader`
  GetMultipartFiles parses and returns the post files array.
- `func (r *Request) Context() context.Context`
  Context is alias for function GetCtx.
- `func (r *Request) GetCtx() context.Context`
  GetCtx retrieves and returns the request's context.
- `func (r *Request) GetNeverDoneCtx() context.Context`
  GetNeverDoneCtx creates and returns a never done context object,
- `func (r *Request) SetCtx(ctx context.Context)`
  SetCtx custom context for current request.
- `func (r *Request) GetCtxVar(key any, def ...any) *gvar.Var`
  GetCtxVar retrieves and returns a Var with a given key name.
- `func (r *Request) SetCtxVar(key any, value any)`
  SetCtxVar sets custom parameter to context with key-value pairs.
- `func (r *Request) GetUploadFile(name string) *UploadFile`
  GetUploadFile retrieves and returns the uploading file with specified form name.
- `func (r *Request) GetUploadFiles(name string) UploadFiles`
  GetUploadFiles retrieves and returns multiple uploading files with specified ...
- `func (r *Request) SetForm(key string, value any)`
  SetForm sets custom form value with key-value pairs.
- `func (r *Request) GetForm(key string, def ...any) *gvar.Var`
  GetForm retrieves and returns parameter `key` from form.
- `func (r *Request) GetFormMap(kvMap ...map[string]any) map[string]any`
  GetFormMap retrieves and returns all form parameters passed from client as map.
- `func (r *Request) GetFormMapStrStr(kvMap ...map[string]any) map[string]string`
  GetFormMapStrStr retrieves and returns all form parameters passed from client...
- `func (r *Request) GetFormMapStrVar(kvMap ...map[string]any) map[string]*gvar.Var`
  GetFormMapStrVar retrieves and returns all form parameters passed from client...
- `func (r *Request) GetFormStruct(pointer any, mapping ...map[string]string) error`
  GetFormStruct retrieves all form parameters passed from client and converts t...
- `func (r *Request) GetHandlerResponse() any`
  GetHandlerResponse retrieves and returns the handler response object and its ...
- `func (r *Request) GetServeHandler() *HandlerItemParsed`
  GetServeHandler retrieves and returns the user defined handler used to serve ...
- `func (r *Request) GetPage(totalSize int, pageSize int) *gpage.Page`
  GetPage creates and returns the pagination object for given `totalSize` and `...
- `func (r *Request) SetParam(key string, value any)`
  SetParam sets custom parameter with key-value pairs.
- `func (r *Request) SetParamMap(data map[string]any)`
  SetParamMap sets custom parameter with key-value pair maps.
- `func (r *Request) GetParam(key string, def ...any) *gvar.Var`
  GetParam returns custom parameter with a given name `key`.
- `func (r *Request) SetQuery(key string, value any)`
  SetQuery sets custom query value with key-value pairs.
- `func (r *Request) GetQuery(key string, def ...any) *gvar.Var`
  GetQuery retrieves and return parameter with the given name `key` from query ...
- `func (r *Request) GetQueryMap(kvMap ...map[string]any) map[string]any`
  GetQueryMap retrieves and returns all parameters passed from the client using...
- `func (r *Request) GetQueryMapStrStr(kvMap ...map[string]any) map[string]string`
  GetQueryMapStrStr retrieves and returns all parameters passed from the client...
- `func (r *Request) GetQueryMapStrVar(kvMap ...map[string]any) map[string]*gvar.Var`
  GetQueryMapStrVar retrieves and returns all parameters passed from the client...
- `func (r *Request) GetQueryStruct(pointer any, mapping ...map[string]string) error`
  GetQueryStruct retrieves all parameters passed from the client using the HTTP...
- `func (r *Request) GetRequest(key string, def ...any) *gvar.Var`
  GetRequest retrieves and returns the parameter named `key` passed from the cl...
- `func (r *Request) GetRequestMap(kvMap ...map[string]any) map[string]any`
  GetRequestMap retrieves and returns all parameters passed from the client and...
- `func (r *Request) GetRequestMapStrStr(kvMap ...map[string]any) map[string]string`
  GetRequestMapStrStr retrieve and returns all parameters passed from the clien...
- `func (r *Request) GetRequestMapStrVar(kvMap ...map[string]any) map[string]*gvar.Var`
  GetRequestMapStrVar retrieve and returns all parameters passed from the clien...
- `func (r *Request) GetRequestStruct(pointer any, mapping ...map[string]string) error`
  GetRequestStruct retrieves all parameters passed from the client and custom p...
- `func (r *Request) GetRouterMap() map[string]string`
  GetRouterMap retrieves and returns a copy of the router map.
- `func (r *Request) GetRouter(key string, def ...any) *gvar.Var`
  GetRouter retrieves and returns the router value with given key name `key`.
- `func (r *Request) SetView(view *gview.View)`
  SetView sets template view engine object for this request.
- `func (r *Request) GetView() *gview.View`
  GetView returns the template view engine object for this request.
- `func (r *Request) Assigns(data gview.Params)`
  Assigns binds multiple template variables to current request.
- `func (r *Request) Assign(key string, value any)`
  Assign binds a template variable to current request.

### UploadFile

**类型**: struct

UploadFile wraps the multipart uploading file with more and convenient features.


**方法**:

- `func (f UploadFile) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (f *UploadFile) Save(dirPath string, randomlyRename ...bool) (filename string, err error)`
  Save saves the single uploading file to directory path and returns the saved ...

### UploadFiles

**类型**: type

UploadFiles is an array type of *UploadFile.


**方法**:

- `func (fs UploadFiles) Save(dirPath string, randomlyRename ...bool) (filenames []string, err error)`
  Save saves all uploading files to specified directory path and returns the sa...

### Response

**类型**: struct

Response is the http response manager.
Note that it implements the http.ResponseWriter interface with buffering feature.


**方法**:

- `func (r *Response) ServeFile(path string, allowIndex ...bool)`
  ServeFile serves the file to the response.
- `func (r *Response) ServeFileDownload(path string, name ...string)`
  ServeFileDownload serves file downloading to the response.
- `func (r *Response) RedirectTo(location string, code ...int)`
  RedirectTo redirects the client to another location.
- `func (r *Response) RedirectBack(code ...int)`
  RedirectBack redirects the client back to referer.
- `func (r *Response) ServeContent(name string, modTime time.Time, content io.ReadSeeker)`
  ServeContent replies to the request using the content in the
- `func (r *Response) Flush()`
  Flush outputs the buffer content to the client and clears the buffer.
- `func (r *Response) DefaultCORSOptions() CORSOptions`
  DefaultCORSOptions returns the default CORS options,
- `func (r *Response) CORS(options CORSOptions)`
  CORS sets custom CORS options.
- `func (r *Response) CORSAllowedOrigin(options CORSOptions) bool`
  CORSAllowedOrigin CORSAllowed checks whether the current request origin is al...
- `func (r *Response) CORSDefault()`
  CORSDefault sets CORS with default CORS options,
- `func (r *Response) WriteTpl(tpl string, params ...gview.Params) error`
  WriteTpl parses and responses given template file.
- `func (r *Response) WriteTplDefault(params ...gview.Params) error`
  WriteTplDefault parses and responses the default template file.
- `func (r *Response) WriteTplContent(content string, params ...gview.Params) error`
  WriteTplContent parses and responses the template content.
- `func (r *Response) ParseTpl(tpl string, params ...gview.Params) (string, error)`
  ParseTpl parses given template file `tpl` with given template variables `params`
- `func (r *Response) ParseTplDefault(params ...gview.Params) (string, error)`
  ParseTplDefault parses the default template file with params.
- `func (r *Response) ParseTplContent(content string, params ...gview.Params) (string, error)`
  ParseTplContent parses given template file `file` with given template paramet...
- `func (r *Response) Write(content ...any)`
  Write writes `content` to the response buffer.
- `func (r *Response) WriteExit(content ...any)`
  WriteExit writes `content` to the response buffer and exits executing of curr...
- `func (r *Response) WriteOver(content ...any)`
  WriteOver overwrites the response buffer with `content`.
- `func (r *Response) WriteOverExit(content ...any)`
  WriteOverExit overwrites the response buffer with `content` and exits executing
- `func (r *Response) Writef(format string, params ...any)`
  Writef writes the response with fmt.Sprintf.
- `func (r *Response) WritefExit(format string, params ...any)`
  WritefExit writes the response with fmt.Sprintf and exits executing of curren...
- `func (r *Response) Writeln(content ...any)`
  Writeln writes the response with `content` and new line.
- `func (r *Response) WritelnExit(content ...any)`
  WritelnExit writes the response with `content` and new line and exits executing
- `func (r *Response) Writefln(format string, params ...any)`
  Writefln writes the response with fmt.Sprintf and new line.
- `func (r *Response) WriteflnExit(format string, params ...any)`
  WriteflnExit writes the response with fmt.Sprintf and new line and exits exec...
- `func (r *Response) WriteJson(content any)`
  WriteJson writes `content` to the response with JSON format.
- `func (r *Response) WriteJsonExit(content any)`
  WriteJsonExit writes `content` to the response with JSON format and exits exe...
- `func (r *Response) WriteJsonP(content any)`
  WriteJsonP writes `content` to the response with JSONP format.
- `func (r *Response) WriteJsonPExit(content any)`
  WriteJsonPExit writes `content` to the response with JSONP format and exits e...
- `func (r *Response) WriteXml(content any, rootTag ...string)`
  WriteXml writes `content` to the response with XML format.
- `func (r *Response) WriteXmlExit(content any, rootTag ...string)`
  WriteXmlExit writes `content` to the response with XML format and exits execu...
- `func (r *Response) WriteStatus(status int, content ...any)`
  WriteStatus writes HTTP `status` and `content` to the response.
- `func (r *Response) WriteStatusExit(status int, content ...any)`
  WriteStatusExit writes HTTP `status` and `content` to the response and exits ...

### CORSOptions

**类型**: struct

CORSOptions is the options for CORS feature.
See https://www.w3.org/TR/cors/ .


### ServerConfig

**类型**: struct

ServerConfig is the HTTP Server configuration manager.


### Cookie

**类型**: struct

Cookie for HTTP COOKIE management.


**方法**:

- `func (c *Cookie) Map() map[string]string`
  Map returns the cookie items as map[string]string.
- `func (c *Cookie) Contains(key string) bool`
  Contains checks if given key exists and not expire in cookie.
- `func (c *Cookie) Set(key string, value string)`
  Set sets cookie item with default domain, path and expiration age.
- `func (c *Cookie) SetCookie(key string, value string, domain string, path string, maxAge time.Duration, options ...CookieOptions)`
  SetCookie sets cookie item with given domain, path and expiration age.
- `func (c *Cookie) SetHttpCookie(httpCookie *http.Cookie)`
  SetHttpCookie sets cookie with *http.Cookie.
- `func (c *Cookie) GetSessionId() string`
  GetSessionId retrieves and returns the session id from cookie.
- `func (c *Cookie) SetSessionId(id string)`
  SetSessionId sets session id in the cookie.
- `func (c *Cookie) Get(key string, def ...string) *gvar.Var`
  Get retrieves and returns the value with specified key.
- `func (c *Cookie) Remove(key string)`
  Remove deletes specified key and its value from cookie using default domain a...
- `func (c *Cookie) RemoveCookie(key string, domain string, path string)`
  RemoveCookie deletes specified key and its value from cookie using given doma...
- `func (c *Cookie) Flush()`
  Flush outputs the cookie items to the client.

### CookieOptions

**类型**: struct

CookieOptions provides security config for cookies


### Domain

**类型**: struct

Domain is used for route register for domains.


**方法**:

- `func (d *Domain) BindHandler(pattern string, handler any)`
  BindHandler binds the handler for the specified pattern.
- `func (d *Domain) BindObject(pattern string, obj any, methods ...string)`
  BindObject binds the object for the specified pattern.
- `func (d *Domain) BindObjectMethod(pattern string, obj any, method string)`
  BindObjectMethod binds the method for the specified pattern.
- `func (d *Domain) BindObjectRest(pattern string, obj any)`
  BindObjectRest binds the RESTful API for the specified pattern.
- `func (d *Domain) BindHookHandler(pattern string, hook HookName, handler HandlerFunc)`
  BindHookHandler binds the hook handler for the specified pattern.
- `func (d *Domain) BindHookHandlerByMap(pattern string, hookMap map[HookName]HandlerFunc)`
  BindHookHandlerByMap binds the hook handler for the specified pattern.
- `func (d *Domain) BindStatusHandler(status int, handler HandlerFunc)`
  BindStatusHandler binds the status handler for the specified pattern.
- `func (d *Domain) BindStatusHandlerByMap(handlerMap map[int]HandlerFunc)`
  BindStatusHandlerByMap binds the status handler for the specified pattern.
- `func (d *Domain) BindMiddleware(pattern string, handlers ...HandlerFunc)`
  BindMiddleware binds the middleware for the specified pattern.
- `func (d *Domain) BindMiddlewareDefault(handlers ...HandlerFunc)`
  BindMiddlewareDefault binds the default middleware for the specified pattern.
- `func (d *Domain) Use(handlers ...HandlerFunc)`
  Use adds middleware to the domain.
- `func (d *Domain) EnablePProf(pattern ...string)`
  EnablePProf enables PProf feature for server of specified domain.
- `func (d *Domain) Group(prefix string, groups ...func) *RouterGroup`
  Group creates and returns a RouterGroup object, which is bound to a specified...

### Plugin

**类型**: interface

Plugin is the interface for server plugin.


### RouterGroup

**类型**: struct

**方法**:

- `func (g *RouterGroup) Group(prefix string, groups ...func) *RouterGroup`
  Group creates and returns a subgroup of the current router group.
- `func (g *RouterGroup) Clone() *RouterGroup`
  Clone returns a new router group which is a clone of the current group.
- `func (g *RouterGroup) Bind(handlerOrObject ...any) *RouterGroup`
  Bind does batch route registering feature for a router group.
- `func (g *RouterGroup) ALL(pattern string, object any, params ...any) *RouterGroup`
  ALL register an http handler to give the route pattern and all http methods.
- `func (g *RouterGroup) ALLMap(m map[string]any)`
  ALLMap registers http handlers for http methods using map.
- `func (g *RouterGroup) Map(m map[string]any)`
  Map registers http handlers for http methods using map.
- `func (g *RouterGroup) GET(pattern string, object any, params ...any) *RouterGroup`
  GET registers an http handler to give the route pattern and the http method: ...
- `func (g *RouterGroup) PUT(pattern string, object any, params ...any) *RouterGroup`
  PUT registers an http handler to give the route pattern and the http method: ...
- `func (g *RouterGroup) POST(pattern string, object any, params ...any) *RouterGroup`
  POST registers an http handler to give the route pattern and the http method:...
- `func (g *RouterGroup) DELETE(pattern string, object any, params ...any) *RouterGroup`
  DELETE registers an http handler to give the route pattern and the http metho...
- `func (g *RouterGroup) PATCH(pattern string, object any, params ...any) *RouterGroup`
  PATCH registers an http handler to give the route pattern and the http method...
- `func (g *RouterGroup) HEAD(pattern string, object any, params ...any) *RouterGroup`
  HEAD registers an http handler to give the route pattern and the http method:...
- `func (g *RouterGroup) CONNECT(pattern string, object any, params ...any) *RouterGroup`
  CONNECT registers an http handler to give the route pattern and the http meth...
- `func (g *RouterGroup) OPTIONS(pattern string, object any, params ...any) *RouterGroup`
  OPTIONS register an http handler to give the route pattern and the http metho...
- `func (g *RouterGroup) TRACE(pattern string, object any, params ...any) *RouterGroup`
  TRACE registers an http handler to give the route pattern and the http method...
- `func (g *RouterGroup) REST(pattern string, object any) *RouterGroup`
  REST registers an http handler to give the route pattern according to REST rule.
- `func (g *RouterGroup) Hook(pattern string, hook HookName, handler HandlerFunc) *RouterGroup`
  Hook registers a hook to given route pattern.
- `func (g *RouterGroup) Middleware(handlers ...HandlerFunc) *RouterGroup`
  Middleware binds one or more middleware to the router group.

### Session

**类型**: type

Session is actually an alias of gsession.Session,
which is bound to a single request.


### WebSocket

**类型**: struct

WebSocket wraps the underlying websocket connection
and provides convenient functions.

Deprecated: will be removed in the future, please use third-party websocket library instead.


## 函数

### SupportedMethods

```go
func SupportedMethods() []string
```

SupportedMethods returns all supported HTTP methods.

### BuildParams

```go
func BuildParams(params any, noUrlEncode ...bool) encodedParamStr string
```

BuildParams builds the request string for the http client. The `params` can be type of:
string/[]byte/map/struct/*struct.

The optional parameter `noUrlEncode` specifies whether to ignore the url encoding for the data.

### MiddlewareCORS

```go
func MiddlewareCORS(r *Request)
```

MiddlewareCORS is a middleware handler for CORS with default options.

### MiddlewareGzip

```go
func MiddlewareGzip(r *Request)
```

MiddlewareGzip is a middleware that compresses HTTP response using gzip compression.
Note that it does not compress responses if:
1. The response is already compressed (Content-Encoding header is set)
2. The client does not accept gzip compression
3. The response body length is too small (less than 1KB)

To disable compression for specific routes, you can use the group middleware:

	group.Group("/api", func(group *ghttp.RouterGroup) {
	    group.Middleware(ghttp.MiddlewareGzip) // Enable GZIP for /api routes
	})

### MiddlewareHandlerResponse

```go
func MiddlewareHandlerResponse(r *Request)
```

MiddlewareHandlerResponse is the default middleware handling handler response object and its error.

### MiddlewareJsonBody

```go
func MiddlewareJsonBody(r *Request)
```

MiddlewareJsonBody validates and returns request body whether JSON format.

### MiddlewareNeverDoneCtx

```go
func MiddlewareNeverDoneCtx(r *Request)
```

MiddlewareNeverDoneCtx sets the context never done for current process.

### RequestFromCtx

```go
func RequestFromCtx(ctx context.Context) *Request
```

RequestFromCtx retrieves and returns the Request object from context.

### GetServer

```go
func GetServer(name ...any) *Server
```

GetServer creates and returns a server instance using given name and default configurations.
Note that the parameter `name` should be unique for different servers. It returns an existing
server instance if given `name` is already existing in the server mapping.

### Wait

```go
func Wait()
```

Wait blocks to wait for all servers done.
It's commonly used in multiple server situation.

### RestartAllServer

```go
func RestartAllServer(ctx context.Context, newExeFilePath string) error
```

RestartAllServer restarts all the servers of the process gracefully.
The optional parameter `newExeFilePath` specifies the new binary file for creating process.

### ShutdownAllServer

```go
func ShutdownAllServer(ctx context.Context) error
```

ShutdownAllServer shuts down all servers of current process gracefully.

### NewConfig

```go
func NewConfig() ServerConfig
```

NewConfig creates and returns a ServerConfig object with default configurations.
Note that, do not define this default configuration to local package variable, as there are
some pointer attributes that may be shared in different servers.

### ConfigFromMap

```go
func ConfigFromMap(m map[string]any) (ServerConfig, error)
```

ConfigFromMap creates and returns a ServerConfig object with given map and
default configuration object.

### GetCookie

```go
func GetCookie(r *Request) *Cookie
```

GetCookie creates or retrieves a cookie object with given request.
It retrieves and returns an existing cookie object if it already exists with given request.
It creates and returns a new cookie object if it does not exist with given request.

### StartPProfServer

```go
func StartPProfServer(address string, pattern ...string) (s *Server, err error)
```

StartPProfServer starts and runs a new server for pprof in another goroutine.

### WrapF

```go
func WrapF(f http.HandlerFunc) HandlerFunc
```

WrapF is a helper function for wrapping http.HandlerFunc and returns a ghttp.HandlerFunc.

### WrapH

```go
func WrapH(h http.Handler) HandlerFunc
```

WrapH is a helper function for wrapping http.Handler and returns a ghttp.HandlerFunc.

## 内部依赖

- `github.com/gogf/gf/v2`
- `container/garray`
- `container/glist`
- `container/gmap`
- `container/gset`
- `container/gtype`
- `container/gvar`
- `debug/gdebug`
- `encoding/gbase64`
- `encoding/ghtml`
- `encoding/gjson`
- `encoding/gurl`
- `encoding/gxml`
- `errors/gcode`
- `errors/gerror`
- `internal/consts`
- `internal/httputil`
- `internal/instance`
- `internal/intlog`
- `internal/json`
- `internal/reflection`
- `internal/utils`
- `net/ghttp/internal/graceful`
- `net/ghttp/internal/response`
- `net/ghttp/internal/swaggerui`
- `net/gipv4`
- `net/goai`
- `net/gsvc`
- `net/gtrace`
- `os/gcache`
- `os/gcfg`
- `os/gctx`
- `os/genv`
- `os/gfile`
- `os/glog`
- `os/gmetric`
- `os/gproc`
- `os/gres`
- `os/gsession`
- `os/gspath`
- `os/gstructs`
- `os/gtime`
- `os/gtimer`
- `os/gview`
- `text/gregex`
- `text/gstr`
- `util/gconv`
- `util/gmeta`
- `util/gmode`
- `util/gpage`
- `util/grand`
- `util/gtag`
- `util/guid`
- `util/gutil`
- `util/gvalid`

## 外部依赖

- `github.com/gorilla/websocket`
- `github.com/olekukonko/tablewriter`
- `github.com/olekukonko/tablewriter/renderer`
- `github.com/olekukonko/tablewriter/tw`
- `go.opentelemetry.io/otel`
- `go.opentelemetry.io/otel/attribute`
- `go.opentelemetry.io/otel/codes`
- `go.opentelemetry.io/otel/propagation`
- `go.opentelemetry.io/otel/trace`

