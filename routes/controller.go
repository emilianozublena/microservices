package routes

import (
	"context"

	routesgrpc "github.com/emilianozublena/microservices/api/grpc/v1/routes"
	"github.com/emilianozublena/microservices/routific"
	"gopkg.in/mgo.v2/bson"
)

// Controller holds the service and implements the controller interface
type Controller struct {
	RouteService
}

// NewRoutesController will create a new value of type *Controller
func NewRoutesController(service RouteService) *Controller {
	return &Controller{
		service,
	}
}

// CreateRoute will call the inner service that will have all the business logic to serve this gRPC call and will be in charge of constructing Requests and Responses
func (controller *Controller) CreateRoute(ctx context.Context, req *routesgrpc.CreateRouteRequest) (resp *routesgrpc.CreateRouteResponse, err error) {
	route := &Route{
		DriverID:   bson.ObjectId(req.GetRoute().GetDriverId()),
		CustomerID: bson.ObjectId(req.GetRoute().GetCustomerId()),
		OrderID:    bson.ObjectId(req.GetRoute().GetOrderId()),
		Lat:        req.GetRoute().GetLat(),
		Lng:        req.GetRoute().GetLng(),
	}
	current := routific.CurrentRoute{
		Lat: req.GetCurrentPosition().GetLat(),
		Lng: req.GetCurrentPosition().GetLng(),
	}

	err = controller.RouteService.CreateInRoutific(route, current)
	if err != nil {
		return &routesgrpc.CreateRouteResponse{}, err
	}
	grpcRoute := createGrpcRouteFromRoute(route)
	response := &routesgrpc.CreateRouteResponse{
		Route: grpcRoute,
	}
	mapSolutionToRoute(route, grpcRoute)

	return response, nil
}

// ReadRoute will call the inner service that will have all the business logic to serve this gRPC call
func (controller *Controller) ReadRoute(ctx context.Context, req *routesgrpc.ReadRouteRequest) (*routesgrpc.ReadRouteResponse, error) {
	route := &Route{}
	err := controller.RouteService.ReadRoute(bson.ObjectId(req.GetRouteId()), route)
	if err != nil {
		return &routesgrpc.ReadRouteResponse{}, err
	}
	grpcRoute := createGrpcRouteFromRoute(route)
	response := &routesgrpc.ReadRouteResponse{
		Route: grpcRoute,
	}
	mapSolutionToRoute(route, grpcRoute)

	return response, nil
}

// GetRoutesByDriver will call the inner service that will have all the business logic to serve this gRPC call
func (controller *Controller) GetRoutesByDriver(ctx context.Context, req *routesgrpc.GetRoutesByDriverRequest) (*routesgrpc.GetRoutesByDriverResponse, error) {
	xr, err := controller.RouteService.GetRoutesByDriver(bson.ObjectId(req.GetDriverId()))
	if err != nil {
		return &routesgrpc.GetRoutesByDriverResponse{}, err
	}
	grpcXRoutes := createGrpcRoutesFromSlice(xr)
	response := &routesgrpc.GetRoutesByDriverResponse{
		Route: grpcXRoutes,
	}

	return response, nil
}

// createGrpcRouteFromRoute will return the route grpc message after mapping it to the internal Route struct
func createGrpcRouteFromRoute(r *Route) *routesgrpc.Route {
	return &routesgrpc.Route{
		Id:         r.GetId().String(),
		DriverId:   string(r.DriverID),
		CustomerId: string(r.CustomerID),
		OrderId:    string(r.OrderID),
		Lat:        r.Lat,
		Lng:        r.Lng,
		Solution: &routesgrpc.Solution{
			Status:          r.Solution.Status,
			TotalTravelTime: float32(r.Solution.TotalTravelTime),
			Visits:          make(map[string]*routesgrpc.Location),
		},
	}
}

// mapSolutionToRoute will map the solution we have in our internal struct (from db) to the grpc message
func mapSolutionToRoute(r *Route, gr *routesgrpc.Route) {
	for k, v := range r.Solution.Solution {
		for _, solution := range v {
			gr.Solution.Visits[k] = &routesgrpc.Location{
				LocationId:   solution.LocationID,
				LocationName: solution.LocationName,
			}
		}
	}
}

// createGrpcRoutesFromSlice will create a slice of *routesgrpc.Route after mapping all values from []Route
func createGrpcRoutesFromSlice(xr []Route) []*routesgrpc.Route {
	response := []*routesgrpc.Route{}
	for _, v := range xr {
		r := createGrpcRouteFromRoute(&v)
		mapSolutionToRoute(&v, r)
		response = append(response, r)
	}
	return response
}
