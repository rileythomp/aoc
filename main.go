package main

import (
	"fmt"
	"os"

	"github.com/rileythomp/aoc/mkday"
	"github.com/rileythomp/aoc/stats"
	"github.com/rileythomp/aoc/submit"
)

func printUsage() {
	fmt.Println("Advent of Code CLI")
}

func main() {
	args := os.Args[1:]
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

	if prog == "submissions" {
		stats.RunSubmissions()
	} else if prog == "mkday" {
		mkday.RunMkday(args[1:])
	} else if prog == "submit" {
		submit.RunSubmit(args[1:])
	}
}
