package experiments

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type JsonTestSuite struct {
	suite.Suite
}

func (suite *JsonTestSuite) TestExperiment() {

	payload := `
		{
			"id": "5f04d5e5-a39b-4c02-ac44-10131c3dea06",
			"label": "Develop",
			"applicationName": "develop",
			"startTime": "2018-12-17T03:00:00Z",
			"endTime": "2018-12-31T03:00:00Z",
			"samplingPercent": 1.0,
			"description": "I just need an application to start developing my lil' client.",
			"hypothesisIsCorrect": "",
			"results": "",
			"rule": "",
			"ruleJson": "",
			"creationTime": "2018-12-17T20:24:22Z",
			"modificationTime": "2018-12-17T22:11:55Z",
			"state": "RUNNING",
			"isPersonalizationEnabled": false,
			"modelName": "",
			"modelVersion": "",
			"isRapidExperiment": true,
			"userCap": 80000,
			"creatorID": "admin",
			"tags": null,
			"buckets": [],
			"exclusionIdList": null,
			"experimentPageList": null,
			"priority": null
		}
	`

	experiment := &Experiment{}

	err := json.Unmarshal([]byte(payload), experiment)
	if suite.NoError(err) {

		expectedStartTime, _ := time.Parse(time.RFC3339, "2018-12-17T03:00:00Z")
		expectedEndTime, _ := time.Parse(time.RFC3339, "2018-12-31T03:00:00Z")
		expectedCreationTime, _ := time.Parse(time.RFC3339, "2018-12-17T20:24:22Z")
		expectedModificationTime, _ := time.Parse(time.RFC3339, "2018-12-17T22:11:55Z")

		expected := &Experiment{
			ID:                       "5f04d5e5-a39b-4c02-ac44-10131c3dea06",
			Label:                    "Develop",
			ApplicationName:          "develop",
			StartTime:                &expectedStartTime,
			EndTime:                  &expectedEndTime,
			SamplingPercent:          1.0,
			Description:              "I just need an application to start developing my lil' client.",
			HypothesisIsCorrect:      "",
			Results:                  "",
			Rule:                     "",
			RuleJSON:                 "",
			CreationTime:             &expectedCreationTime,
			ModificationTime:         &expectedModificationTime,
			State:                    "RUNNING",
			IsPersonalizationEnabled: false,
			ModelName:                "",
			ModelVersion:             "",
			IsRapidExperiment:        true,
			UserCap:                  80000,
			CreatorID:                "admin",
			Tags:                     nil,
			Buckets:                  []*Bucket{},
			ExclusionIDList:          nil,
			ExperimentPageList:       nil,
			Priority:                 0.0,
		}

		suite.EqualValues(
			expected,
			experiment,
		)
	}

}

func TestJsonTestSuite(t *testing.T) {
	suite.Run(t, new(JsonTestSuite))
}
