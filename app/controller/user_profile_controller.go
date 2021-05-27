package controller

import (
	"github.com/labstack/echo"
	"github.com/sidie88/lemonilo/app/request"
	"github.com/sidie88/lemonilo/app/service"
	"github.com/sidie88/lemonilo/storage"
	"net/http"
)

type UserProfileController struct {
	UserProfileService service.UserProfileService
}

func NewUserProfileController(cache *storage.UserProfileConcurrentMap) *UserProfileController {
	return &UserProfileController{
		UserProfileService: service.NewUserProfileServiceImpl(cache),
	}
}

func (c UserProfileController) CreateProfile(ctx echo.Context) (err error) {
	profileRequest := new(request.UserProfileRequest)
	if err = ctx.Bind(profileRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	profile, err := c.UserProfileService.CreateProfile(profileRequest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, profile)
}

func (c UserProfileController) FindAllProfiles(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, c.UserProfileService.GetProfiles())
}

func (c UserProfileController) FindProfileByUserId(ctx echo.Context) error {
	userId := ctx.Param("id")
	profile, err := c.UserProfileService.GetProfile(userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	if profile == nil {
		return ctx.JSON(http.StatusNoContent, profile)
	}
	return ctx.JSON(http.StatusOK, profile)
}

func (c UserProfileController) UpdateProfile(ctx echo.Context) (err error) {
	profileRequest := new(request.UserProfileRequest)
	userId := ctx.Param("id")
	if err = ctx.Bind(profileRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	profile, err := c.UserProfileService.UpdateProfile(profileRequest, userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, profile)
}

func (c UserProfileController) DeleteProfile(ctx echo.Context) (err error) {
	userId := ctx.Param("id")

	profile, err := c.UserProfileService.DeleteProfile(userId)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, profile)
}