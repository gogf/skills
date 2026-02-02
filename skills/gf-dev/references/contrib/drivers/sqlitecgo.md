# sqlitecgo

> Package: `github.com/gogf/gf/contrib/drivers/sqlitecgo/v2`

```go
import "github.com/gogf/gf/contrib/drivers/sqlitecgo/v2"
```

## 概述

Package sqlitecgo implements gdb.Driver, which supports operations for database SQLite.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/drivers/sqlitecgo/v2`

## 源文件

- `sqlitecgo.go`
- `sqlitecgo_do_filter.go`
- `sqlitecgo_format_upsert.go`
- `sqlitecgo_open.go`
- `sqlitecgo_table_fields.go`
- `sqlitecgo_tables.go`

## 类型定义

### Driver

**类型**: struct

Driver is the driver for sqlite database.


**方法**:

- `func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error)`
  New creates and returns a database object for sqlite.
- `func (d *Driver) GetChars() (charLeft string, charRight string)`
  GetChars returns the security char for this type of database.
- `func (d *Driver) DoFilter(ctx context.Context, link gdb.Link, sql string, args []any) (newSql string, newArgs []any, err error)`
  DoFilter deals with the sql string before commits it to underlying sql driver.
- `func (d *Driver) FormatUpsert(columns []string, list gdb.List, option gdb.DoInsertOption) (string, error)`
  FormatUpsert returns SQL clause of type upsert for SQLite.
- `func (d *Driver) Open(config *gdb.ConfigNode) (db *sql.DB, err error)`
  Open creates and returns an underlying sql.DB object for sqlite.
- `func (d *Driver) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*gdb.TableField, err error)`
  TableFields retrieves and returns the fields' information of specified table ...
- `func (d *Driver) Tables(ctx context.Context, schema ...string) (tables []string, err error)`
  Tables retrieves and returns the tables of current schema.

## 函数

### New

```go
func New() gdb.Driver
```

New create and returns a driver that implements gdb.Driver, which supports operations for SQLite.

## 内部依赖

- `database/gdb`
- `encoding/gurl`
- `errors/gcode`
- `errors/gerror`
- `os/gfile`
- `text/gstr`
- `util/gconv`
- `util/gutil`

## 外部依赖

- `github.com/mattn/go-sqlite3`

