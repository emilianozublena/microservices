// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	bongo "github.com/go-bongo/bongo"
	bson "gopkg.in/mgo.v2/bson"

	mock "github.com/stretchr/testify/mock"
)

// Bongo is an autogenerated mock type for the Bongo type
type Bongo struct {
	mock.Mock
}

// Find provides a mock function with given fields: collectionName, query
func (_m *Bongo) Find(collectionName string, query interface{}) *bongo.ResultSet {
	ret := _m.Called(collectionName, query)

	var r0 *bongo.ResultSet
	if rf, ok := ret.Get(0).(func(string, interface{}) *bongo.ResultSet); ok {
		r0 = rf(collectionName, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bongo.ResultSet)
		}
	}

	return r0
}

// FindByID provides a mock function with given fields: collectionName, id, doc
func (_m *Bongo) FindByID(collectionName string, id bson.ObjectId, doc interface{}) error {
	ret := _m.Called(collectionName, id, doc)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bson.ObjectId, interface{}) error); ok {
		r0 = rf(collectionName, id, doc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: collectionName, doc
func (_m *Bongo) Save(collectionName string, doc bongo.Document) error {
	ret := _m.Called(collectionName, doc)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bongo.Document) error); ok {
		r0 = rf(collectionName, doc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
