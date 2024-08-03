package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *AppServer) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", s.HelloWorldHandler)

	return r
}

func (s *AppServer) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}
