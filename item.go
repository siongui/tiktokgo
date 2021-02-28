package tiktokgo

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
)

// ItemPageNextData is the JSON data embedded in the script tag with
// id=__NEXT_DATA__ in the item page.
type ItemPageNextData struct {
	Props struct {
		PageProps struct {
			ServerCode int64 `json:"serverCode"`
			ItemInfo   struct {
				ItemStruct TiktokItem `json:"itemStruct"`
			} `json:"itemInfo"`
		} `json:"pageProps"`
	} `json:"props"`
}

// GetWebItemPageNextData returns JSON data embedded in the HTML of item page.
// The parameters username and id is the name and item id embedded in the URL.
// For example, the following URL
//
//   https://www.tiktok.com/@tiktok/video/6932896546379861253
//
// The username is tiktok (@ not included).
// The id is 6932896546379861253.
func GetWebItemPageNextData(username, id string) (nd ItemPageNextData, err error) {
	url := "https://www.tiktok.com/@" + username + "/video/" + id

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
