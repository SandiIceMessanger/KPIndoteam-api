package usecase

import (
	"context"
	"log"

	"chatnews-api/api/model"
	repository "chatnews-api/api/repository"
	usecase "chatnews-api/api/usecase"
	"chatnews-api/lib/logging"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type CategoryUseCase struct {
	config       *model.EnvConfig
	categoryRepo repository.CategoryRepositoryI
}

func NewCategoryUseCase(config *model.EnvConfig, categoryRepo repository.CategoryRepositoryI) usecase.CategoryUseCaseI {
	return &CategoryUseCase{
		config:       config,
		categoryRepo: categoryRepo,
	}
}

func (cuc *CategoryUseCase) GetDataUC(ctx context.Context) (resp model.GetDataCategoryResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	list, err := cuc.categoryRepo.GetAllDataCategory(ctx)
	if err != nil {
		log.Println("failed to show data category with default log")
		return list, err
	}

	return list, err
}

func (cuc *CategoryUseCase) InsertDataUC(ctx context.Context, req model.DataCategoryRequest) (resp bool, err error) {
	//check if context is nil
	if ctx == nil {
		ctx = context.Background()
	}

	if req.Name == "" {
		err = errors.New("failed to add data category ")
		logging.Info(err)
		return false, err
	}

	//insert data
	err = cuc.categoryRepo.InsertDataCategory(ctx, req)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (cuc *CategoryUseCase) UpdateDataUC(ctx context.Context, req model.DataCategoryRequest) (resp bool, err error) {
	//check if context is nil
	if ctx == nil {
		ctx = context.Background()
	}

	//update data
	err = cuc.categoryRepo.UpdateDataCategory(ctx, req)
	if err != nil {
		return false, err
	}
	zaplogger, _ := zap.NewProduction()
	defer zaplogger.Sync()
	zaplogger.Info("success to update product",
		zap.String("category ID", req.ID.Hex()),
	)

	return true, nil
}

func (cuc *CategoryUseCase) DeleteDataUC(ctx context.Context, req model.DataCategoryRequest) (resp bool, err error) {
	//check if context is nil
	if ctx == nil {
		ctx = context.Background()
	}

	//update data
	err = cuc.categoryRepo.DeleteDataCategory(ctx, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
