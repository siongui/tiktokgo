package tiktokdl

import (
	"os"
	"os/exec"
)

// Not used. Keep here for reference.
func wget(url, filepath string) error {
	// run shell command `wget URL -O filepath -T 7`
	cmd := exec.Command("wget", url, "-O", filepath, "-T", "7")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Not used. Keep here for reference.
func wgetCookies(url, filepath string, headers, cookies map[string]string) error {
	if len(cookies) == 0 {
		return wget(url, filepath)
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

// Not used. Keep here for reference.
// IsCommandAvailable checks if the command is available.
func isCommandAvailable(name string) bool {
	cmd := exec.Command("/bin/sh", "-c", "command -v "+name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
