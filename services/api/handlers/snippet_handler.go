package handlers

import (
	"strconv"

	"codetype-api/models"
	"codetype/snippet-engine"

	"github.com/gofiber/fiber/v2"
)

type SnippetHandler struct{}

func NewSnippetHandler() *SnippetHandler {
	return &SnippetHandler{}
}

func (h *SnippetHandler) GetSnippets(c *fiber.Ctx) error {
	language := c.Query("language")
	difficulty := c.Query("difficulty")
	limitStr := c.Query("limit", "10")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	snippets, err := snippetengine.GetSnippets(language, difficulty, limit)
	if err != nil {
		response := models.ApiResponse[[]models.CodeSnippet]{
			Success: false,
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Convert snippets to models format
	modelSnippets := make([]models.CodeSnippet, len(snippets))
	for i, snippet := range snippets {
		modelSnippets[i] = models.CodeSnippet{
			ID:             snippet.ID,
			Language:       snippet.Language,
			Difficulty:     snippet.Difficulty,
			Title:          snippet.Title,
			Description:    snippet.Description,
			Lines:          snippet.Lines,
			ExpectedOutput: snippet.ExpectedOutput,
			Tags:           snippet.Tags,
		}
	}

	response := models.ApiResponse[[]models.CodeSnippet]{
		Success: true,
		Data:    modelSnippets,
	}

	return c.JSON(response)
}
