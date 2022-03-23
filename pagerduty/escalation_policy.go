package pagerduty

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/petetanton/go-pagerduty-sdk/pagerduty/model"
)

func (c *Client) CreateEscalationPolicy(policy *model.EscalationPolicy) (*model.EscalationPolicy, error) {
	policy.Type = TypeEscalationPolicy
	reader, err := c.objectToJson(policy, TypeEscalationPolicy)
	if err != nil {
		return nil, err
	}

	response, err := c.post(fmt.Sprintf("%s/%s", c.cfg.ApiUrl, TypeEscalationPolicies), reader)
	if err != nil {
		return nil, err
	}

	var out *model.EscalationPolicy
	err = response.unmarshallResponse(&out, TypeEscalationPolicy)
	if err != nil {
		return nil, err
	}

	c.escalationPolicyCache.WriteEscalationPolicy(policy)

	return out, nil
}

func (c *Client) DeleteEscalationPolicy(id string) error {
	c.escalationPolicyCache.RemoveEscalationPolicy(id)
	return c.delete(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeEscalationPolicies, id))
}

func (c *Client) GetEscalationPolicy(id string) (*model.EscalationPolicy, error) {
	if c.shouldCacheType(TypeEscalationPolicy) {
		policy := c.escalationPolicyCache.ReadEscalationPolicy(id)
		if policy != nil {
			return policy, nil
		}
	}

	response, err := c.get(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeEscalationPolicies, id), DefaultPagerDutyRequest())
	if err != nil {
		return nil, err
	}

	if response == nil {
		return nil, nil
	}

	var escalationPolicy *model.EscalationPolicy
	err = response.unmarshallResponse(&escalationPolicy, TypeEscalationPolicy)
	if err != nil {
		return nil, err
	}

	c.escalationPolicyCache.WriteEscalationPolicy(escalationPolicy)

	return escalationPolicy, nil
}

func (c *Client) UpdateEscalationPolicy(policy *model.EscalationPolicy) (*model.EscalationPolicy, error) {
	if policy.Id == "" {
		return nil, errors.New("id needs to be set on a policy to update it")
	}
	policy.Type = TypeEscalationPolicy
	reader, err := c.objectToJson(policy, TypeEscalationPolicy)
	if err != nil {
		return nil, err
	}

	response, err := c.put(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeEscalationPolicies, policy.Id), reader)
	if err != nil {
		return nil, err
	}

	var out *model.EscalationPolicy
	err = response.unmarshallResponse(&out, TypeEscalationPolicy)
	if err != nil {
		return nil, err
	}

	c.escalationPolicyCache.WriteEscalationPolicy(policy)

	return out, nil
}

func (c *Client) ListEscalationPolicies() ([]*model.EscalationPolicy, error) {
	var escalationPolicies []*model.EscalationPolicy
	var response = &PagerDutyResponse{}
	var err error

	for response.hasMore() {
		response, err = c.get(fmt.Sprintf("%s/%s", c.cfg.ApiUrl, TypeEscalationPolicies), PagerDutyRequest{
			Offset: response.nextOffset(),
			Limit:  100,
		})
		if err != nil {
			return nil, err
		}

		var responseEscalationPolicies []*model.EscalationPolicy

		err = json.Unmarshal(response.body.Path(TypeEscalationPolicies).Bytes(), &responseEscalationPolicies)
		if err != nil {
			return nil, err
		}

		escalationPolicies = append(escalationPolicies, responseEscalationPolicies...)
	}

	c.escalationPolicyCache.WriteEscalationPolicies(escalationPolicies)

	return escalationPolicies, nil
}
