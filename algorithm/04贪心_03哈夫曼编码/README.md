# 贪心之哈夫曼编码

## 题目介绍

哈夫曼编码（Huffman Coding）是一种贪心算法，常用于数据压缩。其目标是将频繁出现的字符用较短的编码表示，出现频率较低的字符用较长的编码表示，从而使得总体编码长度最短。

**输入**：一组字符及其出现的频率。

**输出**：一个哈夫曼编码，使得总编码长度最短。

## 解题思路

哈夫曼编码的基本思想是：根据字符出现的频率，使用优先队列（或最小堆）来构建一棵哈夫曼树。通过合并出现频率最小的两个节点，不断迭代，最终生成的树可以得到每个字符的最优编码。

### 步骤

1. 将每个字符及其频率放入一个最小堆（优先队列）。
2. 从堆中取出两个最小频率的节点，合并成一个新节点，并将新节点插入堆中。
3. 重复上述步骤，直到堆中只剩一个节点，这个节点就是哈夫曼树的根节点。
4. 从根节点开始遍历，生成每个字符的哈夫曼编码。

## 示例

假设我们有以下字符和它们的频率：

```text
字符: A, 频率: 5
字符: B, 频率: 9
字符: C, 频率: 12
字符: D, 频率: 13
字符: E, 频率: 16
字符: F, 频率: 45
```

```golang

package main

import (
 "container/heap"
 "fmt"
)

// 定义哈夫曼树节点
type HuffmanNode struct {
 char     rune
 frequency int
 left     *HuffmanNode
 right    *HuffmanNode
}

// 创建一个最小堆
type MinHeap []*HuffmanNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].frequency < h[j].frequency }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
 *h = append(*h, x.(*HuffmanNode))
}

func (h *MinHeap) Pop() interface{} {
 old := *h
 n := len(old)
 x := old[n-1]
 *h = old[0 : n-1]
 return x
}

// 构建哈夫曼树
func buildHuffmanTree(frequencies map[rune]int) *HuffmanNode {
 h := &MinHeap{}
 heap.Init(h)

 // 创建叶子节点
 for char, freq := range frequencies {
  heap.Push(h, &HuffmanNode{char: char, frequency: freq})
 }

 // 合并节点直到只剩一个
 for h.Len() > 1 {
  // 取出最小的两个节点
  left := heap.Pop(h).(*HuffmanNode)
  right := heap.Pop(h).(*HuffmanNode)

  // 创建新的父节点
  newNode := &HuffmanNode{
   frequency: left.frequency + right.frequency,
   left:      left,
   right:     right,
  }

  // 将新节点放回堆中
  heap.Push(h, newNode)
 }

 return heap.Pop(h).(*HuffmanNode)
}

// 生成哈夫曼编码
func generateHuffmanCodes(root *HuffmanNode, code string, codes map[rune]string) {
 if root == nil {
  return
 }

 // 如果是叶子节点，记录字符的编码
 if root.left == nil && root.right == nil {
  codes[root.char] = code
 }

 // 递归生成编码
 generateHuffmanCodes(root.left, code+"0", codes)
 generateHuffmanCodes(root.right, code+"1", codes)
}

func main() {
 // 输入字符及其频率
 frequencies := map[rune]int{
  'A': 5,
  'B': 9,
  'C': 12,
  'D': 13,
  'E': 16,
  'F': 45,
 }

 // 构建哈夫曼树
 root := buildHuffmanTree(frequencies)

 // 生成哈夫曼编码
 codes := make(map[rune]string)
 generateHuffmanCodes(root, "", codes)

 // 输出哈夫曼编码
 fmt.Println("Huffman Codes:")
 for char, code := range codes {
  fmt.Printf("%c: %s\n", char, code)
 }
}

```

贪心算法：每次选择频率最小的节点，合并它们，确保生成的树具有最小的总编码长度。
时间复杂度：由于每次合并操作都涉及堆的操作（O(log n)），总的时间复杂度为 O(n log n)，其中 n 是字符数。
空间复杂度：O(n)，用于存储哈夫曼树的节点和哈夫曼编码
