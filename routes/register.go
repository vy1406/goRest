package routes

import (
	"net/http"
	"strconv"

	"example.com/rest/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Coulndt parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Coulndt fetch event"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt register user for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User registered for event", "event": event})
}

func cancelRegistration(context *gin.Context) {

}
