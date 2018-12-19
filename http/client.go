package http

import (
	"context"

	"github.com/inloco/go-wasabi/assignments"
	"github.com/inloco/go-wasabi/experiments"
)

type HttpClient struct {
	address string

	login    string
	password string
}

func NewHttpClient(address, login, password string) *HttpClient {
	return &HttpClient{
		address:  address,
		login:    login,
		password: password,
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
