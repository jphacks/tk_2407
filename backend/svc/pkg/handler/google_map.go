package handler

import (
	"backend/pkg/google_maps"
	"backend/svc/pkg/domain"
	"backend/svc/pkg/openapi"
	"backend/svc/pkg/query"
	"backend/svc/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"googlemaps.github.io/maps"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type GoogleMapHandler struct {
	q      *query.Query
	client google_maps.GoogleMaps
}

func NewGoogleMapHandler(q *query.Query, client google_maps.GoogleMaps) *GoogleMapHandler {
	return &GoogleMapHandler{
		q:      q,
		client: client,
	}
}

func (h *GoogleMapHandler) GetApiV1Spots(c *gin.Context) {
	var err error
	var lon, lat float64
	if paramValue := c.Query("longitude"); paramValue != "" {
		lon, err = strconv.ParseFloat(paramValue, 64)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid format for parameter longitude: %w", err))
			return
		}
	} else {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("query argument longitude is required, but not found"))
		return
	}

	if paramValue := c.Query("latitude"); paramValue != "" {
		lat, err = strconv.ParseFloat(paramValue, 64)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid format for parameter latitude: %w", err))
			return
		}
	} else {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("query argument latitude is required, but not found"))
		return
	}

	results, err := h.client.GetPlacesFromCoords(c, lat, lon)
	if err != nil {
		return
	}

	for _, result := range results {
		fmt.Println(result.Name)
	}

	resps := make([]openapi.Spot, 0)
	for _, result := range results {
		resps = append(resps, openapi.Spot{
			Address:          result.FormattedAddress,
			GoogleMapId:      result.ID,
			GoogleMapPlaceId: result.PlaceID,
			Latitude:         result.Geometry.Location.Lat,
			Longitude:        result.Geometry.Location.Lng,
			Name:             result.Name,
			PhotoUrl:         "",
			Types:            result.Types,
		})
	}
	c.JSON(http.StatusOK, &openapi.SuccessLocationRes{
		Spots: &resps,
	})
	go func(results []maps.PlacesSearchResult) {
		targets := make([]*domain.GmPlace, 0)
		for _, result := range results {
			var rating *float32
			if result.Rating != 0 {
				rating = &result.Rating
			}
			targets = append(targets, &domain.GmPlace{
				ID:                result.ID,
				PlaceID:           result.PlaceID,
				Name:              result.Name,
				FormattedAddress:  result.FormattedAddress,
				Icon:              result.Icon,
				Rating:            rating,
				UserRatingsTotal:  int32(result.UserRatingsTotal),
				PriceLevel:        util.GetPtr(int32(result.PriceLevel)),
				Vicinity:          result.Vicinity,
				PermanentlyClosed: result.PermanentlyClosed,
				BusinessStatus:    result.BusinessStatus,
				LocationLatitude:  result.Geometry.Location.Lat,
				LocationLongitude: result.Geometry.Location.Lng,
				Types:             strings.Join(result.Types, ","),
			})
		}
		err := h.q.GmPlace.Save(targets...)
		if err != nil {
			log.Printf("failed to create gm place: %v", err)
			return
		}
		targetPhotos := make([]*domain.GmPlacePhoto, 0)
		for _, result := range results {
			for _, photo := range result.Photos {
				targetPhotos = append(targetPhotos, &domain.GmPlacePhoto{
					GmPlaceID:      &result.ID,
					PhotoReference: photo.PhotoReference,
					Width:          util.GetPtr(int32(photo.Width)),
					Height:         util.GetPtr(int32(photo.Height)),
				})
			}
		}
		err = h.q.GmPlacePhoto.Save(targetPhotos...)
	}(results)
}