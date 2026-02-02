# gtree

> Package: `github.com/gogf/gf/v2/container/gtree`

```go
import "github.com/gogf/gf/v2/container/gtree"
```

## 概述

Package gtree provides concurrent-safe/unsafe tree containers.

## 源文件

- `gtree.go`
- `gtree_avltree.go`
- `gtree_btree.go`
- `gtree_k_v_avltree.go`
- `gtree_k_v_btree.go`
- `gtree_k_v_redblacktree.go`
- `gtree_redblacktree.go`

## 类型定义

### AVLTree

**类型**: struct

AVLTree holds elements of the AVL tree.


**方法**:

- `func (tree *AVLTree) Clone() *AVLTree`
  Clone clones and returns a new tree from current tree.
- `func (tree *AVLTree) Set(key any, value any)`
  Set sets key-value pair into the tree.
- `func (tree *AVLTree) Sets(data map[any]any)`
  Sets batch sets key-values to the tree.
- `func (tree *AVLTree) SetIfNotExist(key any, value any) bool`
  SetIfNotExist sets `value` to the map if the `key` does not exist, and then r...
- `func (tree *AVLTree) SetIfNotExistFunc(key any, f func) bool`
  SetIfNotExistFunc sets value with return value of callback function `f`, and ...
- `func (tree *AVLTree) SetIfNotExistFuncLock(key any, f func) bool`
  SetIfNotExistFuncLock sets value with return value of callback function `f`, ...
- `func (tree *AVLTree) Get(key any) value any`
  Get searches the `key` in the tree and returns its associated `value` or nil ...
- `func (tree *AVLTree) GetOrSet(key any, value any) any`
  GetOrSet returns its `value` of `key`, or sets value with given `value` if it...
- `func (tree *AVLTree) GetOrSetFunc(key any, f func) any`
  GetOrSetFunc returns its `value` of `key`, or sets value with returned value ...
- `func (tree *AVLTree) GetOrSetFuncLock(key any, f func) any`
  GetOrSetFuncLock returns its `value` of `key`, or sets value with returned va...
- `func (tree *AVLTree) GetVar(key any) *gvar.Var`
  GetVar returns a gvar.Var with the value by given `key`.
- `func (tree *AVLTree) GetVarOrSet(key any, value any) *gvar.Var`
  GetVarOrSet returns a gvar.Var with result from GetVarOrSet.
- `func (tree *AVLTree) GetVarOrSetFunc(key any, f func) *gvar.Var`
  GetVarOrSetFunc returns a gvar.Var with result from GetOrSetFunc.
- `func (tree *AVLTree) GetVarOrSetFuncLock(key any, f func) *gvar.Var`
  GetVarOrSetFuncLock returns a gvar.Var with result from GetOrSetFuncLock.
- `func (tree *AVLTree) Search(key any) (value any, found bool)`
  Search searches the tree with given `key`.
- `func (tree *AVLTree) Contains(key any) bool`
  Contains checks and returns whether given `key` exists in the tree.
- `func (tree *AVLTree) Size() int`
  Size returns number of nodes in the tree.
- `func (tree *AVLTree) IsEmpty() bool`
  IsEmpty returns true if the tree does not contain any nodes.
- `func (tree *AVLTree) Remove(key any) value any`
  Remove removes the node from the tree by `key`, and returns its associated va...
- `func (tree *AVLTree) Removes(keys []any)`
  Removes batch deletes key-value pairs from the tree by `keys`.
- `func (tree *AVLTree) Clear()`
  Clear removes all nodes from the tree.
- `func (tree *AVLTree) Keys() []any`
  Keys returns all keys from the tree in order by its comparator.
- `func (tree *AVLTree) Values() []any`
  Values returns all values from the true in order by its comparator based on t...
- `func (tree *AVLTree) Replace(data map[any]any)`
  Replace clears the data of the tree and sets the nodes by given `data`.
- `func (tree *AVLTree) Print()`
  Print prints the tree to stdout.
- `func (tree *AVLTree) String() string`
  String returns a string representation of container.
