package models

type ShortURL struct {
	BaseModel
	OriginalURL string `json:"original_url" gorm:"not null"`
	ShortURL    string `json:"short_url" gorm:"unique;not null"`
	AccessCount int    `json:"access_count" gorm:"default:0"`
}
