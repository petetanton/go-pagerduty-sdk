package pagerduty

import (
	"fmt"
)

func (c *Client) CreateUser(user *User) (*User, error) {
	reader, err := c.objectToJson(user, TypeUser)
	if err != nil {
		return nil, err
	}

	response, err := c.post(fmt.Sprintf("%s/%s", c.cfg.ApiUrl, TypeUsers), reader)
	if err != nil {
		return nil, err
	}

	var out *User
	err = response.unmarshallResponse(&out, TypeUser)

	return out, err
}

func (c *Client) GetUser(id string) (*User, error) {
	response, err := c.get(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeUsers, id), DefaultPagerDutyRequest())
	if err != nil {
		return nil, err
	}

	var user *User
	err = response.unmarshallResponse(&user, TypeUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *Client) ListUsers() ([]*User, error) {
	var users []*User
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

		var innerUser []*User
		err = response.unmarshallResponse(&innerUser, TypeUsers)
		if err != nil {
			return nil, err
		}

		users = append(users, innerUser...)
	}

	return users, err
}

func (c *Client) UpdateUser(user *User) (*User, error) {
	user.Type = TypeUser
	reader, err := c.objectToJson(user, TypeUser)
	if err != nil {
		return nil, err
	}

	response, err := c.put(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeUsers, user.Id), reader)
	if err != nil {
		return nil, err
	}

	var out *User
	err = response.unmarshallResponse(&out, TypeUser)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) DeleteUser(id string) error {
	return c.delete(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeUsers, id))
}

func (c *Client) CreateUserNotificationRule(userId string, rule *NotificationRule) (*NotificationRule, error) {
	reader, err := c.objectToJson(rule, TypeNotificationRule)
	if err != nil {
		return nil, err
	}

	response, err := c.post(fmt.Sprintf("%s/%s/%s/%s", c.cfg.ApiUrl, TypeUsers, userId, TypeNotificationRules), reader)
	if err != nil {
		return nil, err
	}

	var out *NotificationRule
	err = response.unmarshallResponse(&out, TypeNotificationRule)

	return out, err
}

func (c *Client) GetUserNotificationRule(userId, notificationRuleId string) (*NotificationRule, error) {
	response, err := c.get(fmt.Sprintf("%s/%s/%s/%s/%s", c.cfg.ApiUrl, TypeUsers, userId, TypeNotificationRules, notificationRuleId), DefaultPagerDutyRequest())
	if err != nil {
		return nil, err
	}

	var out *NotificationRule
	err = response.unmarshallResponse(&out, TypeNotificationRule)

	return out, err
}

func (c *Client) ListUsersNotificationRules(userId string) ([]*NotificationRule, error) {
	var notificationRules []*NotificationRule
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

		var innerRule []*NotificationRule
		err = response.unmarshallResponse(&innerRule, TypeNotificationRules)
		if err != nil {
			return nil, err
		}

		notificationRules = append(notificationRules, innerRule...)
	}

	return notificationRules, err
}

func (c *Client) UpdateNotificationRule(user *User, notificationRule *NotificationRule) (*NotificationRule, error) {
	reader, err := c.objectToJson(notificationRule, TypeNotificationRule)
	if err != nil {
		return nil, err
	}

	response, err := c.put(fmt.Sprintf("%s/%s/%s/%s", c.cfg.ApiUrl, TypeUsers, user.Id, TypeNotificationRules), reader)
	if err != nil {
		return nil, err
	}

	var out *NotificationRule
	err = response.unmarshallResponse(&out, TypeNotificationRule)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) DeleteNotificationRule(userId, notificationRuleId string) error {
	return c.delete(fmt.Sprintf("%s/%s/%s/%s/%s", c.cfg.ApiUrl, TypeUsers, userId, TypeNotificationRules, notificationRuleId))
}

func (c *Client) CreateUserContactMethod(userId string, contactMethod *ContactMethod) (*ContactMethod, error) {
	reader, err := c.objectToJson(contactMethod, TypeContactMethod)
	if err != nil {
		return nil, err
	}

	response, err := c.post(fmt.Sprintf("%s/%s/%s/%s", c.cfg.ApiUrl, TypeUsers, userId, TypeContactMethods), reader)
	if err != nil {
		return nil, err
	}

	var out *ContactMethod
	err = response.unmarshallResponse(&out, TypeContactMethod)

	return out, err
}

func (c *Client) UpdateUserContactMethod(user *User, contactMethod *ContactMethod) (*ContactMethod, error) {
	reader, err := c.objectToJson(contactMethod, TypeContactMethod)
	if err != nil {
		return nil, err
	}

	response, err := c.put(fmt.Sprintf("%s/%s/%s/%s", c.cfg.ApiUrl, TypeUsers, user.Id, TypeContactMethods), reader)
	if err != nil {
		return nil, err
	}

	var out *ContactMethod
	err = response.unmarshallResponse(&out, TypeContactMethod)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) DeleteUserContactMethod(userId, contactMethodId string) error {
	return c.delete(fmt.Sprintf("%s/%s/%s/%s/%s", c.cfg.ApiUrl, TypeUsers, userId, TypeContactMethods, contactMethodId))
}

func (c *Client) GetUserContactMethod(userId, contactMethodId string) (*ContactMethod, error) {
	response, err := c.get(fmt.Sprintf("%s/%s/%s/%s/%s", c.cfg.ApiUrl, TypeUsers, userId, TypeContactMethods, contactMethodId), DefaultPagerDutyRequest())
	if err != nil {
		return nil, err
	}

	var out *ContactMethod
	err = response.unmarshallResponse(&out, TypeContactMethod)

	return out, err
}
