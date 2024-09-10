package auth

import "github.com/dgrijalva/jwt-go"

type JwtCustomClaims struct {
	UserId string `json:"id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}
