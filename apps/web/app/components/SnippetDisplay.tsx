import { CodeSnippet } from '../../../../packages/shared/types'
import { getLanguageDisplayName } from '../lib/utils'
import { Play } from 'lucide-react'

interface SnippetDisplayProps {
  snippets: CodeSnippet[]
  onStart: (snippet: CodeSnippet) => void
}

export function SnippetDisplay({ snippets, onStart }: SnippetDisplayProps) {
  if (snippets.length === 0) {
    return (
      <div className="text-center py-12">
        <p className="text-gray-500">No snippets available for the selected criteria.</p>
      </div>
    )
  }

  return (
    <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
      {snippets.map((snippet) => (
        <div key={snippet.id} className="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
          <div className="p-6">
            <div className="flex justify-between items-start mb-4">
              <div>
                <h3 className="font-semibold text-gray-900 mb-2">{snippet.title}</h3>
                <p className="text-sm text-gray-600 mb-3">{snippet.description}</p>
                <div className="flex gap-2 mb-4">
                  <span className="px-2 py-1 text-xs bg-blue-100 text-blue-800 rounded">
                    {getLanguageDisplayName(snippet.language)}
                  </span>
                  <span className="px-2 py-1 text-xs bg-gray-100 text-gray-800 rounded capitalize">
                    {snippet.difficulty}
                  </span>
                </div>
              </div>
            </div>
            
            <div className="bg-gray-50 rounded-md p-3 mb-4">
              <pre className="text-sm font-mono text-gray-800 overflow-x-auto">
                {snippet.lines.slice(0, 6).map((line, index) => (
                  <div key={index} className="whitespace-pre">
                    {line || ' '}
                  </div>
                ))}
                {snippet.lines.length > 6 && (
                  <div className="text-gray-500 text-xs mt-2">
                    ...and {snippet.lines.length - 6} more lines
                  </div>
                )}
              </pre>
            </div>

            <div className="flex justify-between items-center">
              <div className="flex gap-1 flex-wrap">
                {snippet.tags.slice(0, 2).map((tag) => (
                  <span key={tag} className="px-2 py-1 text-xs bg-green-100 text-green-800 rounded">
                    {tag}
                  </span>
                ))}
              </div>
              <button
                onClick={() => onStart(snippet)}
                className="flex items-center gap-2 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
              >
                <Play size={16} />
                Start
              </button>
            </div>
          </div>
        </div>
      ))}
    </div>
  )
}