package test

import (
	"net/http"

	"github.com/inloco/go-wasabi/assignments"
)

const (
	generateAssignmentPath = "/api/v1/assignments/applications/test/experiments/JustToTestGenerateAssignment/users/userID1"

	generateAssignemntResponseStatus  = http.StatusOK
	generateAssignmentResponsePayload = `{
		"cache": true,
		"payload": "behavior=likethis",
		"assignment": "Experiment",
		"context": "PROD",
		"status": "NEW_ASSIGNMENT"
	}`
)

func handleGenerateAssignemnt(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte(generateAssignmentResponsePayload))
	w.WriteHeader(generateAssignemntResponseStatus)

}

var GenerateAssignmentResponseResource = &assignments.Assignment{
	Cache:      true,
	Payload:    "behavior=likethis",
	Status:     assignments.StatusNewAssignment,
	Assignment: "Experiment",
	Context:    "PROD",
}
