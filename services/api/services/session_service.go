package services

import (
	"errors"
	"strings"
	"sync"
	"time"

	"codetype-api/models"
	"codetype/snippet-engine"

	"github.com/google/uuid"
)

type SessionService struct {
	sessions map[string]*models.TypingSession
	mutex    sync.RWMutex
}

func NewSessionService() *SessionService {
	return &SessionService{
		sessions: make(map[string]*models.TypingSession),
	}
}

func (s *SessionService) StartSession(snippetID string) (*models.TypingSession, error) {
	// Validate snippet exists
	_, err := snippetengine.GetSnippetByID(snippetID)
	if err != nil {
		return nil, errors.New("snippet not found")
	}

	sessionID := uuid.New().String()
	session := &models.TypingSession{
		ID:              sessionID,
		SnippetID:       snippetID,
		StartTime:       time.Now().UnixMilli(),
		CurrentPosition: 0,
		TypedCharacters: []string{},
		Errors:          []models.TypingError{},
		Completed:       false,
	}

	s.mutex.Lock()
	s.sessions[sessionID] = session
	s.mutex.Unlock()

	return session, nil
}

func (s *SessionService) FinishSession(req models.FinishSessionRequest) (*models.SessionResult, error) {
	s.mutex.RLock()
	session, exists := s.sessions[req.SessionID]
	s.mutex.RUnlock()

	if !exists {
		return nil, errors.New("session not found")
	}

	// Get the original snippet
	snippet, err := snippetengine.GetSnippetByID(session.SnippetID)
	if err != nil {
		return nil, errors.New("snippet not found")
	}

	// Update session with final data
	session.EndTime = time.Now().UnixMilli()
	session.TypedCharacters = req.TypedCharacters
	session.Errors = req.Errors
	session.Completed = true

	// Calculate results manually since we have type issues
	expectedText := strings.Join(snippet.Lines, "\n")
	typedText := strings.Join(session.TypedCharacters, "")

	totalCharacters := len(expectedText)
	completionTime := session.EndTime - session.StartTime
	totalErrors := len(session.Errors)

	// Calculate correct characters
	correctCharacters := 0
	minLen := len(expectedText)
	if len(typedText) < minLen {
		minLen = len(typedText)
	}
	for i := 0; i < minLen; i++ {
		if expectedText[i] == typedText[i] {
			correctCharacters++
		}
	}

	// Calculate CPM and accuracy
	timeInMinutes := float64(completionTime) / (1000.0 * 60.0)
	cpm := float64(len(typedText)) / timeInMinutes
	accuracy := float64(correctCharacters) / float64(totalCharacters) * 100

	// Simple syntax score
	syntaxScore := 100.0 - float64(totalErrors)*2.0
	if syntaxScore < 0 {
		syntaxScore = 0
	}

	// Calculate final score
	baseScore := accuracy * 10
	speedBonus := 0.0
	if cpm > 200 {
		speedBonus = (cpm - 200) * 2.0
	}
	accuracyBonus := 0.0
	if accuracy > 95 {
		accuracyBonus = (accuracy - 95) * 10.0
	}
	syntaxPenalty := 100.0 - syntaxScore

	difficultyMultiplier := 1.0
	switch snippet.Difficulty {
	case "easy":
		difficultyMultiplier = 1.0
	case "medium":
		difficultyMultiplier = 1.2
	case "hard":
		difficultyMultiplier = 1.5
	}

	finalScore := (baseScore + speedBonus + accuracyBonus - syntaxPenalty) * difficultyMultiplier

	result := &models.SessionResult{
		SessionID:         session.ID,
		SnippetID:         snippet.ID,
		CompletionTime:    completionTime,
		TotalCharacters:   totalCharacters,
		CorrectCharacters: correctCharacters,
		TotalErrors:       totalErrors,
		CPM:               cpm,
		Accuracy:          accuracy,
		SyntaxScore:       syntaxScore,
		FinalScore:        finalScore,
		Breakdown: models.ScoreBreakdown{
			BaseScore:            baseScore,
			SpeedBonus:           speedBonus,
			AccuracyBonus:        accuracyBonus,
			SyntaxPenalty:        syntaxPenalty,
			DifficultyMultiplier: difficultyMultiplier,
		},
	}

	// Clean up session after some time (in a real app, you'd use a proper cache)
	go func() {
		time.Sleep(10 * time.Minute)
		s.mutex.Lock()
		delete(s.sessions, req.SessionID)
		s.mutex.Unlock()
	}()

	return result, nil
}

func (s *SessionService) GetSession(sessionID string) (*models.TypingSession, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	session, exists := s.sessions[sessionID]
	if !exists {
		return nil, errors.New("session not found")
	}

	return session, nil
}
