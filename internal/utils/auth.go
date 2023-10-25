package utils

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = "r4hasia"

func GenerateToken(data map[string]interface{}) (string, error) {
	claims := jwt.MapClaims(data)
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(strToken string) error {
	errResponse := errors.New("sign in to process")

	token, _ := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return errResponse
	}

	return nil
}

func HashPassword(p string) (string, error) {
	salt := 8
	password := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(password, salt)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(h, p string) error {
	hash, pass := []byte(h), []byte(p)

	if err := bcrypt.CompareHashAndPassword(hash, pass); err != nil {
		return err
	}

	return nil
}
