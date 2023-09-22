package handlers

import (
	"encoding/json"
	"go-restapi-v2/service"
	"net/http"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h UserHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	//var users []model.User
	users, err := h.userService.GetAllUsersService()

	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode(users)
}
