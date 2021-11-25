package routes

import (
	"project/factory"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	presenter := factory.Init()
	e := echo.New()

	e.GET("/creators", presenter.CreatorPresentation.GetAllCreator)

	return e
}