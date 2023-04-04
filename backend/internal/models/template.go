package models

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Template struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Language    string             `json:"language" bson:"language"`
	Path        string             `json:"path"`
}

type Metadata struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	TemplateID primitive.ObjectID `json:"templateId" bson:"templateId"`
	Type       string             `json:"type" bson:"type"`
	Key        string             `json:"accessor" bson:"key"`
	Default    string             `json:"default" bson:"default"`
	Options    json.RawMessage    `json:"options" bson:"options"`
	Order      int                `json:"-" bson:"order"`
}
