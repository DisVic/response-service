package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AIService struct {
	url string
}

type AIResponse struct {
	Description string `json:"description"`
	Text        string `json:"text"`
}

func NewAIService(url string) *AIService {
	return &AIService{url: url}
}

func (s *AIService) FetchData() (*AIResponse, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса к AI: %w", err)
	}
	defer resp.Body.Close()

	var data AIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа AI: %w", err)
	}
	return &data, nil
}
