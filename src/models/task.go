package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Task -> model
type Task struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Text      string             `json:"text" binding:"required" bson:"text"`
	Completed bool               `json:"completed" bson:"completed"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
}
