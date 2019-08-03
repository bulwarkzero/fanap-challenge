package gateway

import (
	filterService "bulwark/fanap-challenge/app/filter/registry/http"

	"github.com/labstack/echo/v4"
)

// Init register services
func Init(e *echo.Echo) {
	filterService.Register(e)
}
