package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myapp/internal/db"
	"myapp/internal/generator"
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

	templateResponse.Total = len(templateResponse.Templates)

	return c.JSON(http.StatusOK, templateResponse)
}

func CreateTemplate(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var template models.Template

	if err := c.Bind(&template); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	newTemplate := models.Template{
		ID:          primitive.NewObjectID(),
		Name:        template.Name,
		Description: template.Description,
		Language:    template.Language,
	}
	templateCollection := db.GetCollection(db.DBManager(), "template")

	result, err := templateCollection.InsertOne(ctx, newTemplate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, result)
}

func GetFields(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var fieldsResponse []models.Metadata
	templateId := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(templateId)
	fieldCollection := db.GetCollection(db.DBManager(), "field")
	results, err := fieldCollection.Find(ctx, bson.M{"templateId": objId})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	for results.Next(ctx) {
		var field models.Metadata
		if err = results.Decode(&field); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		fieldsResponse = append(fieldsResponse, field)
	}

	return c.JSON(http.StatusOK, fieldsResponse)
}

func CreateField(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var field models.Metadata

	if err := c.Bind(&field); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	templateId := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(templateId)

	newField := models.Metadata{
		ID:         primitive.NewObjectID(),
		TemplateID: objId,
		Type:       field.Type,
		Key:        field.Key,
		Default:    field.Default,
		Options:    field.Options,
	}
	fieldCollection := db.GetCollection(db.DBManager(), "field")

	result, err := fieldCollection.InsertOne(ctx, newField)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, result)
}

func CreateJSONConfig(c echo.Context) error {
	var configs []generator.Input
	templateId := c.Param("id")

	if err := c.Bind(&configs); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	g := generator.Generator{}

	_, err := g.Generate(c.Request().Context(), templateId, configs)
	if err != nil {
		return c.String(500, err.Error())
	}

	return c.JSON(http.StatusCreated, string(""))
}
