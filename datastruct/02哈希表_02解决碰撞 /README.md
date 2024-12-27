# 定义

哈希表中解决碰撞
在哈希表中，碰撞是指两个不同的键通过哈希函数映射到相同的索引位置。解决碰撞的方法主要有以下几种：

## 1. 开放地址法

通过探测下一个空闲位置存储发生碰撞的键值对。

示例

```go
package main

import (
    "fmt"
)

const TableSize = 7

type HashTable struct {
    table [TableSize]int
}

// 哈希函数
func hash(key int) int {
    return key % TableSize
}

// 插入元素，使用线性探测
func (h *HashTable) Insert(key int) {
    index := hash(key)
    for h.table[index] != 0 { // 线性探测
        index = (index + 1) % TableSize
    }
    h.table[index] = key
}

// 打印哈希表
func (h *HashTable) Print() {
    fmt.Println(h.table)
}

func main() {
    h := &HashTable{}
    h.Insert(10)
    h.Insert(17) // 发生碰撞
    h.Insert(24) // 再次碰撞
    h.Print()
}

```

## 2. 链地址法（分离链表法）
为每个哈希表槽位维护一个链表，用来存储所有哈希到该位置的键值对。

```golang
package main

import "fmt"

// 链表节点
type Node struct {
    key  int
    next *Node
}

type HashTable struct {
    table []*Node
}

const TableSize = 7

// 哈希函数
func hash(key int) int {
    return key % TableSize
}

// 插入元素，使用链地址法
func (h *HashTable) Insert(key int) {
    index := hash(key)
    newNode := &Node{key: key}
    if h.table[index] == nil {
        h.table[index] = newNode
    } else {
        current := h.table[index]
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
}

// 打印哈希表
func (h *HashTable) Print() {
    for i, node := range h.table {
        fmt.Printf("%d: ", i)
        for node != nil {
            fmt.Printf("%d -> ", node.key)
            node = node.next
        }
        fmt.Println("nil")
    }
}

func main() {
    h := &HashTable{table: make([]*Node, TableSize)}
    h.Insert(10)
    h.Insert(17) // 发生碰撞
    h.Insert(24) // 再次碰撞
    h.Print()
}


```

## 特点

特点
开放地址法
优点: 简单，无需额外内存。
缺点: 当装载因子（存储元素数量/表大小）过高时，性能会显著下降，插入与查找变慢。
链地址法
优点: 哈希表可以处理高装载因子，性能更稳定。
缺点: 需要额外的内存管理，链表操作相对复杂。
