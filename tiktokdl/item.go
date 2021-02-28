package tiktokdl

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/siongui/tiktokgo"
)

// DownloadItem downloads user's tiktok video.
func DownloadItem(item tiktokgo.TiktokItem) (err error) {
	cookies, err := tiktokgo.GetCookies()
	if err != nil {
		return
	}

	err = CreateDirIfNotExist(UserDir(item.Author))
	if err != nil {
		return
	}

	itempath := UserItemFilePath(item)
	if _, err := os.Stat(itempath); os.IsNotExist(err) {
		return SaveItemBytes(item.Video.PlayAddr, itempath, tiktokgo.GetHeaders(), cookies)
		//println(item.Video.PlayAddr)
		//println(item.Video.DownloadAddr)
		//println(itempath)
	} else {
		if err != nil {
			return err
		}
	}
	return
}

// DownloadItemByUsernameVideoId downloads the video item given the username
// (@ not included) and video id in the item URL.
func DownloadItemByUsernameVideoId(username, videoId string) (err error) {
	nd, err := tiktokgo.GetWebItemPageNextData(username, videoId)
	if err != nil {
		return
	}

	if nd.Props.PageProps.ServerCode != 200 {
		err = errors.New("response serverCode not 200: " + strconv.FormatInt(nd.Props.PageProps.ServerCode, 10))
		return
	}

	item := nd.Props.PageProps.ItemInfo.ItemStruct
	return DownloadItem(item)
}

// DownloadItemByUrl downloads video item given the the URL of the video item.
func DownloadItemByUrl(itemUrl string) (err error) {
	urlnoquery := StripQueryString(itemUrl)
	ss := strings.Split(urlnoquery, "/")
	if len(ss) != 6 {
		err = errors.New("video item URL format seems not correct: " + itemUrl)
		return
	}
	if ss[0] != "https:" {
		err = errors.New("video item URL not https: " + itemUrl)
		return
	}
	if !strings.HasPrefix(ss[3], "@") {
		err = errors.New("video item URL has no @ character: " + itemUrl)
		return
	}
	if ss[4] != "video" {
		err = errors.New("video item URL has the string 'video': " + itemUrl)
		return
	}

	return DownloadItemByUsernameVideoId(strings.TrimPrefix(ss[3], "@"), ss[5])
}
