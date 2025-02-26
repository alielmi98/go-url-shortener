package handler

import (
	"net/http"
	"strconv"

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
// @Param id path int true "ShortnUrl ID"
// @Param Request body dto.UpdateShortnUrlRequest true "UpdateShortnUrlRequest"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 404 {object} helper.BaseHttpResponse "Failed"
// @Failure 500 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/shorten/{id} [put]
func (h *ShortnUrlsHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	var updateReqDTO dto.UpdateShortnUrlRequest
	if err := c.ShouldBindJSON(&updateReqDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	response, err := h.usecase.UpdateShortUrl(c, id, &updateReqDTO)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, helper.Success))
}

// Delete godoc
// @Summary Delete shortn url
// @Description Delete shortn url
// @Tags shortn_urls
// @Accept  json
// @Produce  json
// @Param id path int true "ShortnUrl ID"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 404 {object} helper.BaseHttpResponse "Failed"
// @Failure 500 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/shorten/{id} [delete]
func (h *ShortnUrlsHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	err := h.usecase.DeleteShortUrl(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, helper.Success))
}
