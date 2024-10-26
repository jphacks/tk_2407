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

  const [lastUpdate, setLastUpdate] = useState<number>(Date.now())

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
        const now = Date.now()
        // Update only if 0.8 seconds (800 ms) has passed since the last update
        // fix me
        if (now - lastUpdate >= 800) {
          setLocation({
            value: {
              latitude: position.coords.latitude,
              longitude: position.coords.longitude,
            },
            error: false,
          })
          setLastUpdate(now) // Update the last update timestamp
        }
      },
      (error: GeolocationPositionError) => {
        setLocation({
          message: error.message,
          error: true,
        })
      },
      options
    )

    // Cleanup function
    return () => {
      navigator.geolocation.clearWatch(watcher)
    }
  }, [options, lastUpdate])

  return locationState
}

export default useCurrentLocation
