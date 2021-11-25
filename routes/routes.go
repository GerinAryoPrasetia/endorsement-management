package routes

import (
	"github.com/labstack/echo/v4"
	"project/factory"
)

func New() *echo.Echo {

	presenter := factory.Init()
	e := echo.New()

	e.GET("/creators", presenter.CreatorPresentation.GetAllCreator)

	return e
}