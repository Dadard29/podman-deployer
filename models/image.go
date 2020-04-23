package models

import "time"

type Image struct {
	ID       string        `json:"id"`
	Names    []string      `json:"names"`
	Digest   string        `json:"digest"`
	Digests  []string      `json:"digests"`
	CreatedAt  time.Time     `json:"createdAt"`
	Size     int           `json:"size"`
	Readonly bool          `json:"readonly"`
	History  []interface{} `json:"history"`
}
