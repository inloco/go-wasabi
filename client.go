package wasabi

import (
	"context"

	"github.com/inloco/go-wasabi/assignments"
	"github.com/inloco/go-wasabi/experiments"
)

type Client interface {
	GenerateAssignment(ctx context.Context, experimentName string, userID string) (*assignments.Assignment, error)

	CreateExperiment(ctx context.Context, experiment *experiments.Experiment) (*experiments.Experiment, error)
	GetExperiments(ctx context.Context) ([]*experiments.Experiment, error)
	GetExperimentByID(ctx context.Context, experimentID string) (*assignments.Assignment, error)
}
