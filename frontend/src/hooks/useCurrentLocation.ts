import { useState, useEffect } from 'react'
import { Location } from '@/types/Location'

type LocationState =
  | {
      value: Location
      error: false
    }
  | {
      message: string
      error: true
    }

export function useCurrentLocation(): LocationState {
  const [locationState, setLocation] = useState<LocationState>({
    message: 'something went wrong',
    error: true,
  })

  useEffect(() => {
    if (!navigator.geolocation) {
      setLocation({
        message: 'Geolocationはこのブラウザでサポートされていません。',
        error: true,
      })
      return
    }

    const watcherId = navigator.geolocation.watchPosition(
      (position: GeolocationPosition) => {
        setLocation({
          value: {
            latitude: position.coords.latitude,
            longitude: position.coords.longitude,
          },
          error: false,
        })
      },
      (error: GeolocationPositionError) => {
        setLocation({
          message: error.message,
          error: true,
        })
      },
      {
        enableHighAccuracy: true,
        timeout: 5000,
      }
    )

    // クリーンアップ関数
    return () => {
      navigator.geolocation.clearWatch(watcherId)
    }
  }, [])

  return locationState
}

export default useCurrentLocation
