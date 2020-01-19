package main

type Stack []rune

func (s *Stack) Push(v rune) {
	*s = append(*s, v)
}

func (s *Stack) Pop() rune {
	ret := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return ret
}

func isValid(s string) bool {
	stack := new(Stack)
	parenthesesMap := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, c := range s {
		if _, ok := parenthesesMap[c]; !ok {
			stack.Push(c)
		} else if len(*stack) <= 0 || parenthesesMap[c] != stack.Pop() {
			return false
		}
	}
	return len(*stack) == 0
}
