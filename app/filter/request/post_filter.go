package request

import "bulwark/fanap-challenge/models"

// FilterPostRequest ...
type FilterPostRequest struct {
	Main  *models.Rectangle   `json:"main" validate:"required"`
	Input []*models.Rectangle `json:"input" validate:"required"`
}
