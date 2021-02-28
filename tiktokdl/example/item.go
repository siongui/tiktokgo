package main

import (
	"github.com/siongui/tiktokgo/tiktokdl"
)

func main() {
	// Example to download video item (with watermark) by URL.
	err := tiktokdl.DownloadItemByUrl("https://www.tiktok.com/@tiktok/video/6878464336357084422?lang=en&is_copy_url=1&is_from_webapp=v2")
	if err != nil {
		panic(err)
	}
}
