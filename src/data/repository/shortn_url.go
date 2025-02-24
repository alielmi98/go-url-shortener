package repository

import (
	"context"
	"log"

	"github.com/alielmi98/go-url-shortener/constants"
	"github.com/alielmi98/go-url-shortener/data/db"
	"github.com/alielmi98/go-url-shortener/data/models"
	"gorm.io/gorm"
)

type ShortUrlRepository interface {
	Create(ctx context.Context, model *models.ShortURL) (*models.ShortURL, error)
	Exists(shortUrl string) (bool, error)
}

type shortUrlRepository struct {
	db *gorm.DB
}

func NewShortUrlRepository() ShortUrlRepository {
	return &shortUrlRepository{
		db: db.GetDb(),
	}
}

func (r *shortUrlRepository) Create(ctx context.Context, model *models.ShortURL) (*models.ShortURL, error) {
	tx := r.db.WithContext(ctx).Begin()
	err := tx.
		Create(model).
		Error
	if err != nil {
		tx.Rollback()
		log.Printf("Caller:%s Level:%s Msg:%s ", constants.Postgres, constants.Insert, err.Error())
		return nil, err
	}
	tx.Commit()

	return model, nil
}

func (r *shortUrlRepository) Exists(shortUrl string) (bool, error) {
	var exists bool
	if err := r.db.Model(&models.ShortURL{}).
		Select("count(*) > 0").
		Where("short_url = ?", shortUrl).
		Find(&exists).
		Error; err != nil {
		log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.Select, err.Error())
		return false, err
	}
	return exists, nil
}
