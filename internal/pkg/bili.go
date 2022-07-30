package pkg

// https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/user/space.md#%E6%8A%95%E7%A8%BF

// URLPrefix bilibili video prefix
const URLPrefix string = "https://www.bilibili.com/video/"

// RequestURL api request url
const RequestURL string = "https://api.bilibili.com/x/space/arc/search"

type biliResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	TTL     int              `json:"ttl"`
	Data    biliResponseData `json:"data"`
}

type biliResponseData struct {
	List           dataList           `json:"list"`
	Page           dataPage           `json:"page"`
	EpisodicButton dataEpisodicButton `json:"episodic_button"`
}

type dataList struct {
	TList tList       `json:"tlist"`
	VList []videoData `json:"vlist"`
}

type tList struct {
	// ignore this
}

type videoData struct {
	Comment        int    `json:"comment"`
	TypeID         int    `json:"typeid"`
	Play           int    `json:"play"`
	Pic            string `json:"pic"`
	SubTitle       string `json:"subtitle"`
	Description    string `json:"description"`
	Copyright      string `json:"copyright"`
	Title          string `json:"title"`
	Review         int    `json:"review"`
	Author         string `json:"author"`
	Mid            int64  `json:"mid"`
	Created        int64  `json:"created"`
	Length         string `json:"length"`
	VideoReview    int    `json:"video_review"`
	AID            int64  `json:"aid"`
	BVID           string `json:"bvid"`
	HideClick      bool   `json:"hide_click"`
	IsPay          int    `json:"is_pay"`
	IsUnionVideo   int    `json:"is_union_video"`
	IsSteinsGate   int    `json:"is_steins_gate"`
	IsLivePlayback int    `json:"is_live_playback"`
}

type dataPage struct {
	PageNumber int `json:"pn"`
	PageSize   int `json:"ps"`
	Count      int `json:"count"`
}

type dataEpisodicButton struct {
	Text string `json:"text"`
	URI  string `json:"uri"`
}

// VideoRecord video record
type VideoRecord struct {
	URL     string
	Pic     string
	Title   string
	Author  string
	Created int64
}
