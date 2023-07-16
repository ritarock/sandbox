package main

import (
	"crypto/subtle"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var health string
	var user string
	var pass string
	var endpoint string
	var port int

	health = "/health"
	user = "hoge"
	pass = "foo"
	endpoint = "0.0.0.0"
	port = 8080

	run(health, user, pass, endpoint, port)
}

func run(health, user, pass, endpoint string, port int) {
	e := echo.New()

	if health != "" {
		e.GET(health, func(c echo.Context) error {
			c.Response().Header().Set("Context-Type", "text/plain")
			return c.JSON(http.StatusOK, "ok")
		})
	}

	if user != "" && pass != "" {
		e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(user), []byte(username)) == 1 &&
				subtle.ConstantTimeCompare([]byte(pass), []byte(password)) == 1 {
				return true, nil
			}
			return false, nil
		}))
	}

	e.Logger.Fatal(e.Start(endpoint + ":" + strconv.Itoa(port)))
}
