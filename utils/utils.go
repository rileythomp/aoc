package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func AddCss(html []byte) []byte {
	return []byte(strings.Replace(
		string(html),
		"/static/style.css?26",
		"https://adventofcode.com/static/style.css",
		-1,
	))
}

func GetAoC(uri string) ([]byte, error) {
	urlObj, err := url.ParseRequestURI(uri)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	client.Jar, err = cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client.Jar.SetCookies(urlObj, []*http.Cookie{
		{Name: "session", Value: os.Getenv("AOC_TOKEN")},
	})
	resp, err := client.Get(uri)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return body, nil
}

func PostAoC(uri string, form url.Values) ([]byte, error) {
	var (
		client = &http.Client{}
		err    error
	)
	client.Jar, err = cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	urlObj, err := url.ParseRequestURI(uri)
	if err != nil {
		return nil, err
	}
	client.Jar.SetCookies(urlObj, []*http.Cookie{
		{Name: "session", Value: os.Getenv("AOC_TOKEN")},
	})
	req, err := http.NewRequest("POST", uri, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
	if err != nil {
		return nil, err
	}
	return body, nil
}

func WriteFileBytes(name string, content []byte) error {
	if err := os.WriteFile(name, content, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func WriteFileString(name, str string) error {
	f, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(str); err != nil {
		return err
	}
	return nil
}

func GetFile(path string) ([]byte, error) {
	boilerplate, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return boilerplate, nil
}

func OpenFile(path string) error {
	cmd := exec.Command("open", path)
	if _, err := cmd.Output(); err != nil {
		fmt.Printf("There was an error opening %s: %s\n", path, err.Error())
		return err
	}
	return nil
}
