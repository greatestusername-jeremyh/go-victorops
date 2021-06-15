package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type GetHttpCheck struct {
	ID        int
	Name      string `json:"name"`
	Type      string
	Frequency int       `json:"frequency"`
	Paused    bool      `json:"paused"`
	Muted     bool      `json:"muted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Links     struct {
		Self      string `json:"self"`
		SelfHTML  string `json:"self_html"`
		Metrics   string `json:"metrics"`
		LastRun   string `json:"last_run"`
		Runs      string `json:"runs"`
		ShareLink string `json:"share_link"`
	} `json:"links"`
	Status struct {
		LastCode           int       `json:"last_code"`
		LastMessage        string    `json:"last_message"`
		LastResponseTime   int       `json:"last_response_time"`
		LastRunAt          time.Time `json:"last_run_at"`
		LastFailureAt      time.Time `json:"last_failure_at"`
		LastAlertAt        time.Time `json:"last_alert_at"`
		HasFailure         bool      `json:"has_failure"`
		HasLocationFailure bool      `json:"has_location_failure"`
	} `json:"status"`
	Notifications struct {
		Sms                     bool
		Call                    bool
		Email                   bool
		NotifyAfterFailureCount int  `json:"notify_after_failure_count"`
		NotifyOnLocationFailure bool `json:"notify_on_location_failure"`
		Muted                   bool `json:"muted"`
		NotifyWho               []struct {
			Sms             bool
			Call            bool
			Email           bool
			CustomUserEmail bool `json:"custom_user_email"`
			Type            string
			Links           struct {
				SelfHTML string `json:"self_html"`
			} `json:"links"`
			ID int
		} `json:"notify_who"`
		NotificationWindows []interface{} `json:"notification_windows"`
		Escalations         []interface{} `json:"escalations"`
	} `json:"notifications"`
	ResponseTimeMonitorMilliseconds int `json:"response_time_monitor_milliseconds"`
	HTTPRequestHeaders              struct {
		UserAgent string `json:"User-Agent"`
	} `json:"http_request_headers"`
	RoundRobin   bool `json:"round_robin"`
	AutoRetry    bool `json:"auto_retry"`
	Enabled      bool `json:"enabled"`
	Integrations []struct {
		ID   int
		Name string
	} `json:"integrations"`
	URL                 string `json:"url"`
	UserAgent           string `json:"user_agent"`
	AutoUpdateUserAgent bool   `json:"auto_update_user_agent"`
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
	WaitForFullMetrics bool `json:"wait_for_full_metrics"`
	Tags               []struct {
		ID   int
		Name string
	} `json:"tags"`
	BlackoutPeriods []interface{} `json:"blackout_periods"`
	Locations       []struct {
		ID          int
		WorldRegion string `json:"world_region"`
		RegionCode  string `json:"region_code"`
		Name        string
	} `json:"locations"`
	Steps             []interface{} `json:"steps"`
	JavascriptFiles   []interface{} `json:"javascript_files"`
	ThresholdMonitors []interface{} `json:"threshold_monitors"`
	ExcludedFiles     []interface{} `json:"excluded_files"`
	Cookies           []interface{} `json:"cookies"`
	Connection        struct {
		DownloadBandwidth int     `json:"download_bandwidth"`
		UploadBandwidth   int     `json:"upload_bandwidth"`
		Latency           int     `json:"latency"`
		PacketLoss        float32 `json:"packet_loss"`
	} `json:"connection"`
}

func parseGetHttpCheckResponse(response string) (*GetHttpCheck, error) {
	// Parse the response and return the user object
	var httpcheck GetHttpCheck
	err := json.Unmarshal([]byte(response), &httpcheck)
	if err != nil {
		return nil, err
	}

	return &httpcheck, err
}

func (c Client) GetHttpCheck(id int) (*GetHttpCheck, *RequestDetails, error) {
	details, err := c.makePublicAPICall("GET", fmt.Sprintf("/v2/checks/http/%d", id), bytes.NewBufferString("{}"), nil)

	// Check for errors
	if err != nil {
		return nil, details, err
	}

	httpcheck, err := parseGetHttpCheckResponse(details.ResponseBody)
	if err != nil {
		return httpcheck, details, err
	}

	return httpcheck, details, nil
}
