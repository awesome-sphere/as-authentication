package jwt

import (
	"time"

	"github.com/awesome-sphere/as-authentication/db/models"
	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(user models.User) string {
	claims := &authenticationClaim{
		user.ID,
		user.IsAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // FIXME: change expire time
			Issuer:    ISSUER,
			IssuedAt:  time.Now().Unix(),
			Subject:   user.Username,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return ""
	}
	return t
}
