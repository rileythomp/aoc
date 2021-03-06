package main

import (
	"fmt"
	"os"

	"github.com/rileythomp/aoc/src/mkday"
	"github.com/rileythomp/aoc/src/stats"
	"github.com/rileythomp/aoc/src/submit"
)

type Program interface {
	Run(args []string) error
	PrintUsage()
	GetArgs(args []string) ([]string, bool)
}

func printUsage() {
	fmt.Println("Advent of Code CLI")
	fmt.Println("Usage:")
	fmt.Println("./aoc <program> [arguments]")
	fmt.Println("<program> can be one of submit, mkday or stats")
	fmt.Println("Use -h or --help for instructions.")
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printUsage()
		return
	}
	prog := "mkday"
	for i, arg := range args {
		if i == 0 && (arg == "-h" || arg == "--help") {
			printUsage()
			return
		}
		if i == 0 {
			prog = args[0]
		}
	}

	progs := map[string]Program{
		"stats":  &stats.Stats{},
		"mkday":  &mkday.Mkday{},
		"submit": &submit.Submit{},
	}
	if p, ok := progs[prog]; ok {
		if err := p.Run(args[1:]); err != nil {
			fmt.Println(err)
			p.PrintUsage()
		}
	} else {
		printUsage()
	}
}
