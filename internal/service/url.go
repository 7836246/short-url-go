package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/orca-zhang/ecache"
	"github.com/twmb/murmur3"
	"log"
	"short-url-go/internal/dao"
	"short-url-go/internal/model"
	"short-url-go/internal/utils"
)

var (
	longCache = ecache.NewLRUCache(1000, 200)
)

// GenerateShortUrl 生成短网址的函数
func GenerateShortUrl(url string) (string, error) {
	// 输入参数验证
	if url == "" || len(url) > 450 {
		return "", fmt.Errorf("longUrl太长或不能为空")
	}
	// 尝试从缓存中获取长网址
	if longUrlInterface, ok := longCache.Get(url); ok {
		longUrl := longUrlInterface.(string) // 类型断言为 string
		// 在这里可以直接使用 longUrl，不需要再进行类型断言
		if longUrl != "" {
			return longUrl, nil
		}
	}
	return getShortUrl(url, getLongUrlRandom(url))
}

func GetLongUrl(shortUrl string) (string, error) {
	if shortUrl == "" {
		return "", fmt.Errorf("shortUrl 不能为空")
	}
	longUrl, err := dao.GetLongUrl(shortUrl)
	if err != nil {
		return "找不到对应的长链接", nil
	}
	longCache.Put(shortUrl, longUrl)
	return longUrl, nil
}

// getLongUrlRandom 获取随机的长网址
func getLongUrlRandom(longUrl string) string {
	return longUrl + utils.RandomString(6) // 解决冲突多的问题，随机字符串
}

func getShortUrl(rawUrl string, longUrl string) (string, error) {
	// 计算MurmurHash64
	hash := murmur3.StringSum64(longUrl)
	// 将hash转换为Base62
	base62 := utils.EncodeBase62(hash)
	if base62 == "" {
		return "", fmt.Errorf("base62编码失败")
	}
	// 拼接短网址
	shortUrl := base62[len(base62)-6:]
	urlObj := model.ShortUrl{
		Lurl: rawUrl,
		Surl: shortUrl,
	}
	// 模拟插入数据库
	err := dao.InsertShortUrl(urlObj)
	if err != nil {
		// 处理Hash冲突
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("hash冲突 触发唯一索引 rawUrl = %s, longUrl = %s, shortUrl = %s", rawUrl, longUrl, shortUrl)
			return getShortUrl(rawUrl, getLongUrlRandom(longUrl)) // 递归调用，生成新的短网址
		}
		return "", err
	}
	// 插入成功后更新缓存
	longCache.Put(rawUrl, shortUrl)
	return shortUrl, nil
}
