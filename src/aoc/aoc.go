package aoc

import "strconv"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Stack []interface{}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str interface{}) {
	*s = append(*s, str)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return ""
	}
	i := len(*s) - 1
	str := (*s)[i]
	*s = (*s)[:i]
	return str
}

func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		return ""
	}
	str := (*s)[len(*s)-1]
	return str
}

func StrToNums(str string) []int {
	nums := []int{}
	for _, c := range str {
		num, _ := strconv.Atoi(string(c))
		nums = append(nums, num)
	}
	return nums
}
