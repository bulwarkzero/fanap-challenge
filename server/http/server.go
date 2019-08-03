package httpserver

import (
	"bulwark/fanap-challenge/server/http/gateway"
	"bulwark/fanap-challenge/server/http/middleware"

	"fmt"

	"gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo/v4"
)

// Config http server config schema
type Config struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
	Debug   bool   `json:"debug"`
}

// Serve serves the server
func Serve(cnf Config) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Debug = cnf.Debug

	e.Validator = &Validator{validator: validator.New()}

	middleware.Init(e)
	gateway.Init(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", cnf.Address, cnf.Port)))

	return e
}
