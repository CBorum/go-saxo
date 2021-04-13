package referencedata

type GetTradingScheduleResponse struct {
	Sessions []struct {
		EndTime   string `json:"EndTime"`
		StartTime string `json:"StartTime"`
		State     string `json:"State"`
	} `json:"Sessions"`
	TimeZone             float64 `json:"TimeZone"`
	TimeZoneAbbreviation string  `json:"TimeZoneAbbreviation"`
	TimeZoneOffset       string  `json:"TimeZoneOffset"`
}
