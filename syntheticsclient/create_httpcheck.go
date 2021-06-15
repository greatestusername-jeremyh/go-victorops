package main

import (
	"bytes"
	"encoding/json"
	"time"
)

type CreateHttpCheck struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Type      string    `json:"type,omitempty"`
	Frequency int       `json:"frequency,omitempty"`
	Paused    bool      `json:"paused,omitempty"`
	Muted     bool      `json:"muted,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Links     struct {
		Self     string `json:"self,omitempty"`
		SelfHTML string `json:"self_html,omitempty"`
		Metrics  string `json:"metrics,omitempty"`
		LastRun  string `json:"last_run,omitempty"`
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
		UserAgent string `json:"User-Agent,omitempty"`
	} `json:"http_request_headers,omitempty"`
	Notifications struct {
		Sms       bool `json:"sms,omitempty"`
		Email     bool `json:"email,omitempty"`
		Call      bool `json:"call,omitempty"`
		NotifyWho []struct {
			Sms   bool `json:"sms,omitempty"`
			Email bool `json:"email,omitempty"`
			Call  bool `json:"call,omitempty"`
			Links struct {
			} `json:"links,omitempty"`
		} `json:"notify_who,omitempty"`
		NotifyAfterFailureCount int  `json:"notify_after_failure_count,omitempty"`
		NotifyOnLocationFailure bool `json:"notify_on_location_failure,omitempty"`
		NotificationWindows     []struct {
		} `json:"notification_windows,omitempty"`
		Escalations []struct {
			Sms          bool `json:"sms,omitempty"`
			Email        bool `json:"email,omitempty"`
			Call         bool `json:"call,omitempty"`
			AfterMinutes int  `json:"after_minutes,omitempty"`
			NotifyWho    []struct {
				Links struct {
				} `json:"links,omitempty"`
			} `json:"notify_who,omitempty"`
			IsRepeat bool `json:"is_repeat,omitempty"`
		} `json:"escalations,omitempty"`
		Muted bool `json:"muted,omitempty"`
	} `json:"notifications,omitempty"`
	URL             string `json:"url,omitempty"`
	HTTPMethod      string `json:"http_method,omitempty"`
	SuccessCriteria []struct {
		ActionType       string    `json:"action_type,omitempty"`
		ComparisonString string    `json:"comparison_string,omitempty"`
		CreatedAt        time.Time `json:"created_at,omitempty"`
		UpdatedAt        time.Time `json:"updated_at,omitempty"`
	} `json:"success_criteria,omitempty"`
	Connection struct {
		UploadBandwidth   int     `json:"upload_bandwidth,omitempty"`
		DownloadBandwidth int     `json:"download_bandwidth,omitempty"`
		Latency           int     `json:"latency,omitempty"`
		PacketLoss        float32 `json:"packet_loss,omitempty"`
	} `json:"connection,omitempty"`
}

func parseCreateHttpCheckResponse(response string) (*CreateHttpCheck, error) {

	var createHttpCheck CreateHttpCheck
	JSONResponse := []byte(response)
	err := json.Unmarshal(JSONResponse, &createHttpCheck)
	if err != nil {
		return nil, err
	}

	return &createHttpCheck, err
}

// CreateTeam creates a team in the victorops organization
func (c Client) CreateHttpCheck(httpCheckDetails *CreateHttpCheck) (*CreateHttpCheck, *RequestDetails, error) {

	body, err := json.Marshal(httpCheckDetails)
	if err != nil {
		return nil, nil, err
	}

	// Make the request
	details, err := c.makePublicAPICall("POST", "/v2/checks/http", bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, details, err
	}

	newHttpCheck, err := parseCreateHttpCheckResponse(details.ResponseBody)
	if err != nil {
		return newHttpCheck, details, err
	}

	return newHttpCheck, details, nil
}
