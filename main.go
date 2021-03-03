package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(errors.New("mock")).String())
}
