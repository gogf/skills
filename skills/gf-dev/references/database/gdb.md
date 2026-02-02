# gdb

> Package: `github.com/gogf/gf/v2/database/gdb`

```go
import "github.com/gogf/gf/v2/database/gdb"
```

## 概述

Package gdb provides ORM features for popular relationship databases.

## 源文件

- `gdb.go`
- `gdb_converter.go`
- `gdb_core.go`
- `gdb_core_config.go`
- `gdb_core_ctx.go`
- `gdb_core_link.go`
- `gdb_core_stats.go`
- `gdb_core_structure.go`
- `gdb_core_trace.go`
- `gdb_core_transaction.go`
- `gdb_core_txcore.go`
- `gdb_core_underlying.go`
- `gdb_core_utility.go`
- `gdb_driver_default.go`
- `gdb_driver_wrapper.go`
- `gdb_driver_wrapper_db.go`
- `gdb_func.go`
- `gdb_model.go`
- `gdb_model_builder.go`
- `gdb_model_builder_where.go`
- `gdb_model_builder_where_prefix.go`
- `gdb_model_builder_whereor.go`
- `gdb_model_builder_whereor_prefix.go`
- `gdb_model_cache.go`
- `gdb_model_delete.go`
- `gdb_model_fields.go`
- `gdb_model_hook.go`
- `gdb_model_insert.go`
- `gdb_model_join.go`
- `gdb_model_lock.go`
- `gdb_model_option.go`
- `gdb_model_order_group.go`
- `gdb_model_select.go`
- `gdb_model_sharding.go`
- `gdb_model_soft_time.go`
- `gdb_model_transaction.go`
- `gdb_model_update.go`
- `gdb_model_utility.go`
- `gdb_model_where.go`
- `gdb_model_where_prefix.go`
- `gdb_model_whereor.go`
- `gdb_model_whereor_prefix.go`
- `gdb_model_with.go`
- `gdb_result.go`
- `gdb_schema.go`
- `gdb_statement.go`
- `gdb_type_record.go`
- `gdb_type_result.go`
- `gdb_type_result_scanlist.go`

## 类型定义

### DB

**类型**: interface

DB defines the interfaces for ORM operations.


### TX

**类型**: interface

TX defines the interfaces for ORM transaction operations.


### StatsItem

**类型**: interface

StatsItem defines the stats information for a configuration node.


### Core

**类型**: struct

Core is the base struct for database management.


**方法**:

- `func (c *Core) GetCore() *Core`
  GetCore returns the underlying *Core object.
- `func (c *Core) Ctx(ctx context.Context) DB`
  Ctx is a chaining function, which creates and returns a new DB that is a shal...
- `func (c *Core) GetCtx() context.Context`
  GetCtx returns the context for current DB.
- `func (c *Core) GetCtxTimeout(ctx context.Context, timeoutType ctxTimeoutType) (context.Context, context.CancelFunc)`
  GetCtxTimeout returns the context and cancel function for specified timeout t...
- `func (c *Core) Close(ctx context.Context) err error`
  Close closes the database and prevents new queries from starting.
- `func (c *Core) Master(schema ...string) (*sql.DB, error)`
  Master creates and returns a connection from master node if master-slave conf...
- `func (c *Core) Slave(schema ...string) (*sql.DB, error)`
  Slave creates and returns a connection from slave node if master-slave config...
- `func (c *Core) GetAll(ctx context.Context, sql string, args ...any) (Result, error)`
  GetAll queries and returns data records from database.
- `func (c *Core) DoSelect(ctx context.Context, link Link, sql string, args ...any) (result Result, err error)`
  DoSelect queries and returns data records from database.
- `func (c *Core) GetOne(ctx context.Context, sql string, args ...any) (Record, error)`
  GetOne queries and returns one record from database.
- `func (c *Core) GetArray(ctx context.Context, sql string, args ...any) (Array, error)`
  GetArray queries and returns data values as slice from database.
- `func (c *Core) GetScan(ctx context.Context, pointer any, sql string, args ...any) error`
  GetScan queries one or more records from database and converts them to given ...
- `func (c *Core) GetValue(ctx context.Context, sql string, args ...any) (Value, error)`
  GetValue queries and returns the field value from database.
- `func (c *Core) GetCount(ctx context.Context, sql string, args ...any) (int, error)`
  GetCount queries and returns the count from database.
- `func (c *Core) Union(unions ...*Model) *Model`
  Union does "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." statement.
- `func (c *Core) UnionAll(unions ...*Model) *Model`
  UnionAll does "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." sta...
- `func (c *Core) PingMaster() error`
  PingMaster pings the master node to check authentication or keeps the connect...
- `func (c *Core) PingSlave() error`
  PingSlave pings the slave node to check authentication or keeps the connectio...
- `func (c *Core) Insert(ctx context.Context, table string, data any, batch ...int) (sql.Result, error)`
  Insert does "INSERT INTO ..." statement for the table.
- `func (c *Core) InsertIgnore(ctx context.Context, table string, data any, batch ...int) (sql.Result, error)`
  InsertIgnore does "INSERT IGNORE INTO ..." statement for the table.
- `func (c *Core) InsertAndGetId(ctx context.Context, table string, data any, batch ...int) (int64, error)`
  InsertAndGetId performs action Insert and returns the last insert id that aut...
- `func (c *Core) Replace(ctx context.Context, table string, data any, batch ...int) (sql.Result, error)`
  Replace does "REPLACE INTO ..." statement for the table.
- `func (c *Core) Save(ctx context.Context, table string, data any, batch ...int) (sql.Result, error)`
  Save does "INSERT INTO ... ON DUPLICATE KEY UPDATE..." statement for the table.
- `func (c *Core) DoInsert(ctx context.Context, link Link, table string, list List, option DoInsertOption) (result sql.Result, err error)`
  DoInsert inserts or updates data for given table.
- `func (c *Core) Update(ctx context.Context, table string, data any, condition any, args ...any) (sql.Result, error)`
  Update does "UPDATE ... " statement for the table.
- `func (c *Core) DoUpdate(ctx context.Context, link Link, table string, data any, condition string, args ...any) (result sql.Result, err error)`
  DoUpdate does "UPDATE ... " statement for the table.
- `func (c *Core) Delete(ctx context.Context, table string, condition any, args ...any) (result sql.Result, err error)`
  Delete does "DELETE FROM ... " statement for the table.
- `func (c *Core) DoDelete(ctx context.Context, link Link, table string, condition string, args ...any) (result sql.Result, err error)`
  DoDelete does "DELETE FROM ... " statement for the table.
- `func (c *Core) FilteredLink() string`
  FilteredLink retrieves and returns filtered `linkInfo` that can be using for
- `func (c *Core) MarshalJSON() ([]byte, error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (c *Core) HasTable(name string) (bool, error)`
  HasTable determine whether the table name exists in the database.
- `func (c *Core) GetInnerMemCache() *gcache.Cache`
  GetInnerMemCache retrieves and returns the inner memory cache object.
- `func (c *Core) SetTableFields(ctx context.Context, table string, fields map[string]*TableField, schema ...string) error`
- `func (c *Core) GetTablesWithCache() ([]string, error)`
  GetTablesWithCache retrieves and returns the table names of current database ...
- `func (c *Core) IsSoftCreatedFieldName(fieldName string) bool`
  IsSoftCreatedFieldName checks and returns whether given field name is an auto...
- `func (c *Core) FormatSqlBeforeExecuting(sql string, args []any) (newSql string, newArgs []any)`
  FormatSqlBeforeExecuting formats the sql string and its arguments before exec...
- `func (c *Core) SetLogger(logger glog.ILogger)`
  SetLogger sets the logger for orm.
- `func (c *Core) GetLogger() glog.ILogger`
  GetLogger returns the (logger) of the orm.
- `func (c *Core) SetMaxIdleConnCount(n int)`
  SetMaxIdleConnCount sets the maximum number of connections in the idle
- `func (c *Core) SetMaxOpenConnCount(n int)`
  SetMaxOpenConnCount sets the maximum number of open connections to the database.
- `func (c *Core) SetMaxConnLifeTime(d time.Duration)`
  SetMaxConnLifeTime sets the maximum amount of time a connection may be reused.
- `func (c *Core) SetMaxIdleConnTime(d time.Duration)`
  SetMaxIdleConnTime sets the maximum amount of time a connection may be idle b...
- `func (c *Core) GetConfig() *ConfigNode`
  GetConfig returns the current used node configuration.
- `func (c *Core) SetDebug(debug bool)`
  SetDebug enables/disables the debug mode.
- `func (c *Core) GetDebug() bool`
  GetDebug returns the debug value.
- `func (c *Core) GetCache() *gcache.Cache`
  GetCache returns the internal cache object.
- `func (c *Core) GetGroup() string`
  GetGroup returns the group string configured.
- `func (c *Core) SetDryRun(enabled bool)`
  SetDryRun enables/disables the DryRun feature.
- `func (c *Core) GetDryRun() bool`
  GetDryRun returns the DryRun value.
- `func (c *Core) GetPrefix() string`
  GetPrefix returns the table prefix string configured.
- `func (c *Core) GetSchema() string`
  GetSchema returns the schema configured.
