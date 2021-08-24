package pkg

type EscalationPolicy struct {
	ApiObject
	Name                       string            `json:"name"`
	EscalationRules            []*EscalationRule `json:"escalation_rules,omitempty"`
	Services                   []*Service        `json:"services"`
	NumLoops                   uint              `json:"num_loops"`
	Teams                      []*Team           `json:"teams"`
	Description                string            `json:"description"`
	OnCallHandoffNotifications string            `json:"on_call_handoff_notifications"`
	//	privilege
}

type EscalationRule struct {
	Id                       string                  `json:"id,omitempty"`
	EscalationDelayInMinutes uint                    `json:"escalation_delay_in_minutes"`
	Targets                  []*EscalationRuleTarget `json:"targets"`
}

type EscalationRuleTarget struct {
	ApiObject
}
