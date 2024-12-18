# 定义

有限状态自动机（Finite State Machine，简称 FSM）是一种数学模型，表示一个系统在不同状态之间的转换。FSM可以通过状态图来表示，其中每个状态与输入事件相关联，并根据输入事件做出状态的转换。对于字符串匹配问题，有限状态自动机通过逐步匹配输入字符串中的每个字符，状态转换来判断是否匹配成功。

对于字符串匹配问题，FSM通常会根据模式串的每个字符创建状态，并根据字符的输入决定是否进入下一个状态。

---

## 示例

golang

```golang
package main

import "fmt"

// 构建有限状态自动机
type FSM struct {
    transitions map[int]map[rune]int // 状态转移
    finalState  int                   // 最终状态
}

// 构造 FSM 状态机
func NewFSM(pattern string) *FSM {
    fsm := &FSM{
        transitions: make(map[int]map[rune]int),
        finalState:  len(pattern),
    }

    currentState := 0

    // 构建状态机的转移表
    for i := 0; i < len(pattern); i++ {
        currentChar := rune(pattern[i])

        if _, ok := fsm.transitions[currentState]; !ok {
            fsm.transitions[currentState] = make(map[rune]int)
        }

        fsm.transitions[currentState][currentChar] = currentState + 1
        currentState++
    }

    return fsm
}

// 使用 FSM 进行字符串匹配
func (fsm *FSM) Match(text string) bool {
    currentState := 0
    for _, char := range text {
        // 检查当前字符在当前状态下的转移
        if nextState, ok := fsm.transitions[currentState][char]; ok {
            currentState = nextState
        } else {
            currentState = 0 // 如果没有匹配到，重新开始匹配
        }

        // 如果到达了最终状态，说明匹配成功
        if currentState == fsm.finalState {
            return true
        }
    }
    return false
}

func main() {
    pattern := "abc"
    fsm := NewFSM(pattern)
    text := "abcdef"

    if fsm.Match(text) {
        fmt.Println("匹配成功")
    } else {
        fmt.Println("匹配失败")
    }
}
```

## 特点

- 状态转移清晰：有限状态自动机通过状态和输入字符的映射来决定状态的变化，适合模式匹配。
- 效率较高：构建状态机后，匹配过程时间复杂度为O(n)，非常高效。
- 适合于确定性匹配：FSM适合匹配规则固定、字符转移明确的字符串匹配问题。
- 灵活性较差：对于动态变化的模式串，有限状态自动机的构建过程可能不如KMP算法灵活。
