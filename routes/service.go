// Package routes exposes features for creating & reading routes
package routes

import (
	"errors"

	"github.com/emilianozublena/microservices/database"
	"github.com/emilianozublena/microservices/routific"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

// ValidationError represents a Bongo validation error
var ValidationError error

// RouteService defines the interface for the driver service
type RouteService interface {
	UpdateOrCreateRoute(d *Route) []error
	ReadRoute(ID bson.ObjectId, r *Route) error
	GetRoutesByDriver(driverID bson.ObjectId) ([]Route, error)
	CreateInRoutific(r *Route, current routific.CurrentRoute) error
}

// Route is the struct that defines a Route entity
type Route struct {
	bongo.DocumentBase `bson:",inline"`
	DriverID           bson.ObjectId `bson:"driver_id" json:"driver_id"`
	CustomerID         bson.ObjectId `bson:"customer_id" json:"customer_id"`
	OrderID            bson.ObjectId `bson:"order_id" json:"order_id"`
	Lat                float64
	Lng                float64
	Solution           routific.VehicleRoutingResponse
}

// Service is the base struct that holds db connection and interfaces the given service
type Service struct {
	database.Bongo
	routific.Service
}

// NewService will handle the creation of a new service
func NewService(conn database.Bongo, routificAPI routific.Service) *Service {
	return &Service{
		conn,
		routificAPI,
	}
}

// UpdateOrCreateRoute tries to persist a given Route struct. It'll check for any validation errors and return if any
func (s *Service) UpdateOrCreateRoute(r *Route) []error {
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
	results := s.Find("routes", bson.M{"driver_id": driverID})
	route := &Route{}
	rs := []Route{}
	for results.Next(route) {
		rs = append(rs, *route)
	}

	return rs, nil
}

// CreateInRoutific will send a HTTP request to Routific API's which will return a valid solution for that route+driver
func (s *Service) CreateInRoutific(r *Route, current routific.CurrentRoute) error {
	destination := routific.DestinationRoute{
		Lat: r.Lat,
		Lng: r.Lng,
	}
	solution, err := s.GetVehicleRoute(r.DriverID, current, destination)
	if err != nil {
		return err
	}
	r.Solution = solution
	s.UpdateOrCreateRoute(r)
	return nil
}

// Validate is a hook that will get triggered by bongo, this method will handle validation
func (r *Route) Validate(*bongo.Collection) []error {
	xe := []error{}
	if r.DriverID == "" {
		xe = append(xe, errors.New("You need to specify the driver ID"))
	}
	if r.CustomerID == "" {
		xe = append(xe, errors.New("You need to specify the customer ID"))
	}
	if r.OrderID == "" {
		xe = append(xe, errors.New("You need to specify the order ID"))
	}
	if r.Lat == 0 {
		xe = append(xe, errors.New("You need to specify the latitude"))
	}
	if r.Lng == 0 {
		xe = append(xe, errors.New("You need to specify the longitude"))
	}

	if len(xe) > 0 {
		return xe
	}

	return nil
}
