package referencedata

type GetTimeZonesResponse struct {
	Data []struct {
		DisplayName    string `json:"DisplayName"`
		TimeZoneID     string `json:"TimeZoneId"`
		TimeZoneOffset string `json:"TimeZoneOffset"`
		ZoneName       string `json:"ZoneName"`
	} `json:"Data"`
}
