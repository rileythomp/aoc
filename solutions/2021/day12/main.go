package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rileythomp/aoc/src/aoc"
)

type Node struct {
	name string
	nbrs []string
}

func removeNeighbour(node Node, nbr string) {
	for i, n := range node.nbrs {
		if n == nbr {
			node.nbrs[i] = "*"
		}
	}
}

func containsNode(path Stack, node string) bool {
	for _, s := range path {
		if s.name == node {
			return true
		}
	}
	return false
}

func createGraph(strs []string) map[string][]string {
	nodeGraph := map[string][]string{}
	for _, line := range strs {
		nodes := strings.Split(line, "-")
		node1, node2 := nodes[0], nodes[1]
		if _, ok := nodeGraph[node1]; ok && node2 != "start" && node1 != "end" {
			nodeGraph[node1] = append(nodeGraph[node1], node2)
		} else if node2 != "start" && node1 != "end" {
			nodeGraph[node1] = []string{node2}
		}
		if _, ok := nodeGraph[node2]; ok && node1 != "start" && node2 != "end" {
			nodeGraph[node2] = append(nodeGraph[node2], node1)
		} else if node1 != "start" && node2 != "end" {
			nodeGraph[node2] = []string{node1}
		}
	}
	return nodeGraph
}

func part1(strs []string) int {
	paths := 0
	nodeGraph := createGraph(strs)
	curPath := Stack{Node{name: "start", nbrs: append([]string{}, nodeGraph["start"]...)}}
	for !curPath.IsEmpty() {
		curNode := curPath.Top()
		if curNode.name == "end" {
			paths++
			curPath.Pop()
			continue
		}
		added := false
		for _, nbr := range curNode.nbrs {
			if aoc.IsUppercase(nbr) || (aoc.IsLowercase(nbr) && !containsNode(curPath, nbr)) {
				added = true
				curPath.Push(Node{name: nbr, nbrs: append([]string{}, nodeGraph[nbr]...)})
				if aoc.IsLowercase(nbr) {
					removeNeighbour(curNode, nbr)
				}
				break
			} else {
				removeNeighbour(curNode, nbr)
			}
		}
		if !added {
			prev, cur := curPath.Pop(), curPath.Top()
			removeNeighbour(cur, prev.name)
		}
	}
	return paths
}

func smallCaveTwice(path Stack) bool {
	smallCaves := make(map[string]bool)
	for _, s := range path {
		if aoc.IsLowercase(s.name) {
			if _, ok := smallCaves[s.name]; ok {
				return true
			} else {
				smallCaves[s.name] = true
			}
		}
	}
	return false
}

func part2(strs []string) int {
	paths := 0
	nodeGraph := createGraph(strs)
	curPath := Stack{Node{name: "start", nbrs: append([]string{}, nodeGraph["start"]...)}}
	for !curPath.IsEmpty() {
		curNode := curPath.Top()
		if curNode.name == "end" {
			paths++
			prev, cur := curPath.Pop(), curPath.Top()
			removeNeighbour(cur, prev.name)
			continue
		}
		added := false
		for _, nbr := range curNode.nbrs {
			if aoc.IsUppercase(nbr) {
				added = true
				curPath.Push(Node{name: nbr, nbrs: append([]string{}, nodeGraph[nbr]...)})
				break
			} else if aoc.IsLowercase(nbr) && (!smallCaveTwice(curPath) || !containsNode(curPath, nbr)) {
				added = true
				curPath.Push(Node{name: nbr, nbrs: append([]string{}, nodeGraph[nbr]...)})
				if smallCaveTwice(curPath) {
					removeNeighbour(curNode, nbr)
				}
				break
			} else {
				removeNeighbour(curNode, nbr)
			}
		}
		if !added {
			prev, cur := curPath.Pop(), curPath.Top()
			removeNeighbour(cur, prev.name)
		}
	}
	return paths
}

func main() {
	level, name := getArgs()
	if level == "" || name == "" {
		return
	}

	file, _ := os.Open(name)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	strs := []string{}
	for scanner.Scan() {
		strs = append(strs, scanner.Text())
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
		level = "1"
		name  = "input.txt"
	)
	for i, arg := range args {
		if i == 0 && (level == "1" || level == "2") {
			level = arg
		} else if i == 0 {
			fmt.Printf("Level must be 1 or 2, got %s\n", arg)
			return "", ""
		} else if i == 1 {
			name = arg
		}
	}
	return level, name
}

type Stack []Node

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(e Node) {
	*s = append(*s, e)
}

func (s *Stack) Pop() Node {
	if s.IsEmpty() {
		return Node{}
	}
	i := len(*s) - 1
	e := (*s)[i]
	*s = (*s)[:i]
	return e
}

func (s *Stack) Top() Node {
	if s.IsEmpty() {
		return Node{}
	}
	e := (*s)[len(*s)-1]
	return e
}
