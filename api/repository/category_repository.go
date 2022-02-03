package crud

import (
	"context"

	"chatnews-api/api/model"
)

type CategoryRepositoryI interface {
	GetAllDataCategory(ctx context.Context) (crudResp model.GetDataCategoryResponse, err error)
	InsertDataCategory(ctx context.Context, req model.DataCategoryRequest) error
	UpdateDataCategory(ctx context.Context, req model.DataCategoryRequest) error
	DeleteDataCategory(ctx context.Context, req model.DataCategoryRequest) error
}
