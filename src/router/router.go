package router

import (
	// "net/http"
	"github.com/gin-gonic/gin"
	"cms-comment/src/module/comment"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"service": "Api Cms Comment",
			"status":  1,
		})
	})

	initService(router)

	return router
}

func initService(ginRouter *gin.Engine) {
	comment.New("api/v1/").Router(ginRouter)
	
}