- `func (c *Core) InjectIgnoreResult(ctx context.Context) context.Context`
- `func (c *Core) GetIgnoreResultFromCtx(ctx context.Context) bool`
- `func (c *Core) Stats(ctx context.Context) []StatsItem`
  Stats retrieves and returns the pool stat for all nodes that have been establ...
- `func (c *Core) GetFieldTypeStr(ctx context.Context, fieldName string, table string, schema string) string`
  GetFieldTypeStr retrieves and returns the field type string for certain field...
- `func (c *Core) GetFieldType(ctx context.Context, fieldName string, table string, schema string) *TableField`
  GetFieldType retrieves and returns the field type object for certain field by...
- `func (c *Core) ConvertDataForRecord(ctx context.Context, value any, table string) (map[string]any, error)`
  ConvertDataForRecord is a very important function, which does converting for ...
- `func (c *Core) ConvertValueForField(ctx context.Context, fieldType string, fieldValue any) (any, error)`
  ConvertValueForField converts value to the type of the record field.
- `func (c *Core) GetFormattedDBTypeNameForField(fieldType string) (typeName string, typePattern string)`
  GetFormattedDBTypeNameForField retrieves and returns the formatted database t...
- `func (c *Core) CheckLocalTypeForField(ctx context.Context, fieldType string, _ any) (LocalType, error)`
  CheckLocalTypeForField checks and returns corresponding type for given db type.
- `func (c *Core) ConvertValueForLocal(ctx context.Context, fieldType string, fieldValue any) (any, error)`
  ConvertValueForLocal converts value to local Golang type of value according f...
- `func (c *Core) Begin(ctx context.Context) (tx TX, err error)`
  Begin starts and returns the transaction object.
- `func (c *Core) BeginWithOptions(ctx context.Context, opts TxOptions) (tx TX, err error)`
  BeginWithOptions starts and returns the transaction object with given options.
- `func (c *Core) Transaction(ctx context.Context, f func) err error`
  Transaction wraps the transaction logic using function `f`.
- `func (c *Core) TransactionWithOptions(ctx context.Context, opts TxOptions, f func) err error`
  TransactionWithOptions wraps the transaction logic with propagation options u...
- `func (c *Core) Query(ctx context.Context, sql string, args ...any) (result Result, err error)`
  Query commits one query SQL to underlying driver and returns the execution re...
- `func (c *Core) DoQuery(ctx context.Context, link Link, sql string, args ...any) (result Result, err error)`
  DoQuery commits the sql string and its arguments to underlying driver
- `func (c *Core) Exec(ctx context.Context, sql string, args ...any) (result sql.Result, err error)`
  Exec commits one query SQL to underlying driver and returns the execution res...
- `func (c *Core) DoExec(ctx context.Context, link Link, sql string, args ...any) (result sql.Result, err error)`
  DoExec commits the sql string and its arguments to underlying driver
- `func (c *Core) DoFilter(ctx context.Context, link Link, sql string, args []any) (newSql string, newArgs []any, err error)`
  DoFilter is a hook function, which filters the sql and its arguments before i...
- `func (c *Core) DoCommit(ctx context.Context, in DoCommitInput) (out DoCommitOutput, err error)`
  DoCommit commits current sql and arguments to underlying sql driver.
- `func (c *Core) Prepare(ctx context.Context, sql string, execOnMaster ...bool) (*Stmt, error)`
  Prepare creates a prepared statement for later queries or executions.
- `func (c *Core) DoPrepare(ctx context.Context, link Link, sql string) (stmt *Stmt, err error)`
  DoPrepare calls prepare function on given link object and returns the stateme...
- `func (c *Core) FormatUpsert(columns []string, list List, option DoInsertOption) (string, error)`
  FormatUpsert formats and returns SQL clause part for upsert statement.
- `func (c *Core) RowsToResult(ctx context.Context, rows *sql.Rows) (Result, error)`
  RowsToResult converts underlying data record type sql.Rows to Result type.
- `func (c *Core) OrderRandomFunction() string`
  OrderRandomFunction returns the SQL function for random ordering.
- `func (c *Core) GetDB() DB`
  GetDB returns the underlying DB.
- `func (c *Core) GetLink(ctx context.Context, master bool, schema string) (Link, error)`
  GetLink creates and returns the underlying database link object with transact...
- `func (c *Core) MasterLink(schema ...string) (Link, error)`
  MasterLink acts like function Master but with additional `schema` parameter s...
- `func (c *Core) SlaveLink(schema ...string) (Link, error)`
  SlaveLink acts like function Slave but with additional `schema` parameter spe...
- `func (c *Core) QuoteWord(s string) string`
  QuoteWord checks given string `s` a word,
- `func (c *Core) QuoteString(s string) string`
  QuoteString quotes string with quote chars. Strings like:
- `func (c *Core) QuotePrefixTableName(table string) string`
  QuotePrefixTableName adds prefix string and quotes chars for the table.
- `func (c *Core) GetChars() (charLeft string, charRight string)`
  GetChars returns the security char for current database.
- `func (c *Core) Tables(ctx context.Context, schema ...string) (tables []string, err error)`
  Tables retrieves and returns the tables of current schema.
- `func (c *Core) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*TableField, err error)`
  TableFields retrieves and returns the fields' information of specified table ...
- `func (c *Core) ClearTableFields(ctx context.Context, table string, schema ...string) err error`
  ClearTableFields removes certain cached table fields of current configuration...
- `func (c *Core) ClearTableFieldsAll(ctx context.Context) err error`
  ClearTableFieldsAll removes all cached table fields of current configuration ...
- `func (c *Core) ClearCache(ctx context.Context, table string) err error`
  ClearCache removes cached sql result of certain table.
- `func (c *Core) ClearCacheAll(ctx context.Context) err error`
  ClearCacheAll removes all cached sql result from cache
- `func (c *Core) HasField(ctx context.Context, table string, field string, schema ...string) (bool, error)`
  HasField determine whether the field exists in the table.
- `func (c *Core) GetPrimaryKeys(ctx context.Context, table string, schema ...string) ([]string, error)`
  GetPrimaryKeys retrieves and returns the primary key field names of the speci...
- `func (c *Core) Model(tableNameQueryOrStruct ...any) *Model`
  Model creates and returns a new ORM model from given schema.
- `func (c *Core) Raw(rawSql string, args ...any) *Model`
  Raw creates and returns a model based on a raw sql not a table.
- `func (c *Core) With(objects ...any) *Model`
  With creates and returns an ORM model based on metadata of given object.
- `func (c *Core) Schema(schema string) *Schema`
  Schema creates and returns a schema.

### DoCommitInput

**类型**: struct

DoCommitInput is the input parameters for function DoCommit.


### DoCommitOutput

**类型**: struct

DoCommitOutput is the output parameters for function DoCommit.


### Driver

**类型**: interface

Driver is the interface for integrating sql drivers into package gdb.


### Link

**类型**: interface

Link is a common database function wrapper interface.
Note that, any operation using `Link` will have no SQL logging.


### Sql

**类型**: struct

Sql is the sql recording struct.


### DoInsertOption

**类型**: struct

DoInsertOption is the input struct for function DoInsert.


### TableField

**类型**: struct

TableField is the struct for table field.


### Counter

**类型**: struct

Counter is the type for update count.


### Raw

**类型**: type

### Value

**类型**: type

### Array

**类型**: type

### Record

**类型**: type

**方法**:

- `func (r Record) Json() string`
  Json converts `r` to JSON format content.
- `func (r Record) Xml(rootTag ...string) string`
  Xml converts `r` to XML format content.
- `func (r Record) Map() Map`
  Map converts `r` to map[string]any.
- `func (r Record) GMap() *gmap.StrAnyMap`
  GMap converts `r` to a gmap.
- `func (r Record) Struct(pointer any) error`
  Struct converts `r` to a struct.
- `func (r Record) IsEmpty() bool`
  IsEmpty checks and returns whether `r` is empty.

### Result

**类型**: type

**方法**:

- `func (r Result) IsEmpty() bool`
  IsEmpty checks and returns whether `r` is empty.
- `func (r Result) Len() int`
  Len returns the length of result list.
- `func (r Result) Size() int`
  Size is alias of function Len.
- `func (r Result) Chunk(size int) []Result`
  Chunk splits a Result into multiple Results,
- `func (r Result) Json() string`
  Json converts `r` to JSON format content.
- `func (r Result) Xml(rootTag ...string) string`
  Xml converts `r` to XML format content.
- `func (r Result) List() List`
  List converts `r` to a List.
- `func (r Result) Array(field ...string) Array`
  Array retrieves and returns specified column values as slice.
- `func (r Result) MapKeyValue(key string) map[string]Value`
  MapKeyValue converts `r` to a map[string]Value of which key is specified by `...
- `func (r Result) MapKeyStr(key string) map[string]Map`
  MapKeyStr converts `r` to a map[string]Map of which key is specified by `key`.
- `func (r Result) MapKeyInt(key string) map[int]Map`
  MapKeyInt converts `r` to a map[int]Map of which key is specified by `key`.
- `func (r Result) MapKeyUint(key string) map[uint]Map`
  MapKeyUint converts `r` to a map[uint]Map of which key is specified by `key`.
- `func (r Result) RecordKeyStr(key string) map[string]Record`
  RecordKeyStr converts `r` to a map[string]Record of which key is specified by...
- `func (r Result) RecordKeyInt(key string) map[int]Record`
  RecordKeyInt converts `r` to a map[int]Record of which key is specified by `k...
