package router

import (
	"github.com/alielmi98/go-url-shortener/api/handler"
	"github.com/alielmi98/go-url-shortener/config"
	"github.com/gin-gonic/gin"
)

func Shorten(router *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewShortnUrlsHandler(cfg)
	router.POST("/", h.Create)
}
