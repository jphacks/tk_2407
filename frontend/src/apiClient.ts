import api from '@/api/$api'
import aspida from '@aspida/fetch'

const BASE_URL = process.env.NEXT_PUBLIC_BASE_URL as string

const apiClient = api(
  aspida(undefined, {
    baseURL: BASE_URL,
    credentials: 'include',
  })
)

export default apiClient
