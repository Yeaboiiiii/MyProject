package routes

import (
	"C/Users/anura/OneDrive/Documents/GitHub/MyProject/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"messagea": "could not parse event id", "error": err.Error()})
		return
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messagea": "could not fetch event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messagea": "could fetch events try again later", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)

}
func createEvents(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	event.UserID = context.GetInt64("userId")
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}

func updateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id", "error": err.Error()})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messagea": "could not fetch event", "error": err.Error()})
		return
	}
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to update event", "event": event})
	}
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	updatedEvent.ID = eventID
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messagea": "could not update event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event updated succesfully", "event": updatedEvent})
}
func deleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"messagea": "could not parse event id", "error": err.Error()})
		return
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event while deleting", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to delete event", "event": event})
	}
	err = models.DeleteEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messagea": "could not delete event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}
