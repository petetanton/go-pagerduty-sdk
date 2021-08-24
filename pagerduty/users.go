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

func (c *Client) GetUser() {
	//	https://developer.pagerduty.com/api-reference/reference/REST/openapiv3.json/paths/~1users~1%7Bid%7D/get
	panic("not implemented")
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

func (c *Client) UpdateUser() {
	//	https://developer.pagerduty.com/api-reference/reference/REST/openapiv3.json/paths/~1users~1%7Bid%7D/put
	panic("not implemented")
}

func (c *Client) DeleteUser(id string) error {
	return c.delete(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeUsers, id))
}
