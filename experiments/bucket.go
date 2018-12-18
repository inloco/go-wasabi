package experiments

type Bucket struct {
	Label             string  `json:"label"`
	ExperimentID      string  `json:"experimentID"`
	AllocationPercent float64 `json:"allocationPercent"`
	Description       string  `json:"description"`
	Payload           string  `json:"payload"`
	State             string  `json:"state"`
	IsControl         bool    `json:"isControl"`
}
