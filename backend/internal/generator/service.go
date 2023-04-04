package generator

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"myapp/internal/db"
	"myapp/internal/models"
	"os"
	"os/exec"
	"sort"
)

type Input struct {
	Key   string `json:"accessor"`
	Value string `json:"value"`
}

type Service interface {
	//return path for file and error of generator
	Generate(ctx context.Context, templateID string, input []Input) (string, error)
}

type Generator struct {
	db mongo.Client
}

func (s Generator) Generate(ctx context.Context, templateID string, input []Input) (string, error) {
	template, err := s.getTemplate(ctx, templateID)
	if err != nil {
		return "", err
	}

	filedByKey := make(map[string]models.Metadata)
	for _, v := range template.fields {
		filedByKey[v.Key] = v
	}

	//generate json
	configString := "{"
	for i, v := range input {
		var row string
		if i != len(input)-1 {
			row = fmt.Sprintf(`"%s":"%s",`, v.Key, v.Value)
		} else {
			row = fmt.Sprintf(`"%s":"%s"`, v.Key, v.Value)
		}
		configString = fmt.Sprintf("%s%s", configString, row)
	}
	configString = configString + "}"

	cutter, err := os.Create(template.template.Path + "/cookiecutter.json")
	if err != nil {
		return "", err
	}
	defer cutter.Close()

	_, err = cutter.WriteString(configString)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("cookiecutter", "-f", "--no-input", "tmp/", template.template.Path)
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return "", nil
}

type templateFull struct {
	template models.Template
	fields   []models.Metadata
}

func (s Generator) getTemplate(ctx context.Context, templateID string) (templateFull, error) {
	var template models.Template
	templateCollection := db.GetCollection(db.DBManager(), "template")
	err := templateCollection.FindOne(ctx, bson.M{"_id": templateID}).Decode(&template)

	if err != nil {
		return templateFull{}, err
	}

	objId, _ := primitive.ObjectIDFromHex(templateID)
	fieldCollection := db.GetCollection(db.DBManager(), "field")
	results, err := fieldCollection.Find(ctx, bson.M{"templateId": objId})
	if err != nil {
		return templateFull{}, err
	}
	var fieldsResponse []models.Metadata
	for results.Next(ctx) {
		var field models.Metadata
		if err = results.Decode(&field); err != nil {
			return templateFull{}, err
		}

		fieldsResponse = append(fieldsResponse, field)
	}

	sort.SliceStable(fieldsResponse, func(i, j int) bool {
		return fieldsResponse[i].Order < fieldsResponse[j].Order
	})

	return templateFull{
		template: template,
		fields:   fieldsResponse,
	}, nil
}
