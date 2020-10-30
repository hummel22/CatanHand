package game

import "Catan/objects"

func Score(goods []objects.Good) uint64 {
	score := uint64(0)
	for _, good := range goods {
		score += objects.Values[good]
	}
	return score
}
