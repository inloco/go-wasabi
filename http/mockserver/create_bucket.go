package mockserver

import (
	"encoding/json"
	"net/http"

	"github.com/inloco/go-wasabi/experiments"
	"github.com/inloco/go-wasabi/experiments/fixtures"
)

const (
	createBucketPath           = "/api/v1/experiments/experiment_id/buckets"
	createBucketResponseStatus = http.StatusOK
)

func handleCreateBucket(w http.ResponseWriter, req *http.Request) {
	bucket := fixtures.Buckets()
	bucket[0].State = experiments.BucketStateOpen

	payload, _ := json.Marshal(bucket[0])

	w.Write([]byte(payload))
	w.WriteHeader(createBucketResponseStatus)
}
