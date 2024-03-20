package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Sw0xy/go-rest-api-template/domain"
	"github.com/Sw0xy/go-rest-api-template/utils"
	"github.com/sirupsen/logrus"
)

type SignupController struct {
	SignupUseCase domain.SignupUseCase
}

func (sc *SignupController) Signup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request domain.SignupRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logrus.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, refreshToken, err := sc.SignupUseCase.SignUp(ctx, request)
	if err != nil {
		logrus.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	utils.JSON(w, http.StatusOK, signupResponse)
}
