package mockserver

import (
	"encoding/json"
	"net/http"

	"github.com/inloco/go-wasabi/experiments"
	"github.com/inloco/go-wasabi/experiments/fixtures"
)

const (
	getExperimentBucketsPath           = "/api/v1/experiments/experiment01/buckets"
	getExperimentBucketsResponseStatus = http.StatusOK
)

func handleGetExperimentBuckets(w http.ResponseWriter, req *http.Request) {
	experiment := &experiments.Experiment{
		Buckets: fixtures.Buckets(),
	}

	payload, _ := json.Marshal(experiment)

	w.Write([]byte(payload))
	w.WriteHeader(getExperimentBucketsResponseStatus)
}
