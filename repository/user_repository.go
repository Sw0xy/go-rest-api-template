package repository

import (
	"context"
	"errors"

	"github.com/Sw0xy/go-rest-api-template/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUsers(ctx context.Context) ([]*models.User, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
	DeleteUserById(ctx context.Context, id string) error
	UpdateUser(ctx context.Context, user *models.User) error
}

type userRepository struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Database, collection string) UserRepository {
	return &userRepository{
		db: db.Collection(collection),
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {

	_, err := r.db.InsertOne(ctx, user)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {

	filter := bson.D{{Key: "email", Value: email}}
	var user *models.User
	err := r.db.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (r *userRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: objectID}}

	var user *models.User
	err = r.db.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (r *userRepository) GetUsers(ctx context.Context) ([]*models.User, error) {

	filter := bson.D{}
	cursor, err := r.db.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var results []*models.User
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, err
}

func (r *userRepository) DeleteUserById(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: objectID}}

	res, err := r.db.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("id not found")
	}

	return err
}

func (r *userRepository) UpdateUser(ctx context.Context, user *models.User) error {

	/* objectID, err := primitive.ObjectIDFromHex(ctx.Value("user_id").(string))
	if err != nil {
		return err
	} */
	if user.Password != "" {
		encryptedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(user.Password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			return err
		}
		user.Password = string(encryptedPassword)
	}

	filter := bson.D{{Key: "_id", Value: user.Id}}

	res, err := r.db.UpdateOne(ctx, filter, user)
	if err != nil {
		return err
	}

	if res.UpsertedCount == 0 {
		return errors.New("id not found")
	}

	return err
}
