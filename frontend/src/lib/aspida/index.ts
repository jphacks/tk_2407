import api from '@/api/$api'
import aspida from '@aspida/fetch'

export const apiClient = api(
  aspida(undefined, { baseURL: process.env.API_BASE_URL })
)
