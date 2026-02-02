# gfile

> Package: `github.com/gogf/gf/v2/os/gfile`

```go
import "github.com/gogf/gf/v2/os/gfile"
```

## 概述

Package gfile provides easy-to-use operations for file system.

## 源文件

- `gfile.go`
- `gfile_cache.go`
- `gfile_contents.go`
- `gfile_copy.go`
- `gfile_home.go`
- `gfile_match.go`
- `gfile_replace.go`
- `gfile_scan.go`
- `gfile_search.go`
- `gfile_size.go`
- `gfile_sort.go`
- `gfile_source.go`
- `gfile_time.go`

## 类型定义

### CopyOption

**类型**: struct

CopyOption is the option for Copy* functions.


## 函数

### Mkdir

```go
func Mkdir(path string) err error
```

Mkdir creates directories recursively with given `path`.
The parameter `path` is suggested to be an absolute path instead of relative one.

### Create

```go
func Create(path string) (*os.File, error)
```

Create creates a file with given `path` recursively.
The parameter `path` is suggested to be absolute path.

### Open

```go
func Open(path string) (*os.File, error)
```

Open opens file/directory READONLY.

### OpenFile

```go
func OpenFile(path string, flag int, perm os.FileMode) (*os.File, error)
```

OpenFile opens file/directory with custom `flag` and `perm`.
The parameter `flag` is like: O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC, etc.

### OpenWithFlag

```go
func OpenWithFlag(path string, flag int) (*os.File, error)
```

OpenWithFlag opens file/directory with default perm and custom `flag`.
The default `perm` is 0666.
The parameter `flag` is like: O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC, etc.

### OpenWithFlagPerm

```go
func OpenWithFlagPerm(path string, flag int, perm os.FileMode) (*os.File, error)
```

OpenWithFlagPerm opens file/directory with custom `flag` and `perm`.
The parameter `flag` is like: O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC, etc.
The parameter `perm` is like: 0600, 0666, 0777, etc.

### Join

```go
func Join(paths ...string) string
```

Join joins string array paths with file separator of current system.

### Exists

```go
func Exists(path string) bool
```

Exists checks whether given `path` exist.

### IsDir

```go
func IsDir(path string) bool
```

IsDir checks whether given `path` a directory.
Note that it returns false if the `path` does not exist.

### Pwd

```go
func Pwd() string
```

Pwd returns absolute path of current working directory.
Note that it returns an empty string if retrieving current
working directory failed.

### Chdir

```go
func Chdir(dir string) err error
```

Chdir changes the current working directory to the named directory.
If there is an error, it will be of type *PathError.

### IsFile

```go
func IsFile(path string) bool
```

IsFile checks whether given `path` a file, which means it's not a directory.
Note that it returns false if the `path` does not exist.

### Stat

```go
func Stat(path string) (os.FileInfo, error)
```

Stat returns a FileInfo describing the named file.
If there is an error, it will be of type *PathError.

### Move

```go
func Move(src string, dst string) err error
```

Move renames (moves) `src` to `dst` path.
If `dst` already exists and is not a directory, it'll be replaced.

### Rename

```go
func Rename(src string, dst string) error
```

Rename is alias of Move.
See Move.

### DirNames

```go
func DirNames(path string) ([]string, error)
```

DirNames returns sub-file names of given directory `path`.
Note that the returned names are NOT absolute paths.

### Glob

```go
func Glob(pattern string, onlyNames ...bool) ([]string, error)
```

