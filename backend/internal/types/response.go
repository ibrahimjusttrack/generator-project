package types

import (
	"myapp/internal/models"
)

type TemplateAllResponse struct {
	Templates []models.Template `json:"results"`
	Total     int               `json:"total"`
}

type FileCreatedResponse struct {
	Message string `json:"message"`
}
