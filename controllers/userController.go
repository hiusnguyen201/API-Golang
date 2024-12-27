package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api-golang/models"
)

func CreateUserController(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Validation error", "stack": err.Error()})
        return
	}

	if err := models.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Create user failed", "stack": err})
        return
    }

	c.JSON(http.StatusCreated, gin.H{"message": "Create user success","data": user})
}

func GetAllUsersController(c *gin.Context) {
	var users []models.User

	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Get all users",
		"data": users,
	})
}

func GetUserByIdController(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Get one user success",
		"data": user,
	})
}

func UpdateUserByIdController(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	var updatedData models.User
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Validation error", "stack": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(updatedData)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Update user success",
		"data": user,
	})
}

func RemoveUserByIdController(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	models.DB.Model(&user).Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Remove user success",
		"data": user,
	})
}
