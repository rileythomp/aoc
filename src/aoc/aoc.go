package aoc

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type List []interface{}

func (l *List) Len() int {
	return len(*l)
}

func (l *List) IsEmpty() bool {
	return l.Len() == 0
}

func (l *List) Prepend(elem interface{}) {
	*l = append([]interface{}{elem}, *l...)
}

func (l *List) Append(elem interface{}) {
	*l = append(*l, elem)
}

func (l *List) Contains(elem interface{}) bool {
	for _, e := range *l {
		if elem == e {
			return true
		}
	}
	return false
}

func (l *List) At(i int) interface{} {
	if l.IsEmpty() || i < 0 || i >= l.Len() {
		return nil
	}
	return (*l)[i]
}

func (l *List) InsertAt(elem interface{}, i int) {
	elems := make([]interface{}, l.Len()+1)
	for j := range elems {
		if j < i {
			elems[j] = l.At(j)
		} else if j == i {
			elems[j] = elem
		} else {
			elems[j] = l.At(j - 1)
		}
	}
	*l = elems
}

func (l *List) First() interface{} {
	if l.IsEmpty() {
		return nil
	}
	return (*l)[0]
}

func (l *List) Last() interface{} {
	if l.IsEmpty() {
		return nil
	}
	return (*l)[l.Len()-1]
}

func (l *List) Pop() interface{} {
	if l.IsEmpty() {
		return nil
	}
	i := l.Len() - 1
	e := (*l)[i]
	*l = (*l)[:i]
	return e
}

func StrToNums(str string) []int {
	nums := make([]int, len(str))
	for i, c := range str {
		num, _ := strconv.Atoi(string(c))
		nums[i] = num
	}
	return nums
}

func IsUppercase(str string) bool {
	for _, c := range str {
		if !strings.Contains("ABCDEFGHIJKLMNOPQRSTUVWXYZ", string(c)) {
			return false
		}
	}
	return true
}

func IsLowercase(str string) bool {
	for _, c := range str {
		if !strings.Contains("abcdefghijklmnopqrstuvwxyz", string(c)) {
			return false
		}
	}
	return true
}

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Mean(nums []int) float64 {
	return float64(Sum(nums)) / float64(len(nums))
}

func Median(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	return nums[len(nums)/2]
}

func Mode(nums []int) int {
	freqMap := map[int]int{}
	mode, modeFreq := 0, 0
	for _, num := range nums {
		if _, ok := freqMap[num]; ok {
			freqMap[num]++
			if freqMap[num] > modeFreq {
				modeFreq = freqMap[num]
				mode = num
			}
		} else {
			freqMap[num] = 1
		}
	}
	return mode
}

func Largest(nums []int) int {
	max := math.MinInt32
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func Smallest(nums []int) int {
	min := math.MaxInt32
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
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
