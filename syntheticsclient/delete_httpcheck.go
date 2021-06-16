package syntheticsclient

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
