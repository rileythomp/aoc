package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	Crossed bool
	Dist    int
}

type Node struct {
	Name        string
	Visited     bool
	Edges       map[string]*Edge
	LastVisited string
}

func (n *Node) print() {
	fmt.Println("{")
	fmt.Printf("\tName: %s\n", n.Name)
	fmt.Printf("\tVisited: %t\n", n.Visited)
	JSONString, _ := json.MarshalIndent(n.Edges, "	", "	")
	fmt.Printf("\t%s\n", string(JSONString))
	fmt.Println("}")

}

func NewNode(name string) *Node {
	return &Node{
		Name:        name,
		Visited:     false,
		Edges:       map[string]*Edge{},
		LastVisited: "",
	}
}

func printJSON(i interface{}) {
	JSONString, _ := json.MarshalIndent(i, "", "	")
	fmt.Println(string(JSONString))
}

// TODO: This is close, but need to find a way to allow
// going back to the last visited node if the path (stack) is different
// could try a hacky string representation
func findShortestPath(graph map[string]*Node) int {
	minPath := int(^uint(0) / 2)
	for start := range graph {
		graph[start].Visited = true
		stack := &Stack{graph[start]}
		curPath, i := 0, 0
		_ = i
		fmt.Println()
		for !stack.IsEmpty() {
			for i, s := range *stack {
				fmt.Print(s.Name)
				if i != len(*stack)-1 {
					fmt.Print("\n")
				}
			}
			fmt.Printf(" %d\n", curPath)
			fmt.Println()
			curNode := stack.Top()
			// fmt.Printf("processing %s\n", curNode.Name)
			// curNode.print()
			var (
				neighbour string
				// neighbour *Node
				// edge      *Edge
			)
			for name := range curNode.Edges {
				// if !curNode.Edges[n].Crossed && !graph[n].Visited && {
				if name != curNode.LastVisited && !graph[name].Visited {
					neighbour = name
					// neighbour = graph[name]
					// edge = curNode.Edges[name]
					break
				}
			}
			if neighbour != "" {
				// fmt.Printf("%-15s %3d %3d\n", curNode.Name, curPath, edge.Dist)
				stack.Push(graph[neighbour])
				graph[neighbour].Visited = true
				// curNode.Edges[name].Crossed = true
				curPath += curNode.Edges[neighbour].Dist
				curNode.LastVisited = neighbour
			} else {
				// fmt.Printf("%-15s %3d %3d\n", curNode.Name, curPath, 0)
				// if len(*stack) == len(graph) {
				// 	fmt.Println()
				// 	for _, s := range *stack {
				// 		fmt.Println(s.Name)
				// 	}
				// 	fmt.Println(curPath)
				// 	fmt.Println()
				// }
				end := stack.Pop()
				// fmt.Printf("popped %s\n", end.Name)
				curNode.Visited = false
				if !stack.IsEmpty() {
					curPath -= end.Edges[stack.Top().Name].Dist
				}
			}
		}
		break
	}
	printJSON(graph)
	return minPath
}

func initGraph(strs []string) map[string]*Node {
	graph := map[string]*Node{}
	for _, str := range strs {
		parts := strings.Split(str, " ")
		node1, node2, d := parts[0], parts[2], parts[4]
		dist, _ := strconv.Atoi(d)
		if _, ok := graph[node1]; !ok {
			graph[node1] = NewNode(node1)
		}
		if _, ok := graph[node2]; !ok {
			graph[node2] = NewNode(node2)
		}
		edge1 := Edge{
			// Crossed: false,
			Dist: dist,
		}
		edge2 := edge1
		graph[node1].Edges[node2] = &edge1
		graph[node2].Edges[node1] = &edge2
	}
	return graph
}

func part1(strs []string) int {
	return findShortestPath(initGraph(strs))
}

func part2(strs []string) int {
	return 0
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

type Stack []*Node

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(e *Node) {
	*s = append(*s, e)
}

func (s *Stack) Pop() *Node {
	if s.IsEmpty() {
		return nil
	}
	i := len(*s) - 1
	e := (*s)[i]
	*s = (*s)[:i]
	return e
}

func (s *Stack) Top() *Node {
	if s.IsEmpty() {
		return nil
	}
	e := (*s)[len(*s)-1]
	return e
}
