package pkg

import (
	"encoding/json"
	"fmt"
)

func (c *Client) ListTags() ([]*Tag, error) {
	var tags []*Tag
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

		var responseTags []*Tag

		err = json.Unmarshal(response.body.Path(TypeTags).Bytes(), &responseTags)
		if err != nil {
			return nil, err
		}

		tags = append(tags, responseTags...)
	}

	return tags, nil
}

func (c *Client) GetTaggedEntities(tagId, entityType string) ([]*ApiObject, error) {
	var output []*ApiObject
	var response = &PagerDutyResponse{}
	var err error

	for response.hasMore() {
		response, err = c.get(fmt.Sprintf("%s/%s/%s/%s", c.cfg.ApiToken, TypeTags, tagId, entityType), PagerDutyRequest{

			Offset: response.nextOffset(),
			Limit:  100,
		})
		if err != nil {
			return nil, err
		}

		var responseOutput []*ApiObject

		err = json.Unmarshal(response.body.Path(entityType).Bytes(), &responseOutput)
		if err != nil {
			return nil, err
		}

		output = append(output, responseOutput...)
	}

	return output, nil
}
