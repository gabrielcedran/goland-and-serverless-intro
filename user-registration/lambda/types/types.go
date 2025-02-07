package types

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
}

func NewUser(registerUser RegisterUser) (*User, error) {
	// []byte(registerUser.Password) -> transforms a string into a slice of bytes
	// second parameter is the salt (the bigger the number, the long the processing to salt the string however more secure)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), 10)
	if err != nil {
		return nil, err
	}

	return &User{
		Username:     registerUser.Username,
		PasswordHash: string(hashedPassword),
	}, nil
}

func ValidatePassword(hashedPassword, plainTextPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainTextPassword))
	return err == nil
}

func CreateToken(user User) string {
	now := time.Now()
	validUntil := now.Add(time.Hour * 1).Unix()

	claims := jwt.MapClaims{
		"user":    user.Username,
		"expires": validUntil,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims, nil)

	secret := "secret" // this should be store in aws secrets or somewhere safe

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return ""
	}

	return tokenString
}