- `func (tree *AVLTree) MarshalJSON() (jsonBytes []byte, err error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (tree *AVLTree) Map() map[any]any`
  Map returns all key-value pairs as map.
- `func (tree *AVLTree) MapStrAny() map[string]any`
  MapStrAny returns all key-value items as map[string]any.
- `func (tree *AVLTree) Iterator(f func)`
  Iterator is alias of IteratorAsc.
- `func (tree *AVLTree) IteratorFrom(key any, match bool, f func)`
  IteratorFrom is alias of IteratorAscFrom.
- `func (tree *AVLTree) IteratorAsc(f func)`
  IteratorAsc iterates the tree readonly in ascending order with given callback...
- `func (tree *AVLTree) IteratorAscFrom(key any, match bool, f func)`
  IteratorAscFrom iterates the tree readonly in ascending order with given call...
- `func (tree *AVLTree) IteratorDesc(f func)`
  IteratorDesc iterates the tree readonly in descending order with given callba...
- `func (tree *AVLTree) IteratorDescFrom(key any, match bool, f func)`
  IteratorDescFrom iterates the tree readonly in descending order with given ca...
- `func (tree *AVLTree) Left() *AVLTreeNode`
  Left returns the minimum element corresponding to the comparator of the tree ...
- `func (tree *AVLTree) Right() *AVLTreeNode`
  Right returns the maximum element corresponding to the comparator of the tree...
- `func (tree *AVLTree) Floor(key any) (floor *AVLTreeNode, found bool)`
  Floor Finds floor node of the input key, returns the floor node or nil if no ...
- `func (tree *AVLTree) Ceiling(key any) (ceiling *AVLTreeNode, found bool)`
  Ceiling finds ceiling node of the input key, returns the ceiling node or nil ...
- `func (tree *AVLTree) Flip(comparator ...func)`
  Flip exchanges key-value of the tree to value-key.

### AVLTreeNode

**类型**: type

AVLTreeNode is a single element within the tree.


### BTree

**类型**: struct

BTree holds elements of the B-tree.


**方法**:

- `func (tree *BTree) Clone() *BTree`
  Clone clones and returns a new tree from current tree.
- `func (tree *BTree) Set(key any, value any)`
  Set sets key-value pair into the tree.
- `func (tree *BTree) Sets(data map[any]any)`
  Sets batch sets key-values to the tree.
- `func (tree *BTree) SetIfNotExist(key any, value any) bool`
  SetIfNotExist sets `value` to the map if the `key` does not exist, and then r...
- `func (tree *BTree) SetIfNotExistFunc(key any, f func) bool`
  SetIfNotExistFunc sets value with return value of callback function `f`, and ...
- `func (tree *BTree) SetIfNotExistFuncLock(key any, f func) bool`
  SetIfNotExistFuncLock sets value with return value of callback function `f`, ...
- `func (tree *BTree) Get(key any) value any`
  Get searches the `key` in the tree and returns its associated `value` or nil ...
- `func (tree *BTree) GetOrSet(key any, value any) any`
  GetOrSet returns its `value` of `key`, or sets value with given `value` if it...
- `func (tree *BTree) GetOrSetFunc(key any, f func) any`
  GetOrSetFunc returns its `value` of `key`, or sets value with returned value ...
- `func (tree *BTree) GetOrSetFuncLock(key any, f func) any`
  GetOrSetFuncLock returns its `value` of `key`, or sets value with returned va...
- `func (tree *BTree) GetVar(key any) *gvar.Var`
  GetVar returns a gvar.Var with the value by given `key`.
- `func (tree *BTree) GetVarOrSet(key any, value any) *gvar.Var`
  GetVarOrSet returns a gvar.Var with result from GetVarOrSet.
- `func (tree *BTree) GetVarOrSetFunc(key any, f func) *gvar.Var`
  GetVarOrSetFunc returns a gvar.Var with result from GetOrSetFunc.
- `func (tree *BTree) GetVarOrSetFuncLock(key any, f func) *gvar.Var`
  GetVarOrSetFuncLock returns a gvar.Var with result from GetOrSetFuncLock.
- `func (tree *BTree) Search(key any) (value any, found bool)`
  Search searches the tree with given `key`.
