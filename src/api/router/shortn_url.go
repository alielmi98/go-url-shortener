package router

import (
	"github.com/alielmi98/go-url-shortener/api/handler"
	"github.com/alielmi98/go-url-shortener/config"
	"github.com/gin-gonic/gin"
)

func Shorten(router *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewShortnUrlsHandler(cfg)
	router.POST("/", h.Create)
	router.PUT("/:short_code", h.Update)
	router.DELETE("/:short_code", h.Delete)
	router.GET("/:short_code/stats", h.GetByShortCode)
	router.GET("/:short_code", h.RedirectToOriginalURL)
}
