package web

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"net/http/httputil"
	"net/url"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	// api := router.Group("api") // Prefixing all of our URLs to make routing easier in dev

	mountFrontend(router)

	return router
}

func mountFrontend(router *gin.Engine) {
	frontendUrl, _ := url.Parse(viper.GetString("frontendUrl"))
	proxy := httputil.NewSingleHostReverseProxy(frontendUrl)

	router.NoRoute(gin.WrapH(proxy))
}
