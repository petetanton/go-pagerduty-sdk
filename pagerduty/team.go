package pagerduty

import (
	"errors"
	"fmt"
	"strings"
)

func (c *Client) GetTeam(id string) (*Team, error) {
	response, err := c.get(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeTeams, id), DefaultPagerDutyRequest())
	if err != nil {
		return nil, err
	}

	var team *Team
	err = response.unmarshallResponse(&team, TypeTeam)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (c *Client) GetTeamMembers(id string) ([]*TeamMembership, error) {
	response, err := c.get(fmt.Sprintf("%s/%s/%s/members", c.cfg.ApiUrl, TypeTeams, id), &PagerDutyRequest{Limit: 100, Includes: []string{"users"}})
	if err != nil {
		return nil, err
	}

	if response.hasMore() {
		return nil, errors.New("please implement pagination")
	}

	var teamMemberships []*TeamMembership

	err = response.unmarshallResponse(&teamMemberships, "members")
	if err != nil {
		return nil, err
	}

	return teamMemberships, nil
}

func (c *Client) AddUserToTeam(teamId, userId, role string) error {
	_, err := c.put(fmt.Sprintf("%s/%s/%s/%s/%s", c.cfg.ApiUrl, TypeTeams, teamId, TypeUsers, userId), strings.NewReader(fmt.Sprintf("{\"role\":\"%s\"}", role)))
	return err
}

func (c *Client) RemoveUserFromTeam(teamId, userId string) error {
	return c.delete(fmt.Sprintf("%s/%s/%s/%s/%s", c.cfg.ApiUrl, TypeTeams, teamId, TypeUsers, userId))
}
