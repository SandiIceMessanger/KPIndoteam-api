package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	GetDataCategoryResponse struct {
		Data []DataCategory `json:"categories"`
	}
	DataCategory struct {
		ID        primitive.ObjectID `json:"id" xml:"id" bson:"_id,omitempty"`
		Name      string             `json:"name" bson:"name"`
		CreatedBy string             `json:"created_by" bson:"created_by"`
		CreatedAt time.Time          `json:"created_at" bson:"created_at"`
		UpdatedBy string             `json:"update_by" bson:"update_by"`
		UpdatedAt *time.Time         `json:"update_at" bson:"update_at"`
		DeletedBy string             `json:"deleted_by" bson:"deleted_by"`
		DeletedAt *time.Time         `json:"deleted_at" bson:"deleted_at"`
	}

	DataCategoryRequest struct {
		ID        primitive.ObjectID `json:"id" xml:"id" bson:"_id,omitempty"`
		Name      string             `json:"name,omitempty" bson:"name"`
		CreatedBy string             `json:"created_by" bson:"created_by"`
		CreatedAt time.Time          `json:"created_at" bson:"created_at"`
		UpdatedBy string             `json:"update_by" bson:"update_by"`
		UpdatedAt *time.Time         `json:"update_at" bson:"update_at"`
		DeletedBy string             `json:"deleted_by" bson:"deleted_by"`
		DeletedAt *time.Time         `json:"deleted_at" bson:"deleted_at"`
	}
)
