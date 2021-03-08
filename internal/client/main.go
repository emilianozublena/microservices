package client

import (
	"time"

	routesgrpc "github.com/emilianozublena/microservices/api/grpc/v1/routes"
	"github.com/emilianozublena/microservices/routes"
	"github.com/emilianozublena/microservices/routific"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"
)

var defaultRequestTimeout = time.Second * 10

type grpcService struct {
	grpcClient routesgrpc.RoutesServiceClient
}

// NewGRPCService creates a new gRPC user service connection using the specified connection string.
func NewGRPCService(connString string) (routes.RouteService, error) {
	conn, err := grpc.Dial(connString, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &grpcService{grpcClient: routesgrpc.NewRoutesServiceClient(conn)}, nil
}
func (s *grpcService) UpdateOrCreateRoute(d *routes.Route) []error {
	return nil
}

func (s *grpcService) ReadRoute(ID bson.ObjectId, r *routes.Route) error {
	return nil
}

func (s *grpcService) GetRoutesByDriver(driverID bson.ObjectId) ([]routes.Route, error) {
	return []routes.Route{}, nil
}

func (s *grpcService) CreateInRoutific(r *routes.Route, current routific.CurrentRoute) error {
	return nil
}
