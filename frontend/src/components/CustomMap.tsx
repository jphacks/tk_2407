import { Map, Marker } from '@vis.gl/react-google-maps'
import { Location } from '@/types/Location'
import { Circle } from '@/components/Circle'

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
      defaultZoom={15}
      gestureHandling={'greedy'}
      disableDefaultUI={false}
    >
      <Marker
        position={{
          lat: centerLocation.latitude,
          lng: centerLocation.longitude,
        }}
      />
      <Circle
        radius={1000}
        center={{ lat: centerLocation.latitude, lng: centerLocation.longitude }}
        strokeColor={'#0c4cb3'}
        strokeOpacity={1}
        strokeWeight={3}
        fillColor={'#3b82f6'}
        fillOpacity={0.3}
      />
    </Map>
  )
}