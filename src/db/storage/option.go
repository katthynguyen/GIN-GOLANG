package storage

import (
	"context"
	"fmt"
	"log"
	
	"cms-comment/src/infrastructure/database"
	"go.mongodb.org/mongo-driver/bson"
	
)

var Moption = database.MOption
type Option struct{
	Id 		string 		`bson:"_id" 	json:"_id"`
	Name 	string  	`bson:"name" 	json:"name"`
	Value 	[]string  	`bson:"value" 	json:"value"`
	Note 	string 		`bson:"note" 	json:"note"`
}


func Get_Options() {
	resutls, err := Moption.Find(context.Background(), bson.M{"name": "block_users"})
	if err != nil {
		log.Fatal(err)
	}
	
	options := make([]*Option, 0)
	if err = resutls.All(context.Background(), &options); err != nil {
		log.Fatal(err)
	}

	data := make([]string,0)
	for _, v := range options {
		data = append(data,v.Name)
	}
	fmt.Println(data)
}

