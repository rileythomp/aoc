package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

func getInput(year, day string) []byte {
	uri := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, day)
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
	args := os.Args[1:]
	year, day := args[0], args[1]

	input := getInput(year, day)
	boilerplate := getBoilerplate()

	path := fmt.Sprintf("./%s/day%s", year, day)
	_ = os.MkdirAll(path, os.ModePerm)
	_ = os.WriteFile(fmt.Sprintf("%s/%s", path, "test.txt"), []byte{}, os.ModePerm)
	_ = os.WriteFile(fmt.Sprintf("%s/%s", path, "input.txt"), input, os.ModePerm)
	_ = os.WriteFile(fmt.Sprintf("%s/%s", path, "main.go"), boilerplate, os.ModePerm)
}
