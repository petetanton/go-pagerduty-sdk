package model

type ResponsePlay struct {
	ApiObject
	Name               string       `json:"name"`
	Description        string       `json:"description"`
	Team               *ApiObject   `json:"team,omitempty"`
	Subscribers        []*ApiObject `json:"subscribers,omitempty"`
	SubscribersMessage string       `json:"subscribers_message,omitempty"`
	Responders         []*Responder `json:"responders,omitempty"`
	RespondersMessage  string       `json:"responders_message,omitempty"`
	Runnability        string       `json:"runnability"`
	ConferenceNumber   string       `json:"conference_number,omitempty"`
	ConferenceUrl      string       `json:"conference_url,omitempty"`
	ConferenceType     string       `json:"conference_type,omitempty"`
}

type Responder struct {
	Type                       string            `json:"type,omitempty"`
	Id                         string            `json:"id,omitempty"`
	Name                       string            `json:"name,omitempty"`
	Description                string            `json:"description,omitempty"`
	NumLoops                   int               `json:"num_loops,omitempty"`
	OnCallHandoffNotifications string            `json:"on_call_handoff_notifications,omitempty"`
	EscalationRules            []*EscalationRule `json:"escalation_rules,omitempty"`
	Services                   []*ApiObject      `json:"services,omitempty"`
	Teams                      []*ApiObject      `json:"teams,omitempty"`
}
