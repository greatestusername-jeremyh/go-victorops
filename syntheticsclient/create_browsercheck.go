package main

import (
	"bytes"
	"encoding/json"
	"time"
)

type CreateBrowserCheck struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Frequency int       `json:"frequency"`
	Paused    bool      `json:"paused"`
	Muted     bool      `json:"muted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Links     struct {
		Self      string      `json:"self"`
		SelfHTML  string      `json:"self_html"`
		Metrics   string      `json:"metrics"`
		LastRun   interface{} `json:"last_run"`
		Runs      string      `json:"runs"`
		ShareLink string      `json:"share_link"`
	} `json:"links"`
	Status struct {
		LastCode           interface{} `json:"last_code"`
		LastMessage        interface{} `json:"last_message"`
		LastResponseTime   interface{} `json:"last_response_time"`
		LastRunAt          interface{} `json:"last_run_at"`
		LastFailureAt      interface{} `json:"last_failure_at"`
		LastAlertAt        interface{} `json:"last_alert_at"`
		HasFailure         interface{} `json:"has_failure"`
		HasLocationFailure interface{} `json:"has_location_failure"`
	} `json:"status"`
	Notifications struct {
		Sms                     bool `json:"sms"`
		Call                    bool `json:"call"`
		Email                   bool `json:"email"`
		NotifyAfterFailureCount int  `json:"notify_after_failure_count"`
		NotifyOnLocationFailure bool `json:"notify_on_location_failure"`
		Muted                   bool `json:"muted"`
		NotifyWho               []struct {
			Sms             bool        `json:"sms"`
			Call            bool        `json:"call"`
			Email           bool        `json:"email"`
			CustomUserEmail interface{} `json:"custom_user_email"`
			Type            string      `json:"type"`
			Links           struct {
				SelfHTML string `json:"self_html"`
			} `json:"links"`
			ID int `json:"id"`
		} `json:"notify_who"`
		NotificationWindows []interface{} `json:"notification_windows"`
		Escalations         []interface{} `json:"escalations"`
	} `json:"notifications"`
	ResponseTimeMonitorMilliseconds interface{} `json:"response_time_monitor_milliseconds"`
	HTTPRequestHeaders              struct {
		UserAgent string `json:"User-Agent"`
	} `json:"http_request_headers"`
	RoundRobin          bool          `json:"round_robin"`
	AutoRetry           bool          `json:"auto_retry"`
	Enabled             bool          `json:"enabled"`
	Integrations        []interface{} `json:"integrations"`
	URL                 string        `json:"url"`
	UserAgent           string        `json:"user_agent"`
	AutoUpdateUserAgent bool          `json:"auto_update_user_agent"`
	Viewport            struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"viewport"`
	EnforceSslValidation bool `json:"enforce_ssl_validation"`
	Browser              struct {
		Label string `json:"label"`
		Code  string `json:"code"`
	} `json:"browser"`
	DNSOverrides struct {
	} `json:"dns_overrides"`
	WaitForFullMetrics bool          `json:"wait_for_full_metrics"`
	Tags               []interface{} `json:"tags"`
	BlackoutPeriods    []interface{} `json:"blackout_periods"`
	Locations          []struct {
		ID          int    `json:"id"`
		WorldRegion string `json:"world_region"`
		RegionCode  string `json:"region_code"`
		Name        string `json:"name"`
	} `json:"locations"`
	Steps             []interface{} `json:"steps"`
	JavascriptFiles   []interface{} `json:"javascript_files"`
	ThresholdMonitors []interface{} `json:"threshold_monitors"`
	ExcludedFiles     []interface{} `json:"excluded_files"`
	Cookies           []interface{} `json:"cookies"`
	Connection        struct {
		DownloadBandwidth int `json:"download_bandwidth"`
		UploadBandwidth   int `json:"upload_bandwidth"`
		Latency           int `json:"latency"`
		PacketLoss        float32 `json:"packet_loss"`
	} `json:"connection"`
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

// CreateTeam creates a team in the victorops organization
func (c Client) CreateBrowserCheck(browserCheckDetails string) (*CreateBrowserCheck, *RequestDetails, error) {
	
	body := bytes.NewBufferString(browserCheckDetails)

	// Make the request
	details, err := c.makePublicAPICall("POST", "/v2/checks/real_browsers", body, nil)
	if err != nil {
		return nil, details, err
	}

	newBrowserCheck, err := parseCreateBrowserCheckResponse(details.ResponseBody)
	if err != nil {
		return newBrowserCheck, details, err
	}

	return newBrowserCheck, details, nil
}