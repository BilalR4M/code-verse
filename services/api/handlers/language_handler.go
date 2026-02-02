package handlers

import (
	"github.com/gofiber/fiber/v2"
	"codetype-api/models"
)

type LanguageHandler struct{}

func NewLanguageHandler() *LanguageHandler {
	return &LanguageHandler{}
}

func (h *LanguageHandler) GetLanguages(c *fiber.Ctx) error {
	languages := []models.Language{
		{
			ID:        "python",
			Name:      "Python",
			Extension: ".py",
		},
		{
			ID:        "javascript",
			Name:      "JavaScript",
			Extension: ".js",
		},
		{
			ID:        "go",
			Name:      "Go",
			Extension: ".go",
		},
	}

	response := models.ApiResponse[[]models.Language]{
		Success: true,
		Data:    languages,
	}

	return c.JSON(response)
}