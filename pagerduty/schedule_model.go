package pagerduty

type Schedule struct {
	ApiObject
	Description          string           `json:"description,omitempty"`
	EscalationPolicies   []*ApiObject     `json:"escalation_policies,omitempty"`
	FinalSchedule        *SubSchedule     `json:"final_schedule,omitempty"`
	Name                 string           `json:"name,omitempty"`
	OverridesSubSchedule *SubSchedule     `json:"overrides_subschedule,omitempty"`
	Schedule             *Schedule        `json:"schedule,omitempty"`
	ScheduleLayers       []*ScheduleLayer `json:"schedule_layers,omitempty"`
	TimeZone             string           `json:"time_zone,omitempty"`
	Users                []*ApiObject     `json:"users,omitempty"`
	Teams                []*ApiObject     `json:"teams,omitempty"`
}

type SubSchedule struct {
	Name                       string                `json:"name,omitempty"`
	RenderedCoveragePercentage float64               `json:"rendered_coverage_percentage,omitempty"`
	RenderedScheduleEntries    []*ScheduleLayerEntry `json:"rendered_schedule_entries,omitempty"`
}

type ScheduleLayerEntry struct {
	End   string     `json:"end,omitempty"`
	Start string     `json:"start,omitempty"`
	User  *ApiObject `json:"user,omitempty"`
}

type ScheduleLayer struct {
	End                        string                  `json:"end,omitempty"`
	Id                         string                  `json:"id,omitempty"`
	Name                       string                  `json:"name,omitempty"`
	RenderedCoveragePercentage float64                 `json:"rendered_coverage_percentage,omitempty"`
	RenderedScheduleEntries    []*ScheduleLayerEntry   `json:"rendered_schedule_entries,omitempty"`
	Restrictions               []*Restriction          `json:"restrictions,omitempty"`
	RotationTurnLengthSeconds  int                     `json:"rotation_turn_length_seconds,omitempty"`
	RotationVirtualStart       string                  `json:"rotation_virtual_start,omitempty"`
	Start                      string                  `json:"start,omitempty"`
	Users                      []*UserReferenceWrapper `json:"users,omitempty"`
}

type Restriction struct {
	DurationSeconds uint   `json:"duration_seconds,omitempty"`
	StartDayOfWeek  uint   `json:"start_day_of_week,omitempty"`
	StartTimeOfDay  string `json:"start_time_of_day,omitempty"`
	Type            string `json:"type,omitempty"`
}

type UserReferenceWrapper struct {
	User *UserReference `json:"user,omitempty"`
}

type UserReference ApiObject
