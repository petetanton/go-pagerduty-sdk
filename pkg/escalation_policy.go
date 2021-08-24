package pkg

import (
	"encoding/json"
	"fmt"
)

func (c *Client) CreateEscalationPolicy(policy *EscalationPolicy) (*EscalationPolicy, error) {
	policy.Type = TypeEscalationPolicy
	reader, err := c.objectToJson(policy, TypeEscalationPolicy)
	if err != nil {
		return nil, err
	}

	response, err := c.post(fmt.Sprintf("%s/%s", c.cfg.ApiUrl, TypeEscalationPolicies), reader)
	if err != nil {
		return nil, err
	}

	var out *EscalationPolicy
	err = response.unmarshallResponse(&out, TypeEscalationPolicy)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) DeleteEscalationPolicy(id string) error {
	return c.delete(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeEscalationPolicies, id))
}

func (c *Client) GetEscalationPolicy(id string) (*EscalationPolicy, error) {
	response, err := c.get(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeEscalationPolicies, id), DefaultPagerDutyRequest())
	if err != nil {
		return nil, err
	}

	var escalationPolicy *EscalationPolicy
	err = response.unmarshallResponse(&escalationPolicy, TypeEscalationPolicy)
	if err != nil {
		return nil, err
	}

	return escalationPolicy, nil
}

func (c *Client) UpdateEscalationPolicy(policy *EscalationPolicy) (*EscalationPolicy, error) {
	policy.Type = TypeEscalationPolicy
	reader, err := c.objectToJson(policy, TypeEscalationPolicy)
	if err != nil {
		return nil, err
	}

	response, err := c.put(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeEscalationPolicies, policy.Id), reader)
	if err != nil {
		return nil, err
	}

	var out *EscalationPolicy
	err = response.unmarshallResponse(&out, TypeEscalationPolicy)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) ListEscalationPolicies() ([]*EscalationPolicy, error) {
	var escalationPolicies []*EscalationPolicy
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

		var responseEscalationPolicies []*EscalationPolicy

		err = json.Unmarshal(response.body.Path(TypeEscalationPolicies).Bytes(), &responseEscalationPolicies)
		if err != nil {
			return nil, err
		}

		escalationPolicies = append(escalationPolicies, responseEscalationPolicies...)
	}

	return escalationPolicies, nil
}
