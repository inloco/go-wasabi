package experiments

import (
	"time"
)

type Experiment struct {
	ID                       string            `json:"id,omitempty"`
	Label                    string            `json:"label,omitempty"`
	ApplicationName          string            `json:"applicationName,omitempty"`
	StartTime                *time.Time        `json:"startTime,omitempty"`
	EndTime                  *time.Time        `json:"endTime,omitempty"`
	SamplingPercent          float64           `json:"samplingPercent,omitempty"`
	Description              string            `json:"description,omitempty"`
	HypothesisIsCorrect      string            `json:"hypothesisIsCorrect,omitempty"`
	Results                  string            `json:"results,omitempty"`
	Rule                     string            `json:"rule,omitempty"`
	RuleJSON                 string            `json:"ruleJson,omitempty"`
	CreationTime             *time.Time        `json:"creationTime,omitempty"`
	ModificationTime         *time.Time        `json:"modificationTime,omitempty"`
	State                    ExperimentState   `json:"state,omitempty"`
	IsPersonalizationEnabled bool              `json:"isPersonalizationEnabled,omitempty"`
	ModelName                string            `json:"modelName,omitempty"`
	ModelVersion             string            `json:"modelVersion,omitempty"`
	IsRapidExperiment        bool              `json:"isRapidExperiment,omitempty"`
	UserCap                  float64           `json:"userCap,omitempty"`
	CreatorID                string            `json:"creatorID,omitempty"`
	Tags                     []string          `json:"tags,omitempty"`
	Buckets                  []*Bucket         `json:"buckets,omitempty"`
	ExclusionIDList          []string          `json:"exclusionIdList,omitempty"`
	ExperimentPageList       []*ExperimentPage `json:"experimentPageList,omitempty"`
	Priority                 float64           `json:"priority,omitempty"`
}

const (
	ExperimentStateDraft      = "DRAFT"
	ExperimentStateRunning    = "RUNNING"
	ExperimentStatePaused     = "PAUSED"
	ExperimentStateTerminated = "TERMINATED"
	ExperimentStateDeleted    = "DELETED"
)

type ExperimentState string
