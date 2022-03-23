package cache

import (
	"errors"
	"sync"

	"github.com/petetanton/go-pagerduty-sdk/pagerduty/model"
	"github.com/sirupsen/logrus"
)

type UserCache struct {
	m      sync.RWMutex
	users  map[string]*model.User
	logger logrus.FieldLogger
}

func NewUserCache(logger logrus.FieldLogger) *UserCache {
	return &UserCache{
		m:      sync.RWMutex{},
		users:  make(map[string]*model.User),
		logger: logger.WithField("package", "cache"),
	}
}

func (uc *UserCache) ReadUser(id string) *model.User {
	uc.logger.Debugf("reading user %s from cache", id)
	uc.m.RLock()
	defer uc.m.RUnlock()
	return uc.users[id]
}

func (uc *UserCache) WriteUser(user *model.User) error {
	uc.logger.Debugf("writing user %s to cache", user)
	if user == nil || user.Id == "" {
		return errors.New("a user cannot be cached without an ID")
	}
	uc.m.Lock()
	defer uc.m.Unlock()
	uc.users[user.Id] = user

	return nil
}

func (uc *UserCache) WriteUsers(users []*model.User) error {
	uc.logger.Debugf("writing users [%s] to cache", users)

	uc.m.Lock()
	defer uc.m.Unlock()
	for _, user := range users {
		if user == nil || user.Id == "" {
			return errors.New("a user cannot be cached without an ID")
		}
		uc.users[user.Id] = user
	}

	return nil
}

func (uc *UserCache) RemoveUser(id string) {
	uc.logger.Debugf("removing user %s from cache", id)
	uc.m.Lock()
	defer uc.m.Unlock()
	uc.users[id] = nil
}
