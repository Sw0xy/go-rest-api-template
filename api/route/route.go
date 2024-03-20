package route

import (
	"os"
	"time"

	"github.com/Sw0xy/go-rest-api-template/api/middleware"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(r *mux.Router, db *mongo.Database, timeout time.Duration) {
	public := r.PathPrefix("/api").Subrouter()
	protectedRouter := r.PathPrefix("/api").Subrouter()

	public.Use(middleware.LoggerMiddleware)
	protectedRouter.Use(middleware.JwtAuthMiddleware(os.Getenv("ACCESS_TOKEN_SECRET")))

	NewSignupRouter(public, db, timeout)
	NewLoginRouter(public, db, timeout)
	NewUserRouter(protectedRouter, db, timeout)
	NewRefreshTokenRouter(public, db, timeout)
}
