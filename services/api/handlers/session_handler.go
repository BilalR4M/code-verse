package handlers

import (
	"github.com/gofiber/fiber/v2"
	"codetype-api/models"
	"codetype-api/services"
)

type SessionHandler struct {
	sessionService *services.SessionService
}

func NewSessionHandler(sessionService *services.SessionService) *SessionHandler {
	return &SessionHandler{
		sessionService: sessionService,
	}
}

func (h *SessionHandler) StartSession(c *fiber.Ctx) error {
	var req models.StartSessionRequest
	if err := c.BodyParser(&req); err != nil {
		response := models.ApiResponse[models.TypingSession]{
			Success: false,
			Error:   "Invalid request body",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if req.SnippetID == "" {
		response := models.ApiResponse[models.TypingSession]{
			Success: false,
			Error:   "Snippet ID is required",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	session, err := h.sessionService.StartSession(req.SnippetID)
	if err != nil {
		response := models.ApiResponse[models.TypingSession]{
			Success: false,
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := models.ApiResponse[models.TypingSession]{
		Success: true,
		Data:    *session,
	}

	return c.JSON(response)
}

func (h *SessionHandler) FinishSession(c *fiber.Ctx) error {
	var req models.FinishSessionRequest
	if err := c.BodyParser(&req); err != nil {
		response := models.ApiResponse[models.SessionResult]{
			Success: false,
			Error:   "Invalid request body",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if req.SessionID == "" {
		response := models.ApiResponse[models.SessionResult]{
			Success: false,
			Error:   "Session ID is required",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := h.sessionService.FinishSession(req)
	if err != nil {
		response := models.ApiResponse[models.SessionResult]{
			Success: false,
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := models.ApiResponse[models.SessionResult]{
		Success: true,
		Data:    *result,
	}

	return c.JSON(response)
}