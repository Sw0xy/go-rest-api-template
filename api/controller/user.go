package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sw0xy/go-rest-api-template/domain"
	"github.com/Sw0xy/go-rest-api-template/models"
	"github.com/Sw0xy/go-rest-api-template/utils"
	log "github.com/sirupsen/logrus"
)

type UserController struct {
	UserUseCase domain.UserUseCase
}

func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "user_id", r.Context().Value("user_id"))

	users, err := uc.UserUseCase.GetUsers(ctx)
	if err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	utils.JSON(w, http.StatusOK, users)
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "user_id", r.Context().Value("user_id"))

	id := fmt.Sprintf("%v", ctx.Value("user_id"))

	user, err := uc.UserUseCase.GetUserById(ctx, id)
	if err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	utils.JSON(w, http.StatusOK, user)

}

func (uc *UserController) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "user_id", r.Context().Value("user_id"))

	id := fmt.Sprintf("%v", ctx.Value("user_id"))

	err := uc.UserUseCase.DeleteUserById(ctx, id)
	if err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	utils.JSON(w, http.StatusOK, true)

}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "user_id", r.Context().Value("user_id"))

	var user *models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err := uc.UserUseCase.UpdateUser(ctx, user)
	if err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	utils.JSON(w, http.StatusOK, "Success")
}
