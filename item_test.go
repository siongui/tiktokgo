package tiktokgo

import (
	"testing"
)

func TestGetWebItemPageNextData(t *testing.T) {
	nd, err := GetWebItemPageNextData("tiktok", "6932896546379861253")
	if err != nil {
		t.Error(err)
		return
	}

	if nd.Props.PageProps.ServerCode != 200 {
		t.Error(nd.Props.PageProps.ServerCode)
		return
	}

	item := nd.Props.PageProps.ItemInfo.ItemStruct
	t.Log(item)
}
