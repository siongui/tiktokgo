package main

import (
	"github.com/siongui/tiktokgo/tiktokdl"
)

func main() {
	// Example to download user avatar photo and latest 5 video items (with watermark).
	err := tiktokdl.DownloadUserPageNextData("tiktok")
	if err != nil {
		panic(err)
	}
}
