package service

import (
	"errors"
	"github.com/sidie88/lemonilo/storage"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(userId, password string) (bool, error)
}

type AuthServiceImpl struct {
	cache *storage.UserProfileConcurrentMap
}

func NewAuthServiceImpl(cache *storage.UserProfileConcurrentMap) AuthService{
	return &AuthServiceImpl{cache: cache}
}

func (as *AuthServiceImpl) Login(userId, password string) (bool, error) {
	load, ok := as.cache.Load(userId)
	if !ok {
		return false, errors.New("user id was not found")
	}
	err := bcrypt.CompareHashAndPassword([]byte(load.Password), []byte(password))

	if err == nil {
		load.IsLogin = true
		as.cache.Store(userId, load)
		return true, nil
	}

	return false, errors.New("password is not match")

}