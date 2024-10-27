import { Location } from '@/types/Location'
import {
  AdvancedMarker,
  InfoWindow,
  useAdvancedMarkerRef,
} from '@vis.gl/react-google-maps'
import { CustomPin } from '@/components/CustomPin'
import { useState } from 'react'
import { SpeechBubble } from '@/components/SpeechBubble'

export function CustomMarker({
  location,
  title,
  description,
}: {
  location: Location
  title: string
  description: string
}) {
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
          style={{ width: 'auto', height: 'auto' }}
          headerContent={
            <h2 className="text-xl font-bold mb-2 text-gray-900">{title}</h2>
          }
        >
          <SpeechBubble description={description} />
        </InfoWindow>
      )}
    </AdvancedMarker>
  )
}
