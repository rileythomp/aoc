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

func getInput(year, day int) []byte {
	uri := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	urlObj, _ := url.ParseRequestURI(uri)
	client := &http.Client{}
	client.Jar, _ = cookiejar.New(nil)
	client.Jar.SetCookies(urlObj, []*http.Cookie{
		{
			Name:  "session",
			Value: os.Getenv("AOC_COOKIE"),
		},
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

func main() {
	year, _, initDay := time.Now().Date()
	day := initDay
	for day == initDay {
		time.Sleep(time.Second)
		_, _, day = time.Now().Date()
	}

	input := getInput(year, day)
	boilerplate := getBoilerplate()

	path := fmt.Sprintf("./%d/day%d", year, day)
	_ = os.MkdirAll(path, os.ModePerm)
	_ = os.WriteFile(fmt.Sprintf("%s/%s", path, "test.txt"), []byte{}, os.ModePerm)
	_ = os.WriteFile(fmt.Sprintf("%s/%s", path, "input.txt"), input, os.ModePerm)
	_ = os.WriteFile(fmt.Sprintf("%s/%s", path, "main.go"), boilerplate, os.ModePerm)
}
