package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"short-url-go/internal/global"
)

func GetSeoHome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, global.G_CONFIG.Seo)
}
