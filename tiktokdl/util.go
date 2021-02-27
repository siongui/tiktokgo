package tiktokdl

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

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

// Wget calls shell command *wget* to download. The reason to use wget is that
// wget supports automatically resume download and timeout.
func Wget(url, filepath string) error {
	// run shell command `wget URL -O filepath -T 7`
	cmd := exec.Command("wget", url, "-O", filepath, "-T", "7")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// WgetCookies calls wget with cookies header. FIXME: not working
func WgetCookies(url, filepath string, headers, cookies map[string]string) error {
	if len(cookies) == 0 {
		return Wget(url, filepath)
	}

	// Ref: https://superuser.com/a/854094
	// Ref: https://stackoverflow.com/a/7618268
	var opts []string
	opts = append(opts, url)
	opts = append(opts, "-O")
	opts = append(opts, filepath)
	opts = append(opts, "-T")
	opts = append(opts, "7")
	//opts = append(opts, "--no-cookies")
	str := `--header="Cookie: `
	for name, value := range cookies {
		str = str + name + "=" + value + "; "
	}
	str = str[:len(str)-2] + `"`
	println(str)
	opts = append(opts, str)

	for name, value := range headers {
		s := `--header="` + name + ": " + value + `"`
		println(s)
		opts = append(opts, s)
	}

	cmd := exec.Command("wget", opts...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// IsCommandAvailable checks if the command is available.
func IsCommandAvailable(name string) bool {
	cmd := exec.Command("/bin/sh", "-c", "command -v "+name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// CreateDirIfNotExist creates dir if not exist.
func CreateDirIfNotExist(dir string) (err error) {
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
	}
	return
}
