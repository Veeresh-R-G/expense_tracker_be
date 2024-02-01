package server

import (
	"backend/model"
	helper "backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(corsMiddleware())
	// r.Use(middleware.JWTMiddleware())
	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)
	r.GET("/getAllUsers", s.FindAllUsers)
	r.GET("/getUser/:id", s.FindUserbyUUID)
	r.POST("/add", s.AddUser)
	r.POST("/login", s.Login)
	r.PATCH("/update/dailySpends", s.UpdateDailySpends)
	r.PATCH("/update/totalAmout", s.UpdateAmout)

	return r
}

func (s *Server) FindUserbyUUID(c *gin.Context) {
	tokenString, exists := c.Request.Header["Authorization"]
	if !exists || len(tokenString) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
		return
	}

	id := c.Param("id")
	s.db.GetUserbyUUID(id, c)
}

func (s *Server) Login(c *gin.Context) {
	var User model.Users
	c.BindJSON(&User)
	s.db.LoginUser(User, c)
}

func (s *Server) UpdateDailySpends(c *gin.Context) {
	var User model.Users
	c.BindJSON(&User)
	s.db.UpdateDailySpends(&User, c)
}

func (s *Server) UpdateAmout(c *gin.Context) {
	var User model.Users
	c.BindJSON(&User)
	s.db.UpdateTotalAmount(&User, c)
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
