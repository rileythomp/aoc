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
	_ = createFiles(year, day)
}

func getYearAndDay() (string, string) {
	args := os.Args[1:]
	var year, day string
	if len(args) > 1 {
		year, day = args[0], args[1]
	} else {
		fmt.Println("Waiting until problem is released at midnight...")
		curYear, _, initDay := time.Now().Date()
		curDay, seconds := initDay, 0
		for curDay == initDay {
			time.Sleep(time.Second)
			seconds++
			_, _, curDay = time.Now().Date()
		}
		year, day = fmt.Sprint(curYear), fmt.Sprint(curDay)
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
