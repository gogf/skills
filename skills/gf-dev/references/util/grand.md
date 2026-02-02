# grand

> Package: `github.com/gogf/gf/v2/util/grand`

```go
import "github.com/gogf/gf/v2/util/grand"
```

## 概述

Package grand provides high performance random bytes/number/string generation functionality.

## 源文件

- `grand.go`
- `grand_buffer.go`

## 函数

### Intn

```go
func Intn(max int) int
```

Intn returns an int number which is between 0 and max: [0, max).

Note that:
1. The `max` can only be greater than 0, or else it returns `max` directly;
2. The result is greater than or equal to 0, but less than `max`;
3. The result number is 32bit and less than math.MaxUint32.

### B

```go
func B(n int) []byte
```

B retrieves and returns random bytes of given length `n`.

### N

```go
func N(min int, max int) int
```

N returns a random int between min and max: [min, max].
The `min` and `max` also support negative numbers.

### S

```go
func S(n int, symbols ...bool) string
```

S returns a random string which contains digits and letters, and its length is `n`.
The optional parameter `symbols` specifies whether the result could contain symbols,
which is false in default.

### D

```go
func D(min time.Duration, max time.Duration) time.Duration
```

D returns a random time.Duration between min and max: [min, max].

### Str

```go
func Str(s string, n int) string
```

Str randomly picks and returns `n` count of chars from given string `s`.
It also supports unicode string like Chinese/Russian/Japanese, etc.

### Digits

```go
func Digits(n int) string
```

Digits returns a random string which contains only digits, and its length is `n`.

### Letters

```go
func Letters(n int) string
```

Letters returns a random string which contains only letters, and its length is `n`.

### Symbols

```go
func Symbols(n int) string
```

Symbols returns a random string which contains only symbols, and its length is `n`.

### Perm

```go
func Perm(n int) []int
```

Perm returns, as a slice of n int numbers, a pseudo-random permutation of the integers [0,n).
TODO performance improving for large slice producing.

### Meet

```go
func Meet(num int, total int) bool
```

Meet randomly calculate whether the given probability `num`/`total` is met.

### MeetProb

```go
func MeetProb(prob float32) bool
```

MeetProb randomly calculate whether the given probability is met.

## 内部依赖

- `errors/gcode`
- `errors/gerror`

