package storage

import (
	"context"
	"fmt"
	"log"

	"cms-comment/src/infrastructure/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var MCmeta = database.MCmetas

type StructureObject struct {
	Type        string `bson:"type" 		json:"type"`
	StructureId string `bson:"structure_id" json:"structure_id"`
	Name        string `bson:"name" 		json:"name"`
}

type CommentMetadata struct {
	Id             primitive.ObjectID `bson:"_id" 				json:"_id"`
	ContentId      string             `bson:"content_id"		json: "content_id" `
	ContentTitle   string             `bson:"content_title" 	json:"content_title"`
	ContentType    string             `bson:"content_type" 		json:"content_type"`
	CommentType    string             `bson:"comment_type" 		json:"comment_type"`
	Enable         bool               `bson:"enable" 			json:"enable"`
	RealtimeEnable bool               `bson:"realtime_enable" 	json:"realtime_enable"`
	Structure      []StructureObject  `bson:"structure" 		json:"structure"`
}

func GetCommentMetadat() {
	resutls, err := MCmeta.Find(context.Background(), bson.M{"content_id": "58b5354017dc13551fa89c01"})
	if err != nil {
		log.Fatal(err)
	}

	cmetas := make([]*CommentMetadata, 0)
	if err = resutls.All(context.Background(), &cmetas); err != nil {
		log.Fatal(err)
	}
	fmt.Println(cmetas)

	for _, v := range cmetas {
		fmt.Println(v)
	}
}
