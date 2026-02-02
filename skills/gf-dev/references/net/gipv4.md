# gipv4

> Package: `github.com/gogf/gf/v2/net/gipv4`

```go
import "github.com/gogf/gf/v2/net/gipv4"
```

## 源文件

- `gipv4.go`
- `gipv4_convert.go`
- `gipv4_ip.go`
- `gipv4_lookup.go`
- `gipv4_mac.go`

## 函数

### Validate

```go
func Validate(ip string) bool
```

Validate checks whether given `ip` a valid IPv4 address.

### ParseAddress

```go
func ParseAddress(address string) (string, int)
```

ParseAddress parses `address` to its ip and port.
Eg: 192.168.1.1:80 -> 192.168.1.1, 80

### GetSegment

```go
func GetSegment(ip string) string
```

GetSegment returns the segment of given ip address.
Eg: 192.168.2.102 -> 192.168.2

### IpToLongBigEndian

```go
func IpToLongBigEndian(ip string) uint32
```

IpToLongBigEndian converts ip address to an uint32 integer with big endian.

### LongToIpBigEndian

```go
func LongToIpBigEndian(long uint32) string
```

LongToIpBigEndian converts an uint32 integer ip address to its string type address with big endian.

### IpToLongLittleEndian

```go
func IpToLongLittleEndian(ip string) uint32
```

IpToLongLittleEndian converts ip address to an uint32 integer with little endian.

### LongToIpLittleEndian

```go
func LongToIpLittleEndian(long uint32) string
```

LongToIpLittleEndian converts an uint32 integer ip address to its string type address with little endian.

### Ip2long

```go
func Ip2long(ip string) uint32
```

Ip2long converts ip address to an uint32 integer.

Deprecated: Use IpToLongBigEndian instead.

### Long2ip

```go
func Long2ip(long uint32) string
```

Long2ip converts an uint32 integer ip address to its string type address.

Deprecated: Use LongToIpBigEndian instead.

### GetIpArray

```go
func GetIpArray() (ips []string, err error)
```

GetIpArray retrieves and returns all the ip of current host.

### MustGetIntranetIp

```go
func MustGetIntranetIp() string
```

MustGetIntranetIp performs as GetIntranetIp, but it panics if any error occurs.

### GetIntranetIp

```go
func GetIntranetIp() (ip string, err error)
```

GetIntranetIp retrieves and returns the first intranet ip of current machine.

### GetIntranetIpArray

```go
func GetIntranetIpArray() (ips []string, err error)
```

GetIntranetIpArray retrieves and returns the intranet ip list of current machine.

### IsIntranet

```go
func IsIntranet(ip string) bool
```

IsIntranet checks and returns whether given ip an intranet ip.

Local: 127.0.0.1
A: 10.0.0.0--10.255.255.255
B: 172.16.0.0--172.31.255.255
C: 192.168.0.0--192.168.255.255

### GetHostByName

```go
func GetHostByName(hostname string) (string, error)
```

GetHostByName returns the IPv4 address corresponding to a given Internet host name.

### GetHostsByName

```go
func GetHostsByName(hostname string) ([]string, error)
```

GetHostsByName returns a list of IPv4 addresses corresponding to a given Internet
host name.

### GetNameByAddr

```go
func GetNameByAddr(ipAddress string) (string, error)
```

GetNameByAddr returns the Internet host name corresponding to a given IP address.

### GetMac

```go
func GetMac() (mac string, err error)
```

GetMac retrieves and returns the first mac address of current host.

### GetMacArray

```go
func GetMacArray() (macs []string, err error)
```

GetMacArray retrieves and returns all the mac address of current host.

## 内部依赖

- `errors/gerror`

