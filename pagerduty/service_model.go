package pagerduty

type Service struct {
	ApiObject
	Name                    string                   `json:"name"`
	Description             string                   `json:"description"`
	AutoResolveTimeout      uint                     `json:"auto_resolve_timeout"`
	AcknowledgementTimeout  uint                     `json:"acknowledgement_timeout"`
	Status                  string                   `json:"status"`
	EscalationPolicy        *ApiObject               `json:"escalation_policy,omitempty"`
	IncidentUrgencyRule     *IncidentUrgencyRule     `json:"incident_urgency_rule,omitempty"`
	SupportHours            *SupportHours            `json:"support_hours,omitempty"`
	ScheduledActions        []*ScheduledAction       `json:"scheduled_actions,omitempty"`
	AlertCreation           string                   `json:"alert_creation"`
	AlertGroupingParameters *AlertGroupingParameters `json:"alert_grouping_parameters,omitempty"`
}

type IncidentUrgencyRule struct {
	Type                string                   `json:"type"`
	Urgency             string                   `json:"urgency,omitempty"`
	DuringSupportHours  *SupportHoursUrgencyRule `json:"during_support_hours,omitempty"`
	OutsideSupportHours *SupportHoursUrgencyRule `json:"outside_support_hours,omitempty"`
}

type SupportHoursUrgencyRule struct {
	Type    string `json:"type"`
	Urgency string `json:"urgency"`
}

type SupportHours struct {
	Type       string `json:"type"`
	TimeZone   string `json:"time_zone"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	DaysOfWeek []uint `json:"days_of_week"`
}

type ScheduledAction struct {
	Type      string            `json:"type"`
	At        map[string]string `json:"at"`
	ToUrgency string            `json:"to_urgency"`
}

type AlertGroupingParameters struct {
	Type   string                 `json:"type"`
	Config map[string]interface{} `json:"config"`
}
