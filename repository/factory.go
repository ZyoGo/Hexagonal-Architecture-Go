package repository

import (
	"github.com/w33h/Hexagonal-Architecture-Go/repository/user"
	"github.com/w33h/Hexagonal-Architecture-Go/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) repository.UserRepository {
	var userRepo repository.UserRepository

	if dbCon.Driver == util.Postgres {
		userRepo = repository.NewPostgresRepository(dbCon.Postgres)
	}

	return userRepo
}
