package mkday

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/rileythomp/aoc/utils"
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("./aoc mkday <day> <year>")
	fmt.Println("Defaults:")
	fmt.Println("day:  next day")
	fmt.Println("year: current year")
	fmt.Println("So if no arguments are passed, it will wait until the next puzzle is released at midnight")
}

func RunMkday(args []string) {
	year, day, ok := getYearAndDay(args)
	if !ok {
		return
	}

	err := createFiles(year, day)
	if err != nil {
		return
	}

	_ = openProblem(year, day)
}

func getYearAndDay(args []string) (string, string, bool) {
	y, _, d := time.Now().Date()
	var (
		year = fmt.Sprint(y)
		day  = fmt.Sprint(d + 1)
	)
	for i, arg := range args {
		if arg == "-h" || arg == "--help" {
			printUsage()
			return "", "", false
		}
		if i == 0 {
			day = arg
		} else if i == 1 {
			year = arg
		}
	}
	if year == fmt.Sprint(y) && day == fmt.Sprint(d+1) {
		fmt.Println("Waiting until problem is released at midnight...")
		curDay, seconds := d, 0
		for curDay < d+1 {
			time.Sleep(time.Second)
			seconds++
			_, _, curDay = time.Now().Date()
		}
		fmt.Printf("Waited for %d minutes and %d seconds\n", seconds/60, seconds%60)
	}
	return year, day, true
}

func createFiles(year, day string) error {
	fmt.Printf("Downloading %s day %s...\n", year, day)
	path := fmt.Sprintf("./%s/day%s", year, day)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating %s: %s", path, err.Error())
		return err
	}
	uri := fmt.Sprintf("https://adventofcode.com/%s/day/%s", year, day)
	problem, err := utils.GetAoC(uri)
	if err != nil {
		return err
	}
	input, err := utils.GetAoC(uri + "/input")
	if err != nil {
		return err
	}
	boilerplate, err := getBoilerplate()
	if err != nil {
		return err
	}
	files := []struct {
		Name    string
		Content []byte
	}{
		{Name: "problem.html", Content: problem},
		{Name: "input.txt", Content: input},
		{Name: "test.txt", Content: []byte{}},
		{Name: "main.go", Content: boilerplate},
	}
	for _, file := range files {
		file.Content = []byte(strings.Replace(string(file.Content), "/static/style.css?26", "https://adventofcode.com/static/style.css", 1))
		err = os.WriteFile(fmt.Sprintf("%s/%s", path, file.Name), file.Content, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating %s/%s\n", path, file.Name)
			return err
		}
		fmt.Printf("Created %s/%s\n", path, file.Name)
	}
	return nil
}

func getBoilerplate() ([]byte, error) {
	boilerplate, err := ioutil.ReadFile("./mkday/boilerplate.txt")
	if err != nil {
		return nil, err
	}
	return boilerplate, nil
}

func openProblem(year, day string) error {
	filePath := fmt.Sprintf("./%s/day%s/problem.html", year, day)
	cmd := exec.Command("open", filePath)
	_, err := cmd.Output()
	if err != nil {
		fmt.Printf("There was an error opening %s: %s\n", filePath, err.Error())
		return err
	}
	return nil
}
