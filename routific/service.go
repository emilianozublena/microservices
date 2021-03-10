// Package routific exposes a wrapper for making calls to Routific's Engine API
package routific

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

// VehicleRoutingAPIUrl is a constant that holds the url for vehicle routing problem solver
const VehicleRoutingAPIUrl = "https://api.routific.com/v1/vrp"
const accessToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfaWQiOiI1MzEzZDZiYTNiMDBkMzA4MDA2ZTliOGEiLCJpYXQiOjEzOTM4MDkwODJ9.PR5qTHsqPogeIIe0NyH2oheaGR-SJXDsxPTcUQNq90E"

// Service interface has all the methods needed for communicating to Routific's Engine API
type Service interface {
	GetVehicleRoute(driverID bson.ObjectId, currentRoute CurrentRoute, destinationRoute DestinationRoute) (VehicleRoutingResponse, error)
}

// Routific struct will hold http client and implement API interface
type Routific struct {
	Client HTTPClient
}

//Position represents any given position as lat+lng
type Position struct {
	ID  string
	Lat float64
	Lng float64
}

// CurrentRoute represents the current location as lat+lng for a given driver
type CurrentRoute Position

// DestinationRoute represents a destination route as lat+lng
type DestinationRoute Position

// VehicleRoutingResponse holds a valid response from Vehicle routing Routific Engine API
type VehicleRoutingResponse struct {
	Status          string                        `json:"status"`
	TotalTravelTime float64                       `bson:"total_travel_time" json:"total_travel_time"`
	Solution        map[string][]LocationResponse `json:"solution"`
}

// LocationResponse struct holds each location in a given solution for a driver to go through
type LocationResponse struct {
	LocationID   string `bson:"location_id" json:"location_id"`
	LocationName string `bson:"location_name" json:"location_name"`
}

// VehicleRoutingRequest struct holds values for a vehicle routing request
type VehicleRoutingRequest struct {
	Visits map[string]OrderRequest   `json:"visits"`
	Fleet  map[string]VehicleRequest `json:"fleet"`
}

// OrderRequest represents the order for a given vehicle routing reequest
type OrderRequest struct {
	Location Position `json:"location"`
}

// VehicleRequest represents the vehicle for a given vehicle routing reequest
type VehicleRequest struct {
	StartLocation Position `json:"start_location"`
}

// NewService will return a new value of type interface Service
func NewService(client HTTPClient) *Routific {
	return &Routific{
		Client: client,
	}
}

// GetVehicleRoute will make a call to Routific's Engine API Vehicle routing endpoint and try to get a valid solution back
func (r *Routific) GetVehicleRoute(driverID bson.ObjectId, currentRoute CurrentRoute, destinationRoute DestinationRoute) (VehicleRoutingResponse, error) {
	requestBody := VehicleRoutingRequest{
		Visits: make(map[string]OrderRequest),
		Fleet:  make(map[string]VehicleRequest),
	}
	requestBody.Visits[destinationRoute.ID] = OrderRequest{
		Location: Position{
			Lat: destinationRoute.Lat,
			Lng: destinationRoute.Lng,
		},
	}
	requestBody.Fleet[currentRoute.ID] = VehicleRequest{
		StartLocation: Position{
			Lat: currentRoute.Lat,
			Lng: currentRoute.Lng,
		},
	}
	jsonBytes, err := json.Marshal(requestBody)
	if err != nil {
		return VehicleRoutingResponse{}, err
	}

	httpRequest, _ := http.NewRequest("POST", VehicleRoutingAPIUrl, bytes.NewReader(jsonBytes))
	resp, err := r.Client.Do(httpRequest)
	//resp, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return VehicleRoutingResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	vehicleRoutingResponse := &VehicleRoutingResponse{}

	err = json.Unmarshal(body, vehicleRoutingResponse)

	if err != nil {
		return VehicleRoutingResponse{}, err
	}

	return *vehicleRoutingResponse, nil
}
