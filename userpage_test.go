package tiktokgo

import (
	"testing"
)

func TestGetWebUserPageNextData(t *testing.T) {
	nd, err := GetWebUserPageNextData("tiktok")
	if err != nil {
		t.Error(err)
		return
	}

	if nd.Props.PageProps.ServerCode != 200 {
		t.Error(nd.Props.PageProps.ServerCode)
		return
	}

	userinfo := nd.Props.PageProps.UserInfo
	if userinfo.User.Id != "107955" {
		t.Error(userinfo.User.Id)
		return
	}

	items := nd.Props.PageProps.Items
	t.Log(userinfo)
	for _, item := range items {
		t.Log(item.Video.PlayAddr)
		t.Log(item.Video.DownloadAddr)
	}
}
