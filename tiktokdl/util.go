package tiktokdl

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// SaveItemBytes makes HTTP GET request to the given URL and saves the content
// to local file.
func SaveItemBytes(url, localFilePath string, headers, cookies map[string]string) (err error) {
	log.Println("Downloading ", url, " to ", localFilePath)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	for name, value := range cookies {
		req.AddCookie(&http.Cookie{Name: name, Value: value})
	}

	// From https://github.com/drawrowfly/tiktok-scraper
	//   - *User-Agent* header
	//   - *referer* header
	//   - *tt_webid_v2* cookie
	// must be the same both when accessing API and when downloading video.
	// Otherwise HTTP 403 status code will be returned when downloading
	// video.
	for name, value := range headers {
		req.Header.Set(name, value)
	}
	// without the following referer header, HTTP response status code will
	// be 403 when we try to download video.
	req.Header.Set("referer", "https://www.tiktok.com/")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	log.Println("HTTP response status code: ", resp.StatusCode)

	if resp.StatusCode != 200 {
		err = errors.New(url +
			"\nresp.StatusCode: " + strconv.Itoa(resp.StatusCode))
		return
	}

	f, err := os.Create(localFilePath)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return
}

// CreateDirIfNotExist creates dir if not exist.
func CreateDirIfNotExist(dir string) (err error) {
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
	}
	return
}
