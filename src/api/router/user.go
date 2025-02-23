package router

import (
	"github.com/alielmi98/go-url-shortener/api/handler"
	"github.com/alielmi98/go-url-shortener/config"

	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewUsersHandler(cfg)
	router.POST("/register-by-username", h.RegisterByUsername)
	router.POST("/login-by-username", h.LoginByUsername)

}
