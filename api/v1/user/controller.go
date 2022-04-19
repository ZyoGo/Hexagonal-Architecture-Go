package user

import (
	"github.com/labstack/echo/v4"
	"github.com/w33h/Hexagonal-Architecture-Go/business/user"
	"github.com/w33h/Hexagonal-Architecture-Go/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/w33h/Hexagonal-Architecture-Go/api/v1/user/request"
	"github.com/w33h/Hexagonal-Architecture-Go/api/v1/user/response"
	"github.com/w33h/Hexagonal-Architecture-Go/helpers"
	service "github.com/w33h/Hexagonal-Architecture-Go/service/user"
)

type Controller struct {
	service user.UserService
}

func NewController(service user.UserService) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) GetUser(c echo.Context) error {
	reqId, _ := strconv.Atoi(c.Param("id"))

	user, err := ctrl.service.FindById(reqId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &helpers.ResponseFormatter{
			Status:  http.StatusBadRequest,
			Message: "Failed",
			Data:    err.Error(),
		})
	}

	userResponse := response.ToUserResponse(*user)

	return c.JSON(http.StatusOK, &helpers.ResponseFormatter{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    userResponse,
	})
}

func (ctrl *Controller) CreateUser(c echo.Context) error {
	createUserRequest := new(request.CreateUserRequest)
	if err := c.Bind(&createUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, &helpers.ResponseFormatter{
			Status:  http.StatusBadRequest,
			Message: "Failed",
			Data:    err.Error(),
		})
	}

	user, err := ctrl.service.Create(*createUserRequest.ToUpsertUserSpec())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &helpers.ResponseFormatter{
			Status:  http.StatusInternalServerError,
			Message: "Failed",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, &helpers.ResponseFormatter{
		Status:  http.StatusCreated,
		Message: "Failed",
		Data:    user,
	})
}

func (ctrl *Controller) GetUsers(c echo.Context) error {
	users, err := ctrl.service.FindAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, &helpers.ResponseFormatter{
			Status:  http.StatusBadRequest,
			Message: "Failed",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &helpers.ResponseFormatter{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    users,
	})
}