- `func (tree *BTree) Contains(key any) bool`
  Contains checks and returns whether given `key` exists in the tree.
- `func (tree *BTree) Size() int`
  Size returns number of nodes in the tree.
- `func (tree *BTree) IsEmpty() bool`
  IsEmpty returns true if tree does not contain any nodes
- `func (tree *BTree) Remove(key any) value any`
  Remove removes the node from the tree by `key`, and returns its associated va...
- `func (tree *BTree) Removes(keys []any)`
  Removes batch deletes key-value pairs from the tree by `keys`.
- `func (tree *BTree) Clear()`
  Clear removes all nodes from the tree.
- `func (tree *BTree) Keys() []any`
  Keys returns all keys from the tree in order by its comparator.
- `func (tree *BTree) Values() []any`
  Values returns all values from the true in order by its comparator based on t...
- `func (tree *BTree) Replace(data map[any]any)`
  Replace clears the data of the tree and sets the nodes by given `data`.
- `func (tree *BTree) Map() map[any]any`
  Map returns all key-value pairs as map.
- `func (tree *BTree) MapStrAny() map[string]any`
  MapStrAny returns all key-value items as map[string]any.
- `func (tree *BTree) Print()`
  Print prints the tree to stdout.
- `func (tree *BTree) String() string`
  String returns a string representation of container (for debugging purposes)
- `func (tree *BTree) MarshalJSON() (jsonBytes []byte, err error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (tree *BTree) Iterator(f func)`
  Iterator is alias of IteratorAsc.
- `func (tree *BTree) IteratorFrom(key any, match bool, f func)`
  IteratorFrom is alias of IteratorAscFrom.
- `func (tree *BTree) IteratorAsc(f func)`
  IteratorAsc iterates the tree readonly in ascending order with given callback...
- `func (tree *BTree) IteratorAscFrom(key any, match bool, f func)`
  IteratorAscFrom iterates the tree readonly in ascending order with given call...
- `func (tree *BTree) IteratorDesc(f func)`
  IteratorDesc iterates the tree readonly in descending order with given callba...
- `func (tree *BTree) IteratorDescFrom(key any, match bool, f func)`
  IteratorDescFrom iterates the tree readonly in descending order with given ca...
- `func (tree *BTree) Height() int`
  Height returns the height of the tree.
- `func (tree *BTree) Left() *BTreeEntry`
  Left returns the minimum element corresponding to the comparator of the tree ...
- `func (tree *BTree) Right() *BTreeEntry`
  Right returns the maximum element corresponding to the comparator of the tree...

### BTreeEntry

**类型**: type

BTreeEntry represents the key-value pair contained within nodes.


### NilChecker

**类型**: type

NilChecker is a function that checks whether the given value is nil.


### AVLKVTree

**类型**: struct

AVLKVTree holds elements of the AVL tree.


### AVLKVTreeNode

**类型**: struct

AVLKVTreeNode is a single element within the tree.


### BKVTree

**类型**: struct

BKVTree holds elements of the B-tree.


### BKVTreeEntry

**类型**: struct

BKVTreeEntry represents the key-value pair contained within nodes.


### RedBlackKVTree

**类型**: struct

RedBlackKVTree holds elements of the red-black tree.


### RedBlackKVTreeNode

**类型**: struct

RedBlackKVTreeNode is a single element within the tree.


### RedBlackTree

**类型**: struct

RedBlackTree holds elements of the red-black tree.


**方法**:

- `func (tree *RedBlackTree) SetComparator(comparator func)`
  SetComparator sets/changes the comparator for sorting.
- `func (tree *RedBlackTree) Clone() *RedBlackTree`
  Clone clones and returns a new tree from current tree.
- `func (tree *RedBlackTree) Set(key any, value any)`
  Set sets key-value pair into the tree.
- `func (tree *RedBlackTree) Sets(data map[any]any)`
  Sets batch sets key-values to the tree.
- `func (tree *RedBlackTree) SetIfNotExist(key any, value any) bool`
  SetIfNotExist sets `value` to the map if the `key` does not exist, and then r...
