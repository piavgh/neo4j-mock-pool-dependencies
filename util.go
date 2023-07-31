package main

import (
	"math/rand"
	"time"
)

func pickRandomDependency(dependencies []*Contract) *Contract {
	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(len(dependencies))
	pick := dependencies[randomIndex]

	return pick
}
