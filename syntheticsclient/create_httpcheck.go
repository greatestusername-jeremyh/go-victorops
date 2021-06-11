package main

import (
	"bytes"
	"encoding/json"
	"time"
)

type CreateHttpCheck struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Frequency int       `json:"frequency"`
	Paused    bool      `json:"paused"`
	Muted     bool      `json:"muted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Links     struct {
		Self     string `json:"self"`
		SelfHTML string `json:"self_html"`
		Metrics  string `json:"metrics"`
		LastRun  string `json:"last_run"`
	} `json:"links"`
	Tags []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"tags"`
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
	RoundRobin      bool `json:"round_robin"`
	AutoRetry       bool `json:"auto_retry"`
	Enabled         bool `json:"enabled"`
	BlackoutPeriods []struct {
	} `json:"blackout_periods"`
	Locations []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		WorldRegion string `json:"world_region"`
		RegionCode  string `json:"region_code"`
	} `json:"locations"`
	Integrations []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"integrations"`
	HTTPRequestHeaders struct {
		UserAgent string `json:"User-Agent"`
	} `json:"http_request_headers"`
	Notifications struct {
		Sms       bool `json:"sms"`
		Email     bool `json:"email"`
		Call      bool `json:"call"`
		NotifyWho []struct {
			Sms   bool `json:"sms"`
			Email bool `json:"email"`
			Call  bool `json:"call"`
			Links struct {
			} `json:"links"`
		} `json:"notify_who"`
		NotifyAfterFailureCount int  `json:"notify_after_failure_count"`
		NotifyOnLocationFailure bool `json:"notify_on_location_failure"`
		NotificationWindows     []struct {
		} `json:"notification_windows"`
		Escalations []struct {
			Sms          bool `json:"sms"`
			Email        bool `json:"email"`
			Call         bool `json:"call"`
			AfterMinutes int  `json:"after_minutes"`
			NotifyWho    []struct {
				Links struct {
				} `json:"links"`
			} `json:"notify_who"`
			IsRepeat bool `json:"is_repeat"`
		} `json:"escalations"`
		Muted bool `json:"muted"`
	} `json:"notifications"`
	URL             string `json:"url"`
	HTTPMethod      string `json:"http_method"`
	SuccessCriteria []struct {
		ActionType       string    `json:"action_type"`
		ComparisonString string    `json:"comparison_string"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
	} `json:"success_criteria"`
	Connection struct {
		UploadBandwidth   int `json:"upload_bandwidth"`
		DownloadBandwidth int `json:"download_bandwidth"`
		Latency           int `json:"latency"`
		PacketLoss        float32 `json:"packet_loss"`
	} `json:"connection"`
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
func (c Client) CreateHttpCheck(httpCheckDetails string) (*CreateHttpCheck, *RequestDetails, error) {
	
	body := bytes.NewBufferString(httpCheckDetails)

	// Make the request
	details, err := c.makePublicAPICall("POST", "/v2/checks/http", body, nil)
	if err != nil {
		return nil, details, err
	}

	newHttpCheck, err := parseCreateHttpCheckResponse(details.ResponseBody)
	if err != nil {
		return newHttpCheck, details, err
	}

	return newHttpCheck, details, nil
}