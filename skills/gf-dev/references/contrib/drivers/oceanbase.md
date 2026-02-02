# oceanbase

> Package: `github.com/gogf/gf/contrib/drivers/oceanbase/v2`

```go
import "github.com/gogf/gf/contrib/drivers/oceanbase/v2"
```

## 概述

Package oceanbase implements gdb.Driver, which supports operations for database OceanBase.

## 所属模块

此包属于子模块: `github.com/gogf/gf/contrib/drivers/oceanbase/v2`

## 源文件

- `oceanbase.go`

## 类型定义

### Driver

**类型**: struct

Driver is the driver for OceanBase database.

OceanBase is a distributed relational database developed by Ant Group. It supports both MySQL and Oracle
protocol modes. This driver uses the MySQL protocol to communicate with OceanBase database in MySQL
compatibility mode.

Although OceanBase is compatible with MySQL protocol, it is packaged as a separate driver component
rather than reusing the mysql adapter directly. This design allows for future extensibility,
such as implementing OceanBase-specific features like distributed transactions or Oracle mode support.


## 函数

### New

```go
func New() gdb.Driver
```

New creates and returns a driver that implements gdb.Driver, which supports operations for OceanBase.

## 内部依赖

- `github.com/gogf/gf/contrib/drivers/mysql/v2`
- `database/gdb`
- `frame/g`

