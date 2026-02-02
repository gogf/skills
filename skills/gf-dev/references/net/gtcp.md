# gtcp

> Package: `github.com/gogf/gf/v2/net/gtcp`

```go
import "github.com/gogf/gf/v2/net/gtcp"
```

## 概述

Package gtcp provides TCP server and client implementations.

## 源文件

- `gtcp.go`
- `gtcp_conn.go`
- `gtcp_conn_pkg.go`
- `gtcp_func.go`
- `gtcp_func_pkg.go`
- `gtcp_pool.go`
- `gtcp_pool_pkg.go`
- `gtcp_server.go`

## 类型定义

### Conn

**类型**: struct

Conn is the TCP connection object.


**方法**:

- `func (c *Conn) Send(data []byte, retry ...Retry) error`
  Send writes data to remote address.
- `func (c *Conn) Recv(length int, retry ...Retry) ([]byte, error)`
  Recv receives and returns data from the connection.
- `func (c *Conn) RecvLine(retry ...Retry) ([]byte, error)`
  RecvLine reads data from the connection until reads char '\n'.
- `func (c *Conn) RecvTill(til []byte, retry ...Retry) ([]byte, error)`
  RecvTill reads data from the connection until reads bytes `til`.
- `func (c *Conn) RecvWithTimeout(length int, timeout time.Duration, retry ...Retry) (data []byte, err error)`
  RecvWithTimeout reads data from the connection with timeout.
- `func (c *Conn) SendWithTimeout(data []byte, timeout time.Duration, retry ...Retry) err error`
  SendWithTimeout writes data to the connection with timeout.
- `func (c *Conn) SendRecv(data []byte, length int, retry ...Retry) ([]byte, error)`
  SendRecv writes data to the connection and blocks reading response.
- `func (c *Conn) SendRecvWithTimeout(data []byte, length int, timeout time.Duration, retry ...Retry) ([]byte, error)`
  SendRecvWithTimeout writes data to the connection and reads response with tim...
- `func (c *Conn) SetDeadline(t time.Time) err error`
  SetDeadline sets the deadline for current connection.
- `func (c *Conn) SetDeadlineRecv(t time.Time) err error`
  SetDeadlineRecv sets the deadline of receiving for current connection.
- `func (c *Conn) SetDeadlineSend(t time.Time) err error`
  SetDeadlineSend sets the deadline of sending for current connection.
- `func (c *Conn) SetBufferWaitRecv(bufferWaitDuration time.Duration)`
  SetBufferWaitRecv sets the buffer waiting timeout when reading all data from ...
- `func (c *Conn) SendPkg(data []byte, option ...PkgOption) error`
  SendPkg send data using simple package protocol.
- `func (c *Conn) SendPkgWithTimeout(data []byte, timeout time.Duration, option ...PkgOption) err error`
  SendPkgWithTimeout writes data to connection with timeout using simple packag...
- `func (c *Conn) SendRecvPkg(data []byte, option ...PkgOption) ([]byte, error)`
  SendRecvPkg writes data to connection and blocks reading response using simpl...
- `func (c *Conn) SendRecvPkgWithTimeout(data []byte, timeout time.Duration, option ...PkgOption) ([]byte, error)`
  SendRecvPkgWithTimeout writes data to connection and reads response with time...
- `func (c *Conn) RecvPkg(option ...PkgOption) (result []byte, err error)`
  RecvPkg receives data from connection using simple package protocol.
- `func (c *Conn) RecvPkgWithTimeout(timeout time.Duration, option ...PkgOption) (data []byte, err error)`
  RecvPkgWithTimeout reads data from connection with timeout using simple packa...

### PkgOption

**类型**: struct

PkgOption is package option for simple protocol.


### Retry

**类型**: struct

### PoolConn

**类型**: struct

PoolConn is a connection with pool feature for TCP.
Note that it is NOT a pool or connection manager, it is just a TCP connection object.


**方法**:

- `func (c *PoolConn) Close() error`
  Close puts back the connection to the pool if it's active,
- `func (c *PoolConn) Send(data []byte, retry ...Retry) error`
  Send writes data to the connection. It retrieves a new connection from its po...
- `func (c *PoolConn) Recv(length int, retry ...Retry) ([]byte, error)`
  Recv receives data from the connection.
- `func (c *PoolConn) RecvLine(retry ...Retry) ([]byte, error)`
  RecvLine reads data from the connection until reads char '\n'.
