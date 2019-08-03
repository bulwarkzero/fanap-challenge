package http

import (
	"bulwark/fanap-challenge/app/filter/delivery/http"
	"bulwark/fanap-challenge/app/filter/repository"
	"bulwark/fanap-challenge/app/filter/usecase"
	"bulwark/fanap-challenge/database"

	"github.com/labstack/echo/v4"
)

// Register ...
func Register(e *echo.Echo) {
	db := database.GetInstance()
	overlapRepo := repository.NewOverlapRepo(db)
	uc := usecase.New(overlapRepo)
	http.Register(e, uc)
}
