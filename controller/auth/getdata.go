package auth

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/febry3/go-auth-test/models"
	"github.com/labstack/echo/v4"
)

func GetData(e *echo.Echo, db *sql.DB) {
	e.GET("v1/data", func(ctx echo.Context) error {
		user := ctx.Get("User").(models.AuthClaimsJwt)
		fmt.Print(user)
		return ctx.JSON(http.StatusOK, nil)
	})
}
