package services

import (
	"errors"

	"github.com/alielmi98/go-url-shortener/pkg/service_errors"
	"github.com/google/uuid"
)

type ShortUrlGenerator struct {
	length int
}

func NewShortUrlGenerator(length int) *ShortUrlGenerator {
	return &ShortUrlGenerator{
		length: length,
	}
}

func (g *ShortUrlGenerator) GenerateUniqueShortCode(checkExists func(string) (bool, error)) (string, error) {
	if g.length <= 0 {
		return "", &service_errors.ServiceError{
			EndUserMessage:   "Invalid length",
			TechnicalMessage: "The length of the short URL must be greater than 0",
			Err:              errors.New("invalid length"),
		}
	}

	for {
		shorCode := uuid.New().String()[:g.length]

		exists, err := checkExists(shorCode)
		if err != nil {
			return "", &service_errors.ServiceError{
				EndUserMessage:   "Internal server error",
				TechnicalMessage: "Failed to check if short URL exists",
				Err:              err,
			}
		}
		if !exists {
			return shorCode, nil
		}

	}
}
