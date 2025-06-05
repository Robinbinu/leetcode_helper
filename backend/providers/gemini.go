package providers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/leetcode-helper/api/models"
)

// GeminiProvider implements the GenAIProvider interface for Google's Gemini
type GeminiProvider struct {
	baseURL string
}

// GeminiRequest represents a request to the Gemini API
type GeminiRequest struct {
	Contents         []GeminiContent        `json:"contents"`
	GenerationConfig GeminiGenerationConfig `json:"generationConfig"`
}

// GeminiContent represents content in the Gemini API request
type GeminiContent struct {
	Parts []GeminiPart `json:"parts"`
	Role  string       `json:"role"`
}

// GeminiPart represents a part in the Gemini API content
type GeminiPart struct {
	Text string `json:"text"`
}

// GeminiGenerationConfig represents generation configuration for Gemini
type GeminiGenerationConfig struct {
	Temperature     float64 `json:"temperature"`
	MaxOutputTokens int     `json:"maxOutputTokens"`
}

// GeminiResponse represents a response from the Gemini API
type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

// NewGeminiProvider creates a new Gemini provider
func NewGeminiProvider() *GeminiProvider {
	return &GeminiProvider{
		baseURL: "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent",
	}
}

// GetName returns the name of the provider
func (p *GeminiProvider) GetName() string {
	return "gemini"
}

// ValidateAPIKey checks if the API key is valid
func (p *GeminiProvider) ValidateAPIKey(apiKey string) bool {
	return len(apiKey) > 0
}

// GenerateSolution generates a solution for a LeetCode problem
func (p *GeminiProvider) GenerateSolution(ctx context.Context, problem string, language string, userLevel string, apiKey string) (*models.SolutionResponse, error) {
	if !p.ValidateAPIKey(apiKey) {
		return nil, errors.New("invalid API key")
	}

	prompt := generatePrompt(problem, language, userLevel)

	// Create the request
	reqBody := GeminiRequest{
		Contents: []GeminiContent{
			{
				Role: "user",
				Parts: []GeminiPart{
					{
						Text: prompt,
					},
				},
			},
		},
		GenerationConfig: GeminiGenerationConfig{
			Temperature:     0.2,
			MaxOutputTokens: 1000,
		},
	}

	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	// Create HTTP request with API key as query parameter
	url := fmt.Sprintf("%s?key=%s", p.baseURL, apiKey)
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(reqJSON))
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse response
	var geminiResp GeminiResponse
	if err := json.NewDecoder(resp.Body).Decode(&geminiResp); err != nil {
		return nil, err
	}

	// Check for API error
	if resp.StatusCode != http.StatusOK {
		errorMsg := "Gemini API error"
		if geminiResp.Error.Message != "" {
			errorMsg = geminiResp.Error.Message
		}
		return nil, errors.New(errorMsg)
	}

	// Check if we have a response
	if len(geminiResp.Candidates) == 0 {
		return nil, errors.New("no response from Gemini API")
	}

	// Parse the content
	if len(geminiResp.Candidates[0].Content.Parts) == 0 {
		return nil, errors.New("empty response from Gemini API")
	}

	content := geminiResp.Candidates[0].Content.Parts[0].Text
	log.Println(content)

	// Try to extract JSON from the content
	var solution models.SolutionResponse

	// First, try to parse the entire content as JSON
	err = json.Unmarshal([]byte(content), &solution)
	if err != nil {
		// If that fails, try to extract JSON from markdown code blocks
		jsonStr := extractJSONFromMarkdown(content)
		if jsonStr != "" {
			// // Remove outer markers
			cleanStr := strings.ReplaceAll(jsonStr, "```jsonStart", "")
			cleanStr = strings.ReplaceAll(cleanStr, "```jsonEnd", "")

			// Remove code block markers inside the code field
			cleanStr = strings.ReplaceAll(cleanStr, "```", "")

			cleanStr = strings.TrimSpace(cleanStr)
			log.Println("cleanStr: ", cleanStr)

			err = json.Unmarshal([]byte(cleanStr), &solution)
			if err != nil {
				log.Println(err)
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
