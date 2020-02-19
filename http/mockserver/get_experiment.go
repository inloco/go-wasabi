package mockserver

import (
	"encoding/json"
	"net/http"

	"github.com/inloco/go-wasabi/experiments/fixtures"
	uuid "github.com/satori/go.uuid"
)

const (
	getExperimentByIDPath           = "/api/v1/experiments/experiment01"
	getExperimentByIDResponseStatus = http.StatusOK
)

func handleGetExperimentByID(w http.ResponseWriter, req *http.Request) {
	experiment := fixtures.Experiment()
	experiment.ID = uuid.NewV4().String()

	payload, _ := json.Marshal(experiment)

	w.Write([]byte(payload))
	w.WriteHeader(getExperimentByIDResponseStatus)
}
