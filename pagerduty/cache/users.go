package cache

import (
	"github.com/petetanton/go-pagerduty-sdk/pagerduty/model"
	"sync"
)

type UserCache struct {
	m     sync.RWMutex
	users map[string]*model.User
}

func NewUserCache() *UserCache {
	return &UserCache{
		m:     sync.RWMutex{},
		users: make(map[string]*model.User),
	}
}

func (uc *UserCache) ReadUser(id string) *model.User {
	uc.m.RLock()
	defer uc.m.RUnlock()
	return uc.users[id]
}

func (uc *UserCache) WriteUser(user *model.User) {
	uc.m.Lock()
	defer uc.m.Unlock()
	uc.users[user.Id] = user
}

func (uc *UserCache) WriteUsers(users []*model.User) {
	uc.m.Lock()
	defer uc.m.Unlock()
	for _, user := range users {
		uc.users[user.Id] = user
	}
}

func (uc *UserCache) RemoveUser(id string) {
	uc.m.Lock()
	defer uc.m.Unlock()
	uc.users[id] = nil
}
