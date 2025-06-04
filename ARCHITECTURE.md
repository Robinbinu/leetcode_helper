# LeetCode Helper Architecture

## Overview

This document outlines the architecture for the LeetCode Helper application, which consists of a Go (Gin-based) backend API and an AstroJS frontend. The application helps users solve LeetCode problems by leveraging various GenAI providers.

## System Components

### 1. Backend (Go/Gin)

The backend is a RESTful API built with Go and the Gin framework. It provides endpoints for processing LeetCode problems and generating solutions using various GenAI providers.

#### Key Features:
- **Provider-Agnostic GenAI Service**: Abstract interface for multiple AI providers
- **Per-Request API Key Handling**: API keys are passed with each request
- **Extensible Architecture**: Easy addition of new GenAI providers
- **Prompt Template Management**: Standardized templates for different problem types
- **Response Formatting**: Consistent JSON structure for frontend consumption

#### Core Components:
- **API Layer**: Gin-based REST endpoints
- **Service Layer**: Business logic for processing problems
- **Provider Layer**: Abstraction for different GenAI providers
- **Models**: Data structures for requests and responses
- **Utils**: Helper functions and common utilities

### 2. Frontend (AstroJS)

The frontend is built with AstroJS and provides a user interface for inputting LeetCode problems, selecting GenAI providers, and viewing solutions.

#### Key Features:
- **Provider Selection**: Dropdown for choosing GenAI provider
- **API Key Management**: Local storage of API keys
- **Problem Input**: Text area for entering LeetCode problem statements
- **Language Selection**: Options for programming language preference
- **User Level Selection**: Options for user experience level
- **Structured Output Display**: Formatted display of code, explanations, and hints

## Data Flow

1. User inputs problem text, selects language, level, and GenAI provider in the frontend
2. Frontend sends request to backend with problem details and API key
3. Backend processes request and forwards to appropriate GenAI provider
4. GenAI provider returns response
5. Backend formats response and returns to frontend
6. Frontend displays structured output to user

## API Endpoints

### POST /api/solve
- **Description**: Process a LeetCode problem and generate a solution
- **Request Body**:
  ```json
  {
    "problem_text": "string",
    "language": "string",
    "user_level": "string",
    "provider": "string",
    "api_key": "string"
  }
  ```
- **Response**:
  ```json
  {
    "explanation": "string",
    "code": "string",
    "hints": ["string"],
    "time_complexity": "string",
    "space_complexity": "string"
  }
  ```

### GET /api/providers
- **Description**: Get list of supported GenAI providers
- **Response**:
  ```json
  {
    "providers": ["string"]
  }
  ```

## GenAI Provider Interface

The backend implements a common interface for all GenAI providers:

```go
type GenAIProvider interface {
    GenerateSolution(ctx context.Context, problem string, language string, userLevel string, apiKey string) (*SolutionResponse, error)
    GetName() string
    ValidateAPIKey(apiKey string) bool
}
```

This interface allows for easy addition of new providers without changing the core application logic.

## Security Considerations

- API keys are never stored on the backend
- Frontend stores API keys in local storage (encrypted when possible)
- HTTPS is used for all communications
- Input validation is performed on both frontend and backend

## Extensibility

New GenAI providers can be added by:
1. Implementing the GenAIProvider interface
2. Registering the new provider in the provider registry
3. Adding the provider to the frontend dropdown

## Dependencies

### Backend:
- Go 1.21+
- Gin Web Framework
- net/http for API requests
- encoding/json for JSON handling

### Frontend:
- AstroJS
- React (via Astro integration)
- TailwindCSS for styling
- localStorage for API key management