- `func (r Result) RecordKeyUint(key string) map[uint]Record`
  RecordKeyUint converts `r` to a map[uint]Record of which key is specified by ...
- `func (r Result) Structs(pointer any) err error`
  Structs converts `r` to struct slice.
- `func (r Result) ScanList(structSlicePointer any, bindToAttrName string, relationAttrNameAndFields ...string) err error`
  ScanList converts `r` to struct slice which contains other complex struct att...

### Map

**类型**: type

### List

**类型**: type

### CatchSQLManager

**类型**: struct

### SelectType

**类型**: type

### InsertOption

**类型**: type

### SqlType

**类型**: type

### LocalType

**类型**: type

LocalType is a type that defines the local storage type of a field value.
It is used to specify how the field value should be processed locally.


### Config

**类型**: type

Config is the configuration management object.


### ConfigGroup

**类型**: type

ConfigGroup is a slice of configuration node for specified named group.


### ConfigNode

**类型**: struct

ConfigNode is configuration for one node.


### Role

**类型**: type

### Propagation

**类型**: type

Propagation defines transaction propagation behavior.


### TxOptions

**类型**: struct

TxOptions defines options for transaction control.


### TXCore

**类型**: struct

TXCore is the struct for transaction management.


**方法**:

- `func (tx *TXCore) Ctx(ctx context.Context) TX`
  Ctx sets the context for current transaction.
- `func (tx *TXCore) GetCtx() context.Context`
  GetCtx returns the context for current transaction.
- `func (tx *TXCore) GetDB() DB`
  GetDB returns the DB for current transaction.
- `func (tx *TXCore) GetSqlTX() *sql.Tx`
  GetSqlTX returns the underlying transaction object for current transaction.
- `func (tx *TXCore) Commit() error`
  Commit commits current transaction.
- `func (tx *TXCore) Rollback() error`
  Rollback aborts current transaction.
- `func (tx *TXCore) IsClosed() bool`
  IsClosed checks and returns this transaction has already been committed or ro...
- `func (tx *TXCore) Begin() error`
  Begin starts a nested transaction procedure.
- `func (tx *TXCore) SavePoint(point string) error`
  SavePoint performs `SAVEPOINT xxx` SQL statement that saves transaction at cu...
- `func (tx *TXCore) RollbackTo(point string) error`
  RollbackTo performs `ROLLBACK TO SAVEPOINT xxx` SQL statement that rollbacks ...
- `func (tx *TXCore) Transaction(ctx context.Context, f func) err error`
  Transaction wraps the transaction logic using function `f`.
- `func (tx *TXCore) TransactionWithOptions(ctx context.Context, opts TxOptions, f func) err error`
  TransactionWithOptions wraps the transaction logic with propagation options u...
- `func (tx *TXCore) Query(sql string, args ...any) (result Result, err error)`
  Query does query operation on transaction.
- `func (tx *TXCore) Exec(sql string, args ...any) (sql.Result, error)`
  Exec does none query operation on transaction.
- `func (tx *TXCore) Prepare(sql string) (*Stmt, error)`
  Prepare creates a prepared statement for later queries or executions.
- `func (tx *TXCore) GetAll(sql string, args ...any) (Result, error)`
  GetAll queries and returns data records from database.
- `func (tx *TXCore) GetOne(sql string, args ...any) (Record, error)`
  GetOne queries and returns one record from database.
- `func (tx *TXCore) GetStruct(obj any, sql string, args ...any) error`
  GetStruct queries one record from database and converts it to given struct.
- `func (tx *TXCore) GetStructs(objPointerSlice any, sql string, args ...any) error`
  GetStructs queries records from database and converts them to given struct.
- `func (tx *TXCore) GetScan(pointer any, sql string, args ...any) error`
  GetScan queries one or more records from database and converts them to given ...
- `func (tx *TXCore) GetValue(sql string, args ...any) (Value, error)`
  GetValue queries and returns the field value from database.
- `func (tx *TXCore) GetCount(sql string, args ...any) (int64, error)`
  GetCount queries and returns the count from database.
- `func (tx *TXCore) Insert(table string, data any, batch ...int) (sql.Result, error)`
  Insert does "INSERT INTO ..." statement for the table.
- `func (tx *TXCore) InsertIgnore(table string, data any, batch ...int) (sql.Result, error)`
  InsertIgnore does "INSERT IGNORE INTO ..." statement for the table.
- `func (tx *TXCore) InsertAndGetId(table string, data any, batch ...int) (int64, error)`
  InsertAndGetId performs action Insert and returns the last insert id that aut...
- `func (tx *TXCore) Replace(table string, data any, batch ...int) (sql.Result, error)`
  Replace does "REPLACE INTO ..." statement for the table.
- `func (tx *TXCore) Save(table string, data any, batch ...int) (sql.Result, error)`
  Save does "INSERT INTO ... ON DUPLICATE KEY UPDATE..." statement for the table.
- `func (tx *TXCore) Update(table string, data any, condition any, args ...any) (sql.Result, error)`
  Update does "UPDATE ... " statement for the table.
- `func (tx *TXCore) Delete(table string, condition any, args ...any) (sql.Result, error)`
  Delete does "DELETE FROM ... " statement for the table.
- `func (tx *TXCore) QueryContext(ctx context.Context, sql string, args ...any) (*sql.Rows, error)`
  QueryContext implements interface function Link.QueryContext.
- `func (tx *TXCore) ExecContext(ctx context.Context, sql string, args ...any) (sql.Result, error)`
  ExecContext implements interface function Link.ExecContext.
- `func (tx *TXCore) PrepareContext(ctx context.Context, sql string) (*sql.Stmt, error)`
  PrepareContext implements interface function Link.PrepareContext.
- `func (tx *TXCore) IsOnMaster() bool`
  IsOnMaster implements interface function Link.IsOnMaster.
- `func (tx *TXCore) IsTransaction() bool`
  IsTransaction implements interface function Link.IsTransaction.
- `func (tx *TXCore) Raw(rawSql string, args ...any) *Model`
- `func (tx *TXCore) Model(tableNameQueryOrStruct ...any) *Model`
  Model acts like Core.Model except it operates on transaction.
- `func (tx *TXCore) With(object any) *Model`
  With acts like Core.With except it operates on transaction.

### DriverDefault

**类型**: struct

DriverDefault is the default driver for mysql database, which does nothing.


**方法**:

- `func (d *DriverDefault) New(core *Core, node *ConfigNode) (DB, error)`
  New creates and returns a database object for mysql.
- `func (d *DriverDefault) Open(config *ConfigNode) (db *sql.DB, err error)`
  Open creates and returns an underlying sql.DB object for mysql.
- `func (d *DriverDefault) PingMaster() error`
  PingMaster pings the master node to check authentication or keeps the connect...
- `func (d *DriverDefault) PingSlave() error`
  PingSlave pings the slave node to check authentication or keeps the connectio...

### DriverWrapper

**类型**: struct

DriverWrapper is a driver wrapper for extending features with embedded driver.


**方法**:

- `func (d *DriverWrapper) New(core *Core, node *ConfigNode) (DB, error)`
  New creates and returns a database object for mysql.

### DriverWrapperDB

**类型**: struct

DriverWrapperDB is a DB wrapper for extending features with embedded DB.


**方法**:

- `func (d *DriverWrapperDB) Open(node *ConfigNode) (db *sql.DB, err error)`
  Open creates and returns an underlying sql.DB object for pgsql.
- `func (d *DriverWrapperDB) Tables(ctx context.Context, schema ...string) (tables []string, err error)`
  Tables retrieves and returns the tables of current schema.
- `func (d *DriverWrapperDB) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*TableField, err error)`
  TableFields retrieves and returns the fields' information of specified table ...
- `func (d *DriverWrapperDB) DoInsert(ctx context.Context, link Link, table string, list List, option DoInsertOption) (result sql.Result, err error)`
  DoInsert inserts or updates data for given table.

### Model

**类型**: struct

Model is core struct implementing the DAO for ORM.


**方法**:

- `func (m *Model) Raw(rawSql string, args ...any) *Model`
  Raw sets current model as a raw sql model.
- `func (m *Model) Partition(partitions ...string) *Model`
  Partition sets Partition name.
- `func (m *Model) Ctx(ctx context.Context) *Model`
  Ctx sets the context for current operation.
- `func (m *Model) GetCtx() context.Context`
  GetCtx returns the context for current Model.
- `func (m *Model) As(as string) *Model`
  As sets an alias name for current table.
- `func (m *Model) DB(db DB) *Model`
  DB sets/changes the db object for current operation.
- `func (m *Model) TX(tx TX) *Model`
  TX sets/changes the transaction for current operation.
- `func (m *Model) Schema(schema string) *Model`
  Schema sets the schema for current operation.
- `func (m *Model) Clone() *Model`
  Clone creates and returns a new model which is a Clone of current model.
