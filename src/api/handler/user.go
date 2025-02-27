package handler

import (
	"net/http"

	"github.com/alielmi98/go-url-shortener/api/dto"
	"github.com/alielmi98/go-url-shortener/api/helper"
	"github.com/alielmi98/go-url-shortener/config"
	"github.com/alielmi98/go-url-shortener/constants"
	"github.com/alielmi98/go-url-shortener/services"
	"github.com/alielmi98/go-url-shortener/usecase"

	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	usecase *usecase.UserUsecase
	config  *config.Config
	service *services.TokenService
}

func NewUsersHandler(cfg *config.Config) *UsersHandler {
	usecase := usecase.NewUserUsecase(cfg)
	config := cfg
	service := services.NewTokenService(cfg)
	return &UsersHandler{
		usecase: usecase,
		config:  config,
		service: service,
	}
}

// RegisterByUsername godoc
// @Summary Register by username
// @Description Register by username
// @Tags users
// @Accept  json
// @Produce  json
// @Param Request body dto.RegisterUserByUsernameRequest true "RegisterUserByUsernameRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/register-by-username [post]
func (h *UsersHandler) RegisterByUsername(c *gin.Context) {
	req := new(dto.RegisterUserByUsernameRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	err := h.usecase.RegisterByUsername(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, helper.Success))
}

// LoginByUsername godoc
// @Summary Login by username
// @Description Login by username
// @Tags users
// @Accept  json
// @Produce  json
// @Param Request body dto.LoginByUsernameRequest true "LoginByUsernameRequest"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 401 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/login-by-username [post]
func (h *UsersHandler) LoginByUsername(c *gin.Context) {
	req := new(dto.LoginByUsernameRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	token, err := h.usecase.LoginByUsername(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	// Set the refresh token in a cookie
	c.SetCookie(constants.RefreshTokenCookieName, token.RefreshToken, int(h.config.JWT.RefreshTokenExpireDuration*60), "/", h.config.Server.Domin, true, true)

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(token, true, helper.Success))
}

// RefreshToken godoc
// @Summary Refresh token
// @Description Refresh token
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 401 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/refresh-token [get]
func (h *UsersHandler) RefreshToken(c *gin.Context) {
	token, err := h.service.RefreshToken(c)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	// Set the refresh token in a cookie
	c.SetCookie(constants.RefreshTokenCookieName, token.RefreshToken, int(h.config.JWT.RefreshTokenExpireDuration*60), "/", h.config.Server.Domin, true, true)

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(token, true, helper.Success))
}
