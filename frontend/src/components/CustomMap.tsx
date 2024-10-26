import { Map } from '@vis.gl/react-google-maps'
import { Location } from '@/types/Location'

type CustomMapProps = {
  centerLocation: Location
}

export function CustomMap(props: CustomMapProps) {
  const { centerLocation } = props
  return (
    <Map
      style={{ width: '100vw', height: '100vh' }}
      defaultCenter={{
        lat: centerLocation.latitude,
        lng: centerLocation.longitude,
      }}
      defaultZoom={3}
      gestureHandling={'greedy'}
      disableDefaultUI={true}
    />
  )
}
