# mysql

> Package: `github.com/gogf/gf/contrib/drivers/mysql/v2`

```go
import "github.com/gogf/gf/contrib/drivers/mysql/v2"
```

## 概述

Package mysql implements gdb.Driver, which supports operations for database MySQL.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/drivers/mysql/v2`

## 源文件

- `mysql.go`
- `mysql_do_filter.go`
- `mysql_open.go`
- `mysql_table_fields.go`
- `mysql_tables.go`

## 类型定义

### Driver

**类型**: struct

Driver is the driver for mysql database.


**方法**:

- `func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error)`
  New creates and returns a database object for mysql.
- `func (d *Driver) GetChars() (charLeft string, charRight string)`
  GetChars returns the security char for this type of database.
- `func (d *Driver) DoFilter(ctx context.Context, link gdb.Link, sql string, args []any) (newSql string, newArgs []any, err error)`
  DoFilter handles the sql before posts it to database.
- `func (d *Driver) Open(config *gdb.ConfigNode) (db *sql.DB, err error)`
  Open creates and returns an underlying sql.DB object for mysql.
- `func (d *Driver) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*gdb.TableField, err error)`
  TableFields retrieves and returns the fields' information of specified table ...
- `func (d *Driver) Tables(ctx context.Context, schema ...string) (tables []string, err error)`
  Tables retrieves and returns the tables of current schema.

## 函数

### New

```go
func New() gdb.Driver
```

New create and returns a driver that implements gdb.Driver, which supports operations for MySQL.

## 内部依赖

- `database/gdb`
- `errors/gcode`
- `errors/gerror`
- `frame/g`
- `util/gutil`

## 外部依赖

- `github.com/go-sql-driver/mysql`

