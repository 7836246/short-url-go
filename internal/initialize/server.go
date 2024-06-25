package initialize

import (
	"fmt"
	"short-url-go/internal/global"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := Routers()

	address := fmt.Sprintf(":%d", global.G_CONFIG.Http.Port)
	s := initServer(address, Router)
	global.G_LOG.Error(s.ListenAndServe().Error())
}
