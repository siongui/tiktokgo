package tiktokdl

import (
	"net/url"
	"path/filepath"

	"github.com/siongui/tiktokgo"
)

var defaultDownloadDir = "TikTok"

// StripQueryString removes the query string in the URL.
func StripQueryString(inputUrl string) string {
	u, err := url.Parse(inputUrl)
	if err != nil {
		panic(err)
	}
	u.RawQuery = ""
	return u.String()
}

// SetDefaultDownloadDir sets the directory which saves the downloaded files.
func SetDefaultDownloadDir(s string) {
	defaultDownloadDir = s
}

// UserDir returns default directory path for given user.
func UserDir(user tiktokgo.TiktokUser) string {
	return filepath.Join(defaultDownloadDir, user.UniqueId)
}

// UserFilenameString returns user string which is used in the file name of
// downloaded files.
func UserFilenameString(user tiktokgo.TiktokUser) string {
	return user.UniqueId + "-" + user.Id + "-" + user.Nickname
}

// UserAvatarFilePath returns the file path of saved file of user avatar photo.
func UserAvatarFilePath(user tiktokgo.TiktokUser) string {
	filename := UserFilenameString(user) + "-tiktok-avatar" + filepath.Ext(StripQueryString(user.AvatarLarger))
	return filepath.Join(UserDir(user), filename)
}

// UserItemFilePath returns the file path of saved file of video item.
func UserItemFilePath(item tiktokgo.TiktokItem) string {
	user := item.Author
	filename := UserFilenameString(user) + "-tiktok-video-" + GetRFC3339Time(item.CreateTime) + ".mp4"
	return filepath.Join(UserDir(user), filename)
}
