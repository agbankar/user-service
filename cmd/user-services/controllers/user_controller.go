package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"user-service/internal/app_errors"
	"user-service/internal/model"
	"user-service/internal/service"
	"user-service/internal/validations"
	"user-service/respond"
)

type UserController struct {
	UserService    service.UserService
	AppValidations *validations.AppValidations
}

func NewUserController(UserService service.UserService, AppValidator *validations.AppValidations) *UserController {
	return &UserController{
		UserService:    UserService,
		AppValidations: AppValidator,
	}
}

func (u *UserController) GetBonus(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	bonus, _ := u.UserService.CalculateBonus(userId)
	userBonus := model.UserBonus{
		UserID: userId,
		Bonus:  bonus,
	}
	fmt.Println("UserController::GetBonus ", userId)
	respond.With(w, http.StatusOK, userBonus)
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	if !u.AppValidations.ApplyStaticValidations(w, r, user) {
		return
	}
	err := u.UserService.CreateUser(user)
	if err != nil {
		respond.WithErrors(w, http.StatusBadRequest, app_errors.AppError{
			ErrorCode:        "NA",
			ErrorDescription: err.Error(),
			MoreInfo:         "",
		})
	}
	respond.With(w, http.StatusOK, user)

}
