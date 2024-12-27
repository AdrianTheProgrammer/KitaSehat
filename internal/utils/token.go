package utils

import (
	"KitaSehat_Backend/internal/features/users"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenUtilityInterface interface {
	GenerateToken(users.User) (string, error)
	DecodeToken(*jwt.Token) users.User
}

type TokenUtility struct{}

func NewTokenUtility() TokenUtilityInterface {
	return &TokenUtility{}
}

func (tu *TokenUtility) GenerateToken(loginData users.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = loginData.UserID
	claims["username"] = loginData.Username
	claims["email"] = loginData.Email
	claims["phone_number"] = loginData.PhoneNumber
	claims["avatar"] = loginData.Avatar
	claims["access_level"] = loginData.AccessLevel
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	process := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := process.SignedString([]byte(os.Getenv("jwtkey")))
	if err != nil {
		return "", err
	}

	return result, nil
}

func (tu *TokenUtility) DecodeToken(token *jwt.Token) users.User {
	claims := token.Claims.(jwt.MapClaims)

	return users.User{
		UserID:      claims["id"].(int),
		Username:    claims["username"].(string),
		Email:       claims["email"].(string),
		PhoneNumber: claims["phone_number"].(string),
		Avatar:      claims["avatar"].(string),
		AccessLevel: claims["access_level"].(string),
	}
}
