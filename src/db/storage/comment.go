package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"cms-comment/src/infrastructure/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var MCment =  database.MCment

type Comment struct {
	Id          	primitive.ObjectID  `bson:"_id" 			json:"_id"`
	MetaId      	string    			`bson:"meta_id" 		json:"meta_id"`
	Content      	string    			`bson:"content" 		json:"content"`
	ParentId    	string    			`bson:"parent_id" 		json:"parent_id"`
	CreatedTime 	time.Time 			`bson:"created_time" 	json:"created_time"`
	UserName    	string    			`bson:"user_name" 		json:"user_name"`
	PinTop      	int       			`bson:"pin_top" 		json:"pin_top"`
	NLike        	int       			`bson:"nlike" 			json:"nlike"`
	NReply       	int       			`bson:"nreply" 			json:"nreply"`
	UComment     	string   			`bson:"ucomment" 		json:"ucomment"`
	ULiked       	int       			`bson:"uliked" 			json:"uliked"`
	UReported    	bool      			`bson:"ureported" 		json:"ureported"`
	CType        	string   			`bson:"ctype" 			json:"ctype"`
	IsHyperlink 	int       			`bson:"is_hyperlink" 	json:"is_hyperlink"`
	IsAdmin     	int       			`bson:"is_admin" 		json:"is_admin"`
}
 

func GetCmentById() []string{
	resutls, err := MCment.Find(context.Background(), bson.M{"content_id" : "5f042b892089bd009c8934a5"})
	if err != nil {
		log.Fatal(err)
	}
	
	comments := make([]*Comment,0)
	if err = resutls.All(context.Background(), &comments); err != nil {
		log.Fatal(err)
	}
	fmt.Println(comments)
	var data []string
	for _,v := range comments{
		fmt.Println(v)
		data = append(data, v.Content)
	}
	if data == nil {
		return nil
	}
	return data
}
