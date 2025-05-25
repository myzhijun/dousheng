package respToDTO

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/model/dto"
)

func GetUserDTO(user *model.User) *dto.UserDTO {
	return &dto.UserDTO{
		ID:            user.ID,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
}

func GetUserListDTO(users []model.User) *[]dto.UserDTO {
	userList := make([]dto.UserDTO, len(users))
	for i := 0; i < len(users); i++ {
		user := users[i]
		userList[i] = *GetUserDTO(&user)
	}
	return &userList
}
