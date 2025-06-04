# LeetCode Helper Documentation

## Overview

LeetCode Helper is a full-stack application that helps users solve LeetCode problems using various GenAI providers. The application consists of a Go (Gin-based) backend API and an AstroJS frontend with a sleek black, yellow, and white theme. Users can input LeetCode problems, select their preferred programming language and AI provider, and receive structured solutions with explanations, code, and hints.

## Architecture

The application follows a clean, modular architecture:

### Backend (Go/Gin)

- **Provider Abstraction**: Generic interface for multiple AI providers
- **API Endpoints**: RESTful endpoints for solving problems and listing providers
- **Per-Request API Key Handling**: API keys are passed with each request, not stored on the server

### Frontend (AstroJS)

- **React Components**: UI components for input, output, and API key management
- **Local Storage**: Secure storage of API keys in the browser
- **Responsive Design**: Works on both desktop and mobile devices
- **Custom Theme**: Black, yellow, and white color scheme for a distinctive look

## Installation

### Prerequisites

- Go 1.23 or higher
- Node.js 16 or higher
- npm or yarn

### Backend Setup

1. Clone the repository
2. Navigate to the backend directory:
   ```bash
   cd leetcode-helper/backend
   ```
3. Install Go dependencies:
   ```bash
   go mod tidy
   ```
4. Run the backend server:
   ```bash
   go run main.go
   ```
   The server will start on port 8080.

### Frontend Setup

1. Navigate to the frontend directory:
   ```bash
   cd leetcode-helper/frontend
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Run the development server:
   ```bash
   npm run dev
   ```
   The frontend will be available at http://localhost:3000.

## Usage

1. Open the application in your browser
2. Enter a LeetCode problem in the text area
3. Select your preferred programming language and experience level
4. Choose an AI provider from the dropdown
5. Enter your API key for the selected provider
6. Click "Solve Problem" to generate a solution
7. View the explanation, code, and hints in the tabbed interface

## API Endpoints

### POST /api/solve

Generates a solution for a LeetCode problem.

**Request Body:**
```json
{
  "problem_text": "Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.",
  "language": "javascript",
  "user_level": "intermediate",
  "provider": "openai",
  "api_key": "your-api-key"
}
```

**Response:**
```json
{
  "explanation": "To solve this problem, we can use a hash map to store the numbers we've seen so far...",
  "code": "function twoSum(nums, target) {\n  const map = {};\n  for (let i = 0; i < nums.length; i++) {\n    const complement = target - nums[i];\n    if (map[complement] !== undefined) {\n      return [map[complement], i];\n    }\n    map[nums[i]] = i;\n  }\n  return [];\n}",
  "hints": [
    "Consider using a hash map to store values you've already seen",
    "For each number, check if its complement (target - num) exists in the hash map",
    "Time complexity can be reduced to O(n) with the right approach"
  ],
  "timeComplexity": "O(n)",
  "spaceComplexity": "O(n)"
}
```

### GET /api/providers

Returns a list of available AI providers.

**Response:**
```json
{
  "providers": ["openai", "gemini", "claude", "groq"]
}
```

## Supported AI Providers

The application currently supports the following AI providers:

1. **OpenAI (GPT-4)**
   - API Key Format: Starts with "sk-"
   - API Documentation: https://platform.openai.com/docs/api-reference

2. **Google Gemini**
   - API Key Format: Any valid API key
   - API Documentation: https://ai.google.dev/docs

3. **Anthropic Claude**
   - API Key Format: Starts with "sk-"
   - API Documentation: https://docs.anthropic.com/claude/reference/getting-started-with-the-api

4. **Groq (Llama 3)**
   - API Key Format: Starts with "gsk_"
   - API Documentation: https://console.groq.com/docs/quickstart

## Adding New AI Providers

To add a new AI provider:

1. Create a new file in the `providers` directory implementing the `GenAIProvider` interface
2. Register the provider in `main.go`
3. Add the provider to the frontend dropdown in `ProblemInputForm.tsx`

Example implementation:

```go
package providers

import (
    "context"
    "github.com/leetcode-helper/api/models"
)

// NewProviderProvider creates a new provider
func NewProviderProvider() *ProviderProvider {
    return &ProviderProvider{
        baseURL: "https://api.provider.com/v1/completions",
    }
}

// GetName returns the name of the provider
func (p *ProviderProvider) GetName() string {
    return "provider"
}

// ValidateAPIKey checks if the API key is valid
func (p *ProviderProvider) ValidateAPIKey(apiKey string) bool {
    return len(apiKey) > 0
}

// GenerateSolution generates a solution for a LeetCode problem
func (p *ProviderProvider) GenerateSolution(ctx context.Context, problem string, language string, userLevel string, apiKey string) (*models.SolutionResponse, error) {
    // Implementation here
}
```

## Theme Customization

The application uses a black, yellow, and white theme defined in the Tailwind configuration. The main color variables are:

- **Primary (Yellow)**: #FFD700
- **Secondary (Black)**: #000000
- **Background (White)**: #FFFFFF

To modify the theme:

1. Edit the `tailwind.config.cjs` file in the frontend directory
2. Update the color values in the `theme.extend.colors` section
3. Rebuild the frontend with `npm run build`

## Security Considerations

- API keys are stored only in the user's browser local storage
- API keys are never stored on the server
- API keys are sent directly from the frontend to the AI provider
- HTTPS should be used in production to secure API key transmission

## License

This project is licensed under the MIT License - see the LICENSE file for details.
