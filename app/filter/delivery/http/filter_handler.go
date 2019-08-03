package http

import (
	"bulwark/fanap-challenge/app/filter"

	requestTypes "bulwark/fanap-challenge/app/filter/request"
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

// FilterService filter service type
type FilterService struct {
	filterUsecase filter.Usecase
}

// NewFilterService creates new filter service
func NewFilterService(fu filter.Usecase) *FilterService {
	return &FilterService{
		filterUsecase: fu,
	}
}

// FindOverlapsAndStore finds overlaps in given inputs and store those
func (s *FilterService) FindOverlapsAndStore(ctx echo.Context) error {
	req := &requestTypes.FilterPostRequest{}
	if err := ctx.Bind(req); err != nil {
		return err
	}

	if err := ctx.Validate(req); err != nil {
		_err := err.(validator.ValidationErrors)

		return _err
	}

	err := s.filterUsecase.FindOverlapsAndStore(req.Main, req.Input)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, "Ok")
}

// FetchOverlaps fetch overlaps
func (s *FilterService) FetchOverlaps(ctx echo.Context) error {
	overlaps, err := s.filterUsecase.FetchOverlaps()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, overlaps)
}

// Register register filter service
func Register(e *echo.Echo, uc filter.Usecase) {
	filterService := NewFilterService(uc)

	e.POST("/", filterService.FindOverlapsAndStore)
	e.GET("/", filterService.FetchOverlaps)
}
