# gres

> Package: `github.com/gogf/gf/v2/os/gres`

```go
import "github.com/gogf/gf/v2/os/gres"
```

## 概述

Package gres provides resource management and packing/unpacking feature between files and bytes.

## 源文件

- `gres.go`
- `gres_file.go`
- `gres_func.go`
- `gres_func_zip.go`
- `gres_http_file.go`
- `gres_instance.go`
- `gres_resource.go`

## 类型定义

### File

**类型**: struct

File is a file in a zip file.


**方法**:

- `func (f *File) Name() string`
  Name returns the name of the file.
- `func (f *File) Open() (io.ReadCloser, error)`
  Open returns a ReadCloser that provides access to the File's contents.
- `func (f *File) Content() []byte`
  Content returns the content of the file.
- `func (f *File) FileInfo() os.FileInfo`
  FileInfo returns an os.FileInfo for the FileHeader.
- `func (f *File) Export(dst string, option ...ExportOption) error`
  Export exports and saves all its sub files to specified system path `dst` rec...
- `func (f File) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (f *File) Close() error`
  Close implements interface of http.File.
- `func (f *File) Readdir(count int) ([]os.FileInfo, error)`
  Readdir implements Readdir interface of http.File.
- `func (f *File) Stat() (os.FileInfo, error)`
  Stat implements Stat interface of http.File.
- `func (f *File) Read(b []byte) (n int, err error)`
  Read implements the io.Reader interface.
- `func (f *File) Seek(offset int64, whence int) (n int64, err error)`
  Seek implements the io.Seeker interface.

### Option

**类型**: struct

Option contains the extra options for Pack functions.


### Resource

**类型**: struct

Resource is the resource manager for the file system.


**方法**:

- `func (r *Resource) Add(content string, prefix ...string) error`
  Add unpacks and adds the `content` into current resource object.
- `func (r *Resource) Load(path string, prefix ...string) error`
  Load loads, unpacks and adds the data from `path` into current resource object.
- `func (r *Resource) Get(path string) *File`
  Get returns the file with given path.
- `func (r *Resource) GetWithIndex(path string, indexFiles []string) *File`
  GetWithIndex searches file with `path`, if the file is directory
- `func (r *Resource) GetContent(path string) []byte`
  GetContent directly returns the content of `path`.
- `func (r *Resource) Contains(path string) bool`
  Contains checks whether the `path` exists in current resource object.
- `func (r *Resource) IsEmpty() bool`
  IsEmpty checks and returns whether the resource manager is empty.
- `func (r *Resource) ScanDir(path string, pattern string, recursive ...bool) []*File`
  ScanDir returns the files under the given path, the parameter `path` should b...
- `func (r *Resource) ScanDirFile(path string, pattern string, recursive ...bool) []*File`
  ScanDirFile returns all sub-files with absolute paths of given `path`,
- `func (r *Resource) Export(src string, dst string, option ...ExportOption) error`
  Export exports and saves specified path `srcPath` and all its sub files to sp...
- `func (r *Resource) Dump()`
  Dump prints the files of current resource object.

### ExportOption

**类型**: struct

ExportOption is the option for function Export.


## 函数

### Add

```go
func Add(content string, prefix ...string) error
```

Add unpacks and adds the `content` into the default resource object.
The unnecessary parameter `prefix` indicates the prefix
for each file storing into current resource object.

### Load

```go
func Load(path string, prefix ...string) error
```

Load loads, unpacks and adds the data from `path` into the default resource object.
The unnecessary parameter `prefix` indicates the prefix
for each file storing into current resource object.

### Get

```go
func Get(path string) *File
```

Get returns the file with given path.

### GetWithIndex

```go
func GetWithIndex(path string, indexFiles []string) *File
```

GetWithIndex searches file with `path`, if the file is directory
it then does index files searching under this directory.

GetWithIndex is usually used for http static file service.

### GetContent

```go
func GetContent(path string) []byte
```

