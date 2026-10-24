package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/paderinpavel/detalkipro-intelligence/internal/passport"
)

func main() {
	p := passport.PartPassport{
		PartID:       "DPI-001",
		OriginalHash: "placeholder-hash-for-now",
		Spec: map[string]string{
			"type":  "LED_driver",
			"model": "XYZ-100",
		},
		History: []passport.Event{
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

	data, _ := json.MarshalIndent(p, "", "  ")
	fmt.Println(string(data))

	err := os.WriteFile("passport.json", data, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error writing passport.json: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✅ passport.json created")
}
