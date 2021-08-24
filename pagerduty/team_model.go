package pagerduty

type Team struct {
	ApiObject
}

type TeamMembership struct {
	User *User  `json:"user"`
	Role string `json:"role"`
}
