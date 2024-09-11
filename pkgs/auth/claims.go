package auth

import "github.com/dgrijalva/jwt-go"

var JwtSecretKey = []byte("your-secret-key")

// JwtCustomClaims defines the structure of JWT claims
type JwtCustomClaims struct {
	ID    string `json:"id"`
	Role  string `json:"role"`
	Phone string `json:"phone,omitempty"`
	jwt.StandardClaims
}
