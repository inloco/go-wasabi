package http

import (
	"fmt"
)

const (
	generateAssignmentPathFormat = "/api/v1/assignments/applications/%s/experiments/%s/users/%s"

	createExperimentPathFormat  = "/api/v1/experiments"
	createBucketPathFormat      = "/api/v1/experiments/%s/buckets"
	getExperimentsPathFormat    = "/api/v1/experiments"
	getExperimentByIDPathFormat = "/api/v1/experiments/%s"
)

func generateAssignmentPath(applicationName, experimentLabel, userID string) string {
	return fmt.Sprintf(
		generateAssignmentPathFormat,
		applicationName,
		experimentLabel,
		userID,
	)
}

func createExperimentPath() string {
	return createExperimentPathFormat
}

func createBucketPath(experimentID string) string {
	return fmt.Sprintf(createBucketPathFormat, experimentID)
}

func getExperimentsPath(applicationName, experimentLabel, userID string) string {
	return getExperimentsPathFormat
}

func getExperimentByIDPath(experimentID string) string {
	return fmt.Sprintf(
		getExperimentByIDPathFormat,
		experimentID,
	)
}
