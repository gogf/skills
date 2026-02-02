# gxml

> Package: `github.com/gogf/gf/v2/encoding/gxml`

```go
import "github.com/gogf/gf/v2/encoding/gxml"
```

## 概述

Package gxml provides accessing and converting for XML content.

## 源文件

- `gxml.go`

## 函数

### Decode

```go
func Decode(content []byte) (map[string]any, error)
```

Decode parses `content` into and returns as map.

### DecodeWithoutRoot

```go
func DecodeWithoutRoot(content []byte) (map[string]any, error)
```

DecodeWithoutRoot parses `content` into a map, and returns the map without root level.

### XMLEscapeChars

```go
func XMLEscapeChars(b ...bool)
```

XMLEscapeChars forces escaping invalid characters in attribute and element values.
NOTE: this is brute force with NO interrogation of '&' being escaped already; if it is
then '&amp;' will be re-escaped as '&amp;amp;'.

	The values are:
	"   &quot;
	'   &apos;
	<   &lt;
	>   &gt;
	&   &amp;

Note: if XMLEscapeCharsDecoder(true) has been called - or the default, 'false,' value
has been toggled to 'true' - then XMLEscapeChars(true) is ignored.  If XMLEscapeChars(true)
has already been called before XMLEscapeCharsDecoder(true), XMLEscapeChars(false) is called
to turn escape encoding on mv.Xml, etc., to prevent double escaping ampersands, '&'.

### Encode

```go
func Encode(m map[string]any, rootTag ...string) ([]byte, error)
```

Encode encodes map `m` to an XML format content as bytes.
The optional parameter `rootTag` is used to specify the XML root tag.

### EncodeWithIndent

```go
func EncodeWithIndent(m map[string]any, rootTag ...string) ([]byte, error)
```

EncodeWithIndent encodes map `m` to an XML format content as bytes with indent.
The optional parameter `rootTag` is used to specify the XML root tag.

### ToJson

```go
func ToJson(content []byte) ([]byte, error)
```

ToJson converts `content` as XML format into JSON format bytes.

## 内部依赖

- `encoding/gcharset`
- `errors/gerror`
- `text/gregex`

## 外部依赖

- `github.com/clbanning/mxj/v2`

