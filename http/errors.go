package http

import (
	"encoding/json"
	"fmt"
)

type UnexpectedResponse struct {
	statusCode int
	message    string
}

func (e *UnexpectedResponse) Error() string {
	return fmt.Sprintf("wasabi responded unexpectedly with code %d: %s", e.statusCode, e.message)
}

type remoteErrorModel struct {
	Error remoteErrorInfo `json:"error"`
}

type remoteErrorInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *HttpClient) handleUnexpectedResponse(statusCode int, body []byte) error {

	err := &UnexpectedResponse{statusCode: statusCode}

	remoteErrorModel := &remoteErrorModel{}
	if marshalingErr := json.Unmarshal(body, remoteErrorModel); marshalingErr != nil {
		return err
	}

	err.message = remoteErrorModel.Error.Message
	return err
}
