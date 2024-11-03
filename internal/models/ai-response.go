package models

type AIResponse struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	Text        string `json:"text"`
}
