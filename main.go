package main

import (
	"fmt"
	"go-restapi-v2/database"
	"go-restapi-v2/handlers"
	"go-restapi-v2/repository"
	"go-restapi-v2/service"
	"net/http"
)

func main() {
	db := database.NewMySQLDatabase()
	userRepository := repository.NewRepository(db.Instance)
	userService := service.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	router := http.NewServeMux()

	router.HandleFunc("/users", userHandler.GetAllUsersHandler)

	//http.Handle("/users", userHandler.GetAllUsersHandler)

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Printf("Error al iniciar el servidor: %v\n", err)
	}
}
