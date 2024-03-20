package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Sw0xy/go-rest-api-template/domain"
	"github.com/Sw0xy/go-rest-api-template/utils"
	log "github.com/sirupsen/logrus"
)

type RefreshTokenController struct {
	RefreshTokenUseCase domain.RefreshTokenUseCase
}

func (rtc *RefreshTokenController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request domain.RefreshTokenRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, refreshToken, err := rtc.RefreshTokenUseCase.RefreshToken(ctx, request)
	if err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshTokenResponse := domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	utils.JSON(w, http.StatusOK, refreshTokenResponse)

}
