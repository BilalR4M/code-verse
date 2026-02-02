import { Language } from '../../../../packages/shared/types'
import { getLanguageDisplayName } from '../lib/utils'

interface LanguageSelectorProps {
  languages: Language[]
  selected: string
  onChange: (language: string) => void
}

export function LanguageSelector({ languages, selected, onChange }: LanguageSelectorProps) {
  return (
    <div className="flex flex-col items-center">
      <label className="text-sm font-medium text-gray-700 mb-2">Language</label>
      <select
        value={selected}
        onChange={(e) => onChange(e.target.value)}
        className="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 text-gray-900 bg-white"
      >
        {languages.map((language) => (
          <option key={language.id} value={language.id} className="text-gray-900 bg-white">
            {getLanguageDisplayName(language.id)}
          </option>
        ))}
      </select>
    </div>
  )
}