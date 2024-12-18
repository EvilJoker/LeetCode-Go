# 定义

双端队列: 一种特殊的队列数据结构，支持在队列的两端进行插入和删除操作。它结合了栈和队列的特点，可以高效地在头部和尾部操作数据。

## 示例

```golang
package main

import "fmt"

// 双端队列结构
type Deque struct {
 data []int
}

// 从队列头部插入元素
func (d *Deque) PushFront(value int) {
 d.data = append([]int{value}, d.data...)
}

// 从队列尾部插入元素
func (d *Deque) PushBack(value int) {
 d.data = append(d.data, value)
}

// 从队列头部删除元素
func (d *Deque) PopFront() int {
 if len(d.data) == 0 {
  panic("Deque is empty")
 }
 front := d.data[0]
 d.data = d.data[1:]
 return front
}

// 从队列尾部删除元素
func (d *Deque) PopBack() int {
 if len(d.data) == 0 {
  panic("Deque is empty")
 }
 back := d.data[len(d.data)-1]
 d.data = d.data[:len(d.data)-1]
 return back
}

// 打印双端队列
func (d *Deque) Print() {
 fmt.Println(d.data)
}

func main() {
 // 初始化双端队列
 d := &Deque{}

 // 插入元素
 d.PushFront(10)
 d.PushBack(20)
 d.PushFront(5)

 // 打印队列状态
 d.Print() // 输出: [5 10 20]

 // 删除元素
 fmt.Println("PopFront:", d.PopFront()) // 输出: 5
 fmt.Println("PopBack:", d.PopBack())   // 输出: 20

 // 打印队列状态
 d.Print() // 输出: [10]
}

```

## 特点

支持双向操作，既可以作为栈也可以作为队列使用。
插入和删除操作时间复杂度为 O(1)（取决于底层实现）。
适用于需要频繁在两端操作的场景，例如滑动窗口问题或任务调度。
