package syntheticsclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type GetChecks struct {
	CurrentPage  int         `json:"current_page"`
	PerPage      int         `json:"per_page"`
	NextPage     interface{} `json:"next_page"`
	PreviousPage interface{} `json:"previous_page"`
	TotalPages   int         `json:"total_pages"`
	TotalCount   int         `json:"total_count"`
	Checks       []struct {
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
		Status struct {
			LastCode           int         `json:"last_code"`
			LastMessage        interface{} `json:"last_message"`
			LastResponseTime   int         `json:"last_response_time"`
			LastRunAt          time.Time   `json:"last_run_at"`
			LastFailureAt      time.Time   `json:"last_failure_at"`
			LastAlertAt        time.Time   `json:"last_alert_at"`
			HasFailure         bool        `json:"has_failure"`
			HasLocationFailure interface{} `json:"has_location_failure"`
		} `json:"status"`
		Tags []interface{} `json:"tags"`
	} `json:"checks"`
}

// Leaving off "Enabled" filter setting. Can be added later if required.
type GetChecksOptions struct {
	Type    string `json:"type"`
	PerPage int    `json:"per_page"`
	Page    int    `json:"page"`
	Muted   bool   `json:"muted"`
}

func parseChecksResponse(response string) (*GetChecks, error) {
	// Parse the response and return the check object
	var checks GetChecks
	err := json.Unmarshal([]byte(response), &checks)
	if err != nil {
		return nil, err
	}

	return &checks, err
}

// GetChecks returns all checks
func (c Client) GetChecks(params *GetChecksOptions) (*GetChecks, *RequestDetails, error) {
	// Check for default params
	if params.Type == "" {
		params.Type = "all"
	}
	if params.Page == 0 {
		params.Page = int(1)
	}
	if params.PerPage == 0 {
		params.PerPage = int(50)
	}

	// Make the request
	details, err := c.makePublicAPICall(
		"GET",
		fmt.Sprintf("/v2/checks?type=%s&page=%d&per_page=%d&muted=%t", params.Type, params.Page, params.PerPage, params.Muted),
		bytes.NewBufferString("{}"),
		nil)

	// Check for errors
	if err != nil {
		return nil, details, err
	}

	check, err := parseChecksResponse(details.ResponseBody)
	if err != nil {
		return check, details, err
	}

	return check, details, nil
}
