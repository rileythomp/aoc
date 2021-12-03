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
	MinPath int
}

type Node struct {
	Name    string
	Visited bool
	Edges   map[string]*Edge
}

func printJSON(i interface{}) {
	JSONString, _ := json.MarshalIndent(i, "", "	")
	fmt.Println(string(JSONString))
}

func findShortestPath(graph map[string]*Node) int {
	minPath := int(^uint(0) / 2)
	for _, node := range graph {
		node.Visited = true
		stack := &Stack{node}
		curPath, i := 0, 0
		_ = i
		for !stack.IsEmpty() {
			curNode := stack.Top()
			fmt.Printf("processing %s\n", curNode.Name)
			var (
				neighbour *Node
				edge      *Edge
			)
			for name := range curNode.Edges {
				if !curNode.Edges[name].Crossed && !graph[name].Visited {
					neighbour = graph[name]
					edge = curNode.Edges[name]
					break
				}
			}
			if edge != nil {
				// fmt.Printf("%-15s %3d %3d\n", curNode.Name, curPath, edge.Dist)
				stack.Push(neighbour)
				neighbour.Visited = true
				edge.Crossed = true
				curPath += edge.Dist
			} else {
				// fmt.Printf("%-15s %3d %3d\n", curNode.Name, curPath, 0)
				if len(*stack) == len(graph) {
					fmt.Println()
					for _, s := range *stack {
						fmt.Println(s.Name)
					}
					fmt.Println(curPath)
					fmt.Println()
				}
				end := stack.Pop()
				fmt.Printf("popped %s\n", end.Name)
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
			node := &Node{
				Name:    node1,
				Visited: false,
				Edges:   map[string]*Edge{},
			}
			graph[node1] = node
		}
		if _, ok := graph[node2]; !ok {
			node := &Node{
				Name:    node2,
				Visited: false,
				Edges:   map[string]*Edge{},
			}
			graph[node2] = node
		}
		edge1 := Edge{
			Crossed: false,
			Dist:    dist,
			MinPath: dist,
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
