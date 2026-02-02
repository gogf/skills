# gudp

> Package: `github.com/gogf/gf/v2/net/gudp`

```go
import "github.com/gogf/gf/v2/net/gudp"
```

## 概述

Package gudp provides UDP server and client implementations.

## 源文件

- `gudp.go`
- `gudp_conn.go`
- `gudp_conn_client_conn.go`
- `gudp_conn_server_conn.go`
- `gudp_func.go`
- `gudp_server.go`

## 类型定义

### Retry

**类型**: struct

Retry holds the retry options.
TODO replace with standalone retry package.


### ClientConn

**类型**: struct

ClientConn holds the client side connection.


**方法**:

- `func (c *ClientConn) Send(data []byte, retry ...Retry) err error`
  Send writes data to remote address.
- `func (c *ClientConn) SendRecv(data []byte, receive int, retry ...Retry) ([]byte, error)`
  SendRecv writes data to connection and blocks reading response.

### ServerConn

**类型**: struct

ServerConn holds the server side connection.


**方法**:

- `func (c *ServerConn) Send(data []byte, remoteAddr *net.UDPAddr, retry ...Retry) err error`
  Send writes data to remote address.

### Server

**类型**: struct

Server is the UDP server.


**方法**:

- `func (s *Server) SetAddress(address string)`
  SetAddress sets the server address for UDP server.
- `func (s *Server) SetHandler(handler ServerHandler)`
  SetHandler sets the connection handler for UDP server.
- `func (s *Server) Close() err error`
  Close closes the connection.
- `func (s *Server) Run() error`
  Run starts listening UDP connection.
- `func (s *Server) GetListenedAddress() string`
  GetListenedAddress retrieves and returns the address string which are listene...
- `func (s *Server) GetListenedPort() int`
  GetListenedPort retrieves and returns one port which is listened to by curren...

### ServerHandler

**类型**: type

ServerHandler handles all server connections.


## 函数

### NewClientConn

```go
func NewClientConn(remoteAddress string, localAddress ...string) (*ClientConn, error)
```

NewClientConn creates UDP connection to `remoteAddress`.
The optional parameter `localAddress` specifies the local address for connection.

### NewServerConn

```go
func NewServerConn(listenedConn *net.UDPConn) *ServerConn
```

NewServerConn creates an udp connection that listens to `localAddress`.

### NewNetConn

```go
func NewNetConn(remoteAddress string, localAddress ...string) (*net.UDPConn, error)
```

NewNetConn creates and returns a *net.UDPConn with given addresses.

### Send

```go
func Send(address string, data []byte, retry ...Retry) error
```

Send writes data to `address` using UDP connection and then closes the connection.
Note that it is used for short connection usage.

### SendRecv

```go
func SendRecv(address string, data []byte, receive int, retry ...Retry) ([]byte, error)
```

SendRecv writes data to `address` using UDP connection, reads response and then closes the connection.
Note that it is used for short connection usage.

### MustGetFreePort

```go
func MustGetFreePort() port int
```

MustGetFreePort performs as GetFreePort, but it panics if any error occurs.

Deprecated: the port might be used soon after they were returned, please use `:0` as the listening
address which asks system to assign a free port instead.

### GetFreePort

```go
func GetFreePort() (port int, err error)
```

GetFreePort retrieves and returns a port that is free.

Deprecated: the port might be used soon after they were returned, please use `:0` as the listening
address which asks system to assign a free port instead.

### GetFreePorts

```go
func GetFreePorts(count int) (ports []int, err error)
```

GetFreePorts retrieves and returns specified number of ports that are free.

Deprecated: the ports might be used soon after they were returned, please use `:0` as the listening
address which asks system to assign a free port instead.

### GetServer

```go
func GetServer(name ...any) *Server
```

GetServer creates and returns an udp server instance with given name.

### NewServer

```go
func NewServer(address string, handler ServerHandler, name ...string) *Server
```

NewServer creates and returns an udp server.
The optional parameter `name` is used to specify its name, which can be used for
GetServer function to retrieve its instance.

## 内部依赖

- `container/gmap`
- `errors/gcode`
- `errors/gerror`
- `text/gstr`
- `util/gconv`

