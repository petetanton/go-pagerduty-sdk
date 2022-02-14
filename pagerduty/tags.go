package pagerduty

import (
	"encoding/json"
	"fmt"
	"github.com/petetanton/go-pagerduty-sdk/pagerduty/model"
)

func (c *Client) ListTags() ([]*model.Tag, error) {
	var tags []*model.Tag
	var response = &PagerDutyResponse{}
	var err error

	for response.hasMore() {
		response, err = c.get(fmt.Sprintf("%s/%s", c.cfg.ApiUrl, TypeTags), PagerDutyRequest{
			Offset: response.nextOffset(),
			Limit:  100,
		})
		if err != nil {
			return nil, err
		}

		var responseTags []*model.Tag

		err = json.Unmarshal(response.body.Path(TypeTags).Bytes(), &responseTags)
		if err != nil {
			return nil, err
		}

		tags = append(tags, responseTags...)
	}

	return tags, nil
}

func (c *Client) GetTaggedEntities(tagId, entityType string) ([]*model.ApiObject, error) {
	var output []*model.ApiObject
	var response = &PagerDutyResponse{}
	var err error

	for response.hasMore() {
		response, err = c.get(fmt.Sprintf("%s/%s/%s/%s", c.cfg.ApiUrl, TypeTags, tagId, entityType), PagerDutyRequest{

			Offset: response.nextOffset(),
			Limit:  100,
		})
		if err != nil {
			return nil, err
		}

		var responseOutput []*model.ApiObject

		err = json.Unmarshal(response.body.Path(entityType).Bytes(), &responseOutput)
		if err != nil {
			return nil, err
		}

		output = append(output, responseOutput...)
	}

	return output, nil
}

func (c *Client) GetTagsOnEntity(entityId, entityType string) ([]*model.Tag, error) {
	var output []*model.Tag
	var response = &PagerDutyResponse{}
	var err error

	response, err = c.get(fmt.Sprintf("%s/%s/%s/%s", c.cfg.ApiUrl, entityType, entityId, TypeTags), PagerDutyRequest{

		Offset: response.nextOffset(),
		Limit:  100,
	})
	if err != nil {
		return nil, err
	}

	var responseOutput []*model.Tag

	err = json.Unmarshal(response.body.Path(TypeTags).Bytes(), &responseOutput)
	if err != nil {
		return nil, err
	}

	output = append(output, responseOutput...)

	return output, nil
}
