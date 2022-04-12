package helper

type UserResponse struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//func ToUserResponse(user business.Users) UserResponse {
//	return UserResponse{
//		Id: user.Id,
//		Email: user.Email,
//		Password: user.Password,
//	}
//}

//func ToUserResponses(users []business.Users) []UserResponse {
//	var userResponses []UserResponse
//	for _, user := range users {
//		userResponses = append(userResponses, ToUserResponse(user))
//	}
//}