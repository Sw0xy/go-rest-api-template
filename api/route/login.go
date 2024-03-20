package route

import (
	"time"

	"github.com/Sw0xy/go-rest-api-template/api/controller"
	"github.com/Sw0xy/go-rest-api-template/repository"
	"github.com/Sw0xy/go-rest-api-template/usecase"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewLoginRouter(r *mux.Router, db *mongo.Database, timeout time.Duration) {
	ur := repository.NewUserRepository(db, "users")
	lc := &controller.LoginController{
		LoginUseCase: usecase.NewLoginUseCase(ur, timeout),
	}

	r.HandleFunc("/login", lc.Login).Methods("POST")
}
