package http

import (
	"fmt"
)

const (
	generateAssignmentPathFormat = "/v1/assignments/applications/%s/experiments/%s/users/%s"

	createExperimentPathFormat  = "/v1/experiments"
	getExperimentsPathFormat    = "/v1/experiments"
	getExperimentByIDPathFormat = "/v1/experiments/%s"
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

func getExperimentsPath(applicationName, experimentLabel, userID string) string {
	return getExperimentsPathFormat
}

func getExperimentByIDPath(experimentID string) string {
	return fmt.Sprintf(
		getExperimentByIDPathFormat,
		experimentID,
	)
}
