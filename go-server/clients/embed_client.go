package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type EmbedResponse struct {
	Vector []float64 `json:"vector"`
}

type EmbedRequest struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// EmbedClient takes raw text, sends it to the embedder service, and returns the vector.
func EmbedClient(text string) ([]float64, error) {
	var embedResp EmbedResponse

	embedUrl := os.Getenv("EMBED_API_URL")
	if embedUrl == "" {
		return nil, fmt.Errorf("embedding service URL not configured")
	}

	// Prepare request payload
	payload := map[string]string{"text": text}
	reqData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request data: %w", err)
	}

	// Send POST request
	resp, err := http.Post(embedUrl, "application/json", bytes.NewReader(reqData))
	if err != nil {
		return nil, fmt.Errorf("failed to make request to embedding service: %w", err)
	}
	defer resp.Body.Close()

	// Decode response
	err = json.NewDecoder(resp.Body).Decode(&embedResp)
	if err != nil {
		return nil, fmt.Errorf("failed to decode embedding response: %w", err)
	}

	return embedResp.Vector, nil
}
