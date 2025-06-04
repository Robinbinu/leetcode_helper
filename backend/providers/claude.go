package providers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/leetcode-helper/api/models"
)

// ClaudeProvider implements the GenAIProvider interface for Anthropic's Claude
type ClaudeProvider struct {
	baseURL string
}

// ClaudeRequest represents a request to the Claude API
type ClaudeRequest struct {
	Model     string          `json:"model"`
	Messages  []ClaudeMessage `json:"messages"`
	MaxTokens int             `json:"max_tokens"`
}

// ClaudeMessage represents a message in the Claude API request
type ClaudeMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ClaudeResponse represents a response from the Claude API
type ClaudeResponse struct {
	Content []struct {
		Text string `json:"text"`
	} `json:"content"`
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

// NewClaudeProvider creates a new Claude provider
func NewClaudeProvider() *ClaudeProvider {
	return &ClaudeProvider{
		baseURL: "https://api.anthropic.com/v1/messages",
	}
}

// GetName returns the name of the provider
func (p *ClaudeProvider) GetName() string {
	return "claude"
}

// ValidateAPIKey checks if the API key is valid
func (p *ClaudeProvider) ValidateAPIKey(apiKey string) bool {
	return len(apiKey) > 0 && strings.HasPrefix(apiKey, "sk-")
}

// GenerateSolution generates a solution for a LeetCode problem
func (p *ClaudeProvider) GenerateSolution(ctx context.Context, problem string, language string, userLevel string, apiKey string) (*models.SolutionResponse, error) {
	if !p.ValidateAPIKey(apiKey) {
		return nil, errors.New("invalid API key")
	}

	prompt := generatePrompt(problem, language, userLevel)
	
	// Create the request
	reqBody := ClaudeRequest{
		Model: "claude-3-opus-20240229",
		Messages: []ClaudeMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		MaxTokens: 4096,
	}
	
	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	
	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, "POST", p.baseURL, bytes.NewBuffer(reqJSON))
	if err != nil {
		return nil, err
	}
	
	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")
	
	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	// Parse response
	var claudeResp ClaudeResponse
	if err := json.NewDecoder(resp.Body).Decode(&claudeResp); err != nil {
		return nil, err
	}
	
	// Check for API error
	if resp.StatusCode != http.StatusOK {
		errorMsg := "Claude API error"
		if claudeResp.Error.Message != "" {
			errorMsg = claudeResp.Error.Message
		}
		return nil, errors.New(errorMsg)
	}
	
	// Check if we have a response
	if len(claudeResp.Content) == 0 {
		return nil, errors.New("no response from Claude API")
	}
	
	// Parse the content
	content := claudeResp.Content[0].Text
	
	// Try to extract JSON from the content
	var solution models.SolutionResponse
	
	// First, try to parse the entire content as JSON
	err = json.Unmarshal([]byte(content), &solution)
	if err != nil {
		// If that fails, try to extract JSON from markdown code blocks
		jsonStr := extractJSONFromMarkdown(content)
		if jsonStr != "" {
			err = json.Unmarshal([]byte(jsonStr), &solution)
			if err != nil {
				return nil, errors.New("failed to parse solution from response")
			}
		} else {
			// If no JSON found, create a basic response from the text
			return &models.SolutionResponse{
				Explanation: content,
				Code:        extractCodeFromMarkdown(content),
				Hints:       []string{"No structured hints available"},
			}, nil
		}
	}
	
	return &solution, nil
}
