package leetcode

/*
题目：括号校验，看看 ()[]{} 是否成对
思路：
使用栈 遇到左括号，就入栈，遇到有括号就出栈，如果不匹配就报错
*/

type Stack struct {
	Data []byte
	Len  int
}

func (s *Stack) Push(b byte) {
	s.Data = append(s.Data, b)
	s.Len += 1
}

func (s *Stack) Pop() byte {
	if s.Len == 0 {
		return 0
	}
	s.Len -= 1
	ret := s.Data[s.Len]
	s.Data = s.Data[:s.Len]

	return ret
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	sbytes := []byte(s)

	stack := Stack{}

	for _, c := range sbytes {

		if c == '(' || c == '[' || c == '{' {
			stack.Push(c)
			continue
		}

		if c == ')' {
			peer := stack.Pop()
			if peer != '(' || peer == 0 {
				return false
			}
			continue
		}

		if c == ']' {
			peer := stack.Pop()
			if peer != '[' || peer == 0 {
				return false
			}

			continue
		}

		if c == '}' {
			peer := stack.Pop()
			if peer != '{' || peer == 0 {
				return false
			}
			continue
		}

	}

	// 最后应该完美抵消才对

	if len(stack.Data) != 0 {
		return false
	}
	return true

}
func isValidbk(s string) bool {
	// 空字符串直接返回 true
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if (v == '[') || (v == '(') || (v == '{') {
			stack = append(stack, v)
		} else if ((v == ']') && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			((v == ')') && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			((v == '}') && len(stack) > 0 && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
