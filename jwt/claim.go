package jwt

import "github.com/dgrijalva/jwt-go"

type authenticationClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       int    `json:"id"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.StandardClaims
}
