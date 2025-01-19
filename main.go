package main

import (
	"net/http"
	"rest_api/project/db"
	"rest_api/project/models"

	"github.com/gin-gonic/gin"
)
func main(){
	db.InitDB()
	server:=gin.Default()
	server.GET("/events",getEvenets)
	server.POST("/events",createEvent)
	server.Run(":8080")
}

func getEvenets(context *gin.Context){
	events:=models.GetAllEvents()
	context.JSON(http.StatusOK,events)//gin.H{"message":"hello"}
}

func createEvent(context *gin.Context){
	var event models.Event
	err:=context.ShouldBindJSON(&event)//work like scan 

	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"could not parse data"})
		return 
	}

	event.Id=1
	event.UserId=1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "created event!","event":event})

}