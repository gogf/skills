# gcron

> Package: `github.com/gogf/gf/v2/os/gcron`

```go
import "github.com/gogf/gf/v2/os/gcron"
```

## 概述

Package gcron implements a cron pattern parser and job runner.

## 源文件

- `gcron.go`
- `gcron_cron.go`
- `gcron_entry.go`
- `gcron_schedule.go`
- `gcron_schedule_check.go`
- `gcron_schedule_fix.go`
- `gcron_schedule_next.go`

## 类型定义

### Cron

**类型**: struct

Cron stores all the cron job entries.


**方法**:

- `func (c *Cron) SetLogger(logger glog.ILogger)`
  SetLogger sets the logger for cron.
- `func (c *Cron) GetLogger() glog.ILogger`
  GetLogger returns the logger in the cron.
- `func (c *Cron) AddEntry(ctx context.Context, pattern string, job JobFunc, times int, isSingleton bool, name ...string) (*Entry, error)`
  AddEntry creates and returns a new Entry object.
- `func (c *Cron) Add(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error)`
  Add adds a timed task.
- `func (c *Cron) AddSingleton(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error)`
  AddSingleton adds a singleton timed task.
- `func (c *Cron) AddTimes(ctx context.Context, pattern string, times int, job JobFunc, name ...string) (*Entry, error)`
  AddTimes adds a timed task which can be run specified times.
- `func (c *Cron) AddOnce(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error)`
  AddOnce adds a timed task which can be run only once.
- `func (c *Cron) DelayAddEntry(ctx context.Context, delay time.Duration, pattern string, job JobFunc, times int, isSingleton bool, name ...string)`
  DelayAddEntry adds a timed task after `delay` time.
- `func (c *Cron) DelayAdd(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string)`
  DelayAdd adds a timed task after `delay` time.
- `func (c *Cron) DelayAddSingleton(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string)`
  DelayAddSingleton adds a singleton timed task after `delay` time.
- `func (c *Cron) DelayAddOnce(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string)`
  DelayAddOnce adds a timed task after `delay` time.
- `func (c *Cron) DelayAddTimes(ctx context.Context, delay time.Duration, pattern string, times int, job JobFunc, name ...string)`
  DelayAddTimes adds a timed task after `delay` time.
- `func (c *Cron) Search(name string) *Entry`
  Search returns a scheduled task with the specified `name`.
- `func (c *Cron) Start(name ...string)`
  Start starts running the specified timed task named `name`.
- `func (c *Cron) Stop(name ...string)`
  Stop stops running the specified timed task named `name`.
- `func (c *Cron) StopGracefully()`
  StopGracefully Blocks and waits all current running jobs done.
- `func (c *Cron) StopGracefullyNonBlocking() context.Context`
  StopGracefullyNonBlocking stops all running tasks gracefully without blocking,
- `func (c *Cron) Remove(name string)`
  Remove deletes scheduled task which named `name`.
- `func (c *Cron) Close()`
  Close stops and closes current cron.
- `func (c *Cron) Size() int`
  Size returns the size of the timed tasks.
- `func (c *Cron) Entries() []*Entry`
  Entries return all timed tasks as slice(order by registered time asc).

### JobFunc

**类型**: type

JobFunc is the timing called job function in cron.


### Entry

**类型**: struct

Entry is timing task entry.


**方法**:

- `func (e *Entry) IsSingleton() bool`
  IsSingleton return whether this entry is a singleton timed task.
- `func (e *Entry) SetSingleton(enabled bool)`
  SetSingleton sets the entry running in singleton mode.
- `func (e *Entry) SetTimes(times int)`
  SetTimes sets the times which the entry can run.
- `func (e *Entry) Status() int`
  Status returns the status of entry.
- `func (e *Entry) SetStatus(status int) int`
  SetStatus sets the status of the entry.
- `func (e *Entry) Start()`
  Start starts running the entry.
- `func (e *Entry) Stop()`
  Stop stops running the entry.
- `func (e *Entry) Close()`
  Close stops and removes the entry from cron.

## 函数

### SetLogger

```go
func SetLogger(logger glog.ILogger)
```

SetLogger sets the global logger for cron.

### GetLogger

```go
func GetLogger() glog.ILogger
```

GetLogger returns the global logger in the cron.

### Add

```go
func Add(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error)
```

Add adds a timed task to default cron object.
A unique `name` can be bound with the timed task.
It returns and error if the `name` is already used.

### AddSingleton

```go
func AddSingleton(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error)
```

AddSingleton adds a singleton timed task, to default cron object.
A singleton timed task is that can only be running one single instance at the same time.
A unique `name` can be bound with the timed task.
It returns and error if the `name` is already used.

### AddOnce

```go
func AddOnce(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error)
```

AddOnce adds a timed task which can be run only once, to default cron object.
A unique `name` can be bound with the timed task.
It returns and error if the `name` is already used.

### AddTimes

```go
func AddTimes(ctx context.Context, pattern string, times int, job JobFunc, name ...string) (*Entry, error)
```

AddTimes adds a timed task which can be run specified times, to default cron object.
A unique `name` can be bound with the timed task.
It returns and error if the `name` is already used.

### DelayAdd

```go
func DelayAdd(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string)
```

DelayAdd adds a timed task to default cron object after `delay` time.

### DelayAddSingleton

```go
func DelayAddSingleton(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string)
```

DelayAddSingleton adds a singleton timed task after `delay` time to default cron object.

### DelayAddOnce

```go
func DelayAddOnce(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string)
```

DelayAddOnce adds a timed task after `delay` time to default cron object.
This timed task can be run only once.

### DelayAddTimes

```go
func DelayAddTimes(ctx context.Context, delay time.Duration, pattern string, times int, job JobFunc, name ...string)
```

DelayAddTimes adds a timed task after `delay` time to default cron object.
This timed task can be run specified times.

### Search

```go
func Search(name string) *Entry
```

Search returns a scheduled task with the specified `name`.
It returns nil if no found.

### Remove

```go
func Remove(name string)
```

Remove deletes scheduled task which named `name`.

### Size

```go
func Size() int
```

Size returns the size of the timed tasks of default cron.

### Entries

```go
func Entries() []*Entry
```

Entries return all timed tasks as slice.

### Start

```go
func Start(name ...string)
```

Start starts running the specified timed task named `name`.
If no`name` specified, it starts the entire cron.

### Stop

```go
func Stop(name ...string)
```

Stop stops running the specified timed task named `name`.
If no`name` specified, it stops the entire cron.

### StopGracefully

```go
func StopGracefully()
```

StopGracefully Blocks and waits all current running jobs done.

### StopGracefullyNonBlocking

```go
func StopGracefullyNonBlocking() context.Context
```

StopGracefullyNonBlocking stops all running tasks gracefully without blocking,
returning a context that callers can use to wait for completion.

### New

```go
func New() *Cron
```

New returns a new Cron object with default settings.

## 内部依赖

- `container/garray`
- `container/gmap`
- `container/gtype`
- `errors/gcode`
- `errors/gerror`
- `internal/intlog`
- `os/glog`
- `os/gtime`
- `os/gtimer`
- `text/gregex`
- `util/gconv`

