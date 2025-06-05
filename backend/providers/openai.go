package providers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/leetcode-helper/api/models"
)

// OpenAIProvider implements the GenAIProvider interface for OpenAI
type OpenAIProvider struct {
	baseURL string
}

// OpenAIRequest represents a request to the OpenAI API
type OpenAIRequest struct {
	Model    string          `json:"model"`
	Messages []OpenAIMessage `json:"messages"`
}

// OpenAIMessage represents a message in the OpenAI API request
type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OpenAIResponse represents a response from the OpenAI API
type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

// NewOpenAIProvider creates a new OpenAI provider
func NewOpenAIProvider() *OpenAIProvider {
	return &OpenAIProvider{
		baseURL: "https://api.openai.com/v1/chat/completions",
	}
}

// GetName returns the name of the provider
func (p *OpenAIProvider) GetName() string {
	return "openai"
}

// ValidateAPIKey checks if the API key is valid
func (p *OpenAIProvider) ValidateAPIKey(apiKey string) bool {
	return len(apiKey) > 0 && strings.HasPrefix(apiKey, "sk-")
}

// GenerateSolution generates a solution for a LeetCode problem
func (p *OpenAIProvider) GenerateSolution(ctx context.Context, problem string, language string, userLevel string, apiKey string) (*models.SolutionResponse, error) {
	if !p.ValidateAPIKey(apiKey) {
		return nil, errors.New("invalid API key")
	}

	prompt := generatePrompt(problem, language, userLevel)

	// Create the request
	reqBody := OpenAIRequest{
		Model: "gpt-4",
		Messages: []OpenAIMessage{
			{
				Role:    "system",
				Content: "You are a coding expert that helps solve LeetCode problems. Provide detailed explanations, efficient code solutions, and helpful hints. Format your response as JSON with fields: explanation, code, hints (array), timeComplexity, and spaceComplexity.Keep your responses short and precise",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
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
	var openAIResp OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&openAIResp); err != nil {
		return nil, err
	}

	// Check for API error
	if resp.StatusCode != http.StatusOK {
		errorMsg := "OpenAI API error"
		if openAIResp.Error.Message != "" {
			errorMsg = openAIResp.Error.Message
		}
		return nil, errors.New(errorMsg)
	}

	// Check if we have a response
	if len(openAIResp.Choices) == 0 {
		return nil, errors.New("no response from OpenAI API")
	}

	// Parse the content as JSON
	content := openAIResp.Choices[0].Message.Content

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

// Helper function to generate a prompt for the AI
func generatePrompt(problem string, language string, userLevel string) string {
	jsonBlockStart := "```jsonStart"
	jsonBlockEnd := "```jsonEnd"
	return fmt.Sprintf(`
	Solve the following LeetCode problem:
	
	Problem:
	%s
	
	Programming Language: %s
	User Experience Level: %s
	
	Please provide a detailed solution with:
	1. A clear explanation of the approach in points
	2. Efficient code implementation in %s
	3. Helpful hints for understanding key concepts
	4. Time and space complexity analysis in points
	
	Format your response as a JSON object with the following structure
	%s
	{
	  "explanation": "Detailed explanation of the solution approach",
	  "code": "Complete code solution",
	  "hints": ["Hint 1", "Hint 2", "Hint 3"],
	  "timeComplexity": "O(n)",
	  "spaceComplexity": "O(1)"
	}
	%s
	`, problem, language, userLevel, language, jsonBlockStart, jsonBlockEnd)
}

// Helper function to extract JSON from markdown
func extractJSONFromMarkdown(markdown string) string {
	// Look for JSON code blocks
	// jsonBlockStart := "START"
	// jsonBlockEnd := "END"
	jsonBlockStart := "```jsonStart"
	jsonBlockEnd := "```jsonEnd"

	startIdx := strings.Index(markdown, jsonBlockStart)
	if startIdx == -1 {
		// Try without language specifier
		jsonBlockStart = "```"
		startIdx = strings.Index(markdown, jsonBlockStart)
	}

	if startIdx != -1 {
		startIdx += len(jsonBlockStart)
		endIdx := strings.Index(markdown[startIdx:], jsonBlockEnd)
		if endIdx != -1 {
			return strings.TrimSpace(markdown[startIdx : startIdx+endIdx])
		}
	}

	return ""
}

// Helper function to extract code from markdown
func extractCodeFromMarkdown(markdown string) string {
	// Look for code blocks
	codeBlockStart := "```"
	codeBlockEnd := "```"

	startIdx := strings.Index(markdown, codeBlockStart)
	if startIdx != -1 {
		startIdx += len(codeBlockStart)
		endIdx := strings.Index(markdown[startIdx:], codeBlockEnd)
		if endIdx != -1 {
			return strings.TrimSpace(markdown[startIdx : startIdx+endIdx])
		}
	}

	return ""
}
