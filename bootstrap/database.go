package bootstrap

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitDb(env *Env) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info("Initiate MongoDB connection\n")

	opts := options.Client().ApplyURI(env.MongoUri)

	mongoClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("connection error: %v", err)
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Errorf("ping mongodb error: %v", err)
	}

	log.Info("Successfully connected to MongoDB server(s)\n")

	return mongoClient.Database("backend")
}

func CloseDb(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}

	log.Info("Closed MongoDB connection")
}
