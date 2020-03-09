package fixtures

import (
	"time"

	"github.com/inloco/go-wasabi/experiments"
)

const (
	startDate     = "2020-03-04"
	layoutDateISO = "2006-01-02"
)

func Experiment() *experiments.Experiment {
	now, _ := time.Parse(layoutDateISO, "2020-03-04")
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

func ExperimentToUpdateState() *experiments.Experiment {
	experiment := Experiment()
	experiment.ID = "experiment_id"
	experiment.State = experiments.ExperimentStateDraft

	return experiment
}

func ExperimentToUpdate() *experiments.Experiment {
	experiment := Experiment()
	experiment.ID = "experiment_update_id"

	return experiment
}