- `func (c *PoolConn) RecvTill(til []byte, retry ...Retry) ([]byte, error)`
  RecvTill reads data from the connection until reads bytes `til`.
- `func (c *PoolConn) RecvWithTimeout(length int, timeout time.Duration, retry ...Retry) (data []byte, err error)`
  RecvWithTimeout reads data from the connection with timeout.
- `func (c *PoolConn) SendWithTimeout(data []byte, timeout time.Duration, retry ...Retry) err error`
  SendWithTimeout writes data to the connection with timeout.
- `func (c *PoolConn) SendRecv(data []byte, receive int, retry ...Retry) ([]byte, error)`
  SendRecv writes data to the connection and blocks reading response.
- `func (c *PoolConn) SendRecvWithTimeout(data []byte, receive int, timeout time.Duration, retry ...Retry) ([]byte, error)`
  SendRecvWithTimeout writes data to the connection and reads response with tim...
- `func (c *PoolConn) SendPkg(data []byte, option ...PkgOption) err error`
  SendPkg sends a package containing `data` to the connection.
- `func (c *PoolConn) RecvPkg(option ...PkgOption) ([]byte, error)`
  RecvPkg receives package from connection using simple package protocol.
- `func (c *PoolConn) RecvPkgWithTimeout(timeout time.Duration, option ...PkgOption) (data []byte, err error)`
  RecvPkgWithTimeout reads data from connection with timeout using simple packa...
- `func (c *PoolConn) SendPkgWithTimeout(data []byte, timeout time.Duration, option ...PkgOption) err error`
  SendPkgWithTimeout writes data to connection with timeout using simple packag...
- `func (c *PoolConn) SendRecvPkg(data []byte, option ...PkgOption) ([]byte, error)`
  SendRecvPkg writes data to connection and blocks reading response using simpl...
- `func (c *PoolConn) SendRecvPkgWithTimeout(data []byte, timeout time.Duration, option ...PkgOption) ([]byte, error)`
  SendRecvPkgWithTimeout reads data from connection with timeout using simple p...

### Server

**类型**: struct

Server is a TCP server.


**方法**:

- `func (s *Server) SetAddress(address string)`
  SetAddress sets the listening address for server.
- `func (s *Server) GetAddress() string`
  GetAddress get the listening address for server.
- `func (s *Server) SetHandler(handler func)`
  SetHandler sets the connection handler for server.
- `func (s *Server) SetTLSKeyCrt(crtFile string, keyFile string) error`
  SetTLSKeyCrt sets the certificate and key file for TLS configuration of server.
- `func (s *Server) SetTLSConfig(tlsConfig *tls.Config)`
  SetTLSConfig sets the TLS configuration of server.
- `func (s *Server) Close() error`
  Close closes the listener and shutdowns the server.
- `func (s *Server) Run() err error`
  Run starts running the TCP Server.
- `func (s *Server) GetListenedAddress() string`
  GetListenedAddress retrieves and returns the address string which are listene...
- `func (s *Server) GetListenedPort() int`
  GetListenedPort retrieves and returns one port which is listened to by curren...

## 函数

### NewConn

```go
func NewConn(addr string, timeout ...time.Duration) (*Conn, error)
```

NewConn creates and returns a new connection with given address.

### NewConnTLS

```go
func NewConnTLS(addr string, tlsConfig *tls.Config) (*Conn, error)
```

NewConnTLS creates and returns a new TLS connection
with given address and TLS configuration.

### NewConnKeyCrt

```go
func NewConnKeyCrt(addr string, crtFile string, keyFile string) (*Conn, error)
```

NewConnKeyCrt creates and returns a new TLS connection
with given address and TLS certificate and key files.

### NewConnByNetConn

```go
func NewConnByNetConn(conn net.Conn) *Conn
```

NewConnByNetConn creates and returns a TCP connection object with given net.Conn object.

### NewNetConn

```go
func NewNetConn(address string, timeout ...time.Duration) (net.Conn, error)
```

NewNetConn creates and returns a net.Conn with given address like "127.0.0.1:80".
The optional parameter `timeout` specifies the timeout for dialing connection.

### NewNetConnTLS

```go
func NewNetConnTLS(address string, tlsConfig *tls.Config, timeout ...time.Duration) (net.Conn, error)
```

NewNetConnTLS creates and returns a TLS net.Conn with given address like "127.0.0.1:80".
The optional parameter `timeout` specifies the timeout for dialing connection.

### NewNetConnKeyCrt