- `func (tree *RedBlackTree) SetIfNotExistFunc(key any, f func) bool`
  SetIfNotExistFunc sets value with return value of callback function `f`, and ...
- `func (tree *RedBlackTree) SetIfNotExistFuncLock(key any, f func) bool`
  SetIfNotExistFuncLock sets value with return value of callback function `f`, ...
- `func (tree *RedBlackTree) Get(key any) value any`
  Get searches the `key` in the tree and returns its associated `value` or nil ...
- `func (tree *RedBlackTree) GetOrSet(key any, value any) any`
  GetOrSet returns its `value` of `key`, or sets value with given `value` if it...
- `func (tree *RedBlackTree) GetOrSetFunc(key any, f func) any`
  GetOrSetFunc returns its `value` of `key`, or sets value with returned value ...
- `func (tree *RedBlackTree) GetOrSetFuncLock(key any, f func) any`
  GetOrSetFuncLock returns its `value` of `key`, or sets value with returned va...
- `func (tree *RedBlackTree) GetVar(key any) *gvar.Var`
  GetVar returns a gvar.Var with the value by given `key`.
- `func (tree *RedBlackTree) GetVarOrSet(key any, value any) *gvar.Var`
  GetVarOrSet returns a gvar.Var with result from GetVarOrSet.
- `func (tree *RedBlackTree) GetVarOrSetFunc(key any, f func) *gvar.Var`
  GetVarOrSetFunc returns a gvar.Var with result from GetOrSetFunc.
- `func (tree *RedBlackTree) GetVarOrSetFuncLock(key any, f func) *gvar.Var`
  GetVarOrSetFuncLock returns a gvar.Var with result from GetOrSetFuncLock.
- `func (tree *RedBlackTree) Search(key any) (value any, found bool)`
  Search searches the tree with given `key`.
- `func (tree *RedBlackTree) Contains(key any) bool`
  Contains checks and returns whether given `key` exists in the tree.
- `func (tree *RedBlackTree) Size() int`
  Size returns number of nodes in the tree.
- `func (tree *RedBlackTree) IsEmpty() bool`
  IsEmpty returns true if tree does not contain any nodes.
- `func (tree *RedBlackTree) Remove(key any) value any`
  Remove removes the node from the tree by `key`, and returns its associated va...
- `func (tree *RedBlackTree) Removes(keys []any)`
  Removes batch deletes key-value pairs from the tree by `keys`.
- `func (tree *RedBlackTree) Clear()`
  Clear removes all nodes from the tree.
- `func (tree *RedBlackTree) Keys() []any`
  Keys returns all keys from the tree in order by its comparator.
- `func (tree *RedBlackTree) Values() []any`
  Values returns all values from the true in order by its comparator based on t...
- `func (tree *RedBlackTree) Replace(data map[any]any)`
  Replace clears the data of the tree and sets the nodes by given `data`.
- `func (tree *RedBlackTree) Print()`
  Print prints the tree to stdout.
- `func (tree *RedBlackTree) String() string`
  String returns a string representation of container
- `func (tree RedBlackTree) MarshalJSON() (jsonBytes []byte, err error)`
  MarshalJSON implements the interface MarshalJSON for json.Marshal.
- `func (tree *RedBlackTree) Map() map[any]any`
  Map returns all key-value pairs as map.
- `func (tree *RedBlackTree) MapStrAny() map[string]any`
  MapStrAny returns all key-value items as map[string]any.
- `func (tree *RedBlackTree) Iterator(f func)`
  Iterator is alias of IteratorAsc.
- `func (tree *RedBlackTree) IteratorFrom(key any, match bool, f func)`
  IteratorFrom is alias of IteratorAscFrom.
- `func (tree *RedBlackTree) IteratorAsc(f func)`
  IteratorAsc iterates the tree readonly in ascending order with given callback...
- `func (tree *RedBlackTree) IteratorAscFrom(key any, match bool, f func)`
  IteratorAscFrom iterates the tree readonly in ascending order with given call...
- `func (tree *RedBlackTree) IteratorDesc(f func)`
  IteratorDesc iterates the tree readonly in descending order with given callba...
