package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sidie88/lemonilo/app/controller"
	"github.com/sidie88/lemonilo/storage"
)

type Server struct {
	*controller.UserProfileController
	*controller.AuthController
}

func main() {
	e := echo.New()
	server := Server{
		controller.NewUserProfileController(storage.GetSingletonUserProfileCache()),
		controller.NewAuthController(storage.GetSingletonUserProfileCache()),
	}

	g := e.Group("lemonilo", middleware.BasicAuth(server.Login))

	g.POST("/user-profile", server.UserProfileController.CreateProfile)
	g.GET("/user-profile", server.UserProfileController.FindAllProfiles)
	g.GET("/user-profile/:id", server.UserProfileController.FindProfileByUserId)
	g.PUT("/user-profile/:id", server.UserProfileController.UpdateProfile)
	g.DELETE("/user-profile/:id", server.UserProfileController.DeleteProfile)

	e.Logger.Fatal(e.Start(":8080"))
}
