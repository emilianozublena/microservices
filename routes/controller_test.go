package routes

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/emilianozublena/microservices/database"
	"github.com/emilianozublena/microservices/routific"
)

// @todo: write tests for controller

func TestNewRoutesController(t *testing.T) {
	//Given we have a valid RouteService
	routeService := NewService(&database.Connection{}, &routific.Routific{})
	//When we try to create a new controller
	controller := NewRoutesController(routeService)
	//Then we assert we got a valid struct back
	expected := "*routes.Controller"
	got := reflect.TypeOf(controller).String()
	if got != expected {
		t.Error("Expected", expected, "Got", got)
	}
}

func BenchmarkNewRoutesController(t *testing.B) {
	routeService := NewService(&database.Connection{}, &routific.Routific{})
	for i := 0; i < t.N; i++ {
		NewRoutesController(routeService)
	}
}

func ExampleNewRoutesController() {
	routeService := NewService(&database.Connection{}, &routific.Routific{})
	controller := NewRoutesController(routeService)
	fmt.Printf("%T", controller)
	// Output: *routes.Controller
}