Glob returns the names of all files matching pattern or nil
if there is no matching file. The syntax of patterns is the same
as in Match. The pattern may describe hierarchical names such as
/usr/*/bin/ed (assuming the Separator is '/').

Glob ignores file system errors such as I/O errors reading directories.
The only possible returned error is ErrBadPattern, when pattern
is malformed.

### Remove

```go
func Remove(path string) err error
```

Remove deletes all file/directory with `path` parameter.
If parameter `path` is directory, it deletes it recursively.

It does nothing if given `path` does not exist or is empty.

Deprecated:
As the name Remove for files deleting is ambiguous,
please use RemoveFile or RemoveAll for explicit usage instead.

### RemoveFile

```go
func RemoveFile(path string) err error
```

RemoveFile removes the named file or (empty) directory.

### RemoveAll

```go
func RemoveAll(path string) err error
```

RemoveAll removes path and any children it contains.
It removes everything it can but returns the first error
it encounters. If the path does not exist, RemoveAll
returns nil (no error).

### IsReadable

```go
func IsReadable(path string) bool
```

IsReadable checks whether given `path` is readable.

### IsWritable

```go
func IsWritable(path string) bool
```

IsWritable checks whether given `path` is writable.

TODO improve performance; use golang.org/x/sys to cross-plat-form

### Chmod

```go
func Chmod(path string, mode os.FileMode) err error
```

Chmod is alias of os.Chmod.
See os.Chmod.

### Abs

```go
func Abs(path string) string
```

Abs returns an absolute representation of path.
If the path is not absolute it will be joined with the current
working directory to turn it into an absolute path. The absolute
path name for a given file is not guaranteed to be unique.
Abs calls Clean on the result.

### RealPath

```go
func RealPath(path string) string
```

RealPath converts the given `path` to its absolute path
and checks if the file path exists.
If the file does not exist, return an empty string.

### SelfPath

```go
func SelfPath() string
```

SelfPath returns absolute file path of current running process(binary).

### SelfName

```go
func SelfName() string
```

SelfName returns file name of current running process(binary).

### SelfDir

```go
func SelfDir() string
```

SelfDir returns absolute directory path of current running process(binary).

### Basename

```go
func Basename(path string) string
```

Basename returns the last element of path, which contains file extension.
Trailing path separators are removed before extracting the last element.
If the path is empty, Base returns ".".
If the path consists entirely of separators, Basename returns a single separator.

Example:
Basename("/var/www/file.js") -> file.js
Basename("file.js")          -> file.js

### Name

```go
func Name(path string) string
```

Name returns the last element of path without file extension.

Example:
Name("/var/www/file.js") -> file
Name("file.js")          -> file

### Dir

```go
func Dir(path string) string
```

Dir returns all but the last element of path, typically the path's directory.
After dropping the final element, Dir calls Clean on the path and trailing
slashes are removed.
If the `path` is empty, Dir returns ".".
If the `path` is ".", Dir treats the path as current working directory.
If the `path` consists entirely of separators, Dir returns a single separator.
The returned path does not end in a separator unless it is the root directory.

Example:
Dir("/var/www/file.js") -> "/var/www"
Dir("file.js")          -> "."

### IsEmpty

```go
func IsEmpty(path string) bool
```

IsEmpty checks whether the given `path` is empty.
If `path` is a folder, it checks if there's any file under it.
If `path` is a file, it checks if the file size is zero.

Note that it returns true if `path` does not exist.

### Ext

```go
func Ext(path string) string
```

Ext returns the file name extension used by path.
The extension is the suffix beginning at the final dot
in the final element of path; it is empty if there is
no dot.
Note: the result contains symbol '.'.

Example:
Ext("main.go")  => .go
Ext("api.json") => .json

### ExtName

```go
func ExtName(path string) string
```

ExtName is like function Ext, which returns the file name extension used by path,
but the result does not contain symbol '.'.

Example:
ExtName("main.go")  => go
ExtName("api.json") => json

### Temp

```go
func Temp(names ...string) string
```

Temp retrieves and returns the temporary directory of current system.

The optional parameter `names` specifies the sub-folders/sub-files,
which will be joined with current system separator and returned with the path.

### GetContentsWithCache

```go
func GetContentsWithCache(path string, duration ...time.Duration) string
```

GetContentsWithCache returns string content of given file by `path` from cache.
If there's no content in the cache, it will read it from disk file specified by `path`.
The parameter `expire` specifies the caching time for this file content in seconds.

### GetBytesWithCache

```go
func GetBytesWithCache(path string, duration ...time.Duration) []byte
```

GetBytesWithCache returns []byte content of given file by `path` from cache.
If there's no content in the cache, it will read it from disk file specified by `path`.
The parameter `expire` specifies the caching time for this file content in seconds.

### GetContents

```go
func GetContents(path string) string
```

GetContents returns the file content of `path` as string.
It returns en empty string if it fails reading.

### GetBytes

```go
func GetBytes(path string) []byte
```

GetBytes returns the file content of `path` as []byte.
It returns nil if it fails reading.

### Truncate

```go
func Truncate(path string, size int) err error
```

Truncate truncates file of `path` to given size by `size`.

### PutContents

```go
func PutContents(path string, content string) error
```

PutContents puts string `content` to file of `path`.
It creates file of `path` recursively if it does not exist.

### PutContentsAppend

```go
func PutContentsAppend(path string, content string) error
```

PutContentsAppend appends string `content` to file of `path`.
It creates file of `path` recursively if it does not exist.

### PutBytes

```go
func PutBytes(path string, content []byte) error
```

PutBytes puts binary `content` to file of `path`.
It creates file of `path` recursively if it does not exist.

### PutBytesAppend

```go
func PutBytesAppend(path string, content []byte) error
```

PutBytesAppend appends binary `content` to file of `path`.
It creates file of `path` recursively if it does not exist.

### GetNextCharOffset

```go
func GetNextCharOffset(reader io.ReaderAt, char byte, start int64) int64
```

GetNextCharOffset returns the file offset for given `char` starting from `start`.

### GetNextCharOffsetByPath

```go
func GetNextCharOffsetByPath(path string, char byte, start int64) int64
```

GetNextCharOffsetByPath returns the file offset for given `char` starting from `start`.
It opens file of `path` for reading with os.O_RDONLY flag and default perm.

### GetBytesTilChar

```go
func GetBytesTilChar(reader io.ReaderAt, char byte, start int64) ([]byte, int64)
```

GetBytesTilChar returns the contents of the file as []byte
until the next specified byte `char` position.

Note: Returned value contains the character of the last position.

### GetBytesTilCharByPath

```go
func GetBytesTilCharByPath(path string, char byte, start int64) ([]byte, int64)
```

GetBytesTilCharByPath returns the contents of the file given by `path` as []byte
until the next specified byte `char` position.
It opens file of `path` for reading with os.O_RDONLY flag and default perm.

Note: Returned value contains the character of the last position.

### GetBytesByTwoOffsets

```go
func GetBytesByTwoOffsets(reader io.ReaderAt, start int64, end int64) []byte
```

GetBytesByTwoOffsets returns the binary content as []byte from `start` to `end`.
Note: Returned value does not contain the character of the last position, which means
it returns content range as [start, end).

### GetBytesByTwoOffsetsByPath

```go
func GetBytesByTwoOffsetsByPath(path string, start int64, end int64) []byte
```

GetBytesByTwoOffsetsByPath returns the binary content as []byte from `start` to `end`.
Note: Returned value does not contain the character of the last position, which means
it returns content range as [start, end).
It opens file of `path` for reading with os.O_RDONLY flag and default perm.

### ReadLines

```go
func ReadLines(file string, callback func) error
```

ReadLines reads file content line by line, which is passed to the callback function `callback` as string.
It matches each line of text, separated by chars '\r' or '\n', stripped any trailing end-of-line marker.

Note that the parameter passed to callback function might be an empty value, and the last non-empty line
will be passed to callback function `callback` even if it has no newline marker.

### ReadLinesBytes

```go
func ReadLinesBytes(file string, callback func) error
```

ReadLinesBytes reads file content line by line, which is passed to the callback function `callback` as []byte.
It matches each line of text, separated by chars '\r' or '\n', stripped any trailing end-of-line marker.

Note that the parameter passed to callback function might be an empty value, and the last non-empty line
will be passed to callback function `callback` even if it has no newline marker.

### Copy

```go
func Copy(src string, dst string, option ...CopyOption) error
```

Copy file/directory from `src` to `dst`.

If `src` is file, it calls CopyFile to implements copy feature,
or else it calls CopyDir.

If `src` is file, but `dst` already exists and is a folder,
it then creates a same name file of `src` in folder `dst`.

Eg:
Copy("/tmp/file1", "/tmp/file2") => /tmp/file1 copied to /tmp/file2
Copy("/tmp/dir1",  "/tmp/dir2")  => /tmp/dir1  copied to /tmp/dir2
Copy("/tmp/file1", "/tmp/dir2")  => /tmp/file1 copied to /tmp/dir2/file1
Copy("/tmp/dir1",  "/tmp/file2") => error

### CopyFile

```go
func CopyFile(src string, dst string, option ...CopyOption) err error
```

CopyFile copies the contents of the file named `src` to the file named
by `dst`. The file will be created if it does not exist. If the
destination file exists, all it's contents will be replaced by the contents
of the source file. The file mode will be copied from the source and
the copied data is synced/flushed to stable storage.
Thanks: https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04

### CopyDir

```go
func CopyDir(src string, dst string, option ...CopyOption) err error
```

CopyDir recursively copies a directory tree, attempting to preserve permissions.

Note that, the Source directory must exist and symlinks are ignored and skipped.

### Home

```go
func Home(names ...string) (string, error)
```

Home returns absolute path of current user's home directory.
The optional parameter `names` specifies the sub-folders/sub-files,
which will be joined with current system separator and returned with the path.

### MatchGlob

```go
func MatchGlob(pattern string, name string) (bool, error)
```

MatchGlob reports whether name matches the shell pattern.
It extends filepath.Match (https://pkg.go.dev/path/filepath#Match)
with support for "**" (globstar) pattern, similar to bash's globstar
(https://www.gnu.org/software/bash/manual/html_node/The-Shopt-Builtin.html)
and gitignore patterns (https://git-scm.com/docs/gitignore#_pattern_format).

Pattern syntax:
  - '*'         matches any sequence of non-separator characters
  - '**'        matches any sequence of characters including separators (globstar)
  - '?'         matches any single non-separator character
  - '[abc]'     matches any character in the bracket
  - '[a-z]'     matches any character in the range
  - '[^abc]'    matches any character not in the bracket (negation)
  - '[^a-z]'    matches any character not in the range (negation)

Globstar rules:
  - "**" only has globstar semantics when it appears as a complete path component
    (e.g., "a/**/b", "**/a", "a/**", "**").
  - Patterns like "a**b" or "**a" treat "**" as two regular "*" wildcards,
    matching only within a single path component.
  - Both "/" and "\" are treated as path separators (cross-platform support).

Error handling:
  - Returns an error for malformed patterns (e.g., unclosed brackets "[abc").
  - Errors from filepath.Match are propagated.

Example:

	MatchGlob("src/**/*.go", "src/foo/bar/main.go")  => true, nil
	MatchGlob("*.go", "main.go")                     => true, nil
	MatchGlob("**", "any/path/file.go")              => true, nil
	MatchGlob("a**b", "axxb")                        => true, nil  (** as two *)
	MatchGlob("a**b", "a/b")                         => false, nil (no separator match)
	MatchGlob("[abc]", "a")                          => true, nil
	MatchGlob("[", "a")                              => false, error (malformed)

