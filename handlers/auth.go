package handlers

import (
	"blindsig/internal"
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"net/netip"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/mileusna/useragent"
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
	Ip_address net.IP
	Browser    string
	Os         string
	Expires    time.Time
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

func getSessions(uuid string) []Session {
	db, _ := internal.DB()
	rows, _ := db.Query(context.Background(), "SELECT * FROM sessions WHERE uuid = $1", uuid)

	sessions, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (Session, error) {
		thisSession := new(Session)

		err := row.Scan(&thisSession.Session_id, &thisSession.UUID, &thisSession.Ip_address, &thisSession.Browser, &thisSession.Os, &thisSession.Expires)
		return *thisSession, err
	})

	if err != nil {
		fmt.Printf("CollectRows error: %v", err)
		return nil
	}
	return sessions
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

		sessions := getSessions(thisUser.UUID)

		for i := range sessions {
			fmt.Println(sessions[i])
		}

		userIP := c.Request().Header["Cf-Connecting-Ip"]
		parsedIp, _ := netip.ParseAddr(userIP[0])
		userAgent := useragent.Parse(c.Request().UserAgent())
		userOs := userAgent.OS
		userBrowser := userAgent.Name
		fmt.Println(parsedIp.String() + "\n" + userBrowser + "\n" + userOs)

		//_, err := db.Exec(c.Request().Context(), "insert into sessions(uuid,ip_address,browser,os) values($1,$2,$3,$4)", thisUser.UUID, parsedIp, userBrowser, userOs)
		//if err != nil {
		//	fmt.Println(err)
		//}
	}

	return nil
}
