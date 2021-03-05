package routes

import (
	"fmt"
	"net"

	routesgrpc "github.com/emilianozublena/microservices/api/grpc/v1/routes"
	"github.com/emilianozublena/microservices/database"
	"github.com/emilianozublena/microservices/internal"
	"github.com/emilianozublena/microservices/routific"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Init will initialize the gRPC server for this service
func Init() {
	// configure our core service
	grpcAddress := internal.GetEnv("GRPC_ADDR", "localhost:9999")
	conn := database.Connect()
	routificAPI := routific.NewService(&routific.Client{})
	routesService := NewService(conn, routificAPI)
	// configure our gRPC service controller
	routesController := NewRoutesController(routesService)
	// start a gRPC server
	server := grpc.NewServer()
	routesgrpc.RegisterRoutesServiceServer(server, routesController)
	reflection.Register(server)
	con, err := net.Listen("tcp", grpcAddress)
	fmt.Println("gRPC server init on address ", grpcAddress)
	if err != nil {
		panic(err)
	}
	err = server.Serve(con)
	if err != nil {
		panic(err)
	}
}
