package model

type Team struct {
	ApiObject
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type TeamMembership struct {
	User *User  `json:"user"`
	Role string `json:"role"`
}

type ListTeamMembersResponse struct {
	ApiListObject
	Members []*TeamMembership `json:"members"`
}