### ReplaceFile

```go
func ReplaceFile(search string, replace string, path string) error
```

ReplaceFile replaces content for file `path`.

### ReplaceFileFunc

```go
func ReplaceFileFunc(f func, path string) error
```

ReplaceFileFunc replaces content for file `path` with callback function `f`.

### ReplaceDir

```go
func ReplaceDir(search string, replace string, path string, pattern string, recursive ...bool) error
```

ReplaceDir replaces content for files under `path`.
The parameter `pattern` specifies the file pattern which matches to be replaced.
It does replacement recursively if given parameter `recursive` is true.

### ReplaceDirFunc

```go
func ReplaceDirFunc(f func, path string, pattern string, recursive ...bool) error
```

ReplaceDirFunc replaces content for files under `path` with callback function `f`.
The parameter `pattern` specifies the file pattern which matches to be replaced.
It does replacement recursively if given parameter `recursive` is true.

### ScanDir

```go
func ScanDir(path string, pattern string, recursive ...bool) ([]string, error)
```

ScanDir returns all sub-files with absolute paths of given `path`,
It scans directory recursively if given parameter `recursive` is true.

The pattern parameter `pattern` supports multiple file name patterns,
using the ',' symbol to separate multiple patterns.

### ScanDirFunc

```go
func ScanDirFunc(path string, pattern string, recursive bool, handler func) ([]string, error)
```

