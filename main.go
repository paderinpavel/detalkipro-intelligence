package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Импортируем типы прямо из файла passport.go (они в той же папке)
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

func main() {
	p := PartPassport{
		PartID:       "DPI-001",
		OriginalHash: "",
		Spec: map[string]string{
			"type":  "LED_driver",
			"model": "XYZ-100",
		},
		History: []Event{
			{
				Event:     "received",
				Location:  "PVS-05",
				Engineer:  "Ivanov",
				Timestamp: time.Now(),
				Evidence:  "photo_001.jpg",
			},
		},
		Signature: "engineer_signature_placeholder",
	}

	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error marshaling: %v\n", err)
		os.Exit(1)
	}

	hash := sha256.Sum256(data)
	p.OriginalHash = hex.EncodeToString(hash[:])

	finalData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error marshaling final: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(finalData))

	err = os.WriteFile("passport.json", finalData, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error writing passport.json: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✅ passport.json created with real SHA-256 hash")
}
