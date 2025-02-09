package handlers

import (
	"blindsig/internal"
	"blindsig/views"

	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	app.GET("/", HomeHandler)
	app.GET("/.env", RickRoll)

	// Auth HTML API Routing
	app.POST("/api/v1/html/auth", htmlLogIn)

	// JSON API routing
	app.GET("/api/v1/json/publickey", apiJsonPubkey)
	app.POST("/api/v1/json/reqBlindSignature", apiJsonReqBlindSignature)

}

func HomeHandler(c echo.Context) error {
	return internal.Render(c, http.StatusOK, views.Home())
}

func RickRoll(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "https://www.youtube.com/watch?v=dQw4w9WgXcQ")
}
