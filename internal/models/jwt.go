package models

import "github.com/golang-jwt/jwt"

type JWTClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}
