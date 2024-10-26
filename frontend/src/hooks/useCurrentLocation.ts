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

type Options = PositionOptions

export function useCurrentLocation(options: Options = {}): LocationState {
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

    const watcher = navigator.geolocation.watchPosition(
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
      options
    )

    // クリーンアップ関数
    return () => {
      navigator.geolocation.clearWatch(watcher)
    }
  }, [options])

  return locationState
}

export default useCurrentLocation
