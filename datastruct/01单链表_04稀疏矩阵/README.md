# 定义

稀疏链表: 是一种特殊的链表结构，用于存储稀疏数据，即大部分值为空（如零或空值）且仅少量非空值需要存储的场景。稀疏链表通过节省存储空间提高效率。

## 示例

```golang
package main

import "fmt"

// 定义稀疏链表节点
type SparseNode struct {
    index int        // 数据位置
    value int        // 非空值
    next  *SparseNode // 指向下一个节点
}

// 打印稀疏链表
func printSparseList(head *SparseNode) {
    for head != nil {
        fmt.Printf("Index: %d, Value: %d -> ", head.index, head.value)
        head = head.next
    }
    fmt.Println("nil")
}

func main() {
    // 创建稀疏链表，存储 0 0 5 0 0 3 0
    head := &SparseNode{index: 2, value: 5}
    head.next = &SparseNode{index: 5, value: 3}

    printSparseList(head)
}
```

## 特点

+ 节省空间: 仅存储非空数据，避免浪费存储空间。
+ 适合稀疏数据结构: 在大部分元素为空的情况下，效率更高。
+ 灵活性: 通过链表动态扩展，无需预分配大块空间。