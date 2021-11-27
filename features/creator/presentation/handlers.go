package presentation

import (
	"net/http"
	"project/features/creator"
	presentation_response "project/features/creator/presentation/response"
	"strconv"

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

// func (ch *CreatorHandler) GetCreatorByName(c echo.Context) error {
// 	result := ch.creatorBussiness.GetCreatorByName("")

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success",
// 		"data": presentation_response.FromCoreSlice(result),
// 	})
// }

func (ch *CreatorHandler) GetCreatorByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}}

	data, err := ch.creatorBussiness.GetCreatorByID(creator.Core{ID: id})
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data": presentation_response.ToCreatorResponse(data),
	})
}