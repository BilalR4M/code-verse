package statsengine

import (
	"math"
	"strings"
	"unicode"
)

package statsengine

import (
	"math"
	"reflect"
	"strings"
	"unicode"
)

type ScoreBreakdown struct {
	BaseScore            float64 `json:"baseScore"`
	SpeedBonus           float64 `json:"speedBonus"`
	AccuracyBonus        float64 `json:"accuracyBonus"`
	SyntaxPenalty        float64 `json:"syntaxPenalty"`
	DifficultyMultiplier float64 `json:"difficultyMultiplier"`
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

type TypingError struct {
	Position  int    `json:"position"`
	Expected  string `json:"expected"`
	Typed     string `json:"typed"`
	Timestamp int64  `json:"timestamp"`
	Corrected bool   `json:"corrected"`
}

func CalculateScore(sessionInterface interface{}, snippetInterface interface{}) (*SessionResult, error) {
	// Use reflection to extract data from the interfaces
	sessionVal := reflect.ValueOf(sessionInterface)
	if sessionVal.Kind() == reflect.Ptr {
		sessionVal = sessionVal.Elem()
	}
	
	snippetVal := reflect.ValueOf(snippetInterface)
	if snippetVal.Kind() == reflect.Ptr {
		snippetVal = snippetVal.Elem()
	}
	
	// Extract session data
	sessionID := sessionVal.FieldByName("ID").String()
	snippetID := sessionVal.FieldByName("SnippetID").String()
	startTime := sessionVal.FieldByName("StartTime").Int()
	endTime := sessionVal.FieldByName("EndTime").Int()
	typedCharsField := sessionVal.FieldByName("TypedCharacters")
	errorsField := sessionVal.FieldByName("Errors")
	
	// Extract snippet data
	difficulty := snippetVal.FieldByName("Difficulty").String()
	linesField := snippetVal.FieldByName("Lines")
	
	// Convert typed characters
	typedChars := make([]string, typedCharsField.Len())
	for i := 0; i < typedCharsField.Len(); i++ {
		typedChars[i] = typedCharsField.Index(i).String()
	}
	
	// Convert lines
	lines := make([]string, linesField.Len())
	for i := 0; i < linesField.Len(); i++ {
		lines[i] = linesField.Index(i).String()
	}
	
	// Convert errors
	errors := make([]TypingError, errorsField.Len())
	for i := 0; i < errorsField.Len(); i++ {
		errorVal := errorsField.Index(i)
		errors[i] = TypingError{
			Position:  int(errorVal.FieldByName("Position").Int()),
			Expected:  errorVal.FieldByName("Expected").String(),
			Typed:     errorVal.FieldByName("Typed").String(),
			Timestamp: errorVal.FieldByName("Timestamp").Int(),
			Corrected: errorVal.FieldByName("Corrected").Bool(),
		}
	}

	// Join all lines to get the expected text
	expectedText := strings.Join(lines, "\n")
	typedText := strings.Join(typedChars, "")

	// Basic metrics
	totalCharacters := len(expectedText)
	completionTime := endTime - startTime // milliseconds
	totalErrors := len(errors)

	// Calculate correct characters
	correctCharacters := calculateCorrectCharacters(expectedText, typedText)

	// Calculate CPM (Characters Per Minute)
	timeInMinutes := float64(completionTime) / (1000.0 * 60.0)
	cpm := float64(len(typedText)) / timeInMinutes

	// Calculate accuracy
	accuracy := float64(correctCharacters) / float64(totalCharacters) * 100

	// Calculate syntax score (penalty for whitespace/formatting errors)
	syntaxScore := calculateSyntaxScore(expectedText, typedText, errors)

	// Calculate final score with breakdown
	breakdown := calculateScoreBreakdown(cpm, accuracy, syntaxScore, difficulty, totalErrors)

	result := &SessionResult{
		SessionID:         sessionID,
		SnippetID:         snippetID,
		CompletionTime:    completionTime,
		TotalCharacters:   totalCharacters,
		CorrectCharacters: correctCharacters,
		TotalErrors:       totalErrors,
		CPM:               cpm,
		Accuracy:          accuracy,
		SyntaxScore:       syntaxScore,
		FinalScore:        breakdown.BaseScore + breakdown.SpeedBonus + breakdown.AccuracyBonus - breakdown.SyntaxPenalty,
		Breakdown:         breakdown,
	}

	// Apply difficulty multiplier to final score
	result.FinalScore *= breakdown.DifficultyMultiplier

	return result, nil
}

func calculateCorrectCharacters(expected, typed string) int {
	correct := 0
	minLen := int(math.Min(float64(len(expected)), float64(len(typed))))

	for i := 0; i < minLen; i++ {
		if expected[i] == typed[i] {
			correct++
		}
	}

	return correct
}

func calculateSyntaxScore(expected, typed string, errors []TypingError) float64 {
	baseScore := 100.0
	penalty := 0.0

	// Count different types of errors
	whitespaceErrors := 0
	symbolErrors := 0
	indentationErrors := 0

	for _, err := range errors {
		if isWhitespace(err.Expected) || isWhitespace(err.Typed) {
			whitespaceErrors++
		} else if isSymbol(err.Expected) || isSymbol(err.Typed) {
			symbolErrors++
		} else if isIndentationError(expected, err.Position) {
			indentationErrors++
		}
	}

	// Apply penalties
	penalty += float64(whitespaceErrors) * 2.0    // Whitespace errors are important in code
	penalty += float64(symbolErrors) * 3.0       // Symbols are critical
	penalty += float64(indentationErrors) * 4.0  // Indentation is crucial for readability

	return math.Max(0, baseScore-penalty)
}

func calculateScoreBreakdown(cpm, accuracy, syntaxScore float64, difficulty string, totalErrors int) ScoreBreakdown {
	// Base score from accuracy
	baseScore := accuracy * 10 // Max 1000 points for perfect accuracy

	// Speed bonus (CPM above 200 gets bonus points)
	speedBonus := 0.0
	if cpm > 200 {
		speedBonus = (cpm - 200) * 2.0 // 2 points per CPM above 200
	}

	// Accuracy bonus (above 95% gets bonus)
	accuracyBonus := 0.0
	if accuracy > 95 {
		accuracyBonus = (accuracy - 95) * 10.0 // 10 points per percent above 95%
	}

	// Syntax penalty
	syntaxPenalty := 100.0 - syntaxScore

	// Difficulty multiplier
	difficultyMultiplier := 1.0
	switch difficulty {
	case "easy":
		difficultyMultiplier = 1.0
	case "medium":
		difficultyMultiplier = 1.2
	case "hard":
		difficultyMultiplier = 1.5
	}

	return ScoreBreakdown{
		BaseScore:            baseScore,
		SpeedBonus:           speedBonus,
		AccuracyBonus:        accuracyBonus,
		SyntaxPenalty:        syntaxPenalty,
		DifficultyMultiplier: difficultyMultiplier,
	}
}

func isWhitespace(char string) bool {
	if len(char) != 1 {
		return false
	}
	return unicode.IsSpace(rune(char[0]))
}

func isSymbol(char string) bool {
	if len(char) != 1 {
		return false
	}
	r := rune(char[0])
	return unicode.IsPunct(r) || unicode.IsSymbol(r)
}

func isIndentationError(text string, position int) bool {
	if position == 0 {
		return false
	}

	// Check if the error is at the beginning of a line (after a newline)
	for i := position - 1; i >= 0; i-- {
		if text[i] == '\n' {
			return true
		}
		if text[i] != ' ' && text[i] != '\t' {
			return false
		}
	}

	// Check if it's at the very beginning of the text
	for i := 0; i < position && i < len(text); i++ {
		if text[i] != ' ' && text[i] != '\t' {
			return false
		}
	}

	return true
}