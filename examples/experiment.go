package main

import (
	"context"
	"fmt"
	"time"

	"github.com/inloco/go-wasabi"
	"github.com/inloco/go-wasabi/experiments"
)

/*
	This file create an experiment with two buckets: control (20%) and treatment (80%)
	and start the experiment
*/
func main() {
	client := wasabi.NewHttpClient("URL", "application_name", "username", "password")

	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)
	res, err := client.CreateExperiment(context.Background(), &experiments.Experiment{
		ApplicationName: "develop",
		Label:           "nome",
		SamplingPercent: 1,
		StartTime:       &now,
		EndTime:         &tomorrow,
		Description:     "descrição de teste",
	})
	if err != nil {
		panic(err)
	}

	buckets := []*experiments.Bucket{
		&experiments.Bucket{
			AllocationPercent: 0.2,
			IsControl:         true,
			Label:             experiments.BucketLabelControl,
			Payload:           "do_not_deliver", // payload recognized by 4apps backend
			ExperimentID:      res.ID,
			State:             experiments.BucketStateOpen,
		},
		&experiments.Bucket{
			AllocationPercent: 0.8,
			Label:             experiments.BucketLabelTreatment,
			ExperimentID:      res.ID,
			State:             experiments.BucketStateOpen,
		},
	}

	_, err = client.CreateBucket(context.Background(), buckets[0])
	if err != nil {
		panic(err)
	}

	_, err = client.CreateBucket(context.Background(), buckets[1])
	if err != nil {
		panic(err)
	}

	// activate the experiment
	_, err = client.UpdateExperimentState(context.Background(), res.ID, experiments.ExperimentStateRunning)
	if err != nil {
		panic(err)
	}

	fmt.Println("Experimento criado com sucesso!")
}
