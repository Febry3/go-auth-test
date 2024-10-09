package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterController(e *echo.Echo, db *sql.DB) {
	e.POST("v1/auth/register", func(ctx echo.Context) error {
		var request RegisterRequest

		json.NewDecoder(ctx.Request().Body).Decode(&request)

		hasdhedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		_, err = db.Exec(
			"INSERT INTO users (name, email, password) values (?, ?, ?)",
			request.Name,
			request.Email,
			hasdhedPassword,
		)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})
}
