package dto

type CreateShortnUrlRequest struct {
	OriginalURL string `json:"original_url" binding:"required"`
}

type UpdateShortnUrlRequest struct {
	OriginalURL string `json:"original_url" binding:"required"`
	AccessCount int    `json:"access_count"`
}

type ShortnUrlResponse struct {
	Id          int    `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortCode   string `json:"short_code"`
	AccessCount int    `json:"access_count"`
}
