package comment

import (
	"cms-comment/src/db/storage"
	"cms-comment/src/model"
	"context"
	"net/http"
)
type Service interface {
	GetCmentById(ctx context.Context,request *Comment) interface{}
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) GetCmentById(ctx context.Context, request *Comment) interface{} {

	var comment = storage.GetCmentById()
	if comment == nil{
		return model.BasicResponse{
			Code:    http.StatusNotFound,
			Message: "Not found Comment",
			Data:comment,
		}
	}
	return model.BasicResponse{
		Code:    http.StatusOK,
		Message: "successfully",
		Data: comment,
	}
}
