package mkday

import (
	"fmt"
	"os"
	"time"

	"github.com/rileythomp/aoc/utils"
)

type Mkday struct{}

func (m *Mkday) PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("./aoc mkday <day> <year>")
	fmt.Println("Defaults:")
	fmt.Println("day:  next day")
	fmt.Println("year: current year")
	fmt.Println("So if no arguments are passed, it will wait until the next puzzle is released at midnight")
}

func (m *Mkday) Run(args []string) error {
	args, ok := m.GetArgs(args)
	if !ok {
		return nil
	}
	year, day := args[0], args[1]

	if err := m.createFiles(year, day); err != nil {
		return err
	}

	path := fmt.Sprintf("./solutions/%s/day%s/problem.html", year, day)
	if err = utils.OpenFile(path); err != nil {
		return err
	}

	return nil
}

func (m *Mkday) GetArgs(args []string) ([]string, bool) {
	y, _, d := time.Now().Date()
	year, day := fmt.Sprint(y), fmt.Sprint(d+1)
	for i, arg := range args {
		if arg == "-h" || arg == "--help" {
			m.PrintUsage()
			return []string{}, false
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
	return []string{year, day}, true
}

func (m *Mkday) createFiles(year, day string) error {
	fmt.Printf("Downloading %s day %s...\n", year, day)
	path := fmt.Sprintf("./solutions/%s/day%s", year, day)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
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
	boilerplate, err := utils.GetFile("./mkday/boilerplate.txt")
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
		file.Content = utils.AddCss(file.Content)
		fileName := fmt.Sprintf("%s/%s", path, file.Name)
		if err = utils.WriteFile(fileName, fileContent); err != nil {
			fmt.Printf("Error creating %s/%s\n", path, file.Name)
			return err
		}
		fmt.Printf("Created %s/%s\n", path, file.Name)
	}
	return nil
}
