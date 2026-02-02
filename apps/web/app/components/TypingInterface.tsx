'use client'

import { useState, useEffect, useCallback } from 'react'
import { CodeSnippet, TypingSession, TypingError } from '../../../../packages/shared/types'
import { getCharacterStates } from '../lib/utils'
import { CHAR_STATES } from '../lib/constants'
import { ArrowLeft, RotateCcw } from 'lucide-react'

interface TypingInterfaceProps {
  snippet: CodeSnippet
  session: TypingSession
  onComplete: (typedCharacters: string[], errors: TypingError[], completionTime: number) => void
  onReset: () => void
}

export function TypingInterface({ snippet, session, onComplete, onReset }: TypingInterfaceProps) {
  const [typedText, setTypedText] = useState('')
  const [currentPosition, setCurrentPosition] = useState(0)
  const [errors, setErrors] = useState<TypingError[]>([])
  const [startTime] = useState(Date.now())
  const [isCompleted, setIsCompleted] = useState(false)

  const fullText = snippet.lines.join('\n')
  
  const handleKeyPress = useCallback((event: KeyboardEvent) => {
    if (isCompleted) return

    // Prevent paste operations
    if (event.ctrlKey && (event.key === 'v' || event.key === 'V')) {
      event.preventDefault()
      return
    }

    // Handle backspace
    if (event.key === 'Backspace') {
      event.preventDefault()
      if (typedText.length > 0) {
        setTypedText(prev => prev.slice(0, -1))
        setCurrentPosition(prev => Math.max(0, prev - 1))
      }
      return
    }

    // Handle regular characters
    if (event.key.length === 1 || event.key === 'Tab' || event.key === 'Enter') {
      event.preventDefault()
      
      let keyToAdd = event.key
      if (event.key === 'Enter') keyToAdd = '\n'
      if (event.key === 'Tab') keyToAdd = '\t'
      
      const expectedChar = fullText[currentPosition]
      const newTypedText = typedText + keyToAdd
      
      // Track errors
      if (expectedChar !== keyToAdd) {
        const error: TypingError = {
          position: currentPosition,
          expected: expectedChar,
          typed: keyToAdd,
          timestamp: Date.now(),
          corrected: false
        }
        setErrors(prev => [...prev, error])
      }
      
      setTypedText(newTypedText)
      setCurrentPosition(prev => prev + 1)
      
      // Check if completed
      if (newTypedText === fullText) {
        setIsCompleted(true)
        const completionTime = Date.now() - startTime
        onComplete(newTypedText.split(''), errors, completionTime)
      }
    }
  }, [typedText, currentPosition, fullText, errors, startTime, isCompleted, onComplete])

  useEffect(() => {
    window.addEventListener('keydown', handleKeyPress)
    return () => {
      window.removeEventListener('keydown', handleKeyPress)
    }
  }, [handleKeyPress])

  const characterStates = getCharacterStates(snippet.lines, typedText, currentPosition)
  const accuracy = typedText.length > 0 ? ((typedText.length - errors.length) / typedText.length * 100) : 100
  const timeElapsed = Math.floor((Date.now() - startTime) / 1000)
  const cpm = timeElapsed > 0 ? Math.floor((typedText.length / timeElapsed) * 60) : 0

  return (
    <div className="max-w-6xl mx-auto">
      <div className="flex justify-between items-center mb-6">
        <button
          onClick={onReset}
          className="flex items-center gap-2 px-4 py-2 text-gray-600 hover:text-gray-800 transition-colors"
        >
          <ArrowLeft size={20} />
          Back to snippets
        </button>
        
        <div className="flex gap-6">
          <div className="text-sm text-gray-600">
            <span className="font-medium">Time:</span> {timeElapsed}s
          </div>
          <div className="text-sm text-gray-600">
            <span className="font-medium">CPM:</span> {cpm}
          </div>
          <div className="text-sm text-gray-600">
            <span className="font-medium">Accuracy:</span> {accuracy.toFixed(1)}%
          </div>
        </div>

        <button
          onClick={onReset}
          className="flex items-center gap-2 px-4 py-2 text-gray-600 hover:text-gray-800 transition-colors"
        >
          <RotateCcw size={20} />
          Reset
        </button>
      </div>

      <div className="bg-white rounded-lg shadow-lg border border-gray-200 p-6">
        <div className="mb-4">
          <h2 className="text-xl font-semibold text-gray-900 mb-2">{snippet.title}</h2>
          <p className="text-gray-600">{snippet.description}</p>
        </div>

        <div className="typing-area bg-gray-50 rounded-lg p-6 font-mono text-sm leading-relaxed">
          {snippet.lines.map((line, lineIndex) => {
            const lineStart = snippet.lines.slice(0, lineIndex).join('\n').length + (lineIndex > 0 ? 1 : 0)
            const lineEnd = lineStart + line.length
            const lineChars = characterStates.slice(lineStart, lineEnd + (lineIndex < snippet.lines.length - 1 ? 1 : 0))

            return (
              <div key={lineIndex} className="code-line min-h-[1.5rem]">
                {line.length === 0 ? (
                  <span className="text-gray-400">{'<empty line>'}</span>
                ) : (
                  lineChars.map((charInfo, charIndex) => {
                    if (charInfo.char === '\n' && lineIndex < snippet.lines.length - 1) {
                      return null // Don't render newline characters explicitly
                    }
                    
                    let className = 'char-untyped'
                    switch (charInfo.state) {
                      case CHAR_STATES.CORRECT:
                        className = 'char-correct'
                        break
                      case CHAR_STATES.INCORRECT:
                        className = 'char-incorrect'
                        break
                      case CHAR_STATES.CURRENT:
                        className = 'char-current'
                        break
                    }

                    return (
                      <span key={charIndex} className={className}>
                        {charInfo.char === ' ' ? '·' : charInfo.char}
                      </span>
                    )
                  })
                )}
                {/* Show cursor at the end of line if that's the current position */}
                {currentPosition === lineEnd + (lineIndex < snippet.lines.length - 1 ? 1 : 0) && (
                  <span className="typing-cursor"></span>
                )}
              </div>
            )
          })}
        </div>

        <div className="mt-4 text-sm text-gray-500">
          <p>Type the code exactly as shown. Press Tab for indentation, use spaces where shown.</p>
          <p>Progress: {typedText.length}/{fullText.length} characters</p>
        </div>
      </div>
    </div>
  )
}