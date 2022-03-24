package pagerduty

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/petetanton/go-pagerduty-sdk/pagerduty/model"
)

func (c *Client) CreateService(service *model.Service) (*model.Service, error) {
	rBytes, err := json.Marshal(service)
	if err != nil {
		return nil, err
	}

	response, err := c.post(fmt.Sprintf("%s/%s", c.cfg.ApiUrl, TypeServices), bytes.NewReader(rBytes))
	if err != nil {
		return nil, fmt.Errorf("got an error when trying to create service: %s. Msg: %s", string(rBytes), err.Error())
	}

	if response == nil {
		return nil, nil
	}

	var out *model.Service
	err = response.unmarshallResponse(&out, TypeService)

	return out, err
}

func (c *Client) GetService(id string) (*model.Service, error) {
	response, err := c.get(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeServices, id), DefaultPagerDutyRequest())
	if err != nil {
		return nil, err
	}

	if response == nil {
		return nil, nil
	}

	var service *model.Service
	err = response.unmarshallResponse(&service, TypeService)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (c *Client) ListServices() ([]*model.Service, error) {
	var services []*model.Service
	var response = &PagerDutyResponse{}
	var err error

	for response.hasMore() {
		response, err = c.get(fmt.Sprintf("%s/%s", c.cfg.ApiUrl, TypeServices), PagerDutyRequest{
			Limit:  100,
			Offset: response.nextOffset(),
		})
		if err != nil {
			return nil, err
		}

		var innerServices []*model.Service
		err = response.unmarshallResponse(&innerServices, TypeServices)
		if err != nil {
			return nil, err
		}

		services = append(services, innerServices...)
	}

	return services, err
}

func (c *Client) UpdateService(service *model.Service) (*model.Service, error) {
	service.Type = TypeService
	reader, err := c.objectToJson(service, TypeService)
	if err != nil {
		return nil, err
	}

	response, err := c.put(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeServices, service.Id), reader)
	if err != nil {
		return nil, err
	}

	if response == nil {
		return nil, nil
	}

	var out *model.Service
	err = response.unmarshallResponse(&out, TypeService)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) DeleteService(id string) error {
	return c.delete(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeServices, id))
}
