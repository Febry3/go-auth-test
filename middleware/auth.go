package middleware

import (
	"net/http"
	"strings"

	"github.com/febry3/go-auth-test/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		authHeader := ctx.Request().Header.Get("Authorization")

		if authHeader == "" {
			return ctx.String(http.StatusUnauthorized, "token is empty")
		}

		authArr := strings.Split(authHeader, " ")

		if len(authArr) != 2 {
			return ctx.String(http.StatusUnauthorized, "token is invalid")
		}

		var tokenClaim models.AuthClaimsJwt

		token, err := jwt.ParseWithClaims(authArr[1], &tokenClaim, func(t *jwt.Token) (interface{}, error) {
			return []byte("Test"), nil
		})

		if err != nil {
			return ctx.String(http.StatusUnauthorized, err.Error())
		}

		if !token.Valid {
			return ctx.String(http.StatusUnauthorized, "token is not valid")
		}
		ctx.Set("User", tokenClaim)

		return next(ctx)
	}
}
