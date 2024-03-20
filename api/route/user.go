package route

import (
	"time"

	"github.com/Sw0xy/go-rest-api-template/api/controller"
	"github.com/Sw0xy/go-rest-api-template/repository"
	"github.com/Sw0xy/go-rest-api-template/usecase"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRouter(r *mux.Router, db *mongo.Database, timeout time.Duration) {
	ur := repository.NewUserRepository(db, "users")
	uc := &controller.UserController{
		UserUseCase: usecase.NewUserUseCase(ur, timeout),
	}
	group := r.PathPrefix("/user").Subrouter()
	group.HandleFunc("/all", uc.GetUsers).Methods("GET")
	group.HandleFunc("", uc.GetUserById).Methods("GET")
	group.HandleFunc("", uc.UpdateUser).Methods("PUT")
	group.HandleFunc("", uc.DeleteUserById).Methods("DELETE")

}
