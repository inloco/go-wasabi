package http

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/inloco/go-wasabi/assignments"
	"github.com/inloco/go-wasabi/experiments"
)

type HttpClient struct {
	http.Client

	address string

	applicationName string

	login    string
	password string
}

func NewHttpClient(address, applicationName, login, password string) *HttpClient {
	return &HttpClient{
		address:         address,
		applicationName: applicationName,
		login:           login,
		password:        password,
	}
}

func (c *HttpClient) GenerateAssignment(ctx context.Context, experimentLabel string, userID string) (*assignments.Assignment, error) {

	url := c.address + generateAssignmentPath(c.applicationName, experimentLabel, userID)

	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	switch {
	case err != nil:
		return nil, err
	case res.StatusCode != http.StatusOK:
		return nil, c.handleUnexpectedResponse(res.StatusCode, body)
	}

	assignment := &assignments.Assignment{}
	err = json.Unmarshal(body, assignment)

	return assignment, err
}

func (c *HttpClient) CreateExperiment(ctx context.Context, experiment *experiments.Experiment) (*experiments.Experiment, error) {
	return nil, nil
}

func (c *HttpClient) GetExperiments(ctx context.Context) ([]*experiments.Experiment, error) {
	return nil, nil
}

func (c *HttpClient) GetExperimentByID(ctx context.Context, experimentID string) (*assignments.Assignment, error) {
	return nil, nil
}
