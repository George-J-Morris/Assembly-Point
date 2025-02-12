package main

import (
	"blindsig/handlers"
	"blindsig/internal"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	internal.InitDBPool()

	e := echo.New()
	//e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("assemblypoint.org", "www.assemblypoint.org")
	// Cache certificates to avoid issues with rate limits (https://letsencrypt.org/docs/rate-limits)
	//e.AutoTLSManager.Cache = autocert.DirCache("./certs")

	e.IPExtractor = echo.ExtractIPFromXFFHeader()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Static("/assets", "assets")

	handlers.SetupRoutes(e)

	// Currently, website is proxied through cloudflare, which does the tls between
	// cloudflare and the client. This certificate is the origin certificate issued
	// by cloudflare which covers tls between the origin and cloudflare.
	e.Logger.Fatal(e.StartTLS(":443", "./certs/cert.pem", "./certs/priv_key.pem"))
}