```go
func NewNetConnKeyCrt(addr string, crtFile string, keyFile string, timeout ...time.Duration) (net.Conn, error)
```

NewNetConnKeyCrt creates and returns a TLS net.Conn with given TLS certificate and key files
and address like "127.0.0.1:80". The optional parameter `timeout` specifies the timeout for
dialing connection.

### Send

```go
func Send(address string, data []byte, retry ...Retry) error
```

Send creates connection to `address`, writes `data` to the connection and then closes the connection.
The optional parameter `retry` specifies the retry policy when fails in writing data.

### SendRecv

```go
func SendRecv(address string, data []byte, length int, retry ...Retry) ([]byte, error)
```

SendRecv creates connection to `address`, writes `data` to the connection, receives response
and then closes the connection.

The parameter `length` specifies the bytes count waiting to receive. It receives all buffer content
and returns if `length` is -1.

The optional parameter `retry` specifies the retry policy when fails in writing data.

### SendWithTimeout

```go
func SendWithTimeout(address string, data []byte, timeout time.Duration, retry ...Retry) error
```

SendWithTimeout does Send logic with writing timeout limitation.

### SendRecvWithTimeout

```go
func SendRecvWithTimeout(address string, data []byte, receive int, timeout time.Duration, retry ...Retry) ([]byte, error)
```

SendRecvWithTimeout does SendRecv logic with reading timeout limitation.

### LoadKeyCrt

```go
func LoadKeyCrt(crtFile string, keyFile string) (*tls.Config, error)
```

LoadKeyCrt creates and returns a TLS configuration object with given certificate and key files.

### MustGetFreePort

```go
func MustGetFreePort() int
```

MustGetFreePort performs as GetFreePort, but it panics is any error occurs.

### GetFreePort

```go
func GetFreePort() (port int, err error)
```

GetFreePort retrieves and returns a port that is free.

### GetFreePorts

```go
func GetFreePorts(count int) (ports []int, err error)
```

GetFreePorts retrieves and returns specified number of ports that are free.

### SendPkg

```go
func SendPkg(address string, data []byte, option ...PkgOption) error
```

SendPkg sends a package containing `data` to `address` and closes the connection.
The optional parameter `option` specifies the package options for sending.

### SendRecvPkg

```go
func SendRecvPkg(address string, data []byte, option ...PkgOption) ([]byte, error)
```

SendRecvPkg sends a package containing `data` to `address`, receives the response
and closes the connection. The optional parameter `option` specifies the package options for sending.

### SendPkgWithTimeout

```go
func SendPkgWithTimeout(address string, data []byte, timeout time.Duration, option ...PkgOption) error
```

SendPkgWithTimeout sends a package containing `data` to `address` with timeout limitation
and closes the connection. The optional parameter `option` specifies the package options for sending.

### SendRecvPkgWithTimeout

```go
func SendRecvPkgWithTimeout(address string, data []byte, timeout time.Duration, option ...PkgOption) ([]byte, error)
```

SendRecvPkgWithTimeout sends a package containing `data` to `address`, receives the response with timeout limitation
and closes the connection. The optional parameter `option` specifies the package options for sending.

### NewPoolConn

```go
func NewPoolConn(addr string, timeout ...time.Duration) (*PoolConn, error)
```

NewPoolConn creates and returns a connection with pool feature.

### GetServer

```go
func GetServer(name ...any) *Server
```

GetServer returns the TCP server with specified `name`,
or it returns a new normal TCP server named `name` if it does not exist.
The parameter `name` is used to specify the TCP server

### NewServer

```go
func NewServer(address string, handler func, name ...string) *Server
```

NewServer creates and returns a new normal TCP server.
The parameter `name` is optional, which is used to specify the instance name of the server.

### NewServerTLS

```go
func NewServerTLS(address string, tlsConfig *tls.Config, handler func, name ...string) *Server
```

NewServerTLS creates and returns a new TCP server with TLS support.
The parameter `name` is optional, which is used to specify the instance name of the server.

### NewServerKeyCrt

```go
func NewServerKeyCrt(address string, crtFile string, keyFile string, handler func, name ...string) (*Server, error)
```

NewServerKeyCrt creates and returns a new TCP server with TLS support.
The parameter `name` is optional, which is used to specify the instance name of the server.

## 内部依赖

- `container/gmap`
- `container/gpool`
- `errors/gcode`
- `errors/gerror`
- `os/gfile`
- `text/gstr`
- `util/gconv`

