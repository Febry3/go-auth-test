package main

import (
	"github.com/febry3/go-auth-test/controller/auth"
	"github.com/febry3/go-auth-test/database"
	"github.com/labstack/echo/v4"
)

func main() {
	db := database.InitDb()
	defer db.Close()

	err := db.Ping()

	if err != nil {
		panic(err)
	}

	e := echo.New()

	auth.RegisterController(e, db)
	auth.LoginController(e, db)

	e.Start(":9090")

}
