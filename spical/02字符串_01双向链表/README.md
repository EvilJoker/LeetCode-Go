# 定义

**静态链表**:  
静态链表是一种特殊的链表实现方式，使用数组代替动态内存分配来存储节点。每个数组元素包含数据和指向下一个元素索引的“指针”，从而模拟链表结构。

## 示例

```golang
package main

import "fmt"

// 定义静态链表节点
type Node struct {
    value int
    next  int // 指向下一个节点的索引
}

func main() {
    // 定义一个数组模拟静态链表
    const maxSize = 5
    list := [maxSize]Node{
        {value: 1, next: 1},
        {value: 2, next: 2},
        {value: 3, next: -1}, // -1 表示链表结束
    }

    // 遍历静态链表
    index := 0
    for index != -1 {
        fmt.Printf("%d -> ", list[index].value)
        index = list[index].next
    }
    fmt.Println("nil")
}
```

## 特点

使用数组存储: 节点存储在数组中，避免动态内存分配。
通过索引模拟指针: 数组元素的 next 字段存储下一个节点的索引，而不是实际地址。
固定大小: 数组大小固定，链表容量受限。
效率更高: 数组比动态分配内存的链表更高效，适用于内存管理较严格的场景。
插入删除较复杂: 需要手动调整 next 指针和管理空闲节点。
