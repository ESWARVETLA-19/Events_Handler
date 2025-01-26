package routes

import (
	"net/http"
	"rest_api/project/models"
	"rest_api/project/utils"

	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context){
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	err=user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created!"})
}

func Login(context *gin.Context){
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	err=user.ValidateCreds()
	if err!=nil{
		context.JSON(http.StatusUnauthorized,err) 
		return
	}
	token,err:=utils.GenerateToken(user.Email,user.ID)
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"user was Not Authenticated"})
	}
	context.JSON(http.StatusOK,gin.H{"token":token})
}

