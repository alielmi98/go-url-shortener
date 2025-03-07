package handler

import (
	"net/http"

	"github.com/alielmi98/go-url-shortener/api/dto"
	"github.com/alielmi98/go-url-shortener/api/helper"
	"github.com/alielmi98/go-url-shortener/config"
	"github.com/alielmi98/go-url-shortener/usecase"
	"github.com/gin-gonic/gin"
)

type ShortnUrlsHandler struct {
	usecase usecase.ShortenUrlUsecase
}

func NewShortnUrlsHandler(cfg *config.Config) *ShortnUrlsHandler {
	return &ShortnUrlsHandler{
		usecase: usecase.NewShortenUrlUsecase(),
	}
}

// Create godoc
// @Summary Create shortn url
// @Description Create shortn url
// @Tags shortn_urls
// @Accept  json
// @Produce  json
// @Param Request body dto.CreateShortnUrlRequest true "CreateShortnUrlRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/shorten [post]
func (h *ShortnUrlsHandler) Create(c *gin.Context) {
	var createReqDTO dto.CreateShortnUrlRequest
	if err := c.ShouldBindJSON(&createReqDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	shortnUrl, err := h.usecase.CreateShortnUrl(c, &createReqDTO)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(shortnUrl, true, helper.Success))

}

// Update godoc
// @Summary Update shortn url
// @Description Update shortn url
// @Tags shortn_urls
// @Accept  json
// @Produce  json
// @Param short_code path string true "ShortnUrl Short Code"
// @Param Request body dto.UpdateShortnUrlRequest true "UpdateShortnUrlRequest"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 404 {object} helper.BaseHttpResponse "Failed"
// @Failure 500 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/shorten/{short_code} [put]
func (h *ShortnUrlsHandler) Update(c *gin.Context) {
	shortCode := c.Params.ByName("short_code")
	var updateReqDTO dto.UpdateShortnUrlRequest
	if err := c.ShouldBindJSON(&updateReqDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	err := h.usecase.UpdateShortUrl(c, shortCode, &updateReqDTO)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("record updated successfully", true, helper.Success))
}

// Delete godoc
// @Summary Delete shortn url
// @Description Delete shortn url
// @Tags shortn_urls
// @Accept  json
// @Produce  json
// @Param short_code path string true "ShortnUrl Short Code"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 404 {object} helper.BaseHttpResponse "Failed"
// @Failure 500 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/shorten/{short_code} [delete]
func (h *ShortnUrlsHandler) Delete(c *gin.Context) {
	shortCode := c.Params.ByName("short_code")
	err := h.usecase.DeleteShortUrl(c, shortCode)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("record deleted successfully", true, helper.Success))

}

// GetByShortCode godoc
// @Summary Get shortn url by short code
// @Description Get shortn url by short code
// @Tags shortn_urls
// @Accept  json
// @Produce  json
// @Param short_code path string true "ShortnUrl Short Code"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 404 {object} helper.BaseHttpResponse "Failed"
// @Failure 500 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/shorten/{short_code}/stats [get]
func (h *ShortnUrlsHandler) GetByShortCode(c *gin.Context) {
	shortCode := c.Params.ByName("short_code")
	response, err := h.usecase.GetByShortCode(c, shortCode)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, helper.Success))
}

// RedirectToOriginalURL godoc
// @Summary Redirect to original URL
// @Description Redirect to original URL using short code
// @Tags shortn_urls
// @Accept  json
// @Produce  json
// @Param short_code path string true "ShortnUrl Short Code"
// @Success 302 {object} helper.BaseHttpResponse "Redirect"
// @Failure 404 {object} helper.BaseHttpResponse "Failed"
// @Failure 500 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/shorten/{short_code} [get]
func (h *ShortnUrlsHandler) RedirectToOriginalURL(c *gin.Context) {
	shortCode := c.Params.ByName("short_code")

	originalURL, err := h.usecase.GetOriginalURL(c, shortCode)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	err = h.usecase.IncrementAccessCount(c, shortCode)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.Redirect(http.StatusFound, originalURL)
}
