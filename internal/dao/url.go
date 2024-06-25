package dao

import (
	"short-url-go/internal/global"
	"short-url-go/internal/model"
)

func InsertShortUrl(urlObj model.ShortUrl) error {
	return global.G_DB.Create(&urlObj).Error
}

func GetLongUrl(shortUrl string) (result string, err error) {
	url := model.ShortUrl{}
	err = global.G_DB.Select("lurl").Where("surl = ?", shortUrl).First(&url).Error
	result = url.Lurl
	return result, err
}
