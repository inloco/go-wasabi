package experiments

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type EncodingTestSuite struct {
	suite.Suite
}

// Tests the decoding of experiments with all fields filled.
// This payload is seen on /experiments index endpoint.
func (suite *EncodingTestSuite) TestExperiment() {

	payload := `
		{
			"id": "5f04d5e5-a39b-4c02-ac44-10131c3dea06",
			"label": "Develop",
			"applicationName": "develop",
			"startTime": "2018-12-17T03:00:00Z",
			"endTime": "2018-12-31T03:00:00Z",
			"samplingPercent": 1.0,
			"description": "I just need an application to start developing my lil' client.",
			"hypothesisIsCorrect": "yes",
			"results": "Experiment Results",
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
			"tags": [
				"tag_00"
			],
			"buckets": [
				{
					"label": "ControlGroup",
					"experimentID": "5f04d5e5-a39b-4c02-ac44-10131c3dea06",
					"allocationPercent": 0.2,
					"description": "Controle",
					"payload": "bucket=control",
					"state": "OPEN",
					"isControl": true
				},
				{
					"label": "ExperimentGroup",
					"experimentID": "5f04d5e5-a39b-4c02-ac44-10131c3dea06",
					"allocationPercent": 0.8,
					"description": "Experiment Group",
					"payload": "group=experiment",
					"state": "OPEN",
					"isControl": false
				}
			],
			"exclusionIdList": [
				"7a2e8071-712e-463d-b26a-fab086b50603"
			],
			"experimentPageList": [
				{
					"name": "www.somepage.com",
					"allowNewAssignment": true
				}
			],
			"priority": 1
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
			HypothesisIsCorrect:      "yes",
			Results:                  "Experiment Results",
			Rule:                     "",
			RuleJSON:                 "",
			CreationTime:             &expectedCreationTime,
			ModificationTime:         &expectedModificationTime,
			State:                    ExperimentStateRunning,
			IsPersonalizationEnabled: false,
			ModelName:                "",
			ModelVersion:             "",
			IsRapidExperiment:        true,
			UserCap:                  80000,
			CreatorID:                "admin",
			Tags:                     []string{"tag_00"},
			Buckets: []*Bucket{
				&Bucket{
					Label:             "ControlGroup",
					ExperimentID:      "5f04d5e5-a39b-4c02-ac44-10131c3dea06",
					AllocationPercent: 0.2,
					Description:       "Controle",
					Payload:           "bucket=control",
					State:             BucketStateOpen,
					IsControl:         true,
				},
				&Bucket{
					Label:             "ExperimentGroup",
					ExperimentID:      "5f04d5e5-a39b-4c02-ac44-10131c3dea06",
					AllocationPercent: 0.8,
					Description:       "Experiment Group",
					Payload:           "group=experiment",
					State:             BucketStateOpen,
					IsControl:         false,
				},
			},
			ExclusionIDList: []string{"7a2e8071-712e-463d-b26a-fab086b50603"},
			ExperimentPageList: []*ExperimentPage{
				&ExperimentPage{
					Name:               "www.somepage.com",
					AllowNewAssignment: true,
				},
			},
			Priority: 1.0,
		}

		suite.EqualValues(
			expected,
			experiment,
		)
	}

}

// Tests the decoding of experiments with just the first level fields filled.
// This payload is seen on /experiments/:id get endpoint.
func (suite *EncodingTestSuite) TestShallowExperiment() {

	payload := `
		{
			"id": "5f04d5e5-a39b-4c02-ac44-10131c3dea06",
			"label": "Develop",
			"applicationName": "develop",
			"startTime": "2018-12-17T03:00:00Z",
			"endTime": "2018-12-31T03:00:00Z",
			"samplingPercent": 1.0,
			"description": "I just need an application to start developing my lil' client.",
			"hypothesisIsCorrect": "yes",
			"results": "Experiment Results",
			"rule": "",
			"ruleJson": "",
			"creationTime": "2018-12-17T20:24:22Z",
			"modificationTime": "2018-12-17T22:11:55Z",			
			"state": "PAUSED",
			"isPersonalizationEnabled": false,
			"modelName": "",
			"modelVersion": "",
			"isRapidExperiment": true,
			"userCap": 80000,
			"creatorID": "admin",
			"tags": [
				"tag_00"
			],
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
			HypothesisIsCorrect:      "yes",
			Results:                  "Experiment Results",
			Rule:                     "",
			RuleJSON:                 "",
			CreationTime:             &expectedCreationTime,
			ModificationTime:         &expectedModificationTime,
			State:                    ExperimentStatePaused,
			IsPersonalizationEnabled: false,
			ModelName:                "",
			ModelVersion:             "",
			IsRapidExperiment:        true,
			UserCap:                  80000,
			CreatorID:                "admin",
			Tags:                     []string{"tag_00"},
			Buckets:                  []*Bucket{},
			ExclusionIDList:          nil,
			ExperimentPageList:       nil,
			Priority:                 0,
		}

		suite.EqualValues(
			expected,
			experiment,
		)
	}

}

func (suite *EncodingTestSuite) TestBuckets() {
	payload := `{
		"buckets": [{
			"label": "control",
			"experimentID": "exp-id",
			"allocationPercent": 0.4,
			"description": "description1",
			"payload": "payload1",
			"state": "OPEN",
			"isControl": true
		},
		{
			"label": "treatment",
			"experimentID": "exp-id",
			"allocationPercent": 0.6,
			"description": "description2",
			"payload": "payload2",
			"state": "OPEN",
			"isControl": false
		}]
	}
	`

	aux := &Experiment{}
	err := json.Unmarshal([]byte(payload), aux)

	if suite.NoError(err) {
		expectedControlBucket := &Bucket{
			Label:             "control",
			ExperimentID:      "exp-id",
			AllocationPercent: 0.4,
			Description:       "description1",
			Payload:           "payload1",
			State:             "OPEN",
			IsControl:         true,
		}

		expectedTreatmentBucket := &Bucket{
			Label:             "treatment",
			ExperimentID:      "exp-id",
			AllocationPercent: 0.6,
			Description:       "description2",
			Payload:           "payload2",
			State:             "OPEN",
			IsControl:         false,
		}

		suite.Contains(
			aux.Buckets,
			expectedControlBucket,
		)

		suite.Contains(
			aux.Buckets,
			expectedTreatmentBucket,
		)
	}
}

func TestJsonTestSuite(t *testing.T) {
	suite.Run(t, new(EncodingTestSuite))
}
