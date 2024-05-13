package common

type URL struct {
	OriginalUrl string `json:"originalUrl"`
	ShortenUrl  string `json:"shortenUrl"`
}

var Urls = make(map[string]string)
