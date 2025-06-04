package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/leetcode-helper/api/api"
	"github.com/leetcode-helper/api/providers"
	"github.com/leetcode-helper/api/services"
)

func main() {
	// Create provider registry
	providerRegistry := providers.NewProviderRegistry()

	// Register providers
	providerRegistry.RegisterProvider(providers.NewOpenAIProvider())
	providerRegistry.RegisterProvider(providers.NewGeminiProvider())
	providerRegistry.RegisterProvider(providers.NewClaudeProvider())
	providerRegistry.RegisterProvider(providers.NewGroqProvider())

	// Create services
	solutionService := services.NewSolutionService(providerRegistry)

	// Create API handlers
	handler := api.NewHandler(solutionService)

	// Create Gin router
	router := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	// Setup routes
	api.SetupRoutes(router, handler)

	// Start server
	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
