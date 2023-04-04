package handlers

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"myapp/internal/db"
	"myapp/internal/generator"
	"myapp/internal/models"
	"myapp/internal/types"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ctx := c.Request().Context()

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

func UploadTemplate(c echo.Context) error {
	ctx := c.Request().Context()
	templateId := c.Param("id")

	file, err := c.FormFile("project")

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	defer src.Close()

	path := filepath.Join("templates", file.Filename)
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	rootPath := unzip(path, "templates")
	templateCollection := db.GetCollection(db.DBManager(), "template")

	id := primitive.NewObjectID()
	err = id.UnmarshalText([]byte(templateId))
	if err != nil {
		fmt.Println(err)
	}
	result, err := templateCollection.UpdateByID(ctx, id, bson.D{{
		"$set", bson.D{
			{
				"path", rootPath,
			},
		},
	}})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	err = os.Remove(path)
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
		ID:          primitive.NewObjectID(),
		TemplateID:  objId,
		Type:        field.Type,
		Key:         field.Key,
		Default:     field.Default,
		Options:     field.Options,
		Description: field.Description,
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

func unzip(filePath string, dst string) string {
	archive, err := zip.OpenReader(filePath)
	if err != nil {
		panic(err)
	}
	defer archive.Close()
	var firstPath *string = nil
	for _, f := range archive.File {

		filePath := filepath.Join(dst, f.Name)
		if firstPath == nil {
			firstPath = &filePath
		}
		fmt.Println("unzipping file ", filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			fmt.Println("invalid file path")
			return ""
		}
		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		dstFile.Close()
		fileInArchive.Close()
	}
	return *firstPath
}
