package main

import (
	// "net/http"
	"rest_api/project/db"
	// "rest_api/project/models"
	// "strconv"
	"rest_api/project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoues(server)
	// server.GET("/events", routes.GetEvents) // GET, POST, PUT, PATCH, DELETE
	// server.GET("/events/:id", routes.GetEventByID)
	// server.POST("/events", routes.CreateEvent)

	server.Run(":8080") // localhost:8080
}

// func getEvents(context *gin.Context) {
// 	events, err := models.GetAllEvents()
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
// 		return
// 	}
// 	context.JSON(http.StatusOK, events)
// }

// func getEvent(context *gin.Context){
// 	id,err:=strconv.ParseInt(context.Param("id"),10,64)
// 	if err!=nil{
// 		context.JSON(http.StatusBadRequest,gin.H{"message":"Invalid event id"})
// 		return
// 	}
// 	event,err:=models.GetEvent(id)
// 	if err!=nil{
// 		context.JSON(http.StatusNotFound,gin.H{"message":"Event not found"})
// 		return
// 	}
// 	context.JSON(http.StatusOK,event)

// }

// func createEvent(context *gin.Context) {
// 	var event models.Event
// 	err := context.ShouldBindJSON(&event)

// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
// 		return
// 	}

// 	event.ID = 1
// 	event.UserID = 1

// 	err = event.Save()

// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
// 		return
// 	}

// 	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
// }

