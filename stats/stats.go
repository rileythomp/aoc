package stats

import (
	"fmt"
	"strings"
	"time"

	"github.com/rileythomp/aoc/utils"
)

type Stats struct{}

func (s *Stats) PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("./aoc stats")
	fmt.Println("Checks the number of solutions every minute for the first hour and then")
	fmt.Println("every hour for the rest of the day and writes the output to <day>stats.csv")
}

func (s *Stats) Run(args []string) error {
	if _, ok := s.GetArgs(args); !ok {
		return nil
	}

	reqs := 0
	for {
		if reqs < 60 {
			time.Sleep(time.Minute)
		} else if reqs-60 < 24 {
			time.Sleep(time.Hour)
		} else {
			break
		}
		reqs++
		if err := writeStats(); err != nil {
			return err
		}
	}

	return nil
}

func (s *Stats) GetArgs(args []string) ([]string, bool) {
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			s.PrintUsage()
			return []string{}, false
		}
	}
	return []string{}, true
}

func writeStats() error {
	uri := "https://adventofcode.com/2021/stats"
	statsData, err := utils.GetAoC(uri)
	if err != nil {
		return err
	}
	_, _, day := time.Now().Date()
	html := string(statsData)
	lines := strings.Split(html, "\n")
	for _, line := range lines {
		if strings.Contains(line, fmt.Sprintf(" %d ", day)) {
			parts := strings.Split(line, " ")
			numspan := strings.Split(parts[6], "<")
			submissions := numspan[0]
			date := time.Now().Format("02 Jan 06 03:04:05PM")
			entry := date + ", " + submissions + "\n"
			name := fmt.Sprintf("%dstats.csv", day)
			if err := utils.WriteFileString(name, entry); err != nil {
				return err
			}
		}
	}
	return nil
}
