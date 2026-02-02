# gfsnotify

> Package: `github.com/gogf/gf/v2/os/gfsnotify`

```go
import "github.com/gogf/gf/v2/os/gfsnotify"
```

## 概述

Package gfsnotify provides a platform-independent interface for file system notifications.

## 源文件

- `gfsnotify.go`
- `gfsnotify_event.go`
- `gfsnotify_filefunc.go`
- `gfsnotify_watcher.go`
- `gfsnotify_watcher_loop.go`

## 类型定义

### Watcher

**类型**: struct

Watcher is the monitor for file changes.


**方法**:

- `func (w *Watcher) Add(path string, callbackFunc func, option ...WatchOption) (callback *Callback, err error)`
  Add monitors `path` with callback function `callbackFunc` to the watcher.
- `func (w *Watcher) AddOnce(name string, path string, callbackFunc func, option ...WatchOption) (callback *Callback, err error)`
  AddOnce monitors `path` with callback function `callbackFunc` only once using...
- `func (w *Watcher) Close()`
  Close closes the watcher.
- `func (w *Watcher) Remove(path string) error`
  Remove removes watching and all callbacks associated with the `path` recursiv...
- `func (w *Watcher) RemoveCallback(callbackID int)`
  RemoveCallback removes callback with given callback id from watcher.

### Callback

**类型**: struct

Callback is the callback function for Watcher.


### Event

**类型**: struct

Event is the event produced by underlying fsnotify.


**方法**:

- `func (e *Event) String() string`
  String returns current event as string.
- `func (e *Event) IsCreate() bool`
  IsCreate checks whether current event contains file/folder create event.
- `func (e *Event) IsWrite() bool`
  IsWrite checks whether current event contains file/folder write event.
- `func (e *Event) IsRemove() bool`
  IsRemove checks whether current event contains file/folder remove event.
- `func (e *Event) IsRename() bool`
  IsRename checks whether current event contains file/folder rename event.
- `func (e *Event) IsChmod() bool`
  IsChmod checks whether current event contains file/folder chmod event.

### WatchOption

**类型**: struct

WatchOption holds the option for watching.


### Op

**类型**: type

Op is the bits union for file operations.


## 函数

### New

```go
func New() (*Watcher, error)
```

New creates and returns a new watcher.
Note that the watcher number is limited by the file handle setting of the system.
Example: fs.inotify.max_user_instances system variable in linux systems.

In most case, you can use the default watcher for usage instead of creating one.

### Add

```go
func Add(path string, callbackFunc func, option ...WatchOption) (callback *Callback, err error)
```

Add monitors `path` using default watcher with callback function `callbackFunc`.

The parameter `path` can be either a file or a directory path.
The optional parameter `recursive` specifies whether monitoring the `path` recursively, which is true in default.

### AddOnce

```go
func AddOnce(name string, path string, callbackFunc func, option ...WatchOption) (callback *Callback, err error)
```

AddOnce monitors `path` using default watcher with callback function `callbackFunc` only once using unique name `name`.

If AddOnce is called multiple times with the same `name` parameter, `path` is only added to monitor once.
It returns error if it's called twice with the same `name`.

The parameter `path` can be either a file or a directory path.
The optional parameter `recursive` specifies whether monitoring the `path` recursively, which is true in default.

### Remove

```go
func Remove(path string) error
```

Remove removes all monitoring callbacks of given `path` from watcher recursively.

### RemoveCallback

```go
func RemoveCallback(callbackID int) error
```

RemoveCallback removes specified callback with given id from watcher.

### Exit

```go
func Exit()
```

Exit is only used in the callback function, which can be used to remove current callback
of itself from the watcher.

## 内部依赖

- `container/glist`
- `container/gmap`
- `container/gqueue`
- `container/gset`
- `container/gtype`
- `errors/gcode`
- `errors/gerror`
- `internal/intlog`
- `os/gcache`

## 外部依赖

- `github.com/fsnotify/fsnotify`

