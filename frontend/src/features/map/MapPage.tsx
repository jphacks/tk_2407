import { APIProvider } from '@vis.gl/react-google-maps'
import { CustomMap } from '@/components/CustomMap'
import { useCurrentLocation } from '@/hooks/useCurrentLocation'
import { SpotList } from '@/components/SpotList'
import { useEffect, useState } from 'react'
import apiClient from '@/apiClient'
import { Spot } from '@/api/@types'

export default function MapPage() {
  const [spots, setSpots] = useState<Spot[]>([])
  const [isOpen, setIsOpen] = useState<boolean>(false)
  const currentLocation = useCurrentLocation()
  const API_KEY = process.env.NEXT_PUBLIC_GOOGLE_MAPS_API_KEY as string
  useEffect(() => {
    const fetchSpots = async () => {
      if (currentLocation.error) return
      const spotsResp = await apiClient.api.v1.spots
        .$get({
          query: {
            latitude: currentLocation.value.latitude,
            longitude: currentLocation.value.longitude,
          },
        })
        .then((res) => {
          console.log(res)
          return res
        })
      if (spotsResp.spots) {
        setSpots(spotsResp.spots)
      } else console.error('No spots found')
    }
    fetchSpots()
  }, [currentLocation])
  if (currentLocation.error) {
    return null
  }
  const toggleIsOpen = () => setIsOpen(!isOpen)

  return (
    <div>
      <h1>Map</h1>
      <div>
        <SpotList spots={spots} isOpen={true} onToggle={toggleIsOpen} />
        <APIProvider apiKey={API_KEY}>
          <CustomMap spots={spots} centerLocation={currentLocation.value} />
        </APIProvider>
      </div>
    </div>
  )
}
