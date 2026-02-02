export interface Language {
  id: string;
  name: string;
  extension: string;
}

export interface Difficulty {
  id: string;
  name: string;
  description: string;
}

export interface CodeSnippet {
  id: string;
  language: string;
  difficulty: string;
  title: string;
  description: string;
  lines: string[];
  expectedOutput?: string;
  tags: string[];
}

export interface TypingSession {
  id: string;
  snippetId: string;
  startTime: number;
  endTime?: number;
  currentPosition: number;
  typedCharacters: string[];
  errors: TypingError[];
  completed: boolean;
}

export interface TypingError {
  position: number;
  expected: string;
  typed: string;
  timestamp: number;
  corrected: boolean;
}

export interface SessionResult {
  sessionId: string;
  snippetId: string;
  completionTime: number;
  totalCharacters: number;
  correctCharacters: number;
  totalErrors: number;
  cpm: number;
  accuracy: number;
  syntaxScore: number;
  finalScore: number;
  breakdown: ScoreBreakdown;
}

export interface ScoreBreakdown {
  baseScore: number;
  speedBonus: number;
  accuracyBonus: number;
  syntaxPenalty: number;
  difficultyMultiplier: number;
}

export interface StartSessionRequest {
  snippetId: string;
}

export interface FinishSessionRequest {
  sessionId: string;
  typedCharacters: string[];
  errors: TypingError[];
  completionTime: number;
}

export interface ApiResponse<T> {
  success: boolean;
  data?: T;
  error?: string;
}

export interface GetSnippetsParams {
  language?: string;
  difficulty?: string;
  limit?: number;
}