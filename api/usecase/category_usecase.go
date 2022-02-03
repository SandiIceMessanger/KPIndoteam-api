package crud

import (
	"context"

	"chatnews-api/api/model"
)

type CategoryUseCaseI interface {
	GetDataUC(ctx context.Context) (resp model.GetDataCategoryResponse, err error)
	InsertDataUC(ctx context.Context, req model.DataCategoryRequest) (resp bool, err error)
	UpdateDataUC(ctx context.Context, req model.DataCategoryRequest) (resp bool, err error)
	DeleteDataUC(ctx context.Context, req model.DataCategoryRequest) (resp bool, err error)
}
