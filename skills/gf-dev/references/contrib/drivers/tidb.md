# tidb

> Package: `github.com/gogf/gf/contrib/drivers/tidb/v2`

```go
import "github.com/gogf/gf/contrib/drivers/tidb/v2"
```

## 概述

Package tidb implements gdb.Driver, which supports operations for database TiDB.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/drivers/tidb/v2`

## 源文件

- `tidb.go`

## 类型定义

### Driver

**类型**: struct

Driver is the driver for TiDB database.

TiDB is an open-source NewSQL database that supports Hybrid Transactional and Analytical Processing (HTAP).
This driver uses the MySQL protocol to communicate with TiDB database, as TiDB is designed to be highly
compatible with the MySQL protocol.

Although TiDB is compatible with MySQL protocol, it is packaged as a separate driver component
rather than reusing the mysql adapter directly. This design allows for future extensibility,
such as implementing TiDB-specific features like distributed transactions or optimizations.


**方法**:

- `func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error)`
  New creates and returns a database object for TiDB.

## 函数

### New

```go
func New() gdb.Driver
```

New creates and returns a driver that implements gdb.Driver, which supports operations for TiDB.

## 内部依赖

- `github.com/gogf/gf/contrib/drivers/mysql/v2`
- `database/gdb`
- `frame/g`

