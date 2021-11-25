package presentation

import (
	"net/http"
	"project/features/creator"
	presentation_response "project/features/creator/presentation/response"

	"github.com/labstack/echo/v4"
)

type CreatorHandler struct {
	creatorBussiness creator.Bussiness
}

func NewCreatorHandler(cb creator.Bussiness) *CreatorHandler {
	return &CreatorHandler{
		creatorBussiness: cb,
	}
}

func (ch *CreatorHandler) GetAllCreator(c echo.Context) error {
	result := ch.creatorBussiness.GetAllData("")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data": presentation_response.FromCoreSlice(result),
	})
}

func (ch *CreatorHandler) GetCreatorByName(c echo.Context) error {
	result := ch.creatorBussiness.GetCreatorByName("")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data": presentation_response.FromCoreSlice(result),
	})
}