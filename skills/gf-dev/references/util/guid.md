# guid

> Package: `github.com/gogf/gf/v2/util/guid`

```go
import "github.com/gogf/gf/v2/util/guid"
```

## 概述

Package guid provides simple and high performance unique id generation functionality.

## 源文件

- `guid.go`

## 函数

### S

```go
func S(data ...[]byte) string
```

S creates and returns a global unique string in 32 bytes that meets most common
usages without strict UUID algorithm. It returns a unique string using default
unique algorithm if no `data` is given.

The specified `data` can be no more than 2 parts. No matter how long each of the
`data` size is, each of them will be hashed into 7 bytes as part of the result.
If given `data` parts is less than 2, the leftover size of the result bytes will
be token by random string.

The returned string is composed with:
1. Default:    MACHash(7) + PID(4) + TimestampNano(12) + Sequence(3) + RandomString(6)
2. CustomData: DataHash(7/14) + TimestampNano(12) + Sequence(3) + RandomString(3/10)

Note that：
 1. The returned length is fixed to 32 bytes for performance purpose.
 2. The custom parameter `data` composed should have unique attribute in your
    business scenario.

## 内部依赖

- `container/gtype`
- `encoding/ghash`
- `errors/gcode`
- `errors/gerror`
- `net/gipv4`
- `util/grand`

