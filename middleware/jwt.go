package middleware

import (
	"project/constant"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(id int, role string) (string, error) {
	claim:= jwt.MapClaims{}

	claim["id"] = id
	claim["role"] = role
	claim["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString([]byte(constant.JWT_KEY))
}

func ExtractClaim(e echo.Context) (claims map[string]interface{}) {
	user := e.Get("user").(*jwt.Token)

	if user.Valid {
		claims = user.Claims.(jwt.MapClaims)
	}

	return
}
