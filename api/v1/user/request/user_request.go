package request

import "github.com/w33h/Hexagonal-Architecture-Go/business/user/spec"

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"'`
	Password string `json:"password" validate:"required"`
}

func (r *CreateUserRequest) ToUpsertUserSpec() *spec.UpsertUserSpec {
	var upsertUserSpec spec.UpsertUserSpec

	upsertUserSpec.Username = r.Username
	upsertUserSpec.Email = r.Email
	upsertUserSpec.Password = r.Password

	return &upsertUserSpec
}
