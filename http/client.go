package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/inloco/go-wasabi/assignments"
	"github.com/inloco/go-wasabi/experiments"
)

const timeFormat = "2006-01-02T15:04:05"

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

	userIDEscaped := url.PathEscape(userID)
	urlPath := c.address + generateAssignmentPath(c.applicationName, experimentLabel, userIDEscaped)

	req, err := http.NewRequest(
		http.MethodGet,
		urlPath,
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	assignment := &assignments.Assignment{}
	err = json.Unmarshal(body, assignment)

	return assignment, err
}

func (c *HttpClient) CreateExperiment(ctx context.Context, experiment *experiments.Experiment) (*experiments.Experiment, error) {
	url := c.address + createExperimentPath()

	startTime, err := removeTimezoneFromTime(experiment.StartTime)
	if err != nil {
		return nil, err
	}

	endTime, err := removeTimezoneFromTime(experiment.EndTime)
	if err != nil {
		return nil, err
	}

	experiment.StartTime = startTime
	experiment.EndTime = endTime

	payload, err := json.Marshal(experiment)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.login, c.password)

	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	experimentCreated := &experiments.Experiment{}
	err = json.Unmarshal(body, experimentCreated)
	return experimentCreated, err
}

func (c *HttpClient) UpdateExperimentState(ctx context.Context, id string, state experiments.ExperimentState) (*experiments.Experiment, error) {
	url := c.address + updateExperimentPath(id)

	updateStateRequest := map[string]experiments.ExperimentState{"state": state}
	payload, err := json.Marshal(updateStateRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.login, c.password)

	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	experimentUpdated := &experiments.Experiment{}
	err = json.Unmarshal(body, experimentUpdated)
	return experimentUpdated, err
}

func (c *HttpClient) CreateBucket(ctx context.Context, bucket *experiments.Bucket) (*experiments.Bucket, error) {
	url := c.address + createBucketPath(bucket.ExperimentID)

	payload, err := json.Marshal(bucket)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.login, c.password)

	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	bucketCreated := &experiments.Bucket{}
	err = json.Unmarshal(body, bucketCreated)
	return bucketCreated, err
}

func (c *HttpClient) GetExperiments(ctx context.Context) ([]*experiments.Experiment, error) {
	return nil, nil
}

func (c *HttpClient) GetExperimentByID(ctx context.Context, experimentID string) (*experiments.Experiment, error) {
	url := c.address + getExperimentByIDPath(experimentID)

	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	experiment := &experiments.Experiment{}
	err = json.Unmarshal(body, experiment)

	return experiment, err
}

func (c *HttpClient) GetExperimentBuckets(ctx context.Context, experimentID string) ([]*experiments.Bucket, error) {
	url := c.address + getExperimentBucketsPath(experimentID)

	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	experimentWithOnlyBuckets := &experiments.Experiment{}
	err = json.Unmarshal(body, experimentWithOnlyBuckets)

	return experimentWithOnlyBuckets.Buckets, err
}

func (c *HttpClient) UpdateExperiment(ctx context.Context, id string, experiment *experiments.Experiment) (*experiments.Experiment, error) {
	url := c.address + updateExperimentPath(id)

	if experiment.StartTime != nil {
		startTime, err := removeTimezoneFromTime(experiment.StartTime)
		if err != nil {
			return nil, err
		}
		experiment.StartTime = startTime
	}

	if experiment.EndTime != nil {
		endTime, err := removeTimezoneFromTime(experiment.EndTime)
		if err != nil {
			return nil, err
		}
		experiment.EndTime = endTime
	}

	payload, err := json.Marshal(experiment)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.login, c.password)

	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	experimentUpdated := &experiments.Experiment{}
	err = json.Unmarshal(body, experimentUpdated)
	return experimentUpdated, err
}
