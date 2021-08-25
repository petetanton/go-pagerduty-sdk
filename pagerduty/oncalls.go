package pagerduty

import (
	"encoding/json"
	"fmt"
	"time"
)

type onCallRequest struct {
	PagerDutyRequest
	UserIds []string `json:"user_ids,omitempty,brackets"`
	Since   string   `json:"since,omitempty"`
	Unitl   string   `json:"until,omitempty"`
}

type OnCall struct {
	UserReference             ApiObject `json:"user"`
	ScheduleReference         ApiObject `json:"schedule"`
	EscalationPolicyReference ApiObject `json:"escalation_policy"`
	Start                     string    `json:"start"`
	End                       string    `json:"end"`
}

func (c *Client) ListOnCallsForUsers(userIds []string, since, until time.Time) ([]*OnCall, error) {
	var onCalls []*OnCall
	var response = &PagerDutyResponse{}
	var err error

	for response.hasMore() {
		response, err = c.get(fmt.Sprintf("%s/oncalls", c.cfg.ApiUrl), onCallRequest{
			PagerDutyRequest: PagerDutyRequest{
				Offset:   response.nextOffset(),
				Limit:    100,
				Includes: nil,
			},
			UserIds: userIds,
			Since:   since.String(),
			Unitl:   until.String(),
		})
		if err != nil {
			return nil, err
		}

		var responseOncalls []*OnCall

		err = json.Unmarshal(response.body.Path("oncalls").Bytes(), &responseOncalls)
		if err != nil {
			return nil, err
		}

		onCalls = append(onCalls, responseOncalls...)
	}

	return onCalls, nil
}
