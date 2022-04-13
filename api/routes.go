package api

import (
	"github.com/labstack/echo/v4"
	userV1 "github.com/w33h/Hexagonal-Architecture-Go/api/v1/user"
)

type Controller struct {
	UserV1Controller *userV1.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	userv1 := e.Group("/v1/user")
	userv1.GET("/:id", controller.UserV1Controller.GetUser)
}