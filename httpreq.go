package tiktokgo

// This file sends HTTP requests and gets HTTP responses.

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var defaultCookies = make(map[string]string)
var defaultCookiesFilePath = "cookies.json"

// SetCookiesFilePath sets the file path of JSON files which stores the cookies.
func SetCookiesFilePath(s string) {
	defaultCookiesFilePath = s
}

// LoadCookies loads cookies stored in JSON format from file.
func LoadCookies(cookieFilePath string) (err error) {
	b, err := ioutil.ReadFile(cookieFilePath)
	if err != nil {
		return
	}

	return json.Unmarshal(b, &defaultCookies)
}

// SaveCookies saves cookies to file in JSON format.
func SaveCookies() error {
	b, err := json.Marshal(defaultCookies)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(defaultCookiesFilePath, b, 0644)
}

// GetCookies returns HTTP request cookies.
func GetCookies() (map[string]string, error) {
	if len(defaultCookies) > 0 {
		return defaultCookies, nil
	}
	err := LoadCookies(defaultCookiesFilePath)
	if os.IsNotExist(err) {
		log.Println("default cookies file not exists. try to get first-time cookies.")
		defaultCookies, err = GetFirstTimeCookies()
		if err == nil {
			err = SaveCookies()
		}
	}
	return defaultCookies, err
}

// FIXME: ok to get cookies, but cannot be used to get correct video addr
func GetFirstTimeCookies() (c map[string]string, err error) {
	url := "https://www.tiktok.com/"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	for name, value := range defaultHeaders {
		req.Header.Set(name, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(url +
			"\nresp.StatusCode: " + strconv.Itoa(resp.StatusCode))
		return
	}

	c = make(map[string]string)
	// Ref: https://stackoverflow.com/a/53023010
	// Ref: https://gist.github.com/rowland/984989
	for _, cookie := range resp.Cookies() {
		c[cookie.Name] = cookie.Value
	}
	return
}

var defaultHeaders = map[string]string{
	"Accept":     "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36",
}

// GetHeaders returns HTTP request headers.
func GetHeaders() map[string]string {
	return defaultHeaders
}

// SetUserAgent sets User-Agent header in HTTP requests.
func SetUserAgent(s string) {
	defaultHeaders["User-Agent"] = s
}

// SendHttpRequest sends HTTP request.
func SendHttpRequest(url, method string, cookies, headers map[string]string) (b []byte, err error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return
	}

	for name, value := range cookies {
		req.AddCookie(&http.Cookie{Name: name, Value: value})
	}

	for name, value := range headers {
		req.Header.Set(name, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(url +
			"\nresp.StatusCode: " + strconv.Itoa(resp.StatusCode))
		return
	}

	return ioutil.ReadAll(resp.Body)
}
