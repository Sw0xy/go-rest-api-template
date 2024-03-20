package route

import (
	"time"

	"github.com/Sw0xy/go-rest-api-template/api/controller"
	"github.com/Sw0xy/go-rest-api-template/repository"
	"github.com/Sw0xy/go-rest-api-template/usecase"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSignupRouter(r *mux.Router, db *mongo.Database, timeout time.Duration) {
	ur := repository.NewUserRepository(db, "users")
	sc := controller.SignupController{
		SignupUseCase: usecase.NewSignupUseCase(ur, timeout),
	}

	r.HandleFunc("/signup", sc.Signup).Methods("POST")
}
