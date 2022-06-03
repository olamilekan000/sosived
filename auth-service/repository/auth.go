package repository

import (
	"github.com/olamilekan000/sosivio-app/auth-service/data"
)

type UserRepository struct {
	client *data.InmemoryDb
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type VerifyAuthTokenReq struct {
	Token string `json:"token"`
}

type APIResponse struct {
	Code    uint        `json:"statusCode,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

type AuthResponse struct {
	Message string      `json:"message,omitempty"`
	User    interface{} `json:"user,omitempty"`
	Token   string      `json:"token,omitempty"`
}

type User struct {
	User interface{} `json:"user,omitempty"`
}

func (u UserRepository) FindOne(name string) (*data.User, int) {
	user, _ := u.client.FindOne(name)

	return user, 1
}

func NewUserRepository() UserRepository {
	return UserRepository{
		client: data.AllUsers,
	}
}
