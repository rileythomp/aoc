package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	}
	i := len(*s) - 1
	str := (*s)[i]
	*s = (*s)[:i]
	return str
}

func (s *Stack) Top() string {
	if s.IsEmpty() {
		return ""
	}
	str := (*s)[len(*s)-1]
	return str
}

type Wire struct {
	Val  uint16
	Gate string
}

func processGatePart(gatePart string, wires map[string]Wire) (uint16, bool) {
	if num, err := strconv.Atoi(gatePart); err == nil {
		return uint16(num), true
	}
	if wire, ok := wires[gatePart]; ok {
		return wire.Val, true
	}
	return 0, false
}

func processGate(gate string, wires map[string]Wire) (uint16, string, string) {
	// X
	if !strings.Contains(gate, " ") {
		val, ok := processGatePart(gate, wires)
		if ok {
			return val, "", ""
		}
		return 0, gate, ""
	}
	gateParts := strings.Split(gate, " ")
	// NOT X
	if strings.Contains(gate, "NOT") {
		val, ok := processGatePart(gateParts[1], wires)
		if ok {
			return ^val, "", ""
		}
		return 0, gateParts[1], ""
	}
	if len(gateParts) < 3 {
		fmt.Printf("Unexpected gate: %s, %s\n", gate, gateParts)
		return 0, "", ""
	}
	l, lOk := processGatePart(gateParts[0], wires)
	r, rOk := processGatePart(gateParts[2], wires)
	if strings.Contains(gate, "AND") && lOk && rOk {
		return l & r, "", ""
	}
	if strings.Contains(gate, "OR") && lOk && rOk {
		return l | r, "", ""
	}
	if strings.Contains(gate, "LSHIFT") && lOk && rOk {
		return l << r, "", ""
	}
	if strings.Contains(gate, "RSHIFT") && rOk && lOk {
		return l >> r, "", ""
	}
	lw, rw := "", ""
	if !lOk {
		lw = gateParts[0]
	}
	if !rOk {
		rw = gateParts[2]
	}
	return 0, lw, rw
}

func getGate(w string, strs []string) string {
	for _, str := range strs {
		parts := strings.Split(str, " -> ")
		gate, wire := parts[0], parts[1]
		if w == wire {
			return gate
		}
	}
	return ""
}

func part1(strs []string) int {
	wires := map[string]Wire{}
	stack := &Stack{"a"}
	for !stack.IsEmpty() {
		w := stack.Top()
		var gate string
		if wire, ok := wires[w]; ok {
			gate = wire.Gate
		} else {
			gate = getGate(w, strs)
		}
		if gate == "" {
			fmt.Printf("could not find gate for %s\n", w)
			return -1
		}
		val, l, r := processGate(gate, wires)
		if l == "" && r == "" {
			wires[w] = Wire{
				Val:  val,
				Gate: gate,
			}
			stack.Pop()
		} else {
			wires[w] = Wire{
				Val:  0,
				Gate: gate,
			}
			if l != "" {
				stack.Push(l)
			}
			if r != "" {
				stack.Push(r)
			}
		}
	}
	a, _ := wires["a"]
	return int(a.Val)
}

func part2(strs []string) int {
	for i, str := range strs {
		if strings.Contains(str, " -> b") {
			strs[i] = fmt.Sprintf("%d -> b", part1(strs))
			break
		}
	}
	return part1(strs)
}

func main() {
	level, fileName := getArgs()
	if level == "" || fileName == "" {
		return
	}

	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	strs := []string{}
	nums := []int{}
	for scanner.Scan() {
		str := scanner.Text()
		strs = append(strs, str)
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}

	if level == "1" {
		fmt.Println(part1(strs))
	} else if level == "2" {
		fmt.Println(part2(strs))
	}
}

func getArgs() (string, string) {
	args := os.Args[1:]
	var (
		level    = "1"
		fileName = "input.txt"
	)
	for i, arg := range args {
		if i == 0 && (level == "1" || level == "2") {
			level = arg
		} else if i == 0 {
			fmt.Printf("Level must be 1 or 2, got %s\n", arg)
			return "", ""
		} else if i == 1 {
			fileName = arg
		}
	}
	return level, fileName
}
