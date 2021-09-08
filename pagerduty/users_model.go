package pagerduty

type User struct {
	ApiObject
	Name              string       `json:"name,omitempty"`
	Email             string       `json:"email,omitempty"`
	TimeZone          string       `json:"time_zone,omitempty"`
	Color             string       `json:"color,omitempty"`
	AvatarUrl         string       `json:"avatar_url,omitempty"`
	Role              string       `json:"role,omitempty"`
	Description       string       `json:"description,omitempty"`
	InvitationSent    *bool        `json:"invitation_sent,omitempty"`
	JobTitle          string       `json:"job_title,omitempty"`
	Teams             []*Team      `json:"teams,omitempty"`
	ContactMethods    []*ApiObject `json:"contact_methods,omitempty"`
	NotificationRules []*ApiObject `json:"notification_rules"`
}
