package wasabi

import (
	"context"

	"github.com/inloco/go-wasabi/assignments"
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
