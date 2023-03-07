package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongodb struct {
	Client *mongo.Client
}

type Database struct {
	Db *mongo.Database
}

var uri string = "mongodb://localhost:27017" 
// connect to mongodb
func Init() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("Interupted connection to mongodb")
	}
	return client
}

// database 
func Db() *mongo.Database{
	client := Mongodb{Client: Init()}
	mdb := Database{Db: client.Client.Database("fplayux")}
	return mdb.Db
}

// collection
var Mdb = Db()
var MCment =Mdb.Collection("cments_copy")
var MOption = Mdb.Collection("options_copy")
var MCmetas = Mdb.Collection("cmetas_copy")