package model

import "math/rand"

type SearchResult struct {
	Hits     int64
	Start    int
	Query    string
	PrevFrom int
	NextFrom int
	Items    []interface{}
}

func getCurrentListener(n, c int) int {
	if n == 0 {
		return (rand.Intn(9) + 1) * 100
	}
	return n * 100
}
