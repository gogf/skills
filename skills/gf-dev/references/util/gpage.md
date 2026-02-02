# gpage

> Package: `github.com/gogf/gf/v2/util/gpage`

```go
import "github.com/gogf/gf/v2/util/gpage"
```

## 概述

Package gpage provides useful paging functionality for web pages.

## 源文件

- `gpage.go`

## 类型定义

### Page

**类型**: struct

Page is the pagination implementer.
All the attributes are public, you can change them when necessary.

Deprecated: wrap this pagination html content in business layer.


**方法**:

- `func (p *Page) NextPage() string`
  NextPage returns the HTML content for the next page.
- `func (p *Page) PrevPage() string`
  PrevPage returns the HTML content for the previous page.
- `func (p *Page) FirstPage() string`
  FirstPage returns the HTML content for the first page.
- `func (p *Page) LastPage() string`
  LastPage returns the HTML content for the last page.
- `func (p *Page) PageBar() string`
  PageBar returns the HTML page bar content with link and span tags.
- `func (p *Page) SelectBar() string`
  SelectBar returns the select HTML content for pagination.
- `func (p *Page) GetContent(mode int) string`
  GetContent returns the page content for predefined mode.
- `func (p *Page) GetUrl(page int) string`
  GetUrl parses the UrlTemplate with given page number and returns the URL string.
- `func (p *Page) GetLink(page int, text string, title string) string`
  GetLink returns the HTML link tag `a` content for given page number.

## 函数

### New

```go
func New(totalSize int, pageSize int, currentPage int, urlTemplate string) *Page
```

New creates and returns a pagination manager.
Note that the parameter `urlTemplate` specifies the URL producing template, like:
/user/list/{.page}, /user/list/{.page}.html, /user/list?page={.page}&type=1, etc.
The build-in variable in `urlTemplate` "{.page}" specifies the page number, which will be replaced by certain
page number when producing.

Deprecated: wrap this pagination html content in business layer.

## 内部依赖

- `text/gstr`
- `util/gconv`

