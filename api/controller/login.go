package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Sw0xy/go-rest-api-template/domain"
	"github.com/Sw0xy/go-rest-api-template/utils"
	log "github.com/sirupsen/logrus"
)

type LoginController struct {
	LoginUseCase domain.LoginUseCase
}

func (lc *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request domain.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, refreshToken, err := lc.LoginUseCase.Login(ctx, request)
	if err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	response := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	utils.JSON(w, http.StatusOK, response)

}
