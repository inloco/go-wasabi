package assignments

type Assignment struct {
	Cache bool `json:"cache"`

	Payload    string `json:"payload"`
	Status     Status `json:"status"`
	Assignment string `json:"assignment"`
	Context    string `json:"context"`
}
