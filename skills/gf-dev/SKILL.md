---
name: gf-dev
description: gf 模块开发指南。基于源码深度分析，提供完整的包结构和 API 参考。当用户需要使用此模块进行开发时使用。
allowed-tools: 
disable: true
---

# gf 开发技能

## 模块信息

- **主模块**: `github.com/gogf/gf/v2`
- **子模块数量**: 31
- **包数量**: 114
- **分类数量**: 16

### 模块列表

| 模块路径 | 相对位置 |
|---------|----------|
| `github.com/gogf/gf/cmd/gf/v2` | `cmd/gf` |
| `github.com/gogf/gf/contrib/config/apollo/v2` | `contrib/config/apollo` |
| `github.com/gogf/gf/contrib/config/consul/v2` | `contrib/config/consul` |
| `github.com/gogf/gf/contrib/config/kubecm/v2` | `contrib/config/kubecm` |
| `github.com/gogf/gf/contrib/config/nacos/v2` | `contrib/config/nacos` |
| `github.com/gogf/gf/contrib/config/polaris/v2` | `contrib/config/polaris` |
| `github.com/gogf/gf/contrib/drivers/clickhouse/v2` | `contrib/drivers/clickhouse` |
| `github.com/gogf/gf/contrib/drivers/dm/v2` | `contrib/drivers/dm` |
| `github.com/gogf/gf/contrib/drivers/gaussdb/v2` | `contrib/drivers/gaussdb` |
| `github.com/gogf/gf/contrib/drivers/mariadb/v2` | `contrib/drivers/mariadb` |
| `github.com/gogf/gf/contrib/drivers/mssql/v2` | `contrib/drivers/mssql` |
| `github.com/gogf/gf/contrib/drivers/mysql/v2` | `contrib/drivers/mysql` |
| `github.com/gogf/gf/contrib/drivers/oceanbase/v2` | `contrib/drivers/oceanbase` |
| `github.com/gogf/gf/contrib/drivers/oracle/v2` | `contrib/drivers/oracle` |
| `github.com/gogf/gf/contrib/drivers/pgsql/v2` | `contrib/drivers/pgsql` |
| `github.com/gogf/gf/contrib/drivers/sqlite/v2` | `contrib/drivers/sqlite` |
| `github.com/gogf/gf/contrib/drivers/sqlitecgo/v2` | `contrib/drivers/sqlitecgo` |
| `github.com/gogf/gf/contrib/drivers/tidb/v2` | `contrib/drivers/tidb` |
| `github.com/gogf/gf/contrib/metric/otelmetric/v2` | `contrib/metric/otelmetric` |
| `github.com/gogf/gf/contrib/nosql/redis/v2` | `contrib/nosql/redis` |
| `github.com/gogf/gf/contrib/registry/consul/v2` | `contrib/registry/consul` |
| `github.com/gogf/gf/contrib/registry/etcd/v2` | `contrib/registry/etcd` |
| `github.com/gogf/gf/contrib/registry/file/v2` | `contrib/registry/file` |
| `github.com/gogf/gf/contrib/registry/nacos/v2` | `contrib/registry/nacos` |
| `github.com/gogf/gf/contrib/registry/polaris/v2` | `contrib/registry/polaris` |
| `github.com/gogf/gf/contrib/registry/zookeeper/v2` | `contrib/registry/zookeeper` |
| `github.com/gogf/gf/contrib/rpc/grpcx/v2` | `contrib/rpc/grpcx` |
| `github.com/gogf/gf/contrib/sdk/httpclient/v2` | `contrib/sdk/httpclient` |
| `github.com/gogf/gf/contrib/trace/otlpgrpc/v2` | `contrib/trace/otlpgrpc` |
| `github.com/gogf/gf/contrib/trace/otlphttp/v2` | `contrib/trace/otlphttp` |
| `github.com/gogf/gf/v2` | `.` |

## 包索引

### cmd

| 包名 | 导入路径 | 说明 | Reference |
|------|---------|------|----------|
| gfcmd | `github.com/gogf/gf/cmd/gf/v2/gfcmd` | Package gfcmd provides the management of CLI commands for... | `references/cmd/gf/gfcmd.md` |

### container

