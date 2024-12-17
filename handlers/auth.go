package handlers

import (
	"blindsig/internal"
	"fmt"

	"github.com/labstack/echo/v4"
)

type UserDB struct {
	UUID     string
	Email    string
	Password string
}

type Session struct {
	Session_id string
	UUID       string
}

func CookieTest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		var hasSession bool

		sessionCookie, err := c.Cookie("sessionID")

		if err != nil {
			fmt.Println(err)
			hasSession = false
		} else {
			hasSession = true
			db, _ := internal.DB()

			sessionQuery := ""
			db.Query(c.Request().Context(), sessionQuery)

		}

		fmt.Println(hasSession)
		fmt.Println(sessionCookie)

		return next(c)
	}
}
