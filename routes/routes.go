package routes

import (
	"project/constant"
	"project/factory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	presenter := factory.Init()
	e := echo.New()
	jwt := e.Group("")
	jwt.Use(middleware.JWT([]byte(constant.JWT_KEY)))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))


	e.GET("/creators", presenter.CreatorPresentation.GetAllCreator)
	e.GET("/creators/:id", presenter.CreatorPresentation.GetCreatorByID)

	e.POST("/creator", presenter.CreatorPresentation.RegisterCreator)
	e.DELETE("/creator/:id", presenter.CreatorPresentation.DeleteCreator)
	e.PUT("creator/:id", presenter.CreatorPresentation.UpdateCreator)
	return e
}