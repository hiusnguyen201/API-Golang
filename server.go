package main

import (
	"api-golang/initializers"
	"api-golang/middlewares"
	"api-golang/models"
	"api-golang/routers"
	"net/http"
)

func init() {
	initializers.LoadEnvVariables()
	models.ConnectToDatabase()
}


func main() {
	router := routers.InitRouters()

	// Handle 404
	router.NoRoute(middlewares.NotFoundMiddleware)

	http.ListenAndServe(":3000", router)
}