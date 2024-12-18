# 定义

**循环队列（Circular Queue）**: 使用固定大小的数组实现的队列结构，通过将数组的末尾连接到开头形成循环，使得在有限空间内高效利用存储资源。循环队列通过头尾指针的移动实现插入和删除操作。

## 示例

```golang
package main

import (
 "errors"
 "fmt"
)

// 循环队列结构
type CircularQueue struct {
 data        []int
 head, tail  int
 capacity    int
}

// 初始化循环队列
func NewCircularQueue(size int) *CircularQueue {
 return &CircularQueue{
  data:     make([]int, size+1), // 需要额外的空间区分队列满和空
  head:     0,
  tail:     0,
  capacity: size + 1,
 }
}

// 入队操作
func (q *CircularQueue) Enqueue(value int) error {
 if q.IsFull() {
  return errors.New("queue is full")
 }
 q.data[q.tail] = value
 q.tail = (q.tail + 1) % q.capacity
 return nil
}

// 出队操作
func (q *CircularQueue) Dequeue() (int, error) {
 if q.IsEmpty() {
  return 0, errors.New("queue is empty")
 }
 value := q.data[q.head]
 q.head = (q.head + 1) % q.capacity
 return value, nil
}

// 检查队列是否为空
func (q *CircularQueue) IsEmpty() bool {
 return q.head == q.tail
}

// 检查队列是否已满
func (q *CircularQueue) IsFull() bool {
 return (q.tail+1)%q.capacity == q.head
}

// 打印队列
func (q *CircularQueue) PrintQueue() {
 if q.IsEmpty() {
  fmt.Println("Queue is empty")
  return
 }
 fmt.Print("Queue: ")
 for i := q.head; i != q.tail; i = (i + 1) % q.capacity {
  fmt.Printf("%d ", q.data[i])
 }
 fmt.Println()
}

func main() {
 q := NewCircularQueue(5)

 // 入队
 q.Enqueue(1)
 q.Enqueue(2)
 q.Enqueue(3)
 q.PrintQueue() // 输出: Queue: 1 2 3

 // 出队
 q.Dequeue()
 q.PrintQueue() // 输出: Queue: 2 3

 // 测试循环特性
 q.Enqueue(4)
 q.Enqueue(5)
 q.Enqueue(6)
 q.PrintQueue() // 输出: Queue: 2 3 4 5 6
}
```

## 特点

高效利用数组空间，避免了线性队列因空间限制导致的频繁内存分配。
头尾指针移动时间复杂度为 O(1)。
存储空间固定，需提前定义数组大小，可能存在扩展限制。
