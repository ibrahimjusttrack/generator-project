package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"myapp/internal/db"
	"myapp/internal/models"
	"myapp/internal/types"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func GetAllTemplates(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	templateResponse := types.TemplateAllResponse{}
	templateCollection := db.GetCollection(db.DBManager(), "template")
	results, err := templateCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	for results.Next(ctx) {
		var template models.Template
		if err = results.Decode(&template); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		templateResponse.Templates = append(templateResponse.Templates, template)
	}

	return c.JSON(http.StatusOK, templateResponse)
}
