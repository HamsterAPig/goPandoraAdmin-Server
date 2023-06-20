package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BackstageRouter() http.Handler {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "success",
			})
		})
	}
	return r
}
