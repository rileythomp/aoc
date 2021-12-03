package submit

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"
)

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

func RunSubmit(args []string) {
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
	client := &http.Client{}
	client.Jar, err = cookiejar.New(nil)
	if err != nil {
		fmt.Printf("Error creating cookies: %s\n", err)
		return nil, err
	}
	uri := fmt.Sprintf("https://adventofcode.com/%s/day/%s/answer", year, day)
	urlObj, err := url.ParseRequestURI(uri)
	if err != nil {
		fmt.Printf("There was an parsing the uri %s: %s\n", uri, err.Error())
		return nil, err
	}
	client.Jar.SetCookies(urlObj, []*http.Cookie{
		{Name: "session", Value: os.Getenv("AOC_COOKIE")},
	})
	req, err := http.NewRequest("POST", uri, strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Printf("There was an error creating the POST request for submission to the uri %s: %s\n", uri, err.Error())
		return nil, err
	}
	fmt.Printf("Submitting %s for year %s day %s level %s\n", ans, year, day, level)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("There was an error submitting the answer %s for year %s day %s: %s\n", ans, year, day, err.Error())
		return nil, err
	}
	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
	if err != nil {
		fmt.Printf("There was an error reading the submission response: %s\n", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	return body, nil
}

func createSubmissionFile(resp []byte, year, day string) error {
	submission := "submission.html"
	path := fmt.Sprintf("./%s/day%s", year, day)
	resp = []byte(strings.Replace(string(resp), "/static/style.css?26", "https://adventofcode.com/static/style.css", 1))
	err := os.WriteFile(fmt.Sprintf("%s/%s", path, submission), resp, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating %s/%s\n", path, submission)
		return err
	}
	fmt.Printf("Created %s/%s\n", path, submission)
	return nil
}
