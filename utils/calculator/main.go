package main

import "unicode"

func strToInt(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		num := int(s[i] - '0')
		res = 10*res + num
	}
	return res
}

// 只支援+ - ，沒有 * /
func calculate1(s string) int {
	stk := []int{}
	sign := '+'
	res := 0
	tmpNum := 0
	for index, c := range s {
		if unicode.IsDigit(c) {
			tmpNum = 10*tmpNum + int(c-'0')
		}
		// 遇到 + or - 或者是最後一個str
		if c == '+' || c == '-' || index == len(s)-1 {
			switch sign {
			case '+':
				stk = append(stk, tmpNum)
			case '-':
				stk = append(stk, (-tmpNum))
			}
			// reset
			tmpNum = 0
			// update sign
			sign = c
		}
	}
	for _, num := range stk {
		res += num
	}
	return res
}

// 只支援+ - * / 沒有括號
func calculate2(s string) int {
	stk := []int{}
	sign := '+'
	res := 0
	tmpNum := 0
	for index, c := range s {
		if unicode.IsDigit(c) {
			tmpNum = 10*tmpNum + int(c-'0')
		}
		// 遇到 + or - or * or / 或者是最後一個str
		if c == '+' || c == '-' || c == '*' || c == '/' || index == len(s)-1 {
			switch sign {
			case '+':
				stk = append(stk, tmpNum)

			case '-':
				stk = append(stk, (-tmpNum))

			case '*':
				pre := stk[len(stk)-1]
				stk = stk[:len(stk)-1]
				stk = append(stk, pre*tmpNum)
			case '/':
				pre := stk[len(stk)-1]
				stk = stk[:len(stk)-1]
				stk = append(stk, pre/tmpNum)
			}
			// reset
			tmpNum = 0
			// update sign
			sign = c
		}
	}
	for _, num := range stk {
		res += num
	}
	return res
}

// 處理括號
func calculate3(s string) int {
	// 放括號的 index
	rightIndex := make(map[int]int)
	stack := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			rightIndex[left] = i
		}
	}
	return _calculate3(s, 0, len(s)-1, rightIndex)
}

func _calculate3(s string, startIndex, endIndex int, rightIndex map[int]int) int {
	stk := []int{}
	sign := '+'
	res := 0
	tmpNum := 0
	for i := startIndex; i <= endIndex; i++ {
		c := s[i]
		if unicode.IsDigit(rune(c)) {
			tmpNum = 10*tmpNum + int(c-'0')
		}
		// 遇到括號
		if c == '(' {
			right := rightIndex[i]
			tmpNum = _calculate3(s, i+1, right-1, rightIndex)
			i = right
		}
		// 遇到 + or - or * or / 或者是最後一個str
		if c == '+' || c == '-' || c == '*' || c == '/' || i == endIndex {
			switch sign {
			case '+':
				stk = append(stk, tmpNum)

			case '-':
				stk = append(stk, (-tmpNum))

			case '*':
				pre := stk[len(stk)-1]
				stk = stk[:len(stk)-1]
				stk = append(stk, pre*tmpNum)
			case '/':
				pre := stk[len(stk)-1]
				stk = stk[:len(stk)-1]
				stk = append(stk, pre/tmpNum)
			}
			// reset
			tmpNum = 0
			// update sign
			sign = rune(c)
		}
	}
	for _, num := range stk {
		res += num
	}
	return res
}
