package models

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Template struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Language    string             `json:"language" bson:"language"`
}

type Metadata struct {
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	TemplateId primitive.ObjectID `json:"templateId" bson:"templateId"`
	Type       string             `json:"type" bson:"type"`
	Key        string             `json:"key" bson:"key"`
	Default    string             `json:"default" bson:"default"`
	Options    json.RawMessage    `json:"options" bson:"options"`
}