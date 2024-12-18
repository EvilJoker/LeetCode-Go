# 双向链表

定义: 每个节点包含数据域、指向前一个节点的指针（prev）、指向下一个节点的指针（next）。

## 结构

```go

package main

import "fmt"

// 定义单链表节点
package main

import "fmt"

// 定义双向链表节点
type DNode struct {
    value int
    prev  *DNode
    next  *DNode
}

// 打印双向链表
func printDList(head *DNode) {
    for head != nil {
        fmt.Printf("%d <-> ", head.value)
        head = head.next
    }
    fmt.Println("nil")
}

func main() {
    // 创建双向链表 1 <-> 2 <-> 3
    node1 := &DNode{value: 1}
    node2 := &DNode{value: 2}
    node3 := &DNode{value: 3}

    node1.next = node2
    node2.prev = node1
    node2.next = node3
    node3.prev = node2

    printDList(node1)
}

```

特点:
可以在双向（前后）两种方向遍历链表，灵活性更强。
插入和删除节点更高效，因为可以直接访问前驱节点（无需从头遍历找到）。
每个节点需要额外的 prev 指针，占用更多空间。
