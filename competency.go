package model

import (
	"crypto/md5"
	"fmt"
)

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

func (c *Competency) ComputeHash() {
	hashableContents := ""
	hashableContents += c.Title + "\n"
	for _, level := range c.Levels {
		hashableContents += level.Prove
		hashableContents += level.Improve
		hashableContents += level.Summary
		hashableContents += "\n"
	}
	c.Hash = fmt.Sprintf("%x", md5.Sum([]byte(hashableContents)))
}
