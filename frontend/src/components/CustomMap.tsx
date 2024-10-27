import { Map, Marker } from '@vis.gl/react-google-maps'
import { type Location } from '@/types/Location'
import { CustomMarker } from '@/components/CustomMarker'
import { Spot } from '@/api/@types'

type CustomMapProps = {
  spots: Spot[]
  centerLocation: Location
}

export function CustomMap(props: CustomMapProps) {
  const { centerLocation } = props
  console.log('centerLocation', centerLocation)
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
      mapId={'b1b1b1b1b1b1b1b1'}
    >
      <Marker
        position={{
          lat: centerLocation.latitude,
          lng: centerLocation.longitude,
        }}
      ></Marker>
      {props.spots.map((spot) => (
        <CustomMarker
          key={spot.google_map_place_id}
          imageUrl={spot.photo_url}
          id={spot.google_map_place_id}
          title={spot.name}
          location={spot}
          description=""
        ></CustomMarker>
      ))}
    </Map>
  )
}
