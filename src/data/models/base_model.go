package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id int `json:"id" gorm:"primarykey"`

	CreatedAt  time.Time    `json:"created_at"  gorm:"type:TIMESTAMP with time zone;not null"`
	ModifiedAt sql.NullTime `json:"modified_at" gorm:"type:TIMESTAMP with time zone;null"`
	DeletedAt  sql.NullTime `json:"deleted_at" gorm:"type:TIMESTAMP with time zone;null;index"`
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
