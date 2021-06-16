package syntheticsclient

import (
	"bytes"
	"encoding/json"
	"time"
)

type CreateBrowserCheck struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Type      string    `json:"type,omitempty"`
	Frequency int       `json:"frequency,omitempty"`
	Paused    bool      `json:"paused,omitempty"`
	Muted     bool      `json:"muted,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Links     struct {
		Self      string `json:"self,omitempty"`
		SelfHTML  string `json:"self_html,omitempty"`
		Metrics   string `json:"metrics,omitempty"`
		LastRun   string `json:"last_run,omitempty"`
		Runs      string `json:"runs,omitempty"`
		ShareLink string `json:"share_link,omitempty"`
	} `json:"links,omitempty"`
	Tags []struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"tags,omitempty"`
	Status struct {
		LastCode           int       `json:"last_code,omitempty"`
		LastMessage        string    `json:"last_message,omitempty"`
		LastResponseTime   int       `json:"last_response_time,omitempty"`
		LastRunAt          time.Time `json:"last_run_at,omitempty"`
		LastFailureAt      time.Time `json:"last_failure_at,omitempty"`
		LastAlertAt        time.Time `json:"last_alert_at,omitempty"`
		HasFailure         bool      `json:"has_failure,omitempty"`
		HasLocationFailure bool      `json:"has_location_failure,omitempty"`
	} `json:"status,omitempty"`
	RoundRobin      bool `json:"round_robin,omitempty"`
	AutoRetry       bool `json:"auto_retry,omitempty"`
	Enabled         bool `json:"enabled,omitempty"`
	BlackoutPeriods []struct {
		StartDate         string    `json:"start_date,omitempty"`
		EndDate           string    `json:"end_date,omitempty"`
		Timezone          string    `json:"timezone,omitempty"`
		StartTime         time.Time `json:"start_time,omitempty"`
		EndTime           time.Time `json:"end_time,omitempty"`
		RepeatType        string    `json:"repeat_type,omitempty"`
		DurationInMinutes int       `json:"duration_in_minutes,omitempty"`
		IsRepeat          bool      `json:"is_repeat,omitempty"`
		MonthlyRepeatType string    `json:"monthly_repeat_type,omitempty"`
		CreatedAt         time.Time `json:"created_at,omitempty"`
		UpdatedAt         time.Time `json:"updated_at,omitempty"`
	} `json:"blackout_periods,omitempty"`
	Locations []struct {
		ID          int    `json:"id,omitempty"`
		Name        string `json:"name,omitempty"`
		WorldRegion string `json:"world_region,omitempty"`
		RegionCode  string `json:"region_code,omitempty"`
	} `json:"locations,omitempty"`
	Integrations []struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"integrations,omitempty"`
	HTTPRequestHeaders struct {
		ContentType string `json:"Content-Type,omitempty"`
	} `json:"http_request_headers,omitempty"`
	Notifications struct {
		Sms                     bool `json:"sms,omitempty"`
		Call                    bool `json:"call,omitempty"`
		Email                   bool `json:"email,omitempty"`
		NotifyAfterFailureCount int  `json:"notify_after_failure_count,omitempty"`
		NotifyOnLocationFailure bool `json:"notify_on_location_failure,omitempty"`
		Muted                   bool `json:"muted,omitempty"`
		NotifyWho               []struct {
			Sms             bool   `json:"sms,omitempty"`
			Call            bool   `json:"call,omitempty"`
			Email           bool   `json:"email,omitempty"`
			CustomUserEmail string `json:"custom_user_email,omitempty"`
			Type            string `json:"type,omitempty"`
			Links           struct {
				SelfHTML string `json:"self_html,omitempty"`
			} `json:"links,omitempty"`
			ID int `json:"id,omitempty"`
		} `json:"notify_who,omitempty"`
		NotificationWindows []struct {
			StartTime         time.Time `json:"start_time,omitempty"`
			EndTime           time.Time `json:"end_time,omitempty"`
			DurationInMinutes int       `json:"duration_in_minutes,omitempty"`
			TimeZone          string    `json:"time_zone,omitempty"`
		} `json:"notification_windows,omitempty"`
		Escalations []struct {
			Sms          bool `json:"sms,omitempty"`
			Call         bool `json:"call,omitempty"`
			Email        bool `json:"email,omitempty"`
			AfterMinutes int  `json:"after_minutes,omitempty"`
			NotifyWho    []struct {
				ID    int    `json:"id,omitempty"`
				Type  string `json:"type,omitempty"`
				Links struct {
					SelfHTML string `json:"self_html,omitempty"`
				} `json:"links,omitempty"`
				CustomUserEmail string `json:"custom_user_email,omitempty"`
			} `json:"notify_who,omitempty"`
			NotificationWindow []struct {
				StartTime         time.Time `json:"start_time,omitempty"`
				EndTime           time.Time `json:"end_time,omitempty"`
				DurationInMinutes int       `json:"duration_in_minutes,omitempty"`
				TimeZone          string    `json:"time_zone,omitempty"`
			} `json:"notification_window,omitempty"`
			IsRepeat bool `json:"is_repeat,omitempty"`
		} `json:"escalations,omitempty"`
	} `json:"notifications,omitempty"`
	URL                 string `json:"url,omitempty"`
	UserAgent           string `json:"user_agent,omitempty"`
	AutoUpdateUserAgent bool   `json:"auto_update_user_agent,omitempty"`
	Browser             struct {
		Label string `json:"label,omitempty"`
		Code  string `json:"code,omitempty"`
	} `json:"browser,omitempty"`
	Steps []struct {
		ItemMethod   string    `json:"item_method,omitempty"`
		Value        string    `json:"value,omitempty"`
		How          string    `json:"how,omitempty"`
		What         string    `json:"what,omitempty"`
		UpdatedAt    time.Time `json:"updated_at,omitempty"`
		CreatedAt    time.Time `json:"created_at,omitempty"`
		VariableName string    `json:"variable_name,omitempty"`
		Name         string    `json:"name,omitempty"`
		Position     int       `json:"position,omitempty"`
	} `json:"steps,omitempty"`
	Cookies []struct {
		Key    string `json:"key,omitempty"`
		Value  string `json:"value,omitempty"`
		Domain string `json:"domain,omitempty"`
		Path   string `json:"path,omitempty"`
	} `json:"cookies,omitempty"`
	JavascriptFiles []struct {
		ID        int       `json:"id,omitempty"`
		Name      string    `json:"name,omitempty"`
		CreatedAt time.Time `json:"created_at,omitempty"`
		UpdatedAt time.Time `json:"updated_at,omitempty"`
		Links     struct {
			Self string `json:"self,omitempty"`
		} `json:"links,omitempty"`
	} `json:"javascript_files,omitempty"`
	ExcludedFiles []struct {
		ExclusionType string    `json:"exclusion_type,omitempty"`
		PresetName    string    `json:"preset_name,omitempty"`
		URL           string    `json:"url,omitempty"`
		CreatedAt     time.Time `json:"created_at,omitempty"`
		UpdatedAt     time.Time `json:"updated_at,omitempty"`
	} `json:"excluded_files,omitempty"`
	Viewport struct {
		Height int `json:"height,omitempty"`
		Width  int `json:"width,omitempty"`
	} `json:"viewport,omitempty"`
	EnforceSslValidation bool `json:"enforce_ssl_validation,omitempty"`
	ThresholdMonitors    []struct {
		Matcher        string    `json:"matcher,omitempty"`
		MetricName     string    `json:"metric_name,omitempty"`
		ComparisonType string    `json:"comparison_type,omitempty"`
		Value          int       `json:"value,omitempty"`
		CreatedAt      time.Time `json:"created_at,omitempty"`
		UpdatedAt      time.Time `json:"updated_at,omitempty"`
	} `json:"threshold_monitors,omitempty"`
	DNSOverrides struct {
		OriginalDomainCom string `json:"original.domain.com,omitempty"`
		OriginalHostCom   string `json:"original.host.com,omitempty"`
	} `json:"dns_overrides,omitempty"`
	Connection struct {
		UploadBandwidth   int     `json:"upload_bandwidth,omitempty"`
		DownloadBandwidth int     `json:"download_bandwidth,omitempty"`
		Latency           int     `json:"latency,omitempty"`
		PacketLoss        float32 `json:"packet_loss,omitempty"`
	} `json:"connection,omitempty"`
	WaitForFullMetrics bool `json:"wait_for_full_metrics,omitempty"`
}

func parseCreateBrowserCheckResponse(response string) (*CreateBrowserCheck, error) {

	var createBrowserCheck CreateBrowserCheck
	JSONResponse := []byte(response)
	err := json.Unmarshal(JSONResponse, &createBrowserCheck)
	if err != nil {
		return nil, err
	}

	return &createBrowserCheck, err
}

func (c Client) CreateBrowserCheck(browserCheckDetails *CreateBrowserCheck) (*CreateBrowserCheck, *RequestDetails, error) {

	body, err := json.Marshal(browserCheckDetails)
	if err != nil {
		return nil, nil, err
	}

	details, err := c.makePublicAPICall("POST", "/v2/checks/real_browsers", bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, details, err
	}

	newBrowserCheck, err := parseCreateBrowserCheckResponse(details.ResponseBody)
	if err != nil {
		return newBrowserCheck, details, err
	}

	return newBrowserCheck, details, nil
}
