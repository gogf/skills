# gproc

> Package: `github.com/gogf/gf/v2/os/gproc`

```go
import "github.com/gogf/gf/v2/os/gproc"
```

## 概述

Package gproc implements management and communication for processes.

## 源文件

- `gproc.go`
- `gproc_comm.go`
- `gproc_comm_receive.go`
- `gproc_comm_send.go`
- `gproc_manager.go`
- `gproc_must.go`
- `gproc_process.go`
- `gproc_process_newprocess.go`
- `gproc_shell.go`
- `gproc_signal.go`

## 类型定义

### MsgRequest

**类型**: struct

MsgRequest is the request structure for process communication.


### MsgResponse

**类型**: struct

MsgResponse is the response structure for process communication.


### Manager

**类型**: struct

Manager is a process manager maintaining multiple processes.


**方法**:

- `func (m *Manager) NewProcess(path string, args []string, environment []string) *Process`
  NewProcess creates and returns a Process object.
- `func (m *Manager) GetProcess(pid int) *Process`
  GetProcess retrieves and returns a Process object.
- `func (m *Manager) AddProcess(pid int)`
  AddProcess adds a process to current manager.
- `func (m *Manager) RemoveProcess(pid int)`
  RemoveProcess removes a process from current manager.
- `func (m *Manager) Processes() []*Process`
  Processes retrieves and returns all processes in current manager.
- `func (m *Manager) Pids() []int`
  Pids retrieves and returns all process id array in current manager.
- `func (m *Manager) WaitAll()`
  WaitAll waits until all process exit.
- `func (m *Manager) KillAll() error`
  KillAll kills all processes in current manager.
- `func (m *Manager) SignalAll(sig os.Signal) error`
  SignalAll sends a signal `sig` to all processes in current manager.
- `func (m *Manager) Send(data []byte)`
  Send sends data bytes to all processes in current manager.
- `func (m *Manager) SendTo(pid int, data []byte) error`
  SendTo sneds data bytes to specified processe in current manager.
- `func (m *Manager) Clear()`
  Clear removes all processes in current manager.
- `func (m *Manager) Size() int`
  Size returns the size of processes in current manager.

### Process

**类型**: struct

Process is the struct for a single process.


**方法**:

- `func (p *Process) Start(ctx context.Context) (int, error)`
  Start starts executing the process in non-blocking way.
- `func (p *Process) Run(ctx context.Context) error`
  Run executes the process in blocking way.
- `func (p *Process) Pid() int`
  Pid retrieves and returns the PID for the process.
- `func (p *Process) Send(data []byte) error`
  Send sends custom data to the process.
- `func (p *Process) Release() error`
  Release releases any resources associated with the Process p,
- `func (p *Process) Kill() err error`
  Kill causes the Process to exit immediately.
- `func (p *Process) Signal(sig os.Signal) error`
  Signal sends a signal to the Process.

### SigHandler

**类型**: type

SigHandler defines a function type for signal handling.


## 函数

### Pid

```go
func Pid() int
```

Pid returns the pid of current process.

### PPid

```go
func PPid() int
```

PPid returns the custom parent pid if exists, or else it returns the system parent pid.

### PPidOS

```go
func PPidOS() int
```

PPidOS returns the system parent pid of current process.
Note that the difference between PPidOS and PPid function is that the PPidOS returns
the system ppid, but the PPid functions may return the custom pid by gproc if the custom
ppid exists.

### IsChild

```go
func IsChild() bool
```

IsChild checks and returns whether current process is a child process.
A child process is forked by another gproc process.

### SetPPid

```go
func SetPPid(ppid int) error
```

SetPPid sets custom parent pid for current process.

### StartTime

```go
func StartTime() time.Time
```

StartTime returns the start time of current process.

### Uptime

```go
func Uptime() time.Duration
```

Uptime returns the duration which current process has been running

### SearchBinary

```go
func SearchBinary(file string) string
```

SearchBinary searches the binary `file` in current working folder and PATH environment.

### SearchBinaryPath

```go
func SearchBinaryPath(file string) string
```

SearchBinaryPath searches the binary `file` in PATH environment.

### Receive

```go
func Receive(group ...string) *MsgRequest
```

Receive blocks and receives message from other process using local TCP listening.
Note that, it only enables the TCP listening service when this function called.

### Send

```go
func Send(pid int, data []byte, group ...string) error
```

Send sends data to specified process of given pid.

### NewManager

```go
func NewManager() *Manager
```

NewManager creates and returns a new process manager.

### MustShell

```go
func MustShell(ctx context.Context, cmd string, out io.Writer, in io.Reader)
```

MustShell performs as Shell, but it panics if any error occurs.

### MustShellRun

```go
func MustShellRun(ctx context.Context, cmd string)
```

MustShellRun performs as ShellRun, but it panics if any error occurs.

### MustShellExec

```go
func MustShellExec(ctx context.Context, cmd string, environment ...[]string) string
```

MustShellExec performs as ShellExec, but it panics if any error occurs.

### NewProcess

```go
func NewProcess(path string, args []string, environment ...[]string) *Process
```

NewProcess creates and returns a new Process.

### NewProcessCmd

```go
func NewProcessCmd(cmd string, environment ...[]string) *Process
```

NewProcessCmd creates and returns a process with given command and optional environment variable array.

### Shell

```go
func Shell(ctx context.Context, cmd string, out io.Writer, in io.Reader) error
```

Shell executes command `cmd` synchronously with given input pipe `in` and output pipe `out`.
The command `cmd` reads the input parameters from input pipe `in`, and writes its output automatically
to output pipe `out`.

### ShellRun

```go
func ShellRun(ctx context.Context, cmd string) error
```

ShellRun executes given command `cmd` synchronously and outputs the command result to the stdout.

### ShellExec

```go
func ShellExec(ctx context.Context, cmd string, environment ...[]string) (result string, err error)
```

ShellExec executes given command `cmd` synchronously and returns the command result.

### AddSigHandler

```go
func AddSigHandler(handler SigHandler, signals ...os.Signal)
```

AddSigHandler adds custom signal handler for custom one or more signals.

### AddSigHandlerShutdown

```go
func AddSigHandlerShutdown(handler ...SigHandler)
```

AddSigHandlerShutdown adds custom signal handler for shutdown signals:
syscall.SIGINT,
syscall.SIGQUIT,
syscall.SIGKILL,
syscall.SIGTERM,
syscall.SIGABRT.

### Listen

```go
func Listen()
```

Listen blocks and does signal listening and handling.

## 内部依赖

- `github.com/gogf/gf/v2`
- `container/gmap`
- `container/gqueue`
- `container/gtype`
- `errors/gcode`
- `errors/gerror`
- `internal/intlog`
- `internal/json`
- `net/gtcp`
- `net/gtrace`
- `os/genv`
- `os/gfile`
- `os/glog`
- `text/gstr`
- `util/gconv`
- `util/gutil`

## 外部依赖

- `go.opentelemetry.io/otel`
- `go.opentelemetry.io/otel/propagation`
- `go.opentelemetry.io/otel/trace`

