package utils

import (
	"be/entity"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type CohereService struct {
	apiKey string
	client *http.Client
}

type CohereResponse struct {
	Results []struct {
		Text string `json:"text"`
	} `json:"results"`
}

func NewCohereService() *CohereService {
	return &CohereService{
		apiKey: os.Getenv("COHERE_API_KEY"),
		client: &http.Client{},
	}
}

func (s *CohereService) GetItemRecommendations(ctx context.Context, userQuery string, items []entity.Item) ([]entity.Item, error) {
	if len(items) == 0 {
		return []entity.Item{}, nil
	}

	// Tạo prompt mới yêu cầu Cohere trả về tên sản phẩm, phân tách bằng dấu phẩy
	prompt := fmt.Sprintf(`Given these Items:\n%s. %s.\nReturn only the Item names, separated by commas.`,
		formatItemsForPrompt(items), userQuery)

	reqBody := map[string]interface{}{
		"prompt":      prompt,
		"max_tokens":  100,
		"temperature": 0.7,
		"model":       "command",
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST",
		"https://api.cohere.ai/v1/generate",
		strings.NewReader(string(reqBodyBytes)))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+s.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cohere API returned status code: %d", resp.StatusCode)
	}

	var cohereResp CohereResponse
	if err := json.NewDecoder(resp.Body).Decode(&cohereResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	if len(cohereResp.Results) == 0 {
		if len(items) > 5 {
			return items[:5], nil
		}
		return items, nil
	}

	return filterRecommendedItems(items, cohereResp.Results[0].Text), nil
}

func formatItemsForPrompt(items []entity.Item) string {
	var sb strings.Builder
	for _, item := range items {
		sb.WriteString(fmt.Sprintf("- %s: %s\n",
			item.Name, item.Description))
	}
	return sb.String()
}

func filterRecommendedItems(items []entity.Item, recommendation string) []entity.Item {
	names := strings.Split(recommendation, ",")
	nameMap := map[string]bool{}
	for _, n := range names {
		nameMap[strings.TrimSpace(n)] = true
	}

	var result []entity.Item
	for _, item := range items {
		if nameMap[item.Name] {
			result = append(result, item)
		}
	}
	return result
}
