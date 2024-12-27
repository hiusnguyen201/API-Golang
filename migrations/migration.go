package main

import (
	"api-golang/initializers"
	"api-golang/models"
)

func init() {
	initializers.LoadEnvVariables()
	models.ConnectToDatabase()
}

func main() {
	models.DB.AutoMigrate(&models.User{})
}