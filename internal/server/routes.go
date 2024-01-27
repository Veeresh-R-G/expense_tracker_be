package server

import (
	"backend/model"
	helper "backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)
	r.GET("/users", s.HelloUsers)
	r.GET("/getAllUsers", s.FindAllUsers)
	r.POST("/add", s.AddUser)

	return r
}

func (s *Server) FindAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.GetUsers(c))

}

func (s *Server) AddUser(c *gin.Context) {
	var user model.Users
	c.BindJSON(&user)
	s.db.InsertUser(user, c)
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
