package cache

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/petetanton/go-pagerduty-sdk/pagerduty/model"
)

func Test_UserCache(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	t.Run("it returns null if not in cache", func(t *testing.T) {
		uc := NewUserCache(logger.WithField("test", t.Name()))
		user := uc.ReadUser("does-not-exist")
		assert.Nil(t, user, "expected user to be nil")
	})

	t.Run("it caches a user and then returns it", func(t *testing.T) {
		uc := NewUserCache(logger.WithField("test", t.Name()))
		user := model.User{
			ApiObject: model.ApiObject{
				Id: "someId",
			},
			Name: "Real cool person",
		}
		err := uc.WriteUser(&user)
		require.NoError(t, err)
		readUser := uc.ReadUser("someId")

		assert.Equal(t, user.Id, readUser.Id)
		assert.Equal(t, user.Name, readUser.Name)
	})

	t.Run("it does not cache a user without an ID", func(t *testing.T) {
		uc := NewUserCache(logger.WithField("test", t.Name()))
		user := model.User{
			Name: "Real cool person",
		}
		err := uc.WriteUser(&user)
		assert.Error(t, err, "some interesting err")
	})
}
