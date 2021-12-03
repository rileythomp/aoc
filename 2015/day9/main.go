package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	Dist    int
	MinPath int
}

type Node struct {
	Visited bool
	Edges   map[string]*Edge
}

func findShortestPath(graph map[string]*Node) int {
	minPath := int(^uint(0) / 2)
	// have a graph, now find shortest path
	// for name := range graph {
	// 	node := graph[name]
	// 	visited := 1
	// 	for visited != len(graph) {
	// 		// pick an unvisited neighbour
	// 		for name, edge := range node.Edges {

	// 		}
	// 	}

	// }
	return minPath
}

func part1(strs []string) int {
	graph := map[string]*Node{}
	for _, str := range strs {
		parts := strings.Split(str, " ")
		node1, node2, d := parts[0], parts[2], parts[4]
		dist, _ := strconv.Atoi(d)
		if _, ok := graph[node1]; !ok {
			node := &Node{
				Visited: false,
				Edges:   map[string]*Edge{},
			}
			graph[node1] = node
		}
		if _, ok := graph[node2]; !ok {
			node := &Node{
				Visited: false,
				Edges:   map[string]*Edge{},
			}
			graph[node2] = node
		}
		edge := &Edge{
			Dist:    dist,
			MinPath: dist,
		}
		graph[node1].Edges[node2] = edge
		graph[node2].Edges[node1] = edge
	}

	// JSONString, err := json.MarshalIndent(graph, "", "	")
	// if err != nil {
	// 	return -1
	// }
	// fmt.Println(string(JSONString))

	return findShortestPath(graph)
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