- `func (m *Model) Master() *Model`
  Master marks the following operation on master node.
- `func (m *Model) Slave() *Model`
  Slave marks the following operation on slave node.
- `func (m *Model) Safe(safe ...bool) *Model`
  Safe marks this model safe or unsafe. If safe is true, it clones and returns ...
- `func (m *Model) Args(args ...any) *Model`
  Args sets custom arguments for model operation.
- `func (m *Model) Handler(handlers ...ModelHandler) *Model`
  Handler calls each of `handlers` on current Model and returns a new Model.
- `func (m *Model) Builder() *WhereBuilder`
  Builder creates and returns a WhereBuilder. Please note that the builder is c...
- `func (m *Model) Cache(option CacheOption) *Model`
  Cache sets the cache feature for the model. It caches the result of the sql, ...
- `func (m *Model) PageCache(countOption CacheOption, dataOption CacheOption) *Model`
  PageCache sets the cache feature for pagination queries. It allows to configure
- `func (m *Model) Delete(where ...any) (result sql.Result, err error)`
  Delete does "DELETE FROM ... " statement for the model.
- `func (m *Model) Fields(fieldNamesOrMapStruct ...any) *Model`
  Fields appends `fieldNamesOrMapStruct` to the operation fields of the model, ...
- `func (m *Model) FieldsPrefix(prefixOrAlias string, fieldNamesOrMapStruct ...any) *Model`
  FieldsPrefix performs as function Fields but add extra prefix for each field.
- `func (m *Model) FieldsEx(fieldNamesOrMapStruct ...any) *Model`
  FieldsEx appends `fieldNamesOrMapStruct` to the excluded operation fields of ...
- `func (m *Model) FieldsExPrefix(prefixOrAlias string, fieldNamesOrMapStruct ...any) *Model`
  FieldsExPrefix performs as function FieldsEx but add extra prefix for each fi...
- `func (m *Model) FieldCount(column string, as ...string) *Model`
  FieldCount formats and appends commonly used field `COUNT(column)` to the sel...
- `func (m *Model) FieldSum(column string, as ...string) *Model`
  FieldSum formats and appends commonly used field `SUM(column)` to the select ...
- `func (m *Model) FieldMin(column string, as ...string) *Model`
  FieldMin formats and appends commonly used field `MIN(column)` to the select ...
- `func (m *Model) FieldMax(column string, as ...string) *Model`
  FieldMax formats and appends commonly used field `MAX(column)` to the select ...
- `func (m *Model) FieldAvg(column string, as ...string) *Model`
  FieldAvg formats and appends commonly used field `AVG(column)` to the select ...
- `func (m *Model) GetFieldsStr(prefix ...string) string`
  GetFieldsStr retrieves and returns all fields from the table, joined with cha...
- `func (m *Model) GetFieldsExStr(fields string, prefix ...string) (string, error)`
  GetFieldsExStr retrieves and returns fields which are not in parameter `field...
- `func (m *Model) HasField(field string) (bool, error)`
  HasField determine whether the field exists in the table.
- `func (m *Model) Hook(hook HookHandler) *Model`
  Hook sets the hook functions for current model.
- `func (m *Model) Batch(batch int) *Model`
  Batch sets the batch operation number for the model.
- `func (m *Model) Data(data ...any) *Model`
  Data sets the operation data for the model.
- `func (m *Model) OnConflict(onConflict ...any) *Model`
  OnConflict sets the primary key or index when columns conflicts occurs.
- `func (m *Model) OnDuplicate(onDuplicate ...any) *Model`
  OnDuplicate sets the operations when columns conflicts occurs.
- `func (m *Model) OnDuplicateEx(onDuplicateEx ...any) *Model`
  OnDuplicateEx sets the excluding columns for operations when columns conflict...
- `func (m *Model) Insert(data ...any) (result sql.Result, err error)`
  Insert does "INSERT INTO ..." statement for the model.
- `func (m *Model) InsertAndGetId(data ...any) (lastInsertId int64, err error)`
  InsertAndGetId performs action Insert and returns the last insert id that aut...
- `func (m *Model) InsertIgnore(data ...any) (result sql.Result, err error)`
  InsertIgnore does "INSERT IGNORE INTO ..." statement for the model.
- `func (m *Model) Replace(data ...any) (result sql.Result, err error)`
  Replace does "REPLACE INTO ..." statement for the model.
- `func (m *Model) Save(data ...any) (result sql.Result, err error)`
  Save does "INSERT INTO ... ON DUPLICATE KEY UPDATE..." statement for the model.
- `func (m *Model) LeftJoin(tableOrSubQueryAndJoinConditions ...string) *Model`
  LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
- `func (m *Model) RightJoin(tableOrSubQueryAndJoinConditions ...string) *Model`
  RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
- `func (m *Model) InnerJoin(tableOrSubQueryAndJoinConditions ...string) *Model`
  InnerJoin does "INNER JOIN ... ON ..." statement on the model.
- `func (m *Model) LeftJoinOnField(table string, field string) *Model`
  LeftJoinOnField performs as LeftJoin, but it joins both tables with the `same...
- `func (m *Model) RightJoinOnField(table string, field string) *Model`
  RightJoinOnField performs as RightJoin, but it joins both tables with the `sa...
- `func (m *Model) InnerJoinOnField(table string, field string) *Model`
  InnerJoinOnField performs as InnerJoin, but it joins both tables with the `sa...
- `func (m *Model) LeftJoinOnFields(table string, firstField string, operator string, secondField string) *Model`
  LeftJoinOnFields performs as LeftJoin. It specifies different fields and comp...
- `func (m *Model) RightJoinOnFields(table string, firstField string, operator string, secondField string) *Model`
  RightJoinOnFields performs as RightJoin. It specifies different fields and co...
- `func (m *Model) InnerJoinOnFields(table string, firstField string, operator string, secondField string) *Model`
  InnerJoinOnFields performs as InnerJoin. It specifies different fields and co...
- `func (m *Model) Lock(lockClause string) *Model`
  Lock sets a custom lock clause for the current operation.
- `func (m *Model) LockUpdate() *Model`
  LockUpdate sets the lock for update for current operation.
- `func (m *Model) LockUpdateSkipLocked() *Model`
  LockUpdateSkipLocked sets the lock for update with skip locked behavior for c...
- `func (m *Model) LockShared() *Model`
  LockShared sets the lock in share mode for current operation.
- `func (m *Model) OmitEmpty() *Model`
  OmitEmpty sets optionOmitEmpty option for the model, which automatically filers
- `func (m *Model) OmitEmptyWhere() *Model`
  OmitEmptyWhere sets optionOmitEmptyWhere option for the model, which automati...
- `func (m *Model) OmitEmptyData() *Model`
  OmitEmptyData sets optionOmitEmptyData option for the model, which automatica...
- `func (m *Model) OmitNil() *Model`
  OmitNil sets optionOmitNil option for the model, which automatically filers
- `func (m *Model) OmitNilWhere() *Model`
  OmitNilWhere sets optionOmitNilWhere option for the model, which automaticall...
- `func (m *Model) OmitNilData() *Model`
  OmitNilData sets optionOmitNilData option for the model, which automatically ...
- `func (m *Model) Order(orderBy ...any) *Model`
  Order sets the "ORDER BY" statement for the model.
- `func (m *Model) OrderAsc(column string) *Model`
  OrderAsc sets the "ORDER BY xxx ASC" statement for the model.
- `func (m *Model) OrderDesc(column string) *Model`
  OrderDesc sets the "ORDER BY xxx DESC" statement for the model.
- `func (m *Model) OrderRandom() *Model`
  OrderRandom sets the "ORDER BY RANDOM()" statement for the model.
- `func (m *Model) Group(groupBy ...string) *Model`
  Group sets the "GROUP BY" statement for the model.
- `func (m *Model) All(where ...any) (Result, error)`
  All does "SELECT FROM ..." statement for the model.
- `func (m *Model) AllAndCount(useFieldForCount bool) (result Result, totalCount int, err error)`
  AllAndCount retrieves all records and the total count of records from the model.
- `func (m *Model) Chunk(size int, handler ChunkHandler)`
  Chunk iterates the query result with given `size` and `handler` function.
- `func (m *Model) One(where ...any) (Record, error)`
  One retrieves one record from table and returns the result as map type.
- `func (m *Model) Array(fieldsAndWhere ...any) (Array, error)`
  Array queries and returns data values as slice from database.
- `func (m *Model) Scan(pointer any, where ...any) error`
  Scan automatically calls Struct or Structs function according to the type of ...
- `func (m *Model) ScanAndCount(pointer any, totalCount *int, useFieldForCount bool) err error`
  ScanAndCount scans a single record or record array that matches the given con...
- `func (m *Model) ScanList(structSlicePointer any, bindToAttrName string, relationAttrNameAndFields ...string) err error`
  ScanList converts `r` to struct slice which contains other complex struct att...
- `func (m *Model) Value(fieldsAndWhere ...any) (Value, error)`
  Value retrieves a specified record value from table and returns the result as...
- `func (m *Model) Count(where ...any) (int, error)`
  Count does "SELECT COUNT(x) FROM ..." statement for the model.
