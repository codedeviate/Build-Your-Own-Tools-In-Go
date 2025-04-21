// utils/random.go
package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomChoice returns a random element from a slice.
func RandomChoice(choices []string) string {
	return choices[rand.Intn(len(choices))]
}
