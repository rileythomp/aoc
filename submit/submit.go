package submit

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/rileythomp/aoc/utils"
)

type Submit struct{}

func (s *Submit) PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("./aoc submit <level> <input> <day> <year>")
	fmt.Println("Defaults:")
	fmt.Println("level: 1")
	fmt.Println("input: test.txt")
	fmt.Println("day:   current year")
	fmt.Println("year:  current day")
}

func (s *Submit) Run(args []string) error {
	args, ok := s.GetArgs(args)
	if !ok {
		return nil
	}
	level, input, day, year := args[0], args[1], args[2], args[3]

	ans, err := getAnswer(input, level, year, day)
	if err != nil {
		return err
	}

	resp, err := submitAnswer(ans, level, year, day)
	if err != nil {
		return err
	}

	if err = createSubmissionFile(resp, year, day); err != nil {
		return err
	}

	path := fmt.Sprintf("./solutions/%s/day%s/submission.html", year, day)
	if err = utils.OpenFile(path); err != nil {
		return err
	}

	return nil
}

func (s *Submit) GetArgs(args []string) ([]string, bool) {
	y, _, d := time.Now().Date()
	level, input, year, day := "1", "input.txt", fmt.Sprint(y), fmt.Sprint(d)
	for i, arg := range args {
		if arg == "-h" || arg == "--help" {
			s.PrintUsage()
			return []string{}, false
		}
		if i == 0 && (arg == "1" || arg == "2") {
			level = arg
		} else if i == 0 {
			fmt.Println("level must be 1 or 2")
			return []string{}, false
		} else if i == 1 {
			input = arg
		} else if i == 2 {
			day = arg
		} else if i == 3 {
			year = arg
		}
	}
	path := fmt.Sprintf("./solutions/%s/day%s", year, day)
	inputPath := path + "/" + input
	_, err := os.Stat(inputPath)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("Input file %s does not exist\n", inputPath)
		return []string{}, false
	} else if err != nil {
		fmt.Printf("Unexpected error with input file %s: %s\n", inputPath, err.Error())
		return []string{}, false
	}
	return []string{level, input, day, year}, true
}

func getAnswer(input, level, year, day string) (string, error) {
	cmd := exec.Command("go", "run", "main.go", input, level)
	dir := fmt.Sprintf("./solutions/%s/day%s/", year, day)
	cmd.Dir = dir
	output, err := cmd.Output()
	if err != nil {
		progPath := dir + "main.go"
		inputPath := dir + input
		fmt.Printf("There was an error running %s on level %s with input %s: %s\n", progPath, level, inputPath, err.Error())
		return "", err
	}
	ans := string(output)
	ans = strings.Replace(ans, " ", "", -1)
	ans = strings.Replace(ans, "\n", "", -1)
	ans = strings.Replace(ans, "\t", "", -1)
	return ans, err
}

func submitAnswer(ans, level, year, day string) ([]byte, error) {
	form := url.Values{}
	form.Add("level", level)
	form.Add("answer", ans)
	fmt.Printf("Submitting %s for day %s %s\n", ans, day, year)
	uri := fmt.Sprintf("https://adventofcode.com/%s/day/%s/answer", year, day)
	body, err := utils.PostAoC(uri, form)
	if err != nil {
		fmt.Println("There was an error submitting the answer")
		return nil, err
	}
	return body, nil
}

func createSubmissionFile(resp []byte, year, day string) error {
	submission := "submission.html"
	path := fmt.Sprintf("./solutions/%s/day%s", year, day)
	resp = utils.AddCss(resp)
	fileName := fmt.Sprintf("%s/%s", path, submission)
	if err := utils.WriteFile(name, resp); err != nil {
		fmt.Printf("Error creating %s\n", fileName)
		return err
	}
	fmt.Printf("Created %s/%s\n", path, submission)
	return nil
}
