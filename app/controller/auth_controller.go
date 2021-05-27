package controller

import (
	"github.com/labstack/echo"
	"github.com/sidie88/lemonilo/app/service"
	"github.com/sidie88/lemonilo/storage"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(cache *storage.UserProfileConcurrentMap) *AuthController  {
	return &AuthController{
		AuthService: service.NewAuthServiceImpl(cache),
	}
}

func (auth *AuthController) Login(username, password string, ctx echo.Context) (bool, error){
	return auth.AuthService.Login(username, password)
}