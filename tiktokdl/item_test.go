package tiktokdl

import (
	"testing"
)

func TestDownloadItemByUrl(t *testing.T) {
	if err := DownloadItemByUrl("https://www.tiktok.com/@tiktok/video/6878464336357084422?lang=en&is_copy_url=1&is_from_webapp=v2"); err != nil {
		t.Error(err)
	}
}
