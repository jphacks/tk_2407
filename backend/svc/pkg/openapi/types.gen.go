// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package openapi

// HealthRes defines model for HealthRes.
type HealthRes struct {
	Message string `json:"message"`
}

// Spot defines model for Spot.
type Spot struct {
	Address          string   `json:"address"`
	GoogleMapId      string   `json:"google_map_id"`
	GoogleMapPlaceId string   `json:"google_map_place_id"`
	Latitude         float64  `json:"latitude"`
	Longitude        float64  `json:"longitude"`
	Name             string   `json:"name"`
	PhotoUrl         string   `json:"photo_url"`
	SpotId           string   `json:"spot_id"`
	Types            []string `json:"types"`
}

// SuccessLocationRes defines model for SuccessLocationRes.
type SuccessLocationRes struct {
	Spots *[]Spot `json:"spots"`
}

// SuccessLoginRes defines model for SuccessLoginRes.
type SuccessLoginRes struct {
	UserId string `json:"userId"`
}

// SuccessMessageCreateRes defines model for SuccessMessageCreateRes.
type SuccessMessageCreateRes struct {
	Content   string `json:"content"`
	MessageId string `json:"messageId"`
	PhotoUrl  string `json:"photoUrl"`
	SpotId    string `json:"spotId"`
	UserId    string `json:"userId"`
}

// SuccessMessageRes defines model for SuccessMessageRes.
type SuccessMessageRes struct {
	Messages []struct {
		AuthorName string `json:"author_name"`
		Content    string `json:"content"`
		Id         string `json:"id"`
		Stamps     *map[string]struct {
			Count     int    `json:"count"`
			IsReacted bool   `json:"is_reacted"`
			Type      string `json:"type"`
		} `json:"stamps"`
	} `json:"messages"`
}

// SuccessSignupRes defines model for SuccessSignupRes.
type SuccessSignupRes struct {
	UserId string `json:"userId"`
}

// SuccessStampRes defines model for SuccessStampRes.
type SuccessStampRes struct {
	Message string `json:"message"`
}

// SuccessUserRes defines model for SuccessUserRes.
type SuccessUserRes struct {
	Email    string `json:"email"`
	UserId   string `json:"userId"`
	Username string `json:"username"`
}

// PostApiV1LoginJSONBody defines parameters for PostApiV1Login.
type PostApiV1LoginJSONBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// PostApiV1MessageJSONBody defines parameters for PostApiV1Message.
type PostApiV1MessageJSONBody struct {
	Content  string `json:"content"`
	PhotoUrl string `json:"photoUrl"`
	SpotId   string `json:"spotId"`
	UserId   string `json:"userId"`
}

// PostApiV1MessageStampJSONBody defines parameters for PostApiV1MessageStamp.
type PostApiV1MessageStampJSONBody struct {
	MessageId string `json:"messageId"`
	StampType string `json:"stampType"`
	UserId    string `json:"userId"`
}

// PostApiV1SignupJSONBody defines parameters for PostApiV1Signup.
type PostApiV1SignupJSONBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

// GetApiV1SpotsParams defines parameters for GetApiV1Spots.
type GetApiV1SpotsParams struct {
	// Longitude The longitude of the location.
	Longitude float64 `form:"longitude" json:"longitude"`

	// Latitude The latitude of the location.
	Latitude float64 `form:"latitude" json:"latitude"`
}

// PostApiV1LoginJSONRequestBody defines body for PostApiV1Login for application/json ContentType.
type PostApiV1LoginJSONRequestBody PostApiV1LoginJSONBody

// PostApiV1MessageJSONRequestBody defines body for PostApiV1Message for application/json ContentType.
type PostApiV1MessageJSONRequestBody PostApiV1MessageJSONBody

// PostApiV1MessageStampJSONRequestBody defines body for PostApiV1MessageStamp for application/json ContentType.
type PostApiV1MessageStampJSONRequestBody PostApiV1MessageStampJSONBody

// PostApiV1SignupJSONRequestBody defines body for PostApiV1Signup for application/json ContentType.
type PostApiV1SignupJSONRequestBody PostApiV1SignupJSONBody
