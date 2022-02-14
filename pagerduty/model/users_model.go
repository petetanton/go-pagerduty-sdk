package model

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

type NotificationRule struct {
	ApiObject
	StartDelayInMinutes uint       `json:"start_delay_in_minutes"`
	ContactMethod       *ApiObject `json:"contact_method"`
	Urgency             string     `json:"urgency"`
}

type ContactMethod struct {
	ApiObject
	Label       string `json:"label,omitempty"`
	Address     string `json:"address,omitempty"`
	BlackListed bool   `json:"blacklisted,omitempty"`

	// Email contact method options
	SendShortEmail bool `json:"send_short_email,omitempty"`

	// Phone contact method options
	CountryCode int  `json:"country_code,omitempty"`
	Enabled     bool `json:"enabled,omitempty"`

	// Push contact method options
	DeviceType string                    `json:"device_type,omitempty"`
	Sounds     []*PushContactMethodSound `json:"sounds,omitempty"`
	CreatedAt  string                    `json:"created_at,omitempty"`
}

type PushContactMethodSound struct {
	Type string `json:"type,omitempty"`
	File string `json:"file,omitempty"`
}

type OnCall struct {
	UserReference             ApiObject `json:"user"`
	ScheduleReference         ApiObject `json:"schedule"`
	EscalationPolicyReference ApiObject `json:"escalation_policy"`
	Start                     string    `json:"start"`
	End                       string    `json:"end"`
	EscalationLevel           uint      `json:"escalation_level"`
}
