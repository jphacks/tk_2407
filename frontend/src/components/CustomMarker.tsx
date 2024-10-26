import { Location } from '@/types/Location'
import { AdvancedMarker } from '@vis.gl/react-google-maps'
import { CustomPin } from '@/components/CustomPin'

export function CustomMarker({ location }: { location: Location }) {
  return (
    <AdvancedMarker
      position={{ lat: location.latitude, lng: location.longitude }}
    >
      <CustomPin imageUrl="/saru.png" />
    </AdvancedMarker>
  )
}
