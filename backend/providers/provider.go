package providers

import (
	"context"
	"errors"
	"sync"

	"github.com/leetcode-helper/api/models"
)

// GenAIProvider defines the interface that all AI providers must implement
type GenAIProvider interface {
	// GenerateSolution generates a solution for a LeetCode problem
	GenerateSolution(ctx context.Context, problem string, language string, userLevel string, apiKey string) (*models.SolutionResponse, error)
	
	// GetName returns the name of the provider
	GetName() string
	
	// ValidateAPIKey checks if the API key is valid
	ValidateAPIKey(apiKey string) bool
}

// ProviderRegistry manages the available GenAI providers
type ProviderRegistry struct {
	providers map[string]GenAIProvider
	mu        sync.RWMutex
}

// NewProviderRegistry creates a new provider registry
func NewProviderRegistry() *ProviderRegistry {
	return &ProviderRegistry{
		providers: make(map[string]GenAIProvider),
	}
}

// RegisterProvider adds a provider to the registry
func (r *ProviderRegistry) RegisterProvider(provider GenAIProvider) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.providers[provider.GetName()] = provider
}

// GetProvider returns a provider by name
func (r *ProviderRegistry) GetProvider(name string) (GenAIProvider, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	provider, exists := r.providers[name]
	if !exists {
		return nil, errors.New("provider not found")
	}
	
	return provider, nil
}

// GetProviderNames returns a list of all registered provider names
func (r *ProviderRegistry) GetProviderNames() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	names := make([]string, 0, len(r.providers))
	for name := range r.providers {
		names = append(names, name)
	}
	
	return names
}
