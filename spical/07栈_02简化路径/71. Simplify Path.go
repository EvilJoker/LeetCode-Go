package leetcode

import (
	"path/filepath"
	"strings"
)

/*
题目： 简化路径
"/a//b////c/d//././/.."

需要会识别 . ..，还需要将 // 变成 /
思路： 使用栈
1. 先 split
2. 将 目录 push 到 栈中，遇到 . 和/ 啥也不干。
3. 遇到 .. 进行 pop
*/
type Stack struct {
	Data []string
	Len  int
}

func (st *Stack) Push(s string) {
	st.Data = append(st.Data, s)
	st.Len += 1
}

func (st *Stack) Pop() string {
	if st.Len == 0 {
		return ""
	}
	st.Len -= 1
	result := st.Data[st.Len]
	st.Data = st.Data[:st.Len]
	return result

}

func simplifyPath(path string) string {

	dirs := strings.Split(path, "/")

	stack := Stack{}
	for _, dir := range dirs {
		if dir == "" {
			continue
		}
		// 放入正常路径
		if dir != ".." && dir != "." {
			stack.Push(dir)
			continue
		}
		// 弹出
		if dir == ".." {

			stack.Pop()
		}

	}
	return "/" + strings.Join(stack.Data, "/")

}

// 解法一
func simplifyPathbk(path string) string {
	arr := strings.Split(path, "/")
	stack := make([]string, 0)
	var res string
	for i := 0; i < len(arr); i++ {
		cur := arr[i]
		//cur := strings.TrimSpace(arr[i]) 更加严谨的做法应该还要去掉末尾的空格
		if cur == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if cur != "." && len(cur) > 0 {
			stack = append(stack, arr[i])
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	res = strings.Join(stack, "/")
	return "/" + res
}

// 解法二 golang 的官方库 API
func simplifyPath1(path string) string {
	return filepath.Clean(path)
}
