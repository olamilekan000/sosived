package services

import (
	"strings"

	helpers "github.com/olamilekan000/sosivio-app/auth-service/helpers"
	"github.com/olamilekan000/sosivio-app/auth-service/interfaces"
	"github.com/olamilekan000/sosivio-app/auth-service/repository"
)

type AuthService struct {
	repo interfaces.UserService
	jwt  interfaces.JWTInterfaces
}

func (u AuthService) Login(req repository.LoginRequest) (*helpers.SuccessResponse, *helpers.AppError) {
	user, res := u.repo.FindOne(req.Username)

	if res < 1 {
		return nil, helpers.NewNotFoundError(
			"User not found",
		)
	}

	if user.Password != req.Password {
		return nil, helpers.NewAuthorizationError(
			"Invalid password",
		)
	}

	token, err := u.jwt.CreateToken(user.Name)

	if err != nil {
		return nil, helpers.NewUnexpectedError("An unexpected error occured")
	}

	return helpers.Success("Authentication was successful", &repository.AuthResponse{Token: token}), nil
}

func (u AuthService) VerifyAuthToken(token string) (*helpers.SuccessResponse, *helpers.AppError) {

	splitedToken := strings.Split(token, " ")

	if splitedToken[0] != "Bearer" {
		return nil, helpers.NewAuthorizationError("Inavlid Token")
	}

	rep, err := u.jwt.VerifyJWTToken(splitedToken[1])

	if err != nil {
		return nil, helpers.NewUnexpectedError(err.Error())
	}

	return helpers.Success("Validation was successful", &rep), nil
}

func NewAuthService(repo interfaces.UserService, jwt interfaces.JWTInterfaces) AuthService {
	return AuthService{repo: repo, jwt: jwt}
}
