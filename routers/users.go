package routers

import (
	"api-golang/controllers"

	"github.com/gin-gonic/gin"
)

func SetUsersRouter(router *gin.RouterGroup) *gin.RouterGroup {
	usersRouter := router.Group("/users") 
	
	usersRouter.POST("/create-user", controllers.CreateUserController)
	usersRouter.GET("/get-users", controllers.GetAllUsersController)
	usersRouter.GET("/get-user-by-id/:id", controllers.GetUserByIdController)
	usersRouter.PUT("/update-user-by-id/:id", controllers.UpdateUserByIdController)
	usersRouter.DELETE("/remove-user-by-id/:id", controllers.RemoveUserByIdController)

	return usersRouter
}