package models

import "time"

type Image struct {
	ID       string        `json:"id"`
	Names    []string      `json:"names"`
	Digest   string        `json:"digest"`
	Digests  []string      `json:"digests"`
	Created  time.Time     `json:"created"`
	Size     int           `json:"size"`
	Readonly bool          `json:"readonly"`
	History  []interface{} `json:"history"`
}
