package tiktokgo

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
)

type UserPageNextData struct {
	Props struct {
		PageProps struct {
			ServerCode         int64          `json:"serverCode"`
			UserInfo           TiktokUserInfo `json:"userInfo"`
			Items              []TiktokItem   `json:"items"`
			VideoListHasMore   bool           `json:"videoListHasMore"`
			VideoListMaxCursor int64          `json:"videoListMaxCursor"`
			VideoListMinCursor int64          `json:"videoListMinCursor"`
		} `json:"pageProps"`
	} `json:"props"`
}

// GetWebUserPageNextData returns JSON data embedded in the HTML of user page.
// The parameter username is the name embedded in the URL. For example, the
// following URL
//
//   https://www.tiktok.com/@tiktok
//
// The username is tiktok (@ not included).
func GetWebUserPageNextData(username string) (nd UserPageNextData, err error) {
	//url := "https://www.tiktok.com/@" + username + "?"
	url := "https://www.tiktok.com/@" + username

	cookies, err := GetCookies()
	if err != nil {
		return
	}
	b, err := SendHttpRequest(url, http.MethodGet, cookies, GetHeaders())
	if err != nil {
		return
	}

	//println(string(b))

	pattern := regexp.MustCompile(`<script id="__NEXT_DATA__" type="application\/json" crossorigin="anonymous">(.*?)<\/script>`)
	matches := pattern.FindSubmatch(b)
	//println(len(matches))
	if len(matches) != 2 {
		err = errors.New("trouble getting __NEXT_DATA__")
		return
	}

	//println(string(matches[1]))

	err = json.Unmarshal(matches[1], &nd)
	return
}
