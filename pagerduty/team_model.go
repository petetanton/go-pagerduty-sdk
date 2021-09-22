package pagerduty

type Team struct {
	ApiObject
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TeamMembership struct {
	User *User  `json:"user"`
	Role string `json:"role"`
}
