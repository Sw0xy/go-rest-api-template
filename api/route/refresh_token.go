package route

import (
	"time"

	"github.com/Sw0xy/go-rest-api-template/api/controller"
	"github.com/Sw0xy/go-rest-api-template/repository"
	"github.com/Sw0xy/go-rest-api-template/usecase"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRefreshTokenRouter(r *mux.Router, db *mongo.Database, timeout time.Duration) {
	ur := repository.NewUserRepository(db, "users")
	rtc := &controller.RefreshTokenController{
		RefreshTokenUseCase: usecase.NewRefreshTokenUseCase(ur, timeout),
	}

	r.HandleFunc("/refresh_token", rtc.RefreshToken).Methods("POST")
}
