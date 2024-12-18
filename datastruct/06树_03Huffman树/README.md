# 定义

**哈夫曼树**（Huffman Tree）是一种用于数据压缩的树结构，它通过对输入数据中各个字符的出现频率进行编码，生成一棵最优的二叉树。哈夫曼树的特点是频率越高的字符，编码越短；频率越低的字符，编码越长，从而有效减少了数据的存储空间。哈夫曼树常用于文件压缩、图像压缩等领域。

哈夫曼树的构建步骤通常是：

1. 统计每个字符的出现频率。
2. 将每个字符作为一个节点，并根据其频率构建最小堆。
3. 从最小堆中取出两个频率最小的节点，合并为一个新节点，新节点的频率为两个子节点的频率之和。
4. 将新节点插入堆中，直到堆中只剩下一个节点，这个节点即为哈夫曼树的根节点。

---

## 示例

```golang
package main

import (
 "container/heap"
 "fmt"
)

// 定义哈夫曼树的节点
type HuffmanNode struct {
 char  rune
 freq  int
 left  *HuffmanNode
 right *HuffmanNode
}

// 定义一个最小堆，用于构建哈夫曼树
type MinHeap []*HuffmanNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h MinHeap) Swap(i, j int) {
 h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
 *h = append(*h, x.(*HuffmanNode))
}

func (h *MinHeap) Pop() interface{} {
 old := *h
 n := len(old)
 item := old[n-1]
 *h = old[0 : n-1]
 return item
}

// 构建哈夫曼树
func buildHuffmanTree(freqMap map[rune]int) *HuffmanNode {
 h := &MinHeap{}
 heap.Init(h)

 // 将所有字符的频率插入到最小堆中
 for char, freq := range freqMap {
  heap.Push(h, &HuffmanNode{char: char, freq: freq})
 }

 // 合并节点，直到堆中只剩下一个节点
 for h.Len() > 1 {
  // 取出两个频率最小的节点
  left := heap.Pop(h).(*HuffmanNode)
  right := heap.Pop(h).(*HuffmanNode)

  // 创建一个新的节点，将两个节点合并
  newNode := &HuffmanNode{
   freq:  left.freq + right.freq,
   left:  left,
   right: right,
  }

  heap.Push(h, newNode)
 }

 // 返回最终的哈夫曼树根节点
 return heap.Pop(h).(*HuffmanNode)
}

// 遍历哈夫曼树生成编码
func generateHuffmanCodes(node *HuffmanNode, prefix string, codes map[rune]string) {
 if node == nil {
  return
 }

 // 如果是叶子节点，则保存对应字符的哈夫曼编码
 if node.left == nil && node.right == nil {
  codes[node.char] = prefix
  return
 }

 // 递归生成左子树和右子树的编码
 generateHuffmanCodes(node.left, prefix+"0", codes)
 generateHuffmanCodes(node.right, prefix+"1", codes)
}

func main() {
 // 输入字符串
 text := "this is an example for huffman encoding"

 // 统计字符频率
 freqMap := make(map[rune]int)
 for _, char := range text {
  freqMap[char]++
 }

 // 构建哈夫曼树
 root := buildHuffmanTree(freqMap)

 // 生成哈夫曼编码
 codes := make(map[rune]string)
 generateHuffmanCodes(root, "", codes)

 // 打印哈夫曼编码
 fmt.Println("Huffman Codes:")
 for char, code := range codes {
  fmt.Printf("%c: %s\n", char, code)
 }
}
```

## 特点

- 最优性：哈夫曼树可以生成最优的前缀编码，对于出现频率较高的字符，赋予较短的编码，对于出现频率较低的字符，赋予较长的编码，从而压缩数据大小。
- 适用范围：适用于文件压缩、数据传输等需要减少存储空间的场景。
- 构建过程：通过构建最小堆来生成哈夫曼树，时间复杂度为O(n log n)，其中n是字符的种类数。
- 压缩效果：哈夫曼编码能有效减少数据的冗余，尤其是在字符频率分布不均的情况下。
- 无前缀：哈夫曼编码生成的编码具有前缀性，即不会出现一个编码是另一个编码的前缀。
