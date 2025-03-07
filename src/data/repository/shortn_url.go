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
	Update(ctx context.Context, shortCode string, model *models.ShortURL) error
	Delete(ctx context.Context, shortCode string) error
	Exists(shortUrl string) (bool, error)
	GetByShortCode(ctx context.Context, shortCode string) (*models.ShortURL, error)
	IncrementAccessCount(ctx context.Context, shortCode string) error
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

func (r *shortUrlRepository) Update(ctx context.Context, shortCode string, model *models.ShortURL) error {
	tx := r.db.WithContext(ctx).Begin()
	if err := tx.Model(model).
		Where("short_code = ? ", shortCode).
		Updates(model).
		Error; err != nil {
		tx.Rollback()
		log.Printf("Caller:%s Level:%s Msg:%s ", constants.Postgres, constants.Update, err.Error())
		return err
	}
	tx.Commit()

	return nil
}

func (r *shortUrlRepository) Delete(ctx context.Context, shortCode string) error {
	tx := r.db.WithContext(ctx).Begin()
	model := new(models.ShortURL)
	if err := tx.Where("short_code = ?", shortCode).Delete(model).Error; err != nil {
		tx.Rollback()
		log.Printf("Caller:%s Level:%s Msg:%s ", constants.Postgres, constants.Delete, err.Error())
		return err
	}
	tx.Commit()

	return nil
}

func (r *shortUrlRepository) Exists(shortUrl string) (bool, error) {
	var exists bool
	if err := r.db.Model(&models.ShortURL{}).
		Select("count(*) > 0").
		Where("short_code = ?", shortUrl).
		Find(&exists).
		Error; err != nil {
		log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.Select, err.Error())
		return false, err
	}
	return exists, nil
}

func (r *shortUrlRepository) GetByShortCode(ctx context.Context, shortCode string) (*models.ShortURL, error) {
	model := new(models.ShortURL)
	if err := r.db.WithContext(ctx).
		Where("short_code = ?", shortCode).
		First(model).
		Error; err != nil {
		log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.Select, err.Error())
		return nil, err
	}
	return model, nil
}

func (r *shortUrlRepository) IncrementAccessCount(ctx context.Context, shortCode string) error {
	model := new(models.ShortURL)
	if err := r.db.Model(model).
		Where("short_code = ?", shortCode).
		Update("access_count", gorm.Expr("access_count + ?", 1)).
		Error; err != nil {
		log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.Update, err.Error())
		return err
	}
	return nil
}
