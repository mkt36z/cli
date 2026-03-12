package local

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/mkt36z/cli/internal/api"
)

// contentListResponse is the API response for listing assets.
type contentListResponse struct {
	Assets []struct {
		Name string `json:"name"`
	} `json:"assets"`
}

// contentGetResponse is the API response for getting a single asset.
type contentGetResponse struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// fetchAssetList calls GET /api/v1/content/{assetType} and returns asset names.
func fetchAssetList(client *api.Client, assetType string) ([]string, error) {
	if client == nil {
		return nil, fmt.Errorf("not authenticated. Run `mkt36z auth login` first")
	}

	path := fmt.Sprintf("/api/v1/content/%s", assetType)
	resp, err := client.Do(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	var result contentListResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}

	names := make([]string, len(result.Assets))
	for i, a := range result.Assets {
		names[i] = a.Name
	}
	return names, nil
}

// fetchAssetContent calls GET /api/v1/content/{assetType}/{name} and returns the content.
func fetchAssetContent(client *api.Client, assetType, name string) (string, error) {
	if client == nil {
		return "", fmt.Errorf("not authenticated. Run `mkt36z auth login` first")
	}

	path := fmt.Sprintf("/api/v1/content/%s/%s", assetType, name)
	resp, err := client.Do(context.Background(), "GET", path, nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response: %w", err)
	}

	var result contentGetResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("parsing response: %w", err)
	}

	return result.Content, nil
}
