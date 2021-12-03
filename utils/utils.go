package utils

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

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
