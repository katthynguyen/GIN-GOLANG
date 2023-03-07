package middleware

import (
	
	"github.com/gin-gonic/gin"

)

type Service interface {
	Middleware() gin.HandlerFunc
}

type service struct{}

// func NewService() Service {
// 	return &service{}
// }