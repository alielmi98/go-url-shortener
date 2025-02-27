package usecase

import (
	"context"

	"github.com/alielmi98/go-url-shortener/api/dto"
	"github.com/alielmi98/go-url-shortener/common"
	"github.com/alielmi98/go-url-shortener/data/models"
	"github.com/alielmi98/go-url-shortener/data/repository"
	"github.com/alielmi98/go-url-shortener/services"
	"github.com/mitchellh/mapstructure"
)

type ShortenUrlUsecase interface {
	CreateShortnUrl(ctx context.Context, url *dto.CreateShortnUrlRequest) (*dto.ShortnUrlResponse, error)
	UpdateShortUrl(ctx context.Context, id int, url *dto.UpdateShortnUrlRequest) (*dto.ShortnUrlResponse, error)
	DeleteShortUrl(ctx context.Context, id int) error
	GetByShortCode(ctx context.Context, shortCode string) (*dto.ShortnUrlResponse, error)
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

	shortCode, err := u.service.GenerateUniqueShortCode(func(code string) (bool, error) {
		return u.repo.Exists(code)
	})
	if err != nil {
		return nil, err
	}
	model.ShortCode = shortCode
	model.OriginalURL = url.OriginalURL
	model.AccessCount = url.AccessCount

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

func (u *shortenUrlUsecase) UpdateShortUrl(ctx context.Context, id int, url *dto.UpdateShortnUrlRequest) (*dto.ShortnUrlResponse, error) {
	model := new(models.ShortURL)
	if err := mapstructure.Decode(url, model); err != nil {
		return nil, err
	}

	err := u.repo.Update(ctx, id, model)
	if err != nil {
		return &dto.ShortnUrlResponse{}, err
	}

	response, err := common.TypeConverter[dto.ShortnUrlResponse](model)
	if err != nil {
		return &dto.ShortnUrlResponse{}, err
	}
	return response, err
}

func (u *shortenUrlUsecase) DeleteShortUrl(ctx context.Context, id int) error {
	return u.repo.Delete(ctx, id)
}

func (u *shortenUrlUsecase) GetByShortCode(ctx context.Context, shortCode string) (*dto.ShortnUrlResponse, error) {
	model, err := u.repo.GetByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	response, err := common.TypeConverter[dto.ShortnUrlResponse](model)
	if err != nil {
		return &dto.ShortnUrlResponse{}, err
	}
	return response, err
}
