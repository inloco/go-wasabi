package http

import (
	"context"
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

func (c *HttpClient) GenerateAssignment(ctx context.Context, experimentName string, userID string) (*assignments.Assignment, error) {
	return nil, nil
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
