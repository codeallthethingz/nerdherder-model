package model

type User struct {
	Login                string                         `json:"login"`
	NodeID               string                         `json:"node_id"`
	AvatarURL            string                         `json:"avatar_url"`
	Name                 string                         `json:"name"`
	Token                string                         `json:"token,omitempty"`
	AccessLevel          string                         `json:"accessLevel"`
	CompetenciesInMotion map[string]*CompetencyInMotion `json:"competenciesInMotion"`
}

type CompetencyInMotion struct {
	CompetencyID string            `json:"competencyId"`
	Status       map[string]string `json:"status"`
}
