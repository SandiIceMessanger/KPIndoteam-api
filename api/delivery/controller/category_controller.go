package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"chatnews-api/api/model"
	api "chatnews-api/api/usecase"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	e       *echo.Echo
	usecase api.CategoryUseCaseI
}

func NewCategoryController(e *echo.Echo, usecase api.CategoryUseCaseI) *CategoryController {
	return &CategoryController{
		e:       e,
		usecase: usecase,
	}
}

func (cc *CategoryController) GetData(ec echo.Context) error {

	data, err := cc.usecase.GetDataUC(context.Background())
	if err != nil {
		return ec.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":     "false",
			"message":    "err",
			"error_code": "502",
			"data":       data,
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "success",
		"data":    data,
	})
}

func (cc *CategoryController) InsertData(ec echo.Context) error {

	var req model.DataCategoryRequest
	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":     "false",
			"message":    "err",
			"error_code": "502",
			"data":       "",
		})
	}
	data, err := cc.usecase.InsertDataUC(context.Background(), req)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":     "false",
			"message":    "err",
			"error_code": "502",
			"data":       data,
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "success",
		"data":    data,
	})
}

func (cc *CategoryController) UpdateData(ec echo.Context) error {

	var req model.DataCategoryRequest
	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":     "false",
			"message":    "err",
			"error_code": "502",
			"data":       "",
		})
	}
	data, err := cc.usecase.UpdateDataUC(context.Background(), req)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":     "false",
			"message":    "err",
			"error_code": "502",
			"data":       data,
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "success",
		"data":    data,
	})
}

func (cc *CategoryController) DeleteData(ec echo.Context) error {

	var req model.DataCategoryRequest
	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":     "false",
			"message":    "err",
			"error_code": "502",
			"data":       "",
		})
	}
	data, err := cc.usecase.DeleteDataUC(context.Background(), req)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":     "false",
			"message":    "err",
			"error_code": "502",
			"data":       data,
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "success",
		"data":    data,
	})
}
