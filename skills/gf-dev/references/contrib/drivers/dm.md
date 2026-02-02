# dm

> Package: `github.com/gogf/gf/contrib/drivers/dm/v2`

```go
import "github.com/gogf/gf/contrib/drivers/dm/v2"
```

## 概述

Package dm implements gdb.Driver, which supports operations for database DM.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/drivers/dm/v2`

## 源文件

- `dm.go`
- `dm_convert.go`
- `dm_do_filter.go`
- `dm_do_insert.go`
- `dm_do_query.go`
- `dm_open.go`
- `dm_table_fields.go`
- `dm_tables.go`

## 类型定义

### Driver

**类型**: struct

Driver is the driver for dm database.


**方法**:

- `func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error)`
  New creates and returns a database object for dm.
- `func (d *Driver) GetChars() (charLeft string, charRight string)`
  GetChars returns the security char for this type of database.
- `func (d *Driver) ConvertValueForField(ctx context.Context, fieldType string, fieldValue any) (any, error)`
  ConvertValueForField converts value to the type of the record field.
- `func (d *Driver) DoFilter(ctx context.Context, link gdb.Link, sql string, args []any) (newSql string, newArgs []any, err error)`
  DoFilter deals with the sql string before commits it to underlying sql driver.
- `func (d *Driver) DoInsert(ctx context.Context, link gdb.Link, table string, list gdb.List, option gdb.DoInsertOption) (result sql.Result, err error)`
  DoInsert inserts or updates data for given table.
- `func (d *Driver) Open(config *gdb.ConfigNode) (db *sql.DB, err error)`
  Open creates and returns an underlying sql.DB object for pgsql.
- `func (d *Driver) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*gdb.TableField, err error)`
  TableFields retrieves and returns the fields' information of specified table ...
- `func (d *Driver) Tables(ctx context.Context, schema ...string) (tables []string, err error)`
  Tables retrieves and returns the tables of current schema.

## 函数

### New

```go
func New() gdb.Driver
```

New create and returns a driver that implements gdb.Driver, which supports operations for dm.

## 内部依赖

- `container/gmap`
- `container/gset`
- `database/gdb`
- `errors/gcode`
- `errors/gerror`
- `frame/g`
- `os/gtime`
- `text/gregex`
- `text/gstr`
- `util/gutil`

## 外部依赖

- `gitee.com/chunanyong/dm`

