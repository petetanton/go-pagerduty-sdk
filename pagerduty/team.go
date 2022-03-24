package pagerduty

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/petetanton/go-pagerduty-sdk/pagerduty/model"
)

func (c *Client) CreateTeam(team *model.Team) (*model.Team, error) {
	reader, err := c.objectToJson(team, TypeTeam)
	if err != nil {
		return nil, err
	}

	response, err := c.post(fmt.Sprintf("%s/%s", c.cfg.ApiUrl, TypeTeams), reader)
	if err != nil {
		return nil, err
	}

	if response == nil {
		return nil, nil
	}

	var out *model.Team
	err = response.unmarshallResponse(&out, TypeTeam)

	return out, err
}

func (c *Client) GetTeam(id string) (*model.Team, error) {
	response, err := c.get(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeTeams, id), DefaultPagerDutyRequest())
	if err != nil {
		return nil, err
	}

	if response == nil {
		return nil, nil
	}

	var team *model.Team
	err = response.unmarshallResponse(&team, TypeTeam)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (c *Client) GetTeamMembers(id string) ([]*model.TeamMembership, error) {
	response, err := c.get(fmt.Sprintf("%s/%s/%s/members", c.cfg.ApiUrl, TypeTeams, id), &PagerDutyRequest{Limit: 100, Includes: []string{"users"}})
	if err != nil {
		return nil, err
	}

	if response.hasMore() {
		return nil, errors.New("please implement pagination")
	}

	var teamMemberships []*model.TeamMembership

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

func (c *Client) UpdateTeam(team *model.Team) (*model.Team, error) {
	team.Type = TypeTeam
	reader, err := c.objectToJson(team, TypeTeam)
	if err != nil {
		return nil, err
	}

	response, err := c.put(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeTeams, team.Id), reader)
	if err != nil {
		return nil, err
	}

	if response == nil {
		return nil, nil
	}

	var out *model.Team
	err = response.unmarshallResponse(&out, TypeTeam)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) DeleteTeam(id string) error {
	return c.delete(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeTeams, id))
}

func (c *Client) ListTeams() ([]*model.Team, error) {
	var teams []*model.Team
	var response = &PagerDutyResponse{}
	var err error

	for response.hasMore() {
		response, err = c.get(fmt.Sprintf("%s/%s", c.cfg.ApiUrl, TypeTeams), PagerDutyRequest{
			Limit:  100,
			Offset: response.nextOffset(),
		})
		if err != nil {
			return nil, err
		}

		var innerTeam []*model.Team
		err = response.unmarshallResponse(&innerTeam, TypeTeams)
		if err != nil {
			return nil, err
		}

		teams = append(teams, innerTeam...)
	}

	return teams, err
}

func (c *Client) AddEscalationPolicyToTeam(teamId, epId string) error {
	_, err := c.put(fmt.Sprintf("%s/%s/%s/%s/%s", c.cfg.ApiUrl, TypeTeams, teamId, TypeEscalationPolicies, epId), nil)
	return err
}

func (c *Client) ListTeamMembers(teamId string) ([]*model.TeamMembership, error) {
	var members []*model.TeamMembership
	var response = &PagerDutyResponse{}
	var err error

	for response.hasMore() {
		response, err = c.get(fmt.Sprintf("%s/%s/%s/members", c.cfg.ApiUrl, TypeTeams, teamId), PagerDutyRequest{
			Limit:  100,
			Offset: response.nextOffset(),
		})
		if err != nil {
			return nil, err
		}

		if response == nil {
			return nil, fmt.Errorf("got no memberships for team %s", teamId)
		}

		var result *model.ListTeamMembersResponse
		err := json.Unmarshal(response.body.Bytes(), &result)
		if err != nil {
			return nil, err
		}

		members = append(members, result.Members...)
	}

	return members, nil
}
