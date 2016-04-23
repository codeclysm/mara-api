package main

import (
	"net/http"
	"time"

	"github.com/codeclysm/mara-api/auth"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
)

// Server contains the handlers for the rest api
type Server struct {
	Auth auth.Client
}

// New creates a new engine
func (s *Server) New() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	r.POST("/login", s.Login)
	return r
}

// Login exchanges a user+password for a jwt token to be used as an Authorization
// header in the other api calls.
//
//     POST /login HTTP/1.1
//     Content-Type: application/json
//     {"user": "admin", "password": "p4$$w0rd"}
//
//     {"token": "AKAJSKSJJSPSPSKS"}
func (s *Server) Login(c *gin.Context) {
	data := new(auth.LoginData)
	err := c.BindJSON(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := s.Auth.Login(data)
	if err != nil {
		switch err := err.(type) {
		default:
			c.JSON(http.StatusInternalServerError, err.Error())
		case auth.ErrorNotFound:
			c.JSON(http.StatusNotFound, err.Error())
		case auth.ErrorWrongPass:
			c.JSON(http.StatusBadRequest, err.Error())
		}
		return
	}
	c.JSON(http.StatusCreated, gin.H{"token": token})
}
