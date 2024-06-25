package initialize

import (
	"github.com/gin-gonic/gin"
	"short-url-go/internal/global"
	"short-url-go/internal/handler"
	"short-url-go/internal/middleware"
)

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	Router.LoadHTMLFiles("/web/dist//*")                        // 添加资源路径
	Router.Static("/_nuxt", "./web/dist/_nuxt/")                // 添加资源路径
	Router.StaticFile("/favicon.ico", "./web/dist/favicon.ico") // 添加资源路径
	Router.StaticFile("/", "web/dist/index.html")               //前端接口

	// 跨域，如需跨域可以打开下面的注释
	Router.Use(middleware.Cors()) // 直接放行全部跨域请求

	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group(global.G_CONFIG.Http.RouterPrefix)
	PublicGroup.GET("/url/generate", handler.GenerateShortUrl)
	PublicGroup.GET("/url/getLong", handler.GetLongUrl)
	PublicGroup.GET("/seo/json/home", handler.GetSeoHome)
	global.G_LOG.Info("路由注册成功")
	return Router
}