- `func (tree *RedBlackTree) IteratorDescFrom(key any, match bool, f func)`
  IteratorDescFrom iterates the tree readonly in descending order with given ca...
- `func (tree *RedBlackTree) Left() *RedBlackTreeNode`
  Left returns the minimum element corresponding to the comparator of the tree ...
- `func (tree *RedBlackTree) Right() *RedBlackTreeNode`
  Right returns the maximum element corresponding to the comparator of the tree...
- `func (tree *RedBlackTree) Floor(key any) (floor *RedBlackTreeNode, found bool)`
  Floor Finds floor node of the input key, returns the floor node or nil if no ...
- `func (tree *RedBlackTree) Ceiling(key any) (ceiling *RedBlackTreeNode, found bool)`
  Ceiling finds ceiling node of the input key, returns the ceiling node or nil ...
- `func (tree *RedBlackTree) Flip(comparator ...func)`
  Flip exchanges key-value of the tree to value-key.
- `func (tree *RedBlackTree) UnmarshalJSON(b []byte) error`
  UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
- `func (tree *RedBlackTree) UnmarshalValue(value any) err error`
  UnmarshalValue is an interface implement which sets any type of value for map.

### RedBlackTreeNode

**类型**: type

RedBlackTreeNode is a single element within the tree.


## 函数

### NewAVLTree

```go
func NewAVLTree(comparator func, safe ...bool) *AVLTree
```

NewAVLTree instantiates an AVL tree with the custom key comparator.

The parameter `safe` is used to specify whether using tree in concurrent-safety,
which is false in default.

### NewAVLTreeFrom

```go
func NewAVLTreeFrom(comparator func, data map[any]any, safe ...bool) *AVLTree
```

NewAVLTreeFrom instantiates an AVL tree with the custom key comparator and data map.

The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.

### NewBTree

```go
func NewBTree(m int, comparator func, safe ...bool) *BTree
```

NewBTree instantiates a B-tree with `m` (maximum number of children) and a custom key comparator.
The parameter `safe` is used to specify whether using tree in concurrent-safety,
which is false in default.
Note that the `m` must be greater or equal than 3, or else it panics.

### NewBTreeFrom

```go
func NewBTreeFrom(m int, comparator func, data map[any]any, safe ...bool) *BTree
```

NewBTreeFrom instantiates a B-tree with `m` (maximum number of children), a custom key comparator and data map.
The parameter `safe` is used to specify whether using tree in concurrent-safety,
which is false in default.

### NewAVLKVTree

```go
func NewAVLKVTree(comparator func, safe ...bool) *
```

NewAVLKVTree instantiates an AVL tree with the custom key comparator.

The parameter `safe` is used to specify whether using tree in concurrent-safety,
which is false in default.

### NewAVLKVTreeWithChecker

```go
func NewAVLKVTreeWithChecker(comparator func, checker , safe ...bool) *
```

NewAVLKVTreeWithChecker instantiates an AVL tree with the custom key comparator and nil checker.
The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.
The parameter `checker` is used to specify whether the given value is nil.

### NewAVLKVTreeFrom

```go
func NewAVLKVTreeFrom(comparator func, data map[K]V, safe ...bool) *
```

NewAVLKVTreeFrom instantiates an AVL tree with the custom key comparator and data map.

The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.

### NewAVLKVTreeWithCheckerFrom

```go
func NewAVLKVTreeWithCheckerFrom(comparator func, data map[K]V, checker , safe ...bool) *
```

NewAVLKVTreeWithCheckerFrom instantiates an AVL tree with the custom key comparator, nil checker and data map.
The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.
The parameter `checker` is used to specify whether the given value is nil.

### NewBKVTree

```go
func NewBKVTree(m int, comparator func, safe ...bool) *
```

NewBKVTree instantiates a B-tree with `m` (maximum number of children) and a custom key comparator.
The parameter `safe` is used to specify whether using tree in concurrent-safety,
which is false in default.
Note that the `m` must be greater or equal than 3, or else it panics.

### NewBKVTreeWithChecker

```go
func NewBKVTreeWithChecker(m int, comparator func, checker , safe ...bool) *
```

