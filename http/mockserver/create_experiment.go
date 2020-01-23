package mockserver

import (
	"encoding/json"
	"net/http"

	"github.com/inloco/go-wasabi/experiments/fixtures"
	uuid "github.com/satori/go.uuid"
)

const (
	createExperimentPath           = "/api/v1/experiments"
	createExperimentResponseStatus = http.StatusOK
)

func handleCreateExperiment(w http.ResponseWriter, req *http.Request) {
	experiment := fixtures.Experiment()
	experiment.ID = uuid.NewV4().String()

	payload, _ := json.Marshal(experiment)

	w.Write([]byte(payload))
	w.WriteHeader(createExperimentResponseStatus)
}
