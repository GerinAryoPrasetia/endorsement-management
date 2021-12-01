package presentation

import (
	"net/http"
	"project/features/creator"
	"project/features/creator/presentation/request"
	"project/features/creator/presentation/response"
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

func (ch *CreatorHandler) RegisterCreator(c echo.Context) error {
	reqData := request.CreatorRequest{}

	err:= c.Bind(&reqData)
	if err!= nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	_, err = ch.creatorBussiness.RegisterCreator(request.ToCreatorCore(reqData))	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func (ch *CreatorHandler) UpdateCreator(c echo.Context) error {
	var creatorData request.CreatorRequest
	err := c.Bind(&creatorData)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func (ch *CreatorHandler) DeleteCreator(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return response.NewErrorResponse(c, err.Error(), http.StatusBadRequest)
	}

	err = ch.creatorBussiness.DeleteCreator(creator.Core{ID: id})

	if err != nil {
		return response.NewErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "berhasil menghapus creator",
	})
}