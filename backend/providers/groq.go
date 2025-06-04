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

// GroqProvider implements the GenAIProvider interface for Groq
type GroqProvider struct {
	baseURL string
}

// GroqRequest represents a request to the Groq API
type GroqRequest struct {
	Model     string        `json:"model"`
	Messages  []GroqMessage `json:"messages"`
	MaxTokens int           `json:"max_tokens"`
}

// GroqMessage represents a message in the Groq API request
type GroqMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// GroqResponse represents a response from the Groq API
type GroqResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

// NewGroqProvider creates a new Groq provider
func NewGroqProvider() *GroqProvider {
	return &GroqProvider{
		baseURL: "https://api.groq.com/openai/v1/chat/completions",
	}
}

// GetName returns the name of the provider
func (p *GroqProvider) GetName() string {
	return "groq"
}

// ValidateAPIKey checks if the API key is valid
func (p *GroqProvider) ValidateAPIKey(apiKey string) bool {
	return len(apiKey) > 0 && strings.HasPrefix(apiKey, "gsk_")
}

// GenerateSolution generates a solution for a LeetCode problem
func (p *GroqProvider) GenerateSolution(ctx context.Context, problem string, language string, userLevel string, apiKey string) (*models.SolutionResponse, error) {
	if !p.ValidateAPIKey(apiKey) {
		return nil, errors.New("invalid API key")
	}

	prompt := generatePrompt(problem, language, userLevel)
	
	// Create the request
	reqBody := GroqRequest{
		Model: "llama3-70b-8192",
		Messages: []GroqMessage{
			{
				Role:    "system",
				Content: "You are a coding expert that helps solve LeetCode problems. Provide detailed explanations, efficient code solutions, and helpful hints. Format your response as JSON with fields: explanation, code, hints (array), timeComplexity, and spaceComplexity.",
			},
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
	req.Header.Set("Authorization", "Bearer "+apiKey)
	
	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	// Parse response
	var groqResp GroqResponse
	if err := json.NewDecoder(resp.Body).Decode(&groqResp); err != nil {
		return nil, err
	}
	
	// Check for API error
	if resp.StatusCode != http.StatusOK {
		errorMsg := "Groq API error"
		if groqResp.Error.Message != "" {
			errorMsg = groqResp.Error.Message
		}
		return nil, errors.New(errorMsg)
	}
	
	// Check if we have a response
	if len(groqResp.Choices) == 0 {
		return nil, errors.New("no response from Groq API")
	}
	
	// Parse the content
	content := groqResp.Choices[0].Message.Content
	
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
