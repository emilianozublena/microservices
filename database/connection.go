package database

import (
	"log"
	"os"

	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

// Connection holds bongo conn
type Connection struct {
	connection *bongo.Connection
}

// Bongo represents the interface of Bongo's ODM so we can decouple our code from Bongo itself (this was done to be able to mock bongo methods as well)
type Bongo interface {
	Save(collectionName string, doc bongo.Document) error
	FindByID(collectionName string, id bson.ObjectId, doc interface{}) error
	Find(collectionName string, query interface{}) *bongo.ResultSet
}

// Connect will try to connect to mongodb using Bongo ODM
func Connect() *Connection {
	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	database := os.Getenv("MONGODB_DATABASE")
	config := &bongo.Config{
		ConnectionString: connectionString,
		Database:         database,
	}

	connection, err := bongo.Connect(config)

	if err != nil {
		log.Fatal(err)
	}

	return &Connection{
		connection,
	}
}

// Save will persist a given bongo document in mongodb
func (c *Connection) Save(collectionName string, doc bongo.Document) error {
	return c.connection.Collection(collectionName).Save(doc)
}

// FindByID will retrieve a document by a given id
func (c *Connection) FindByID(collectionName string, id bson.ObjectId, doc interface{}) error {
	return c.connection.Collection(collectionName).FindById(id, doc)
}

// Find will return a *bongo.ResultSet containing all available routes for a single driver by its ObjectID
func (c *Connection) Find(collectionName string, query interface{}) *bongo.ResultSet {
	return c.connection.Collection(collectionName).Find(query)
}
