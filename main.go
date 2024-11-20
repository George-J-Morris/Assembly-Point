package main

import (
	"blindsig/handlers"
	"blindsig/internal"

	_ "database/sql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/acme/autocert"
)

func main() {

	scssPath := "./assets/scss/theme.scss"
	cssOutPAth := "./assets/css"
	internal.TranspileBootstrapCss(scssPath, cssOutPAth)

	e := echo.New()
	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("assemblypoint.org", "www.assemblypoint.org")
	// Cache certificates to avoid issues with rate limits (https://letsencrypt.org/docs/rate-limits)
	e.AutoTLSManager.Cache = autocert.DirCache("./certs")

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Static("/assets", "assets")

	handlers.SetupRoutes(e)

	e.Logger.Fatal(e.StartAutoTLS(":443"))
}
