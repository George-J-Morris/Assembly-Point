package handlers

import (
	"blindsig/internal"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserDB struct {
	UUID          string
	Email         string
	Password      string
	Organisations []string
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

// HTML Handlers
func htmlLogIn(c echo.Context) error {
	logInQuery := "SELECT * FROM users WHERE email = $1"

	thisUser := new(UserDB)

	db, _ := internal.DB()
	db.QueryRow(c.Request().Context(), logInQuery, strings.ToLower(c.FormValue("username"))).Scan(&thisUser.UUID, &thisUser.Email, &thisUser.Password, &thisUser.Organisations)

	passCompare := bcrypt.CompareHashAndPassword([]byte(thisUser.Password), []byte(c.FormValue("password")))

	if passCompare != nil {
		fmt.Println(passCompare)
	} else {
		fmt.Println("Correct Password")
	}

	return nil
}
