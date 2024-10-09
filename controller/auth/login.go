package auth

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/febry3/go-auth-test/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string
}

func LoginController(e *echo.Echo, db *sql.DB) {
	e.POST("v1/auth/login", func(ctx echo.Context) error {
		var request LoginRequest

		json.NewDecoder(ctx.Request().Body).Decode(&request)

		row := db.QueryRow("select id, name, email, password from users where email=?", request.Email)

		if row.Err() != nil {
			return ctx.String(http.StatusInternalServerError, row.Err().Error())
		}

		var retId int
		var retName, retEmail, retPassword string

		err := row.Scan(&retId, &retName, &retEmail, &retPassword)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return ctx.String(http.StatusUnauthorized, "email is not registered")
			}
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		err = bcrypt.CompareHashAndPassword([]byte(retPassword), []byte(request.Password))

		if err != nil {
			return ctx.String(http.StatusUnauthorized, err.Error())
		}

		tokenClaims := models.AuthClaimsJwt{
			UserId:    retId,
			UserName:  retName,
			UserEmail: retEmail,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
		tokenStr, err := token.SignedString([]byte("Test"))

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		response := LoginResponse{
			AccessToken: tokenStr,
		}

		return ctx.JSON(http.StatusOK, response)
	})
}
