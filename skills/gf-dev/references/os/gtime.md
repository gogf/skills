# gtime

> Package: `github.com/gogf/gf/v2/os/gtime`

```go
import "github.com/gogf/gf/v2/os/gtime"
```

## 概述

Package gtime provides functionality for measuring and displaying time.

## 源文件

- `gtime.go`
- `gtime_format.go`
- `gtime_sql.go`
- `gtime_time.go`
- `gtime_time_wrapper.go`
- `gtime_time_zone.go`

## 类型定义

### Time

**类型**: struct

Time is a wrapper for time.Time for additional features.


**方法**:

- `func (t *Time) Timestamp() int64`
  Timestamp returns the timestamp in seconds.
- `func (t *Time) TimestampMilli() int64`
  TimestampMilli returns the timestamp in milliseconds.
- `func (t *Time) TimestampMicro() int64`
  TimestampMicro returns the timestamp in microseconds.
- `func (t *Time) TimestampNano() int64`
  TimestampNano returns the timestamp in nanoseconds.
- `func (t *Time) TimestampStr() string`
  TimestampStr is a convenience method which retrieves and returns
- `func (t *Time) TimestampMilliStr() string`
  TimestampMilliStr is a convenience method which retrieves and returns
- `func (t *Time) TimestampMicroStr() string`
  TimestampMicroStr is a convenience method which retrieves and returns
- `func (t *Time) TimestampNanoStr() string`
  TimestampNanoStr is a convenience method which retrieves and returns
- `func (t *Time) Month() int`
  Month returns the month of the year specified by t.
- `func (t *Time) Second() int`
  Second returns the second offset within the minute specified by t,
- `func (t *Time) Millisecond() int`
  Millisecond returns the millisecond offset within the second specified by t,
- `func (t *Time) Microsecond() int`
  Microsecond returns the microsecond offset within the second specified by t,
- `func (t *Time) Nanosecond() int`
  Nanosecond returns the nanosecond offset within the second specified by t,
- `func (t *Time) String() string`
  String returns current time object as string.
- `func (t *Time) IsZero() bool`
  IsZero reports whether t represents the zero time instant,
- `func (t *Time) Clone() *Time`
  Clone returns a new Time object which is a clone of current time object.
- `func (t *Time) Add(d time.Duration) *Time`
  Add adds the duration to current time.
- `func (t *Time) AddStr(duration string) (*Time, error)`
  AddStr parses the given duration as string and adds it to current time.
- `func (t *Time) UTC() *Time`
  UTC converts current time to UTC timezone.
- `func (t *Time) ISO8601() string`
  ISO8601 formats the time as ISO8601 and returns it as string.
- `func (t *Time) RFC822() string`
  RFC822 formats the time as RFC822 and returns it as string.
- `func (t *Time) AddDate(years int, months int, days int) *Time`
  AddDate adds year, month and day to the time.
