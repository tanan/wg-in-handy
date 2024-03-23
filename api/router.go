package api

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	h := NewHandler()
	r.GET("/_health", h.Health)
	v1 := r.Group("/v1")
	{
		v1.GET("/interface", h.ShowInterface)
		v1.GET("/users", h.GetUsers)
		v1.POST("/users", h.CreateUser)
		v1.GET("/users/:user", h.GetUser)
		v1.GET("/users/:user/download", h.DownloadUserConfig)
	}
	return r
}
