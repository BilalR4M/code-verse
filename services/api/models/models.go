package models

type Language struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Extension string `json:"extension"`
}

type Difficulty struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CodeSnippet struct {
	ID             string   `json:"id"`
	Language       string   `json:"language"`
	Difficulty     string   `json:"difficulty"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Lines          []string `json:"lines"`
	ExpectedOutput string   `json:"expectedOutput,omitempty"`
	Tags           []string `json:"tags"`
}

type TypingError struct {
	Position  int    `json:"position"`
	Expected  string `json:"expected"`
	Typed     string `json:"typed"`
	Timestamp int64  `json:"timestamp"`
	Corrected bool   `json:"corrected"`
}

type TypingSession struct {
	ID              string        `json:"id"`
	SnippetID       string        `json:"snippetId"`
	StartTime       int64         `json:"startTime"`
	EndTime         int64         `json:"endTime,omitempty"`
	CurrentPosition int           `json:"currentPosition"`
	TypedCharacters []string      `json:"typedCharacters"`
	Errors          []TypingError `json:"errors"`
	Completed       bool          `json:"completed"`
}

type SessionResult struct {
	SessionID         string         `json:"sessionId"`
	SnippetID         string         `json:"snippetId"`
	CompletionTime    int64          `json:"completionTime"`
	TotalCharacters   int            `json:"totalCharacters"`
	CorrectCharacters int            `json:"correctCharacters"`
	TotalErrors       int            `json:"totalErrors"`
	CPM               float64        `json:"cpm"`
	Accuracy          float64        `json:"accuracy"`
	SyntaxScore       float64        `json:"syntaxScore"`
	FinalScore        float64        `json:"finalScore"`
	Breakdown         ScoreBreakdown `json:"breakdown"`
}

type ScoreBreakdown struct {
	BaseScore            float64 `json:"baseScore"`
	SpeedBonus           float64 `json:"speedBonus"`
	AccuracyBonus        float64 `json:"accuracyBonus"`
	SyntaxPenalty        float64 `json:"syntaxPenalty"`
	DifficultyMultiplier float64 `json:"difficultyMultiplier"`
}

type StartSessionRequest struct {
	SnippetID string `json:"snippetId"`
}

type FinishSessionRequest struct {
	SessionID       string        `json:"sessionId"`
	TypedCharacters []string      `json:"typedCharacters"`
	Errors          []TypingError `json:"errors"`
	CompletionTime  int64         `json:"completionTime"`
}

type GetSnippetsParams struct {
	Language   string `query:"language"`
	Difficulty string `query:"difficulty"`
	Limit      int    `query:"limit"`
}

type ApiResponse[T any] struct {
	Success bool   `json:"success"`
	Data    T      `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}
