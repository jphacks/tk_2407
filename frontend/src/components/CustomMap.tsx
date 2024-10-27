import { Map, Marker } from '@vis.gl/react-google-maps'
import { type Location } from '@/types/Location'
import apiClient from '@/apiClient'
import React, { useEffect, useState } from 'react'
import { CustomMarker } from '@/components/CustomMarker'
import { Spot } from '@/api/@types'

type CustomMapProps = {
  centerLocation: Location
}

export function CustomMap(props: CustomMapProps) {
  const { centerLocation } = props
  console.log('centerLocation', centerLocation)
  const [spots, setSpots] = useState<Spot[]>([])
  useEffect(() => {
    const fetchSpots = async () => {
      const spotsResp = await apiClient.api.v1.spots
        .$get({
          query: {
            latitude: centerLocation.latitude,
            longitude: centerLocation.longitude,
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
  }, [centerLocation])
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
      {spots.map((spot) => (
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
