package routes

import (
	"C/Users/anura/OneDrive/Documents/GitHub/MyProject/models"
	"C/Users/anura/OneDrive/Documents/GitHub/MyProject/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could nnnot parse data", "error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully", "event": user})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data", "error": err.Error()})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized access", "error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "user login successfully", "token": token})

}
