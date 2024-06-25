package v1

type GenerateShortUrlReq struct {
	Url string `json:"url"` // 网址
}

type GetLongUrlReq struct {
	ShortUrl string `json:"url"`
}
