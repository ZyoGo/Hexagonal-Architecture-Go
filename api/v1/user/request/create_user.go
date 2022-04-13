package request

import "github.com/w33h/Hexagonal-Architecture-Go/business/user/spec"

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) ToUpsertUserSpec() *spec.UpsertUserSpec {
	var upsertUserSpec spec.UpsertUserSpec
	upsertUserSpec.Email = r.Email
	upsertUserSpec.Password = r.Password

	return &upsertUserSpec
}
