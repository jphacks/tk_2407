import { Location } from '@/types/Location'
import {
  AdvancedMarker,
  InfoWindow,
  useAdvancedMarkerRef,
} from '@vis.gl/react-google-maps'
import { CustomPin } from '@/components/CustomPin'
import { useState } from 'react'

export function CustomMarker({ location }: { location: Location }) {
  const [isOpenInfowindow, setIsOpenInfowindow] = useState<boolean>(false)
  const [markerRef, marker] = useAdvancedMarkerRef()

  return (
    <AdvancedMarker
      ref={markerRef}
      onClick={() =>
        setIsOpenInfowindow((isOpenInfowindow) => !isOpenInfowindow)
      }
      position={{ lat: location.latitude, lng: location.longitude }}
      title={'AdvancedMarker that opens an Infowindow when clicked.'}
    >
      <CustomPin imageUrl="/saru.png" />
      {isOpenInfowindow && (
        <InfoWindow
          anchor={marker}
          maxWidth={200}
          onCloseClick={() => setIsOpenInfowindow(false)}
        >
          This is an example for the combined with an Infowindow.
        </InfoWindow>
      )}
    </AdvancedMarker>
  )
}
