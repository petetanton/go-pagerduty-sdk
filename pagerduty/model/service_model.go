package model

type Service struct {
	ApiObject
	Name                    string                   `json:"name,omitempty"`
	Description             string                   `json:"description,omitempty"`
	AutoResolveTimeout      uint                     `json:"auto_resolve_timeout"`
	AcknowledgementTimeout  uint                     `json:"acknowledgement_timeout"`
	Status                  string                   `json:"status,omitempty"`
	EscalationPolicy        *ApiObject               `json:"escalation_policy,omitempty"`
	IncidentUrgencyRule     *IncidentUrgencyRule     `json:"incident_urgency_rule,omitempty"`
	SupportHours            *SupportHours            `json:"support_hours,omitempty"`
	ScheduledActions        []*ScheduledAction       `json:"scheduled_actions,omitempty"`
	Addons                  []*ApiObject             `json:"addons,omitempty"`
	AlertCreation           string                   `json:"alert_creation,omitempty"`
	AlertGrouping           string                   `json:"alert_grouping"`
	AlertGroupingTimeout    uint                     `json:"alert_grouping_timeout,omitempty"`
	AlertGroupingParameters *AlertGroupingParameters `json:"alert_grouping_parameters,omitempty"`
	CreatedAt               string                   `json:"created_at,omitempty"`
	Integrations            []*ApiObject             `json:"integrations,omitempty"`
	LastIncidentTimestamp   string                   `json:"last_incident_timestamp,omitempty"`
	Teams                   []*ApiObject             `json:"teams,omitempty"`
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
	Type   string               `json:"type,omitempty"`
	Config *AlertGroupingConfig `json:"config,omitempty"`
}

type AlertGroupingConfig struct {
	Timeout   int      `json:"timeout,omitempty"`
	Aggregate string   `json:"aggregate,omitempty"`
	Fields    []string `json:"fields,omitempty"`
}
