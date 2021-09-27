package pagerduty

import (
	"encoding/json"
	"fmt"
	"github.com/petetanton/go-pagerduty-sdk/pagerduty/model"
	"time"
)

type onCallRequest struct {
	PagerDutyRequest
	UserIds []string `url:"user_ids,omitempty,brackets"`
	Since   string   `url:"since,omitempty"`
	Until   string   `url:"until,omitempty"`
}

func (c *Client) ListOnCallsForUsers(userIds []string, since, until time.Time) ([]*model.OnCall, error) {
	var onCalls []*model.OnCall
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

		var responseOncalls []*model.OnCall

		err = json.Unmarshal(response.body.Path("oncalls").Bytes(), &responseOncalls)
		if err != nil {
			return nil, err
		}

		onCalls = append(onCalls, responseOncalls...)
	}

	return onCalls, nil
}
