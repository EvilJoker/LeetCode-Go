# 定义

队列的链表实现: 使用链表来实现队列结构，每个节点存储一个元素和指向下一个节点的指针。链表实现的队列可以动态扩展，不受固定大小限制，适用于需要频繁插入和删除的场景。

## 示例

```golang
package main

import "fmt"

// 定义链表节点
type Node struct {
 value int
 next  *Node
}

// 定义队列
type Queue struct {
 front *Node // 队列头部
 rear  *Node // 队列尾部
}
入
// 队操作
func (q *Queue) Enqueue(value int) {
 newNode := &Node{value: value}
 if q.rear == nil {
  // 队列为空时
  q.front = newNode
  q.rear = newNode
 } else {
  q.rear.next = newNode
  q.rear = newNode
 }
}

// 出队操作
func (q *Queue) Dequeue() (int, error) {
 if q.front == nil {
  return 0, fmt.Errorf("queue is empty")
 }
 value := q.front.value
 q.front = q.front.next
 if q.front == nil {
  q.rear = nil // 队列为空时更新尾部指针
 }
 return value, nil
}

// 打印队列
func (q *Queue) Print() {
 current := q.front
 for current != nil {
  fmt.Printf("%d ", current.value)
  current = current.next
 }
 fmt.Println()
}

func main() {
 queue := &Queue{}

 // 入队
 queue.Enqueue(10)
 queue.Enqueue(20)
 queue.Enqueue(30)
 fmt.Print("Queue after enqueues: ")
 queue.Print()

 // 出队
 value, _ := queue.Dequeue()
 fmt.Printf("Dequeued: %d\n", value)
 fmt.Print("Queue after dequeue: ")
 queue.Print()
}
```

## 特点

动态大小：链表实现的队列可以根据需要动态扩展，不受固定大小限制。
操作高效：插入和删除操作在 O(1) 时间内完成。
内存利用率高：只分配实际需要的内存，但需要额外的指针存储空间。
适用于需要频繁的入队和出队操作的场景。
