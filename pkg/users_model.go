package pkg

type User struct {
	ApiObject
	Name              string       `json:"name,omitempty"`
	Email             string       `json:"email,omitempty"`
	Role              string       `json:"role,omitempty"`
	Description       string       `json:"description,omitempty"`
	JobTitle          string       `json:"job_title,omitempty"`
	Teams             []*Team      `json:"teams,omitempty"`
	ContactMethods    []*ApiObject `json:"contact_methods,omitempty"`
	NotificationRules []*ApiObject `json:"notification_rules"`
}
