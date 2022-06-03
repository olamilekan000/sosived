package interfaces

import (
	"github.com/olamilekan000/sosivio-app/auth-service/data"
	"github.com/olamilekan000/sosivio-app/auth-service/helpers"
	"github.com/olamilekan000/sosivio-app/auth-service/repository"
)

type UserService interface {
	// Login(name string) data.User
	FindOne(name string) (*data.User, int)
}

type AuthService interface {
	Login(name repository.LoginRequest) (*helpers.SuccessResponse, *helpers.AppError)
	VerifyAuthToken(token string) (*helpers.SuccessResponse, *helpers.AppError)
}

type Messenger interface {
	SendMessage(reqToken string) (*helpers.SuccessResponse, *helpers.AppError)
}

type JWTInterfaces interface {
	VerifyJWTToken(jwtToken string) (*repository.AuthResponse, error)
	CreateToken(username string) (string, error)
}