| 包名 | 导入路径 | 说明 | Reference |
|------|---------|------|----------|
| garray | `github.com/gogf/gf/v2/container/garray` | Package garray provides most commonly used array containe... | `references/container/garray.md` |
| glist | `github.com/gogf/gf/v2/container/glist` | Package glist provides most commonly used doubly linked l... | `references/container/glist.md` |
| gmap | `github.com/gogf/gf/v2/container/gmap` | Package gmap provides most commonly used map container wh... | `references/container/gmap.md` |
| gpool | `github.com/gogf/gf/v2/container/gpool` | Package gpool provides object-reusable concurrent-safe pool. | `references/container/gpool.md` |
| gqueue | `github.com/gogf/gf/v2/container/gqueue` | Package gqueue provides dynamic/static concurrent-safe qu... | `references/container/gqueue.md` |
| gring | `github.com/gogf/gf/v2/container/gring` | Package gring provides a concurrent-safe/unsafe ring(circ... | `references/container/gring.md` |
| gset | `github.com/gogf/gf/v2/container/gset` | Package gset provides kinds of concurrent-safe/unsafe sets. | `references/container/gset.md` |
| gtree | `github.com/gogf/gf/v2/container/gtree` | Package gtree provides concurrent-safe/unsafe tree contai... | `references/container/gtree.md` |
| gtype | `github.com/gogf/gf/v2/container/gtype` | Package gtype provides high performance and concurrent-sa... | `references/container/gtype.md` |
| gvar | `github.com/gogf/gf/v2/container/gvar` | Package gvar provides an universal variable type, like ru... | `references/container/gvar.md` |

### contrib

| 包名 | 导入路径 | 说明 | Reference |
|------|---------|------|----------|
| apollo | `github.com/gogf/gf/contrib/config/apollo/v2` | Package apollo implements gcfg.Adapter using apollo service. | `references/contrib/config/apollo.md` |
| consul | `github.com/gogf/gf/contrib/config/consul/v2` | Package consul implements gcfg.Adapter using consul service. | `references/contrib/config/consul.md` |
| kubecm | `github.com/gogf/gf/contrib/config/kubecm/v2` | Package kubecm implements gcfg.Adapter using kubernetes c... | `references/contrib/config/kubecm.md` |
| nacos | `github.com/gogf/gf/contrib/config/nacos/v2` | Package nacos implements gcfg.Adapter using nacos service. | `references/contrib/config/nacos.md` |
| polaris | `github.com/gogf/gf/contrib/config/polaris/v2` | Package polaris implements gcfg.Adapter using polaris ser... | `references/contrib/config/polaris.md` |
| clickhouse | `github.com/gogf/gf/contrib/drivers/clickhouse/v2` | Package clickhouse implements gdb.Driver, which supports ... | `references/contrib/drivers/clickhouse.md` |
| dm | `github.com/gogf/gf/contrib/drivers/dm/v2` | Package dm implements gdb.Driver, which supports operatio... | `references/contrib/drivers/dm.md` |
| gaussdb | `github.com/gogf/gf/contrib/drivers/gaussdb/v2` | Package gaussdb implements gdb.Driver, which supports ope... | `references/contrib/drivers/gaussdb.md` |
| mariadb | `github.com/gogf/gf/contrib/drivers/mariadb/v2` | Package mariadb implements gdb.Driver, which supports ope... | `references/contrib/drivers/mariadb.md` |
| mssql | `github.com/gogf/gf/contrib/drivers/mssql/v2` | Package mssql implements gdb.Driver, which supports opera... | `references/contrib/drivers/mssql.md` |
| mysql | `github.com/gogf/gf/contrib/drivers/mysql/v2` | Package mysql implements gdb.Driver, which supports opera... | `references/contrib/drivers/mysql.md` |
| oceanbase | `github.com/gogf/gf/contrib/drivers/oceanbase/v2` | Package oceanbase implements gdb.Driver, which supports o... | `references/contrib/drivers/oceanbase.md` |
| oracle | `github.com/gogf/gf/contrib/drivers/oracle/v2` | Package oracle implements gdb.Driver, which supports oper... | `references/contrib/drivers/oracle.md` |
| pgsql | `github.com/gogf/gf/contrib/drivers/pgsql/v2` | Package pgsql implements gdb.Driver, which supports opera... | `references/contrib/drivers/pgsql.md` |
| sqlite | `github.com/gogf/gf/contrib/drivers/sqlite/v2` | Package sqlite implements gdb.Driver, which supports oper... | `references/contrib/drivers/sqlite.md` |
| sqlitecgo | `github.com/gogf/gf/contrib/drivers/sqlitecgo/v2` | Package sqlitecgo implements gdb.Driver, which supports o... | `references/contrib/drivers/sqlitecgo.md` |
| tidb | `github.com/gogf/gf/contrib/drivers/tidb/v2` | Package tidb implements gdb.Driver, which supports operat... | `references/contrib/drivers/tidb.md` |
| otelmetric | `github.com/gogf/gf/contrib/metric/otelmetric/v2` | Package otelmetric provides metric functionalities using ... | `references/contrib/metric/otelmetric.md` |
| redis | `github.com/gogf/gf/contrib/nosql/redis/v2` | Package redis provides gredis.Adapter implements using go... | `references/contrib/nosql/redis.md` |
| consul | `github.com/gogf/gf/contrib/registry/consul/v2` | Package consul implements service Registry and Discovery ... | `references/contrib/registry/consul.md` |
| etcd | `github.com/gogf/gf/contrib/registry/etcd/v2` | Package etcd implements service Registry and Discovery us... | `references/contrib/registry/etcd.md` |
| file | `github.com/gogf/gf/contrib/registry/file/v2` | Package file implements service Registry and Discovery us... | `references/contrib/registry/file.md` |
| nacos | `github.com/gogf/gf/contrib/registry/nacos/v2` | Package nacos implements service Registry and Discovery u... | `references/contrib/registry/nacos.md` |
| polaris | `github.com/gogf/gf/contrib/registry/polaris/v2` | Package polaris implements service Registry and Discovery... | `references/contrib/registry/polaris.md` |
| zookeeper | `github.com/gogf/gf/contrib/registry/zookeeper/v2` | Package zookeeper implements service Registry and Discove... | `references/contrib/registry/zookeeper.md` |
| grpcx | `github.com/gogf/gf/contrib/rpc/grpcx/v2` | Package grpcx provides grpc service functionalities. | `references/contrib/rpc/grpcx.md` |
| httpclient | `github.com/gogf/gf/contrib/sdk/httpclient/v2` | Package httpclient provides http client used for SDK. | `references/contrib/sdk/httpclient.md` |
| otlpgrpc | `github.com/gogf/gf/contrib/trace/otlpgrpc/v2` | Package otlpgrpc provides gtrace.Tracer implementation us... | `references/contrib/trace/otlpgrpc.md` |
| otlphttp | `github.com/gogf/gf/contrib/trace/otlphttp/v2` | Package otlphttp provides gtrace.Tracer implementation us... | `references/contrib/trace/otlphttp.md` |

### crypto

| 包名 | 导入路径 | 说明 | Reference |
|------|---------|------|----------|
| gaes | `github.com/gogf/gf/v2/crypto/gaes` | Package gaes provides useful API for AES encryption/decry... | `references/crypto/gaes.md` |
| gcrc32 | `github.com/gogf/gf/v2/crypto/gcrc32` | Package gcrc32 provides useful API for CRC32 encryption a... | `references/crypto/gcrc32.md` |
| gdes | `github.com/gogf/gf/v2/crypto/gdes` | Package gdes provides useful API for DES encryption/decry... | `references/crypto/gdes.md` |
| gmd5 | `github.com/gogf/gf/v2/crypto/gmd5` | Package gmd5 provides useful API for MD5 encryption algor... | `references/crypto/gmd5.md` |
| grsa | `github.com/gogf/gf/v2/crypto/grsa` | Package grsa provides useful API for RSA encryption/decry... | `references/crypto/grsa.md` |
| gsha1 | `github.com/gogf/gf/v2/crypto/gsha1` | Package gsha1 provides useful API for SHA1 encryption alg... | `references/crypto/gsha1.md` |
| gsha256 | `github.com/gogf/gf/v2/crypto/gsha256` | Package gsha256 provides useful API for SHA256 encryption... | `references/crypto/gsha256.md` |

### database

| 包名 | 导入路径 | 说明 | Reference |
|------|---------|------|----------|
| gdb | `github.com/gogf/gf/v2/database/gdb` | Package gdb provides ORM features for popular relationshi... | `references/database/gdb.md` |
| gredis | `github.com/gogf/gf/v2/database/gredis` | Package gredis provides convenient client for redis server. | `references/database/gredis.md` |

### debug

| 包名 | 导入路径 | 说明 | Reference |
|------|---------|------|----------|
| gdebug | `github.com/gogf/gf/v2/debug/gdebug` | Package gdebug contains facilities for programs to debug ... | `references/debug/gdebug.md` |

### encoding

| 包名 | 导入路径 | 说明 | Reference |
|------|---------|------|----------|
| gbase64 | `github.com/gogf/gf/v2/encoding/gbase64` | Package gbase64 provides useful API for BASE64 encoding/d... | `references/encoding/gbase64.md` |
| gbinary | `github.com/gogf/gf/v2/encoding/gbinary` | Package gbinary provides useful API for handling binary/b... | `references/encoding/gbinary.md` |
| gcharset | `github.com/gogf/gf/v2/encoding/gcharset` | Package gcharset implements character-set conversion func... | `references/encoding/gcharset.md` |
| gcompress | `github.com/gogf/gf/v2/encoding/gcompress` | Package gcompress provides kinds of compression algorithm... | `references/encoding/gcompress.md` |
| ghash | `github.com/gogf/gf/v2/encoding/ghash` | Package ghash provides some classic hash functions(uint32... | `references/encoding/ghash.md` |
| ghtml | `github.com/gogf/gf/v2/encoding/ghtml` | Package ghtml provides useful API for HTML content handling. | `references/encoding/ghtml.md` |
| gini | `github.com/gogf/gf/v2/encoding/gini` | Package gini provides accessing and converting for INI co... | `references/encoding/gini.md` |
| gjson | `github.com/gogf/gf/v2/encoding/gjson` | Package gjson provides convenient API for JSON/XML/INI/YA... | `references/encoding/gjson.md` |
| gproperties | `github.com/gogf/gf/v2/encoding/gproperties` | Package gproperties provides accessing and converting for... | `references/encoding/gproperties.md` |
| gtoml | `github.com/gogf/gf/v2/encoding/gtoml` | Package gtoml provides accessing and converting for TOML ... | `references/encoding/gtoml.md` |
| gurl | `github.com/gogf/gf/v2/encoding/gurl` | Package gurl provides useful API for URL handling. | `references/encoding/gurl.md` |
| gxml | `github.com/gogf/gf/v2/encoding/gxml` | Package gxml provides accessing and converting for XML co... | `references/encoding/gxml.md` |
| gyaml | `github.com/gogf/gf/v2/encoding/gyaml` | Package gyaml provides accessing and converting for YAML ... | `references/encoding/gyaml.md` |

### errors

| 包名 | 导入路径 | 说明 | Reference |
|------|---------|------|----------|
| gcode | `github.com/gogf/gf/v2/errors/gcode` | Package gcode provides universal error code definition an... | `references/errors/gcode.md` |
| gerror | `github.com/gogf/gf/v2/errors/gerror` | Package gerror provides rich functionalities to manipulat... | `references/errors/gerror.md` |

### frame

| 包名 | 导入路径 | 说明 | Reference |
|------|---------|------|----------|
| g | `github.com/gogf/gf/v2/frame/g` | Package g provides commonly used type/function defines an... | `references/frame/g.md` |
| gins | `github.com/gogf/gf/v2/frame/gins` | Package gins provides instances and core components manag... | `references/frame/gins.md` |

### i18n

| 包名 | 导入路径 | 说明 | Reference |
|------|---------|------|----------|
| gi18n | `github.com/gogf/gf/v2/i18n/gi18n` | Package gi18n implements internationalization and localiz... | `references/i18n/gi18n.md` |

### net

| 包名 | 导入路径 | 说明 | Reference |
|------|---------|------|----------|
| gclient | `github.com/gogf/gf/v2/net/gclient` | Package gclient provides convenient http client functiona... | `references/net/gclient.md` |
| ghttp | `github.com/gogf/gf/v2/net/ghttp` | Package ghttp provides powerful http server and simple cl... | `references/net/ghttp.md` |
| gipv4 | `github.com/gogf/gf/v2/net/gipv4` |  | `references/net/gipv4.md` |
| gipv6 | `github.com/gogf/gf/v2/net/gipv6` | Package gipv6 provides useful API for IPv6 address handling. | `references/net/gipv6.md` |
| goai | `github.com/gogf/gf/v2/net/goai` | Package goai implements and provides document generating ... | `references/net/goai.md` |
| gsel | `github.com/gogf/gf/v2/net/gsel` | Package gsel provides selector definition and implements. | `references/net/gsel.md` |
| gsvc | `github.com/gogf/gf/v2/net/gsvc` | Package gsvc provides service registry and discovery defi... | `references/net/gsvc.md` |
| gtcp | `github.com/gogf/gf/v2/net/gtcp` | Package gtcp provides TCP server and client implementations. | `references/net/gtcp.md` |
| gtrace | `github.com/gogf/gf/v2/net/gtrace` | Package gtrace provides convenience wrapping functionalit... | `references/net/gtrace.md` |
| gudp | `github.com/gogf/gf/v2/net/gudp` | Package gudp provides UDP server and client implementations. | `references/net/gudp.md` |

### os

| 包名 | 导入路径 | 说明 | Reference |
|------|---------|------|----------|
| gbuild | `github.com/gogf/gf/v2/os/gbuild` | Package gbuild manages the build-in variables from "gf bu... | `references/os/gbuild.md` |
| gcache | `github.com/gogf/gf/v2/os/gcache` | Package gcache provides kinds of cache management for pro... | `references/os/gcache.md` |
| gcfg | `github.com/gogf/gf/v2/os/gcfg` | Package gcfg provides reading, caching and managing for c... | `references/os/gcfg.md` |
| gcmd | `github.com/gogf/gf/v2/os/gcmd` | Package gcmd provides console operations, like options/ar... | `references/os/gcmd.md` |
| gcron | `github.com/gogf/gf/v2/os/gcron` | Package gcron implements a cron pattern parser and job ru... | `references/os/gcron.md` |
| gctx | `github.com/gogf/gf/v2/os/gctx` | Package gctx wraps context.Context and provides extra con... | `references/os/gctx.md` |
| genv | `github.com/gogf/gf/v2/os/genv` | Package genv provides operations for environment variable... | `references/os/genv.md` |
| gfile | `github.com/gogf/gf/v2/os/gfile` | Package gfile provides easy-to-use operations for file sy... | `references/os/gfile.md` |
| gfpool | `github.com/gogf/gf/v2/os/gfpool` | Package gfpool provides io-reusable pool for file pointer. | `references/os/gfpool.md` |
| gfsnotify | `github.com/gogf/gf/v2/os/gfsnotify` | Package gfsnotify provides a platform-independent interfa... | `references/os/gfsnotify.md` |
| glog | `github.com/gogf/gf/v2/os/glog` | Package glog implements powerful and easy-to-use leveled ... | `references/os/glog.md` |
| gmetric | `github.com/gogf/gf/v2/os/gmetric` | Package gmetric provides interface definitions and simple... | `references/os/gmetric.md` |
| gmlock | `github.com/gogf/gf/v2/os/gmlock` | Package gmlock implements a concurrent-safe memory-based ... | `references/os/gmlock.md` |
| gmutex | `github.com/gogf/gf/v2/os/gmutex` | Package gmutex inherits and extends sync.Mutex and sync.R... | `references/os/gmutex.md` |
| gproc | `github.com/gogf/gf/v2/os/gproc` | Package gproc implements management and communication for... | `references/os/gproc.md` |
| gres | `github.com/gogf/gf/v2/os/gres` | Package gres provides resource management and packing/unp... | `references/os/gres.md` |
| grpool | `github.com/gogf/gf/v2/os/grpool` | Package grpool implements a goroutine reusable pool. | `references/os/grpool.md` |
| gsession | `github.com/gogf/gf/v2/os/gsession` | Package gsession implements manager and storage features ... | `references/os/gsession.md` |
| gspath | `github.com/gogf/gf/v2/os/gspath` | Package gspath implements file index and search for folders. | `references/os/gspath.md` |
| gstructs | `github.com/gogf/gf/v2/os/gstructs` | Package gstructs provides functions for struct informatio... | `references/os/gstructs.md` |
| gtime | `github.com/gogf/gf/v2/os/gtime` | Package gtime provides functionality for measuring and di... | `references/os/gtime.md` |
| gtimer | `github.com/gogf/gf/v2/os/gtimer` | Package gtimer implements timer for interval/delayed jobs... | `references/os/gtimer.md` |
| gview | `github.com/gogf/gf/v2/os/gview` | Package gview implements a template engine based on text/... | `references/os/gview.md` |

### 根包

| 包名 | 导入路径 | 说明 | Reference |
|------|---------|------|----------|
| gf | `github.com/gogf/gf/v2` |  | `references/root.md` |

### test

| 包名 | 导入路径 | 说明 | Reference |
|------|---------|------|----------|
| gtest | `github.com/gogf/gf/v2/test/gtest` | Package gtest provides convenient test utilities for unit... | `references/test/gtest.md` |

### text

| 包名 | 导入路径 | 说明 | Reference |
|------|---------|------|----------|
| gregex | `github.com/gogf/gf/v2/text/gregex` | Package gregex provides high performance API for regular ... | `references/text/gregex.md` |
| gstr | `github.com/gogf/gf/v2/text/gstr` | Package gstr provides functions for string handling. | `references/text/gstr.md` |

### util

| 包名 | 导入路径 | 说明 | Reference |
|------|---------|------|----------|
| gconv | `github.com/gogf/gf/v2/util/gconv` | Package gconv implements powerful and convenient converti... | `references/util/gconv.md` |
| gmeta | `github.com/gogf/gf/v2/util/gmeta` | Package gmeta provides embedded meta data feature for str... | `references/util/gmeta.md` |
| gmode | `github.com/gogf/gf/v2/util/gmode` | Package gmode provides release mode management for project. | `references/util/gmode.md` |
| gpage | `github.com/gogf/gf/v2/util/gpage` | Package gpage provides useful paging functionality for we... | `references/util/gpage.md` |
| grand | `github.com/gogf/gf/v2/util/grand` | Package grand provides high performance random bytes/numb... | `references/util/grand.md` |
| gtag | `github.com/gogf/gf/v2/util/gtag` | Package gtag providing tag content storing for struct. | `references/util/gtag.md` |
| guid | `github.com/gogf/gf/v2/util/guid` | Package guid provides simple and high performance unique ... | `references/util/guid.md` |
| gutil | `github.com/gogf/gf/v2/util/gutil` | Package gutil provides utility functions. | `references/util/gutil.md` |
| gvalid | `github.com/gogf/gf/v2/util/gvalid` | Package gvalid implements powerful and useful data/form v... | `references/util/gvalid.md` |

## 使用指南

当用户询问某个包的用法时：

1. **查找对应的 Reference 文件**：根据上方索引定位对应的 `.md` 文件
2. **读取包详情**：使用 `read_file` 读取 reference 文件获取完整的类型和函数定义
3. **提供示例代码**：基于包的 API 提供准确的使用示例

## 快速参考

### 核心类型

**garray** (`github.com/gogf/gf/v2/container/garray`):
  - `Array` (struct)
  - `IntArray` (struct)
  - `StrArray` (struct)
  - `TArray` (struct)
  - `SortedArray` (struct)
  - ... (3 more)

**glist** (`github.com/gogf/gf/v2/container/glist`):
  - `List` (struct)
  - `Element` (type)
  - `TElement` (struct)
  - `TList` (struct)

**gmap** (`github.com/gogf/gf/v2/container/gmap`):
  - `Map` (type)
  - `HashMap` (type)
  - `AnyAnyMap` (struct)
  - `IntAnyMap` (struct)
  - `IntIntMap` (struct)
  - ... (10 more)

**gpool** (`github.com/gogf/gf/v2/container/gpool`):
  - `Pool` (struct)
  - `NewFunc` (type)
  - `ExpireFunc` (type)
  - `TPool` (struct)
  - `TPoolNewFunc` (type)
  - ... (1 more)

**gqueue** (`github.com/gogf/gf/v2/container/gqueue`):
  - `Queue` (struct)
  - `TQueue` (struct)

**gring** (`github.com/gogf/gf/v2/container/gring`):
  - `Ring` (struct)
  - `TRing` (struct)

**gset** (`github.com/gogf/gf/v2/container/gset`):
  - `Set` (struct)
  - `IntSet` (struct)
  - `StrSet` (struct)
  - `NilChecker` (type)
  - `TSet` (struct)

**gtree** (`github.com/gogf/gf/v2/container/gtree`):
  - `AVLTree` (struct)
  - `AVLTreeNode` (type)
  - `BTree` (struct)
  - `BTreeEntry` (type)
  - `NilChecker` (type)
  - ... (8 more)

**gtype** (`github.com/gogf/gf/v2/container/gtype`):
  - `Any` (type)
  - `Bool` (struct)
  - `Byte` (struct)
  - `Bytes` (struct)
  - `Float32` (struct)
  - ... (9 more)

**gvar** (`github.com/gogf/gf/v2/container/gvar`):
  - `Var` (struct)
  - `MapOption` (type)
  - `Vars` (type)

**apollo** (`github.com/gogf/gf/contrib/config/apollo/v2`):
  - `Config` (struct)
  - `Client` (struct)
  - `ApolloAdapterCtx` (struct)

**consul** (`github.com/gogf/gf/contrib/config/consul/v2`):
  - `Config` (struct)
  - `Client` (struct)
  - `ConsulAdapterCtx` (struct)

**kubecm** (`github.com/gogf/gf/contrib/config/kubecm/v2`):
  - `Client` (struct)
  - `Config` (struct)
  - `KubecmAdapterCtx` (struct)

**nacos** (`github.com/gogf/gf/contrib/config/nacos/v2`):
  - `Config` (struct)
  - `Client` (struct)
  - `NacosAdapterCtx` (struct)

**polaris** (`github.com/gogf/gf/contrib/config/polaris/v2`):
  - `Config` (struct)
  - `Client` (struct)
  - `PolarisAdapterCtx` (struct)

**clickhouse** (`github.com/gogf/gf/contrib/drivers/clickhouse/v2`):
  - `Driver` (struct)

**dm** (`github.com/gogf/gf/contrib/drivers/dm/v2`):
  - `Driver` (struct)

**gaussdb** (`github.com/gogf/gf/contrib/drivers/gaussdb/v2`):
  - `Driver` (struct)
  - `Result` (struct)

**mariadb** (`github.com/gogf/gf/contrib/drivers/mariadb/v2`):
  - `Driver` (struct)

**mssql** (`github.com/gogf/gf/contrib/drivers/mssql/v2`):
  - `Driver` (struct)
  - `Result` (struct)

**mysql** (`github.com/gogf/gf/contrib/drivers/mysql/v2`):
  - `Driver` (struct)

**oceanbase** (`github.com/gogf/gf/contrib/drivers/oceanbase/v2`):
  - `Driver` (struct)

**oracle** (`github.com/gogf/gf/contrib/drivers/oracle/v2`):
  - `Driver` (struct)
  - `Result` (struct)

**pgsql** (`github.com/gogf/gf/contrib/drivers/pgsql/v2`):
  - `Driver` (struct)
  - `Result` (struct)

**sqlite** (`github.com/gogf/gf/contrib/drivers/sqlite/v2`):
  - `Driver` (struct)

**sqlitecgo** (`github.com/gogf/gf/contrib/drivers/sqlitecgo/v2`):
  - `Driver` (struct)

**tidb** (`github.com/gogf/gf/contrib/drivers/tidb/v2`):
  - `Driver` (struct)

**otelmetric** (`github.com/gogf/gf/contrib/metric/otelmetric/v2`):
  - `Option` (interface)

**redis** (`github.com/gogf/gf/contrib/nosql/redis/v2`):
  - `Redis` (struct)
  - `Conn` (struct)
  - `GroupGeneric` (struct)
  - `GroupHash` (struct)
  - `GroupList` (struct)
  - ... (5 more)

**consul** (`github.com/gogf/gf/contrib/registry/consul/v2`):
  - `Registry` (struct)
  - `Option` (type)
  - `Watcher` (struct)

**etcd** (`github.com/gogf/gf/contrib/registry/etcd/v2`):
  - `Registry` (struct)
  - `Option` (struct)
  - `Service` (struct)

**file** (`github.com/gogf/gf/contrib/registry/file/v2`):
  - `Registry` (struct)
  - `Service` (struct)
  - `Watcher` (struct)

**nacos** (`github.com/gogf/gf/contrib/registry/nacos/v2`):
  - `Registry` (struct)
  - `Config` (struct)
  - `Watcher` (struct)

**polaris** (`github.com/gogf/gf/contrib/registry/polaris/v2`):
  - `Option` (type)
  - `Registry` (struct)
  - `Service` (struct)
  - `Watcher` (struct)

**zookeeper** (`github.com/gogf/gf/contrib/registry/zookeeper/v2`):
  - `Content` (struct)
  - `Option` (type)
  - `Registry` (struct)

**grpcx** (`github.com/gogf/gf/contrib/rpc/grpcx/v2`):
  - `GrpcServer` (struct)
  - `Service` (struct)
  - `GrpcServerConfig` (struct)

**httpclient** (`github.com/gogf/gf/contrib/sdk/httpclient/v2`):
  - `Client` (struct)
  - `Config` (struct)
  - `Handler` (interface)
  - `DefaultHandler` (struct)

**gdb** (`github.com/gogf/gf/v2/database/gdb`):
  - `DB` (interface)
  - `TX` (interface)
  - `StatsItem` (interface)
  - `Core` (struct)
  - `DoCommitInput` (struct)
  - ... (55 more)

**gredis** (`github.com/gogf/gf/v2/database/gredis`):
  - `AdapterFunc` (type)
  - `Adapter` (interface)
  - `AdapterGroup` (interface)
  - `RedisRawClient` (type)
  - `AdapterOperation` (interface)
  - ... (28 more)

**gbinary** (`github.com/gogf/gf/v2/encoding/gbinary`):
  - `Bit` (type)

**gjson** (`github.com/gogf/gf/v2/encoding/gjson`):
  - `ContentType` (type)
  - `Json` (struct)
  - `Options` (struct)

**gcode** (`github.com/gogf/gf/v2/errors/gcode`):
  - `Code` (interface)

**gerror** (`github.com/gogf/gf/v2/errors/gerror`):
  - `IEqual` (interface)
  - `ICode` (interface)
  - `IStack` (interface)
  - `ICause` (interface)
  - `ICurrent` (interface)
  - ... (4 more)

**g** (`github.com/gogf/gf/v2/frame/g`):
  - `Var` (type)
  - `Ctx` (type)
  - `Meta` (type)
  - `Map` (type)
  - `MapAnyAny` (type)
  - ... (32 more)

**gi18n** (`github.com/gogf/gf/v2/i18n/gi18n`):
  - `Manager` (struct)
  - `Options` (struct)

**gclient** (`github.com/gogf/gf/v2/net/gclient`):
  - `Client` (struct)
  - `HandlerFunc` (type)
  - `Response` (struct)
  - `WebSocketClient` (struct)

**ghttp** (`github.com/gogf/gf/v2/net/ghttp`):
  - `Server` (struct)
  - `Router` (struct)
  - `RouterItem` (struct)
  - `HandlerFunc` (type)
  - `HandlerItem` (struct)
  - ... (18 more)

**goai** (`github.com/gogf/gf/v2/net/goai`):
  - `OpenApiV3` (struct)
  - `AddInput` (struct)
  - `Callback` (type)
  - `Callbacks` (type)
  - `CallbackRef` (struct)
  - ... (52 more)

**gsel** (`github.com/gogf/gf/v2/net/gsel`):
  - `Builder` (interface)
  - `Selector` (interface)
  - `Node` (interface)
  - `Nodes` (type)
  - `DoneFunc` (type)
  - ... (2 more)

**gsvc** (`github.com/gogf/gf/v2/net/gsvc`):
  - `Registry` (interface)
  - `Registrar` (interface)
  - `Discovery` (interface)
  - `Watcher` (interface)
  - `Service` (interface)
  - ... (7 more)

**gtcp** (`github.com/gogf/gf/v2/net/gtcp`):
  - `Conn` (struct)
  - `PkgOption` (struct)
  - `Retry` (struct)
  - `PoolConn` (struct)
  - `Server` (struct)

**gtrace** (`github.com/gogf/gf/v2/net/gtrace`):
  - `Baggage` (struct)
  - `Carrier` (type)
  - `Span` (struct)
  - `Tracer` (struct)

**gudp** (`github.com/gogf/gf/v2/net/gudp`):
  - `Retry` (struct)
  - `ClientConn` (struct)
  - `ServerConn` (struct)
  - `Server` (struct)
  - `ServerHandler` (type)

**gbuild** (`github.com/gogf/gf/v2/os/gbuild`):
  - `BuildInfo` (struct)

**gcache** (`github.com/gogf/gf/v2/os/gcache`):
  - `Func` (type)
  - `Adapter` (interface)
  - `AdapterMemory` (struct)
  - `AdapterRedis` (struct)
  - `Cache` (struct)

**gcfg** (`github.com/gogf/gf/v2/os/gcfg`):
  - `Config` (struct)
  - `Adapter` (interface)
  - `WatcherAdapter` (interface)
  - `AdapterContent` (struct)
  - `AdapterContentCtx` (struct)
  - ... (4 more)

**gcmd** (`github.com/gogf/gf/v2/os/gcmd`):
  - `Command` (struct)
  - `Function` (type)
  - `FuncWithValue` (type)
  - `Argument` (struct)
  - `ParserOption` (struct)
  - ... (1 more)

**gcron** (`github.com/gogf/gf/v2/os/gcron`):
  - `Cron` (struct)
  - `JobFunc` (type)
  - `Entry` (struct)

**gctx** (`github.com/gogf/gf/v2/os/gctx`):
  - `Ctx` (type)
  - `StrKey` (type)

**gfile** (`github.com/gogf/gf/v2/os/gfile`):
  - `CopyOption` (struct)

**gfpool** (`github.com/gogf/gf/v2/os/gfpool`):
  - `Pool` (struct)
  - `File` (struct)

**gfsnotify** (`github.com/gogf/gf/v2/os/gfsnotify`):
  - `Watcher` (struct)
  - `Callback` (struct)
  - `Event` (struct)
  - `WatchOption` (struct)
  - `Op` (type)

**glog** (`github.com/gogf/gf/v2/os/glog`):
  - `ILogger` (interface)
  - `Logger` (struct)
  - `Config` (struct)
  - `Handler` (type)
  - `HandlerInput` (struct)
  - ... (1 more)

**gmetric** (`github.com/gogf/gf/v2/os/gmetric`):
  - `MetricType` (type)
  - `Provider` (interface)
  - `MeterPerformer` (interface)
  - `MetricOption` (struct)
  - `Metric` (interface)
  - ... (32 more)

**gmlock** (`github.com/gogf/gf/v2/os/gmlock`):
  - `Locker` (struct)

**gmutex** (`github.com/gogf/gf/v2/os/gmutex`):
  - `Mutex` (struct)
  - `RWMutex` (struct)

**gproc** (`github.com/gogf/gf/v2/os/gproc`):
  - `MsgRequest` (struct)
  - `MsgResponse` (struct)
  - `Manager` (struct)
  - `Process` (struct)
  - `SigHandler` (type)

**gres** (`github.com/gogf/gf/v2/os/gres`):
  - `File` (struct)
  - `Option` (struct)
  - `Resource` (struct)
  - `ExportOption` (struct)

**grpool** (`github.com/gogf/gf/v2/os/grpool`):
  - `Func` (type)
  - `RecoverFunc` (type)
  - `Pool` (struct)

**gsession** (`github.com/gogf/gf/v2/os/gsession`):
  - `Manager` (struct)
  - `Session` (struct)
  - `Storage` (interface)
  - `StorageBase` (struct)
  - `StorageFile` (struct)
  - ... (3 more)

**gspath** (`github.com/gogf/gf/v2/os/gspath`):
  - `SPath` (struct)
  - `SPathCacheItem` (struct)

**gstructs** (`github.com/gogf/gf/v2/os/gstructs`):
  - `Type` (struct)
  - `Field` (struct)
  - `FieldsInput` (struct)
  - `FieldMapInput` (struct)
  - `RecursiveOption` (type)

**gtime** (`github.com/gogf/gf/v2/os/gtime`):
  - `Time` (struct)

**gtimer** (`github.com/gogf/gf/v2/os/gtimer`):
  - `Timer` (struct)
  - `TimerOptions` (struct)
  - `Entry` (struct)
  - `JobFunc` (type)

**gview** (`github.com/gogf/gf/v2/os/gview`):
  - `View` (struct)
  - `Params` (type)
  - `FuncMap` (type)
  - `Config` (struct)
  - `Option` (type)
  - ... (1 more)

**gtest** (`github.com/gogf/gf/v2/test/gtest`):
  - `T` (struct)

**gstr** (`github.com/gogf/gf/v2/text/gstr`):
  - `CaseType` (type)

**gconv** (`github.com/gogf/gf/v2/util/gconv`):
  - `Converter` (interface)
  - `ConverterForBasic` (interface)
  - `ConverterForTime` (interface)
  - `ConverterForInt` (interface)
  - `ConverterForUint` (interface)
  - ... (15 more)

**gmeta** (`github.com/gogf/gf/v2/util/gmeta`):
  - `Meta` (struct)

**gpage** (`github.com/gogf/gf/v2/util/gpage`):
  - `Page` (struct)

**gutil** (`github.com/gogf/gf/v2/util/gutil`):
  - `Comparator` (type)
  - `DumpOption` (struct)
  - `OriginValueAndKindOutput` (type)
  - `OriginTypeAndKindOutput` (type)

**gvalid** (`github.com/gogf/gf/v2/util/gvalid`):
  - `CustomMsg` (type)
  - `Error` (interface)
  - `RuleFunc` (type)
  - `RuleFuncInput` (struct)
  - `Validator` (struct)

