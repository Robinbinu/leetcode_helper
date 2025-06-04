package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leetcode-helper/api/models"
	"github.com/leetcode-helper/api/services"
)

// Handler contains the API handlers
type Handler struct {
	solutionService *services.SolutionService
}

// NewHandler creates a new API handler
func NewHandler(solutionService *services.SolutionService) *Handler {
	return &Handler{
		solutionService: solutionService,
	}
}

// SolveProblem handles the /api/solve endpoint
func (h *Handler) SolveProblem(c *gin.Context) {
	var req models.SolveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid request: " + err.Error(),
		})
		return
	}

	// Validate required fields
	if req.ProblemText == "" || req.Language == "" || req.UserLevel == "" || req.Provider == "" || req.APIKey == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Missing required fields",
		})
		return
	}

	// Generate solution
	solution, err := h.solutionService.GenerateSolution(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to generate solution: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, solution)
}

// GetProviders handles the /api/providers endpoint
func (h *Handler) GetProviders(c *gin.Context) {
	providers := h.solutionService.GetAvailableProviders()
	c.JSON(http.StatusOK, models.ProviderListResponse{
		Providers: providers,
	})
}

// SetupRoutes sets up the API routes
func SetupRoutes(router *gin.Engine, handler *Handler) {
	api := router.Group("/api")
	{
		api.POST("/solve", handler.SolveProblem)
		api.GET("/providers", handler.GetProviders)
	}
}
