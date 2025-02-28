package usecase

import (
	"context"
	"log"
	"time"

	"github.com/alielmi98/go-url-shortener/api/dto"
	"github.com/alielmi98/go-url-shortener/common"
	"github.com/alielmi98/go-url-shortener/constants"
	"github.com/alielmi98/go-url-shortener/data/cache"
	"github.com/alielmi98/go-url-shortener/data/models"
	"github.com/alielmi98/go-url-shortener/data/repository"
	"github.com/alielmi98/go-url-shortener/services"
	"github.com/go-redis/redis/v7"
	"github.com/mitchellh/mapstructure"
)

type ShortenUrlUsecase interface {
	CreateShortnUrl(ctx context.Context, url *dto.CreateShortnUrlRequest) (*dto.ShortnUrlResponse, error)
	UpdateShortUrl(ctx context.Context, id int, url *dto.UpdateShortnUrlRequest) (*dto.ShortnUrlResponse, error)
	DeleteShortUrl(ctx context.Context, shortCode string) error
	GetByShortCode(ctx context.Context, shortCode string) (*dto.ShortnUrlResponse, error)
	IncrementAccessCount(ctx context.Context, shortCode string) error
}

// ShortenUrlUsecase implementation
type shortenUrlUsecase struct {
	repo    repository.ShortUrlRepository
	service *services.ShortUrlGenerator
	cache   *redis.Client
}

func NewShortenUrlUsecase() ShortenUrlUsecase {
	return &shortenUrlUsecase{
		repo:    repository.NewShortUrlRepository(),
		service: services.NewShortUrlGenerator(6),
		cache:   cache.GetRedis(),
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

	createdModel, err := u.repo.Create(ctx, model)
	if err != nil {
		return nil, err
	}

	response, err := common.TypeConverter[dto.ShortnUrlResponse](createdModel)
	if err != nil {
		return &dto.ShortnUrlResponse{}, err
	}

	err = cache.Set(u.cache, response.ShortCode, response, 24*time.Hour)
	if err != nil {
		log.Printf("Caller:%s Level:%s Msg:%s ", constants.Redis, constants.Insert, err.Error())
		return response, nil
	}
	return response, nil
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

	err = cache.Set(u.cache, response.ShortCode, response, 24*time.Hour)
	if err != nil {
		log.Printf("Caller:%s Level:%s Msg:%s ", constants.Redis, constants.Insert, err.Error())
		return response, nil
	}
	return response, nil
}

func (u *shortenUrlUsecase) DeleteShortUrl(ctx context.Context, shortCode string) error {
	err := u.repo.Delete(ctx, shortCode)
	if err != nil {
		return err
	}
	u.cache.Del(shortCode)
	return nil
}

func (u *shortenUrlUsecase) GetByShortCode(ctx context.Context, shortCode string) (*dto.ShortnUrlResponse, error) {
	response, err := cache.Get[*dto.ShortnUrlResponse](u.cache, shortCode)
	if err == nil && response != nil {
		return response, nil
	}

	model, err := u.repo.GetByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	response, err = common.TypeConverter[dto.ShortnUrlResponse](model)
	if err != nil {
		return &dto.ShortnUrlResponse{}, err
	}
	err = cache.Set(u.cache, response.ShortCode, response, 24*time.Hour)
	if err != nil {
		log.Printf("Caller:%s Level:%s Msg:%s ", constants.Redis, constants.Insert, err.Error())
		return response, nil
	}
	return response, nil
}

func (u *shortenUrlUsecase) IncrementAccessCount(ctx context.Context, shortCode string) error {
	return u.repo.IncrementAccessCount(ctx, shortCode)
}
