package comment

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	
)

type Handler struct {
	CommentVersionAPI string
	service           Service
}

func New(CommentVersionAPI string) *Handler {
	return &Handler{
		CommentVersionAPI: CommentVersionAPI,
		service:           NewService(),
	}
}
func (handler *Handler) Router(route *gin.Engine) {
	group := route.Group(handler.CommentVersionAPI)
	{
		group.GET("/comment",handler.GetCmentById)
	}
}

func (handler *Handler)GetCmentById(c *gin.Context) {
	var request Comment
	resp := handler.service.GetCmentById(context.Background(), &request)
	c.JSON(http.StatusOK, resp)
}