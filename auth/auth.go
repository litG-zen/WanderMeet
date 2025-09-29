package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAuthToken(userID int) string {
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    userID,
		"iat":        time.Now().Unix(),
		"token_type": "auth-token",
		"exp":        time.Now().Add(JWT_AUTH_TOKEN_EXP_DELTA * time.Second).Unix(),
	})

	fmt.Println(token)

	signedToken, err := token.SignedString([]byte(JWT_AUTH_SECRET_KEY))
	if err != nil {
		return ""
	}
	return signedToken
}

func GenerateRefreshToken(userID int) string {
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    userID,
		"iat":        time.Now().Unix(),
		"token_type": "refresh-token",
		"exp":        time.Now().Add(JWT_REFRESH_TOKEN_EXP_DELTA * time.Second).Unix(),
	})

	fmt.Println(token)

	signedToken, err := token.SignedString([]byte(JWT_AUTH_SECRET_KEY))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return signedToken
}
