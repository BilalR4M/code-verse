import { CharState, CHAR_STATES } from './constants'

export interface CharacterInfo {
  char: string
  state: CharState
  lineIndex: number
  charIndex: number
}

export function getCharacterStates(
  lines: string[], 
  typedText: string, 
  currentPosition: number
): CharacterInfo[] {
  const fullText = lines.join('\n')
  const characters: CharacterInfo[] = []
  
  let lineIndex = 0
  let charIndex = 0
  
  for (let i = 0; i < fullText.length; i++) {
    const char = fullText[i]
    let state: CharState = CHAR_STATES.UNTYPED
    
    if (i < typedText.length) {
      state = typedText[i] === char ? CHAR_STATES.CORRECT : CHAR_STATES.INCORRECT
    } else if (i === currentPosition) {
      state = CHAR_STATES.CURRENT
    }
    
    characters.push({
      char,
      state,
      lineIndex,
      charIndex,
    })
    
    if (char === '\n') {
      lineIndex++
      charIndex = 0
    } else {
      charIndex++
    }
  }
  
  return characters
}

export function formatTime(milliseconds: number): string {
  const seconds = Math.floor(milliseconds / 1000)
  const minutes = Math.floor(seconds / 60)
  const remainingSeconds = seconds % 60
  
  if (minutes > 0) {
    return `${minutes}m ${remainingSeconds}s`
  }
  return `${remainingSeconds}s`
}

export function formatScore(score: number): string {
  return Math.round(score).toLocaleString()
}

export function getLanguageDisplayName(languageId: string): string {
  const names: Record<string, string> = {
    'python': 'Python',
    'javascript': 'JavaScript',
    'go': 'Go',
  }
  return names[languageId] || languageId
}