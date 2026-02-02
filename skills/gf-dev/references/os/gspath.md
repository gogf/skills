# gspath

> Package: `github.com/gogf/gf/v2/os/gspath`

```go
import "github.com/gogf/gf/v2/os/gspath"
```

## 概述

Package gspath implements file index and search for folders.

## 源文件

- `gspath.go`
- `gspath_cache.go`

## 类型定义

### SPath

**类型**: struct

SPath manages the path searching feature.


**方法**:

- `func (sp *SPath) Set(path string) (realPath string, err error)`
  Set deletes all other searching directories and sets the searching directory ...
- `func (sp *SPath) Add(path string) (realPath string, err error)`
  Add adds more searching directory to the manager.
- `func (sp *SPath) Search(name string, indexFiles ...string) (filePath string, isDir bool)`
  Search searches file `name` in the manager.
- `func (sp *SPath) Remove(path string)`
  Remove deletes the `path` from cache files of the manager.
- `func (sp *SPath) Paths() []string`
  Paths returns all searching directories.
- `func (sp *SPath) AllPaths() []string`
  AllPaths returns all paths cached in the manager.
- `func (sp *SPath) Size() int`
  Size returns the count of the searching directories.

### SPathCacheItem

**类型**: struct

SPathCacheItem is a cache item for searching.


## 函数

### New

```go
func New(path string, cache bool) *SPath
```

New creates and returns a new path searching manager.

### Get

```go
func Get(root string, cache bool) *SPath
```

Get creates and returns an instance of searching manager for given path.
The parameter `cache` specifies whether using cache feature for this manager.
If cache feature is enabled, it asynchronously and recursively scans the path
and updates all sub files/folders to the cache using package gfsnotify.

### Search

```go
func Search(root string, name string, indexFiles ...string) (filePath string, isDir bool)
```

Search searches file `name` under path `root`.
The parameter `root` should be an absolute path. It will not automatically
convert `root` to absolute path for performance reason.
The optional parameter `indexFiles` specifies the searching index files when the result is a directory.
For example, if the result `filePath` is a directory, and `indexFiles` is [index.html, main.html], it will also
search [index.html, main.html] under `filePath`. It returns the absolute file path if any of them found,
or else it returns `filePath`.

### SearchWithCache

```go
func SearchWithCache(root string, name string, indexFiles ...string) (filePath string, isDir bool)
```

SearchWithCache searches file `name` under path `root` with cache feature enabled.
The parameter `root` should be an absolute path. It will not automatically
convert `root` to absolute path for performance reason.
The optional parameter `indexFiles` specifies the searching index files when the result is a directory.
For example, if the result `filePath` is a directory, and `indexFiles` is [index.html, main.html], it will also
search [index.html, main.html] under `filePath`. It returns the absolute file path if any of them found,
or else it returns `filePath`.

## 内部依赖

- `container/garray`
- `container/gmap`
- `errors/gcode`
- `errors/gerror`
- `internal/intlog`
- `os/gfile`
- `os/gfsnotify`
- `text/gstr`