- `func (m *Model) Exist(where ...any) (bool, error)`
  Exist does "SELECT 1 FROM ... LIMIT 1" statement for the model.
- `func (m *Model) CountColumn(column string) (int, error)`
  CountColumn does "SELECT COUNT(x) FROM ..." statement for the model.
- `func (m *Model) Min(column string) (float64, error)`
  Min does "SELECT MIN(x) FROM ..." statement for the model.
- `func (m *Model) Max(column string) (float64, error)`
  Max does "SELECT MAX(x) FROM ..." statement for the model.
- `func (m *Model) Avg(column string) (float64, error)`
  Avg does "SELECT AVG(x) FROM ..." statement for the model.
- `func (m *Model) Sum(column string) (float64, error)`
  Sum does "SELECT SUM(x) FROM ..." statement for the model.
- `func (m *Model) Union(unions ...*Model) *Model`
  Union does "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." statement ...
- `func (m *Model) UnionAll(unions ...*Model) *Model`
  UnionAll does "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." sta...
- `func (m *Model) Limit(limit ...int) *Model`
  Limit sets the "LIMIT" statement for the model.
- `func (m *Model) Offset(offset int) *Model`
  Offset sets the "OFFSET" statement for the model.
- `func (m *Model) Distinct() *Model`
  Distinct forces the query to only return distinct results.
- `func (m *Model) Page(page int, limit int) *Model`
  Page sets the paging number for the model.
- `func (m *Model) Having(having any, args ...any) *Model`
  Having sets the having statement for the model.
- `func (m *Model) Sharding(config ShardingConfig) *Model`
  Sharding creates a sharding model with given sharding configuration.
- `func (m *Model) ShardingValue(value any) *Model`
  ShardingValue sets the sharding value for routing
- `func (m *Model) SoftTime(option SoftTimeOption) *Model`
  SoftTime sets the SoftTimeOption to customize soft time feature for Model.
- `func (m *Model) Unscoped() *Model`
  Unscoped disables the soft time feature for insert, update and delete operati...
- `func (m *Model) Transaction(ctx context.Context, f func) err error`
  Transaction wraps the transaction logic using function `f`.
- `func (m *Model) TransactionWithOptions(ctx context.Context, opts TxOptions, f func) err error`
  TransactionWithOptions executes transaction with options.
- `func (m *Model) Update(dataAndWhere ...any) (result sql.Result, err error)`
  Update does "UPDATE ... " statement for the model.
- `func (m *Model) UpdateAndGetAffected(dataAndWhere ...any) (affected int64, err error)`
  UpdateAndGetAffected performs update statement and returns the affected rows ...
- `func (m *Model) Increment(column string, amount any) (sql.Result, error)`
  Increment increments a column's value by a given amount.
- `func (m *Model) Decrement(column string, amount any) (sql.Result, error)`
  Decrement decrements a column's value by a given amount.
- `func (m *Model) QuoteWord(s string) string`
  QuoteWord checks given string `s` a word,
- `func (m *Model) TableFields(tableStr string, schema ...string) (fields map[string]*TableField, err error)`
  TableFields retrieves and returns the fields' information of specified table ...
- `func (m *Model) Where(where any, args ...any) *Model`
  Where sets the condition statement for the builder. The parameter `where` can...
- `func (m *Model) Wheref(format string, args ...any) *Model`
  Wheref builds condition string using fmt.Sprintf and arguments.
- `func (m *Model) WherePri(where any, args ...any) *Model`
  WherePri does the same logic as Model.Where except that if the parameter `where`
- `func (m *Model) WhereLT(column string, value any) *Model`
  WhereLT builds `column < value` statement.
- `func (m *Model) WhereLTE(column string, value any) *Model`
  WhereLTE builds `column <= value` statement.
- `func (m *Model) WhereGT(column string, value any) *Model`
  WhereGT builds `column > value` statement.
- `func (m *Model) WhereGTE(column string, value any) *Model`
  WhereGTE builds `column >= value` statement.
- `func (m *Model) WhereBetween(column string, min any, max any) *Model`
  WhereBetween builds `column BETWEEN min AND max` statement.
- `func (m *Model) WhereLike(column string, like string) *Model`
  WhereLike builds `column LIKE like` statement.
- `func (m *Model) WhereIn(column string, in any) *Model`
  WhereIn builds `column IN (in)` statement.
- `func (m *Model) WhereNull(columns ...string) *Model`
  WhereNull builds `columns[0] IS NULL AND columns[1] IS NULL ...` statement.
- `func (m *Model) WhereNotBetween(column string, min any, max any) *Model`
  WhereNotBetween builds `column NOT BETWEEN min AND max` statement.
- `func (m *Model) WhereNotLike(column string, like any) *Model`
  WhereNotLike builds `column NOT LIKE like` statement.
- `func (m *Model) WhereNot(column string, value any) *Model`
  WhereNot builds `column != value` statement.
- `func (m *Model) WhereNotIn(column string, in any) *Model`
  WhereNotIn builds `column NOT IN (in)` statement.
- `func (m *Model) WhereNotNull(columns ...string) *Model`
  WhereNotNull builds `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` s...
- `func (m *Model) WhereExists(subQuery *Model) *Model`
  WhereExists builds `EXISTS (subQuery)` statement.
- `func (m *Model) WhereNotExists(subQuery *Model) *Model`
  WhereNotExists builds `NOT EXISTS (subQuery)` statement.
- `func (m *Model) WherePrefix(prefix string, where any, args ...any) *Model`
  WherePrefix performs as Where, but it adds prefix to each field in where stat...
- `func (m *Model) WherePrefixLT(prefix string, column string, value any) *Model`
  WherePrefixLT builds `prefix.column < value` statement.
- `func (m *Model) WherePrefixLTE(prefix string, column string, value any) *Model`
  WherePrefixLTE builds `prefix.column <= value` statement.
- `func (m *Model) WherePrefixGT(prefix string, column string, value any) *Model`
  WherePrefixGT builds `prefix.column > value` statement.
- `func (m *Model) WherePrefixGTE(prefix string, column string, value any) *Model`
  WherePrefixGTE builds `prefix.column >= value` statement.
- `func (m *Model) WherePrefixBetween(prefix string, column string, min any, max any) *Model`
  WherePrefixBetween builds `prefix.column BETWEEN min AND max` statement.
- `func (m *Model) WherePrefixLike(prefix string, column string, like any) *Model`
  WherePrefixLike builds `prefix.column LIKE like` statement.
- `func (m *Model) WherePrefixIn(prefix string, column string, in any) *Model`
  WherePrefixIn builds `prefix.column IN (in)` statement.
