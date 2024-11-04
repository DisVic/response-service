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
	Descr string `json:"descr"`
	Text  string `json:"text"`
}

func NewAIService(url string) *AIService {
	return &AIService{url: url}
}

func (s *AIService) FetchData(query string) (*AIResponse, error) {
	fullURL := fmt.Sprintf("%s?query=%s", s.url, query)

	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса к AI-сервису: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка: AI-сервис вернул статус %d", resp.StatusCode)
	}

	var data AIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа AI: %w", err)
	}
	return &data, nil
}