- `func (t *Time) Round(d time.Duration) *Time`
  Round returns the result of rounding t to the nearest multiple of d (since th...
- `func (t *Time) Truncate(d time.Duration) *Time`
  Truncate returns the result of rounding t down to a multiple of d (since the ...
- `func (t *Time) Equal(u *Time) bool`
  Equal reports whether t and u represent the same time instant.
- `func (t *Time) Before(u *Time) bool`
  Before reports whether the time instant t is before u.
- `func (t *Time) After(u *Time) bool`
  After reports whether the time instant t is after u.
- `func (t *Time) Sub(u *Time) time.Duration`
  Sub returns the duration t-u. If the result exceeds the maximum (or minimum)
- `func (t *Time) StartOfMinute() *Time`
  StartOfMinute clones and returns a new time of which the seconds is set to 0.
- `func (t *Time) StartOfHour() *Time`
  StartOfHour clones and returns a new time of which the hour, minutes and seco...
- `func (t *Time) StartOfDay() *Time`
  StartOfDay clones and returns a new time which is the start of day, its time ...
- `func (t *Time) StartOfWeek() *Time`
  StartOfWeek clones and returns a new time which is the first day of week and ...
- `func (t *Time) StartOfMonth() *Time`
  StartOfMonth clones and returns a new time which is the first day of the mont...
- `func (t *Time) StartOfQuarter() *Time`
  StartOfQuarter clones and returns a new time which is the first day of the qu...
- `func (t *Time) StartOfHalf() *Time`
  StartOfHalf clones and returns a new time which is the first day of the half ...
- `func (t *Time) StartOfYear() *Time`
  StartOfYear clones and returns a new time which is the first day of the year ...
- `func (t *Time) EndOfMinute(withNanoPrecision ...bool) *Time`
  EndOfMinute clones and returns a new time of which the seconds is set to 59.
- `func (t *Time) EndOfHour(withNanoPrecision ...bool) *Time`
  EndOfHour clones and returns a new time of which the minutes and seconds are ...
- `func (t *Time) EndOfDay(withNanoPrecision ...bool) *Time`
  EndOfDay clones and returns a new time which is the end of day the and its ti...
- `func (t *Time) EndOfWeek(withNanoPrecision ...bool) *Time`
  EndOfWeek clones and returns a new time which is the end of week and its time...
- `func (t *Time) EndOfMonth(withNanoPrecision ...bool) *Time`
  EndOfMonth clones and returns a new time which is the end of the month and it...
- `func (t *Time) EndOfQuarter(withNanoPrecision ...bool) *Time`
  EndOfQuarter clones and returns a new time which is end of the quarter and it...
- `func (t *Time) EndOfHalf(withNanoPrecision ...bool) *Time`
  EndOfHalf clones and returns a new time which is the end of the half year and...
- `func (t *Time) EndOfYear(withNanoPrecision ...bool) *Time`
  EndOfYear clones and returns a new time which is the end of the year and its ...
- `func (t Time) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (t *Time) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (t *Time) UnmarshalText(data []byte) error`
  UnmarshalText implements the encoding.TextUnmarshaler interface.
- `func (t *Time) NoValidation()`
  NoValidation marks this struct object will not be validated by package gvalid.
- `func (t *Time) DeepCopy() any`
  DeepCopy implements interface for deep copy of current type.
- `func (t *Time) ToLocation(location *time.Location) *Time`
  ToLocation converts current time to specified location.
- `func (t *Time) ToZone(zone string) (*Time, error)`
  ToZone converts current time to specified zone like: Asia/Shanghai.
- `func (t *Time) Local() *Time`
  Local converts the time to local timezone.

## 函数

### Timestamp

```go
func Timestamp() int64
```

Timestamp retrieves and returns the timestamp in seconds.

### TimestampMilli

```go
func TimestampMilli() int64
```

TimestampMilli retrieves and returns the timestamp in milliseconds.

### TimestampMicro

```go
func TimestampMicro() int64
```

TimestampMicro retrieves and returns the timestamp in microseconds.

### TimestampNano

```go
func TimestampNano() int64
```

TimestampNano retrieves and returns the timestamp in nanoseconds.

### TimestampStr

```go
func TimestampStr() string
```

TimestampStr is a convenience method which retrieves and returns
the timestamp in seconds as string.

### TimestampMilliStr

```go
func TimestampMilliStr() string
```

TimestampMilliStr is a convenience method which retrieves and returns
the timestamp in milliseconds as string.

### TimestampMicroStr

```go
func TimestampMicroStr() string
```

TimestampMicroStr is a convenience method which retrieves and returns
the timestamp in microseconds as string.

### TimestampNanoStr

```go
func TimestampNanoStr() string
```

TimestampNanoStr is a convenience method which retrieves and returns
the timestamp in nanoseconds as string.

### Date

```go
func Date() string
```

Date returns current date in string like "2006-01-02".

### Datetime

```go
func Datetime() string
```

Datetime returns current datetime in string like "2006-01-02 15:04:05".

### ISO8601

```go
func ISO8601() string
```

ISO8601 returns current datetime in ISO8601 format like "2006-01-02T15:04:05-07:00".

### RFC822

```go
func RFC822() string
```

RFC822 returns current datetime in RFC822 format like "Mon, 02 Jan 06 15:04 MST".

### StrToTime

```go
func StrToTime(str string, format ...string) (*Time, error)
```

StrToTime converts string to *Time object. It also supports timestamp string.
The parameter `format` is unnecessary, which specifies the format for converting like "Y-m-d H:i:s".
If `format` is given, it acts as same as function StrToTimeFormat.
If `format` is not given, it converts string as a "standard" datetime string.
Note that, it fails and returns error if there's no date string in `str`.

### ConvertZone

```go
func ConvertZone(strTime string, toZone string, fromZone ...string) (*Time, error)
```

ConvertZone converts time in string `strTime` from `fromZone` to `toZone`.
The parameter `fromZone` is unnecessary, it is current time zone in default.

### StrToTimeFormat

```go
func StrToTimeFormat(str string, format string) (*Time, error)
```

StrToTimeFormat parses string `str` to *Time object with given format `format`.
The parameter `format` is like "Y-m-d H:i:s".

### StrToTimeLayout

```go
func StrToTimeLayout(str string, layout string) (*Time, error)
```

StrToTimeLayout parses string `str` to *Time object with given format `layout`.
The parameter `layout` is in stdlib format like "2006-01-02 15:04:05".

### ParseTimeFromContent

```go
func ParseTimeFromContent(content string, format ...string) *Time
```

ParseTimeFromContent retrieves time information for content string, it then parses and returns it
as *Time object.
It returns the first time information if there are more than one time string in the content.
It only retrieves and parses the time information with given first matched `format` if it's passed.

### ParseDuration

```go
func ParseDuration(s string) (duration time.Duration, err error)
```

ParseDuration parses a duration string.
A duration string is a possibly signed sequence of
decimal numbers, each with optional fraction and a unit suffix,
such as "300ms", "-1.5h", "1d" or "2h45m".
Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h", "d".

Very note that it supports unit "d" more than function time.ParseDuration.

### FuncCost

```go
func FuncCost(f func) time.Duration
```

FuncCost calculates the cost time of function `f` in nanoseconds.

### New

```go
func New(param ...any) *Time
```

New creates and returns a Time object with given parameter.
The optional parameter is the time object which can be type of: time.Time/*time.Time, string or integer.
Example:
New("2024-10-29")
New(1390876568)
New(t) // The t is type of time.Time.

### Now

```go
func Now() *Time
```

Now creates and returns a time object of now.

### NewFromTime

```go
func NewFromTime(t time.Time) *Time
```

NewFromTime creates and returns a Time object with given time.Time object.

### NewFromStr

```go
func NewFromStr(str string) *Time
```

NewFromStr creates and returns a Time object with given string.
Note that it returns nil if there's error occurs.

### NewFromStrFormat

```go
func NewFromStrFormat(str string, format string) *Time
```

NewFromStrFormat creates and returns a Time object with given string and
custom format like: Y-m-d H:i:s.
Note that it returns nil if there's error occurs.

### NewFromStrLayout

```go
func NewFromStrLayout(str string, layout string) *Time
```

NewFromStrLayout creates and returns a Time object with given string and
stdlib layout like: 2006-01-02 15:04:05.
Note that it returns nil if there's error occurs.

### NewFromTimeStamp

```go
func NewFromTimeStamp(timestamp int64) *Time
```

NewFromTimeStamp creates and returns a Time object with given timestamp,
which can be in seconds to nanoseconds.
Eg: 1600443866 and 1600443866199266000 are both considered as valid timestamp number.

### SetTimeZone

```go
func SetTimeZone(zone string) err error
```

SetTimeZone sets the time zone for current whole process.
The parameter `zone` is an area string specifying corresponding time zone,
eg: Asia/Shanghai.

PLEASE VERY NOTE THAT:
1. This should be called before package "time" import.
2. This function should be called once.
3. Please refer to issue: https://github.com/golang/go/issues/34814

## 内部依赖

- `errors/gcode`
- `errors/gerror`
- `internal/intlog`
- `internal/utils`
- `text/gregex`

