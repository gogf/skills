# gaussdb

> Package: `github.com/gogf/gf/contrib/drivers/gaussdb/v2`

```go
import "github.com/gogf/gf/contrib/drivers/gaussdb/v2"
```

## 概述

Package gaussdb implements gdb.Driver, which supports operations for database GaussDB.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/drivers/gaussdb/v2`

## 源文件

- `gaussdb.go`
- `gaussdb_convert.go`
- `gaussdb_do_exec.go`
- `gaussdb_do_filter.go`
- `gaussdb_do_insert.go`
- `gaussdb_open.go`
- `gaussdb_order.go`
- `gaussdb_result.go`
- `gaussdb_table_fields.go`
- `gaussdb_tables.go`

## 类型定义

### Driver

**类型**: struct

Driver is the driver for GaussDB database.


**方法**:

- `func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error)`
  New creates and returns a database object for postgresql.
- `func (d *Driver) GetChars() (charLeft string, charRight string)`
  GetChars returns the security char for this type of database.
- `func (d *Driver) ConvertValueForField(ctx context.Context, fieldType string, fieldValue any) (any, error)`
  ConvertValueForField converts value to database acceptable value.
- `func (d *Driver) CheckLocalTypeForField(ctx context.Context, fieldType string, fieldValue any) (gdb.LocalType, error)`
  CheckLocalTypeForField checks and returns corresponding local golang type for...
- `func (d *Driver) ConvertValueForLocal(ctx context.Context, fieldType string, fieldValue any) (any, error)`
  ConvertValueForLocal converts value to local Golang type of value according f...
- `func (d *Driver) DoExec(ctx context.Context, link gdb.Link, sql string, args ...any) (result sql.Result, err error)`
  DoExec commits the sql string and its arguments to underlying driver
- `func (d *Driver) DoFilter(ctx context.Context, link gdb.Link, sql string, args []any) (newSql string, newArgs []any, err error)`
  DoFilter deals with the sql string before commits it to underlying sql driver.
- `func (d *Driver) DoInsert(ctx context.Context, link gdb.Link, table string, list gdb.List, option gdb.DoInsertOption) (result sql.Result, err error)`
  DoInsert inserts or updates data for given table.
- `func (d *Driver) Open(config *gdb.ConfigNode) (db *sql.DB, err error)`
  Open creates and returns an underlying sql.DB object for GaussDB (openGauss).
- `func (d *Driver) OrderRandomFunction() string`
  OrderRandomFunction returns the SQL function for random ordering.
- `func (d *Driver) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*gdb.TableField, err error)`
  TableFields retrieves and returns the fields' information of specified table ...
- `func (d *Driver) Tables(ctx context.Context, schema ...string) (tables []string, err error)`
  Tables retrieves and returns the tables of current schema.

### Result

**类型**: struct

**方法**:

- `func (pgr Result) RowsAffected() (int64, error)`
- `func (pgr Result) LastInsertId() (int64, error)`

## 函数

### New

```go
func New() gdb.Driver
```

New create and returns a driver that implements gdb.Driver, which supports operations for PostgreSql.

## 内部依赖

- `container/gset`
- `database/gdb`
- `errors/gcode`
- `errors/gerror`
- `frame/g`
- `os/gctx`
- `text/gregex`
- `text/gstr`
- `util/gconv`
- `util/gutil`

## 外部依赖

- `gitee.com/opengauss/openGauss-connector-go-pq`
- `github.com/google/uuid`

