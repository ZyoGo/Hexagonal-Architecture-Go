package modules

import (
	"github.com/w33h/Hexagonal-Architecture-Go/api"
	userV1Controller "github.com/w33h/Hexagonal-Architecture-Go/api/v1/user"
	userRepository "github.com/w33h/Hexagonal-Architecture-Go/repository"
	userService "github.com/w33h/Hexagonal-Architecture-Go/service/user"
	"github.com/w33h/Hexagonal-Architecture-Go/util"
)

func RegisterModules(dbCon *util.DatabaseConnection) api.Controller {
	userPermitRepository := userRepository.RepositoryFactory(dbCon)
	userPermitService := userService.NewService(userPermitRepository)

	userV1PermitController := userV1Controller.NewController(userPermitService)

	controller := api.Controller{
		UserV1Controller: userV1PermitController,
	}

	return controller
}