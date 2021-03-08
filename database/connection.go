package database

import (
	"log"

	"github.com/emilianozublena/microservices/internal"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Connection holds bongo conn
type Connection struct {
	connection *bongo.Connection
}

// Bongo represents the interface of Bongo's ODM so we can decouple our code from Bongo itself
type Bongo interface {
	Save(collectionName string, doc bongo.Document) error
	FindByID(collectionName string, id bson.ObjectId, doc interface{}) error
	Find(collectionName string, query interface{}) *bongo.ResultSet
	Delete(collectionName string, query bson.M) (*mgo.ChangeInfo, error)
}

// Connect will try to connect to mongodb using Bongo ODM
func Connect() *Connection {
	connectionString := internal.GetEnv("MONGODB_CONNECTION_STRING", "mongodb://localhost")
	database := internal.GetEnv("MONGODB_DATABASE", "testing")
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

// Delete will remove documents from any given collection
func (c *Connection) Delete(collectionName string, query bson.M) (*mgo.ChangeInfo, error) {
	return c.connection.Collection(collectionName).Delete(query)
}