ScanDirFunc returns all sub-files with absolute paths of given `path`,
It scans directory recursively if given parameter `recursive` is true.

The pattern parameter `pattern` supports multiple file name patterns, using the ','
symbol to separate multiple patterns.

The parameter `recursive` specifies whether scanning the `path` recursively, which
means it scans its sub-files and appends the files path to result array if the sub-file
is also a folder. It is false in default.

The parameter `handler` specifies the callback function handling each sub-file path of
the `path` and its sub-folders. It ignores the sub-file path if `handler` returns an empty
string, or else it appends the sub-file path to result slice.

### ScanDirFile

```go
func ScanDirFile(path string, pattern string, recursive ...bool) ([]string, error)
```

ScanDirFile returns all sub-files with absolute paths of given `path`,
It scans directory recursively if given parameter `recursive` is true.

The pattern parameter `pattern` supports multiple file name patterns,
using the ',' symbol to separate multiple patterns.

Note that it returns only files, exclusive of directories.

### ScanDirFileFunc

```go
func ScanDirFileFunc(path string, pattern string, recursive bool, handler func) ([]string, error)
```

ScanDirFileFunc returns all sub-files with absolute paths of given `path`,
It scans directory recursively if given parameter `recursive` is true.

The pattern parameter `pattern` supports multiple file name patterns, using the ','
symbol to separate multiple patterns.

