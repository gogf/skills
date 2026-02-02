# gtimer

> Package: `github.com/gogf/gf/v2/os/gtimer`

```go
import "github.com/gogf/gf/v2/os/gtimer"
```

## 概述

Package gtimer implements timer for interval/delayed jobs running and management.

## 源文件

- `gtimer.go`
- `gtimer_entry.go`
- `gtimer_exit.go`
- `gtimer_queue.go`
- `gtimer_queue_heap.go`
- `gtimer_timer.go`
- `gtimer_timer_loop.go`

## 类型定义

### Timer

**类型**: struct

Timer is the timer manager, which uses ticks to calculate the timing interval.


**方法**:

- `func (t *Timer) Add(ctx context.Context, interval time.Duration, job JobFunc) *Entry`
  Add adds a timing job to the timer, which runs in interval of `interval`.
- `func (t *Timer) AddEntry(ctx context.Context, interval time.Duration, job JobFunc, isSingleton bool, times int, status int) *Entry`
  AddEntry adds a timing job to the timer with detailed parameters.
- `func (t *Timer) AddSingleton(ctx context.Context, interval time.Duration, job JobFunc) *Entry`
  AddSingleton is a convenience function for add singleton mode job.
- `func (t *Timer) AddOnce(ctx context.Context, interval time.Duration, job JobFunc) *Entry`
  AddOnce is a convenience function for adding a job which only runs once and t...
- `func (t *Timer) AddTimes(ctx context.Context, interval time.Duration, times int, job JobFunc) *Entry`
  AddTimes is a convenience function for adding a job which is limited running ...
- `func (t *Timer) DelayAdd(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc)`
  DelayAdd adds a timing job after delay of `delay` duration.
- `func (t *Timer) DelayAddEntry(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc, isSingleton bool, times int, status int)`
  DelayAddEntry adds a timing job after delay of `delay` duration.
- `func (t *Timer) DelayAddSingleton(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc)`
  DelayAddSingleton adds a timing job after delay of `delay` duration.
- `func (t *Timer) DelayAddOnce(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc)`
  DelayAddOnce adds a timing job after delay of `delay` duration.
- `func (t *Timer) DelayAddTimes(ctx context.Context, delay time.Duration, interval time.Duration, times int, job JobFunc)`
  DelayAddTimes adds a timing job after delay of `delay` duration.
- `func (t *Timer) Start()`
  Start starts the timer.
- `func (t *Timer) Stop()`
  Stop stops the timer.
- `func (t *Timer) Close()`
  Close closes the timer.

### TimerOptions

**类型**: struct

TimerOptions is the configuration object for Timer.


### Entry

**类型**: struct

Entry is the timing job.


**方法**:

- `func (entry *Entry) Status() int`
  Status returns the status of the job.
- `func (entry *Entry) Run()`
  Run runs the timer job asynchronously.
- `func (entry *Entry) SetStatus(status int) int`
  SetStatus custom sets the status for the job.
- `func (entry *Entry) Start()`
  Start starts the job.
- `func (entry *Entry) Stop()`
  Stop stops the job.
- `func (entry *Entry) Close()`
  Close closes the job, and then it will be removed from the timer.
- `func (entry *Entry) Reset()`
  Reset resets the job, which resets its ticks for next running.
- `func (entry *Entry) IsSingleton() bool`
  IsSingleton checks and returns whether the job in singleton mode.
- `func (entry *Entry) SetSingleton(enabled bool)`
  SetSingleton sets the job singleton mode.
- `func (entry *Entry) Job() JobFunc`
  Job returns the job function of this job.
- `func (entry *Entry) Ctx() context.Context`
  Ctx returns the initialized context of this job.
- `func (entry *Entry) SetTimes(times int)`
  SetTimes sets the limit running times for the job.

### JobFunc

**类型**: type

JobFunc is the timing called job function in timer.


## 函数

### DefaultOptions

```go
func DefaultOptions() TimerOptions
```

DefaultOptions creates and returns a default options object for Timer creation.

### SetTimeout

```go
func SetTimeout(ctx context.Context, delay time.Duration, job JobFunc)
```

SetTimeout runs the job once after duration of `delay`.
It is like the one in javascript.

### SetInterval

```go
func SetInterval(ctx context.Context, interval time.Duration, job JobFunc)
```

SetInterval runs the job every duration of `delay`.
It is like the one in javascript.

### Add

```go
func Add(ctx context.Context, interval time.Duration, job JobFunc) *Entry
```

Add adds a timing job to the default timer, which runs in interval of `interval`.

### AddEntry

```go
func AddEntry(ctx context.Context, interval time.Duration, job JobFunc, isSingleton bool, times int, status int) *Entry
```

AddEntry adds a timing job to the default timer with detailed parameters.

The parameter `interval` specifies the running interval of the job.

The parameter `singleton` specifies whether the job running in singleton mode.
There's only one of the same job is allowed running when its a singleton mode job.

The parameter `times` specifies limit for the job running times, which means the job
exits if its run times exceeds the `times`.

The parameter `status` specifies the job status when it's firstly added to the timer.

### AddSingleton

```go
func AddSingleton(ctx context.Context, interval time.Duration, job JobFunc) *Entry
```

AddSingleton is a convenience function for add singleton mode job.

### AddOnce

```go
func AddOnce(ctx context.Context, interval time.Duration, job JobFunc) *Entry
```

AddOnce is a convenience function for adding a job which only runs once and then exits.

### AddTimes

```go
func AddTimes(ctx context.Context, interval time.Duration, times int, job JobFunc) *Entry
```

AddTimes is a convenience function for adding a job which is limited running times.

### DelayAdd

```go
func DelayAdd(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc)
```

DelayAdd adds a timing job after delay of `interval` duration.
Also see Add.

### DelayAddEntry

```go
func DelayAddEntry(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc, isSingleton bool, times int, status int)
```

DelayAddEntry adds a timing job after delay of `interval` duration.
Also see AddEntry.

### DelayAddSingleton

```go
func DelayAddSingleton(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc)
```

DelayAddSingleton adds a timing job after delay of `interval` duration.
Also see AddSingleton.

### DelayAddOnce

```go
func DelayAddOnce(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc)
```

DelayAddOnce adds a timing job after delay of `interval` duration.
Also see AddOnce.

### DelayAddTimes

```go
func DelayAddTimes(ctx context.Context, delay time.Duration, interval time.Duration, times int, job JobFunc)
```

DelayAddTimes adds a timing job after delay of `interval` duration.
Also see AddTimes.

### Exit

```go
func Exit()
```

Exit is used in timing job internally, which exits and marks it closed from timer.
The timing job will be automatically removed from timer later. It uses "panic-recover"
mechanism internally implementing this feature, which is designed for simplification
and convenience.

### New

```go
func New(options ...TimerOptions) *Timer
```

New creates and returns a Timer.

## 内部依赖

- `container/gtype`
- `errors/gcode`
- `errors/gerror`
- `internal/command`

