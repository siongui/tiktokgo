package tiktokdl

import (
	"net/url"
	"path/filepath"

	"github.com/siongui/tiktokgo"
)

var DefaultDownloadDir = "TikTok"

func StripQueryString(inputUrl string) string {
	u, err := url.Parse(inputUrl)
	if err != nil {
		panic(err)
	}
	u.RawQuery = ""
	return u.String()
}

func UserDir(user tiktokgo.TiktokUser) string {
	return filepath.Join(DefaultDownloadDir, user.UniqueId)
}

func UserAvatarFilePath(user tiktokgo.TiktokUser) string {
	filename := user.UniqueId + "-" + user.Id + "-" + user.Nickname + "-tiktok-avatar" + filepath.Ext(StripQueryString(user.AvatarLarger))
	return filepath.Join(UserDir(user), filename)
}
