package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = "sipantri-key"

type JwtClaims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(userID uint, role string) (string, error) {

	expireTime := time.Now().Add(time.Duration(1) * 24 * time.Hour)
	claims := &JwtClaims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(strToken string) (jwt.MapClaims, error) {
	errResponse := errors.New("pastikan token kamu masih valid")

	token, err := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	if err != nil{
		return nil, errResponse
	}

	mapClaims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, errResponse
	}

	return mapClaims, nil
}
