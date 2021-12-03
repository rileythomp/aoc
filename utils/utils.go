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
	return []byte(strings.Replace(string(html), "/static/style.css?26", "https://adventofcode.com/static/style.css", 1))
}

func GetAoC(uri string) ([]byte, error) {
	urlObj, err := url.ParseRequestURI(uri)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	client.Jar, _ = cookiejar.New(nil)
	client.Jar.SetCookies(urlObj, []*http.Cookie{
		{Name: "session", Value: os.Getenv("AOC_COOKIE")},
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
	var err error
	client := &http.Client{}
	client.Jar, err = cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	urlObj, err := url.ParseRequestURI(uri)
	if err != nil {
		return nil, err
	}
	client.Jar.SetCookies(urlObj, []*http.Cookie{
		{Name: "session", Value: os.Getenv("AOC_COOKIE")},
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

func GetFile(path string) ([]byte, error) {
	boilerplate, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return boilerplate, nil
}

func OpenFile(path string) error {
	cmd := exec.Command("open", path)
	_, err := cmd.Output()
	if err != nil {
		fmt.Printf("There was an error opening %s: %s\n", path, err.Error())
		return err
	}
	return nil
}
