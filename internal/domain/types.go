package domain

type Product struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Category      string   `json:"category"`
	Compatibility []string `json:"compatibility"`
	Analogues     []string `json:"analogues"`
}

type Recommendation struct {
	ProductID     string   `json:"product_id"`
	Name          string   `json:"name"`
	Compatibility []string `json:"compatibility"`
	Analogues     []string `json:"analogues"`
	Score         float64  `json:"score"`
}

type AnalyzeResponse struct {
	Recommendations []Recommendation `json:"recommendations"`
	PhotoHash       string           `json:"photo_hash"`
	EventID         string           `json:"event_id"`
}
