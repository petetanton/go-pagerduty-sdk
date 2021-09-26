package pagerduty

import "fmt"

func (c *Client) CreateSchedule(schedule *Schedule, overflow bool) (*Schedule, error) {
	reader, err := c.objectToJson(schedule, TypeSchedule)
	if err != nil {
		return nil, err
	}

	response, err := c.post(fmt.Sprintf("%s/%s?overflow=%t", c.cfg.ApiUrl, TypeSchedules, overflow), reader)
	if err != nil {
		return nil, err
	}

	var out *Schedule
	err = response.unmarshallResponse(&out, TypeSchedule)

	return out, err
}

func (c *Client) GetSchedule(id string) (*Schedule, error) {
	response, err := c.get(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeSchedules, id), DefaultPagerDutyRequest())
	if err != nil {
		return nil, err
	}

	var service *Schedule
	err = response.unmarshallResponse(&service, TypeSchedule)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (c *Client) ListSchedule() ([]*Schedule, error) {
	var services []*Schedule
	var response = &PagerDutyResponse{}
	var err error

	for response.hasMore() {
		response, err = c.get(fmt.Sprintf("%s/%s", c.cfg.ApiUrl, TypeSchedules), PagerDutyRequest{
			Limit:  100,
			Offset: response.nextOffset(),
		})
		if err != nil {
			return nil, err
		}

		var innerServices []*Schedule
		err = response.unmarshallResponse(&innerServices, TypeSchedules)
		if err != nil {
			return nil, err
		}

		services = append(services, innerServices...)
	}

	return services, err
}

func (c *Client) UpdateSchedule(service *Schedule, overflow bool) (*Schedule, error) {
	service.Type = TypeSchedule
	reader, err := c.objectToJson(service, TypeSchedule)
	if err != nil {
		return nil, err
	}

	response, err := c.put(fmt.Sprintf("%s/%s/%s?overflow=%t", c.cfg.ApiUrl, TypeSchedules, service.Id, overflow), reader)
	if err != nil {
		return nil, err
	}

	var out *Schedule
	err = response.unmarshallResponse(&out, TypeSchedule)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) DeleteSchedule(id string) error {
	return c.delete(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeSchedules, id))
}
