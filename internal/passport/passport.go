package passport

import (
	"encoding/json"
	"time"
)

type Event struct {
	Event     string    `json:"event"`
	Location  string    `json:"location"`
	Engineer  string    `json:"engineer"`
	Timestamp time.Time `json:"timestamp"`
	Evidence  string    `json:"evidence"`
}

type PartPassport struct {
	PartID       string            `json:"part_id"`
	OriginalHash string            `json:"original_hash"`
	Spec         map[string]string `json:"specifications"`
	History      []Event           `json:"history"`
	Signature    string            `json:"signature"`
}
