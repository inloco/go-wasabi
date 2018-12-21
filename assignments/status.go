package assignments

const (
	StatusNewAssignment      = "NEW_ASSIGNMENT"
	StatusExistingAssignment = "EXISTING_ASSIGNMENT"

	StatusExperimentInDraft = "EXPERIMENT_IN_DRAFT_STATE"
	StatusExperimentPaused  = "EXPERIMENT_PAUSED"
)

type Status string
