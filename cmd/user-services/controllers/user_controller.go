package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"user-service/internal/model"
	"user-service/internal/service"
	"user-service/respond"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(UserService service.UserService) *UserController {
	return &UserController{
		UserService: UserService,
	}
}

func (u *UserController) GetBonus(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	bonus, _ := u.UserService.CalculateBonus(userId)
	userBonus := model.UserBonus{
		UserID: userId,
		Bonus:  bonus,
	}
	respond.With(w, r, http.StatusOK, userBonus)
}
