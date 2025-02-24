package usecase

import (
	"context"

	"github.com/alielmi98/go-url-shortener/api/dto"
	"github.com/alielmi98/go-url-shortener/common"
	"github.com/alielmi98/go-url-shortener/data/models"
	"github.com/alielmi98/go-url-shortener/data/repository"
	"github.com/alielmi98/go-url-shortener/services"
)

type ShortenUrlUsecase interface {
	CreateShortnUrl(ctx context.Context, url *dto.CreateShortnUrlRequest) (*dto.ShortnUrlResponse, error)
}

// ShortenUrlUsecase implementation
type shortenUrlUsecase struct {
	repo    repository.ShortUrlRepository
	service *services.ShortUrlGenerator
}

func NewShortenUrlUsecase() ShortenUrlUsecase {
	return &shortenUrlUsecase{
		repo:    repository.NewShortUrlRepository(),
		service: services.NewShortUrlGenerator(6),
	}
}

func (u *shortenUrlUsecase) CreateShortnUrl(ctx context.Context, url *dto.CreateShortnUrlRequest) (*dto.ShortnUrlResponse, error) {
	model := new(models.ShortURL)
	shortURL, err := u.service.Generate()
	if err != nil {
		return nil, err
	}
	model.ShortURL = shortURL
	model.OriginalURL = url.OriginalURL

	createdModel, err := u.repo.Create(ctx, model)
	if err != nil {
		return nil, err
	}

	response, err := common.TypeConverter[dto.ShortnUrlResponse](createdModel)
	if err != nil {
		return &dto.ShortnUrlResponse{}, err
	}

	return response, err
}
