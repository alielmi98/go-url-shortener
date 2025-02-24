package middlewares

import (
	"net/http"
	"strings"

	"github.com/alielmi98/go-url-shortener/api/helper"
	"github.com/alielmi98/go-url-shortener/config"
	"github.com/alielmi98/go-url-shortener/constants"
	"github.com/alielmi98/go-url-shortener/pkg/service_errors"
	"github.com/alielmi98/go-url-shortener/services"
	"github.com/gin-gonic/gin"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	var tokenService = services.NewTokenService(cfg)

	return func(c *gin.Context) {
		var err error
		claimMap := map[string]interface{}{}
		auth := c.GetHeader(constants.AuthorizationHeaderKey)
		token := strings.Split(auth, " ")
		if auth == "" {
			err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenRequired}
		} else {
			claimMap, err = tokenService.GetClaims(token[1])
			if err != nil {
				serviceErr, ok := err.(*service_errors.ServiceError)
				if !ok {
					err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenInvalid}
				} else {
					err = serviceErr
				}
			}
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateBaseResponseWithError(
				nil, false, helper.AuthError, err,
			))
			return
		}

		c.Set(constants.UserIdKey, claimMap[constants.UserIdKey])
		c.Set(constants.UsernameKey, claimMap[constants.UsernameKey])
		c.Set(constants.EmailKey, claimMap[constants.EmailKey])
		c.Set(constants.ExpireTimeKey, claimMap[constants.ExpireTimeKey])

		c.Next()
	}
}