GetContent directly returns the content of `path` in default resource object.

### Contains

```go
func Contains(path string) bool
```

Contains checks whether the `path` exists in the default resource object.

### IsEmpty

```go
func IsEmpty() bool
```

IsEmpty checks and returns whether the resource manager is empty.

### ScanDir

```go
func ScanDir(path string, pattern string, recursive ...bool) []*File
```

ScanDir returns the files under the given path, the parameter `path` should be a folder type.

The pattern parameter `pattern` supports multiple file name patterns,
using the ',' symbol to separate multiple patterns.

It scans directory recursively if given parameter `recursive` is true.

### ScanDirFile

```go
func ScanDirFile(path string, pattern string, recursive ...bool) []*File
```

ScanDirFile returns all sub-files with absolute paths of given `path`,
It scans directory recursively if given parameter `recursive` is true.

Note that it returns only files, exclusive of directories.

### Export

```go
func Export(src string, dst string, option ...ExportOption) error
```

Export exports and saves specified path `src` and all its sub files to specified system path `dst` recursively.

### Dump

```go
func Dump()
```

Dump prints the files of the default resource object.

### Pack

```go
func Pack(srcPaths string, keyPrefix ...string) ([]byte, error)
```

Pack packs the path specified by `srcPaths` into bytes.
The unnecessary parameter `keyPrefix` indicates the prefix for each file
packed into the result bytes.

Note that parameter `srcPaths` supports multiple paths join with ','.

Deprecated: use PackWithOption instead.

### PackWithOption

```go
func PackWithOption(srcPaths string, option Option) ([]byte, error)
```

PackWithOption packs the path specified by `srcPaths` into bytes.

Note that parameter `srcPaths` supports multiple paths join with ','.

### PackToFile

```go
func PackToFile(srcPaths string, dstPath string, keyPrefix ...string) error
```

PackToFile packs the path specified by `srcPaths` to target file `dstPath`.
The unnecessary parameter `keyPrefix` indicates the prefix for each file
packed into the result bytes.

Note that parameter `srcPaths` supports multiple paths join with ','.

Deprecated: use PackToFileWithOption instead.

### PackToFileWithOption

```go
func PackToFileWithOption(srcPaths string, dstPath string, option Option) error
```

PackToFileWithOption packs the path specified by `srcPaths` to target file `dstPath`.

Note that parameter `srcPaths` supports multiple paths join with ','.

### PackToGoFile

```go
func PackToGoFile(srcPath string, goFilePath string, pkgName string, keyPrefix ...string) error
```

PackToGoFile packs the path specified by `srcPaths` to target go file `goFilePath`
with given package name `pkgName`.

The unnecessary parameter `keyPrefix` indicates the prefix for each file
packed into the result bytes.

Note that parameter `srcPaths` supports multiple paths join with ','.

Deprecated: use PackToGoFileWithOption instead.

### PackToGoFileWithOption

```go
func PackToGoFileWithOption(srcPath string, goFilePath string, pkgName string, option Option) error
```

PackToGoFileWithOption packs the path specified by `srcPaths` to target go file `goFilePath`
with given package name `pkgName`.

Note that parameter `srcPaths` supports multiple paths join with ','.

### Unpack

```go
func Unpack(path string) ([]*File, error)
```

Unpack unpacks the content specified by `path` to []*File.

### UnpackContent

```go
func UnpackContent(content string) ([]*File, error)
```

UnpackContent unpacks the content to []*File.

### Instance

```go
func Instance(name ...string) *Resource
```

Instance returns an instance of Resource.
The parameter `name` is the name for the instance.

### New

```go
func New() *Resource
```

New creates and returns a new resource object.

## 内部依赖

- `container/gmap`
- `container/gtree`
- `encoding/gbase64`
- `encoding/gcompress`
- `errors/gerror`
- `internal/fileinfo`
- `internal/intlog`
- `internal/json`
- `os/gfile`
- `os/gtime`
- `text/gregex`
- `text/gstr`

