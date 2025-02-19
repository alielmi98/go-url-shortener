package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id int `gorm:"primarykey"`

	CreatedAt  time.Time    `gorm:"type:TIMESTAMP with time zone;not null"`
	ModifiedAt sql.NullTime `gorm:"type:TIMESTAMP with time zone;null"`
	DeletedAt  sql.NullTime `gorm:"type:TIMESTAMP with time zone;null;index"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now().UTC()
	return
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	return
}

func (m *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	m.DeletedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}

	return
}
