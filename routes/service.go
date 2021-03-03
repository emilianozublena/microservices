package routes

import (
	"github.com/emilianozublena/microservices/database"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

// ValidationError represents a Bongo validation error
var ValidationError error

// RouteService defines the interface for the driver service
type RouteService interface {
	CreateRoute(d *Route) []error
	CreateInRoutific(d *Route)
	ReadRoute(ID bson.ObjectId, r *Route) error
	GetRoutesByDriver(driverID bson.ObjectId) ([]Route, error)
}

// Route is the struct that defines a Route entity
type Route struct {
	bongo.DocumentBase `bson:",inline"`
	driverID           bson.ObjectId
	customerID         bson.ObjectId
	orderID            bson.ObjectId
	lat                float64
	lng                float64
	solution           string
}

// Service is the base struct that holds db connection and interfaces the given service
type Service struct {
	database.Bongo
}

// NewService will handle the creation of a new service
func NewService(conn database.Bongo) *Service {
	return &Service{
		conn,
	}
}

// CreateRoute tries to persist a given Route struct. It'll check for any validation errors and return if any
func (s *Service) CreateRoute(r *Route) []error {
	var errors []error
	err := s.Save("routes", r)

	if vErr, ok := err.(*bongo.ValidationError); ok {
		errors = vErr.Errors
	} else if err != nil {
		errors = []error{
			vErr,
		}
	}

	if errors != nil {
		return errors
	}

	return nil
}

// ReadRoute will try to return a Document from mongodb according to the given bson.ObjectId
func (s *Service) ReadRoute(ID bson.ObjectId, r *Route) error {
	return s.FindByID("routes", ID, r)
}

// GetRoutesByDriver will try to return all available routes for a specific driver
func (s *Service) GetRoutesByDriver(driverID bson.ObjectId) ([]Route, error) {
	results := s.Find("routes", bson.M{"driverID": driverID})
	route := &Route{}
	rs := []Route{}
	for results.Next(route) {
		rs = append(rs, *route)
	}

	return rs, nil
}
