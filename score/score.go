package score

import "github.com/GeoinovaDev/lower-resultys/collection/array"

// Score struct
type Score struct {
	items []interface{}
}

// New create score
func New() *Score {
	return &Score{}
}

// Load carrega array de itens do tipo string
func (s *Score) Load(items []interface{}) *Score {
	s.items = items

	return s
}

// Add adiciona um item
func (s *Score) Add(item interface{}) *Score {
	s.items = append(s.items, item)

	return s
}

// TopLimit retorna os items ordenados por limite
func (s *Score) TopLimit(limit int) []interface{} {
	clones := array.Cut(s.items, limit)

	return processTop(clones)
}

// Top retorna os items ordenados sem limite
// Return array string
func (s *Score) Top() []interface{} {
	return processTop(s.items)
}

func processTop(items []interface{}) []interface{} {
	score := make(map[interface{}]int)

	for i := 0; i < len(items); i++ {
		domain := items[i]
		if _, ok := score[domain]; ok {
			score[domain]++
		} else {
			score[domain] = 0
		}
	}

	p := make([][]interface{}, len(items))
	for k, v := range score {
		p[v] = append(p[v], k)
	}

	r := make([]interface{}, 0)
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p[i]); j++ {
			r = append(r, p[i][j])
		}
	}

	return array.Reverse(r)
}
