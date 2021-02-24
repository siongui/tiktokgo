package tiktokgo

// This file sends HTTP requests and gets HTTP responses.

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

var userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36"

// SetUserAgent sets User-Agent header in HTTP requests.
func SetUserAgent(s string) {
	userAgent = s
}

func SendHttpRequest(url, method string, cookies, headers map[string]string) (b []byte, err error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return
	}

	for name, value := range cookies {
		req.AddCookie(&http.Cookie{Name: name, Value: value})
	}

	req.Header.Set("User-Agent", userAgent)
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
