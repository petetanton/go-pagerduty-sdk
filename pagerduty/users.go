package pagerduty

import (
	"fmt"
	"github.com/petetanton/go-pagerduty-sdk/pagerduty/model"
)

func (c *Client) CreateUser(user *model.User) (*model.User, error) {
	reader, err := c.objectToJson(user, TypeUser)
	if err != nil {
		return nil, err
	}

	response, err := c.post(fmt.Sprintf("%s/%s", c.cfg.ApiUrl, TypeUsers), reader)
	if err != nil {
		return nil, err
	}

	if response == nil {
		return nil, fmt.Errorf("cannot create user %v", user)
	}

	var out *model.User
	err = response.unmarshallResponse(&out, TypeUser)
	if err != nil {
		return nil, err
	}

	return out, c.userCache.WriteUser(out)
}

func (c *Client) GetUser(id string) (*model.User, error) {
	if c.shouldCacheType(TypeUser) {
		user := c.userCache.ReadUser(id)
		if user != nil {
			return user, nil
		}
	}

	response, err := c.get(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeUsers, id), DefaultPagerDutyRequest())
	if err != nil {
		return nil, err
	}

	if response == nil {
		c.cfg.Logger.Warnf("could not find user %s", id)
		return nil, nil
	}

	var user *model.User
	err = response.unmarshallResponse(&user, TypeUser)
	if err != nil {
		return nil, err
	}

	return user, c.userCache.WriteUser(user)
}

func (c *Client) ListUsers() ([]*model.User, error) {
	var users []*model.User
	var response = &PagerDutyResponse{}
	var err error

	for response.hasMore() {
		response, err = c.get(fmt.Sprintf("%s/%s", c.cfg.ApiUrl, TypeUsers), PagerDutyRequest{
			Limit:  100,
			Offset: response.nextOffset(),
		})
		if err != nil {
			return nil, err
		}

		var innerUser []*model.User
		err = response.unmarshallResponse(&innerUser, TypeUsers)
		if err != nil {
			return nil, err
		}

		users = append(users, innerUser...)
	}

	return users, c.userCache.WriteUsers(users)
}

func (c *Client) UpdateUser(user *model.User) (*model.User, error) {
	user.Type = TypeUser
	reader, err := c.objectToJson(user, TypeUser)
	if err != nil {
		return nil, err
	}

	response, err := c.put(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeUsers, user.Id), reader)
	if err != nil {
		return nil, err
	}

	var out *model.User
	err = response.unmarshallResponse(&out, TypeUser)
	if err != nil {
		return nil, err
	}

	return out, c.userCache.WriteUser(out)
}

func (c *Client) DeleteUser(id string) error {
	c.userCache.RemoveUser(id)
	return c.delete(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeUsers, id))
}

func (c *Client) CreateUserNotificationRule(userId string, rule *model.NotificationRule) (*model.NotificationRule, error) {
	reader, err := c.objectToJson(rule, TypeNotificationRule)
	if err != nil {
		return nil, err
	}

	response, err := c.post(fmt.Sprintf("%s/%s/%s/%s", c.cfg.ApiUrl, TypeUsers, userId, TypeNotificationRules), reader)
	if err != nil {
		return nil, err
	}

	var out *model.NotificationRule
	err = response.unmarshallResponse(&out, TypeNotificationRule)

	return out, err
}

func (c *Client) GetUserNotificationRule(userId, notificationRuleId string) (*model.NotificationRule, error) {
	response, err := c.get(fmt.Sprintf("%s/%s/%s/%s/%s", c.cfg.ApiUrl, TypeUsers, userId, TypeNotificationRules, notificationRuleId), DefaultPagerDutyRequest())
	if err != nil {
		return nil, err
	}

	var out *model.NotificationRule
	err = response.unmarshallResponse(&out, TypeNotificationRule)

	return out, err
}

func (c *Client) ListUsersNotificationRules(userId string) ([]*model.NotificationRule, error) {
	var notificationRules []*model.NotificationRule
	var response = &PagerDutyResponse{}
	var err error

	for response.hasMore() {
		response, err = c.get(fmt.Sprintf("%s/%s/%s/%s", c.cfg.ApiUrl, TypeUsers, userId, TypeNotificationRules), PagerDutyRequest{
			Limit:    100,
			Offset:   response.nextOffset(),
			Includes: []string{TypeContactMethods},
		})
		if err != nil {
			return nil, err
		}

		var innerRule []*model.NotificationRule
		err = response.unmarshallResponse(&innerRule, TypeNotificationRules)
		if err != nil {
			return nil, err
		}

		notificationRules = append(notificationRules, innerRule...)
	}

	return notificationRules, err
}

func (c *Client) UpdateNotificationRule(user *model.User, notificationRule *model.NotificationRule) (*model.NotificationRule, error) {
	reader, err := c.objectToJson(notificationRule, TypeNotificationRule)
	if err != nil {
		return nil, err
	}

	response, err := c.put(fmt.Sprintf("%s/%s/%s/%s", c.cfg.ApiUrl, TypeUsers, user.Id, TypeNotificationRules), reader)
	if err != nil {
		return nil, err
	}

	var out *model.NotificationRule
	err = response.unmarshallResponse(&out, TypeNotificationRule)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) DeleteNotificationRule(userId, notificationRuleId string) error {
	return c.delete(fmt.Sprintf("%s/%s/%s/%s/%s", c.cfg.ApiUrl, TypeUsers, userId, TypeNotificationRules, notificationRuleId))
}

func (c *Client) CreateUserContactMethod(userId string, contactMethod *model.ContactMethod) (*model.ContactMethod, error) {
	reader, err := c.objectToJson(contactMethod, TypeContactMethod)
	if err != nil {
		return nil, err
	}

	response, err := c.post(fmt.Sprintf("%s/%s/%s/%s", c.cfg.ApiUrl, TypeUsers, userId, TypeContactMethods), reader)
	if err != nil {
		return nil, err
	}

	var out *model.ContactMethod
	err = response.unmarshallResponse(&out, TypeContactMethod)

	return out, err
}

func (c *Client) UpdateUserContactMethod(user *model.User, contactMethod *model.ContactMethod) (*model.ContactMethod, error) {
	reader, err := c.objectToJson(contactMethod, TypeContactMethod)
	if err != nil {
		return nil, err
	}

	response, err := c.put(fmt.Sprintf("%s/%s/%s/%s", c.cfg.ApiUrl, TypeUsers, user.Id, TypeContactMethods), reader)
	if err != nil {
		return nil, err
	}

	var out *model.ContactMethod
	err = response.unmarshallResponse(&out, TypeContactMethod)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) DeleteUserContactMethod(userId, contactMethodId string) error {
	return c.delete(fmt.Sprintf("%s/%s/%s/%s/%s", c.cfg.ApiUrl, TypeUsers, userId, TypeContactMethods, contactMethodId))
}

func (c *Client) GetUserContactMethod(userId, contactMethodId string) (*model.ContactMethod, error) {
	response, err := c.get(fmt.Sprintf("%s/%s/%s/%s/%s", c.cfg.ApiUrl, TypeUsers, userId, TypeContactMethods, contactMethodId), DefaultPagerDutyRequest())
	if err != nil {
		return nil, err
	}

	var out *model.ContactMethod
	err = response.unmarshallResponse(&out, TypeContactMethod)

	return out, err
}
