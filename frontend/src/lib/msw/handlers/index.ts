import { http, HttpResponse, type RequestHandler } from 'msw'

const healthHandler = http.get('http://localhost:8080/api/v1/health', () => {
  return HttpResponse.json({
    message: 'ok',
  })
})

export const handlers: RequestHandler[] = [healthHandler]
