package server

import (
	helper "backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)
	r.GET("/users", s.HelloUsers)
	return r
}

func (s *Server) HelloUsers(c *gin.Context) {
	date := helper.GetCurrentMonthAndYear()
	resp := make(map[string]string)
	resp["message"] = "Hello Users"

	c.JSON(http.StatusOK, date)
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
