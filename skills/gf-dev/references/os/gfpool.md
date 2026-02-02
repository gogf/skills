# gfpool

> Package: `github.com/gogf/gf/v2/os/gfpool`

```go
import "github.com/gogf/gf/v2/os/gfpool"
```

## 概述

Package gfpool provides io-reusable pool for file pointer.

## 源文件

- `gfpool.go`
- `gfpool_file.go`
- `gfpool_pool.go`

## 类型定义

### Pool

**类型**: struct

Pool pointer pool.


**方法**:

- `func (p *Pool) File() (*File, error)`
  File retrieves file item from the file pointer pool and returns it. It create...
- `func (p *Pool) Close()`
  Close closes current file pointer pool.

### File

**类型**: struct

File is an item in the pool.


**方法**:

- `func (f *File) Stat() (os.FileInfo, error)`
  Stat returns the FileInfo structure describing file.
- `func (f *File) Close(close ...bool) error`
  Close puts the file pointer back to the file pointer pool.

## 函数

### Open

```go
func Open(path string, flag int, perm os.FileMode, ttl ...time.Duration) (file *File, err error)
```

Open creates and returns a file item with given file path, flag and opening permission.
It automatically creates an associated file pointer pool internally when it's called first time.
It retrieves a file item from the file pointer pool after then.

### Get

```go
func Get(path string, flag int, perm os.FileMode, ttl ...time.Duration) file *File
```

Get returns a file item with given file path, flag and opening permission.
It retrieves a file item from the file pointer pool after then.

### New

```go
func New(path string, flag int, perm os.FileMode, ttl ...time.Duration) *Pool
```

New creates and returns a file pointer pool with given file path, flag and opening permission.

Note the expiration logic:
ttl = 0 : not expired;
ttl < 0 : immediate expired after use;
ttl > 0 : timeout expired;
It is not expired in default.

## 内部依赖

- `container/gmap`
- `container/gpool`
- `container/gtype`
- `errors/gerror`
- `os/gfsnotify`

