# mariadb

> Package: `github.com/gogf/gf/contrib/drivers/mariadb/v2`

```go
import "github.com/gogf/gf/contrib/drivers/mariadb/v2"
```

## 概述

Package mariadb implements gdb.Driver, which supports operations for database MariaDB.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/drivers/mariadb/v2`

## 源文件

- `mariadb.go`
- `mariadb_table_fields.go`

## 类型定义

### Driver

**类型**: struct

Driver is the driver for MariaDB database.

MariaDB is a community-developed, commercially supported fork of the MySQL relational database.
This driver uses the MySQL protocol to communicate with MariaDB database, as MariaDB maintains
high compatibility with MySQL protocol.

Although MariaDB is compatible with MySQL protocol, it is packaged as a separate driver component
rather than reusing the mysql adapter directly. This design allows for future extensibility,
such as implementing MariaDB-specific features or optimizations.


**方法**:

- `func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error)`
  New creates and returns a database object for MariaDB.
- `func (d *Driver) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*gdb.TableField, err error)`
  TableFields retrieves and returns the fields' information of specified table ...

## 函数

### New

```go
func New() gdb.Driver
```

New creates and returns a driver that implements gdb.Driver, which supports operations for MariaDB.

## 内部依赖

- `github.com/gogf/gf/contrib/drivers/mysql/v2`
- `database/gdb`
- `frame/g`
- `util/gutil`

