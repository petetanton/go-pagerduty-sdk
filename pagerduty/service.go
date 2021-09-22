package pagerduty

import "fmt"

func (c *Client) CreateService(service *Service) (*Service, error) {
	reader, err := c.objectToJson(service, TypeService)
	if err != nil {
		return nil, err
	}

	response, err := c.post(fmt.Sprintf("%s/%s", c.cfg.ApiUrl, TypeServices), reader)
	if err != nil {
		return nil, err
	}

	var out *Service
	err = response.unmarshallResponse(&out, TypeService)

	return out, err
}

func (c *Client) GetService(id string) (*Service, error) {
	response, err := c.get(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeServices, id), DefaultPagerDutyRequest())
	if err != nil {
		return nil, err
	}

	var service *Service
	err = response.unmarshallResponse(&service, TypeService)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (c *Client) ListServices() ([]*Service, error) {
	var services []*Service
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

		var innerServices []*Service
		err = response.unmarshallResponse(&innerServices, TypeServices)
		if err != nil {
			return nil, err
		}

		services = append(services, innerServices...)
	}

	return services, err
}

func (c *Client) UpdateService(service *Service) (*Service, error) {
	service.Type = TypeService
	reader, err := c.objectToJson(service, TypeService)
	if err != nil {
		return nil, err
	}

	response, err := c.put(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeServices, service.Id), reader)
	if err != nil {
		return nil, err
	}

	var out *Service
	err = response.unmarshallResponse(&out, TypeService)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) DeleteService(id string) error {
	return c.delete(fmt.Sprintf("%s/%s/%s", c.cfg.ApiUrl, TypeServices, id))
}
