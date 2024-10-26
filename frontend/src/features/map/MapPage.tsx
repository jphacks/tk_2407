import React from 'react'
import { APIProvider } from '@vis.gl/react-google-maps'
import { CustomMap } from '@/components/CustomMap'

export default function MapPage() {
  const API_KEY = process.env.NEXT_PUBLIC_GOOGLE_MAPS_API_KEY as string
  console.log(API_KEY)
  return (
    <div>
      <APIProvider apiKey={API_KEY}>
        <CustomMap />
      </APIProvider>
    </div>
  )
}
