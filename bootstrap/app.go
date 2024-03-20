package bootstrap

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	DB  *mongo.Database
	Env *Env
}

func Initialize() App {
	env := InitEnv()
	a := &App{}
	a.Env = env
	a.DB = InitDb(a.Env)

	return *a
}
