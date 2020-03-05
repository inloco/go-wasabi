package wasabi

import (
	"context"

	"github.com/inloco/go-wasabi/assignments"
	"github.com/inloco/go-wasabi/experiments"
	"github.com/stretchr/testify/mock"
)

func NewMockedClient() (*MockedClient, error) {
	return &MockedClient{}, nil
}

type MockedClient struct {
	mock.Mock
}

func (m *MockedClient) GenerateAssignment(ctx context.Context, experimentLabel string, userID string) (*assignments.Assignment, error) {
	args := m.Called(experimentLabel, userID)
	return args.Get(0).(*assignments.Assignment), args.Error(1)
}

func (m *MockedClient) CreateExperiment(ctx context.Context, experiment *experiments.Experiment) (*experiments.Experiment, error) {
	args := m.Called(experiment)
	return args.Get(0).(*experiments.Experiment), args.Error(1)
}

func (m *MockedClient) UpdateExperimentState(ctx context.Context, id string, state experiments.ExperimentState) (*experiments.Experiment, error) {
	args := m.Called(id, state)
	return args.Get(0).(*experiments.Experiment), args.Error(1)
}

func (m *MockedClient) CreateBucket(ctx context.Context, bucket *experiments.Bucket) (*experiments.Bucket, error) {
	args := m.Called(bucket)
	return args.Get(0).(*experiments.Bucket), args.Error(1)
}

func (m *MockedClient) GetExperimentByID(ctx context.Context, experimentID string) (*experiments.Experiment, error) {
	args := m.Called(experimentID)
	return args.Get(0).(*experiments.Experiment), args.Error(1)
}

func (m *MockedClient) GetExperimentBuckets(ctx context.Context, experimentID string) ([]*experiments.Bucket, error) {
	args := m.Called(experimentID)
	return args.Get(0).([]*experiments.Bucket), args.Error(1)
}
