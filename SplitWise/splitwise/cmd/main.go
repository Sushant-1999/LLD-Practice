package main

import (
	"splitwise/handler"
	"splitwise/repository"
	"splitwise/routes"
	"splitwise/service"
)

func main() {

	//Set up Repository
	// Setup Repository
	groupRepo := repository.NewGroupRepository()

	//Set up Service Layer : Main business logic
	groupService := service.NewGroupService(groupRepo)

	//Set up Handler
	groupHandler := handler.NewGroupHandler(groupService)

	//Set up router
	router := routes.SetupRouter(groupHandler)

	//Run the application
	router.Run(":8080")

}
