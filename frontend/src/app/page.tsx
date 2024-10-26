'use client'

import { APIProvider, Map } from '@vis.gl/react-google-maps'

export default function Home() {
  const API_KEY = process.env.NEXT_PUBLIC_GOOGLE_MAPS_API_KEY as string

  return (
    <div>
      <h1>つむぎ</h1>
      <APIProvider apiKey={API_KEY}>
        <Map
          style={{ width: '100vw', height: '100vh' }}
          defaultCenter={{ lat: 22.54992, lng: 0 }}
          defaultZoom={3}
          gestureHandling={'greedy'}
          disableDefaultUI={true}
        />
      </APIProvider>
    </div>
  )
}
