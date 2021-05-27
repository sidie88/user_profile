package storage

import (
	"github.com/sidie88/lemonilo/app/entity"
	"sync"
)

var once sync.Once

type UserProfileConcurrentMap struct {
	sync.RWMutex
	cache map[string]*entity.UserProfile
}

var userProfileCache UserProfileConcurrentMap

func GetSingletonUserProfileCache() *UserProfileConcurrentMap {
	once.Do(func() {
		cache := make(map[string]*entity.UserProfile)

		defaultUser := entity.UserProfile{
			Password: "$2a$04$p15xSyEYn4Re5kUsD./f4emjZbIf8o/4J3TFSZVOuRSoUZjEP5w6y",
		}
		cache["default_user"] = &defaultUser

		userProfileCache = UserProfileConcurrentMap{
			cache: cache,
		}
	})
	return &userProfileCache
}

func (cm *UserProfileConcurrentMap) Load(key string) (value *entity.UserProfile, ok bool) {
	cm.RLock()
	result, ok := cm.cache[key]
	cm.RUnlock()
	return result, ok
}

func (cm *UserProfileConcurrentMap) Delete(key string) {
	cm.Lock()
	delete(cm.cache, key)
	cm.Unlock()
}

func (cm *UserProfileConcurrentMap) Store(key string, value *entity.UserProfile) {
	cm.Lock()
	cm.cache[key] = value
	cm.Unlock()
}


func (cm *UserProfileConcurrentMap) GetAllValues() []*entity.UserProfile {
	profiles := make([]*entity.UserProfile, 0, len(cm.cache))
	for _, v := range cm.cache {
		if v != nil && v.UserID != ""{
			profiles = append(profiles, v)

		}
	}
	return profiles
}
