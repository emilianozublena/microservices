package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	routesgrpc "github.com/emilianozublena/microservices/api/grpc/v1/routes"
	"github.com/emilianozublena/microservices/internal"
	"github.com/emilianozublena/microservices/routes"
	"github.com/emilianozublena/microservices/routific"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"
)

var defaultRequestTimeout = time.Second * 10

type grpcService struct {
	grpcClient routesgrpc.RoutesServiceClient
}

// RouteService represents the interface of this gRPC client as a service
type RouteService interface {
	CreateRoute(r *routes.Route) []error
	ReadRoute(ID bson.ObjectId, r *routes.Route) error
	GetRoutesByDriver(driverID bson.ObjectId) ([]routes.Route, error)
}

func main() {
	// configure our core service
	grpcAddress := internal.GetEnv("GRPC_ADDR", "localhost:1435")
	grpcClient, err := NewGRPCService(grpcAddress)
	if err != nil {
		log.Fatal(err)
	}
	r := &routes.Route{
		DriverID:   bson.NewObjectId(),
		CustomerID: bson.NewObjectId(),
		OrderID:    bson.NewObjectId(),
		Lat:        49.2819229,
		Lng:        -123.1211844,
		Solution:   routific.VehicleRoutingResponse{},
	}
	routes := []routes.Route{}
	var mErr []error
	switch selectedMethod := os.Args[1]; selectedMethod {
	case "create":
		mErr = grpcClient.CreateRoute(r)
	case "read":
		err = grpcClient.ReadRoute(bson.NewObjectId(), r)
	case "byDriver":
		routes, err = grpcClient.GetRoutesByDriver(bson.NewObjectId())
	}
	if err != nil {
		log.Fatal(err)
	}
	if mErr != nil {
		log.Fatal(mErr)
	}
	fmt.Println(r)
	fmt.Println(routes)
}

// NewGRPCService creates a new gRPC user service connection using the specified connection string.
func NewGRPCService(connString string) (RouteService, error) {
	conn, err := grpc.Dial(connString, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &grpcService{grpcClient: routesgrpc.NewRoutesServiceClient(conn)}, nil
}

func (s *grpcService) CreateRoute(r *routes.Route) []error {
	routeRequest, positionRequest := createRouteRequest(r)
	req := &routesgrpc.CreateRouteRequest{
		Route:           routeRequest,
		CurrentPosition: positionRequest,
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	resp, err := s.grpcClient.CreateRoute(ctx, req)
	if err != nil {
		return []error{
			err,
		}
	}
	mapRouteResponse(resp.GetRoute(), r)
	return nil
}

func (s *grpcService) ReadRoute(ID bson.ObjectId, r *routes.Route) error {
	req := &routesgrpc.ReadRouteRequest{
		RouteId: []byte(ID),
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	resp, err := s.grpcClient.ReadRoute(ctx, req)
	if err != nil {
		return err
	}
	mapRouteResponse(resp.GetRoute(), r)
	return nil
}

func (s *grpcService) GetRoutesByDriver(driverID bson.ObjectId) ([]routes.Route, error) {
	results := []routes.Route{}
	req := &routesgrpc.GetRoutesByDriverRequest{
		DriverId: []byte(driverID),
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	resp, err := s.grpcClient.GetRoutesByDriver(ctx, req)
	if err != nil {
		return results, err
	}
	for _, v := range resp.GetRoute() {
		results = append(results, createRouteResponse(v))
	}
	return results, nil
}

func createRouteRequest(r *routes.Route) (*routesgrpc.Route, *routesgrpc.Position) {
	route := &routesgrpc.Route{
		Id:         []byte(r.GetId()),
		DriverId:   []byte(r.DriverID),
		CustomerId: []byte(r.CustomerID),
		OrderId:    []byte(r.OrderID),
		Lat:        r.Lat,
		Lng:        r.Lng,
	}
	position := &routesgrpc.Position{
		Lat: 49.2553636,
		Lng: -123.0873365,
	}

	return route, position
}

func mapRouteResponse(r *routesgrpc.Route, route *routes.Route) {
	route.DriverID = bson.ObjectId(r.GetDriverId())
	route.CustomerID = bson.ObjectId(r.GetCustomerId())
	route.OrderID = bson.ObjectId(r.GetOrderId())
	route.Solution = routific.VehicleRoutingResponse{
		Status:          r.GetSolution().GetStatus(),
		TotalTravelTime: float64(r.GetSolution().GetTotalTravelTime()),
		Solution:        make(map[string][]routific.LocationResponse),
	}
	route.SetId(bson.ObjectId(r.GetId()))
	for i, v := range r.GetSolution().GetVisits() {
		route.Solution.Solution[i] = []routific.LocationResponse{
			{
				LocationID:   v.GetLocationId(),
				LocationName: v.GetLocationName(),
			},
		}
	}
}

func createRouteResponse(r *routesgrpc.Route) routes.Route {
	route := &routes.Route{}
	mapRouteResponse(r, route)
	return *route
}
