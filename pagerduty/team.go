package pagerduty

import (
	"errors"
	"fmt"
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
