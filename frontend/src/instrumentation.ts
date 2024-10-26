export async function register() {
  if (
    process.env.NEXT_RUNTIME === 'nodejs' &&
    process.env.MOCK_ENABLED === 'true'
  ) {
    const { initMocks } = await import('./lib/msw')
    initMocks()
  }
}
