package models

type ShortURL struct {
	BaseModel
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
	AccessCount int    `json:"access_count"`
}
