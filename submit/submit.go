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

type SubmitArgs struct {
	Level string
	Input string
	Year  string
	Day   string
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("./aoc submit <level> <input> <day> <year>")
	fmt.Println("Defaults:")
	fmt.Println("level: 1")
	fmt.Println("input: test.txt")
	fmt.Println("day:  current year")
	fmt.Println("year:  current day")
}

func (s *Submit) Run(args []string) {
	sa, ok := getArgs(args)
	if !ok {
		return
	}

	ans, err := getAnswer(sa.Input, sa.Level, sa.Year, sa.Day)
	if err != nil {
		return
	}

	resp, err := submitAnswer(ans, sa.Level, sa.Year, sa.Day)
	if err != nil {
		return
	}

	err = createSubmissionFile(resp, sa.Year, sa.Day)
	if err != nil {
		return
	}

	_ = openResult(sa.Year, sa.Day)
}

func getArgs(args []string) (SubmitArgs, bool) {
	y, _, d := time.Now().Date()
	sa := SubmitArgs{
		Level: "1",
		Input: "input.txt",
		Year:  fmt.Sprint(y),
		Day:   fmt.Sprint(d),
	}
	for i, arg := range args {
		if arg == "-h" || arg == "--help" {
			printUsage()
			return SubmitArgs{}, false
		}
		if i == 0 && (arg == "1" || arg == "2") {
			sa.Level = arg
		} else if i == 0 {
			fmt.Println("level must be 1 or 2")
			return SubmitArgs{}, false
		} else if i == 1 {
			sa.Input = arg
		} else if i == 2 {
			sa.Day = arg
		} else if i == 3 {
			sa.Year = arg
		}
	}
	path := fmt.Sprintf("./%s/day%s", sa.Year, sa.Day)
	inputPath := path + "/" + sa.Input
	_, err := os.Stat(inputPath)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("Input file %s does not exist\n", inputPath)
		return SubmitArgs{}, false
	} else if err != nil {
		fmt.Printf("Unexpected error with input file %s: %s\n", inputPath, err.Error())
		return SubmitArgs{}, false
	}
	return sa, true
}

func getAnswer(input, level, year, day string) (string, error) {
	cmd := exec.Command("go", "run", "main.go", input, level)
	dir := fmt.Sprintf("./%s/day%s/", year, day)
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
	var err error
	form := url.Values{}
	form.Add("level", level)
	form.Add("answer", ans)
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
	path := fmt.Sprintf("./%s/day%s", year, day)
	resp = utils.AddCss(resp)
	err := os.WriteFile(fmt.Sprintf("%s/%s", path, submission), resp, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating %s/%s\n", path, submission)
		return err
	}
	fmt.Printf("Created %s/%s\n", path, submission)
	return nil
}

func openResult(year, day string) error {
	filePath := fmt.Sprintf("./%s/day%s/submission.html", year, day)
	cmd := exec.Command("open", filePath)
	_, err := cmd.Output()
	if err != nil {
		fmt.Printf("There was an error opening %s: %s\n", filePath, err.Error())
		return err
	}
	return nil
}
