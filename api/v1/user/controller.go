package user

import (
	"github.com/labstack/echo/v4"
	domain "github.com/w33h/Hexagonal-Architecture-Go/business/user"
	"github.com/w33h/Hexagonal-Architecture-Go/helpers"
	"net/http"
	"strconv"
)

type Controller struct {
	service domain.Service
}

func NewController(service domain.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) GetUser(c echo.Context) error {
	reqId, _ := strconv.Atoi(c.Param("id"))

	user, err := ctrl.service.FindById(reqId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &helpers.ResponseFormatter{
			Status: http.StatusBadRequest,
			Message: "Failed",
			Data: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &helpers.ResponseFormatter{
		Status: http.StatusOK,
		Message: "Success",
		Data: user,
	})
}
