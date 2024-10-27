package google_maps

import (
	"backend/pkg/settings"
	"context"
	"googlemaps.github.io/maps"
	"image"
)

type GoogleMaps struct {
	client maps.Client
}

func NewGoogleMaps() (*GoogleMaps, error) {
	conf := settings.Get()
	c, err := maps.NewClient(maps.WithAPIKey(conf.ThirdParty.GoogleMaps.APIKey))
	if err != nil {
		return nil, err
	}
	return &GoogleMaps{
		client: *c,
	}, nil
}

func (g *GoogleMaps) GetPlacesFromCoords(ctx context.Context, lat, lng float64) ([]maps.PlacesSearchResult, error) {
	// This function should return a list of places near the given coordinates
	req := &maps.NearbySearchRequest{
		Location: &maps.LatLng{
			Lat: lat,
			Lng: lng,
		},
		Radius:   1000,
		Language: "ja-JP",
	}
	search, err := g.client.NearbySearch(ctx, req)
	if err != nil {
		return nil, err
	}
	return search.Results, nil
}

// GetPlacePhoto returns photo data for a given photo reference
// note that `maxWidth` and `maxHeight` are the required parameters
func (g *GoogleMaps) GetPlacePhoto(ctx context.Context, photoRef string, maxWidth, maxHeight uint) (image.Image, error) {
	// This function should return a list of photo URLs for the given place ID
	req := &maps.PlacePhotoRequest{
		PhotoReference: photoRef,
		MaxWidth:       maxWidth,
		MaxHeight:      maxHeight,
	}
	resp, err := g.client.PlacePhoto(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Image()
}
