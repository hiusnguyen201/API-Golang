package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserController(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusCreated,
		"message": "Create user",
		"data": nil,
	})
}

func GetAllUsersController(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Get all users",
		"data": nil,
	})
}

func GetUserByIdController(context *gin.Context) {
	id := context.Param("id")
	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Get user by id " + id,
		"data": nil,
	})
}

func UpdateUserByIdController(context *gin.Context) {
	id := context.Param("id")
	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Update user by id "+ id,
		"data": nil,
	})
}

func RemoveUserByIdController(context *gin.Context) {
	id := context.Param("id")
	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Remove user by id "+ id,
		"data": nil,
	})
}
