package stats

import (
	"fmt"
	"strings"
	"time"

	"github.com/rileythomp/aoc/src/utils"
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

	_, _, d := time.Now().Date()
	fmt.Println("Waiting until problem is released at midnight...")
	curDay, seconds := d, 0
	for curDay < d+1 {
		time.Sleep(time.Second)
		seconds++
		_, _, curDay = time.Now().Date()
	}
	fmt.Printf("Waited for %d minutes and %d seconds\n", seconds/60, seconds%60)

	reqs := 0
	year, _, day := time.Now().Date()
	for {
		if reqs < 60 {
			time.Sleep(time.Minute)
		} else if reqs-60 < 24 {
			time.Sleep(time.Hour)
		} else {
			break
		}
		reqs++
		if err := writeStats(day, year); err != nil {
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

func writeStats(day, year int) error {
	uri := fmt.Sprintf("https://adventofcode.com/%d/stats", year)
	statsData, err := utils.GetAoC(uri)
	if err != nil {
		return err
	}
	html := string(statsData)
	lines := strings.Split(html, "\n")
	for _, line := range lines {
		if strings.Contains(line, fmt.Sprintf(" %d ", day)) {
			parts := strings.Fields(line)
			if len(parts) < 7 {
				return fmt.Errorf("unexpected response from %s: %s", uri, line)
			}
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
