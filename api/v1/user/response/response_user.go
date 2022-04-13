package response

import domain "github.com/w33h/Hexagonal-Architecture-Go/business/user"

type UserResponse struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToUserResponse(user domain.Users) UserResponse {
	return UserResponse{
		Id:       user.Id,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ToUserResponses(users []domain.Users) []UserResponse {
	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}
