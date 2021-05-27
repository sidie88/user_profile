package service

import (
	"errors"
	"github.com/sidie88/lemonilo/app/entity"
	"github.com/sidie88/lemonilo/app/request"
	"github.com/sidie88/lemonilo/storage"
	"golang.org/x/crypto/bcrypt"
)

type UserProfileService interface {
	CreateProfile(request *request.UserProfileRequest) (userProfile *entity.UserProfile, err error)
	GetProfiles() []*entity.UserProfile
	GetProfile(userId string) (userProfile *entity.UserProfile, err error)
	UpdateProfile(request *request.UserProfileRequest, userId string) (userProfile *entity.UserProfile, err error)
	DeleteProfile(userId string) (bool, error)
}

type UserProfileServiceImpl struct {
	cache *storage.UserProfileConcurrentMap
}

func NewUserProfileServiceImpl(cache *storage.UserProfileConcurrentMap) UserProfileService {
	return &UserProfileServiceImpl{
		cache,
	}
}

func (ups *UserProfileServiceImpl) CreateProfile(request *request.UserProfileRequest) (*entity.UserProfile, error) {

	profileBuilder := entity.NewBuilder()

	_, ok := ups.cache.Load(request.UserID)
	if ok {
		return nil, errors.New("user_id already used")
	}

	password, fail := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if fail != nil {
		return nil, errors.New("failed to generate password hash")
	}
	userProfile, err := profileBuilder.
		WithUserId(request.UserID).
		WithEmail(request.Email).
		WithAddress(request.Address).
		WithPassword(string(password)).
		Build()

	if err == nil {
		ups.cache.Store(request.UserID, &userProfile)
	}
	return &userProfile, err

}

func (ups *UserProfileServiceImpl) GetProfiles() []*entity.UserProfile {
	return ups.cache.GetAllValues()
}

func (ups *UserProfileServiceImpl) GetProfile(userId string) (*entity.UserProfile, error){
	loaded, ok := ups.cache.Load(userId)

	if !ok {
		return nil, errors.New("user_id was not found")
	}
	return loaded, nil
}

func (ups *UserProfileServiceImpl) UpdateProfile(request *request.UserProfileRequest, userId string) (*entity.UserProfile, error) {
	loaded, ok := ups.cache.Load(userId)

	if !ok {
		return nil, errors.New("user_id was not found")
	}

	return updateCache(userId, request, loaded, ups)
}

func updateCache(userId string, request *request.UserProfileRequest, loaded *entity.UserProfile, ups *UserProfileServiceImpl) (*entity.UserProfile, error){
	password, fail := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if fail != nil {
		return nil, errors.New("failed to generate password hash")
	}
	loaded.UserID = userId
	loaded.Email = request.Email
	loaded.Address = request.Address
	loaded.Password = string(password)
	ups.cache.Store(userId, loaded)
	return loaded, nil
}

func (ups *UserProfileServiceImpl) DeleteProfile(userId string) (bool, error){
	_, ok := ups.cache.Load(userId)

	if !ok {
		return false, errors.New("user_id was not found")
	}

	ups.cache.Delete(userId)

	return true, nil
}