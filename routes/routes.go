package routes

import (
	"C/Users/anura/OneDrive/Documents/GitHub/MyProject/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	server.POST("/events", middleware.Authenticate, createEvents)
	server.PUT("/events/:id", middleware.Authenticate, updateEvent)
	server.DELETE("/events/:id", middleware.Authenticate, deleteEvent)
	server.POST("/events/:id/register", middleware.Authenticate, registerForEvent)
	server.DELETE("/events/:id/register", middleware.Authenticate, cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)

}
