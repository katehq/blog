package serializer

import (
	"blog/models"
)

// User serializer 
type User struct {
	ID   uint   `json:"id"`
	UserName string `json:"username"`
	Avatar   string `json:"avatar"`
}

// UserResponse Response 
type UserResponse struct {
	Response 
	Data User `json:"data"`
}


// BuildUser build user
func BuildUser(user models.User) User {
	return User{
		ID : user.ID,
		UserName: user.Name,
		Avatar: user.Avatar,
	}
}

// BuildUserResponse return response for user
func BuildUserResponse(user models.User) UserResponse {
	return UserResponse {
		Data: BuildUser(user),
	}
}