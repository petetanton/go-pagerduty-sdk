package pagerduty

import (
	"encoding/json"
	"fmt"
	"time"
)

type onCallRequest struct {
	PagerDutyRequest
	UserIds []string `url:"user_ids,omitempty,brackets"`
	Since   string   `url:"since,omitempty"`
	Until   string   `url:"until,omitempty"`
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
			Until:   until.String(),
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
