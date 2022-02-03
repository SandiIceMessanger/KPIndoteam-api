package http

import (
	"chatnews-api/api/delivery/controller"
	api "chatnews-api/api/usecase"
	"chatnews-api/middleware/config"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

func CategoryRouter(e *echo.Echo, categoryUseCase api.CategoryUseCaseI) {

	categoryCtrl := controller.NewCategoryController(e, categoryUseCase)

	r := e.Group("/api/v1")
	r.Use(middleware.JWT([]byte(config.JWTSecret)))

	r.GET("/category/list", categoryCtrl.GetData)
	r.POST("/category/add", categoryCtrl.InsertData)
	r.POST("/category/edit", categoryCtrl.UpdateData)
	r.POST("/category/delete", categoryCtrl.DeleteData)

}
