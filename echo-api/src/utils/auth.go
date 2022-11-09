package auth

import (
	"fmt"
	"time"
	"strings"
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
)

func getVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Token invalid")
	}

	return "Secred", nil
}

func CreateToken(email string) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["email"] = email

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte("Secred"))
}

func ValidateToken() error {
	tokenString := getToken()

	token, erro := jwt.Parse(tokenString, getVerificationKey)

	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("test")
}

func getToken() string {
	token := "Bearer x"

	if len(strings.Splint(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
 
	return ""
}