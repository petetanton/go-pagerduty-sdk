package pagerduty

type ResponsePlay struct {
	ApiObject
	Name               string       `json:"name"`
	Description        string       `json:"description"`
	Team               ApiObject    `json:"team,omitempty"`
	Subscribers        []*ApiObject `json:"subscribers,omitempty"`
	SubscribersMessage string       `json:"subscribers_message,omitempty"`
	Responders         []*ApiObject `json:"responders,omitempty"`
	RespondersMessage  string       `json:"responders_message,omitempty"`
	Runability         string       `json:"runability"`
	ConferenceNumber   string       `json:"conference_number,omitempty"`
	ConferenceUrl      string       `json:"conference_url,omitempty"`
	ConferenceType     string       `json:"conference_type,omitempty"`
}
