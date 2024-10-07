package routes

import (
	"net/http"

	"example.com/rest/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt parse req data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
}
