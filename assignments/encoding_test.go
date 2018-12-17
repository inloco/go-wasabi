package assignments

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/suite"
)

type JsonTestSuite struct {
	suite.Suite
}

func (suite *JsonTestSuite) TestNewAssignment() {
	payload := `
		{
			"cache": true,
			"payload": "group=experiment",
			"assignment": "ExperimentGroup",
			"context": "PROD",
			"status": "NEW_ASSIGNMENT"
		}
	`

	assignment := &Assignment{}

	err := json.Unmarshal([]byte(payload), assignment)
	if suite.NoError(err) {

		expected := &Assignment{
			Cache:      true,
			Payload:    "group=experiment",
			Assignment: "ExperimentGroup",
			Context:    "PROD",
			Status:     StatusNewAssignment,
		}

		suite.EqualValues(
			expected,
			assignment,
		)
	}
}

func (suite *JsonTestSuite) TestExistingAssignment() {
	payload := `
		{
			"cache": true,
			"payload": "group=experiment",
			"assignment": "ExperimentGroup",
			"context": "PROD",
			"status": "EXISTING_ASSIGNMENT"
		}
	`

	assignment := &Assignment{}

	err := json.Unmarshal([]byte(payload), assignment)
	if suite.NoError(err) {

		expected := &Assignment{
			Cache:      true,
			Payload:    "group=experiment",
			Assignment: "ExperimentGroup",
			Context:    "PROD",
			Status:     StatusExistingAssignment,
		}

		suite.EqualValues(
			expected,
			assignment,
		)
	}
}

func (suite *JsonTestSuite) TestInDraftExperiment() {
	payload := `
		{
			"cache": false,
			"status": "EXPERIMENT_IN_DRAFT_STATE"
		}
	`

	assignment := &Assignment{}

	err := json.Unmarshal([]byte(payload), assignment)
	if suite.NoError(err) {

		expected := &Assignment{
			Cache:  false,
			Status: StatusExperimentInDraft,
		}

		suite.EqualValues(
			expected,
			assignment,
		)
	}
}

func TestJsonTestSuite(t *testing.T) {
	suite.Run(t, new(JsonTestSuite))
}
