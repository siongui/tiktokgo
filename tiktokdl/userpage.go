package tiktokdl

import (
	"github.com/siongui/tiktokgo"
)

// DownloadUserPageNextData downloads larger user profile pic and video items in
// the __NEXT_DATA__ of user page HTML.
func DownloadUserPageNextData(username string) (err error) {
	nd, err := tiktokgo.GetWebUserPageNextData(username)
	if err != nil {
		return
	}

	userinfo := nd.Props.PageProps.UserInfo

	err = CreateDirIfNotExist(UserDir(userinfo.User))
	if err != nil {
		return
	}

	err = Wget(userinfo.User.AvatarLarger, UserAvatarFilePath(userinfo.User))
	if err != nil {
		return
	}

	//TODO: download items in __NEXT_DATA__

	return
}
