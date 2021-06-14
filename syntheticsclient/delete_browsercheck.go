package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type DeleteBrowserCheck struct {
	Result  string        `json:"result"`
	Message string        `json:"message"`
	Errors  []interface{} `json:"errors"`
}

func parseDeleteBrowserCheckResponse(response string) (*DeleteBrowserCheck, error) {
	var deleteBrowserCheck DeleteBrowserCheck
	err := json.Unmarshal([]byte(response), &deleteBrowserCheck)
	if err != nil {
		return nil, err
	}

	return &deleteBrowserCheck, err
}

func (c Client) DeleteBrowserCheck(id int) (*DeleteBrowserCheck, error) {
	requestDetails, err := c.makePublicAPICall("DELETE", fmt.Sprintf("/v2/checks/real_browsers/%d", id), bytes.NewBufferString("{}"), nil)
	if err != nil {
		return nil, err
	}

	deleteBrowserCheck, err := parseDeleteBrowserCheckResponse(requestDetails.ResponseBody)
	if err != nil {
		return deleteBrowserCheck, err
	}

	return deleteBrowserCheck, nil
}
