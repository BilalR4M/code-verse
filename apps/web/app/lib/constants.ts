export const DIFFICULTIES = [
  { id: 'easy', name: 'Easy', description: 'Simple code structures and basic syntax' },
  { id: 'medium', name: 'Medium', description: 'Functions, classes, and common patterns' },
  { id: 'hard', name: 'Hard', description: 'Complex algorithms and advanced concepts' },
]

export const CHAR_STATES = {
  UNTYPED: 'untyped',
  CORRECT: 'correct',
  INCORRECT: 'incorrect',
  CURRENT: 'current',
} as const

export type CharState = typeof CHAR_STATES[keyof typeof CHAR_STATES]