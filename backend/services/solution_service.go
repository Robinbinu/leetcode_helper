package services

import (
	"context"
	"errors"

	"github.com/leetcode-helper/api/models"
	"github.com/leetcode-helper/api/providers"
)

// SolutionService handles the business logic for generating solutions
type SolutionService struct {
	providerRegistry *providers.ProviderRegistry
}

// NewSolutionService creates a new solution service
func NewSolutionService(providerRegistry *providers.ProviderRegistry) *SolutionService {
	return &SolutionService{
		providerRegistry: providerRegistry,
	}
}

// GenerateSolution generates a solution for a LeetCode problem
func (s *SolutionService) GenerateSolution(ctx context.Context, req models.SolveRequest) (*models.SolutionResponse, error) {
	// Get the provider
	provider, err := s.providerRegistry.GetProvider(req.Provider)
	if err != nil {
		return nil, errors.New("provider not found: " + req.Provider)
	}

	// Generate solution
	solution, err := provider.GenerateSolution(ctx, req.ProblemText, req.Language, req.UserLevel, req.APIKey)
	if err != nil {
		return nil, err
	}

	return solution, nil
}

// GetAvailableProviders returns a list of available providers
func (s *SolutionService) GetAvailableProviders() []string {
	return s.providerRegistry.GetProviderNames()
}
