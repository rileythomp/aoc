package aoc

import "strconv"

func StrToNums(str string) []int {
	nums := []int{}
	for _, c := range str {
		num, _ := strconv.Atoi(string(c))
		nums = append(nums, num)
	}
	return nums
}

func ContainsInt(list []int, elem int) bool {
	for _, e := range list {
		if elem == e {
			return true
		}
	}
	return false
}

func ContainsStr(list []string, elem string) bool {
	for _, e := range list {
		if elem == e {
			return true
		}
	}
	return false
}

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

func Abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

type Queue []interface{}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Push(e interface{}) {
	*q = append(*q, e)
}

func (q *Queue) Pop() interface{} {
	if q.IsEmpty() {
		return nil
	}
	e := (*q)[0]
	*q = (*q)[1:]
	return e
}

func (s *Queue) Front() interface{} {
	if s.IsEmpty() {
		return nil
	}
	e := (*s)[0]
	return e
}

type Stack []interface{}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(e interface{}) {
	*s = append(*s, e)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	i := len(*s) - 1
	e := (*s)[i]
	*s = (*s)[:i]
	return e
}

func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	e := (*s)[len(*s)-1]
	return e
}
