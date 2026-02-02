# mssql

> Package: `github.com/gogf/gf/contrib/drivers/mssql/v2`

```go
import "github.com/gogf/gf/contrib/drivers/mssql/v2"
```

## 概述

Package mssql implements gdb.Driver, which supports operations for MSSQL.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/drivers/mssql/v2`

## 源文件

- `mssql.go`
- `mssql_do_commit.go`
- `mssql_do_exec.go`
- `mssql_do_filter.go`
- `mssql_do_insert.go`
- `mssql_open.go`
- `mssql_result.go`
- `mssql_table_fields.go`
- `mssql_tables.go`

## 类型定义

### Driver

**类型**: struct

Driver is the driver for SQL server database.


**方法**:

- `func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error)`
  New creates and returns a database object for SQL server.
- `func (d *Driver) GetChars() (charLeft string, charRight string)`
  GetChars returns the security char for this type of database.
- `func (d *Driver) DoCommit(ctx context.Context, in gdb.DoCommitInput) (out gdb.DoCommitOutput, err error)`
  DoCommit commits current sql and arguments to underlying sql driver.
- `func (d *Driver) DoExec(ctx context.Context, link gdb.Link, sqlStr string, args ...interface{}) (result sql.Result, err error)`
  DoExec commits the sql string and its arguments to underlying driver
- `func (d *Driver) GetTableNameFromSql(sqlStr string) table string`
  GetTableNameFromSql get table name from sql statement
- `func (d *Driver) GetInsertOutputSql(ctx context.Context, table string) string`
  GetInsertOutputSql  gen get last_insert_id code
- `func (d *Driver) DoFilter(ctx context.Context, link gdb.Link, sql string, args []any) (newSql string, newArgs []any, err error)`
  DoFilter deals with the sql string before commits it to underlying sql driver.
- `func (d *Driver) DoInsert(ctx context.Context, link gdb.Link, table string, list gdb.List, option gdb.DoInsertOption) (result sql.Result, err error)`
  DoInsert inserts or updates data for given table.
- `func (d *Driver) Open(config *gdb.ConfigNode) (db *sql.DB, err error)`
  Open creates and returns an underlying sql.DB object for mssql.
- `func (d *Driver) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*gdb.TableField, err error)`
  TableFields retrieves and returns the fields' information of specified table ...
- `func (d *Driver) Tables(ctx context.Context, schema ...string) (tables []string, err error)`
  Tables retrieves and returns the tables of current schema.

### Result

**类型**: struct

Result instance of sql.Result


**方法**:

- `func (r *Result) LastInsertId() (int64, error)`
- `func (r *Result) RowsAffected() (int64, error)`

## 函数

### New

```go
func New() gdb.Driver
```

New create and returns a driver that implements gdb.Driver, which supports operations for Mssql.

## 内部依赖

- `container/gset`
- `database/gdb`
- `errors/gcode`
- `errors/gerror`
- `text/gregex`
- `text/gstr`
- `util/gutil`

## 外部依赖

- `github.com/microsoft/go-mssqldb`

