# 定义

广义栈: 是栈的一种扩展形式，可以表示多种数据结构。广义栈的每个元素不仅可以是基本类型的数据，还可以是另一个栈或复杂的嵌套数据结构。

## 分类

```golang
package main

import "fmt"

// 定义广义栈元素
type Element struct {
 isStack bool
 value   interface{}
}

// 定义广义栈
type GeneralizedStack struct {
 stack []Element
}

// 入栈操作
func (gs *GeneralizedStack) Push(elem Element) {
 gs.stack = append(gs.stack, elem)
}

// 出栈操作
func (gs *GeneralizedStack) Pop() Element {
 if len(gs.stack) == 0 {
  panic("Stack is empty")
 }
 top := gs.stack[len(gs.stack)-1]
 gs.stack = gs.stack[:len(gs.stack)-1]
 return top
}

// 打印广义栈
func (gs *GeneralizedStack) Print() {
 for _, elem := range gs.stack {
  if elem.isStack {
   fmt.Print("[Stack], ")
  } else {
   fmt.Printf("%v, ", elem.value)
  }
 }
 fmt.Println()
}

func main() {
 // 初始化广义栈
 gs := &GeneralizedStack{}

 // 添加基本类型元素
 gs.Push(Element{isStack: false, value: 10})
 gs.Push(Element{isStack: false, value: 20})

 // 添加嵌套栈
 nestedStack := &GeneralizedStack{}
 nestedStack.Push(Element{isStack: false, value: 30})
 gs.Push(Element{isStack: true, value: nestedStack})

 // 打印广义栈
 gs.Print()
}

```

## 特点

支持存储基本类型和其他栈，适合表达复杂嵌套结构。
基于栈的后进先出（LIFO）原则，可以进行灵活扩展。
实现较复杂，尤其是在处理嵌套数据时。
