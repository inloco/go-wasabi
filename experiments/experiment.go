package experiments

import (
	"time"
)

type Experiment struct {
	ID                       string            `json:"id"`
	Label                    string            `json:"label"`
	ApplicationName          string            `json:"applicationName"`
	StartTime                *time.Time        `json:"startTime"`
	EndTime                  *time.Time        `json:"endTime"`
	SamplingPercent          float64           `json:"samplingPercent"`
	Description              string            `json:"description"`
	HypothesisIsCorrect      string            `json:"hypothesisIsCorrect"`
	Results                  string            `json:"results"`
	Rule                     string            `json:"rule"`
	RuleJSON                 string            `json:"ruleJson"`
	CreationTime             *time.Time        `json:"creationTime"`
	ModificationTime         *time.Time        `json:"modificationTime"`
	State                    State             `json:"state"`
	IsPersonalizationEnabled bool              `json:"isPersonalizationEnabled"`
	ModelName                string            `json:"modelName"`
	ModelVersion             string            `json:"modelVersion"`
	IsRapidExperiment        bool              `json:"isRapidExperiment"`
	UserCap                  float64           `json:"userCap"`
	CreatorID                string            `json:"creatorID"`
	Tags                     []string          `json:"tags"`
	Buckets                  []*Bucket         `json:"buckets"`
	ExclusionIDList          []*ExclusionID    `json:"exclusionIdList"`
	ExperimentPageList       []*ExperimentPage `json:"experimentPageList"`
	Priority                 float64           `json:"priority"`
}

type ExperimentPage struct {
	Name               string `json:"name"`
	AllowNewAssignment bool   `json:"allowNewAssignment"`
}

type Bucket struct {
	Label             string  `json:"label"`
	ExperimentID      string  `json:"experimentID"`
	AllocationPercent float64 `json:"allocationPercent"`
	Description       string  `json:"description"`
	Payload           string  `json:"payload"`
	State             string  `json:"state"`
	IsControl         bool    `json:"isControl"`
}

type ExclusionID struct {
	RawID string `json:"rawID"`
}
