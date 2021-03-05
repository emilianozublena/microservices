// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package routific

import (
	mock "github.com/stretchr/testify/mock"
	bson "gopkg.in/mgo.v2/bson"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// GetVehicleRoute provides a mock function with given fields: driverID, currentRoute, destinationRoute
func (_m *MockService) GetVehicleRoute(driverID bson.ObjectId, currentRoute CurrentRoute, destinationRoute DestinationRoute) (VehicleRoutingResponse, error) {
	ret := _m.Called(driverID, currentRoute, destinationRoute)

	var r0 VehicleRoutingResponse
	if rf, ok := ret.Get(0).(func(bson.ObjectId, CurrentRoute, DestinationRoute) VehicleRoutingResponse); ok {
		r0 = rf(driverID, currentRoute, destinationRoute)
	} else {
		r0 = ret.Get(0).(VehicleRoutingResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bson.ObjectId, CurrentRoute, DestinationRoute) error); ok {
		r1 = rf(driverID, currentRoute, destinationRoute)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
