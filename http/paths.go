package http

import (
	"fmt"
)

const (
	generateAssignmentPathFormat = "/api/v1/assignments/applications/%s/experiments/%s/users/%s"

	createExperimentPathFormat     = "/api/v1/experiments"
	updateExperimentPathFormat     = "/api/v1/experiments/%s"
	createBucketPathFormat         = "/api/v1/experiments/%s/buckets"
	getExperimentsPathFormat       = "/api/v1/experiments"
	getExperimentByIDPathFormat    = "/api/v1/experiments/%s"
	getExperimentBucketsPathFormat = "/api/v1/experiments/%s/buckets"
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

func updateExperimentPath(experimentID string) string {
	return fmt.Sprintf(updateExperimentPathFormat, experimentID)
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

func getExperimentBucketsPath(experimentID string) string {
	return fmt.Sprintf(
		getExperimentBucketsPathFormat,
		experimentID,
	)
}
