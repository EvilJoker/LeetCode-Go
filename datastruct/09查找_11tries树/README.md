# 定义

Trie（字典树）是一种树形数据结构，用于存储集合中的字符串。它能够通过字符的逐层匹配来实现高效的查找操作。Trie 的每个节点表示字符串中的一个字符，树的路径代表了从根节点到叶节点的字符串。Trie 结构最常用于实现字典、自动补全、词频统计等场景。

---

## 示例

```golang
package main

import "fmt"

// TrieNode 是 Trie 树的节点结构
type TrieNode struct {
    children map[rune]*TrieNode // 子节点 // 字母表。每个节点都有全量的字母表
    isEnd    bool               // 是否是单词的结尾
}

// Trie 是字典树的结构
type Trie struct {
    root *TrieNode
}

// 初始化 Trie
func NewTrie() *Trie {
    return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// 插入单词
// 遍历每个字母，链式的查找每个字母对应的下一个节点。当到达结尾时，返回end 字符
func (t *Trie) Insert(word string) {
    node := t.root
    for _, char := range word {
        if _, exists := node.children[char]; !exists {
            node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
        }
        node = node.children[char]
    }
    node.isEnd = true
}

// 查找单词
// 也是遍历，看看最后一个字母是否有值
func (t *Trie) Search(word string) bool {
    node := t.root
    for _, char := range word {
        if _, exists := node.children[char]; !exists {
            return false
        }
        node = node.children[char]
    }
    return node.isEnd
}

// 前缀匹配
// 匹配前n 个字符
func (t *Trie) StartsWith(prefix string) bool {
    node := t.root
    for _, char := range prefix {
        if _, exists := node.children[char]; !exists {
            return false
        }
        node = node.children[char]
    }
    return true
}

func main() {
    trie := NewTrie()
    trie.Insert("apple")
    trie.Insert("app")

    fmt.Println(trie.Search("apple"))   // true
    fmt.Println(trie.Search("app"))     // true
    fmt.Println(trie.Search("appl"))    // false
    fmt.Println(trie.StartsWith("app")) // true
}
```

## 特点

高效的查找：Trie 在查找操作上时间复杂度为 O(m)，其中 m 为单词的长度。相比哈希表、链表等结构，能够更快地进行查找，尤其适用于前缀匹配。

节省空间：当有大量相同前缀的字符串时，Trie 会共享相同的前缀部分，从而节省内存空间。每个节点只需要存储一个字符，避免了重复存储。

支持前缀匹配：Trie 可以高效地实现前缀匹配功能，在自动补全等应用中非常有用。

单词插入与删除：Trie 可以方便地实现单词的插入和删除操作，操作时间复杂度为 O(m)，其中 m 为单词长度。

适用于词典和自动补全：Trie 结构非常适合实现词典、拼写检查、自动补全等功能。
