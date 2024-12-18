# 定义

堆（Heap）是一种特殊的完全二叉树，具有以下性质：在堆中，每个父节点的值都满足特定的大小关系与其子节点（对于最大堆，父节点的值大于等于子节点的值；对于最小堆，父节点的值小于等于子节点的值）。堆通常通过数组来实现，因为完全二叉树的结构特性可以通过数组索引来方便地表示。

- **最大堆**（Max Heap）：在最大堆中，父节点的值大于或等于其子节点的值，因此根节点的值是最大的。
- **最小堆**（Min Heap）：在最小堆中，父节点的值小于或等于其子节点的值，因此根节点的值是最小的。

堆通常用于优先队列、堆排序等应用。

---

## 示例

**最大堆示例：**

```golang
package main

import (
 "fmt"
)

// 堆结构体
type MaxHeap struct {
 arr []int
}

// 构造最大堆
// 上浮操作， 从当前节点开始，比较当前节点与父节点的大小关系，如果当前节点不符合堆的性质，则交换位置并继续向上调整。
func (h *MaxHeap) heapifyUp(i int) {
// i>0 保证不是 根节点
// (i-1)/2 是父节点

 for i > 0 && h.arr[i] > h.arr[(i-1)/2] {
 // 当子节点大于父节点时，交换值。
  h.arr[i], h.arr[(i-1)/2] = h.arr[(i-1)/2], h.arr[i]
  // 更新为 父节点编号，继续上浮
  i = (i - 1) / 2
 }
}

// 插入元素
//将新元素添加到堆的末尾，增加堆的大小。
//使用 heapifyUp（上浮操作）调整堆，确保堆的性质得以保持。如果新元素大于其父节点（最大堆）或小于其父节点（最小堆），则交换它与父节点的位置，直到堆的性质得到恢复。
func (h *MaxHeap) insert(val int) {
 h.arr = append(h.arr, val)
 h.heapifyUp(len(h.arr) - 1)
}

// 获取堆顶元素
func (h *MaxHeap) peek() int {
 return h.arr[0]
}

// 删除堆顶元素
func (h *MaxHeap) remove() int {
 if len(h.arr) == 0 {
  return -1
 }
 root := h.arr[0]
 h.arr[0] = h.arr[len(h.arr)-1]
 h.arr = h.arr[:len(h.arr)-1]
 h.heapifyDown(0)
 return root
}

// 向下调整
func (h *MaxHeap) heapifyDown(i int) {
 lastIndex := len(h.arr) - 1

 // 左中右找最大节点
 largest := i

// 左节点
 left := 2*i + 1
 if left <= lastIndex && h.arr[left] > h.arr[largest] {
  largest = left
 }
// 右节点
 right := 2*i + 2
 if right <= lastIndex && h.arr[right] > h.arr[largest] {
  largest = right
 }
// 交换节点
 if largest != i {
  h.arr[i], h.arr[largest] = h.arr[largest], h.arr[i]
  // 递归下沉
  h.heapifyDown(largest)
 }
}

func main() {
 h := &MaxHeap{}
 h.insert(10)
 h.insert(20)
 h.insert(5)
 h.insert(30)
 fmt.Println("Heap:", h.arr)
 fmt.Println("Peek:", h.peek())
 fmt.Println("Remove:", h.remove())
 fmt.Println("Heap after remove:", h.arr)
}
```

**最小堆示例：**

```golang
package main

import (
 "fmt"
)

// 堆结构体
type MinHeap struct {
 arr []int
}

// 构造最小堆
func (h *MinHeap) heapifyUp(i int) {
    // 区别就是判断条件
 for i > 0 && h.arr[i] < h.arr[(i-1)/2] {
  h.arr[i], h.arr[(i-1)/2] = h.arr[(i-1)/2], h.arr[i]
  i = (i - 1) / 2
 }
}

// 插入元素
func (h *MinHeap) insert(val int) {
 h.arr = append(h.arr, val)
 h.heapifyUp(len(h.arr) - 1)
}

// 获取堆顶元素
func (h *MinHeap) peek() int {
 return h.arr[0]
}

// 删除堆顶元素
func (h *MinHeap) remove() int {
 if len(h.arr) == 0 {
  return -1
 }
 root := h.arr[0]
 h.arr[0] = h.arr[len(h.arr)-1]
 h.arr = h.arr[:len(h.arr)-1]
 h.heapifyDown(0)
 return root
}

// 向下调整
func (h *MinHeap) heapifyDown(i int) {
 lastIndex := len(h.arr) - 1
 smallest := i

 left := 2*i + 1
 if left <= lastIndex && h.arr[left] < h.arr[smallest] {
  smallest = left
 }

 right := 2*i + 2
 if right <= lastIndex && h.arr[right] < h.arr[smallest] {
  smallest = right
 }

 if smallest != i {
  h.arr[i], h.arr[smallest] = h.arr[smallest], h.arr[i]
  h.heapifyDown(smallest)
 }
}

func main() {
 h := &MinHeap{}
 h.insert(10)
 h.insert(20)
 h.insert(5)
 h.insert(30)
 fmt.Println("Heap:", h.arr)
 fmt.Println("Peek:", h.peek())
 fmt.Println("Remove:", h.remove())
 fmt.Println("Heap after remove:", h.arr)
}

```
