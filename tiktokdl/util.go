package tiktokdl

import (
	"os"
	"os/exec"
)

// Wget calls shell command *wget* to download. The reason to use wget is that
// wget supports automatically resume download and timeout.
func Wget(url, filepath string) error {
	// run shell command `wget URL -O filepath -T 7`
	cmd := exec.Command("wget", url, "-O", filepath, "-T", "7")
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
