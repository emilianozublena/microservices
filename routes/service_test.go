package routes

import (
	"errors"
	"reflect"
	"testing"

	"github.com/emilianozublena/microservices/database"
	"github.com/emilianozublena/microservices/mocks"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

func TestNewService(t *testing.T) {
	//Given we a valid bongo mongodb connection
	conn := &database.Connection{}

	//When we try to create a new driver service
	newService := NewService(conn)

	//Then we assert we got the expected one
	expectedType := reflect.TypeOf(&Service{})
	gotType := reflect.TypeOf(newService)
	if gotType != expectedType {
		t.Error("Expected", expectedType, "Got", gotType)
	}
}

func BenchmarkNewService(t *testing.B) {
	conn := &database.Connection{}
	for i := 0; i < t.N; i++ {
		NewService(conn)
	}
}

func TestCreateRoute(t *testing.T) {
	//Given we have a valid service pointer and a valid route pointer
	bongoMock := &mocks.Bongo{}

	testsTable := []struct {
		routePointer       *Route
		mockError          error
		expectedResultType reflect.Type
	}{
		{
			routePointer: &Route{
				driverID:   bson.NewObjectId(),
				customerID: bson.NewObjectId(),
				orderID:    bson.NewObjectId(),
				lat:        34.567,
				lng:        -56.1234,
				solution:   "some solution",
			},
			mockError:          nil,
			expectedResultType: reflect.TypeOf(nil),
		}, {
			routePointer:       &Route{},
			mockError:          errors.New("Mock error"),
			expectedResultType: reflect.TypeOf([]error{}),
		},
	}
	//When we try to create it
	for _, v := range testsTable {
		route := v.routePointer
		bongoMock.On("Save", "routes", route).Return(v.mockError).Once()
		service := NewService(bongoMock)
		result := service.CreateRoute(route)
		resultType := reflect.TypeOf(result)
		//Then we assert we got expected result type back
		if resultType != v.expectedResultType && result != nil {
			t.Error("Expected", v.expectedResultType, "Got", resultType)
		}
	}
}

func BenchmarkCreateRoute(t *testing.B) {
	bongoMock := &mocks.Bongo{}
	route := &Route{}
	bongoMock.On("Save", "routes", route).Return(nil)
	service := NewService(bongoMock)
	for i := 0; i < t.N; i++ {
		service.CreateRoute(route)
	}
}

func TestReadRoute(t *testing.T) {
	//Given we have a valid service and route id
	bongoMock := &mocks.Bongo{}
	service := NewService(bongoMock)
	testsTable := []struct {
		routeID        bson.ObjectId
		mockError      error
		expectedReturn reflect.Type
	}{
		{
			routeID:        bson.NewObjectId(),
			mockError:      nil,
			expectedReturn: reflect.TypeOf(nil),
		},
		{
			routeID:        bson.NewObjectId(),
			mockError:      errors.New("Mock error"),
			expectedReturn: reflect.TypeOf(errors.New("404")),
		},
	}
	//When we try to retrieve a route
	r := &Route{}
	for _, v := range testsTable {
		bongoMock.On("FindByID", "routes", v.routeID, r).Return(v.mockError).Once()
		err := service.ReadRoute(v.routeID, r)
		//Then we assert we got valid route back or 404 err
		if reflect.TypeOf(err) != v.expectedReturn {
			t.Error("Expected", v.expectedReturn, "Got", err)
		}
	}
}

func BenchmarkReadRoute(t *testing.B) {
	bongoMock := &mocks.Bongo{}
	route := &Route{}
	routeID := bson.NewObjectId()
	bongoMock.On("FindByID", "routes", routeID, route).Return(nil)
	service := NewService(bongoMock)
	for i := 0; i < t.N; i++ {
		service.ReadRoute(routeID, route)
	}
}

func TestGetRoutesByDriver(t *testing.T) {
	//Given we have a valid driver id, a bongo mock and a route service
	bongoMock := &mocks.Bongo{}
	service := NewService(bongoMock)
	testsTable := []struct {
		driverID        bson.ObjectId
		bongoMockReturn *bongo.ResultSet
		expectedReturn  []Route
	}{
		{
			driverID:        bson.NewObjectId(),
			bongoMockReturn: new(bongo.ResultSet),
			expectedReturn:  []Route{},
		},
		{
			driverID:        bson.NewObjectId(),
			bongoMockReturn: new(bongo.ResultSet),
			expectedReturn:  nil,
		},
	}
	//When we try to retrieve routes related that driver
	for _, v := range testsTable {
		bongoMock.On("Find", "routes", bson.M{"driverID": v.driverID}).Return(v.bongoMockReturn).Once()
		routes, _ := service.GetRoutesByDriver(v.driverID)
		expectedType := reflect.TypeOf(v.expectedReturn)
		gotType := reflect.TypeOf(routes)
		//Then we assert we got the expected routes
		if expectedType != gotType {
			t.Error("Expected", expectedType, "Got", gotType)
		}
	}

}
