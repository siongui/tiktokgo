package tiktokgo

import (
	"testing"
)

func TestItemList(t *testing.T) {
	items, err := ItemList("tiktok")
	if err != nil {
		t.Error(err)
		return
	}

	for _, item := range items {
		t.Log(item.Id)
		t.Log(item.Video.PlayAddr)
	}
}
