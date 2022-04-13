package spec

type UpsertUserSpec struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}
