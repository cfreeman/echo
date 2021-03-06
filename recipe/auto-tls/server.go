package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	// e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("<your_domain>")
	// Store the certificate to avoid issues with rate limits (https://letsencrypt.org/docs/rate-limits/)
	// e.AutoTLSManager.Cache = autocert.DirCache("<path to store key and certificate>")
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
			<h1>Welcome to Echo!</h1>
			<h3>TLS certificates automatically installed from Let's Encrypt :)</h3>
		`)
	})
	e.Logger.Fatal(e.StartAutoTLS(":443"))
}
