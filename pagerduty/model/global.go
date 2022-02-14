package model

type ApiObject struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	Summary string `json:"summary"`
	Self    string `json:"self,omitempty"`
	HtmlUrl string `json:"html_url,omitempty"`
}

type ApiListObject struct {
	Limit  uint `json:"limit,omitempty"`
	Offset uint `json:"offset,omitempty"`
	More   bool `json:"more,omitempty"`
	Total  uint `json:"total,omitempty"`
}
