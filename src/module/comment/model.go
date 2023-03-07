package comment

import (
	"time"
)
type Comment struct {
	MetaId      	string    			`json:"meta_id"`
	Content      	string    			`json:"content"`
	ParentId    	string    			`json:"parent_id"`
	CreatedTime 	time.Time 			`json:"created_time"`
	UserName    	string    			`json:"user_name"`
	PinTop      	int       			`json:"pin_top"`
	NLike        	int       			`json:"nlike"`
	NReply       	int       			`json:"nreply"`
	UComment     	string   			`json:"ucomment"`
	ULiked       	int       			`json:"uliked"`
	UReported    	bool      			`json:"ureported"`
	CType        	string   			`json:"ctype"`
	IsHyperlink 	int       			`json:"is_hyperlink"`
	IsAdmin     	int       			`json:"is_admin"`
}

type Option struct{
	Id 		string 		`json:"_id"`
	Name 	string  	`json:"name"`
	Value 	[]string  	`json:"value"`
	Note 	string 		`json:"note"`
}


type StructureObject struct {
	Type        string `json:"type"`
	StructureId string `json:"structure_id"`
	Name        string `json:"name"`
}

type CommentMetadata struct {
	ContentId      string             `json:"content_id" `
	ContentTitle   string             `json:"content_title"`
	ContentType    string             `json:"content_type"`
	CommentType    string             `json:"comment_type"`
	Enable         bool               `json:"enable"`
	RealtimeEnable bool               `json:"realtime_enable"`
	Structure      []StructureObject  `json:"structure"`
}