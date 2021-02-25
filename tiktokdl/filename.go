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

func UserFilenameString(user tiktokgo.TiktokUser) string {
	return user.UniqueId + "-" + user.Id + "-" + user.Nickname
}

func UserAvatarFilePath(user tiktokgo.TiktokUser) string {
	filename := UserFilenameString(user) + "-tiktok-avatar" + filepath.Ext(StripQueryString(user.AvatarLarger))
	return filepath.Join(UserDir(user), filename)
}

func UserItemFilePath(item tiktokgo.TiktokItem) string {
	user := item.Author
	filename := UserFilenameString(user) + "-tiktok-video-" + GetRFC3339Time(item.CreateTime) + ".mp4"
	return filepath.Join(UserDir(user), filename)
}
