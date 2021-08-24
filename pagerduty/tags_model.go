package pagerduty

type Tag struct {
	ApiObject
	Label string `json:"label"`
}
