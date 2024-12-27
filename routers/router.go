package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine { 
	router := gin.Default();
	apiRouters := router.Group("/api/v1") 
	{
		SetUsersRouter(apiRouters)
	}

	return router
}