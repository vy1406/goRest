package routes

import (
	"example.com/rest/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events/:id", getEvent)
	server.GET("/events", getEvents)

	// assign middleware
	// option 1 ( per group )
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	// assign middleware
	// option 2 ( per route )
	// 	authenticated.POST("/events", middlewares.Authenticate, createEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
