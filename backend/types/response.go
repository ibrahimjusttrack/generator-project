package types

import "myapp/models"

type TemplateAllResponse struct {
	Templates []models.Template `json:"results"`
	Total     int               `json:"total"`
}
