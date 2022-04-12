package repository

import (
	domain "github.com/w33h/Hexagonal-Architecture-Go/business/user"
	"github.com/w33h/Hexagonal-Architecture-Go/repository/user"
	"github.com/w33h/Hexagonal-Architecture-Go/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) domain.Repository {
	var userRepo domain.Repository

	if dbCon.Driver == util.Postgres {
		userRepo = user.NewPostgresRepository(dbCon.Postgres)
	}

	return userRepo
}
