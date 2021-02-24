package tiktokgo

import (
	"testing"
)

func TestGetWebUserPageNextData(t *testing.T) {
	nd, err := GetWebUserPageNextData("tiktok")
	if err != nil {
		t.Error(err)
	}

	userinfo := nd.Props.PageProps.UserInfo
	if userinfo.User.Id != "107955" {
		t.Error(userinfo.User.Id)
	}

	items := nd.Props.PageProps.Items
	t.Log(userinfo)
	for _, item := range items {
		t.Log(item)
	}
}
