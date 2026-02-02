# gcompress

> Package: `github.com/gogf/gf/v2/encoding/gcompress`

```go
import "github.com/gogf/gf/v2/encoding/gcompress"
```

## 概述

Package gcompress provides kinds of compression algorithms for binary/bytes data.

## 源文件

- `gcompress.go`
- `gcompress_gzip.go`
- `gcompress_zip.go`
- `gcompress_zlib.go`

## 函数

### Gzip

```go
func Gzip(data []byte, level ...int) ([]byte, error)
```

Gzip compresses `data` using gzip algorithm.
The optional parameter `level` specifies the compression level from
1 to 9 which means from none to the best compression.

Note that it returns error if given `level` is invalid.

### GzipFile

```go
func GzipFile(srcFilePath string, dstFilePath string, level ...int) err error
```

GzipFile compresses the file `src` to `dst` using gzip algorithm.

### GzipPathWriter

```go
func GzipPathWriter(filePath string, writer io.Writer, level ...int) error
```

GzipPathWriter compresses `filePath` to `writer` using gzip compressing algorithm.

Note that the parameter `path` can be either a directory or a file.

### UnGzip

```go
func UnGzip(data []byte) ([]byte, error)
```

UnGzip decompresses `data` with gzip algorithm.

### UnGzipFile

```go
func UnGzipFile(srcFilePath string, dstFilePath string) error
```

UnGzipFile decompresses srcFilePath `src` to `dst` using gzip algorithm.

### ZipPath

```go
func ZipPath(fileOrFolderPaths string, dstFilePath string, prefix ...string) error
```

ZipPath compresses `fileOrFolderPaths` to `dstFilePath` using zip compressing algorithm.

The parameter `paths` can be either a directory or a file, which
supports multiple paths join with ','.
The unnecessary parameter `prefix` indicates the path prefix for zip file.

### ZipPathWriter

```go
func ZipPathWriter(fileOrFolderPaths string, writer io.Writer, prefix ...string) error
```

ZipPathWriter compresses `fileOrFolderPaths` to `writer` using zip compressing algorithm.

Note that the parameter `fileOrFolderPaths` can be either a directory or a file, which
supports multiple paths join with ','.
The unnecessary parameter `prefix` indicates the path prefix for zip file.

### ZipPathContent

```go
func ZipPathContent(fileOrFolderPaths string, prefix ...string) ([]byte, error)
```

ZipPathContent compresses `fileOrFolderPaths` to []byte using zip compressing algorithm.

Note that the parameter `fileOrFolderPaths` can be either a directory or a file, which
supports multiple paths join with ','.
The unnecessary parameter `prefix` indicates the path prefix for zip file.

### UnZipFile

```go
func UnZipFile(zippedFilePath string, dstFolderPath string, zippedPrefix ...string) error
```

UnZipFile decompresses `archive` to `dstFolderPath` using zip compressing algorithm.

The parameter `dstFolderPath` should be a directory.
The optional parameter `zippedPrefix` specifies the unzipped path of `zippedFilePath`,
which can be used to specify part of the archive file to unzip.

### UnZipContent

```go
func UnZipContent(zippedContent []byte, dstFolderPath string, zippedPrefix ...string) error
```

UnZipContent decompresses `zippedContent` to `dstFolderPath` using zip compressing algorithm.

The parameter `dstFolderPath` should be a directory.
The parameter `zippedPrefix` specifies the unzipped path of `zippedContent`,
which can be used to specify part of the archive file to unzip.

### Zlib

```go
func Zlib(data []byte) ([]byte, error)
```

Zlib compresses `data` with zlib algorithm.

### UnZlib

```go
func UnZlib(data []byte) ([]byte, error)
```

UnZlib decompresses `data` with zlib algorithm.

## 内部依赖

- `errors/gerror`
- `internal/intlog`
- `os/gfile`
- `text/gstr`

