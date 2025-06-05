package models

// SolveRequest represents the request body for solving a LeetCode problem
type SolveRequest struct {
	ProblemText string `json:"problem_text" binding:"required"`
	Language    string `json:"language" binding:"required"`
	UserLevel   string `json:"user_level" binding:"required"`
	Provider    string `json:"provider" binding:"required"`
	APIKey      string `json:"api_key" binding:"required"`
}

// SolutionResponse represents the response structure for a solved problem
type SolutionResponse struct {
	Explanation     string   `json:"explanation"`
	Code            string   `json:"code"`
	Hints           []string `json:"hints"`
	TimeComplexity  string   `json:"timeComplexity"`
	SpaceComplexity string   `json:"spaceComplexity"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// ProviderListResponse represents the response for listing available providers
type ProviderListResponse struct {
	Providers []string `json:"providers"`
}
