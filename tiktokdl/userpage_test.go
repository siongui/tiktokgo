package tiktokdl

import (
	"testing"
)

func TestDownloadUserPageNextData(t *testing.T) {
	if err := DownloadUserPageNextData("tiktok"); err != nil {
		t.Error(err)
	}
}
