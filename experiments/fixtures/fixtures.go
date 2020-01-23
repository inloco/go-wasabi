package fixtures

import (
	"time"

	"github.com/inloco/go-wasabi/experiments"
)

func Experiment() *experiments.Experiment {
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)

	return &experiments.Experiment{
		ApplicationName: "production",
		Description:     "experiment of campaign id",
		Label:           "experiment name",
		StartTime:       &now,
		EndTime:         &tomorrow,
	}
}

func Buckets() []*experiments.Bucket {
	return []*experiments.Bucket{
		&experiments.Bucket{
			AllocationPercent: 0.2,
			IsControl:         true,
			Label:             experiments.BucketLabelControl,
			Payload:           "do_not_deliver",
			ExperimentID:      "experiment_id",
		},
		&experiments.Bucket{
			AllocationPercent: 0.8,
			Label:             experiments.BucketLabelTreatment,
			ExperimentID:      "experiment_id",
		},
	}
}
