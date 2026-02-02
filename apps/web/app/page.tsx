'use client'

import { useState, useEffect } from 'react'
import { LanguageSelector } from './components/LanguageSelector'
import { DifficultySelector } from './components/DifficultySelector'
import { SnippetDisplay } from './components/SnippetDisplay'
import { TypingInterface } from './components/TypingInterface'
import { Results } from './components/Results'
import { api } from './lib/api'
import { Language, CodeSnippet, SessionResult, TypingSession } from '../../../packages/shared/types'

export default function Home() {
  const [languages, setLanguages] = useState<Language[]>([])
  const [selectedLanguage, setSelectedLanguage] = useState<string>('')
  const [selectedDifficulty, setSelectedDifficulty] = useState<string>('easy')
  const [snippets, setSnippets] = useState<CodeSnippet[]>([])
  const [currentSnippet, setCurrentSnippet] = useState<CodeSnippet | null>(null)
  const [session, setSession] = useState<TypingSession | null>(null)
  const [results, setResults] = useState<SessionResult | null>(null)
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    loadLanguages()
  }, [])

  useEffect(() => {
    if (selectedLanguage) {
      loadSnippets()
    }
  }, [selectedLanguage, selectedDifficulty])

  const loadLanguages = async () => {
    try {
      const response = await api.getLanguages()
      setLanguages(response.data || [])
      if (response.data && response.data.length > 0) {
        setSelectedLanguage(response.data[0].id)
      }
    } catch (err) {
      setError('Failed to load languages')
      console.error(err)
    }
  }

  const loadSnippets = async () => {
    if (!selectedLanguage) return
    
    try {
      setIsLoading(true)
      const response = await api.getSnippets({
        language: selectedLanguage,
        difficulty: selectedDifficulty,
        limit: 10
      })
      setSnippets(response.data || [])
    } catch (err) {
      setError('Failed to load snippets')
      console.error(err)
    } finally {
      setIsLoading(false)
    }
  }

  const startTest = async (snippet: CodeSnippet) => {
    try {
      setIsLoading(true)
      setError(null)
      setResults(null)
      
      const response = await api.startSession({ snippetId: snippet.id })
      setCurrentSnippet(snippet)
      setSession(response.data!)
    } catch (err) {
      setError('Failed to start session')
      console.error(err)
    } finally {
      setIsLoading(false)
    }
  }

  const completeTest = async (typedCharacters: string[], errors: any[], completionTime: number) => {
    if (!session) return

    try {
      setIsLoading(true)
      const response = await api.finishSession({
        sessionId: session.id,
        typedCharacters,
        errors,
        completionTime
      })
      setResults(response.data!)
      setSession(null)
      setCurrentSnippet(null)
    } catch (err) {
      setError('Failed to finish session')
      console.error(err)
    } finally {
      setIsLoading(false)
    }
  }

  const resetTest = () => {
    setSession(null)
    setCurrentSnippet(null)
    setResults(null)
    setError(null)
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <header className="text-center mb-12">
        <h1 className="text-4xl font-bold text-gray-900 mb-4">
          Code Verse
        </h1>
        <p className="text-xl text-gray-600">
          Practice typing real code snippets and improve your coding speed and accuracy
        </p>
      </header>

      {error && (
        <div className="mb-6 p-4 bg-red-100 border border-red-300 text-red-700 rounded-lg">
          {error}
        </div>
      )}

      {!session && !results && (
        <div className="space-y-6">
          <div className="flex gap-4 justify-center">
            <LanguageSelector
              languages={languages}
              selected={selectedLanguage}
              onChange={setSelectedLanguage}
            />
            <DifficultySelector
              selected={selectedDifficulty}
              onChange={setSelectedDifficulty}
            />
          </div>

          {isLoading ? (
            <div className="text-center py-8">
              <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500 mx-auto"></div>
              <p className="mt-4 text-gray-600">Loading snippets...</p>
            </div>
          ) : (
            <SnippetDisplay
              snippets={snippets}
              onStart={startTest}
            />
          )}
        </div>
      )}

      {session && currentSnippet && (
        <TypingInterface
          snippet={currentSnippet}
          session={session}
          onComplete={completeTest}
          onReset={resetTest}
        />
      )}

      {results && (
        <Results
          results={results}
          onRestart={resetTest}
        />
      )}
    </div>
  )
}