package web

import (
	"avsg/embeds"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"net/http/httputil"
	"net/url"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// 8 MiB file upload should be more than enough for an AV1 save file
	router.MaxMultipartMemory = 8 << 20
	api := router.Group("api") // Prefixing all of our URLs to make routing easier in dev
	api.POST("/encrypt", Encrypt)
	api.POST("/decrypt", Decrypt)

	mountFrontend(router)
	return router
}

func mountFrontend(router *gin.Engine) {
	// In prod, serve the frontend assets directly
	if os.Getenv("GIN_MODE") == "release" {
		router.Use(static.Serve("/", static.EmbedFolder(embeds.Frontend, "dist")))
		return
	}

	// In dev, proxy to our development server
	frontendUrl, _ := url.Parse(viper.GetString("frontendUrl"))
	proxy := httputil.NewSingleHostReverseProxy(frontendUrl)

	router.NoRoute(gin.WrapH(proxy))
}
