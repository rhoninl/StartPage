package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	privateKey = "myStartPage"
	MaxAge     = 3 * 30 * 24 * 60 * 60 //3 Months
)

type CustomClaims struct {
	UserId int64
	jwt.StandardClaims
}

func SetToken(userId int64) (string, error) {
	c := &CustomClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(MaxAge) * time.Second).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString([]byte(privateKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CheckToken(tokenString string) (int64, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Wrong Method:%v", token.Header["alg"])
		}
		return []byte(privateKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		return claims.UserId, nil
	}
	return 0, fmt.Errorf("token is Invalid Or Up to date")
}
func DeleteToken() (string, error) {
	c := &CustomClaims{
		UserId: 0,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: -1,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString([]byte(privateKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
