package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Tgis represents a person resource ein the application.
type Person struct {
	Id     primitive.ObjectID `json:"_id, omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name"`
	Age    int                `json:"age" bson:"age"`
	Gender string             `json:"gender" bson:"gender"`
}
