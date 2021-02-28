package tiktokgo

type TiktokUser struct {
	Id           string `json:"id"`
	ShortId      string `json:"shortId"`
	UniqueId     string `json:"uniqueId"`
	Nickname     string `json:"nickname"`
	AvatarLarger string `json:"avatarLarger"`
	AvatarMedium string `json:"avatarMedium"`
	AvatarThumb  string `json:"avatarThumb"`
	Signature    string `json:"signature"`
	CreateTime   int64  `json:"createTime"`
	Verified     bool   `json:"verified"`
	SecUid       string `json:"secUid"`
	Ftc          bool   `json:"ftc"`
	//TODO: add more properties
	PrivateAccount bool   `json:"privateAccount"`
	Secret         bool   `json:"secret"`
	roomId         string `json:"roomId"`
}

type TiktokVideo struct {
	Id           string   `json:"id"`
	Height       int64    `json:"height"`
	Width        int64    `json:"width"`
	Duration     int64    `json:"duration"`
	Ratio        string   `json:"ratio"`
	Cover        string   `json:"cover"`
	OriginCover  string   `json:"originCover"`
	DynamicCover string   `json:"dynamicCover"`
	PlayAddr     string   `json:"playAddr"`
	DownloadAddr string   `json:"downloadAddr"`
	ShareCover   []string `json:"shareCover"`
	ReflowCover  string   `json:"reflowCover"`
}

type TiktokUserInfo struct {
	User  TiktokUser `json:"user"`
	Stats struct {
		FollowerCount  int64 `json:"followerCount"`
		FollowingCount int64 `json:"followingCount"`
		Heart          int64 `json:"heart"`
		HeartCount     int64 `json:"heartCount"`
		VideoCount     int64 `json:"videoCount"`
		DiggCount      int64 `json:"diggCount"`
	} `json:"stats"`
}

type TiktokItem struct {
	Id           string      `json:"id"`
	Desc         string      `json:"desc"`
	CreateTime   int64       `json:"createTime"`
	ScheduleTime int64       `json:"scheduleTime"`
	Video        TiktokVideo `json:"video"`
	Author       TiktokUser  `json:"author"`
	//TODO: add more properties
	IsAd         bool `json:"isAd"`
	ShareEnabled bool `json:"shareEnabled"`
}
