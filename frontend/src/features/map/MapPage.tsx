import { APIProvider } from '@vis.gl/react-google-maps'
import { CustomMap } from '@/components/CustomMap'
import { useCurrentLocation } from '@/hooks/useCurrentLocation'

export default function MapPage() {
  const API_KEY = process.env.NEXT_PUBLIC_GOOGLE_MAPS_API_KEY as string
  const currentLocation = useCurrentLocation()

  if (currentLocation.error) {
    return null
  }

  return (
    <div>
      <h1>Map</h1>
      <div>
        <APIProvider apiKey={API_KEY}>
          <CustomMap centerLocation={currentLocation.value} />
        </APIProvider>
      </div>
    </div>
  )
}
