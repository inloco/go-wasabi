package mockserver

import (
	"encoding/json"
	"net/http"

	"github.com/inloco/go-wasabi/experiments"

	"github.com/inloco/go-wasabi/experiments/fixtures"
)

const (
	updateExperimentStatePath           = "/api/v1/experiments/experiment_id"
	updateExperimentStateResponseStatus = http.StatusOK
	updateExperimentPath                = "/api/v1/experiments/experiment_update_id"
	updateExperimentResponseStatus      = http.StatusOK
)

func handleUpdateExperimentState(w http.ResponseWriter, req *http.Request) {
	experiment := fixtures.ExperimentToUpdateState()
	experiment.State = experiments.ExperimentStateRunning

	payload, _ := json.Marshal(experiment)

	w.Write([]byte(payload))
	w.WriteHeader(updateExperimentStateResponseStatus)
}

func handleUpdateExperiment(w http.ResponseWriter, req *http.Request) {
	experiment := fixtures.ExperimentToUpdate()

	payload, _ := json.Marshal(experiment)

	w.Write([]byte(payload))
	w.WriteHeader(updateExperimentResponseStatus)
}