- `func (m *Model) WherePrefixNull(prefix string, columns ...string) *Model`
  WherePrefixNull builds `prefix.columns[0] IS NULL AND prefix.columns[1] IS NU...
- `func (m *Model) WherePrefixNotBetween(prefix string, column string, min any, max any) *Model`
  WherePrefixNotBetween builds `prefix.column NOT BETWEEN min AND max` statement.
- `func (m *Model) WherePrefixNotLike(prefix string, column string, like any) *Model`
  WherePrefixNotLike builds `prefix.column NOT LIKE like` statement.
- `func (m *Model) WherePrefixNot(prefix string, column string, value any) *Model`
  WherePrefixNot builds `prefix.column != value` statement.
- `func (m *Model) WherePrefixNotIn(prefix string, column string, in any) *Model`
  WherePrefixNotIn builds `prefix.column NOT IN (in)` statement.
- `func (m *Model) WherePrefixNotNull(prefix string, columns ...string) *Model`
  WherePrefixNotNull builds `prefix.columns[0] IS NOT NULL AND prefix.columns[1...
- `func (m *Model) WhereOr(where any, args ...any) *Model`
  WhereOr adds "OR" condition to the where statement.
- `func (m *Model) WhereOrf(format string, args ...any) *Model`
  WhereOrf builds `OR` condition string using fmt.Sprintf and arguments.
- `func (m *Model) WhereOrLT(column string, value any) *Model`
  WhereOrLT builds `column < value` statement in `OR` conditions.
- `func (m *Model) WhereOrLTE(column string, value any) *Model`
  WhereOrLTE builds `column <= value` statement in `OR` conditions.
- `func (m *Model) WhereOrGT(column string, value any) *Model`
  WhereOrGT builds `column > value` statement in `OR` conditions.
- `func (m *Model) WhereOrGTE(column string, value any) *Model`
  WhereOrGTE builds `column >= value` statement in `OR` conditions.
- `func (m *Model) WhereOrBetween(column string, min any, max any) *Model`
  WhereOrBetween builds `column BETWEEN min AND max` statement in `OR` conditions.
- `func (m *Model) WhereOrLike(column string, like any) *Model`
  WhereOrLike builds `column LIKE like` statement in `OR` conditions.
- `func (m *Model) WhereOrIn(column string, in any) *Model`
  WhereOrIn builds `column IN (in)` statement in `OR` conditions.
- `func (m *Model) WhereOrNull(columns ...string) *Model`
  WhereOrNull builds `columns[0] IS NULL OR columns[1] IS NULL ...` statement i...
- `func (m *Model) WhereOrNotBetween(column string, min any, max any) *Model`
  WhereOrNotBetween builds `column NOT BETWEEN min AND max` statement in `OR` c...
- `func (m *Model) WhereOrNotLike(column string, like any) *Model`
  WhereOrNotLike builds `column NOT LIKE 'like'` statement in `OR` conditions.
- `func (m *Model) WhereOrNot(column string, value any) *Model`
  WhereOrNot builds `column != value` statement.
- `func (m *Model) WhereOrNotIn(column string, in any) *Model`
  WhereOrNotIn builds `column NOT IN (in)` statement.
- `func (m *Model) WhereOrNotNull(columns ...string) *Model`
  WhereOrNotNull builds `columns[0] IS NOT NULL OR columns[1] IS NOT NULL ...` ...
- `func (m *Model) WhereOrPrefix(prefix string, where any, args ...any) *Model`
  WhereOrPrefix performs as WhereOr, but it adds prefix to each field in where ...
- `func (m *Model) WhereOrPrefixLT(prefix string, column string, value any) *Model`
  WhereOrPrefixLT builds `prefix.column < value` statement in `OR` conditions.
- `func (m *Model) WhereOrPrefixLTE(prefix string, column string, value any) *Model`
  WhereOrPrefixLTE builds `prefix.column <= value` statement in `OR` conditions.
- `func (m *Model) WhereOrPrefixGT(prefix string, column string, value any) *Model`
  WhereOrPrefixGT builds `prefix.column > value` statement in `OR` conditions.
- `func (m *Model) WhereOrPrefixGTE(prefix string, column string, value any) *Model`
  WhereOrPrefixGTE builds `prefix.column >= value` statement in `OR` conditions.
- `func (m *Model) WhereOrPrefixBetween(prefix string, column string, min any, max any) *Model`
  WhereOrPrefixBetween builds `prefix.column BETWEEN min AND max` statement in ...
- `func (m *Model) WhereOrPrefixLike(prefix string, column string, like any) *Model`
  WhereOrPrefixLike builds `prefix.column LIKE like` statement in `OR` conditions.
- `func (m *Model) WhereOrPrefixIn(prefix string, column string, in any) *Model`
  WhereOrPrefixIn builds `prefix.column IN (in)` statement in `OR` conditions.
- `func (m *Model) WhereOrPrefixNull(prefix string, columns ...string) *Model`
  WhereOrPrefixNull builds `prefix.columns[0] IS NULL OR prefix.columns[1] IS N...
- `func (m *Model) WhereOrPrefixNotBetween(prefix string, column string, min any, max any) *Model`
  WhereOrPrefixNotBetween builds `prefix.column NOT BETWEEN min AND max` statem...
- `func (m *Model) WhereOrPrefixNotLike(prefix string, column string, like any) *Model`
  WhereOrPrefixNotLike builds `prefix.column NOT LIKE like` statement in `OR` c...
- `func (m *Model) WhereOrPrefixNotIn(prefix string, column string, in any) *Model`
  WhereOrPrefixNotIn builds `prefix.column NOT IN (in)` statement.
- `func (m *Model) WhereOrPrefixNotNull(prefix string, columns ...string) *Model`
  WhereOrPrefixNotNull builds `prefix.columns[0] IS NOT NULL OR prefix.columns[...
- `func (m *Model) WhereOrPrefixNot(prefix string, column string, value any) *Model`
  WhereOrPrefixNot builds `prefix.column != value` statement in `OR` conditions.
- `func (m *Model) With(objects ...any) *Model`
  With creates and returns an ORM model based on metadata of given object.
- `func (m *Model) WithAll() *Model`
  WithAll enables model association operations on all objects that have "with" ...

### ModelHandler

**类型**: type

ModelHandler is a function that handles given Model and returns a new Model that is custom modified.


### ChunkHandler

**类型**: type

ChunkHandler is a function that is used in function Chunk, which handles given Result and error.
It returns true if it wants to continue chunking, or else it returns false to stop chunking.


### WhereBuilder

**类型**: struct

WhereBuilder holds multiple where conditions in a group.


**方法**:

- `func (b *WhereBuilder) Clone() *WhereBuilder`
  Clone clones and returns a WhereBuilder that is a copy of current one.
- `func (b *WhereBuilder) Build() (conditionWhere string, conditionArgs []any)`
  Build builds current WhereBuilder and returns the condition string and parame...
- `func (b *WhereBuilder) Where(where any, args ...any) *WhereBuilder`
  Where sets the condition statement for the builder. The parameter `where` can...
- `func (b *WhereBuilder) Wheref(format string, args ...any) *WhereBuilder`
  Wheref builds condition string using fmt.Sprintf and arguments.
- `func (b *WhereBuilder) WherePri(where any, args ...any) *WhereBuilder`
  WherePri does the same logic as Model.Where except that if the parameter `where`
- `func (b *WhereBuilder) WhereLT(column string, value any) *WhereBuilder`
  WhereLT builds `column < value` statement.
- `func (b *WhereBuilder) WhereLTE(column string, value any) *WhereBuilder`
  WhereLTE builds `column <= value` statement.
- `func (b *WhereBuilder) WhereGT(column string, value any) *WhereBuilder`
  WhereGT builds `column > value` statement.
- `func (b *WhereBuilder) WhereGTE(column string, value any) *WhereBuilder`
  WhereGTE builds `column >= value` statement.
- `func (b *WhereBuilder) WhereBetween(column string, min any, max any) *WhereBuilder`
  WhereBetween builds `column BETWEEN min AND max` statement.
- `func (b *WhereBuilder) WhereLike(column string, like string) *WhereBuilder`
  WhereLike builds `column LIKE like` statement.
- `func (b *WhereBuilder) WhereIn(column string, in any) *WhereBuilder`
  WhereIn builds `column IN (in)` statement.
- `func (b *WhereBuilder) WhereNull(columns ...string) *WhereBuilder`
  WhereNull builds `columns[0] IS NULL AND columns[1] IS NULL ...` statement.
- `func (b *WhereBuilder) WhereNotBetween(column string, min any, max any) *WhereBuilder`
  WhereNotBetween builds `column NOT BETWEEN min AND max` statement.
- `func (b *WhereBuilder) WhereNotLike(column string, like any) *WhereBuilder`
  WhereNotLike builds `column NOT LIKE like` statement.
- `func (b *WhereBuilder) WhereNot(column string, value any) *WhereBuilder`
  WhereNot builds `column != value` statement.
- `func (b *WhereBuilder) WhereNotIn(column string, in any) *WhereBuilder`
  WhereNotIn builds `column NOT IN (in)` statement.
- `func (b *WhereBuilder) WhereNotNull(columns ...string) *WhereBuilder`
  WhereNotNull builds `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` s...
- `func (b *WhereBuilder) WhereExists(subQuery *Model) *WhereBuilder`
  WhereExists builds `EXISTS (subQuery)` statement.
- `func (b *WhereBuilder) WhereNotExists(subQuery *Model) *WhereBuilder`
  WhereNotExists builds `NOT EXISTS (subQuery)` statement.
- `func (b *WhereBuilder) WherePrefix(prefix string, where any, args ...any) *WhereBuilder`
  WherePrefix performs as Where, but it adds prefix to each field in where stat...
- `func (b *WhereBuilder) WherePrefixLT(prefix string, column string, value any) *WhereBuilder`
  WherePrefixLT builds `prefix.column < value` statement.
- `func (b *WhereBuilder) WherePrefixLTE(prefix string, column string, value any) *WhereBuilder`
  WherePrefixLTE builds `prefix.column <= value` statement.
- `func (b *WhereBuilder) WherePrefixGT(prefix string, column string, value any) *WhereBuilder`
  WherePrefixGT builds `prefix.column > value` statement.
- `func (b *WhereBuilder) WherePrefixGTE(prefix string, column string, value any) *WhereBuilder`
  WherePrefixGTE builds `prefix.column >= value` statement.
- `func (b *WhereBuilder) WherePrefixBetween(prefix string, column string, min any, max any) *WhereBuilder`
  WherePrefixBetween builds `prefix.column BETWEEN min AND max` statement.
- `func (b *WhereBuilder) WherePrefixLike(prefix string, column string, like any) *WhereBuilder`
  WherePrefixLike builds `prefix.column LIKE like` statement.
- `func (b *WhereBuilder) WherePrefixIn(prefix string, column string, in any) *WhereBuilder`
  WherePrefixIn builds `prefix.column IN (in)` statement.
- `func (b *WhereBuilder) WherePrefixNull(prefix string, columns ...string) *WhereBuilder`
  WherePrefixNull builds `prefix.columns[0] IS NULL AND prefix.columns[1] IS NU...
- `func (b *WhereBuilder) WherePrefixNotBetween(prefix string, column string, min any, max any) *WhereBuilder`
  WherePrefixNotBetween builds `prefix.column NOT BETWEEN min AND max` statement.
- `func (b *WhereBuilder) WherePrefixNotLike(prefix string, column string, like any) *WhereBuilder`
  WherePrefixNotLike builds `prefix.column NOT LIKE like` statement.
- `func (b *WhereBuilder) WherePrefixNot(prefix string, column string, value any) *WhereBuilder`
  WherePrefixNot builds `prefix.column != value` statement.
- `func (b *WhereBuilder) WherePrefixNotIn(prefix string, column string, in any) *WhereBuilder`
  WherePrefixNotIn builds `prefix.column NOT IN (in)` statement.
- `func (b *WhereBuilder) WherePrefixNotNull(prefix string, columns ...string) *WhereBuilder`
  WherePrefixNotNull builds `prefix.columns[0] IS NOT NULL AND prefix.columns[1...
- `func (b *WhereBuilder) WhereOr(where any, args ...any) *WhereBuilder`
  WhereOr adds "OR" condition to the where statement.
- `func (b *WhereBuilder) WhereOrf(format string, args ...any) *WhereBuilder`
  WhereOrf builds `OR` condition string using fmt.Sprintf and arguments.
- `func (b *WhereBuilder) WhereOrNot(column string, value any) *WhereBuilder`
  WhereOrNot builds `column != value` statement in `OR` conditions.
- `func (b *WhereBuilder) WhereOrLT(column string, value any) *WhereBuilder`
  WhereOrLT builds `column < value` statement in `OR` conditions.
- `func (b *WhereBuilder) WhereOrLTE(column string, value any) *WhereBuilder`
  WhereOrLTE builds `column <= value` statement in `OR` conditions.
- `func (b *WhereBuilder) WhereOrGT(column string, value any) *WhereBuilder`
  WhereOrGT builds `column > value` statement in `OR` conditions.
- `func (b *WhereBuilder) WhereOrGTE(column string, value any) *WhereBuilder`
  WhereOrGTE builds `column >= value` statement in `OR` conditions.
- `func (b *WhereBuilder) WhereOrBetween(column string, min any, max any) *WhereBuilder`
  WhereOrBetween builds `column BETWEEN min AND max` statement in `OR` conditions.
- `func (b *WhereBuilder) WhereOrLike(column string, like any) *WhereBuilder`
  WhereOrLike builds `column LIKE 'like'` statement in `OR` conditions.
- `func (b *WhereBuilder) WhereOrIn(column string, in any) *WhereBuilder`
  WhereOrIn builds `column IN (in)` statement in `OR` conditions.
- `func (b *WhereBuilder) WhereOrNull(columns ...string) *WhereBuilder`
  WhereOrNull builds `columns[0] IS NULL OR columns[1] IS NULL ...` statement i...
- `func (b *WhereBuilder) WhereOrNotBetween(column string, min any, max any) *WhereBuilder`
  WhereOrNotBetween builds `column NOT BETWEEN min AND max` statement in `OR` c...
- `func (b *WhereBuilder) WhereOrNotLike(column string, like any) *WhereBuilder`
  WhereOrNotLike builds `column NOT LIKE like` statement in `OR` conditions.
- `func (b *WhereBuilder) WhereOrNotIn(column string, in any) *WhereBuilder`
  WhereOrNotIn builds `column NOT IN (in)` statement.
- `func (b *WhereBuilder) WhereOrNotNull(columns ...string) *WhereBuilder`
  WhereOrNotNull builds `columns[0] IS NOT NULL OR columns[1] IS NOT NULL ...` ...
- `func (b *WhereBuilder) WhereOrPrefix(prefix string, where any, args ...any) *WhereBuilder`
  WhereOrPrefix performs as WhereOr, but it adds prefix to each field in where ...
- `func (b *WhereBuilder) WhereOrPrefixNot(prefix string, column string, value any) *WhereBuilder`
  WhereOrPrefixNot builds `prefix.column != value` statement in `OR` conditions.
- `func (b *WhereBuilder) WhereOrPrefixLT(prefix string, column string, value any) *WhereBuilder`
  WhereOrPrefixLT builds `prefix.column < value` statement in `OR` conditions.
- `func (b *WhereBuilder) WhereOrPrefixLTE(prefix string, column string, value any) *WhereBuilder`
  WhereOrPrefixLTE builds `prefix.column <= value` statement in `OR` conditions.
- `func (b *WhereBuilder) WhereOrPrefixGT(prefix string, column string, value any) *WhereBuilder`
  WhereOrPrefixGT builds `prefix.column > value` statement in `OR` conditions.
- `func (b *WhereBuilder) WhereOrPrefixGTE(prefix string, column string, value any) *WhereBuilder`
  WhereOrPrefixGTE builds `prefix.column >= value` statement in `OR` conditions.
- `func (b *WhereBuilder) WhereOrPrefixBetween(prefix string, column string, min any, max any) *WhereBuilder`
  WhereOrPrefixBetween builds `prefix.column BETWEEN min AND max` statement in ...
- `func (b *WhereBuilder) WhereOrPrefixLike(prefix string, column string, like any) *WhereBuilder`
  WhereOrPrefixLike builds `prefix.column LIKE 'like'` statement in `OR` condit...
- `func (b *WhereBuilder) WhereOrPrefixIn(prefix string, column string, in any) *WhereBuilder`
  WhereOrPrefixIn builds `prefix.column IN (in)` statement in `OR` conditions.
- `func (b *WhereBuilder) WhereOrPrefixNull(prefix string, columns ...string) *WhereBuilder`
  WhereOrPrefixNull builds `prefix.columns[0] IS NULL OR prefix.columns[1] IS N...
- `func (b *WhereBuilder) WhereOrPrefixNotBetween(prefix string, column string, min any, max any) *WhereBuilder`
  WhereOrPrefixNotBetween builds `prefix.column NOT BETWEEN min AND max` statem...
- `func (b *WhereBuilder) WhereOrPrefixNotLike(prefix string, column string, like any) *WhereBuilder`
  WhereOrPrefixNotLike builds `prefix.column NOT LIKE 'like'` statement in `OR`...
- `func (b *WhereBuilder) WhereOrPrefixNotIn(prefix string, column string, in any) *WhereBuilder`
  WhereOrPrefixNotIn builds `prefix.column NOT IN (in)` statement.
- `func (b *WhereBuilder) WhereOrPrefixNotNull(prefix string, columns ...string) *WhereBuilder`
  WhereOrPrefixNotNull builds `prefix.columns[0] IS NOT NULL OR prefix.columns[...

### WhereHolder

**类型**: struct

WhereHolder is the holder for where condition preparing.


### CacheOption

**类型**: struct

CacheOption is options for model cache control in query.


### HookFuncSelect

**类型**: type

### HookFuncInsert

**类型**: type

### HookFuncUpdate

**类型**: type

### HookFuncDelete

**类型**: type

### HookHandler

**类型**: struct

HookHandler manages all supported hook functions for Model.


### HookSelectInput

**类型**: struct

HookSelectInput holds the parameters for select hook operation.
Note that, COUNT statement will also be hooked by this feature,
which is usually not be interesting for upper business hook handler.


**方法**:

- `func (h *HookSelectInput) Next(ctx context.Context) (result Result, err error)`
  Next calls the next hook handler.

### HookInsertInput

**类型**: struct

HookInsertInput holds the parameters for insert hook operation.


**方法**:

- `func (h *HookInsertInput) Next(ctx context.Context) (result sql.Result, err error)`
  Next calls the next hook handler.

### HookUpdateInput

**类型**: struct

HookUpdateInput holds the parameters for update hook operation.


**方法**:

- `func (h *HookUpdateInput) Next(ctx context.Context) (result sql.Result, err error)`
  Next calls the next hook handler.

### HookDeleteInput

**类型**: struct

HookDeleteInput holds the parameters for delete hook operation.


**方法**:

- `func (h *HookDeleteInput) Next(ctx context.Context) (result sql.Result, err error)`
  Next calls the next hook handler.

### ShardingConfig

**类型**: struct

ShardingConfig defines the configuration for database/table sharding.


### ShardingSchemaConfig

**类型**: struct

ShardingSchemaConfig defines the configuration for database sharding.


### ShardingTableConfig

**类型**: struct

ShardingTableConfig defines the configuration for table sharding


### ShardingRule

**类型**: interface

ShardingRule defines the interface for sharding rules


### DefaultShardingRule

**类型**: struct

DefaultShardingRule implements a simple modulo-based sharding rule


**方法**:

- `func (r *DefaultShardingRule) SchemaName(ctx context.Context, config ShardingSchemaConfig, value any) (string, error)`
  SchemaName implements the default database sharding strategy
- `func (r *DefaultShardingRule) TableName(ctx context.Context, config ShardingTableConfig, value any) (string, error)`
  TableName implements the default table sharding strategy

### SoftTimeType

**类型**: type

SoftTimeType custom defines the soft time field type.


### SoftTimeOption

**类型**: struct

SoftTimeOption is the option to customize soft time feature for Model.


### SoftTimeFieldType

**类型**: type

SoftTimeFieldType represents different soft time field purposes.


### SqlResult

**类型**: struct

SqlResult is execution result for sql operations.
It also supports batch operation result for rowsAffected.


**方法**:

- `func (r *SqlResult) MustGetAffected() int64`
  MustGetAffected returns the affected rows count, if any error occurs, it panics.
- `func (r *SqlResult) MustGetInsertId() int64`
  MustGetInsertId returns the last insert id, if any error occurs, it panics.
- `func (r *SqlResult) RowsAffected() (int64, error)`
  RowsAffected returns the number of rows affected by an
- `func (r *SqlResult) LastInsertId() (int64, error)`
  LastInsertId returns the integer generated by the database

### Schema

**类型**: struct

Schema is a schema object from which it can then create a Model.


### Stmt

**类型**: struct

Stmt is a prepared statement.
A Stmt is safe for concurrent use by multiple goroutines.

If a Stmt is prepared on a Tx or Conn, it will be bound to a single
underlying connection forever. If the Tx or Conn closes, the Stmt will
become unusable and all operations will return an error.
If a Stmt is prepared on a DB, it will remain usable for the lifetime of the
DB. When the Stmt needs to execute on a new underlying connection, it will
prepare itself on the new connection automatically.


**方法**:

- `func (s *Stmt) ExecContext(ctx context.Context, args ...any) (sql.Result, error)`
  ExecContext executes a prepared statement with the given arguments and
- `func (s *Stmt) QueryContext(ctx context.Context, args ...any) (*sql.Rows, error)`
  QueryContext executes a prepared query statement with the given arguments
- `func (s *Stmt) QueryRowContext(ctx context.Context, args ...any) *sql.Row`
  QueryRowContext executes a prepared query statement with the given arguments.
- `func (s *Stmt) Exec(args ...any) (sql.Result, error)`
  Exec executes a prepared statement with the given arguments and
- `func (s *Stmt) Query(args ...any) (*sql.Rows, error)`
  Query executes a prepared query statement with the given arguments
- `func (s *Stmt) QueryRow(args ...any) *sql.Row`
  QueryRow executes a prepared query statement with the given arguments.
- `func (s *Stmt) Close() error`
  Close closes the statement.

## 函数

### Register

```go
func Register(name string, driver Driver) error
```

Register registers custom database driver to gdb.

### New

```go
func New(node ConfigNode) (db DB, err error)
```

New creates and returns an ORM object with given configuration node.

### NewByGroup

```go
func NewByGroup(group ...string) (db DB, err error)
```

NewByGroup creates and returns an ORM object with global configurations.
The parameter `name` specifies the configuration group name,
which is DefaultGroupName in default.

### Instance

```go
func Instance(name ...string) (db DB, err error)
```

Instance returns an instance for DB operations.
The parameter `name` specifies the configuration group name,
which is DefaultGroupName in default.

### GetConverter

```go
func GetConverter() gconv.Converter
```

GetConverter returns the internal type converter for gdb.

### SetConfig

```go
func SetConfig(config Config) error
```

SetConfig sets the global configuration for package.
It will overwrite the old configuration of package.

### SetConfigGroup

```go
func SetConfigGroup(group string, nodes ConfigGroup) error
```

SetConfigGroup sets the configuration for given group.

### AddConfigNode

```go
func AddConfigNode(group string, node ConfigNode) error
```

AddConfigNode adds one node configuration to configuration of given group.

### AddDefaultConfigNode

```go
func AddDefaultConfigNode(node ConfigNode) error
```

AddDefaultConfigNode adds one node configuration to configuration of default group.

### AddDefaultConfigGroup

```go
func AddDefaultConfigGroup(nodes ConfigGroup) error
```

AddDefaultConfigGroup adds multiple node configurations to configuration of default group.

Deprecated: Use SetDefaultConfigGroup instead.

### SetDefaultConfigGroup

```go
func SetDefaultConfigGroup(nodes ConfigGroup) error
```

SetDefaultConfigGroup sets multiple node configurations to configuration of default group.

### GetConfig

```go
func GetConfig(group string) ConfigGroup
```

GetConfig retrieves and returns the configuration of given group.

Deprecated: Use GetConfigGroup instead.

### GetConfigGroup

```go
func GetConfigGroup(group string) (ConfigGroup, error)
```

GetConfigGroup retrieves and returns the configuration of given group.
It returns an error if the group does not exist, or an empty slice if the group exists but has no nodes.

### GetAllConfig

```go
func GetAllConfig() Config
```

GetAllConfig retrieves and returns all configurations.

### SetDefaultGroup

```go
func SetDefaultGroup(name string)
```

SetDefaultGroup sets the group name for default configuration.

### GetDefaultGroup

```go
func GetDefaultGroup() string
```

GetDefaultGroup returns the { name of default configuration.

### IsConfigured

```go
func IsConfigured() bool
```

IsConfigured checks and returns whether the database configured.
It returns true if any configuration exists.

### DefaultTxOptions

```go
func DefaultTxOptions() TxOptions
```

DefaultTxOptions returns the default transaction options.

### WithTX

```go
func WithTX(ctx context.Context, tx TX) context.Context
```

WithTX injects given transaction object into context and returns a new context.

### WithoutTX

```go
func WithoutTX(ctx context.Context, group string) context.Context
```

WithoutTX removed transaction object from context and returns a new context.

### TXFromCtx

```go
func TXFromCtx(ctx context.Context, group string) TX
```

TXFromCtx retrieves and returns transaction object from context.
It is usually used in nested transaction feature, and it returns nil if it is not set previously.

### WithDB

```go
func WithDB(ctx context.Context, db DB) context.Context
```

WithDB injects given db object into context and returns a new context.

### DBFromCtx

```go
func DBFromCtx(ctx context.Context) DB
```

DBFromCtx retrieves and returns DB object from context.

### ToSQL

```go
func ToSQL(ctx context.Context, f func) (sql string, err error)
```

ToSQL formats and returns the last one of sql statements in given closure function
WITHOUT TRULY EXECUTING IT.
Be caution that, all the following sql statements should use the context object passing by function `f`.

### CatchSQL

```go
func CatchSQL(ctx context.Context, f func) (sqlArray []string, err error)
```

CatchSQL catches and returns all sql statements that are EXECUTED in given closure function.
Be caution that, all the following sql statements should use the context object passing by function `f`.

### ListItemValues

```go
func ListItemValues(list any, key any, subKey ...any) values []any
```

ListItemValues retrieves and returns the elements of all item struct/map with key `key`.
Note that the parameter `list` should be type of slice which contains elements of map or struct,
or else it returns an empty slice.

The parameter `list` supports types like:
[]map[string]any
[]map[string]sub-map
[]struct
[]struct:sub-struct
Note that the sub-map/sub-struct makes sense only if the optional parameter `subKey` is given.
See gutil.ListItemValues.

### ListItemValuesUnique

```go
func ListItemValuesUnique(list any, key string, subKey ...any) []any
```

ListItemValuesUnique retrieves and returns the unique elements of all struct/map with key `key`.
Note that the parameter `list` should be type of slice which contains elements of map or struct,
or else it returns an empty slice.
See gutil.ListItemValuesUnique.

### GetInsertOperationByOption

```go
func GetInsertOperationByOption(option InsertOption) string
```

GetInsertOperationByOption returns proper insert option with given parameter `option`.

### MapOrStructToMapDeep

```go
func MapOrStructToMapDeep(value any, omitempty bool) map[string]any
```

MapOrStructToMapDeep converts `value` to map type recursively(if attribute struct is embedded).
The parameter `value` should be type of *map/map/*struct/struct.
It supports embedded struct definition for struct.

### GetPrimaryKeyCondition

```go
func GetPrimaryKeyCondition(primary string, where ...any) newWhereCondition []any
```

GetPrimaryKeyCondition returns a new where condition by primary field name.
The optional parameter `where` is like follows:
123                             => primary=123
[]int{1, 2, 3}                  => primary IN(1,2,3)
"john"                          => primary='john'
[]string{"john", "smith"}       => primary IN('john','smith')
g.Map{"id": g.Slice{1,2,3}}     => id IN(1,2,3)
g.Map{"id": 1, "name": "john"}  => id=1 AND name='john'
etc.

Note that it returns the given `where` parameter directly if the `primary` is empty
or length of `where` > 1.

### FormatSqlWithArgs

```go
func FormatSqlWithArgs(sql string, args []any) string
```

FormatSqlWithArgs binds the arguments to the sql string and returns a complete
sql string, just for debugging.

### FormatMultiLineSqlToSingle

```go
func FormatMultiLineSqlToSingle(sql string) (string, error)
```

FormatMultiLineSqlToSingle formats sql template string into one line.

## 内部依赖

- `github.com/gogf/gf/v2`
- `container/garray`
- `container/gmap`
- `container/gset`
- `container/gtype`
- `container/gvar`
- `encoding/gbinary`
- `encoding/ghash`
- `encoding/gjson`
- `errors/gcode`
- `errors/gerror`
- `internal/empty`
- `internal/intlog`
- `internal/json`
- `internal/reflection`
- `internal/utils`
- `net/gtrace`
- `os/gcache`
- `os/gcmd`
- `os/gctx`
- `os/glog`
- `os/gstructs`
- `os/gtime`
- `text/gregex`
- `text/gstr`
- `util/gconv`
- `util/gmeta`
- `util/grand`
- `util/gtag`
- `util/guid`
- `util/gutil`

## 外部依赖

- `go.opentelemetry.io/otel`
- `go.opentelemetry.io/otel/attribute`
- `go.opentelemetry.io/otel/codes`
- `go.opentelemetry.io/otel/semconv/v1.18.0`
- `go.opentelemetry.io/otel/trace`

