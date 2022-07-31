package dto

import "awesomeProject0511/model"

type UserDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Uid   string `json:"uid"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:  user.Name,
		Email: user.Email,
		Uid:   user.Uid,
	}
}
