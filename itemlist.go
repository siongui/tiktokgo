package tiktokgo

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

// TiktokItemList used to decode JSON data for TikTok itemList.
type TiktokItemList struct {
	StatusCode int64        `json:"statusCode"`
	ItemList   []TiktokItem `json:"itemList"`
	Cursor     string       `json:"cursor"`
	HasMore    bool         `json:"hasMore"`
}

// GetItemListQueryString returns query string for item list api endpoint.
func GetItemListQueryString(userinfo TiktokUserInfo) url.Values {
	v := url.Values{}

	v.Add("aid", "1988")
	v.Add("app_name", "tiktok_web")
	v.Add("device_platform", "web")
	v.Add("referer", "")
	v.Add("root_referer", "")
	v.Add("user_agent", GetHeaders()["User-Agent"])
	v.Add("cookie_enabled", "true")
	v.Add("screen_width", "1920")
	v.Add("screen_height", "1080")
	v.Add("browser_language", "en-US")
	v.Add("browser_platform", "Linux x86_64")
	v.Add("browser_name", "Mozilla")
	// TODO: non-empty browser_version
	v.Add("browser_version", "")
	v.Add("browser_online", "true")
	v.Add("ac", "4g")
	v.Add("timezone_name", "Asia/Taipei")
	v.Add("priority_region", "")
	v.Add("appId", "1180")
	v.Add("region", "TW")
	v.Add("appType", "t")
	v.Add("isAndroid", "false")
	v.Add("isMobile", "false")
	v.Add("isIOS", "false")
	v.Add("OS", "linux")
	// TODO: non-empty did
	v.Add("did", "")
	v.Add("count", "30")
	// TODO: non-empty cursor
	v.Add("cursor", "")
	v.Add("language", "en")
	//v.Add("", "")
	v.Add("id", userinfo.User.Id)
	v.Add("secUid", userinfo.User.SecUid)
	return v
}

// ItemList returns item list of a given user.
func ItemList(username string) (items []TiktokItem, err error) {
	nd, err := GetWebUserPageNextData(username)
	if err != nil {
		return
	}

	if nd.Props.PageProps.ServerCode != 200 {
		err = errors.New("nd.Props.PageProps.ServerCode: " + strconv.FormatInt(nd.Props.PageProps.ServerCode, 10))
		return
	}

	userinfo := nd.Props.PageProps.UserInfo
	v := GetItemListQueryString(userinfo)

	//url := "https://m.tiktok.com/api/post/item_list/?" + v.Encode()
	url := "https://t.tiktok.com/api/post/item_list/?" + v.Encode()
	//println(url)

	cookies, err := GetCookies()
	if err != nil {
		return
	}
	b, err := SendHttpRequest(url, http.MethodGet, cookies, GetHeaders())
	if err != nil {
		return
	}

	//println(string(b))
	il := TiktokItemList{}
	err = json.Unmarshal(b, &il)
	if err != nil {
		return
	}

	if il.StatusCode != 0 {
		err = errors.New("il.StatusCode: " + strconv.FormatInt(il.StatusCode, 10))
		return
	}

	// TODO: check hasMore
	items = il.ItemList
	return
}