The parameter `recursive` specifies whether scanning the `path` recursively, which
means it scans its sub-files and appends the file paths to result array if the sub-file
is also a folder. It is false in default.

The parameter `handler` specifies the callback function handling each sub-file path of
the `path` and its sub-folders. It ignores the sub-file path if `handler` returns an empty
string, or else it appends the sub-file path to result slice.

Note that the parameter `path` for `handler` is not a directory but a file.
It returns only files, exclusive of directories.

### Search

```go
func Search(name string, prioritySearchPaths ...string) (realPath string, err error)
```

Search searches file by name `name` in following paths with priority:
prioritySearchPaths, Pwd()、SelfDir()、MainPkgPath().
It returns the absolute file path of `name` if found, or en empty string if not found.

### Size

```go
func Size(path string) int64
```

Size returns the size of file specified by `path` in byte.

### SizeFormat

```go
func SizeFormat(path string) string
```

SizeFormat returns the size of file specified by `path` in format string.

### ReadableSize

```go
func ReadableSize(path string) string
```

ReadableSize formats size of file given by `path`, for more human readable.

### StrToSize

```go
func StrToSize(sizeStr string) int64
```

StrToSize converts formatted size string to its size in bytes.

### FormatSize

```go
func FormatSize(raw int64) string
```

FormatSize formats size `raw` for more manually readable.

### SortFiles

```go
func SortFiles(files []string) []string
```

SortFiles sorts the `files` in order of: directory -> file.
Note that the item of `files` should be absolute path.

### MainPkgPath

```go
func MainPkgPath() string
```

MainPkgPath returns absolute file path of package main,
which contains the entrance function main.

It's only available in develop environment.

Note1: Only valid for source development environments,
IE only valid for systems that generate this executable.

Note2: When the method is called for the first time, if it is in an asynchronous goroutine,
the method may not get the main package path.

### MTime

```go
func MTime(path string) time.Time
```

MTime returns the modification time of file given by `path` in second.

### MTimestamp

```go
func MTimestamp(path string) int64
```

MTimestamp returns the modification time of file given by `path` in second.

### MTimestampMilli

```go
func MTimestampMilli(path string) int64
```

MTimestampMilli returns the modification time of file given by `path` in millisecond.

## 内部依赖

- `container/garray`
- `container/gtype`
- `errors/gcode`
- `errors/gerror`
- `internal/command`
- `internal/intlog`
- `os/gcache`
- `os/gfsnotify`
- `text/gregex`
- `text/gstr`
- `util/gconv`

