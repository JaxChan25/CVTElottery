package serializer

import "singo/model"

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"created_at"`
}

// BuildUser 序列化用户
func BuildUser(user model.GameUser) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04"),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.GameUser) Response {
	return Response{
		Data: BuildUser(user),
	}
}