NewBKVTreeWithChecker instantiates a B-tree with `m` (maximum number of children), a custom key comparator and nil checker.
The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.
The parameter `checker` is used to specify whether the given value is nil.

### NewBKVTreeFrom

```go
func NewBKVTreeFrom(m int, comparator func, data map[K]V, safe ...bool) *
```

NewBKVTreeFrom instantiates a B-tree with `m` (maximum number of children), a custom key comparator and data map.
The parameter `safe` is used to specify whether using tree in concurrent-safety,
which is false in default.

### NewBKVTreeWithCheckerFrom

```go
func NewBKVTreeWithCheckerFrom(m int, comparator func, data map[K]V, checker , safe ...bool) *
```

NewBKVTreeWithCheckerFrom instantiates a B-tree with `m` (maximum number of children), a custom key comparator, nil checker and data map.
The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.
The parameter `checker` is used to specify whether the given value is nil.

### NewRedBlackKVTree

```go
func NewRedBlackKVTree(comparator func, safe ...bool) *
```

NewRedBlackKVTree instantiates a red-black tree with the custom key comparator.
The parameter `safe` is used to specify whether using tree in concurrent-safety,
which is false in default.

### NewRedBlackKVTreeWithChecker

```go
func NewRedBlackKVTreeWithChecker(comparator func, checker , safe ...bool) *
```

NewRedBlackKVTreeWithChecker instantiates a red-black tree with the custom key comparator and `nilChecker`.
The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.
The parameter `checker` is used to specify whether the given value is nil.

### NewRedBlackKVTreeFrom

```go
func NewRedBlackKVTreeFrom(comparator func, data map[K]V, safe ...bool) *
```

NewRedBlackKVTreeFrom instantiates a red-black tree with the custom key comparator and `data` map.
The parameter `safe` is used to specify whether using tree in concurrent-safety,
which is false in default.

### NewRedBlackKVTreeWithCheckerFrom

```go
func NewRedBlackKVTreeWithCheckerFrom(comparator func, data map[K]V, checker , safe ...bool) *
```

NewRedBlackKVTreeWithCheckerFrom instantiates a red-black tree with the custom key comparator, `data` map and `nilChecker`.
The parameter `safe` is used to specify whether using tree in concurrent-safety, which is false in default.
The parameter `checker` is used to specify whether the given value is nil.

### RedBlackKVTreeInit

```go
func RedBlackKVTreeInit(tree *, comparator func, safe ...bool)
```

RedBlackKVTreeInit instantiates a red-black tree with the custom key comparator.
The parameter `safe` is used to specify whether using tree in concurrent-safety,
which is false in default.

### RedBlackKVTreeInitFrom

```go
func RedBlackKVTreeInitFrom(tree *, comparator func, data map[K]V, safe ...bool)
```

RedBlackKVTreeInitFrom instantiates a red-black tree with the custom key comparator and `data` map.
The parameter `safe` is used to specify whether using tree in concurrent-safety,
which is false in default.

### MarshalJSON

```go
func (tree ) MarshalJSON() (jsonBytes []byte, err error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

### NewRedBlackTree

```go
func NewRedBlackTree(comparator func, safe ...bool) *RedBlackTree
```

NewRedBlackTree instantiates a red-black tree with the custom key comparator.
The parameter `safe` is used to specify whether using tree in concurrent-safety,
which is false in default.

### NewRedBlackTreeFrom

```go
func NewRedBlackTreeFrom(comparator func, data map[any]any, safe ...bool) *RedBlackTree
```

NewRedBlackTreeFrom instantiates a red-black tree with the custom key comparator and `data` map.
The parameter `safe` is used to specify whether using tree in concurrent-safety,
which is false in default.

## 内部依赖

- `container/gvar`
- `internal/json`
- `internal/rwmutex`
- `text/gstr`
- `util/gconv`
- `util/gutil`

## 外部依赖

- `github.com/emirpasic/gods/v2/trees/avltree`
- `github.com/emirpasic/gods/v2/trees/btree`
- `github.com/emirpasic/gods/v2/trees/redblacktree`

