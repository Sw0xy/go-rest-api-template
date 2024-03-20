package utils

import (
	"errors"
	"time"

	"github.com/Sw0xy/go-rest-api-template/domain"
	"github.com/Sw0xy/go-rest-api-template/models"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtCustomClaims struct {
	Id       primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	jwt.RegisteredClaims
}

func CreateAccessToken(user *models.User, secret string, expiry int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claims := &JwtCustomClaims{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func CreateRefreshToken(user *models.User, secret string, expiry int) (refreshToken string, err error) {
	claimsRefresh := &JwtCustomClaims{
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	return token.SignedString([]byte(secret))
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.ErrUnexpectedSigningMethod
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {

	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	idClaim, ok := claims["id"]
	if !ok {
		return "", errors.New("id claim not found")
	}

	id, ok := idClaim.(string)
	if !ok {
		return "", errors.New("id claim is not a string")
	}

	return id, nil
}
