package jwt

import "github.com/dgrijalva/jwt-go"

type authenticationClaim struct {
	UserID  int  `json:"user_id"`
	IsAdmin bool `json:"is_admin"`
	jwt.StandardClaims
}
