import { DIFFICULTIES } from '../lib/constants'

interface DifficultySelectorProps {
  selected: string
  onChange: (difficulty: string) => void
}

export function DifficultySelector({ selected, onChange }: DifficultySelectorProps) {
  return (
    <div className="flex flex-col items-center">
      <label className="text-sm font-medium text-gray-700 mb-2">Difficulty</label>
      <select
        value={selected}
        onChange={(e) => onChange(e.target.value)}
        className="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 text-gray-900 bg-white"
      >
        {DIFFICULTIES.map((difficulty) => (
          <option key={difficulty.id} value={difficulty.id} className="text-gray-900 bg-white">
            {difficulty.name}
          </option>
        ))}
      </select>
    </div>
  )
}