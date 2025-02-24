package services

import (
	"errors"

	"github.com/alielmi98/go-url-shortener/data/repository"
	"github.com/alielmi98/go-url-shortener/pkg/service_errors"
	"github.com/google/uuid"
)

type ShortUrlGenerator struct {
	length int
	repo   repository.ShortUrlRepository
}

func NewShortUrlGenerator(length int) *ShortUrlGenerator {
	return &ShortUrlGenerator{
		length: length,
		repo:   repository.NewShortUrlRepository()}
}

func (g *ShortUrlGenerator) Generate() (string, error) {
	if g.length <= 0 {
		return "", &service_errors.ServiceError{
			EndUserMessage:   "Invalid length",
			TechnicalMessage: "The length of the short URL must be greater than 0",
			Err:              errors.New("invalid length"),
		}
	}

	for {
		shortUrl := uuid.New().String()[:g.length]

		exists, err := g.repo.Exists(shortUrl)
		if err != nil {
			return "", &service_errors.ServiceError{
				EndUserMessage:   "Internal server error",
				TechnicalMessage: "Failed to check if short URL exists",
				Err:              err,
			}
		}

		if !exists {
			return shortUrl, nil
		}
	}
}
