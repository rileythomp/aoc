package stats

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rileythomp/aoc/utils"
)

func RunSubmissions() {
	for {
		if err := writeStats(); err != nil {
			return
		}
		time.Sleep(600 * time.Second)
	}
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
			statsFile := fmt.Sprintf("%dstats.csv", day)
			f, _ := os.OpenFile(statsFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
			defer f.Close()
			if _, err = f.WriteString(entry); err != nil {
				return err
			}
		}
	}
	return nil
}
