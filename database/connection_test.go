package database

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

type baseBongoDoc struct {
	bongo.DocumentBase `bson:",inline"`
	SomeField          string
}

// @todo: Write tests for db connection
func TestConnect(t *testing.T) {
	//Given we have an installation of mongodb
	//When we try to connect
	conn := Connect()
	//Then we assert we got a valid struct with conn back
	expected := "*database.Connection"
	got := reflect.TypeOf(conn).String()
	if got != expected {
		t.Error("Expected", expected, "Got", got)
	}
}

func ExampleConnect() {
	conn := Connect()
	fmt.Printf("%T", conn)
	// Output: *database.Connection
}

func TestSave(t *testing.T) {
	//Given we have a collection name and a bongo document
	conn := setUp()
	collectionName := "SomeCollection"
	bongoDoc := &baseBongoDoc{
		SomeField: "SomeValue",
	}
	//When we try to save it
	conn.Save(collectionName, bongoDoc)
	//We assert it got saved
	if bongoDoc.GetId() == "" {
		t.Error("Expected an id, none returned")
	}
	tearDown(collectionName)
}

func BenchmarkSave(t *testing.B) {
	conn := setUp()
	collectionName := "SomeCollection"
	bongoDoc := &baseBongoDoc{
		SomeField: "SomeValue",
	}
	for i := 0; i < t.N; i++ {
		conn.Save(collectionName, bongoDoc)
	}
	tearDown(collectionName)
}

func ExampleConnection_Save() {
	conn := setUp()
	collectionName := "SomeCollection"
	bongoDoc := &baseBongoDoc{
		SomeField: "SomeValue",
	}
	conn.Save(collectionName, bongoDoc)
	fmt.Println(bongoDoc.SomeField)
	tearDown(collectionName)
	// Output: SomeValue
}

func TestFindByID(t *testing.T) {
	//Given we have an existing doc in our collection
	conn := setUp()
	collectionName := "SomeCollection"
	bongoDoc := &baseBongoDoc{
		SomeField: "SomeValue",
	}
	conn.Save(collectionName, bongoDoc)
	//When we try to retrieve by it's id
	retrievedDoc := &baseBongoDoc{}
	conn.FindByID(collectionName, bongoDoc.GetId(), retrievedDoc)
	//Then we assert we got it back
	if retrievedDoc.GetId() != bongoDoc.GetId() {
		t.Error("Expected", bongoDoc.GetId(), "Got", retrievedDoc.GetId())
	}
	tearDown(collectionName)
}

func BenchmarkFindByID(t *testing.B) {
	conn := setUp()
	collectionName := "SomeCollection"
	bongoDoc := &baseBongoDoc{
		SomeField: "SomeValue",
	}
	conn.Save(collectionName, bongoDoc)
	retrievedDoc := &baseBongoDoc{}
	for i := 0; i < t.N; i++ {
		conn.FindByID(collectionName, bongoDoc.GetId(), retrievedDoc)
	}
	tearDown(collectionName)
}

func ExampleConnection_FindByID() {
	conn := setUp()
	collectionName := "SomeCollection"
	bongoDoc := &baseBongoDoc{
		SomeField: "SomeValue",
	}
	conn.Save(collectionName, bongoDoc)
	retrievedDoc := &baseBongoDoc{}
	conn.FindByID(collectionName, bongoDoc.GetId(), retrievedDoc)
	fmt.Println(retrievedDoc.SomeField)
	tearDown(collectionName)
	// Output: SomeValue
}

func TestFind(t *testing.T) {
	//Given we have a valid doc and conncetion
	conn := setUp()
	collectionName := "SomeCollection"
	bongoDoc := &baseBongoDoc{
		SomeField: "SomeValue",
	}
	retrievedDoc := &baseBongoDoc{}
	conn.Save(collectionName, bongoDoc)
	query := bson.M{"SomeField": "SomeValue"}
	//When we try to get it through using a query
	results := conn.Find(collectionName, query)

	//Then we assert we got it back
	for results.Next(retrievedDoc) {
		if bongoDoc.GetId() != retrievedDoc.GetId() {
			t.Error("Expected", bongoDoc.GetId(), "Got", retrievedDoc.GetId())
		}
	}
	tearDown(collectionName)
}

func BenchmarkFind(t *testing.B) {
	conn := setUp()
	collectionName := "SomeCollection"
	bongoDoc := &baseBongoDoc{
		SomeField: "SomeValue",
	}
	conn.Save(collectionName, bongoDoc)
	query := bson.M{"SomeField": "SomeValue"}
	for i := 0; i < t.N; i++ {
		conn.Find(collectionName, query)
	}
	tearDown(collectionName)
}

func ExampleConnection_Find() {
	conn := setUp()
	collectionName := "SomeCollection"
	bongoDoc := &baseBongoDoc{
		SomeField: "SomeValue",
	}
	conn.Save(collectionName, bongoDoc)
	query := bson.M{"SomeField": "SomeValue"}
	results := conn.Find(collectionName, query)
	fmt.Printf("%T", results)
	tearDown(collectionName)
	// Output: *bongo.ResultSet
}

func TestDelete(t *testing.T) {
	//Given we have a valid doc in our db
	conn := setUp()
	collectionName := "SomeCollection"
	bongoDoc := &baseBongoDoc{
		SomeField: "SomeValue",
	}
	conn.Save(collectionName, bongoDoc)
	//When we try to delete it
	changeInfo, _ := conn.Delete(collectionName, bson.M{"_id": bongoDoc.GetId()})
	//Then we assert it's not there anymore
	if changeInfo.Removed < 1 {
		t.Error("Expected removed to be 1 or more but ", changeInfo, "got back")
	}
	tearDown(collectionName)
}

func BenchmarkDelete(t *testing.B) {
	conn := setUp()
	collectionName := "SomeCollection"
	bongoDoc := &baseBongoDoc{
		SomeField: "SomeValue",
	}
	conn.Save(collectionName, bongoDoc)
	for i := 0; i < t.N; i++ {
		conn.Delete(collectionName, bson.M{"_id": bongoDoc.GetId()})
	}
}

func ExampleConnection_Delete() {
	conn := setUp()
	collectionName := "SomeCollection"
	bongoDoc := &baseBongoDoc{
		SomeField: "SomeValue",
	}
	conn.Save(collectionName, bongoDoc)
	changeInfo, _ := conn.Delete(collectionName, bson.M{"_id": bongoDoc.GetId()})
	fmt.Println(changeInfo.Removed)
	// Output: 1
}

// setUp will get executed before each test/benchmark gets run
// In this case, it'll only connect to db and return a value of type *database.Connection back
func setUp() *Connection {
	return Connect()
}

// tearDown will get executed after each tests/benchmark gets run.
// The idea for it is that it'll clean the testing database
func tearDown(name string) {
	conn := Connect()
	conn.Delete(name, bson.M{})
}
