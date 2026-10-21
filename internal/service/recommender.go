package service

import (
	"detalkipro-intelligence/internal/domain"
	"strings"
)

type Recommender struct {
	products []domain.Product
}

func NewRecommender(products []domain.Product) *Recommender {
	return &Recommender{products: products}
}

func (r *Recommender) SimpleMatch(category string, compatibilityHint string) []domain.Recommendation {
	var recs []domain.Recommendation
	for _, p := range r.products {
		if strings.EqualFold(p.Category, category) {
			score := 1.0
			for _, c := range p.Compatibility {
				if strings.Contains(strings.ToLower(c), strings.ToLower(compatibilityHint)) {
					score = 0.95
				}
			}
			recs = append(recs, domain.Recommendation{
				ProductID:     p.ID,
				Name:          p.Name,
				Compatibility: p.Compatibility,
				Analogues:     p.Analogues,
				Score:         score,
			})
		}
	}
	return recs
}
