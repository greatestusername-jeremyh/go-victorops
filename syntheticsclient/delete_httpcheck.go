package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type DeleteHttpCheck struct {
	Result  string        `json:"result"`
	Message string        `json:"message"`
	Errors  []interface{} `json:"errors"`
}

func parseDeleteHttpCheckResponse(response string) (*DeleteHttpCheck, error) {
	var deleteHttpCheck DeleteHttpCheck
	err := json.Unmarshal([]byte(response), &deleteHttpCheck)
	if err != nil {
		return nil, err
	}

	return &deleteHttpCheck, err
}


// DeleteContact deletes a contact
func (c Client) DeleteHttpCheck(id int) (*DeleteHttpCheck, error) {
	requestDetails, err := c.makePublicAPICall("DELETE", fmt.Sprintf("/v2/checks/http/%d", id), bytes.NewBufferString("{}"), nil)
	if err != nil {
		return nil, err
	}

	deleteHttpCheck, err := parseDeleteHttpCheckResponse(requestDetails.ResponseBody)
	if err != nil {
		return deleteHttpCheck, err
	}

	return deleteHttpCheck, nil
}


// // CreateContact creates a new contact for a user
// func (c Client) UpdateHttpCheck(id int, httpCheckDetails string) (*UpdateHttpCheck, *RequestDetails, error) {
	
// 	body := bytes.NewBufferString(httpCheckDetails)

// 	requestDetails, err := c.makePublicAPICall("PUT", fmt.Sprintf("/v2/checks/http/%d", id), body, nil)
// 	if err != nil {
// 		return nil, requestDetails, err
// 	}

// 	updateHttpCheck, err := parseUpdateHttpCheckResponse(requestDetails.ResponseBody)
// 	if err != nil {
// 		return updateHttpCheck, requestDetails, err
// 	}

// 	return updateHttpCheck, requestDetails, nil
// }