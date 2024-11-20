package handlers

import (
	"blindsig/components"
	"blindsig/internal"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	app.GET("/", HomeHandler)
	app.GET("/.env", RickRoll)

	//API grouo
	app.Group("/api")

	// HTML API routing
	app.POST("/api/html/message", apiHtmlMessage)

	// JSON API routing
	app.GET("/api/json/publickey", apiJsonPubkey)
	app.POST("/api/json/reqBlindSignature", apiJsonReqBlindSignature)

}

func HomeHandler(c echo.Context) error {
	return internal.Render(c, http.StatusOK, components.Home())
}

func RickRoll(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "https://www.youtube.com/watch?v=dQw4w9WgXcQ")
}
