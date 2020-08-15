package model

type User struct {
	Login       string `json:"login"`
	NodeID      string `json:"node_id"`
	AvatarURL   string `json:"avatar_url"`
	Name        string `json:"name"`
	Token       string `json:"token"`
	AccessLevel string `json:"accessLevel"`
}
