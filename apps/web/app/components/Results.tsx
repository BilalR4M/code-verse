import { SessionResult } from '../../../../packages/shared/types'
import { formatTime, formatScore } from '../lib/utils'
import { RotateCcw, Award, Clock, Target, Zap } from 'lucide-react'

interface ResultsProps {
  results: SessionResult
  onRestart: () => void
}

export function Results({ results, onRestart }: ResultsProps) {
  const {
    completionTime,
    totalCharacters,
    correctCharacters,
    totalErrors,
    cpm,
    accuracy,
    syntaxScore,
    finalScore,
    breakdown
  } = results

  const getScoreColor = (score: number) => {
    if (score >= 90) return 'text-green-600'
    if (score >= 70) return 'text-yellow-600'
    return 'text-red-600'
  }

  return (
    <div className="max-w-4xl mx-auto">
      <div className="bg-white rounded-lg shadow-lg border border-gray-200 p-8">
        <div className="text-center mb-8">
          <div className="flex items-center justify-center gap-3 mb-4">
            <Award className="w-8 h-8 text-yellow-500" />
            <h2 className="text-3xl font-bold text-gray-900">Test Complete!</h2>
          </div>
          <div className="text-6xl font-bold text-blue-600 mb-2">
            {formatScore(finalScore)}
          </div>
          <p className="text-gray-600">Final Score</p>
        </div>

        <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
          <div className="text-center p-4 bg-gray-50 rounded-lg">
            <Clock className="w-6 h-6 text-gray-600 mx-auto mb-2" />
            <div className="text-2xl font-bold text-gray-900">
              {formatTime(completionTime)}
            </div>
            <p className="text-sm text-gray-600">Time</p>
          </div>

          <div className="text-center p-4 bg-gray-50 rounded-lg">
            <Zap className="w-6 h-6 text-blue-600 mx-auto mb-2" />
            <div className="text-2xl font-bold text-blue-600">
              {Math.round(cpm)}
            </div>
            <p className="text-sm text-gray-600">CPM</p>
          </div>

          <div className="text-center p-4 bg-gray-50 rounded-lg">
            <Target className="w-6 h-6 text-green-600 mx-auto mb-2" />
            <div className={`text-2xl font-bold ${getScoreColor(accuracy)}`}>
              {accuracy.toFixed(1)}%
            </div>
            <p className="text-sm text-gray-600">Accuracy</p>
          </div>

          <div className="text-center p-4 bg-gray-50 rounded-lg">
            <Award className="w-6 h-6 text-purple-600 mx-auto mb-2" />
            <div className={`text-2xl font-bold ${getScoreColor(syntaxScore)}`}>
              {syntaxScore.toFixed(0)}
            </div>
            <p className="text-sm text-gray-600">Syntax Score</p>
          </div>
        </div>

        <div className="space-y-4 mb-8">
          <h3 className="text-lg font-semibold text-gray-900">Detailed Breakdown</h3>
          
          <div className="bg-gray-50 rounded-lg p-4">
            <div className="grid md:grid-cols-2 gap-4 text-sm">
              <div>
                <p className="text-gray-600">Characters typed: <span className="font-medium text-gray-900">{totalCharacters}</span></p>
                <p className="text-gray-600">Correct characters: <span className="font-medium text-green-600">{correctCharacters}</span></p>
                <p className="text-gray-600">Total errors: <span className="font-medium text-red-600">{totalErrors}</span></p>
              </div>
              <div>
                <p className="text-gray-600">Base score: <span className="font-medium text-gray-900">+{formatScore(breakdown.baseScore)}</span></p>
                {breakdown.speedBonus > 0 && (
                  <p className="text-gray-600">Speed bonus: <span className="font-medium text-green-600">+{formatScore(breakdown.speedBonus)}</span></p>
                )}
                {breakdown.accuracyBonus > 0 && (
                  <p className="text-gray-600">Accuracy bonus: <span className="font-medium text-green-600">+{formatScore(breakdown.accuracyBonus)}</span></p>
                )}
                {breakdown.syntaxPenalty > 0 && (
                  <p className="text-gray-600">Syntax penalty: <span className="font-medium text-red-600">-{formatScore(breakdown.syntaxPenalty)}</span></p>
                )}
                <p className="text-gray-600">Difficulty multiplier: <span className="font-medium text-blue-600">×{breakdown.difficultyMultiplier}</span></p>
              </div>
            </div>
          </div>
        </div>

        <div className="text-center">
          <button
            onClick={onRestart}
            className="flex items-center gap-2 px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors mx-auto"
          >
            <RotateCcw size={20} />
            Try Another Snippet
          </button>
        </div>
      </div>
    </div>
  )
}