package routific

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"

	mock "github.com/stretchr/testify/mock"
	bson "gopkg.in/mgo.v2/bson"
)

func TestNewService(t *testing.T) {
	//Given there's a valid NewService func
	//When we execute it
	service := NewService(Client)
	//Then we assert we got a valid Service interface TYPE
	serviceTypeAsString := reflect.TypeOf(service).String()
	routificTypeAsString := reflect.TypeOf(&Routific{}).String()
	if serviceTypeAsString != routificTypeAsString {
		t.Error("Expected", routificTypeAsString, "Got", serviceTypeAsString)
	}
	//And we have an underlying Client of interface HTTPClient
	if v, ok := service.Client.(HTTPClient); ok != true {
		t.Error("Expected", "HTTPClient", "Got", v)
	}
}

func BenchmarkNewService(t *testing.B) {
	for i := 0; i < t.N; i++ {
		NewService(Client)
	}
}

func ExampleNewService() {
	service := NewService(Client)
	fmt.Printf("%T", service)
	// Output: *routific.Routific
}

func TestGetVehicleRoute(t *testing.T) {
	//Given we have a valid service, current route & destination route
	httpMock := &MockHTTPClient{}
	service := NewService(httpMock)
	driverID := bson.NewObjectId()
	currentRoute := CurrentRoute{
		Lat: 45.675,
		Lng: 34.123,
	}
	destinationRoute := DestinationRoute{
		Lat: 45.678,
		Lng: 34.123,
	}
	responseBody := VehicleRoutingResponse{
		Status:          "success",
		TotalTravelTime: 31.983334,
		Solution:        make(map[string][]LocationResponse),
	}
	responseBody.Solution["vehicle_1"] = []LocationResponse{
		{
			LocationID:   "1",
			LocationName: "Location name 1",
		},
		{
			LocationID:   "2",
			LocationName: "Location name 2",
		},
	}
	responseBytes, _ := json.Marshal(responseBody)
	httpResponse := &http.Response{
		Body: io.NopCloser(bytes.NewBufferString(string(responseBytes))),
	}
	httpResponse.StatusCode = 200
	httpMock.On("Do", mock.Anything).Return(httpResponse, nil).Once()

	//When we try to get the route from Routific's Engine API
	solution, _ := service.GetVehicleRoute(driverID, currentRoute, destinationRoute)

	//Then we assert we got a valid solution back
	if len(solution.Solution) == 0 || solution.Status != "success" {
		t.Error("Expected at least one solution, none given back", solution)
	}
}
