//* The infrastructure layer provides the router, database connection, etc. We'll use the Gin framework for routing.

package routes

import (
	"splitwise/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(groupHandler *handler.GroupHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/groups", groupHandler.GetGroups)
	r.POST("/group/create", groupHandler.CreateGroup)
	return r
}
