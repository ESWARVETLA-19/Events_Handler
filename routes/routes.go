package routes

import "github.com/gin-gonic/gin"

func RegisterRoues(server *gin.Engine){
	server.GET("/events", GetEvents) //route to get all events  // GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id", GetEventByID)//route to get event by single id
	server.POST("/events", CreateEvent)//route to create event
	server.PUT("/events/:id",UpdateEvent)//route to update event
	server.DELETE("/events/:id",DeleteEvent)//route to Delete Event
	server.POST("/signup")//User creation Route
}