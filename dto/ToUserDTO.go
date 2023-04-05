package dto

import "awesomeProject0511/model"

type UserDTO struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	AvatarUrl string `json:"avatar_url"`
	Gender    string `json:"gender"`
	Email     string `json:"email"`
	Age       uint   `json:"age"`
}

func ToUserDTO(user model.User) UserDTO {
	return UserDTO{
		ID:        user.ID,
		Username:  user.Username,
		Nickname:  user.Nickname,
		AvatarUrl: user.AvatarUrl,
		Gender:    user.Gender,
		Email:     user.Email,
		Age:       user.Age,
	}
}
