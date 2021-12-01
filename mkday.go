package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"time"
)

func main() {
	year, day := getYearAndDay()
	if year != "" && day != "" {
		_ = createFiles(year, day)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("go run mkday.go <day> <year>")
	fmt.Println("Defaults:")
	fmt.Println("day:  next day")
	fmt.Println("year: current year")
	fmt.Println("So if no arguments are passed, it will wait until the next puzzle is released at midnight")
}

func getYearAndDay() (string, string) {
	args := os.Args[1:]
	y, _, d := time.Now().Date()
	var (
		year = fmt.Sprint(y)
		day  = fmt.Sprint(d + 1)
	)
	for i, arg := range args {
		if arg == "-h" || arg == "--help" {
			printUsage()
			return "", ""
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
	return year, day
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
	files := []struct {
		Name    string
		Content []byte
	}{
		{Name: "problem.html", Content: getAoC(year, day, uri)},
		{Name: "input.txt", Content: getAoC(year, day, uri+"/input")},
		{Name: "test.txt", Content: []byte{}},
		{Name: "main.go", Content: getBoilerplate()},
	}
	for _, file := range files {
		err = os.WriteFile(fmt.Sprintf("%s/%s", path, file.Name), file.Content, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating %s/%s\n", path, file.Name)
			return err
		}
		fmt.Printf("Created %s/%s\n", path, file.Name)
	}
	return nil
}

func getAoC(year, day, uri string) []byte {
	urlObj, _ := url.ParseRequestURI(uri)
	client := &http.Client{}
	client.Jar, _ = cookiejar.New(nil)
	client.Jar.SetCookies(urlObj, []*http.Cookie{
		{Name: "session", Value: os.Getenv("AOC_COOKIE")},
	})
	resp, _ := client.Get(uri)
	body, _ := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
	defer resp.Body.Close()
	return body
}

func getBoilerplate() []byte {
	boilerplate, _ := ioutil.ReadFile("./boilerplate.go")
	return boilerplate
}
