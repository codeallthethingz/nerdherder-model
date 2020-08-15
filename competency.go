package model

type Competency struct {
	ID           string   `json:"_id,omitempty"`
	Path         string   `json:"path"`
	Title        string   `json:"title"`
	Levels       []*Level `json:"levels"`
	Hash         string   `json:"hash"`
	Organization string   `json:"organization"`
}

type Level struct {
	Prove   string `json:"prove"`
	Improve string `json:"improve"`
	Summary string `json:"summary"`
	Status  string `json:"status"`
}
