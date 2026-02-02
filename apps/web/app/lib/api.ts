import axios from 'axios'
import { 
  Language, 
  CodeSnippet, 
  TypingSession,
  SessionResult,
  StartSessionRequest,
  FinishSessionRequest,
  GetSnippetsParams,
  ApiResponse 
} from '../../../../packages/shared/types'

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api'

const apiClient = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
})

export const api = {
  async getLanguages(): Promise<ApiResponse<Language[]>> {
    const response = await apiClient.get('/languages')
    return response.data
  },

  async getSnippets(params: GetSnippetsParams): Promise<ApiResponse<CodeSnippet[]>> {
    const response = await apiClient.get('/snippets', { params })
    return response.data
  },

  async startSession(request: StartSessionRequest): Promise<ApiResponse<TypingSession>> {
    const response = await apiClient.post('/sessions/start', request)
    return response.data
  },

  async finishSession(request: FinishSessionRequest): Promise<ApiResponse<SessionResult>> {
    const response = await apiClient.post('/sessions/finish', request)
    return response.data
  },
}