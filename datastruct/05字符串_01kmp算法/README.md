# 定义

KMP（Knuth-Morris-Pratt）算法是一种**改进的字符串匹配算法**，通过预处理模式串（Pattern）生成一个**部分匹配表（next数组）**，可以在匹配过程中**避免不必要的重复比较**，使得匹配效率更高。

传统的**暴力匹配**算法在遇到字符不匹配时，会将模式串整体后移一位，时间复杂度是 **O(n × m)**。而KMP算法利用已经匹配的信息，将模式串向右移动多位，从而提高效率，时间复杂度为 **O(n + m)**。

## 核心思想

KMP算法的核心是通过**部分匹配表（next数组）**来记录模式串中**前缀和后缀**匹配的情况。在匹配过程中，利用这些信息**跳过已经匹配的部分**，避免重新比较。

- **前缀**：字符串的前面部分，例如 `abc` 的前缀是 `a`, `ab`。
- **后缀**：字符串的后面部分，例如 `abc` 的后缀是 `c`, `bc`。
- **部分匹配值**：前缀和后缀的**最长公共部分**的长度。

---

## KMP算法的两个步骤  

### 1. 预处理模式串：构建next数组  

next数组记录模式串中**前缀和后缀的最长匹配长度**，用来确定不匹配时模式串的下一个位置。

**构建next数组的规则：**  

- 遍历模式串，检查每个位置前缀和后缀的最长匹配情况。  
- 如果字符匹配，则前缀长度加1。  
- 如果字符不匹配，利用next数组回退到前一个最长匹配位置，继续检查。  

示例模式串：abcab
我们逐步计算每个位置的next值。

|位置|字符|前缀|后缀|最长匹配长度|next值|
|---|---|---|---|---|---|
|0|a|无|无|0|0|
|1|b|a|无|0|0|
|2|c|a, ab|无|0|0|
|3|a|a, ab, abc|无|1 (a)|1|
|4|b|a, ab, abc, abca|b|2 (ab)|2|

next数组为：[0, 0, 0, 1, 2]

### 2. 匹配过程  

- 遍历文本串，每次比较文本串和模式串的字符。  
- 如果字符不匹配，使用next数组将模式串向右移动到合适的位置（跳过已匹配的部分）。  
- 如果字符匹配，继续向后比较。  
- 当模式串全部匹配时，记录匹配位置。  

示例文本串：ababcababc，模式串：abc

#### 第一轮匹配

- 比较 text[0] 和 pattern[0]，匹配，继续。
- 比较 text[1] 和 pattern[1]，匹配，继续。
- 比较 text[2] 和 pattern[2]，匹配，完成一次匹配，记录位置。

#### 第二轮匹配（利用next数组）

- 跳到 text[3] 继续匹配，发现不匹配，利用next回退，不浪费时间。

重复以上过程，直到文本串遍历完毕。

## 示例代码

```golang
package main

import "fmt"

// 计算 next 数组
func computeNext(pattern string) []int {
    next := make([]int, len(pattern)) // 初始化 next 数组
    j := 0                           // 前缀指针

    for i := 1; i < len(pattern); i++ {
        // 当不匹配时，回退到前一个匹配位置
        for j > 0 && pattern[i] != pattern[j] {
            j = next[j-1]
        }
        // 如果字符匹配，前缀长度增加
        if pattern[i] == pattern[j] {
            j++
        }
        next[i] = j
    }
    return next
}

// KMP 字符串匹配算法
func kmpSearch(text, pattern string) []int {
    next := computeNext(pattern) // 生成 next 数组
    positions := []int{}         // 记录匹配位置
    j := 0                       // 模式串指针

    for i := 0; i < len(text); i++ {
        // 字符不匹配时，回退到前一个匹配位置
        for j > 0 && text[i] != pattern[j] {
            j = next[j-1]
        }
        // 字符匹配时，移动指针
        if text[i] == pattern[j] {
            j++
        }
        // 完全匹配，记录位置
        if j == len(pattern) {
            positions = append(positions, i-j+1)
            j = next[j-1] // 继续寻找下一个匹配位置
        }
    }
    return positions
}

func main() {
    text := "ababcababc"
    pattern := "abc"

    positions := kmpSearch(text, pattern)
    fmt.Println("匹配位置:", positions)
}
```

## 特点

- 时间复杂度低：时间复杂度为O(n + m)，n是文本串长度，m是模式串长度。

- 避免重复比较：通过next数组实现跳跃匹配，减少字符比较次数。

- 适合大规模匹配：在长文本和短模式串的匹配场景中，性能优势明显。

- 部分匹配表（next数组）：是KMP算法的核心，记录前缀和后缀的匹配信息。
