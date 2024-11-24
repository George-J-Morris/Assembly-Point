package handlers

import (
	"blindsig/internal"
	views "blindsig/views"

	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	app.GET("/", HomeHandler)
	app.GET("/.env", RickRoll)

	//API group
	app.Group("/api")

	// HTML API routing
	app.GET("/", HomeHandler)
	app.GET("/superuser", SAdminHandler)

	// JSON API routing
	app.GET("/api/json/publickey", apiJsonPubkey)
	app.POST("/api/json/reqBlindSignature", apiJsonReqBlindSignature)

}

func HomeHandler(c echo.Context) error {
	return internal.Render(c, http.StatusOK, views.Home())
}

func SAdminHandler(c echo.Context) error {
	return internal.Render(c, http.StatusOK, views.SuperUser())
}

func RickRoll(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "https://www.youtube.com/watch?v=dQw4w9WgXcQ")
}
