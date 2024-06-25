package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "short-url-go/api/v1"
	"short-url-go/internal/service"
)

func GenerateShortUrl(ctx *gin.Context) {
	// 请求参数
	req := v1.GenerateShortUrlReq{
		Url: ctx.Query("url"),
	}
	resp, err := service.GenerateShortUrl(req.Url)
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func GetLongUrl(ctx *gin.Context) {
	// 请求参数
	req := v1.GetLongUrlReq{
		ShortUrl: ctx.Query("url"),
	}
	resp, err := service.GetLongUrl(req.ShortUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
