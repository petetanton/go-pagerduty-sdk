package pagerduty

import (
	"fmt"
	"github.com/petetanton/go-pagerduty-sdk/pagerduty/model"
	"strings"
)

func (c *Client) CreateResponsePlay(responsePlay *model.ResponsePlay) (*model.ResponsePlay, error) {
	reader, err := c.objectToJson(responsePlay, TypeResponsePlay)
	if err != nil {
		return nil, err
	}

	response, err := c.post(fmt.Sprintf("%s/%s", c.cfg.ApiUrl, TypeResponsePlays), reader)
	if err != nil {
		return nil, err
	}

	var out *model.ResponsePlay
	err = response.unmarshallResponse(&out, TypeResponsePlay)

	return out, err
}

func (c *Client) DeleteResponsePlay(responsePlayId string) error {
	return c.delete(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeResponsePlays, responsePlayId))
}

func (c *Client) GetResponsePlay(id string) (*model.ResponsePlay, error) {
	response, err := c.get(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeResponsePlays, id), DefaultPagerDutyRequest())
	if err != nil {
		return nil, err
	}

	var responsePlay *model.ResponsePlay
	err = response.unmarshallResponse(&responsePlay, TypeResponsePlays)
	if err != nil {
		return nil, err
	}

	return responsePlay, nil
}

func (c *Client) ListResponsePlays() ([]*model.ResponsePlay, error) {
	var responsePlays []*model.ResponsePlay
	var response = &PagerDutyResponse{}
	var err error

	for response.hasMore() {
		response, err = c.get(fmt.Sprintf("%s/%s", c.cfg.ApiUrl, TypeResponsePlays), PagerDutyRequest{
			Limit:  100,
			Offset: response.nextOffset(),
		})
		if err != nil {
			return nil, err
		}

		var innerResponsePlays []*model.ResponsePlay
		err = response.unmarshallResponse(&innerResponsePlays, TypeResponsePlays)
		if err != nil {
			return nil, err
		}

		responsePlays = append(responsePlays, innerResponsePlays...)
	}

	return responsePlays, err
}

func (c *Client) RunResponsePlay(responsePlayId, incidentId string) error {
	bodyStr := fmt.Sprintf("{\"incident\":{\"id\":\"%s\",\"type\":\"incident_reference\"}}", incidentId)

	response, err := c.post(fmt.Sprintf("%s/%s/%s/run", c.cfg.ApiUrl, TypeResponsePlays, responsePlayId), strings.NewReader(bodyStr))
	if err != nil {
		return err
	}

	if response.body.Path("status").String() != "ok" {
		return fmt.Errorf("tried to run response play %s on incident %s but got response from PD: %s", responsePlayId, incidentId, response.body.Path("status").String())
	}

	return nil
}

func (c *Client) UpdateResponsePlay(responsePlay *model.ResponsePlay) (*model.ResponsePlay, error) {
	responsePlay.Type = TypeResponsePlay
	reader, err := c.objectToJson(responsePlay, TypeResponsePlay)
	if err != nil {
		return nil, err
	}

	response, err := c.put(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeResponsePlays, responsePlay.Id), reader)
	if err != nil {
		return nil, err
	}

	var out *model.ResponsePlay
	err = response.unmarshallResponse(&out, TypeResponsePlay)
	if err != nil {
		return nil, err
	}

	return out, nil
}
