package routes

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/emilianozublena/microservices/database"
	"github.com/emilianozublena/microservices/routific"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

func TestNewService(t *testing.T) {
	//Given we a valid bongo mongodb connection
	conn := &database.Connection{}
	routific := &routific.Routific{}

	//When we try to create a new driver service
	newService := NewService(conn, routific)

	//Then we assert we got the expected one
	expectedType := reflect.TypeOf(&Service{})
	gotType := reflect.TypeOf(newService)
	if gotType != expectedType {
		t.Error("Expected", expectedType, "Got", gotType)
	}
}

func BenchmarkNewService(t *testing.B) {
	conn := &database.Connection{}
	routific := &routific.Routific{}
	for i := 0; i < t.N; i++ {
		NewService(conn, routific)
	}
}

func ExampleNewService() {
	conn := database.Connect()
	routific := &routific.Routific{}
	newService := NewService(conn, routific)
	fmt.Printf("%T", newService)
	// Output: *routes.Service
}

func TestUpdateOrCreateRoute(t *testing.T) {
	//Given we have a valid service pointer and a valid route pointer
	connection := setUp()
	routificAPI := &routific.Routific{}
	testsTable := []struct {
		routePointer       *Route
		mockError          error
		expectedResultType reflect.Type
	}{
		{
			routePointer: &Route{
				DriverID:   bson.NewObjectId(),
				CustomerID: bson.NewObjectId(),
				OrderID:    bson.NewObjectId(),
				Lat:        34.567,
				Lng:        -56.1234,
				Solution:   routific.VehicleRoutingResponse{},
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
	service := NewService(connection, routificAPI)
	var route *Route
	for _, v := range testsTable {
		route = v.routePointer
		result := service.UpdateOrCreateRoute(route)
		resultType := reflect.TypeOf(result)
		//Then we assert we got expected result type back
		if resultType != v.expectedResultType && result != nil {
			t.Error("Expected", v.expectedResultType, "Got", resultType)
		}
	}
	tearDown("routes")
}

func BenchmarkUpdateOrCreateRoute(t *testing.B) {
	conn := setUp()
	routificAPI := &routific.Routific{}
	route := &Route{
		DriverID:   bson.NewObjectId(),
		CustomerID: bson.NewObjectId(),
		OrderID:    bson.NewObjectId(),
		Lat:        34.567,
		Lng:        -56.1234,
		Solution:   routific.VehicleRoutingResponse{},
	}
	service := NewService(conn, routificAPI)
	for i := 0; i < t.N; i++ {
		service.UpdateOrCreateRoute(route)
	}
	tearDown("routes")
}

func ExampleService_UpdateOrCreateRoute() {
	conn := database.Connect()
	routificAPI := &routific.Routific{}
	service := NewService(conn, routificAPI)
	route := &Route{
		DriverID:   bson.NewObjectId(),
		CustomerID: bson.NewObjectId(),
		OrderID:    bson.NewObjectId(),
		Lat:        34.567,
		Lng:        -56.1234,
		Solution:   routific.VehicleRoutingResponse{},
	}
	service.UpdateOrCreateRoute(route)
	fmt.Printf("%T", route.Id)
	tearDown("routes")
	// Output: bson.ObjectId
}

func TestReadRoute(t *testing.T) {
	//Given we have a valid service and route id
	conn := setUp()
	routificAPI := &routific.Routific{}
	service := NewService(conn, routificAPI)
	route := &Route{
		DriverID:   bson.NewObjectId(),
		CustomerID: bson.NewObjectId(),
		OrderID:    bson.NewObjectId(),
		Lat:        34.567,
		Lng:        -56.1234,
		Solution:   routific.VehicleRoutingResponse{},
	}
	service.UpdateOrCreateRoute(route)
	testsTable := []struct {
		routeID        bson.ObjectId
		expectedReturn bson.ObjectId
	}{
		{
			routeID:        route.Id,
			expectedReturn: route.Id,
		},
		{
			routeID: bson.NewObjectId(),
		},
	}
	//When we try to retrieve a route
	r := &Route{}
	for _, v := range testsTable {
		err := service.ReadRoute(v.routeID, r)
		//Then we assert we got valid route back or 404 err
		if r.Id == "" || r.Id != route.Id {
			t.Error("Expected", "", "Got", "")
		}
		if r.Id == "" && (err == nil || reflect.TypeOf(err) != reflect.TypeOf(&bongo.DocumentNotFoundError{})) {
			t.Error("Expected", "DocumentNotFoundError", "Got", reflect.TypeOf(err).String())
		}
	}
	tearDown("routes")
}

func BenchmarkReadRoute(t *testing.B) {
	conn := setUp()
	routificAPI := &routific.Routific{}
	service := NewService(conn, routificAPI)
	route := &Route{
		DriverID:   bson.NewObjectId(),
		CustomerID: bson.NewObjectId(),
		OrderID:    bson.NewObjectId(),
		Lat:        34.567,
		Lng:        -56.1234,
		Solution:   routific.VehicleRoutingResponse{},
	}
	service.UpdateOrCreateRoute(route)
	for i := 0; i < t.N; i++ {
		service.ReadRoute(route.Id, route)
	}
	tearDown("routes")
}

func ExampleService_ReadRoute() {
	conn := database.Connect()
	routificAPI := &routific.Routific{}
	service := NewService(conn, routificAPI)
	exampleSolution := routific.VehicleRoutingResponse{
		Status:          "success",
		TotalTravelTime: 31.983334,
		Solution:        make(map[string][]routific.LocationResponse),
	}
	exampleSolution.Solution["vehicle_1"] = []routific.LocationResponse{
		{
			LocationID:   "1",
			LocationName: "Location name 1",
		},
		{
			LocationID:   "2",
			LocationName: "Location name 2",
		},
	}
	createRoute := &Route{
		DriverID:   bson.NewObjectId(),
		CustomerID: bson.NewObjectId(),
		OrderID:    bson.NewObjectId(),
		Lat:        34.567,
		Lng:        -56.1234,
		Solution:   exampleSolution,
	}
	readRoute := &Route{}
	service.UpdateOrCreateRoute(createRoute)
	service.ReadRoute(createRoute.Id, readRoute)
	tearDown("routes")
	fmt.Println(readRoute.Solution)
	// Output: {success 31.983334 map[vehicle_1:[{1 Location name 1} {2 Location name 2}]]}
}

func TestGetRoutesByDriver(t *testing.T) {
	//Given we have a valid route with driver id in our db and a route service
	conn := setUp()
	routificAPI := &routific.Routific{}
	service := NewService(conn, routificAPI)
	route := &Route{
		DriverID:   bson.NewObjectId(),
		CustomerID: bson.NewObjectId(),
		OrderID:    bson.NewObjectId(),
		Lat:        34.567,
		Lng:        -56.1234,
		Solution:   routific.VehicleRoutingResponse{},
	}
	service.UpdateOrCreateRoute(route)

	testsTable := []struct {
		driverID       bson.ObjectId
		expectedReturn []Route
	}{
		{
			driverID: route.DriverID,
			expectedReturn: []Route{
				*route,
			},
		},
		{
			driverID:       bson.NewObjectId(),
			expectedReturn: []Route{},
		},
	}
	//When we try to retrieve routes related that driver
	for _, v := range testsTable {
		routes, _ := service.GetRoutesByDriver(v.driverID)
		//Then we assert we got the expected routes
		if len(v.expectedReturn) == 0 && len(routes) > 0 {
			t.Error("Expected nil", "Got", routes)
		} else if len(v.expectedReturn) > 0 {
			expectedRoute := v.expectedReturn[0]
			firstRoute := routes[0]
			if expectedRoute.Id != firstRoute.Id {
				t.Error("Expected", "", "Got", "")
			}
		}
	}
	tearDown("routes")
}

func BenchmarkGetRouteByDriver(t *testing.B) {
	conn := setUp()
	routificAPI := &routific.Routific{}
	service := NewService(conn, routificAPI)
	route := &Route{
		DriverID:   bson.NewObjectId(),
		CustomerID: bson.NewObjectId(),
		OrderID:    bson.NewObjectId(),
		Lat:        34.567,
		Lng:        -56.1234,
		Solution:   routific.VehicleRoutingResponse{},
	}
	service.UpdateOrCreateRoute(route)
	for i := 0; i < t.N; i++ {
		service.GetRoutesByDriver(route.DriverID)
	}
	tearDown("routes")
}

func ExampleService_GetRoutesByDriver() {
	conn := database.Connect()
	routificAPI := &routific.Routific{}
	service := NewService(conn, routificAPI)
	route := &Route{
		DriverID:   bson.NewObjectId(),
		CustomerID: bson.NewObjectId(),
		OrderID:    bson.NewObjectId(),
		Lat:        34.567,
		Lng:        -56.1234,
		Solution:   routific.VehicleRoutingResponse{},
	}
	service.UpdateOrCreateRoute(route)
	routes, _ := service.GetRoutesByDriver(route.DriverID)
	firstRoute := routes[0]
	fmt.Println(firstRoute.Lat)
	// Output: 34.567
}

func TestCreateInRoutific(t *testing.T) {
	//Given we have a valid route and service
	conn := database.Connect()
	routificAPI := &routific.MockService{}
	service := NewService(conn, routificAPI)
	route, currentRoute, destinationRoute, responseBody := MockForRoutificService()
	routificAPI.On("GetVehicleRoute", route.DriverID, currentRoute, destinationRoute).Return(responseBody, nil).Once()

	//When we try to create it in routific
	err := service.CreateInRoutific(route, currentRoute)

	//Then we assert we have the solution persisted in our route
	if route.Solution.Status != "success" || err != nil {
		t.Error("Expected", "Solution: ", route.Solution, "Got", "Err:", err)
	}
	tearDown("routes")
}

func BenchmarkCreateInRoutific(t *testing.B) {
	conn := database.Connect()
	routificAPI := &routific.MockService{}
	service := NewService(conn, routificAPI)
	route, currentRoute, destinationRoute, responseBody := MockForRoutificService()
	routificAPI.On("GetVehicleRoute", route.DriverID, currentRoute, destinationRoute).Return(responseBody, nil)
	for i := 0; i < t.N; i++ {
		service.CreateInRoutific(route, currentRoute)
	}
	tearDown("routes")
}

func ExampleService_CreateInRoutific() {
	conn := database.Connect()
	routificAPI := &routific.MockService{}
	service := NewService(conn, routificAPI)
	route, currentRoute, destinationRoute, responseBody := MockForRoutificService()
	routificAPI.On("GetVehicleRoute", route.DriverID, currentRoute, destinationRoute).Return(responseBody, nil).Once()
	service.CreateInRoutific(route, currentRoute)
	fmt.Println(route.Solution)
	tearDown("routes")
	// Output: {success 31.983334 map[vehicle_1:[{1 Location name 1} {2 Location name 2}]]}
}

// setUp will get executed before each test/benchmark gets run
// In this case, it'll only connect to db and return a value of type *database.Connection back
func setUp() *database.Connection {
	return database.Connect()
}

// tearDown will get executed after each tests/benchmark gets run.
// The idea for it is that it'll clean the testing database
func tearDown(name string) {
	conn := database.Connect()
	conn.Delete(name, bson.M{})
}

func MockForRoutificService() (*Route, routific.CurrentRoute, routific.DestinationRoute, routific.VehicleRoutingResponse) {
	route := &Route{
		DriverID:   bson.NewObjectId(),
		CustomerID: bson.NewObjectId(),
		OrderID:    bson.NewObjectId(),
		Lat:        34.567,
		Lng:        -56.1234,
		Solution:   routific.VehicleRoutingResponse{},
	}
	currentRoute := routific.CurrentRoute{
		Lat: 45.675,
		Lng: 34.123,
	}
	destinationRoute := routific.DestinationRoute{
		ID:  string(route.CustomerID) + string(route.DriverID) + string(route.OrderID),
		Lat: route.Lat,
		Lng: route.Lng,
	}
	responseBody := routific.VehicleRoutingResponse{
		Status:          "success",
		TotalTravelTime: 31.983334,
		Solution:        make(map[string][]routific.LocationResponse),
	}
	responseBody.Solution["vehicle_1"] = []routific.LocationResponse{
		{
			LocationID:   "1",
			LocationName: "Location name 1",
		},
		{
			LocationID:   "2",
			LocationName: "Location name 2",
		},
	}

	return route, currentRoute, destinationRoute, responseBody
}
