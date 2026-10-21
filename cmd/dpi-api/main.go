package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"detalkipro-intelligence/internal/crypto"
	"detalkipro-intelligence/internal/domain"
	"detalkipro-intelligence/internal/service"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	products := []domain.Product{
		{
			ID:          "tcon-123",
			Name:        "T-Con плата для LG 49UJ630V",
			Category:    "tcon",
			Compatibility: []string{"LG 49UJ630V", "LG 50UJ634V"},
			Analogues:   []string{"tcon-456", "tcon-789"},
		},
		{
			ID:          "power-234",
			Name:        "Плата питания для Samsung UE43NU7100",
			Category:    "power",
			Compatibility: []string{"Samsung UE43NU7100"},
			Analogues:   []string{"power-555"},
		},
	}
	rec := service.NewRecommender(products)

	http.HandleFunc("/api/v1/analyze", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		file, header, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		photoHash := crypto.SHA256(data)

		category := "unknown"
		hint := "unknown"
		name := header.Filename
		if strings.Contains(name, "tcon") {
			category = "tcon"
			hint = strings.ReplaceAll(name, ".jpg", "")
		} else if strings.Contains(name, "power") {
			category = "power"
			hint = strings.ReplaceAll(name, ".jpg", "")
		}

		recommendations := rec.SimpleMatch(category, hint)

		resp := domain.AnalyzeResponse{
			Recommendations: recommendations,
			PhotoHash:       photoHash,
			EventID:         fmt.Sprintf("evt_%s", photoHash[:16]),
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("encode error: %v", err)
		}
	})

	log.Printf("DPI API listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
