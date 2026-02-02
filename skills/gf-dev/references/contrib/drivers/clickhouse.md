# clickhouse

> Package: `github.com/gogf/gf/contrib/drivers/clickhouse/v2`

```go
import "github.com/gogf/gf/contrib/drivers/clickhouse/v2"
```

## 概述

Package clickhouse implements gdb.Driver, which supports operations for database ClickHouse.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/drivers/clickhouse/v2`

## 源文件

- `clickhouse.go`
- `clickhouse_convert.go`
- `clickhouse_do_commit.go`
- `clickhouse_do_delete.go`
- `clickhouse_do_filter.go`
- `clickhouse_do_insert.go`
- `clickhouse_do_update.go`
- `clickhouse_insert.go`
- `clickhouse_open.go`
- `clickhouse_ping.go`
- `clickhouse_table_fields.go`
- `clickhouse_tables.go`
- `clickhouse_transaction.go`

## 类型定义

### Driver

**类型**: struct

Driver is the driver for clickhouse database.


**方法**:

- `func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error)`
  New creates and returns a database object for clickhouse.
- `func (d *Driver) ConvertValueForField(ctx context.Context, fieldType string, fieldValue any) (any, error)`
  ConvertValueForField converts value to the type of the record field.
- `func (d *Driver) DoCommit(ctx context.Context, in gdb.DoCommitInput) (out gdb.DoCommitOutput, err error)`
  DoCommit commits current sql and arguments to underlying sql driver.
- `func (d *Driver) DoDelete(ctx context.Context, link gdb.Link, table string, condition string, args ...any) (result sql.Result, err error)`
  DoDelete does "DELETE FROM ... " statement for the table.
- `func (d *Driver) DoFilter(ctx context.Context, link gdb.Link, originSql string, args []any) (newSql string, newArgs []any, err error)`
  DoFilter handles the sql before posts it to database.
- `func (d *Driver) DoInsert(ctx context.Context, link gdb.Link, table string, list gdb.List, option gdb.DoInsertOption) (result sql.Result, err error)`
  DoInsert inserts or updates data for given table.
- `func (d *Driver) DoUpdate(ctx context.Context, link gdb.Link, table string, data any, condition string, args ...any) (result sql.Result, err error)`
  DoUpdate does "UPDATE ... " statement for the table.
- `func (d *Driver) InsertIgnore(ctx context.Context, table string, data any, batch ...int) (sql.Result, error)`
  InsertIgnore Other queries for modifying data parts are not supported: REPLAC...
- `func (d *Driver) InsertAndGetId(ctx context.Context, table string, data any, batch ...int) (int64, error)`
  InsertAndGetId Other queries for modifying data parts are not supported: REPL...
- `func (d *Driver) Replace(ctx context.Context, table string, data any, batch ...int) (sql.Result, error)`
  Replace Other queries for modifying data parts are not supported: REPLACE, ME...
- `func (d *Driver) Open(config *gdb.ConfigNode) (db *sql.DB, err error)`
  Open creates and returns an underlying sql.DB object for clickhouse.
- `func (d *Driver) PingMaster() error`
  PingMaster pings the master node to check authentication or keeps the connect...
- `func (d *Driver) PingSlave() error`
  PingSlave pings the slave node to check authentication or keeps the connectio...
- `func (d *Driver) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*gdb.TableField, err error)`
  TableFields retrieves and returns the fields' information of specified table ...
- `func (d *Driver) Tables(ctx context.Context, schema ...string) (tables []string, err error)`
  Tables retrieves and returns the tables of current schema.
- `func (d *Driver) Begin(ctx context.Context) (tx gdb.TX, err error)`
  Begin starts and returns the transaction object.
- `func (d *Driver) Transaction(ctx context.Context, f func) error`
  Transaction wraps the transaction logic using function `f`.

## 函数

### New

```go
func New() gdb.Driver
```

New create and returns a driver that implements gdb.Driver, which supports operations for clickhouse.

## 内部依赖

- `database/gdb`
- `errors/gcode`
- `errors/gerror`
- `os/gctx`
- `os/gtime`
- `text/gregex`
- `util/gutil`

## 外部依赖

- `github.com/ClickHouse/clickhouse-go/v2`
- `github.com/google/uuid`
- `github.com/shopspring/decimal`

